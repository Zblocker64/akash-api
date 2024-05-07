// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/market/v1/event.proto

package v1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
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

type EventOrderCreated struct {
	ID OrderID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
}

func (m *EventOrderCreated) Reset()         { *m = EventOrderCreated{} }
func (m *EventOrderCreated) String() string { return proto.CompactTextString(m) }
func (*EventOrderCreated) ProtoMessage()    {}
func (*EventOrderCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_95cc31d2d9808e8a, []int{0}
}
func (m *EventOrderCreated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventOrderCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventOrderCreated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventOrderCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventOrderCreated.Merge(m, src)
}
func (m *EventOrderCreated) XXX_Size() int {
	return m.Size()
}
func (m *EventOrderCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventOrderCreated.DiscardUnknown(m)
}

var xxx_messageInfo_EventOrderCreated proto.InternalMessageInfo

func (m *EventOrderCreated) GetID() OrderID {
	if m != nil {
		return m.ID
	}
	return OrderID{}
}

type EventOrderClosed struct {
	ID OrderID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
}

func (m *EventOrderClosed) Reset()         { *m = EventOrderClosed{} }
func (m *EventOrderClosed) String() string { return proto.CompactTextString(m) }
func (*EventOrderClosed) ProtoMessage()    {}
func (*EventOrderClosed) Descriptor() ([]byte, []int) {
	return fileDescriptor_95cc31d2d9808e8a, []int{1}
}
func (m *EventOrderClosed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventOrderClosed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventOrderClosed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventOrderClosed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventOrderClosed.Merge(m, src)
}
func (m *EventOrderClosed) XXX_Size() int {
	return m.Size()
}
func (m *EventOrderClosed) XXX_DiscardUnknown() {
	xxx_messageInfo_EventOrderClosed.DiscardUnknown(m)
}

var xxx_messageInfo_EventOrderClosed proto.InternalMessageInfo

func (m *EventOrderClosed) GetID() OrderID {
	if m != nil {
		return m.ID
	}
	return OrderID{}
}

type EventBidCreated struct {
	ID    BidID         `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
	Price types.DecCoin `protobuf:"bytes,3,opt,name=price,proto3" json:"price" yaml:"price"`
}

func (m *EventBidCreated) Reset()         { *m = EventBidCreated{} }
func (m *EventBidCreated) String() string { return proto.CompactTextString(m) }
func (*EventBidCreated) ProtoMessage()    {}
func (*EventBidCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_95cc31d2d9808e8a, []int{2}
}
func (m *EventBidCreated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBidCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBidCreated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBidCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBidCreated.Merge(m, src)
}
func (m *EventBidCreated) XXX_Size() int {
	return m.Size()
}
func (m *EventBidCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBidCreated.DiscardUnknown(m)
}

var xxx_messageInfo_EventBidCreated proto.InternalMessageInfo

func (m *EventBidCreated) GetID() BidID {
	if m != nil {
		return m.ID
	}
	return BidID{}
}

func (m *EventBidCreated) GetPrice() types.DecCoin {
	if m != nil {
		return m.Price
	}
	return types.DecCoin{}
}

type EventBidClosed struct {
	ID BidID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
}

func (m *EventBidClosed) Reset()         { *m = EventBidClosed{} }
func (m *EventBidClosed) String() string { return proto.CompactTextString(m) }
func (*EventBidClosed) ProtoMessage()    {}
func (*EventBidClosed) Descriptor() ([]byte, []int) {
	return fileDescriptor_95cc31d2d9808e8a, []int{3}
}
func (m *EventBidClosed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBidClosed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBidClosed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBidClosed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBidClosed.Merge(m, src)
}
func (m *EventBidClosed) XXX_Size() int {
	return m.Size()
}
func (m *EventBidClosed) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBidClosed.DiscardUnknown(m)
}

var xxx_messageInfo_EventBidClosed proto.InternalMessageInfo

func (m *EventBidClosed) GetID() BidID {
	if m != nil {
		return m.ID
	}
	return BidID{}
}

type EventLeaseCreated struct {
	ID    LeaseID       `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
	Price types.DecCoin `protobuf:"bytes,3,opt,name=price,proto3" json:"price" yaml:"price"`
}

