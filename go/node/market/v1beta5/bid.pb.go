// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/market/v1beta5/bid.proto

package v1beta5

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	v1 "pkg.akt.dev/go/node/market/v1"
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

// Bid stores BidID, state of bid and price
type Bid struct {
	ID             v1.BidID       `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
	State          v1.BidState    `protobuf:"varint,2,opt,name=state,proto3,enum=akash.market.v1.BidState" json:"state" yaml:"state"`
	Price          types.DecCoin  `protobuf:"bytes,3,opt,name=price,proto3" json:"price" yaml:"price"`
	CreatedAt      int64          `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"state" yaml:"created_at"`
	ResourcesOffer ResourcesOffer `protobuf:"bytes,5,rep,name=resources_offer,json=resourcesOffer,proto3,castrepeated=ResourcesOffer" json:"resources_offer" yaml:"resources_offer"`
}

func (m *Bid) Reset()      { *m = Bid{} }
func (*Bid) ProtoMessage() {}
func (*Bid) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a051a9ee62f13b0, []int{0}
}
func (m *Bid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bid.Merge(m, src)
}
func (m *Bid) XXX_Size() int {
	return m.Size()
}
func (m *Bid) XXX_DiscardUnknown() {
	xxx_messageInfo_Bid.DiscardUnknown(m)
}

var xxx_messageInfo_Bid proto.InternalMessageInfo

func (m *Bid) GetID() v1.BidID {
	if m != nil {
		return m.ID
	}
	return v1.BidID{}
}

func (m *Bid) GetState() v1.BidState {
	if m != nil {
		return m.State
	}
	return v1.BidStateInvalid
}

func (m *Bid) GetPrice() types.DecCoin {
	if m != nil {
		return m.Price
	}
	return types.DecCoin{}
}

func (m *Bid) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Bid) GetResourcesOffer() ResourcesOffer {
	if m != nil {
		return m.ResourcesOffer
	}
	return nil
}

func init() {
	proto.RegisterType((*Bid)(nil), "akash.market.v1beta5.Bid")
}

func init() { proto.RegisterFile("akash/market/v1beta5/bid.proto", fileDescriptor_9a051a9ee62f13b0) }

var fileDescriptor_9a051a9ee62f13b0 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x3f, 0x6f, 0x13, 0x31,
	0x18, 0xc6, 0xef, 0x0f, 0x41, 0xd4, 0x45, 0x41, 0x9c, 0xaa, 0xea, 0x52, 0x81, 0x1d, 0x8e, 0x25,
	0x2c, 0x3e, 0x25, 0x88, 0xa5, 0x88, 0x81, 0x23, 0x0c, 0x65, 0x41, 0x1c, 0x62, 0x61, 0xa9, 0x9c,
	0xb3, 0x1b, 0xac, 0x90, 0x38, 0xb2, 0x0d, 0x12, 0xdf, 0x82, 0x91, 0xb1, 0x12, 0x1b, 0x9f, 0xa4,
	0x63, 0x47, 0xc4, 0x60, 0xd0, 0x65, 0x41, 0x19, 0xf3, 0x09, 0xd0, 0xd9, 0x86, 0xd0, 0x53, 0xb6,
	0x7b, 0xdf, 0xf7, 0xf7, 0x3e, 0xf7, 0xe8, 0xf1, 0x0b, 0x20, 0x99, 0x11, 0xf5, 0x2e, 0x9f, 0x13,
	0x39, 0x63, 0x3a, 0xff, 0x38, 0x9c, 0x30, 0x4d, 0x1e, 0xe5, 0x13, 0x4e, 0xf1, 0x52, 0x0a, 0x2d,
	0x92, 0x03, 0x3b, 0xc7, 0x6e, 0x8e, 0xfd, 0xfc, 0xe8, 0x60, 0x2a, 0xa6, 0xc2, 0x02, 0x79, 0xf3,
	0xe5, 0xd8, 0x23, 0x58, 0x09, 0x35, 0x17, 0x2a, 0x9f, 0x10, 0xc5, 0xbc, 0xd4, 0x30, 0xaf, 0x04,
	0x5f, 0xf8, 0xf9, 0x83, 0x9d, 0xff, 0x92, 0x4c, 0x89, 0x0f, 0xb2, 0x62, 0x4a, 0x9c, 0x9d, 0x31,
	0xe9, 0xd1, 0x5e, 0x0b, 0xdd, 0x3a, 0xca, 0x7e, 0xc4, 0x20, 0x2e, 0x38, 0x4d, 0x9e, 0x83, 0x88,
	0xd3, 0x34, 0xec, 0x87, 0x83, 0xfd, 0xd1, 0x21, 0x6e, 0xd9, 0xc4, 0x05, 0xa7, 0x27, 0xe3, 0xe2,
	0xee, 0x85, 0x41, 0x41, 0x6d, 0x50, 0x74, 0x32, 0x5e, 0x1b, 0x14, 0x71, 0xba, 0x31, 0x68, 0xef,
	0x13, 0x99, 0xbf, 0x3f, 0xce, 0x38, 0xcd, 0xca, 0x88, 0xd3, 0xe4, 0x05, 0xe8, 0x28, 0x4d, 0x34,
	0x4b, 0xa3, 0x7e, 0x38, 0xe8, 0x8e, 0x7a, 0xbb, 0x94, 0x5e, 0x37, 0x40, 0xd1, 0x5b, 0x1b, 0xe4,
	0xd8, 0x8d, 0x41, 0x37, 0x9d, 0x8a, 0x2d, 0xb3, 0xd2, 0xb5, 0x93, 0x57, 0xa0, 0xb3, 0x94, 0xbc,
	0x62, 0x69, 0x6c, 0x5d, 0xdd, 0xc1, 0x2e, 0x10, 0xdc, 0x04, 0xe2, 0xb3, 0x1b, 0xe2, 0x31, 0xab,
	0x9e, 0x09, 0xbe, 0x70, 0xde, 0x1a, 0x49, 0xbb, 0xb2, 0x95, 0xb4, 0x65, 0x56, 0xba, 0x76, 0xf2,
	0x04, 0x80, 0x4a, 0x32, 0xa2, 0x19, 0x3d, 0x25, 0x3a, 0xbd, 0xd6, 0x0f, 0x07, 0x71, 0x01, 0xff,
	0x37, 0x72, 0xdb, 0x6d, 0x6d, 0xa1, 0xac, 0xdc, 0xf3, 0xc5, 0x53, 0x9d, 0x7c, 0x0d, 0xc1, 0xad,
	0x7f, 0x01, 0x9f, 0xda, 0x84, 0xd3, 0x4e, 0x3f, 0x1e, 0xec, 0x8f, 0xee, 0xe3, 0x5d, 0x2f, 0x8b,
	0x4b, 0x0f, 0xbf, 0x6c, 0xd0, 0xe2, 0x8d, 0xcf, 0xaf, 0xfb, 0xb7, 0xad, 0x6c, 0x7f, 0x6d, 0x50,
	0x5b, 0x75, 0x63, 0xd0, 0xa1, 0x73, 0xd2, 0x1a, 0x64, 0xdf, 0x7e, 0xb6, 0xd7, 0xcb, 0xae, 0xbc,
	0x52, 0x1f, 0xdf, 0xf8, 0x72, 0x8e, 0x82, 0xdf, 0xe7, 0x28, 0x28, 0x1e, 0x5f, 0xd4, 0x30, 0xbc,
	0xac, 0x61, 0xf8, 0xab, 0x86, 0xe1, 0xe7, 0x15, 0x0c, 0x2e, 0x57, 0x30, 0xf8, 0xbe, 0x82, 0xc1,
	0xdb, 0x7b, 0xcb, 0xd9, 0x14, 0x93, 0x99, 0xc6, 0xbc, 0xb9, 0xb9, 0x7c, 0x21, 0x28, 0x6b, 0x5d,
	0xd2, 0xe4, 0xba, 0x3d, 0x90, 0x87, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x51, 0xbd, 0x89,
	0xd4, 0x02, 0x00, 0x00,
}

func (m *Bid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ResourcesOffer) > 0 {
		for iNdEx := len(m.ResourcesOffer) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResourcesOffer[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBid(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.CreatedAt != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.State != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *Bid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovBid(uint64(l))
	if m.State != 0 {
		n += 1 + sovBid(uint64(m.State))
	}
	l = m.Price.Size()
	n += 1 + l + sovBid(uint64(l))
	if m.CreatedAt != 0 {
		n += 1 + sovBid(uint64(m.CreatedAt))
	}
	if len(m.ResourcesOffer) > 0 {
		for _, e := range m.ResourcesOffer {
			l = e.Size()
			n += 1 + l + sovBid(uint64(l))
		}
	}
	return n
}

func sovBid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBid(x uint64) (n int) {
	return sovBid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Bid) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Bid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= v1.BidState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourcesOffer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourcesOffer = append(m.ResourcesOffer, ResourceOffer{})
			if err := m.ResourcesOffer[len(m.ResourcesOffer)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
