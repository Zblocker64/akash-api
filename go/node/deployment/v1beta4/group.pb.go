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

// State is an enum which refers to state of group.
type Group_State int32

const (
	// Prefix should start with 0 in enum. So declaring dummy state.
	GroupStateInvalid Group_State = 0
	// GroupOpen denotes state for group open.
	GroupOpen Group_State = 1
	// GroupOrdered denotes state for group ordered.
	GroupPaused Group_State = 2
	// GroupInsufficientFunds denotes state for group insufficient_funds.
	GroupInsufficientFunds Group_State = 3
	// GroupClosed denotes state for group closed.
	GroupClosed Group_State = 4
)

var Group_State_name = map[int32]string{
	0: "invalid",
	1: "open",
	2: "paused",
	3: "insufficient_funds",
	4: "closed",
}

var Group_State_value = map[string]int32{
	"invalid":            0,
	"open":               1,
	"paused":             2,
	"insufficient_funds": 3,
	"closed":             4,
}

func (x Group_State) String() string {
	return proto.EnumName(Group_State_name, int32(x))
}

func (Group_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a45c04780ffee23e, []int{0, 0}
}

// Group stores group id, state and specifications of a group.
type Group struct {
	// Id is the unique identifier for the group.
	ID v1.GroupID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
	// State represents the state of the group.
	State Group_State `protobuf:"varint,2,opt,name=state,proto3,enum=akash.deployment.v1beta4.Group_State" json:"state" yaml:"state"`
	// GroupSpec holds the specification of a the Group.
	GroupSpec GroupSpec `protobuf:"bytes,3,opt,name=group_spec,json=groupSpec,proto3,castrepeated=GroupSpecs" json:"spec" yaml:"spec"`
	// CreatedAt is the block height at which the deployment was created.
	CreatedAt int64 `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
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

func (m *Group) GetState() Group_State {
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
	proto.RegisterEnum("akash.deployment.v1beta4.Group_State", Group_State_name, Group_State_value)
	proto.RegisterType((*Group)(nil), "akash.deployment.v1beta4.Group")
}

func init() {
	proto.RegisterFile("akash/deployment/v1beta4/group.proto", fileDescriptor_a45c04780ffee23e)
}

var fileDescriptor_a45c04780ffee23e = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x3d, 0x8f, 0xd3, 0x30,
	0x18, 0xc7, 0x93, 0x34, 0x3d, 0x54, 0x97, 0x97, 0x62, 0xf1, 0x12, 0x02, 0x17, 0x47, 0xe1, 0x45,
	0x65, 0x49, 0x44, 0x61, 0xba, 0x09, 0xca, 0x09, 0x54, 0x31, 0x80, 0x7a, 0x12, 0x03, 0x4b, 0xe5,
	0x8b, 0xdd, 0x60, 0xb5, 0x8d, 0xad, 0xc6, 0xad, 0x74, 0xdf, 0x00, 0x75, 0xe2, 0x0b, 0x54, 0x42,
	0x62, 0xe3, 0x13, 0xf0, 0x11, 0x6e, 0xbc, 0x91, 0x29, 0xa0, 0x76, 0x41, 0x1d, 0xbb, 0xb2, 0x20,
	0xdb, 0x05, 0x3a, 0x14, 0xd8, 0xdc, 0xff, 0xff, 0xe7, 0x5f, 0x9f, 0x47, 0x31, 0xb8, 0x83, 0x07,
	0xb8, 0x78, 0x9b, 0x10, 0x2a, 0x86, 0xfc, 0x64, 0x44, 0x73, 0x99, 0x4c, 0x1f, 0x1c, 0x53, 0x89,
	0x1f, 0x25, 0xd9, 0x98, 0x4f, 0x44, 0x2c, 0xc6, 0x5c, 0x72, 0xe8, 0x69, 0x2a, 0xfe, 0x43, 0xc5,
	0x1b, 0xca, 0xbf, 0x92, 0xf1, 0x8c, 0x6b, 0x28, 0x51, 0x27, 0xc3, 0xfb, 0x68, 0x87, 0x75, 0x5b,
	0xe8, 0x37, 0xff, 0xfd, 0xb7, 0x85, 0xa0, 0xa9, 0x21, 0xa3, 0x1f, 0x15, 0x50, 0x7d, 0xae, 0x32,
	0xf8, 0x02, 0x38, 0x8c, 0x78, 0x76, 0x68, 0x37, 0xeb, 0xad, 0x5b, 0xf1, 0x8e, 0x89, 0x62, 0xcd,
	0x75, 0x0e, 0xdb, 0xfb, 0xa7, 0x25, 0xb2, 0x16, 0x25, 0x72, 0x3a, 0x87, 0xab, 0x12, 0x39, 0x8c,
	0xac, 0x4b, 0x54, 0x3b, 0xc1, 0xa3, 0xe1, 0x41, 0xc4, 0x48, 0xd4, 0x75, 0x18, 0x81, 0xaf, 0x41,
	0xb5, 0x90, 0x58, 0x52, 0xcf, 0x09, 0xed, 0xe6, 0xc5, 0xd6, 0xdd, 0xf8, 0x6f, 0x1b, 0x1a, 0x69,
	0x7c, 0xa4, 0xe0, 0xf6, 0x8d, 0x55, 0x89, 0xcc, 0xbd, 0x75, 0x89, 0xce, 0x1b, 0xa3, 0xfe, 0x19,
	0x75, 0x4d, 0x0c, 0x47, 0x00, 0xe8, 0x0d, 0x7a, 0x6a, 0x05, 0xaf, 0xa2, 0x87, 0xbd, 0xfd, 0x1f,
	0xf9, 0x91, 0xa0, 0x69, 0xfb, 0xbe, 0x9a, 0x79, 0x55, 0x22, 0x57, 0x5d, 0x5c, 0x97, 0xa8, 0xbe,
	0xb1, 0x0b, 0x9a, 0x46, 0x9f, 0xbe, 0x22, 0xf0, 0x9b, 0x2c, 0xba, 0xb5, 0xec, 0xd7, 0x19, 0xee,
	0x03, 0x90, 0x8e, 0x29, 0x96, 0x94, 0xf4, 0xb0, 0xf4, 0xdc, 0xd0, 0x6e, 0x56, 0xba, 0xb5, 0x4d,
	0xf2, 0x44, 0x46, 0x9f, 0x6d, 0x50, 0xd5, 0x93, 0xc3, 0x08, 0x9c, 0x63, 0xf9, 0x14, 0x0f, 0x19,
	0x69, 0x58, 0xfe, 0xd5, 0xd9, 0x3c, 0xbc, 0x6c, 0x84, 0xaa, 0xec, 0x98, 0x02, 0x5e, 0x07, 0x2e,
	0x17, 0x34, 0x6f, 0xd8, 0xfe, 0x85, 0xd9, 0x3c, 0xac, 0x69, 0xe0, 0xa5, 0xa0, 0x39, 0xbc, 0x09,
	0xf6, 0x04, 0x9e, 0x14, 0x94, 0x34, 0x1c, 0xff, 0xd2, 0x6c, 0x1e, 0xd6, 0x75, 0xf5, 0x4a, 0x47,
	0xb0, 0x05, 0x20, 0xcb, 0x8b, 0x49, 0xbf, 0xcf, 0x52, 0x46, 0x73, 0xd9, 0xeb, 0x4f, 0x72, 0x52,
	0x34, 0x2a, 0xbe, 0x3f, 0x9b, 0x87, 0xd7, 0xcc, 0x17, 0xd9, 0xaa, 0x9f, 0xa9, 0x56, 0x09, 0xd3,
	0x21, 0x57, 0x42, 0x77, 0x4b, 0xf8, 0x54, 0x47, 0xbe, 0xfb, 0xee, 0x63, 0x60, 0x1d, 0xb8, 0xdf,
	0x3f, 0x20, 0xab, 0xfd, 0xf8, 0x74, 0x11, 0xd8, 0x67, 0x8b, 0xc0, 0xfe, 0xb6, 0x08, 0xec, 0xf7,
	0xcb, 0xc0, 0x3a, 0x5b, 0x06, 0xd6, 0x97, 0x65, 0x60, 0xbd, 0xb9, 0x27, 0x06, 0x59, 0x8c, 0x07,
	0x32, 0x26, 0x74, 0x9a, 0x64, 0x3c, 0xc9, 0x39, 0xa1, 0x3b, 0xde, 0xd3, 0xf1, 0x9e, 0x7e, 0x46,
	0x0f, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x39, 0xf7, 0x33, 0xba, 0xe9, 0x02, 0x00, 0x00,
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
				m.State |= Group_State(b&0x7F) << shift
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
