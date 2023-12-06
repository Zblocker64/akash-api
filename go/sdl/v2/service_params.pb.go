// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/sdl/v2/service_params.proto

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

type ServiceStorageParams struct {
	Mount    string `protobuf:"bytes,1,opt,name=mount,proto3" json:"mount" yaml:"mount"`
	ReadOnly bool   `protobuf:"varint,2,opt,name=read_only,json=readOnly,proto3" json:"readOnly" yaml:"readOnly"`
}

func (m *ServiceStorageParams) Reset()         { *m = ServiceStorageParams{} }
func (m *ServiceStorageParams) String() string { return proto.CompactTextString(m) }
func (*ServiceStorageParams) ProtoMessage()    {}
func (*ServiceStorageParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_25037b2bfca6b25d, []int{0}
}
func (m *ServiceStorageParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceStorageParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServiceStorageParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ServiceStorageParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceStorageParams.Merge(m, src)
}
func (m *ServiceStorageParams) XXX_Size() int {
	return m.Size()
}
func (m *ServiceStorageParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceStorageParams.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceStorageParams proto.InternalMessageInfo

func (m *ServiceStorageParams) GetMount() string {
	if m != nil {
		return m.Mount
	}
	return ""
}

func (m *ServiceStorageParams) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

type ServiceParams struct {
	Storage map[string]*ServiceStorageParams `protobuf:"bytes,1,rep,name=Storage,proto3" json:"storage" yaml:"storage" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ServiceParams) Reset()         { *m = ServiceParams{} }
func (m *ServiceParams) String() string { return proto.CompactTextString(m) }
func (*ServiceParams) ProtoMessage()    {}
func (*ServiceParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_25037b2bfca6b25d, []int{1}
}
func (m *ServiceParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServiceParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ServiceParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceParams.Merge(m, src)
}
func (m *ServiceParams) XXX_Size() int {
	return m.Size()
}
func (m *ServiceParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceParams.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceParams proto.InternalMessageInfo

func (m *ServiceParams) GetStorage() map[string]*ServiceStorageParams {
	if m != nil {
		return m.Storage
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceStorageParams)(nil), "akash.sdl.v2.ServiceStorageParams")
	proto.RegisterType((*ServiceParams)(nil), "akash.sdl.v2.ServiceParams")
	proto.RegisterMapType((map[string]*ServiceStorageParams)(nil), "akash.sdl.v2.ServiceParams.StorageEntry")
}

func init() { proto.RegisterFile("akash/sdl/v2/service_params.proto", fileDescriptor_25037b2bfca6b25d) }

var fileDescriptor_25037b2bfca6b25d = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xc1, 0x6a, 0xea, 0x40,
	0x14, 0x75, 0x14, 0x9f, 0x3a, 0xfa, 0xde, 0x2b, 0xc1, 0x85, 0xb8, 0xc8, 0xd8, 0x59, 0x05, 0xda,
	0x26, 0x90, 0x6e, 0xc4, 0x65, 0xa0, 0xa5, 0x9b, 0xd2, 0x12, 0x77, 0x5d, 0x54, 0x46, 0x33, 0x44,
	0x31, 0xc9, 0xc8, 0x24, 0xa6, 0xe4, 0x2f, 0xba, 0xee, 0x17, 0x75, 0xe9, 0xb2, 0x9b, 0x0e, 0x25,
	0xee, 0x5c, 0xfa, 0x05, 0x25, 0x99, 0xb1, 0x28, 0x74, 0x77, 0xee, 0xb9, 0xe7, 0x1e, 0xce, 0xbd,
	0x17, 0x9e, 0x93, 0x25, 0x89, 0xe7, 0x56, 0xec, 0x05, 0x56, 0x6a, 0x5b, 0x31, 0xe5, 0xe9, 0x62,
	0x46, 0x27, 0x2b, 0xc2, 0x49, 0x18, 0x9b, 0x2b, 0xce, 0x12, 0xa6, 0x75, 0x4a, 0x89, 0x19, 0x7b,
	0x81, 0x99, 0xda, 0xfd, 0xae, 0xcf, 0x7c, 0x56, 0x36, 0xac, 0x02, 0x49, 0x0d, 0x7e, 0x03, 0xb0,
	0x3b, 0x96, 0xc3, 0xe3, 0x84, 0x71, 0xe2, 0xd3, 0xc7, 0xd2, 0x42, 0x1b, 0xc2, 0x7a, 0xc8, 0xd6,
	0x51, 0xd2, 0x03, 0x03, 0x60, 0xb4, 0x1c, 0x9c, 0x0b, 0x54, 0xbf, 0x2f, 0x88, 0x9d, 0x40, 0xb2,
	0xb3, 0x17, 0xa8, 0x93, 0x91, 0x30, 0x18, 0xe1, 0xb2, 0xc4, 0xae, 0xa4, 0xb5, 0x3b, 0xd8, 0xe2,
	0x94, 0x78, 0x13, 0x16, 0x05, 0x59, 0xaf, 0x3a, 0x00, 0x46, 0xd3, 0xb9, 0xc8, 0x05, 0x6a, 0xba,
	0x94, 0x78, 0x0f, 0x51, 0x90, 0xed, 0x04, 0x6a, 0x72, 0x85, 0xf7, 0x02, 0xfd, 0x97, 0x1e, 0x07,
	0x06, 0xbb, 0x3f, 0x4d, 0xfc, 0x09, 0xe0, 0x5f, 0x15, 0x4e, 0xa5, 0x0a, 0x60, 0x43, 0xc5, 0xec,
	0x81, 0x41, 0xcd, 0x68, 0xdb, 0x86, 0x79, 0xbc, 0xa4, 0x79, 0xa2, 0x36, 0x95, 0xf4, 0x26, 0x4a,
	0x78, 0xe6, 0x18, 0xb9, 0x40, 0x87, 0xe1, 0x9d, 0x40, 0x8d, 0x58, 0xc2, 0xbd, 0x40, 0xff, 0x64,
	0x02, 0x45, 0x60, 0xf7, 0xa0, 0xea, 0x3f, 0xc3, 0xce, 0xb1, 0x85, 0x76, 0x06, 0x6b, 0x4b, 0x9a,
	0xc9, 0x8b, 0xb8, 0x05, 0x2c, 0xae, 0x94, 0x92, 0x60, 0x4d, 0xcb, 0x3d, 0xdb, 0x36, 0xfe, 0x35,
	0xcd, 0xc9, 0x61, 0x5d, 0x39, 0x30, 0xaa, 0x0e, 0x81, 0x73, 0xfb, 0x9e, 0xeb, 0x60, 0x93, 0xeb,
	0xe0, 0x2b, 0xd7, 0xc1, 0xeb, 0x56, 0xaf, 0x6c, 0xb6, 0x7a, 0xe5, 0x63, 0xab, 0x57, 0x9e, 0x2e,
	0xfd, 0x45, 0x32, 0x5f, 0x4f, 0xcd, 0x19, 0x0b, 0xad, 0xd2, 0xf2, 0x2a, 0xa2, 0xc9, 0x0b, 0xe3,
	0x4b, 0x55, 0x91, 0xd5, 0xc2, 0xf2, 0x99, 0xfa, 0xfe, 0xf4, 0x4f, 0xf9, 0xcb, 0xeb, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x7f, 0x3d, 0xe1, 0xce, 0x14, 0x02, 0x00, 0x00,
}

func (m *ServiceStorageParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceStorageParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceStorageParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReadOnly {
		i--
		if m.ReadOnly {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Mount) > 0 {
		i -= len(m.Mount)
		copy(dAtA[i:], m.Mount)
		i = encodeVarintServiceParams(dAtA, i, uint64(len(m.Mount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ServiceParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Storage) > 0 {
		for k := range m.Storage {
			v := m.Storage[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintServiceParams(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintServiceParams(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintServiceParams(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintServiceParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovServiceParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ServiceStorageParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Mount)
	if l > 0 {
		n += 1 + l + sovServiceParams(uint64(l))
	}
	if m.ReadOnly {
		n += 2
	}
	return n
}

func (m *ServiceParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Storage) > 0 {
		for k, v := range m.Storage {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovServiceParams(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovServiceParams(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovServiceParams(uint64(mapEntrySize))
		}
	}
	return n
}

func sovServiceParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozServiceParams(x uint64) (n int) {
	return sovServiceParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ServiceStorageParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceParams
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
			return fmt.Errorf("proto: ServiceStorageParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceStorageParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceParams
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
				return ErrInvalidLengthServiceParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReadOnly", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceParams
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
			m.ReadOnly = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipServiceParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServiceParams
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
func (m *ServiceParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceParams
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
			return fmt.Errorf("proto: ServiceParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceParams
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
				return ErrInvalidLengthServiceParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthServiceParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Storage == nil {
				m.Storage = make(map[string]*ServiceStorageParams)
			}
			var mapkey string
			var mapvalue *ServiceStorageParams
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowServiceParams
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowServiceParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthServiceParams
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthServiceParams
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowServiceParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthServiceParams
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthServiceParams
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &ServiceStorageParams{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipServiceParams(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthServiceParams
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Storage[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServiceParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServiceParams
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
func skipServiceParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowServiceParams
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
					return 0, ErrIntOverflowServiceParams
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
					return 0, ErrIntOverflowServiceParams
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
				return 0, ErrInvalidLengthServiceParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupServiceParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthServiceParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthServiceParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowServiceParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupServiceParams = fmt.Errorf("proto: unexpected end of group")
)
