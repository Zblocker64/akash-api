// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/market/v1/bid.proto

package v1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// BidState is an enum which refers to state of bid
type BidState int32

const (
	// Prefix should start with 0 in enum. So declaring dummy state
	BidStateInvalid BidState = 0
	// BidOpen denotes state for bid open
	BidOpen BidState = 1
	// BidMatched denotes state for bid open
	BidActive BidState = 2
	// BidLost denotes state for bid lost
	BidLost BidState = 3
	// BidClosed denotes state for bid closed
	BidClosed BidState = 4
)

var BidState_name = map[int32]string{
	0: "bid_invalid",
	1: "bid_open",
	2: "bid_active",
	3: "bid_lost",
	4: "bid_closed",
}

var BidState_value = map[string]int32{
	"bid_invalid": 0,
	"bid_open":    1,
	"bid_active":  2,
	"bid_lost":    3,
	"bid_closed":  4,
}

func (x BidState) String() string {
	return proto.EnumName(BidState_name, int32(x))
}

func (BidState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3938cb3dd8faff6a, []int{0}
}

// BidID stores owner and all other seq numbers
// A successful bid becomes a Lease(ID).
type BidID struct {
	Owner    string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	DSeq     uint64 `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	GSeq     uint32 `protobuf:"varint,3,opt,name=gseq,proto3" json:"gseq" yaml:"gseq"`
	OSeq     uint32 `protobuf:"varint,4,opt,name=oseq,proto3" json:"oseq" yaml:"oseq"`
	Provider string `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider" yaml:"provider"`
}

func (m *BidID) Reset()      { *m = BidID{} }
func (*BidID) ProtoMessage() {}
func (*BidID) Descriptor() ([]byte, []int) {
	return fileDescriptor_3938cb3dd8faff6a, []int{0}
}
func (m *BidID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BidID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BidID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BidID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidID.Merge(m, src)
}
func (m *BidID) XXX_Size() int {
	return m.Size()
}
func (m *BidID) XXX_DiscardUnknown() {
	xxx_messageInfo_BidID.DiscardUnknown(m)
}

var xxx_messageInfo_BidID proto.InternalMessageInfo

func (m *BidID) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *BidID) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *BidID) GetGSeq() uint32 {
	if m != nil {
		return m.GSeq
	}
	return 0
}

func (m *BidID) GetOSeq() uint32 {
	if m != nil {
		return m.OSeq
	}
	return 0
}

func (m *BidID) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func init() {
	proto.RegisterEnum("akash.market.v1.BidState", BidState_name, BidState_value)
	proto.RegisterType((*BidID)(nil), "akash.market.v1.BidID")
}

func init() { proto.RegisterFile("akash/market/v1/bid.proto", fileDescriptor_3938cb3dd8faff6a) }

