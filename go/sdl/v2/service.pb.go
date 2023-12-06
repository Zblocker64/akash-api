// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/sdl/v2/service.proto

package v2

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Service struct {
	Image        string         `protobuf:"bytes,1,opt,name=image,proto3" json:"image" yaml:"image"`
	Command      []string       `protobuf:"bytes,2,rep,name=command,proto3" json:"command" yaml:"command"`
	Args         []string       `protobuf:"bytes,3,rep,name=args,proto3" json:"args" yaml:"args"`
	Env          []string       `protobuf:"bytes,4,rep,name=env,proto3" json:"env" yaml:"env"`
	Expose       Exposes        `protobuf:"bytes,5,rep,name=expose,proto3,castrepeated=Exposes" json:"expose" yaml:"expose"`
	Dependencies Dependencies   `protobuf:"bytes,6,rep,name=dependencies,proto3,castrepeated=Dependencies" json:"dependencies" yaml:"dependencies"`
	Params       *ServiceParams `protobuf:"bytes,7,opt,name=params,proto3" json:"params" yaml:"params"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1746d0a871f7bbe, []int{0}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Service.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return m.Size()
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Service) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Service) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Service) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Service) GetExpose() Exposes {
	if m != nil {
		return m.Expose
	}
	return nil
}

func (m *Service) GetDependencies() Dependencies {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *Service) GetParams() *ServiceParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "akash.sdl.v2.Service")
}

func init() { proto.RegisterFile("akash/sdl/v2/service.proto", fileDescriptor_f1746d0a871f7bbe) }

var fileDescriptor_f1746d0a871f7bbe = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xb1, 0x8e, 0xd3, 0x40,
	0x10, 0x8d, 0x71, 0x62, 0xeb, 0x36, 0x81, 0xc2, 0x5c, 0x61, 0x02, 0x78, 0x82, 0x25, 0x24, 0x17,
	0x60, 0x4b, 0x39, 0x0a, 0x44, 0x87, 0xb9, 0x43, 0x42, 0x34, 0xc8, 0x74, 0xd7, 0x20, 0x27, 0x5e,
	0xf9, 0xac, 0x8b, 0xbd, 0x96, 0x37, 0x18, 0xee, 0x13, 0xe8, 0xf8, 0x0e, 0xbe, 0x80, 0x4f, 0xb8,
	0xf2, 0x4a, 0xaa, 0x01, 0x39, 0x9d, 0x4b, 0x7f, 0x01, 0xf2, 0xee, 0x5e, 0x88, 0x25, 0xba, 0x7d,
	0xef, 0xcd, 0x9b, 0xb7, 0xb3, 0xb3, 0x64, 0x1e, 0x5f, 0xc6, 0xfc, 0x22, 0xe0, 0xc9, 0x26, 0xa8,
	0x97, 0x01, 0xa7, 0x55, 0x9d, 0xad, 0xa9, 0x5f, 0x56, 0x6c, 0xcb, 0xac, 0x99, 0xd0, 0x7c, 0x9e,
	0x6c, 0xfc, 0x7a, 0x39, 0x3f, 0x4e, 0x59, 0xca, 0x84, 0x10, 0xf4, 0x27, 0x59, 0x33, 0x7f, 0x3c,
	0xf0, 0x27, 0xb4, 0xa4, 0x45, 0x42, 0x8b, 0xf5, 0x95, 0x92, 0x1f, 0x0c, 0x64, 0xfa, 0xb5, 0x64,
	0x5c, 0x75, 0x9f, 0x3f, 0xf9, 0x5f, 0xf2, 0xa7, 0x32, 0xae, 0xe2, 0x9c, 0xcb, 0x12, 0xf7, 0xe7,
	0x98, 0x98, 0x1f, 0xa5, 0x60, 0xbd, 0x24, 0x93, 0x2c, 0x8f, 0x53, 0x6a, 0x6b, 0x0b, 0xcd, 0x3b,
	0x0a, 0xdd, 0x06, 0x61, 0xf2, 0xae, 0x27, 0x5a, 0x04, 0xa9, 0x74, 0x08, 0xb3, 0xab, 0x38, 0xdf,
	0xbc, 0x72, 0x05, 0x74, 0x23, 0x49, 0x5b, 0x21, 0x31, 0xd7, 0x2c, 0xcf, 0xe3, 0x22, 0xb1, 0xef,
	0x2c, 0x74, 0xef, 0x28, 0xf4, 0x1a, 0x04, 0xf3, 0x8d, 0xa4, 0x5a, 0x84, 0x5b, 0xb5, 0x43, 0xb8,
	0x27, 0xfd, 0x8a, 0x70, 0xa3, 0x5b, 0xc9, 0x3a, 0x21, 0xe3, 0xb8, 0x4a, 0xb9, 0xad, 0x8b, 0x06,
	0xd0, 0x20, 0x8c, 0x5f, 0x57, 0x29, 0x6f, 0x11, 0x04, 0xdf, 0x21, 0x4c, 0xa5, 0xb5, 0x47, 0x6e,
	0x24, 0x48, 0xcb, 0x27, 0x3a, 0x2d, 0x6a, 0x7b, 0x2c, 0x3c, 0x8f, 0x1a, 0x04, 0xfd, 0xac, 0xa8,
	0x5b, 0x84, 0x9e, 0xed, 0x10, 0x88, 0x74, 0xd0, 0xa2, 0x76, 0xa3, 0x9e, 0xb2, 0x56, 0xc4, 0x90,
	0x2f, 0x64, 0x4f, 0x16, 0xba, 0x37, 0x5d, 0x1e, 0xfb, 0x87, 0x0b, 0xf0, 0xcf, 0x84, 0x16, 0xbe,
	0xb8, 0x46, 0x18, 0x35, 0x08, 0x86, 0xc4, 0x2d, 0x82, 0x72, 0x75, 0x08, 0x77, 0x55, 0x4b, 0x81,
	0xdd, 0x1f, 0xbf, 0xc1, 0x94, 0x45, 0x3c, 0x52, 0x35, 0xd6, 0x37, 0x8d, 0xcc, 0xf6, 0x5b, 0xca,
	0x28, 0xb7, 0x0d, 0x11, 0x65, 0x0f, 0xa3, 0x4e, 0xf7, 0x7b, 0x0c, 0xdf, 0xab, 0xb8, 0xd9, 0xe9,
	0x81, 0xab, 0x45, 0x18, 0x74, 0xe9, 0x10, 0xee, 0xcb, 0xe8, 0x43, 0xb6, 0xbf, 0xc0, 0xc0, 0x16,
	0x0d, 0x4c, 0xd6, 0x39, 0x31, 0xe4, 0xba, 0x6d, 0x73, 0xa1, 0x79, 0xd3, 0xe5, 0xc3, 0xe1, 0x25,
	0xd4, 0xe6, 0x3f, 0x88, 0x92, 0xf0, 0x69, 0x3f, 0xb2, 0x3c, 0xf7, 0x23, 0x4b, 0xe3, 0xbf, 0x91,
	0x25, 0x76, 0x23, 0x25, 0x84, 0x6f, 0xaf, 0x1b, 0x47, 0xbb, 0x69, 0x1c, 0xed, 0x4f, 0xe3, 0x68,
	0xdf, 0x77, 0xce, 0xe8, 0x66, 0xe7, 0x8c, 0x7e, 0xed, 0x9c, 0xd1, 0xf9, 0xb3, 0x34, 0xdb, 0x5e,
	0x7c, 0x5e, 0xf9, 0x6b, 0x96, 0x07, 0x22, 0xef, 0x79, 0x41, 0xb7, 0x5f, 0x58, 0x75, 0xa9, 0x50,
	0x5c, 0x66, 0x41, 0xca, 0xd4, 0xbf, 0x5c, 0x19, 0xe2, 0x27, 0x9e, 0xfc, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xcc, 0x62, 0x19, 0xb6, 0x28, 0x03, 0x00, 0x00,
}

func (m *Service) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Service) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Service) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Dependencies) > 0 {
		for iNdEx := len(m.Dependencies) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Dependencies[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Expose) > 0 {
		for iNdEx := len(m.Expose) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Expose[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Env) > 0 {
		for iNdEx := len(m.Env) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Env[iNdEx])
			copy(dAtA[i:], m.Env[iNdEx])
			i = encodeVarintService(dAtA, i, uint64(len(m.Env[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Args) > 0 {
		for iNdEx := len(m.Args) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Args[iNdEx])
			copy(dAtA[i:], m.Args[iNdEx])
			i = encodeVarintService(dAtA, i, uint64(len(m.Args[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Command) > 0 {
		for iNdEx := len(m.Command) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Command[iNdEx])
			copy(dAtA[i:], m.Command[iNdEx])
			i = encodeVarintService(dAtA, i, uint64(len(m.Command[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Image) > 0 {
		i -= len(m.Image)
		copy(dAtA[i:], m.Image)
		i = encodeVarintService(dAtA, i, uint64(len(m.Image)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Service) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if len(m.Command) > 0 {
		for _, s := range m.Command {
			l = len(s)
			n += 1 + l + sovService(uint64(l))
		}
	}
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			l = len(s)
			n += 1 + l + sovService(uint64(l))
		}
	}
	if len(m.Env) > 0 {
		for _, s := range m.Env {
			l = len(s)
			n += 1 + l + sovService(uint64(l))
		}
	}
	if len(m.Expose) > 0 {
		for _, e := range m.Expose {
			l = e.Size()
			n += 1 + l + sovService(uint64(l))
		}
	}
	if len(m.Dependencies) > 0 {
		for _, e := range m.Dependencies {
			l = e.Size()
			n += 1 + l + sovService(uint64(l))
		}
	}
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Service) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: Service: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Service: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Command", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Command = append(m.Command, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = append(m.Args, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Env", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Env = append(m.Env, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expose", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Expose = append(m.Expose, Expose{})
			if err := m.Expose[len(m.Expose)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dependencies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dependencies = append(m.Dependencies, Dependency{})
			if err := m.Dependencies[len(m.Dependencies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = &ServiceParams{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
