// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/market/v1beta5/order.proto

package v1beta5

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	v1beta4 "pkg.akt.io/go/node/deployment/v1beta4"
	v1 "pkg.akt.io/go/node/market/v1"
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

// Order stores orderID, state of order and other details
type Order struct {
	ID        v1.OrderID        `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
	State     v1.OrderState     `protobuf:"varint,2,opt,name=state,proto3,enum=akash.market.v1.OrderState" json:"state" yaml:"state"`
	Spec      v1beta4.GroupSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec" yaml:"spec"`
	CreatedAt int64             `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (m *Order) Reset()      { *m = Order{} }
func (*Order) ProtoMessage() {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_a72454f2c693d67f, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetID() v1.OrderID {
	if m != nil {
		return m.ID
	}
	return v1.OrderID{}
}

func (m *Order) GetState() v1.OrderState {
	if m != nil {
		return m.State
	}
	return v1.OrderStateInvalid
}

func (m *Order) GetSpec() v1beta4.GroupSpec {
	if m != nil {
		return m.Spec
	}
	return v1beta4.GroupSpec{}
}

func (m *Order) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Order)(nil), "akash.market.v1beta5.Order")
}

func init() { proto.RegisterFile("akash/market/v1beta5/order.proto", fileDescriptor_a72454f2c693d67f) }

var fileDescriptor_a72454f2c693d67f = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3f, 0x4b, 0xf3, 0x40,
	0x18, 0xc0, 0x93, 0xb4, 0x7d, 0x79, 0x9b, 0x8a, 0x43, 0xe8, 0x10, 0x5b, 0x9a, 0x8b, 0x71, 0xc9,
	0x74, 0xc1, 0x5a, 0x97, 0x3a, 0x19, 0x0a, 0xa5, 0x20, 0x08, 0xa9, 0x93, 0x8b, 0x5c, 0x73, 0x47,
	0x0c, 0x69, 0x7b, 0x21, 0x3d, 0x85, 0x7e, 0x0b, 0x47, 0x17, 0xa1, 0x1f, 0xa7, 0x63, 0x47, 0xa7,
	0x43, 0xd2, 0x45, 0x3a, 0xf6, 0x13, 0xc8, 0x5d, 0x22, 0x4a, 0x71, 0xbb, 0xe7, 0xe1, 0x77, 0xbf,
	0xe7, 0x9f, 0x6e, 0xa3, 0x04, 0x2d, 0x1e, 0xbd, 0x19, 0xca, 0x12, 0xc2, 0xbc, 0xe7, 0xf3, 0x09,
	0x61, 0xe8, 0xd2, 0xa3, 0x19, 0x26, 0x19, 0x4c, 0x33, 0xca, 0xa8, 0xd1, 0x94, 0x04, 0x2c, 0x08,
	0x58, 0x12, 0xad, 0x66, 0x44, 0x23, 0x2a, 0x01, 0x4f, 0xbc, 0x0a, 0xb6, 0xe5, 0x16, 0x36, 0x4c,
	0xd2, 0x29, 0x5d, 0xce, 0xc8, 0xfc, 0xdb, 0xd8, 0xf3, 0xa2, 0x8c, 0x3e, 0xa5, 0x8b, 0x94, 0x84,
	0x25, 0xd9, 0x3e, 0xa8, 0xfb, 0xbb, 0xa4, 0xf3, 0xa6, 0xe9, 0xb5, 0x5b, 0x11, 0x1b, 0x43, 0x5d,
	0x8b, 0xb1, 0xa9, 0xda, 0xaa, 0xdb, 0xe8, 0x9a, 0xf0, 0xa0, 0x13, 0x28, 0x99, 0xd1, 0xc0, 0xef,
	0xac, 0x39, 0x50, 0x72, 0x0e, 0xb4, 0xd1, 0x60, 0xc7, 0x81, 0x16, 0xe3, 0x3d, 0x07, 0xf5, 0x25,
	0x9a, 0x4d, 0xfb, 0x4e, 0x8c, 0x9d, 0x40, 0x8b, 0xb1, 0x71, 0xa3, 0xd7, 0x16, 0x0c, 0x31, 0x62,
	0x6a, 0xb6, 0xea, 0x1e, 0x77, 0xdb, 0x7f, 0xbb, 0xc6, 0x02, 0xf1, 0x4f, 0x76, 0x1c, 0x14, 0xf4,
	0x9e, 0x83, 0xa3, 0xc2, 0x23, 0x43, 0x27, 0x28, 0xd2, 0xc6, 0x9d, 0x5e, 0x15, 0xb3, 0x98, 0x15,
	0xd9, 0xd8, 0x59, 0x29, 0xfb, 0x19, 0xbb, 0x5c, 0x53, 0x0f, 0x0e, 0xc5, 0xd8, 0xe3, 0x94, 0x84,
	0x7e, 0x5b, 0xf4, 0xb8, 0xe3, 0x40, 0x7e, 0xdc, 0x73, 0xd0, 0x28, 0xbd, 0x29, 0x09, 0x9d, 0x40,
	0x26, 0x8d, 0x8e, 0xae, 0x87, 0x19, 0x41, 0x8c, 0xe0, 0x07, 0xc4, 0xcc, 0xaa, 0xad, 0xba, 0x95,
	0xa0, 0x5e, 0x66, 0xae, 0x59, 0xff, 0xff, 0xeb, 0x0a, 0x28, 0x9f, 0x2b, 0xa0, 0xf8, 0x57, 0xeb,
	0xdc, 0x52, 0x37, 0xb9, 0xa5, 0x7e, 0xe4, 0x96, 0xfa, 0xb2, 0xb5, 0x94, 0xcd, 0xd6, 0x52, 0xde,
	0xb7, 0x96, 0x72, 0x7f, 0x9a, 0x26, 0x11, 0x44, 0x09, 0x83, 0xb1, 0xb8, 0x8b, 0x37, 0xa7, 0x98,
	0x1c, 0xdc, 0x76, 0xf2, 0x4f, 0xee, 0xf8, 0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x48, 0x80, 0xed,
	0x3f, 0xfa, 0x01, 0x00, 0x00,
}

func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.State != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovOrder(uint64(l))
	if m.State != 0 {
		n += 1 + sovOrder(uint64(m.State))
	}
	l = m.Spec.Size()
	n += 1 + l + sovOrder(uint64(l))
	if m.CreatedAt != 0 {
		n += 1 + sovOrder(uint64(m.CreatedAt))
	}
	return n
}

func sovOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrder(x uint64) (n int) {
	return sovOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
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
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= v1.OrderState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowOrder
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
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrder
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
func skipOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
				return 0, ErrInvalidLengthOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrder = fmt.Errorf("proto: unexpected end of group")
)
