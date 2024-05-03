// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/market/v1/filters.proto

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

// LeaseFilters defines flags for lease list filtering.
type LeaseFilters struct {
	// Owner is the account bech32 address of the user who owns the deployment.
	// It is a string representing a valid bech32 account address.
	//
	// Example:
	//   "akash1..."
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	// Dseq (deployment sequence number) is a unique numeric identifier for the deployment.
	// It is used to differentiate deployments created by the same owner.
	DSeq uint64 `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	// Gseq (group sequence number) is a unique numeric identifier for the group.
	// It is used to differentiate groups created by the same owner in a deployment.
	GSeq uint32 `protobuf:"varint,3,opt,name=gseq,proto3" json:"gseq" yaml:"gseq"`
	// Oseq (order sequence) distinguishes multiple orders associated with a single deployment.
	// Oseq is incremented when a lease associated with an existing deployment is closed, and a new order is generated.
	OSeq uint32 `protobuf:"varint,4,opt,name=oseq,proto3" json:"oseq" yaml:"oseq"`
	// Provider is the account bech32 address of the provider making the bid.
	// It is a string representing a valid account bech32 address.
	//
	// Example:
	//   "akash1..."
	Provider string `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider" yaml:"provider"`
	// State represents the state of the lease.
	State string `protobuf:"bytes,6,opt,name=state,proto3" json:"state" yaml:"state"`
}

func (m *LeaseFilters) Reset()         { *m = LeaseFilters{} }
func (m *LeaseFilters) String() string { return proto.CompactTextString(m) }
func (*LeaseFilters) ProtoMessage()    {}
func (*LeaseFilters) Descriptor() ([]byte, []int) {
	return fileDescriptor_601e383ab343d6e3, []int{0}
}
func (m *LeaseFilters) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaseFilters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaseFilters.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LeaseFilters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseFilters.Merge(m, src)
}
func (m *LeaseFilters) XXX_Size() int {
	return m.Size()
}
func (m *LeaseFilters) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseFilters.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseFilters proto.InternalMessageInfo

func (m *LeaseFilters) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *LeaseFilters) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *LeaseFilters) GetGSeq() uint32 {
	if m != nil {
		return m.GSeq
	}
	return 0
}

func (m *LeaseFilters) GetOSeq() uint32 {
	if m != nil {
		return m.OSeq
	}
	return 0
}

func (m *LeaseFilters) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *LeaseFilters) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*LeaseFilters)(nil), "akash.market.v1.LeaseFilters")
}

func init() { proto.RegisterFile("akash/market/v1/filters.proto", fileDescriptor_601e383ab343d6e3) }

