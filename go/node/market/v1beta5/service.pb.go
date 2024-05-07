// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/market/v1beta5/service.proto

package v1beta5

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

func init() {
	proto.RegisterFile("akash/market/v1beta5/service.proto", fileDescriptor_f1203af46a0757a8)
}

var fileDescriptor_f1203af46a0757a8 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0xcc, 0x4e, 0x2c,
	0xce, 0xd0, 0xcf, 0x4d, 0x2c, 0xca, 0x4e, 0x2d, 0xd1, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34,
	0xd5, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x01, 0xab, 0xd1, 0x83, 0xa8, 0xd1, 0x83, 0xaa, 0x91, 0x12, 0x4f, 0xce, 0x2f, 0xce, 0xcd, 0x2f,
	0xd6, 0xcf, 0x2d, 0x4e, 0xd7, 0x2f, 0x33, 0x04, 0x51, 0x10, 0xe5, 0x52, 0x8a, 0x58, 0x8d, 0x4c,
	0xca, 0x4c, 0x41, 0x28, 0x51, 0xc6, 0xaa, 0x24, 0x27, 0x35, 0xb1, 0x38, 0x15, 0xae, 0xc8, 0xe8,
	0x3d, 0x33, 0x17, 0xb3, 0x6f, 0x71, 0xba, 0x50, 0x34, 0x17, 0xa7, 0x73, 0x51, 0x6a, 0x62, 0x49,
	0xaa, 0x53, 0x66, 0x8a, 0x90, 0x92, 0x1e, 0x36, 0xc7, 0xe8, 0xf9, 0x16, 0xa7, 0xc3, 0xd5, 0x48,
	0x69, 0x11, 0x56, 0x13, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x14, 0xc1, 0xc5, 0xe1,
	0x9c, 0x93, 0x5f, 0x0c, 0x36, 0x5b, 0x11, 0xb7, 0x3e, 0xa8, 0x12, 0x29, 0x4d, 0x82, 0x4a, 0xe0,
	0x26, 0xa7, 0x73, 0xf1, 0x86, 0x67, 0x96, 0x64, 0xa4, 0x14, 0x25, 0x96, 0xfb, 0x80, 0x3c, 0x26,
	0xa4, 0x86, 0x53, 0x2f, 0x8a, 0x3a, 0x29, 0x3d, 0xe2, 0xd4, 0xc1, 0x2d, 0x4a, 0xe4, 0xe2, 0x86,
	0xf8, 0x0b, 0x62, 0x8d, 0x0a, 0x01, 0xdf, 0x43, 0x2c, 0xd1, 0x21, 0x46, 0x15, 0xdc, 0x8a, 0x38,
	0x2e, 0x2e, 0xb0, 0xff, 0x20, 0x36, 0x28, 0xe3, 0x0f, 0x04, 0x88, 0x05, 0xda, 0x44, 0x28, 0x82,
	0x99, 0x2f, 0xc5, 0xda, 0xf0, 0x7c, 0x83, 0x16, 0xa3, 0x93, 0xf5, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c,
	0x37, 0x1e, 0xcb, 0x31, 0x44, 0x29, 0x16, 0x64, 0xa7, 0xeb, 0x25, 0x66, 0x97, 0xe8, 0x65, 0xe6,
	0xeb, 0xa7, 0xe7, 0xeb, 0xe7, 0xe5, 0xa7, 0xa4, 0xa2, 0xa5, 0x9e, 0x24, 0x36, 0x70, 0xaa, 0x31,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x95, 0x43, 0x0a, 0xd2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateBid defines a method to create a bid given proper inputs.
	CreateBid(ctx context.Context, in *MsgCreateBid, opts ...grpc.CallOption) (*MsgCreateBidResponse, error)
	// CloseBid defines a method to close a bid given proper inputs.
	CloseBid(ctx context.Context, in *MsgCloseBid, opts ...grpc.CallOption) (*MsgCloseBidResponse, error)
	// WithdrawLease withdraws accrued funds from the lease payment
	WithdrawLease(ctx context.Context, in *MsgWithdrawLease, opts ...grpc.CallOption) (*MsgWithdrawLeaseResponse, error)
	// CreateLease creates a new lease
	CreateLease(ctx context.Context, in *MsgCreateLease, opts ...grpc.CallOption) (*MsgCreateLeaseResponse, error)
	// CloseLease defines a method to close an order given proper inputs.
	CloseLease(ctx context.Context, in *MsgCloseLease, opts ...grpc.CallOption) (*MsgCloseLeaseResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateBid(ctx context.Context, in *MsgCreateBid, opts ...grpc.CallOption) (*MsgCreateBidResponse, error) {
	out := new(MsgCreateBidResponse)
	err := c.cc.Invoke(ctx, "/akash.market.v1beta5.Msg/CreateBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CloseBid(ctx context.Context, in *MsgCloseBid, opts ...grpc.CallOption) (*MsgCloseBidResponse, error) {
	out := new(MsgCloseBidResponse)
	err := c.cc.Invoke(ctx, "/akash.market.v1beta5.Msg/CloseBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawLease(ctx context.Context, in *MsgWithdrawLease, opts ...grpc.CallOption) (*MsgWithdrawLeaseResponse, error) {
	out := new(MsgWithdrawLeaseResponse)
	err := c.cc.Invoke(ctx, "/akash.market.v1beta5.Msg/WithdrawLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateLease(ctx context.Context, in *MsgCreateLease, opts ...grpc.CallOption) (*MsgCreateLeaseResponse, error) {
	out := new(MsgCreateLeaseResponse)
	err := c.cc.Invoke(ctx, "/akash.market.v1beta5.Msg/CreateLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CloseLease(ctx context.Context, in *MsgCloseLease, opts ...grpc.CallOption) (*MsgCloseLeaseResponse, error) {
	out := new(MsgCloseLeaseResponse)
	err := c.cc.Invoke(ctx, "/akash.market.v1beta5.Msg/CloseLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateBid defines a method to create a bid given proper inputs.
	CreateBid(context.Context, *MsgCreateBid) (*MsgCreateBidResponse, error)
	// CloseBid defines a method to close a bid given proper inputs.
	CloseBid(context.Context, *MsgCloseBid) (*MsgCloseBidResponse, error)
	// WithdrawLease withdraws accrued funds from the lease payment
	WithdrawLease(context.Context, *MsgWithdrawLease) (*MsgWithdrawLeaseResponse, error)
	// CreateLease creates a new lease
	CreateLease(context.Context, *MsgCreateLease) (*MsgCreateLeaseResponse, error)
	// CloseLease defines a method to close an order given proper inputs.
	CloseLease(context.Context, *MsgCloseLease) (*MsgCloseLeaseResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateBid(ctx context.Context, req *MsgCreateBid) (*MsgCreateBidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBid not implemented")
}
func (*UnimplementedMsgServer) CloseBid(ctx context.Context, req *MsgCloseBid) (*MsgCloseBidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseBid not implemented")
}
func (*UnimplementedMsgServer) WithdrawLease(ctx context.Context, req *MsgWithdrawLease) (*MsgWithdrawLeaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawLease not implemented")
}
func (*UnimplementedMsgServer) CreateLease(ctx context.Context, req *MsgCreateLease) (*MsgCreateLeaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLease not implemented")
}
func (*UnimplementedMsgServer) CloseLease(ctx context.Context, req *MsgCloseLease) (*MsgCloseLeaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseLease not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateBid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.market.v1beta5.Msg/CreateBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateBid(ctx, req.(*MsgCreateBid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CloseBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCloseBid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CloseBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.market.v1beta5.Msg/CloseBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CloseBid(ctx, req.(*MsgCloseBid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawLease)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.market.v1beta5.Msg/WithdrawLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawLease(ctx, req.(*MsgWithdrawLease))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateLease)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.market.v1beta5.Msg/CreateLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateLease(ctx, req.(*MsgCreateLease))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CloseLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCloseLease)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CloseLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.market.v1beta5.Msg/CloseLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CloseLease(ctx, req.(*MsgCloseLease))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "akash.market.v1beta5.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBid",
			Handler:    _Msg_CreateBid_Handler,
		},
		{
			MethodName: "CloseBid",
			Handler:    _Msg_CloseBid_Handler,
		},
		{
			MethodName: "WithdrawLease",
			Handler:    _Msg_WithdrawLease_Handler,
		},
		{
			MethodName: "CreateLease",
			Handler:    _Msg_CreateLease_Handler,
		},
		{
			MethodName: "CloseLease",
			Handler:    _Msg_CloseLease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "akash/market/v1beta5/service.proto",
}
