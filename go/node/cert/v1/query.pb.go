// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/cert/v1/query.proto

package v1

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// CertificateResponse contains a single X509 certificate and its serial number
type CertificateResponse struct {
	Certificate Certificate `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate" yaml:"certificate"`
	Serial      string      `protobuf:"bytes,2,opt,name=serial,proto3" json:"serial" yaml:"serial"`
}

func (m *CertificateResponse) Reset()         { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()    {}
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e5d14d61992fa41, []int{0}
}
func (m *CertificateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CertificateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateResponse.Merge(m, src)
}
func (m *CertificateResponse) XXX_Size() int {
	return m.Size()
}
func (m *CertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateResponse proto.InternalMessageInfo

func (m *CertificateResponse) GetCertificate() Certificate {
	if m != nil {
		return m.Certificate
	}
	return Certificate{}
}

func (m *CertificateResponse) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

// QueryDeploymentsRequest is request type for the Query/Deployments RPC method
type QueryCertificatesRequest struct {
	Filter     CertificateFilter  `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter"`
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCertificatesRequest) Reset()         { *m = QueryCertificatesRequest{} }
func (m *QueryCertificatesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCertificatesRequest) ProtoMessage()    {}
func (*QueryCertificatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e5d14d61992fa41, []int{1}
}
func (m *QueryCertificatesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCertificatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCertificatesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCertificatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCertificatesRequest.Merge(m, src)
}
func (m *QueryCertificatesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCertificatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCertificatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCertificatesRequest proto.InternalMessageInfo

func (m *QueryCertificatesRequest) GetFilter() CertificateFilter {
	if m != nil {
		return m.Filter
	}
	return CertificateFilter{}
}

func (m *QueryCertificatesRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryCertificatesResponse is response type for the Query/Certificates RPC method
type QueryCertificatesResponse struct {
	Certificates CertificatesResponse `protobuf:"bytes,1,rep,name=certificates,proto3,castrepeated=CertificatesResponse" json:"certificates"`
	Pagination   *query.PageResponse  `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCertificatesResponse) Reset()         { *m = QueryCertificatesResponse{} }
func (m *QueryCertificatesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCertificatesResponse) ProtoMessage()    {}
func (*QueryCertificatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e5d14d61992fa41, []int{2}
}
func (m *QueryCertificatesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCertificatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCertificatesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCertificatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCertificatesResponse.Merge(m, src)
}
func (m *QueryCertificatesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCertificatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCertificatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCertificatesResponse proto.InternalMessageInfo

func (m *QueryCertificatesResponse) GetCertificates() CertificatesResponse {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func (m *QueryCertificatesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*CertificateResponse)(nil), "akash.cert.v1.CertificateResponse")
	proto.RegisterType((*QueryCertificatesRequest)(nil), "akash.cert.v1.QueryCertificatesRequest")
	proto.RegisterType((*QueryCertificatesResponse)(nil), "akash.cert.v1.QueryCertificatesResponse")
}

func init() { proto.RegisterFile("akash/cert/v1/query.proto", fileDescriptor_2e5d14d61992fa41) }

var fileDescriptor_2e5d14d61992fa41 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xce, 0x15, 0x88, 0xc4, 0xa5, 0x5d, 0x8e, 0x0e, 0x6e, 0x52, 0xd9, 0xd6, 0x0d, 0xd4, 0x30,
	0xdc, 0x29, 0xa9, 0x58, 0x18, 0x18, 0x8c, 0x54, 0x56, 0xf0, 0xc8, 0x76, 0x49, 0xae, 0x87, 0x15,
	0xd7, 0xe7, 0xfa, 0xae, 0x91, 0xba, 0xb2, 0xb2, 0x20, 0x21, 0xfe, 0x00, 0x23, 0x2b, 0x7f, 0xa2,
	0x62, 0xaa, 0xc4, 0xc2, 0x64, 0x50, 0xc2, 0xd4, 0xb1, 0xbf, 0x00, 0xd9, 0x77, 0x51, 0xcf, 0x90,
	0x28, 0x9b, 0xfd, 0xbe, 0xef, 0x7b, 0xef, 0x7b, 0xef, 0x3b, 0x78, 0xc0, 0x66, 0x4c, 0xbd, 0xa3,
	0x13, 0x5e, 0x6a, 0x3a, 0x1f, 0xd2, 0xf3, 0x0b, 0x5e, 0x5e, 0x92, 0xa2, 0x94, 0x5a, 0xa2, 0xbd,
	0x06, 0x22, 0x35, 0x44, 0xe6, 0xc3, 0xfe, 0xbe, 0x90, 0x42, 0x36, 0x08, 0xad, 0xbf, 0x0c, 0xa9,
	0x7f, 0x28, 0xa4, 0x14, 0x19, 0xa7, 0xac, 0x48, 0x29, 0xcb, 0x73, 0xa9, 0x99, 0x4e, 0x65, 0xae,
	0x2c, 0xfa, 0x74, 0x22, 0xd5, 0x99, 0x54, 0x74, 0xcc, 0x14, 0x37, 0xbd, 0xe9, 0x7c, 0x38, 0xe6,
	0x9a, 0x0d, 0x69, 0xc1, 0x44, 0x9a, 0x37, 0x64, 0xcb, 0xf5, 0xda, 0x4e, 0x9a, 0xb1, 0x06, 0x19,
	0xb4, 0x91, 0xd3, 0x34, 0xd3, 0xbc, 0xb4, 0x23, 0xf0, 0x37, 0x00, 0x1f, 0xbd, 0xe4, 0xa5, 0x4e,
	0x4f, 0xd3, 0x09, 0xd3, 0x3c, 0xe1, 0xaa, 0x90, 0xb9, 0xe2, 0x48, 0xc0, 0xde, 0xe4, 0xae, 0xec,
	0x81, 0x10, 0x44, 0xbd, 0x51, 0x9f, 0xb4, 0x76, 0x22, 0x8e, 0x30, 0x7e, 0x72, 0x55, 0x05, 0x9d,
	0x9b, 0x2a, 0x70, 0x65, 0xb7, 0x55, 0x80, 0x2e, 0xd9, 0x59, 0xf6, 0x1c, 0x3b, 0x45, 0x9c, 0xb8,
	0x14, 0x74, 0x0c, 0xbb, 0x8a, 0x97, 0x29, 0xcb, 0xbc, 0x9d, 0x10, 0x44, 0x0f, 0xe3, 0xc1, 0x4d,
	0x15, 0xd8, 0xca, 0x6d, 0x15, 0xec, 0x19, 0xb9, 0xf9, 0xc7, 0x89, 0x05, 0xf0, 0x17, 0x00, 0xbd,
	0x37, 0xf5, 0x3d, 0x1c, 0x07, 0x2a, 0xe1, 0xe7, 0x17, 0x5c, 0x69, 0xf4, 0x02, 0x76, 0xcd, 0x8e,
	0xd6, 0x75, 0xb8, 0xd9, 0xf5, 0x49, 0xc3, 0x8b, 0xef, 0xd7, 0xde, 0x13, 0xab, 0x42, 0x27, 0x10,
	0xde, 0x5d, 0xb7, 0x71, 0xd5, 0x1b, 0x3d, 0x26, 0x26, 0x0a, 0x52, 0x47, 0x41, 0x4c, 0xcc, 0x36,
	0x0a, 0xf2, 0x9a, 0x09, 0x6e, 0x67, 0x27, 0x8e, 0x12, 0x7f, 0x07, 0xf0, 0x60, 0x8d, 0x49, 0x7b,
	0xe0, 0x29, 0xdc, 0x75, 0xce, 0xa0, 0x3c, 0x10, 0xde, 0x8b, 0x7a, 0x23, 0xbc, 0xd9, 0xeb, 0x4a,
	0x19, 0x1f, 0xd6, 0x6e, 0xbf, 0xfe, 0x0a, 0xf6, 0xd7, 0xf5, 0x4d, 0x5a, 0x5d, 0xd1, 0xab, 0x35,
	0xbb, 0x1c, 0x6d, 0xdd, 0xc5, 0xb6, 0x72, 0xa4, 0xa3, 0xcf, 0x00, 0x3e, 0x68, 0x96, 0x41, 0x1f,
	0x00, 0xdc, 0x75, 0x27, 0xa3, 0xa3, 0x7f, 0x3c, 0x6f, 0x0a, 0xa6, 0x1f, 0x6d, 0x27, 0x9a, 0xc9,
	0x38, 0x7a, 0xff, 0xe3, 0xcf, 0xa7, 0x1d, 0x8c, 0x42, 0xfa, 0xff, 0xab, 0x5e, 0x91, 0x69, 0x96,
	0x2a, 0x1d, 0x3f, 0xbb, 0x5a, 0xf8, 0xe0, 0x7a, 0xe1, 0x83, 0xdf, 0x0b, 0x1f, 0x7c, 0x5c, 0xfa,
	0x9d, 0xeb, 0xa5, 0xdf, 0xf9, 0xb9, 0xf4, 0x3b, 0x6f, 0x07, 0xc5, 0x4c, 0x10, 0x36, 0xd3, 0x64,
	0xca, 0xe7, 0x54, 0x48, 0x9a, 0xcb, 0x29, 0x5f, 0x35, 0x1a, 0x77, 0x9b, 0xd7, 0x7f, 0xfc, 0x37,
	0x00, 0x00, 0xff, 0xff, 0xc5, 0xd6, 0xea, 0x4d, 0xc0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Certificates queries certificates
	Certificates(ctx context.Context, in *QueryCertificatesRequest, opts ...grpc.CallOption) (*QueryCertificatesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Certificates(ctx context.Context, in *QueryCertificatesRequest, opts ...grpc.CallOption) (*QueryCertificatesResponse, error) {
	out := new(QueryCertificatesResponse)
	err := c.cc.Invoke(ctx, "/akash.cert.v1.Query/Certificates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Certificates queries certificates
	Certificates(context.Context, *QueryCertificatesRequest) (*QueryCertificatesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Certificates(ctx context.Context, req *QueryCertificatesRequest) (*QueryCertificatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Certificates not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Certificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCertificatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Certificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.cert.v1.Query/Certificates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Certificates(ctx, req.(*QueryCertificatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "akash.cert.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Certificates",
			Handler:    _Query_Certificates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "akash/cert/v1/query.proto",
}

func (m *CertificateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CertificateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CertificateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Serial) > 0 {
		i -= len(m.Serial)
		copy(dAtA[i:], m.Serial)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Serial)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Certificate.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryCertificatesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCertificatesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCertificatesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Filter.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryCertificatesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCertificatesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCertificatesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Certificates) > 0 {
		for iNdEx := len(m.Certificates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Certificates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CertificateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Certificate.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = len(m.Serial)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCertificatesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Filter.Size()
	n += 1 + l + sovQuery(uint64(l))
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCertificatesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Certificates) > 0 {
		for _, e := range m.Certificates {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CertificateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: CertificateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertificateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Certificate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Serial", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Serial = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCertificatesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCertificatesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCertificatesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Filter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCertificatesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCertificatesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCertificatesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificates = append(m.Certificates, CertificateResponse{})
			if err := m.Certificates[len(m.Certificates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