var fileDescriptor_3938cb3dd8faff6a = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd2, 0xb1, 0x6a, 0xdb, 0x40,
	0x18, 0x07, 0x70, 0x9d, 0x23, 0xb7, 0xce, 0xb9, 0xc1, 0x42, 0xcd, 0x60, 0x0b, 0xaa, 0x13, 0xa6,
	0x83, 0x29, 0x54, 0xc2, 0x64, 0x28, 0xcd, 0x16, 0x35, 0x10, 0x02, 0x85, 0x40, 0xbc, 0x75, 0x09,
	0xb2, 0xef, 0x50, 0x0f, 0xdb, 0xfa, 0x1c, 0xdd, 0xa1, 0xd2, 0x37, 0x28, 0x9e, 0x3a, 0x76, 0x31,
	0x04, 0xba, 0x75, 0xee, 0x43, 0x74, 0x2a, 0xa1, 0x53, 0x27, 0x51, 0xec, 0xa5, 0x78, 0xf4, 0x13,
	0x94, 0xbb, 0xb3, 0x15, 0x3c, 0x74, 0x92, 0xee, 0xff, 0xff, 0x7e, 0x07, 0x07, 0x1f, 0xee, 0x24,
	0xe3, 0x44, 0xbc, 0x8f, 0xa6, 0x49, 0x3e, 0x66, 0x32, 0x2a, 0xfa, 0xd1, 0x90, 0xd3, 0x70, 0x96,
	0x83, 0x04, 0xb7, 0xa5, 0xab, 0xd0, 0x54, 0x61, 0xd1, 0xf7, 0x8e, 0x53, 0x48, 0x41, 0x77, 0x91,
	0xfa, 0x33, 0x63, 0x5e, 0x67, 0x04, 0x62, 0x0a, 0xe2, 0xc6, 0x14, 0xe6, 0x60, 0xaa, 0xee, 0xcf,
	0x1a, 0xae, 0xc7, 0x9c, 0x5e, 0x9e, 0xbb, 0x17, 0xb8, 0x0e, 0x1f, 0x32, 0x96, 0xb7, 0x51, 0x80,
	0x7a, 0x87, 0x71, 0x7f, 0x5d, 0x12, 0x13, 0x6c, 0x4a, 0xf2, 0xe4, 0x63, 0x32, 0x9d, 0x9c, 0x76,
	0xf5, 0xb1, 0xfb, 0xeb, 0xfb, 0xcb, 0xe3, 0xed, 0x1d, 0x67, 0x94, 0xe6, 0x4c, 0x88, 0x81, 0xcc,
	0x79, 0x96, 0x5e, 0x9b, 0x71, 0xf7, 0x04, 0xdb, 0x54, 0xb0, 0xdb, 0x76, 0x2d, 0x40, 0x3d, 0x3b,
	0x26, 0xcb, 0x92, 0xd8, 0xe7, 0x03, 0x76, 0xbb, 0x2e, 0x89, 0xce, 0x37, 0x25, 0x69, 0x9a, 0xeb,
	0xd4, 0xa9, 0x7b, 0xad, 0x43, 0x85, 0x52, 0x85, 0x0e, 0x02, 0xd4, 0x3b, 0x32, 0xe8, 0x62, 0x8b,
	0xd2, 0x3d, 0x94, 0x1a, 0x94, 0x6e, 0x11, 0x28, 0x64, 0x3f, 0xa0, 0xab, 0x2d, 0x82, 0x3d, 0x04,
	0x06, 0xa9, 0x8f, 0x3b, 0xc0, 0x8d, 0x59, 0x0e, 0x05, 0xa7, 0x2c, 0x6f, 0xd7, 0xf5, 0x53, 0x5f,
	0xad, 0x4b, 0x52, 0x65, 0x9b, 0x92, 0xb4, 0x0c, 0xda, 0x25, 0xff, 0x7f, 0x70, 0x85, 0x4e, 0x1b,
	0x5f, 0xee, 0x88, 0xf5, 0xf7, 0x8e, 0x58, 0x2f, 0xbe, 0x21, 0xdc, 0x88, 0x39, 0x1d, 0xc8, 0x44,
	0x32, 0xf7, 0x39, 0x6e, 0x0e, 0x39, 0xbd, 0xe1, 0x59, 0x91, 0x4c, 0x38, 0x75, 0x2c, 0xef, 0xe9,
	0x7c, 0x11, 0xb4, 0x76, 0xf5, 0xa5, 0x89, 0xdd, 0x0e, 0x6e, 0xa8, 0x29, 0x98, 0xb1, 0xcc, 0x41,
	0x5e, 0x73, 0xbe, 0x08, 0x1e, 0xc7, 0x9c, 0x5e, 0xcd, 0x58, 0xe6, 0x3e, 0xc3, 0x58, 0x55, 0xc9,
	0x48, 0xf2, 0x82, 0x39, 0x35, 0xef, 0x68, 0xbe, 0x08, 0x0e, 0x63, 0x4e, 0xcf, 0x74, 0xb0, 0x93,
	0x13, 0x10, 0xd2, 0x39, 0xa8, 0xe4, 0x5b, 0x10, 0x72, 0x27, 0x47, 0x13, 0x10, 0x8c, 0x3a, 0x76,
	0x25, 0xdf, 0xe8, 0xc0, 0xb3, 0x3f, 0x7d, 0xf5, 0xad, 0xf8, 0xf5, 0x8f, 0xa5, 0x8f, 0xee, 0x97,
	0x3e, 0xfa, 0xb3, 0xf4, 0xd1, 0xe7, 0x95, 0x6f, 0xdd, 0xaf, 0x7c, 0xeb, 0xf7, 0xca, 0xb7, 0xde,
	0x91, 0x14, 0x42, 0xbd, 0x5c, 0x34, 0xe4, 0x10, 0x09, 0x3a, 0x8e, 0x32, 0xa0, 0xec, 0x61, 0x09,
	0x87, 0x8f, 0xf4, 0xfe, 0x9c, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xf3, 0x6e, 0x87, 0x9e,
	0x02, 0x00, 0x00,
}

func (m *BidID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BidID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BidID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintBid(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x2a
	}
	if m.OSeq != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.OSeq))
		i--
		dAtA[i] = 0x20
	}
	if m.GSeq != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.GSeq))
		i--
		dAtA[i] = 0x18
	}
	if m.DSeq != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintBid(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBid(dAtA []byte, offset int, v uint64) int {
	offset -= sovBid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BidID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovBid(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovBid(uint64(m.DSeq))
	}
	if m.GSeq != 0 {
		n += 1 + sovBid(uint64(m.GSeq))
	}
	if m.OSeq != 0 {
		n += 1 + sovBid(uint64(m.OSeq))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovBid(uint64(l))
	}
	return n
}

func sovBid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBid(x uint64) (n int) {
	return sovBid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BidID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBid
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
			return fmt.Errorf("proto: BidID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BidID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DSeq", wireType)
			}
			m.DSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GSeq", wireType)
			}
			m.GSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSeq", wireType)
			}
			m.OSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
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
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBid
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
					return 0, ErrIntOverflowBid
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
					return 0, ErrIntOverflowBid
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
				return 0, ErrInvalidLengthBid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBid = fmt.Errorf("proto: unexpected end of group")
)
