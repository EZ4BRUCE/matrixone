package db

import (
	"context"
	"fmt"
	"math/rand"
	"matrixone/pkg/vm/engine/aoe"
	e "matrixone/pkg/vm/engine/aoe/storage"
	"matrixone/pkg/vm/engine/aoe/storage/dbi"
	md "matrixone/pkg/vm/engine/aoe/storage/metadata"
	"matrixone/pkg/vm/engine/aoe/storage/mock/type/chunk"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/panjf2000/ants/v2"

	"github.com/stretchr/testify/assert"
)

var (
	TEST_DB_DIR = "/tmp/db_test"
)

func initDBTest() {
	os.RemoveAll(TEST_DB_DIR)
}

func initDB() *DB {
	rand.Seed(time.Now().UnixNano())
	cfg := &md.Configuration{
		Dir:              TEST_DB_DIR,
		SegmentMaxBlocks: 2,
		BlockMaxRows:     20000,
	}
	opts := &e.Options{}
	opts.Meta.Conf = cfg
	inst, _ := Open(TEST_DB_DIR, opts)
	return inst
}

func TestCreateTable(t *testing.T) {
	initDBTest()
	inst := initDB()
	assert.NotNil(t, inst)
	defer inst.Close()
	tblCnt := rand.Intn(5) + 3
	prefix := "mocktbl_"
	var wg sync.WaitGroup
	names := make([]string, 0)
	for i := 0; i < tblCnt; i++ {
		tableInfo := md.MockTableInfo(2)
		name := fmt.Sprintf("%s%d", prefix, i)
		tablet := aoe.TabletInfo{Table: *tableInfo, Name: name}
		names = append(names, name)
		wg.Add(1)
		go func(w *sync.WaitGroup) {
			_, err := inst.CreateTable(&tablet)
			assert.Nil(t, err)
			w.Done()
		}(&wg)
	}
	wg.Wait()
	ids, err := inst.TableIDs()
	assert.Nil(t, err)
	assert.Equal(t, tblCnt, len(ids))
	for _, name := range names {
		assert.True(t, inst.HasTable(name))
	}
	t.Log(inst.store.MetaInfo.String())
}

func TestCreateDuplicateTable(t *testing.T) {
	initDBTest()
	inst := initDB()
	defer inst.Close()

	tableInfo := md.MockTableInfo(2)
	tablet := aoe.TabletInfo{Table: *tableInfo, Name: "t1"}
	_, err := inst.CreateTable(&tablet)
	assert.Nil(t, err)
	_, err = inst.CreateTable(&tablet)
	assert.NotNil(t, err)
}

func TestAppend(t *testing.T) {
	initDBTest()
	inst := initDB()
	tableInfo := md.MockTableInfo(2)
	tablet := aoe.TabletInfo{Table: *tableInfo, Name: "mocktbl"}
	tid, err := inst.CreateTable(&tablet)
	assert.Nil(t, err)
	tblMeta, err := inst.Opts.Meta.Info.ReferenceTable(tid)
	assert.Nil(t, err)
	blkCnt := 2
	rows := inst.store.MetaInfo.Conf.BlockMaxRows * uint64(blkCnt)
	ck := chunk.MockChunk(tblMeta.Schema.Types(), rows)
	assert.Equal(t, uint64(rows), ck.GetCount())
	logIdx := &md.LogIndex{
		ID:       uint64(0),
		Capacity: ck.GetCount(),
	}
	invalidName := "xxx"
	err = inst.Append(invalidName, ck, logIdx)
	assert.NotNil(t, err)
	insertCnt := 4
	for i := 0; i < insertCnt; i++ {
		err = inst.Append(tableInfo.Name, ck, logIdx)
		assert.Nil(t, err)
		// tbl, err := inst.store.DataTables.GetTable(tid)
		// assert.Nil(t, err)
		// t.Log(tbl.GetCollumn(0).ToString(1000))
	}

	cols := []int{0, 1}
	tbl, _ := inst.store.DataTables.GetTable(tid)
	segIds := tbl.SegmentIds()
	ssCtx := &dbi.GetSnapshotCtx{
		TableName:  tableInfo.Name,
		SegmentIds: segIds,
		Cols:       cols,
	}

	blkCount := 0
	segCount := 0
	ss, err := inst.GetSnapshot(ssCtx)
	assert.Nil(t, err)
	segIt := ss.NewIt()
	assert.NotNil(t, segIt)

	for segIt.Valid() {
		segCount++
		segH := segIt.GetHandle()
		assert.NotNil(t, segH)
		blkIt := segH.NewIt()
		// segH.Close()
		assert.NotNil(t, blkIt)
		for blkIt.Valid() {
			blkCount++
			blkIt.Next()
		}
		blkIt.Close()
		segIt.Next()
	}
	segIt.Close()
	assert.Equal(t, insertCnt, segCount)
	assert.Equal(t, blkCnt*insertCnt, blkCount)
	ss.Close()

	time.Sleep(time.Duration(20) * time.Millisecond)
	t.Log(inst.FsMgr.String())
	t.Log(inst.MTBufMgr.String())
	t.Log(inst.SSTBufMgr.String())
	t.Log(inst.IndexBufMgr.String())
	// t.Log(tbl.GetIndexHolder().String())
	inst.Close()
}