func (m *EventLeaseCreated) Reset()         { *m = EventLeaseCreated{} }
func (m *EventLeaseCreated) String() string { return proto.CompactTextString(m) }
func (*EventLeaseCreated) ProtoMessage()    {}
func (*EventLeaseCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_95cc31d2d9808e8a, []int{4}
}
func (m *EventLeaseCreated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventLeaseCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventLeaseCreated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventLeaseCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventLeaseCreated.Merge(m, src)
}
func (m *EventLeaseCreated) XXX_Size() int {
	return m.Size()
}
func (m *EventLeaseCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventLeaseCreated.DiscardUnknown(m)
}

var xxx_messageInfo_EventLeaseCreated proto.InternalMessageInfo

func (m *EventLeaseCreated) GetID() LeaseID {
	if m != nil {
		return m.ID
	}
	return LeaseID{}
}

func (m *EventLeaseCreated) GetPrice() types.DecCoin {
	if m != nil {
		return m.Price
	}
	return types.DecCoin{}
}

type EventLeaseClosed struct {
	ID LeaseID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
}

func (m *EventLeaseClosed) Reset()         { *m = EventLeaseClosed{} }
func (m *EventLeaseClosed) String() string { return proto.CompactTextString(m) }
func (*EventLeaseClosed) ProtoMessage()    {}
func (*EventLeaseClosed) Descriptor() ([]byte, []int) {
	return fileDescriptor_95cc31d2d9808e8a, []int{5}
}
func (m *EventLeaseClosed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventLeaseClosed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventLeaseClosed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventLeaseClosed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventLeaseClosed.Merge(m, src)
}
func (m *EventLeaseClosed) XXX_Size() int {
	return m.Size()
}
func (m *EventLeaseClosed) XXX_DiscardUnknown() {
	xxx_messageInfo_EventLeaseClosed.DiscardUnknown(m)
}

var xxx_messageInfo_EventLeaseClosed proto.InternalMessageInfo

func (m *EventLeaseClosed) GetID() LeaseID {
	if m != nil {
		return m.ID
	}
	return LeaseID{}
}

func init() {
	proto.RegisterType((*EventOrderCreated)(nil), "akash.market.v1.EventOrderCreated")
	proto.RegisterType((*EventOrderClosed)(nil), "akash.market.v1.EventOrderClosed")
	proto.RegisterType((*EventBidCreated)(nil), "akash.market.v1.EventBidCreated")
	proto.RegisterType((*EventBidClosed)(nil), "akash.market.v1.EventBidClosed")
	proto.RegisterType((*EventLeaseCreated)(nil), "akash.market.v1.EventLeaseCreated")
	proto.RegisterType((*EventLeaseClosed)(nil), "akash.market.v1.EventLeaseClosed")
}

func init() { proto.RegisterFile("akash/market/v1/event.proto", fileDescriptor_95cc31d2d9808e8a) }

