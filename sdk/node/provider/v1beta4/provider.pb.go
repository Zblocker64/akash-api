// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/provider/v1beta4/provider.proto

package v1beta4

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	go_akashd_io_sdk_node_types_attributes_v1 "go.akashd.io/sdk/node/types/attributes/v1"
	v1 "go.akashd.io/sdk/node/types/attributes/v1"
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

// Info
type Info struct {
	EMail   string `protobuf:"bytes,1,opt,name=email,proto3" json:"email" yaml:"email"`
	Website string `protobuf:"bytes,2,opt,name=website,proto3" json:"website" yaml:"website"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb1622664c70e47, []int{0}
}
func (m *Info) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Info.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(m, src)
}
func (m *Info) XXX_Size() int {
	return m.Size()
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetEMail() string {
	if m != nil {
		return m.EMail
	}
	return ""
}

func (m *Info) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

// Provider stores owner and host details
type Provider struct {
	Owner      string                                               `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	HostURI    string                                               `protobuf:"bytes,2,opt,name=host_uri,json=hostUri,proto3" json:"host_uri" yaml:"host_uri"`
	Attributes go_akashd_io_sdk_node_types_attributes_v1.Attributes `protobuf:"bytes,3,rep,name=attributes,proto3,castrepeated=go.akashd.io/sdk/node/types/attributes/v1.Attributes" json:"attributes" yaml:"attributes"`
	Info       Info                                                 `protobuf:"bytes,4,opt,name=info,proto3" json:"info" yaml:"info"`
}

func (m *Provider) Reset()      { *m = Provider{} }
func (*Provider) ProtoMessage() {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb1622664c70e47, []int{1}
}
func (m *Provider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Provider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Provider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Provider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Provider.Merge(m, src)
}
func (m *Provider) XXX_Size() int {
	return m.Size()
}
func (m *Provider) XXX_DiscardUnknown() {
	xxx_messageInfo_Provider.DiscardUnknown(m)
}

var xxx_messageInfo_Provider proto.InternalMessageInfo

func (m *Provider) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Provider) GetHostURI() string {
	if m != nil {
		return m.HostURI
	}
	return ""
}

func (m *Provider) GetAttributes() go_akashd_io_sdk_node_types_attributes_v1.Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Provider) GetInfo() Info {
	if m != nil {
		return m.Info
	}
	return Info{}
}

func init() {
	proto.RegisterType((*Info)(nil), "akash.provider.v1beta4.Info")
	proto.RegisterType((*Provider)(nil), "akash.provider.v1beta4.Provider")
}

func init() {
	proto.RegisterFile("akash/provider/v1beta4/provider.proto", fileDescriptor_cbb1622664c70e47)
}