type InsertReq struct {
	Name     string
	Data     *chunk.Chunk
	LogIndex *md.LogIndex
}

// TODO: When the capacity is not very big and the query concurrency is very high,
// the db will be stuck due to no more space. Need intruduce timeout mechanism later
func TestConcurrency(t *testing.T) {
	initDBTest()
	inst := initDB()
	tableInfo := md.MockTableInfo(2)
	tablet := aoe.TabletInfo{Table: *tableInfo, Name: "mockcon"}
	tid, err := inst.CreateTable(&tablet)
	assert.Nil(t, err)
	tblMeta, err := inst.Opts.Meta.Info.ReferenceTable(tid)
	assert.Nil(t, err)
	blkCnt := inst.store.MetaInfo.Conf.SegmentMaxBlocks
	rows := inst.store.MetaInfo.Conf.BlockMaxRows * blkCnt
	baseCk := chunk.MockChunk(tblMeta.Schema.Types(), rows)
	insertCh := make(chan *InsertReq)
	searchCh := make(chan *dbi.GetSnapshotCtx)

	p, _ := ants.NewPool(40)

	reqCtx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	var searchWg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case req := <-searchCh:
				f := func() {
					defer searchWg.Done()
					{
						ss, err := inst.GetSnapshot(req)
						assert.Nil(t, err)
						segIt := ss.NewIt()
						assert.Nil(t, err)
						if segIt == nil {
							return
						}
						segCnt := 0
						blkCnt := 0
						for segIt.Valid() {
							segCnt++
							sh := segIt.GetHandle()
							blkIt := sh.NewIt()
							for blkIt.Valid() {
								blkCnt++
								blkHandle := blkIt.GetHandle()
								hh := blkHandle.Prefetch()
								hh.Close()
								// blkHandle.Close()
								blkIt.Next()
							}
							blkIt.Close()
							segIt.Next()
						}
						segIt.Close()
						ss.Close()
					}
				}
				p.Submit(f)
			case req := <-insertCh:
				wg.Add(1)
				go func() {
					err := inst.Append(req.Name, req.Data, req.LogIndex)
					assert.Nil(t, err)
					wg.Done()
				}()
			}
		}
	}(reqCtx)

	insertCnt := 8
	var wg2 sync.WaitGroup

	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for i := 0; i < insertCnt; i++ {
			insertReq := &InsertReq{
				Name:     tablet.Name,
				Data:     baseCk,
				LogIndex: &md.LogIndex{ID: uint64(i), Capacity: baseCk.GetCount()},
			}
			insertCh <- insertReq
		}
	}()

	cols := make([]int, 0)
	for i := 0; i < len(tblMeta.Schema.ColDefs); i++ {
		cols = append(cols, i)
	}
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		reqCnt := 20000
		for i := 0; i < reqCnt; i++ {
			tbl, _ := inst.store.DataTables.GetTable(tid)
			for tbl == nil {
				time.Sleep(time.Duration(100) * time.Microsecond)
				tbl, _ = inst.store.DataTables.GetTable(tid)
			}
			segIds := tbl.SegmentIds()
			searchReq := &dbi.GetSnapshotCtx{
				TableName:  tablet.Name,
				SegmentIds: segIds,
				Cols:       cols,
			}
			searchWg.Add(1)
			searchCh <- searchReq
		}
	}()

	wg2.Wait()
	searchWg.Wait()
	cancel()
	wg.Wait()
	time.Sleep(time.Duration(100) * time.Millisecond)
	tbl, _ := inst.store.DataTables.GetTable(tid)
	root := tbl.WeakRefRoot()
	assert.Equal(t, int64(1), root.RefCount())
	opts := &dbi.GetSnapshotCtx{
		TableName: tablet.Name,
		Cols:      cols,
		ScanAll:   true,
	}
	now := time.Now()
	ss, err := inst.GetSnapshot(opts)
	assert.Nil(t, err)
	segIt := ss.NewIt()
	segCnt := 0
	tblkCnt := 0
	for segIt.Valid() {
		segCnt++
		h := segIt.GetHandle()
		blkIt := h.NewIt()
		for blkIt.Valid() {
			tblkCnt++
			blkHandle := blkIt.GetHandle()
			// col0 := blkHandle.GetColumn(0)
			// ctx := index.NewFilterCtx(index.OpEq)
			// ctx.Val = int32(0 + col0.GetColIdx()*100)
			// err = col0.EvalFilter(ctx)
			// assert.Nil(t, err)
			// if col0.GetBlockType() > base.PERSISTENT_BLK {
			// 	assert.False(t, ctx.BoolRes)
			// }
			// ctx.Reset()
			// ctx.Op = index.OpEq
			// ctx.Val = int32(1 + col0.GetColIdx()*100)
			// err = col0.EvalFilter(ctx)
			// assert.Nil(t, err)
			// if col0.GetBlockType() > base.PERSISTENT_BLK {
			// 	assert.True(t, ctx.BoolRes)
			// }
			hh := blkHandle.Prefetch()
			hh.Close()
			// blkHandle.Close()
			blkIt.Next()
		}
		blkIt.Close()
		// h.Close()
		segIt.Next()
	}
	segIt.Close()
	ss.Close()
	assert.Equal(t, insertCnt*int(blkCnt), tblkCnt)
	assert.Equal(t, insertCnt, segCnt)
	assert.Equal(t, int64(1), root.RefCount())

	t.Logf("Takes %v", time.Since(now))
	t.Log(tbl.String())
	time.Sleep(time.Duration(100) * time.Millisecond)

	t.Log(inst.WorkersStatsString())
	t.Log(inst.MTBufMgr.String())
	t.Log(inst.SSTBufMgr.String())
	// t.Log(inst.IndexBufMgr.String())
	// t.Log(tbl.GetIndexHolder().String())
	inst.Close()
}

