// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/deployment/v1beta4/filters.proto

package v1beta4

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	_ "pkg.akt.io/go/node/deployment/v1"
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

// DeploymentFilters defines filters used to filter deployments
type DeploymentFilters struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	DSeq  uint64 `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	State string `protobuf:"bytes,3,opt,name=state,proto3" json:"state" yaml:"state"`
}

func (m *DeploymentFilters) Reset()         { *m = DeploymentFilters{} }
func (m *DeploymentFilters) String() string { return proto.CompactTextString(m) }
func (*DeploymentFilters) ProtoMessage()    {}
func (*DeploymentFilters) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb62db02ae97591f, []int{0}
}
func (m *DeploymentFilters) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeploymentFilters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeploymentFilters.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeploymentFilters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentFilters.Merge(m, src)
}
func (m *DeploymentFilters) XXX_Size() int {
	return m.Size()
}
func (m *DeploymentFilters) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentFilters.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentFilters proto.InternalMessageInfo

func (m *DeploymentFilters) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *DeploymentFilters) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *DeploymentFilters) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*DeploymentFilters)(nil), "akash.deployment.v1beta4.DeploymentFilters")
}

func init() {
	proto.RegisterFile("akash/deployment/v1beta4/filters.proto", fileDescriptor_fb62db02ae97591f)
}

var fileDescriptor_fb62db02ae97591f = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4b, 0xcc, 0x4e, 0x2c,
	0xce, 0xd0, 0x4f, 0x49, 0x2d, 0xc8, 0xc9, 0xaf, 0xcc, 0x4d, 0xcd, 0x2b, 0xd1, 0x2f, 0x33, 0x4c,
	0x4a, 0x2d, 0x49, 0x34, 0xd1, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0x2a, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x92, 0x00, 0xab, 0xd3, 0x43, 0xa8, 0xd3, 0x83, 0xaa, 0x93, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0x2b, 0xd2, 0x07, 0xb1, 0x20, 0xea, 0xa5, 0x24, 0x93, 0xf3, 0x8b, 0x73, 0xf3,
	0x8b, 0xe3, 0x21, 0x12, 0x10, 0x0e, 0x54, 0x4a, 0x05, 0x8b, 0x95, 0x48, 0x3c, 0x88, 0x2a, 0xa5,
	0x13, 0x8c, 0x5c, 0x82, 0x2e, 0x70, 0x41, 0x37, 0x88, 0x63, 0x84, 0xdc, 0xb9, 0x58, 0xf3, 0xcb,
	0xf3, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0x0c, 0x5f, 0xdd, 0x93, 0x87, 0x08,
	0x7c, 0xba, 0x27, 0xcf, 0x53, 0x99, 0x98, 0x9b, 0x63, 0xa5, 0x04, 0xe6, 0x2a, 0x5d, 0xda, 0xa2,
	0x2b, 0x02, 0xb5, 0xd5, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x38, 0xb8, 0xa4, 0x28, 0x33, 0x2f,
	0x3d, 0x08, 0xa2, 0x5c, 0xc8, 0x98, 0x8b, 0x25, 0xa5, 0x38, 0xb5, 0x50, 0x82, 0x49, 0x81, 0x51,
	0x83, 0xc5, 0x49, 0xfe, 0xd1, 0x3d, 0x79, 0x16, 0x97, 0xe0, 0xd4, 0xc2, 0x57, 0xf7, 0xe4, 0xc1,
	0xe2, 0x9f, 0xee, 0xc9, 0x73, 0x43, 0x8c, 0x03, 0xf1, 0x94, 0x82, 0xc0, 0x82, 0x42, 0xfa, 0x5c,
	0xac, 0xc5, 0x25, 0x89, 0x25, 0xa9, 0x12, 0xcc, 0x60, 0xdb, 0x25, 0x41, 0xb6, 0x83, 0x05, 0x10,
	0xb6, 0x83, 0xb9, 0x4a, 0x41, 0x10, 0x61, 0x2b, 0x96, 0x17, 0x0b, 0xe4, 0x19, 0x9c, 0xec, 0x4f,
	0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18,
	0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0xb5, 0x20, 0x3b, 0x5d, 0x2f, 0x31,
	0xbb, 0x44, 0x2f, 0x13, 0x14, 0x78, 0xfa, 0x79, 0xf9, 0x29, 0xa9, 0x58, 0xa2, 0x22, 0x89, 0x0d,
	0x1c, 0x24, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0xce, 0x44, 0xd5, 0xad, 0x01, 0x00,
	0x00,
}

func (m *DeploymentFilters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeploymentFilters) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeploymentFilters) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintFilters(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x1a
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
func (m *DeploymentFilters) Size() (n int) {
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
func (m *DeploymentFilters) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DeploymentFilters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeploymentFilters: illegal tag %d (wire type %d)", fieldNum, wire)
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