var fileDescriptor_95cc31d2d9808e8a = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x4d, 0xf2, 0xfd, 0xc0, 0x97, 0x4f, 0xac, 0x16, 0x91, 0x5a, 0x6b, 0x46, 0xb2, 0x72, 0x35,
	0x43, 0x14, 0x5c, 0x74, 0x99, 0xb6, 0x48, 0x41, 0x10, 0xbb, 0x14, 0x44, 0x26, 0x99, 0xa1, 0x0e,
	0x69, 0x33, 0x25, 0x09, 0x01, 0xdf, 0xc2, 0x47, 0xf0, 0x15, 0x5c, 0xf8, 0x0e, 0x5d, 0x76, 0xe9,
	0x2a, 0x48, 0xba, 0x91, 0x2e, 0xfb, 0x04, 0x92, 0x99, 0xa9, 0x96, 0x16, 0x21, 0x8b, 0x82, 0xbb,
	0xdc, 0x9c, 0x7b, 0xef, 0x39, 0x77, 0xce, 0xbd, 0xe6, 0x21, 0x0e, 0x70, 0x7c, 0x8f, 0x86, 0x38,
	0x0a, 0x68, 0x82, 0x52, 0x07, 0xd1, 0x94, 0x86, 0x09, 0x1c, 0x45, 0x3c, 0xe1, 0xd5, 0x8a, 0x00,
	0xa1, 0x04, 0x61, 0xea, 0xd4, 0xf7, 0xfa, 0xbc, 0xcf, 0x05, 0x86, 0x8a, 0x2f, 0x99, 0x56, 0x3f,
	0xf0, 0x79, 0x3c, 0xe4, 0xf1, 0x9d, 0x04, 0x64, 0xa0, 0x20, 0x4b, 0x46, 0xc8, 0xc3, 0x31, 0x45,
	0xa9, 0xe3, 0xd1, 0x04, 0x3b, 0xc8, 0xe7, 0x2c, 0x5c, 0x94, 0xae, 0xd2, 0x7b, 0x8c, 0x28, 0x68,
	0x4d, 0x19, 0x8f, 0x08, 0x8d, 0xbe, 0x03, 0x07, 0x14, 0xc7, 0x54, 0x82, 0xb6, 0x67, 0xee, 0x76,
	0x8a, 0x29, 0xae, 0x8a, 0x82, 0x56, 0x44, 0x71, 0x42, 0x49, 0xf5, 0xc2, 0x34, 0x18, 0xa9, 0xe9,
	0xc7, 0xfa, 0xc9, 0xff, 0xd3, 0x1a, 0x5c, 0x19, 0x0c, 0x8a, 0xd4, 0x6e, 0xdb, 0x3d, 0x1a, 0x67,
	0x40, 0xcb, 0x33, 0x60, 0x74, 0xdb, 0xb3, 0x0c, 0x18, 0x8c, 0xcc, 0x33, 0xf0, 0xef, 0x01, 0x0f,
	0x07, 0x4d, 0x9b, 0x11, 0xbb, 0x67, 0x30, 0xd2, 0xfc, 0xfd, 0xfe, 0x04, 0x34, 0x1b, 0x9b, 0x3b,
	0x4b, 0x1c, 0x03, 0x1e, 0x6f, 0x9e, 0xe2, 0x59, 0x37, 0x2b, 0x82, 0xc3, 0x65, 0x64, 0x31, 0x45,
	0x67, 0x89, 0x62, 0x7f, 0x8d, 0xc2, 0x65, 0xa4, 0x14, 0x41, 0xf5, 0xda, 0xfc, 0x33, 0x8a, 0x98,
	0x4f, 0x6b, 0xbf, 0x44, 0xa7, 0x06, 0x54, 0xa6, 0x15, 0x36, 0x41, 0x65, 0x13, 0x6c, 0x53, 0xbf,
	0xc5, 0x59, 0x28, 0xfb, 0xcd, 0x32, 0x20, 0x4b, 0xe6, 0x19, 0xd8, 0x92, 0xcd, 0x44, 0x68, 0xf7,
	0xe4, 0x6f, 0xa5, 0xf9, 0xd6, 0xdc, 0xfe, 0x94, 0x2c, 0x1f, 0x65, 0x33, 0x8a, 0x55, 0xfb, 0x17,
	0x5d, 0x59, 0x7b, 0x59, 0xd8, 0x5d, 0xce, 0x5a, 0x91, 0xfa, 0xc3, 0xcf, 0xb2, 0xd8, 0x16, 0x29,
	0xbb, 0xcc, 0xb6, 0x94, 0x57, 0x2d, 0x29, 0xdc, 0xf3, 0x71, 0x6e, 0xe9, 0x93, 0xdc, 0xd2, 0xdf,
	0x72, 0x4b, 0x7f, 0x9c, 0x5a, 0xda, 0x64, 0x6a, 0x69, 0xaf, 0x53, 0x4b, 0xbb, 0x69, 0x8c, 0x82,
	0x3e, 0xc4, 0x41, 0x02, 0x59, 0x71, 0xb5, 0x28, 0xe4, 0x84, 0x7e, 0x1d, 0x8e, 0xf7, 0x57, 0xdc,
	0xcc, 0xd9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0xe0, 0x00, 0x64, 0x09, 0x04, 0x00, 0x00,
}

func (m *EventOrderCreated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventOrderCreated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventOrderCreated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *EventOrderClosed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventOrderClosed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventOrderClosed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *EventBidCreated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBidCreated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBidCreated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *EventBidClosed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBidClosed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBidClosed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *EventLeaseCreated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventLeaseCreated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventLeaseCreated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *EventLeaseClosed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventLeaseClosed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventLeaseClosed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventOrderCreated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func (m *EventOrderClosed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func (m *EventBidCreated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovEvent(uint64(l))
	l = m.Price.Size()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func (m *EventBidClosed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func (m *EventLeaseCreated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovEvent(uint64(l))
	l = m.Price.Size()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func (m *EventLeaseClosed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventOrderCreated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventOrderCreated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventOrderCreated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventOrderClosed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventOrderClosed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventOrderClosed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventBidCreated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventBidCreated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBidCreated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventBidClosed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventBidClosed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBidClosed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventLeaseCreated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventLeaseCreated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventLeaseCreated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventLeaseClosed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventLeaseClosed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventLeaseClosed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