func TestGC(t *testing.T) {
	initDBTest()
	inst := initDB()
	tableInfo := md.MockTableInfo(2)
	tablet := aoe.TabletInfo{Table: *tableInfo, Name: "mockcon"}
	tid, err := inst.CreateTable(&tablet)
	assert.Nil(t, err)
	blkCnt := inst.store.MetaInfo.Conf.SegmentMaxBlocks
	rows := inst.store.MetaInfo.Conf.BlockMaxRows * blkCnt
	tblMeta, err := inst.Opts.Meta.Info.ReferenceTable(tid)
	assert.Nil(t, err)
	baseCk := chunk.MockChunk(tblMeta.Schema.Types(), rows)

	logIdx := &md.LogIndex{
		ID:       uint64(0),
		Capacity: baseCk.GetCount(),
	}

	insertCnt := uint64(4)

	var wg sync.WaitGroup
	{
		for i := uint64(0); i < insertCnt; i++ {
			wg.Add(1)
			go func() {
				inst.Append(tablet.Name, baseCk, logIdx)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	time.Sleep(time.Duration(40) * time.Millisecond)
	t.Log(inst.MTBufMgr.String())
	t.Log(inst.SSTBufMgr.String())
	t.Log(inst.IndexBufMgr.String())
	assert.Equal(t, int(blkCnt*insertCnt*2), inst.SSTBufMgr.NodeCount()+inst.MTBufMgr.NodeCount())
	inst.Close()
}
