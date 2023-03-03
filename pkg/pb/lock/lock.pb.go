// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lock.proto

package lock

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// LockTable describes which CN manages a Table's Locks.
type LockTable struct {
	// Table table id
	Table uint64 `protobuf:"varint,1,opt,name=Table,proto3" json:"Table,omitempty"`
	// ServiceID lock service id, a cn node will only run one instance of LockService
	ServiceID string `protobuf:"bytes,2,opt,name=ServiceID,proto3" json:"ServiceID,omitempty"`
	// Version will incr if CN-Table bind changed.
	Version uint64 `protobuf:"varint,3,opt,name=Version,proto3" json:"Version,omitempty"`
	// Valid false if the service is disabled, and no new service bind this table
	Valid                bool     `protobuf:"varint,4,opt,name=Valid,proto3" json:"Valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockTable) Reset()         { *m = LockTable{} }
func (m *LockTable) String() string { return proto.CompactTextString(m) }
func (*LockTable) ProtoMessage()    {}
func (*LockTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_164ad2988c7acaf1, []int{0}
}
func (m *LockTable) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LockTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LockTable.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LockTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockTable.Merge(m, src)
}
func (m *LockTable) XXX_Size() int {
	return m.Size()
}
func (m *LockTable) XXX_DiscardUnknown() {
	xxx_messageInfo_LockTable.DiscardUnknown(m)
}

var xxx_messageInfo_LockTable proto.InternalMessageInfo

func (m *LockTable) GetTable() uint64 {
	if m != nil {
		return m.Table
	}
	return 0
}

func (m *LockTable) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

func (m *LockTable) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *LockTable) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*LockTable)(nil), "lock.LockTable")
}

func init() { proto.RegisterFile("lock.proto", fileDescriptor_164ad2988c7acaf1) }

var fileDescriptor_164ad2988c7acaf1 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xc9, 0x4f, 0xce,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0x74, 0xd3, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0x92, 0x49, 0xa5,
	0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0x34, 0x29, 0xe5, 0x73, 0x71, 0xfa, 0xe4, 0x27, 0x67,
	0x87, 0x24, 0x26, 0xe5, 0xa4, 0x0a, 0x89, 0x70, 0xb1, 0x82, 0x19, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x2c, 0x41, 0x10, 0x8e, 0x90, 0x0c, 0x17, 0x67, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0xa7,
	0x8b, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x42, 0x40, 0x48, 0x82, 0x8b, 0x3d, 0x2c, 0xb5,
	0xa8, 0x38, 0x33, 0x3f, 0x4f, 0x82, 0x19, 0xac, 0x0b, 0xc6, 0x05, 0x99, 0x16, 0x96, 0x98, 0x93,
	0x99, 0x22, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1, 0x11, 0x04, 0xe1, 0x38, 0xd9, 0x5f, 0x78, 0x28, 0xc7,
	0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x46, 0x21, 0xbb,
	0x38, 0x37, 0xb1, 0xa4, 0x28, 0xb3, 0x22, 0xbf, 0x28, 0x33, 0x3d, 0x33, 0x0f, 0xc6, 0xc9, 0x4b,
	0xd5, 0x2f, 0xc8, 0x4e, 0xd7, 0x2f, 0x48, 0xd2, 0x07, 0x79, 0x30, 0x89, 0x0d, 0xec, 0x70, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x65, 0x67, 0xda, 0xfb, 0x00, 0x00, 0x00,
}

func (m *LockTable) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LockTable) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LockTable) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Valid {
		i--
		if m.Valid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Version != 0 {
		i = encodeVarintLock(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ServiceID) > 0 {
		i -= len(m.ServiceID)
		copy(dAtA[i:], m.ServiceID)
		i = encodeVarintLock(dAtA, i, uint64(len(m.ServiceID)))
		i--
		dAtA[i] = 0x12
	}
	if m.Table != 0 {
		i = encodeVarintLock(dAtA, i, uint64(m.Table))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLock(dAtA []byte, offset int, v uint64) int {
	offset -= sovLock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LockTable) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Table != 0 {
		n += 1 + sovLock(uint64(m.Table))
	}
	l = len(m.ServiceID)
	if l > 0 {
		n += 1 + l + sovLock(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovLock(uint64(m.Version))
	}
	if m.Valid {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLock(x uint64) (n int) {
	return sovLock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LockTable) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LockTable: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LockTable: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Table", wireType)
			}
			m.Table = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Table |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Valid = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipLock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLock
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLock = fmt.Errorf("proto: unexpected end of group")
)