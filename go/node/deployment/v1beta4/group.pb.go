// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/deployment/v1beta4/group.proto

package v1beta4

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	v1 "pkg.akt.dev/go/node/deployment/v1"
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

// GroupState is an enum which refers to state of group
type GroupState int32

const (
	// Prefix should start with 0 in enum. So declaring dummy state
	GroupStateInvalid GroupState = 0
	// GroupOpen denotes state for group open
	GroupOpen GroupState = 1
	// GroupOrdered denotes state for group ordered
	GroupPaused GroupState = 2
	// GroupInsufficientFunds denotes state for group insufficient_funds
	GroupInsufficientFunds GroupState = 3
	// GroupClosed denotes state for group closed
	GroupClosed GroupState = 4
)

var GroupState_name = map[int32]string{
	0: "invalid",
	1: "open",
	2: "paused",
	3: "insufficient_funds",
	4: "closed",
}

var GroupState_value = map[string]int32{
	"invalid":            0,
	"open":               1,
	"paused":             2,
	"insufficient_funds": 3,
	"closed":             4,
}

func (x GroupState) String() string {
	return proto.EnumName(GroupState_name, int32(x))
}

func (GroupState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a45c04780ffee23e, []int{0}
}

// Group stores group id, state and specifications of group
type Group struct {
	ID        v1.GroupID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
	State     GroupState `protobuf:"varint,2,opt,name=state,proto3,enum=akash.deployment.v1beta4.GroupState" json:"state" yaml:"state"`
	GroupSpec GroupSpec  `protobuf:"bytes,3,opt,name=group_spec,json=groupSpec,proto3,castrepeated=GroupSpecs" json:"spec" yaml:"spec"`
	CreatedAt int64      `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_a45c04780ffee23e, []int{0}
}
func (m *Group) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Group.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return m.Size()
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetID() v1.GroupID {
	if m != nil {
		return m.ID
	}
	return v1.GroupID{}
}

func (m *Group) GetState() GroupState {
	if m != nil {
		return m.State
	}
	return GroupStateInvalid
}

func (m *Group) GetGroupSpec() GroupSpec {
	if m != nil {
		return m.GroupSpec
	}
	return GroupSpec{}
}

func (m *Group) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func init() {
	proto.RegisterEnum("akash.deployment.v1beta4.GroupState", GroupState_name, GroupState_value)
	proto.RegisterType((*Group)(nil), "akash.deployment.v1beta4.Group")
}

func init() {
	proto.RegisterFile("akash/deployment/v1beta4/group.proto", fileDescriptor_a45c04780ffee23e)
}

var fileDescriptor_a45c04780ffee23e = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x3f, 0x8f, 0xd3, 0x3c,
	0x1c, 0xc7, 0x93, 0x34, 0x77, 0x8f, 0xea, 0x3e, 0x40, 0xb1, 0xf8, 0x13, 0x0c, 0x17, 0x47, 0xe1,
	0x84, 0x0a, 0x43, 0x22, 0x0a, 0xd3, 0x4d, 0x50, 0x4e, 0xa0, 0x8a, 0x01, 0x54, 0xc4, 0xc2, 0x52,
	0xf9, 0x62, 0x37, 0x58, 0x6d, 0x63, 0xab, 0x71, 0x2b, 0xdd, 0x3b, 0x40, 0x9d, 0x78, 0x03, 0x95,
	0x90, 0xd8, 0x78, 0x0f, 0x88, 0xf5, 0xc6, 0x1b, 0x99, 0x02, 0x6a, 0x17, 0xd4, 0xb1, 0xaf, 0x00,
	0xd9, 0x29, 0xb4, 0x43, 0x81, 0xcd, 0xfd, 0x7e, 0x3f, 0xfe, 0xf4, 0xf7, 0x53, 0x0c, 0x0e, 0x49,
	0x9f, 0xe4, 0x6f, 0x63, 0xca, 0xe4, 0x40, 0x9c, 0x0e, 0x59, 0xa6, 0xe2, 0xc9, 0xfd, 0x13, 0xa6,
	0xc8, 0xc3, 0x38, 0x1d, 0x89, 0xb1, 0x8c, 0xe4, 0x48, 0x28, 0x01, 0x3d, 0x43, 0x45, 0x1b, 0x2a,
	0x5a, 0x53, 0xe8, 0x4a, 0x2a, 0x52, 0x61, 0xa0, 0x58, 0x9f, 0x4a, 0x1e, 0xe1, 0x1d, 0xd6, 0x6d,
	0x21, 0x6a, 0xfc, 0xfd, 0x6f, 0x73, 0xc9, 0x92, 0x92, 0x0c, 0xbf, 0x38, 0x60, 0xef, 0x99, 0xce,
	0xe0, 0x73, 0xe0, 0x70, 0xea, 0xd9, 0x81, 0xdd, 0xa8, 0x35, 0x6f, 0x45, 0x3b, 0x26, 0x8a, 0x0c,
	0xd7, 0x3e, 0x6e, 0x1d, 0x9c, 0x15, 0xd8, 0x9a, 0x17, 0xd8, 0x69, 0x1f, 0x2f, 0x0b, 0xec, 0x70,
	0xba, 0x2a, 0x70, 0xf5, 0x94, 0x0c, 0x07, 0x47, 0x21, 0xa7, 0x61, 0xc7, 0xe1, 0x14, 0xbe, 0x06,
	0x7b, 0xb9, 0x22, 0x8a, 0x79, 0x4e, 0x60, 0x37, 0x2e, 0x36, 0x0f, 0xa3, 0x3f, 0x6d, 0x58, 0x4a,
	0x5f, 0x69, 0xb6, 0x75, 0x63, 0x59, 0xe0, 0xf2, 0xda, 0xaa, 0xc0, 0xff, 0x97, 0x42, 0xf3, 0x33,
	0xec, 0x94, 0x31, 0x1c, 0x02, 0x60, 0x16, 0xe8, 0xea, 0x0d, 0xbc, 0x8a, 0x99, 0xf5, 0xf6, 0xbf,
	0xdc, 0x92, 0x25, 0xad, 0xbb, 0x7a, 0xe4, 0x65, 0x81, 0x5d, 0x7d, 0x71, 0x55, 0xe0, 0xda, 0xda,
	0x2e, 0x59, 0x12, 0x7e, 0xfa, 0x86, 0xc1, 0x6f, 0x32, 0xef, 0x54, 0xd3, 0x5f, 0x67, 0x78, 0x00,
	0x40, 0x32, 0x62, 0x44, 0x31, 0xda, 0x25, 0xca, 0x73, 0x03, 0xbb, 0x51, 0xe9, 0x54, 0xd7, 0xc9,
	0x63, 0x75, 0xe4, 0xfe, 0xf8, 0x80, 0xad, 0x7b, 0x9f, 0x6d, 0x00, 0x36, 0x4b, 0xc0, 0x10, 0xfc,
	0xc7, 0xb3, 0x09, 0x19, 0x70, 0x5a, 0xb7, 0xd0, 0xd5, 0xe9, 0x2c, 0xb8, 0xbc, 0x29, 0xdb, 0x65,
	0x01, 0xaf, 0x03, 0x57, 0x48, 0x96, 0xd5, 0x6d, 0x74, 0x61, 0x3a, 0x0b, 0xaa, 0x06, 0x78, 0x21,
	0x59, 0x06, 0x6f, 0x82, 0x7d, 0x49, 0xc6, 0x39, 0xa3, 0x75, 0x07, 0x5d, 0x9a, 0xce, 0x82, 0x9a,
	0xa9, 0x5e, 0x9a, 0x08, 0x36, 0x01, 0xe4, 0x59, 0x3e, 0xee, 0xf5, 0x78, 0xc2, 0x59, 0xa6, 0xba,
	0xbd, 0x71, 0x46, 0xf3, 0x7a, 0x05, 0xa1, 0xe9, 0x2c, 0xb8, 0x56, 0x7e, 0x9b, 0xad, 0xfa, 0xa9,
	0x6e, 0xb5, 0x30, 0x19, 0x08, 0x2d, 0x74, 0xb7, 0x84, 0x4f, 0x4c, 0x84, 0xdc, 0x77, 0x1f, 0x7d,
	0xab, 0xf5, 0xe8, 0x6c, 0xee, 0xdb, 0xe7, 0x73, 0xdf, 0xfe, 0x3e, 0xf7, 0xed, 0xf7, 0x0b, 0xdf,
	0x3a, 0x5f, 0xf8, 0xd6, 0xd7, 0x85, 0x6f, 0xbd, 0xb9, 0x23, 0xfb, 0x69, 0x44, 0xfa, 0x2a, 0xa2,
	0x6c, 0x12, 0xa7, 0x22, 0xce, 0x04, 0x65, 0x3b, 0xde, 0xd4, 0xc9, 0xbe, 0x79, 0x4a, 0x0f, 0x7e,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x12, 0x4d, 0x64, 0x2b, 0xed, 0x02, 0x00, 0x00,
}

func (m *Group) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Group) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Group) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.GroupSpec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.State != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroup(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGroup(dAtA []byte, offset int, v uint64) int {
	offset -= sovGroup(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Group) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovGroup(uint64(l))
	if m.State != 0 {
		n += 1 + sovGroup(uint64(m.State))
	}
	l = m.GroupSpec.Size()
	n += 1 + l + sovGroup(uint64(l))
	if m.CreatedAt != 0 {
		n += 1 + sovGroup(uint64(m.CreatedAt))
	}
	return n
}

func sovGroup(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGroup(x uint64) (n int) {
	return sovGroup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Group) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
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
			return fmt.Errorf("proto: Group: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Group: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
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
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= GroupState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GroupSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowGroup
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
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroup
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
func skipGroup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGroup
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
					return 0, ErrIntOverflowGroup
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
					return 0, ErrIntOverflowGroup
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
				return 0, ErrInvalidLengthGroup
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGroup
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGroup        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGroup          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGroup = fmt.Errorf("proto: unexpected end of group")
)