var fileDescriptor_cbb1622664c70e47 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x3f, 0x6f, 0x13, 0x31,
	0x14, 0xbf, 0x23, 0x09, 0x09, 0x0e, 0x02, 0x71, 0xaa, 0x50, 0x28, 0x70, 0x8e, 0x8c, 0x40, 0x11,
	0x12, 0x3e, 0xa5, 0x54, 0x02, 0x75, 0xa2, 0x91, 0x2a, 0xe8, 0x50, 0x09, 0x1d, 0xea, 0xc2, 0x12,
	0xdd, 0x71, 0x6e, 0x6a, 0x35, 0x39, 0x47, 0xb6, 0x9b, 0x2a, 0xdf, 0x82, 0x11, 0x31, 0x75, 0x60,
	0x62, 0xe6, 0x43, 0x74, 0xac, 0x98, 0x98, 0x0c, 0xba, 0x2c, 0x28, 0xe3, 0x7d, 0x02, 0x74, 0xb6,
	0x2f, 0x57, 0xaa, 0x6c, 0xfe, 0xfd, 0x79, 0xef, 0xf7, 0x6c, 0x3f, 0xf0, 0x34, 0x3a, 0x89, 0xc4,
	0x71, 0x30, 0xe5, 0x6c, 0x46, 0x13, 0xc2, 0x83, 0x59, 0x3f, 0x26, 0x32, 0xda, 0x5e, 0x11, 0x78,
	0xca, 0x99, 0x64, 0xde, 0x7d, 0x6d, 0xc3, 0x2b, 0xd6, 0xda, 0x36, 0x37, 0x46, 0x6c, 0xc4, 0xb4,
	0x25, 0x28, 0x4e, 0xc6, 0xbd, 0xf9, 0xe0, 0x13, 0x13, 0x13, 0x26, 0x86, 0x46, 0x30, 0xc0, 0x4a,
	0x3d, 0x93, 0x17, 0x47, 0x82, 0x04, 0x91, 0x94, 0x9c, 0xc6, 0xa7, 0x92, 0x88, 0x60, 0xd6, 0xaf,
	0x90, 0x71, 0xa2, 0x39, 0xa8, 0xef, 0xa7, 0x47, 0xcc, 0x7b, 0x0d, 0x1a, 0x64, 0x12, 0xd1, 0x71,
	0xc7, 0xed, 0xba, 0xbd, 0x5b, 0x03, 0x94, 0x29, 0xd8, 0xd8, 0x3b, 0x88, 0xe8, 0x78, 0xa9, 0xa0,
	0x51, 0x72, 0x05, 0x6f, 0xcf, 0xa3, 0xc9, 0x78, 0x07, 0x69, 0x88, 0x42, 0x43, 0x7b, 0xaf, 0x40,
	0xf3, 0x8c, 0xc4, 0x82, 0x4a, 0xd2, 0xb9, 0xa1, 0x6b, 0x1f, 0x2f, 0x15, 0x2c, 0xa9, 0x5c, 0xc1,
	0x3b, 0xa6, 0xc8, 0x12, 0x28, 0x2c, 0x25, 0xf4, 0xb5, 0x06, 0x5a, 0xef, 0xed, 0x55, 0xbd, 0xb7,
	0xa0, 0xc1, 0xce, 0x52, 0xc2, 0x6d, 0x7e, 0xbf, 0x88, 0xd5, 0x44, 0x15, 0xab, 0x21, 0xfa, 0xf9,
	0xe3, 0xc5, 0x86, 0xbd, 0xeb, 0x6e, 0x92, 0x70, 0x22, 0xc4, 0x07, 0xc9, 0x69, 0x3a, 0x0a, 0x8d,
	0xdd, 0xdb, 0x03, 0xad, 0x63, 0x26, 0xe4, 0xf0, 0x94, 0x53, 0x3b, 0xcf, 0xf3, 0x4c, 0xc1, 0xe6,
	0x3b, 0x26, 0xe4, 0x61, 0xb8, 0xbf, 0x54, 0x70, 0x25, 0xe7, 0x0a, 0xde, 0x35, 0x9d, 0x4b, 0x06,
	0x85, 0xcd, 0xe2, 0x78, 0xc8, 0xa9, 0xf7, 0xcd, 0x05, 0xa0, 0x7a, 0xb9, 0x4e, 0xad, 0x5b, 0xeb,
	0xb5, 0xb7, 0x9e, 0x60, 0xf3, 0x41, 0xc5, 0xbb, 0xe2, 0x4a, 0xc5, 0xb3, 0x3e, 0xde, 0x2d, 0xd1,
	0x60, 0x78, 0xa1, 0xa0, 0xb3, 0x54, 0xf0, 0x4a, 0x79, 0xae, 0xe0, 0x3d, 0x93, 0x54, 0x71, 0xe8,
	0xfb, 0x6f, 0xb8, 0x3d, 0x62, 0xa6, 0x63, 0x82, 0x29, 0x0b, 0x44, 0x72, 0x12, 0xa4, 0x2c, 0x21,
	0x81, 0x9c, 0x4f, 0x89, 0xf8, 0xff, 0xdf, 0xaa, 0xfe, 0x22, 0xbc, 0xd2, 0xd8, 0x3b, 0x00, 0x75,
	0x9a, 0x1e, 0xb1, 0x4e, 0xbd, 0xeb, 0xf6, 0xda, 0x5b, 0x8f, 0xf0, 0xfa, 0x05, 0xc2, 0xc5, 0x17,
	0x0f, 0x1e, 0xda, 0xc1, 0x74, 0x45, 0xae, 0x60, 0xdb, 0x8c, 0x54, 0x20, 0x14, 0x6a, 0x72, 0xa7,
	0xf5, 0xe5, 0x1c, 0x3a, 0x7f, 0xcf, 0xa1, 0x33, 0x78, 0x73, 0x91, 0xf9, 0xee, 0x65, 0xe6, 0xbb,
	0x7f, 0x32, 0xdf, 0xfd, 0xbc, 0xf0, 0x9d, 0xcb, 0x85, 0xef, 0xfc, 0x5a, 0xf8, 0xce, 0xc7, 0x67,
	0xeb, 0x87, 0xbe, 0xbe, 0xdb, 0xf1, 0x4d, 0xbd, 0x60, 0x2f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x73, 0x1f, 0x10, 0x96, 0xfc, 0x02, 0x00, 0x00,
}

func (m *Info) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Info) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Info) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.EMail) > 0 {
		i -= len(m.EMail)
		copy(dAtA[i:], m.EMail)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.EMail)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Provider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Provider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Provider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProvider(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProvider(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.HostURI) > 0 {
		i -= len(m.HostURI)
		copy(dAtA[i:], m.HostURI)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.HostURI)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProvider(dAtA []byte, offset int, v uint64) int {
	offset -= sovProvider(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Info) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EMail)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	return n
}

func (m *Provider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.HostURI)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovProvider(uint64(l))
		}
	}
	l = m.Info.Size()
	n += 1 + l + sovProvider(uint64(l))
	return n
}

func sovProvider(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProvider(x uint64) (n int) {
	return sovProvider(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Info) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProvider
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
			return fmt.Errorf("proto: Info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EMail", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EMail = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProvider(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProvider
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
func (m *Provider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProvider
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
			return fmt.Errorf("proto: Provider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Provider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostURI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostURI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, v1.Attribute{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProvider(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProvider
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
func skipProvider(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProvider
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
					return 0, ErrIntOverflowProvider
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
					return 0, ErrIntOverflowProvider
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
				return 0, ErrInvalidLengthProvider
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProvider
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProvider
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProvider        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProvider          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProvider = fmt.Errorf("proto: unexpected end of group")
)