var fileDescriptor_601e383ab343d6e3 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6a, 0xea, 0x40,
	0x14, 0x86, 0x93, 0x6b, 0x94, 0x7b, 0x73, 0xbd, 0x08, 0xc1, 0x45, 0x14, 0xcc, 0x48, 0x56, 0x6e,
	0x6e, 0x82, 0xb8, 0x10, 0xdc, 0x55, 0x4a, 0xdd, 0x14, 0x0a, 0xba, 0xeb, 0xa6, 0x4c, 0x9b, 0xe9,
	0x54, 0xa2, 0x8e, 0xce, 0x0c, 0x29, 0xdd, 0xf6, 0x09, 0xfa, 0x08, 0x7d, 0x88, 0x3e, 0x44, 0x97,
	0xd2, 0x55, 0x57, 0x43, 0x89, 0x9b, 0xe2, 0xd2, 0x27, 0x28, 0x33, 0x27, 0x55, 0x5c, 0x74, 0x95,
	0xfc, 0xff, 0x7f, 0xbe, 0x1f, 0xce, 0x49, 0xdc, 0x16, 0x4e, 0xb1, 0xb8, 0x8b, 0xe7, 0x98, 0xa7,
	0x44, 0xc6, 0x59, 0x37, 0xbe, 0x9d, 0xce, 0x24, 0xe1, 0x22, 0x5a, 0x72, 0x26, 0x99, 0x57, 0x33,
	0x71, 0x04, 0x71, 0x94, 0x75, 0x9b, 0x75, 0xca, 0x28, 0x33, 0x59, 0xac, 0xdf, 0x60, 0xac, 0xd9,
	0xb8, 0x61, 0x62, 0xce, 0xc4, 0x15, 0x04, 0x20, 0x20, 0x0a, 0x1f, 0x4b, 0x6e, 0xf5, 0x9c, 0x60,
	0x41, 0xce, 0xa0, 0xd8, 0x1b, 0xb9, 0x65, 0x76, 0xbf, 0x20, 0xdc, 0xb7, 0xdb, 0x76, 0xe7, 0xcf,
	0xb0, 0xbb, 0x55, 0x08, 0x8c, 0x9d, 0x42, 0xd5, 0x07, 0x3c, 0x9f, 0x0d, 0x42, 0x23, 0xc3, 0xb7,
	0x97, 0xff, 0xf5, 0xa2, 0xea, 0x24, 0x49, 0x38, 0x11, 0x62, 0x22, 0xf9, 0x74, 0x41, 0xc7, 0x30,
	0xee, 0xf5, 0x5c, 0x27, 0x11, 0x64, 0xe5, 0xff, 0x6a, 0xdb, 0x1d, 0x67, 0x88, 0x72, 0x85, 0x9c,
	0xd3, 0x09, 0x59, 0x6d, 0x15, 0x32, 0xfe, 0x4e, 0xa1, 0xbf, 0x50, 0xa7, 0x55, 0x38, 0x36, 0xa6,
	0x86, 0xa8, 0x86, 0x4a, 0x6d, 0xbb, 0xf3, 0x0f, 0xa0, 0x51, 0x01, 0xd1, 0x23, 0x88, 0x02, 0x44,
	0x0b, 0x88, 0x69, 0xc8, 0x39, 0x40, 0x17, 0x05, 0xc4, 0x8e, 0x20, 0x06, 0x90, 0x7e, 0x78, 0x13,
	0xf7, 0xf7, 0x92, 0xb3, 0x6c, 0x9a, 0x10, 0xee, 0x97, 0xcd, 0xaa, 0xfd, 0xad, 0x42, 0x7b, 0x6f,
	0xa7, 0x50, 0x0d, 0xa0, 0x6f, 0xe7, 0xe7, 0x85, 0xf7, 0x90, 0x17, 0xbb, 0x65, 0x21, 0xb1, 0x24,
	0x7e, 0xc5, 0x34, 0x36, 0xf4, 0xf1, 0x8c, 0x71, 0x38, 0x9e, 0x91, 0xe1, 0x18, 0xec, 0x81, 0xf3,
	0xf9, 0x8c, 0xac, 0x61, 0xff, 0x35, 0x0f, 0xec, 0x75, 0x1e, 0xd8, 0x1f, 0x79, 0x60, 0x3f, 0x6d,
	0x02, 0x6b, 0xbd, 0x09, 0xac, 0xf7, 0x4d, 0x60, 0x5d, 0xb6, 0x96, 0x29, 0x8d, 0x70, 0x2a, 0xa3,
	0x84, 0x64, 0x31, 0x65, 0xf1, 0x82, 0x25, 0xe4, 0xf0, 0x37, 0x5c, 0x57, 0xcc, 0x47, 0xec, 0x7d,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x36, 0xfb, 0x56, 0x68, 0x27, 0x02, 0x00, 0x00,
}

func (m *LeaseFilters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseFilters) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseFilters) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintFilters(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintFilters(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x2a
	}
	if m.OSeq != 0 {
		i = encodeVarintFilters(dAtA, i, uint64(m.OSeq))
		i--
		dAtA[i] = 0x20
	}
	if m.GSeq != 0 {
		i = encodeVarintFilters(dAtA, i, uint64(m.GSeq))
		i--
		dAtA[i] = 0x18
	}
	if m.DSeq != 0 {
		i = encodeVarintFilters(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintFilters(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFilters(dAtA []byte, offset int, v uint64) int {
	offset -= sovFilters(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LeaseFilters) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovFilters(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovFilters(uint64(m.DSeq))
	}
	if m.GSeq != 0 {
		n += 1 + sovFilters(uint64(m.GSeq))
	}
	if m.OSeq != 0 {
		n += 1 + sovFilters(uint64(m.OSeq))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovFilters(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovFilters(uint64(l))
	}
	return n
}

func sovFilters(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFilters(x uint64) (n int) {
	return sovFilters(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LeaseFilters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFilters
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
			return fmt.Errorf("proto: LeaseFilters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseFilters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilters
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
				return ErrInvalidLengthFilters
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFilters
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
					return ErrIntOverflowFilters
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
					return ErrIntOverflowFilters
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
					return ErrIntOverflowFilters
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
					return ErrIntOverflowFilters
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
				return ErrInvalidLengthFilters
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFilters
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilters
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
				return ErrInvalidLengthFilters
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFilters
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFilters(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFilters
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
func skipFilters(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFilters
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
					return 0, ErrIntOverflowFilters
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
					return 0, ErrIntOverflowFilters
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
				return 0, ErrInvalidLengthFilters
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFilters
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFilters
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFilters        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFilters          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFilters = fmt.Errorf("proto: unexpected end of group")
)
