// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/escrow/v1/types.proto

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

// State stores state for an escrow account
type AccountState int32

const (
	// AccountStateInvalid is an invalid state
	AccountStateInvalid AccountState = 0
	// AccountOpen is the state when an account is open
	AccountOpen AccountState = 1
	// AccountClosed is the state when an account is closed
	AccountClosed AccountState = 2
	// AccountOverdrawn is the state when an account is overdrawn
	AccountOverdrawn AccountState = 3
)

var AccountState_name = map[int32]string{
	0: "invalid",
	1: "open",
	2: "closed",
	3: "overdrawn",
}

var AccountState_value = map[string]int32{
	"invalid":   0,
	"open":      1,
	"closed":    2,
	"overdrawn": 3,
}

func (x AccountState) String() string {
	return proto.EnumName(AccountState_name, int32(x))
}

func (AccountState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_50b386ecf43f1d04, []int{0}
}

// Payment State
type FractionalPayment_State int32

const (
	// PaymentStateInvalid is the state when the payment is invalid
	PaymentStateInvalid FractionalPayment_State = 0
	// PaymentStateOpen is the state when the payment is open
	PaymentOpen FractionalPayment_State = 1
	// PaymentStateClosed is the state when the payment is closed
	PaymentClosed FractionalPayment_State = 2
	// PaymentStateOverdrawn is the state when the payment is overdrawn
	PaymentOverdrawn FractionalPayment_State = 3
)

var FractionalPayment_State_name = map[int32]string{
	0: "invalid",
	1: "open",
	2: "closed",
	3: "overdrawn",
}

var FractionalPayment_State_value = map[string]int32{
	"invalid":   0,
	"open":      1,
	"closed":    2,
	"overdrawn": 3,
}

func (x FractionalPayment_State) String() string {
	return proto.EnumName(FractionalPayment_State_name, int32(x))
}

func (FractionalPayment_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_50b386ecf43f1d04, []int{2, 0}
}

// AccountID is the account identifier
type AccountID struct {
	Scope string `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope" yaml:"scope"`
	XID   string `protobuf:"bytes,2,opt,name=xid,proto3" json:"xid" yaml:"xid"`
}

func (m *AccountID) Reset()         { *m = AccountID{} }
func (m *AccountID) String() string { return proto.CompactTextString(m) }
func (*AccountID) ProtoMessage()    {}
func (*AccountID) Descriptor() ([]byte, []int) {
	return fileDescriptor_50b386ecf43f1d04, []int{0}
}
func (m *AccountID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountID.Merge(m, src)
}
func (m *AccountID) XXX_Size() int {
	return m.Size()
}
func (m *AccountID) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountID.DiscardUnknown(m)
}

var xxx_messageInfo_AccountID proto.InternalMessageInfo

func (m *AccountID) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *AccountID) GetXID() string {
	if m != nil {
		return m.XID
	}
	return ""
}

// Account stores state for an escrow account
type Account struct {
	// unique identifier for this escrow account
	ID AccountID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
	// bech32 encoded account address of the owner of this escrow account
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	// current state of this escrow account
	State AccountState `protobuf:"varint,3,opt,name=state,proto3,enum=akash.escrow.v1.AccountState" json:"state" yaml:"state"`
	// unspent coins received from the owner's wallet
	Balance types.DecCoin `protobuf:"bytes,4,opt,name=balance,proto3" json:"balance" yaml:"balance"`
	// total coins spent by this account
	Transferred types.DecCoin `protobuf:"bytes,5,opt,name=transferred,proto3" json:"transferred" yaml:"transferred"`
	// block height at which this account was last settled
	SettledAt int64 `protobuf:"varint,6,opt,name=settled_at,json=settledAt,proto3" json:"settledAt" yaml:"settledAt"`
	// bech32 encoded account address of the depositor.
	// If depositor is same as the owner, then any incoming coins are added to the Balance.
	// If depositor isn't same as the owner, then any incoming coins are added to the Funds.
	Depositor string `protobuf:"bytes,7,opt,name=depositor,proto3" json:"depositor" yaml:"depositor"`
	// Funds are unspent coins received from the (non-Owner) Depositor's wallet.
	// If there are any funds, they should be spent before spending the Balance.
	Funds types.DecCoin `protobuf:"bytes,8,opt,name=funds,proto3" json:"funds" yaml:"funds"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_50b386ecf43f1d04, []int{1}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetID() AccountID {
	if m != nil {
		return m.ID
	}
	return AccountID{}
}

func (m *Account) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Account) GetState() AccountState {
	if m != nil {
		return m.State
	}
	return AccountStateInvalid
}

func (m *Account) GetBalance() types.DecCoin {
	if m != nil {
		return m.Balance
	}
	return types.DecCoin{}
}

func (m *Account) GetTransferred() types.DecCoin {
	if m != nil {
		return m.Transferred
	}
	return types.DecCoin{}
}

func (m *Account) GetSettledAt() int64 {
	if m != nil {
		return m.SettledAt
	}
	return 0
}

func (m *Account) GetDepositor() string {
	if m != nil {
		return m.Depositor
	}
	return ""
}

func (m *Account) GetFunds() types.DecCoin {
	if m != nil {
		return m.Funds
	}
	return types.DecCoin{}
}

// Payment stores state for a payment
type FractionalPayment struct {
	AccountID AccountID               `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"accountID" yaml:"accountID"`
	PaymentID string                  `protobuf:"bytes,2,opt,name=payment_id,json=paymentId,proto3" json:"paymentID" yaml:"paymentID"`
	Owner     string                  `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	State     FractionalPayment_State `protobuf:"varint,4,opt,name=state,proto3,enum=akash.escrow.v1.FractionalPayment_State" json:"state" yaml:"state"`
	Rate      types.DecCoin           `protobuf:"bytes,5,opt,name=rate,proto3" json:"rate" yaml:"rate"`
	Balance   types.DecCoin           `protobuf:"bytes,6,opt,name=balance,proto3" json:"balance" yaml:"balance"`
	Withdrawn types.Coin              `protobuf:"bytes,7,opt,name=withdrawn,proto3" json:"withdrawn" yaml:"withdrawn"`
}

func (m *FractionalPayment) Reset()         { *m = FractionalPayment{} }
func (m *FractionalPayment) String() string { return proto.CompactTextString(m) }
func (*FractionalPayment) ProtoMessage()    {}
func (*FractionalPayment) Descriptor() ([]byte, []int) {
	return fileDescriptor_50b386ecf43f1d04, []int{2}
}
func (m *FractionalPayment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FractionalPayment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FractionalPayment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FractionalPayment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FractionalPayment.Merge(m, src)
}
func (m *FractionalPayment) XXX_Size() int {
	return m.Size()
}
func (m *FractionalPayment) XXX_DiscardUnknown() {
	xxx_messageInfo_FractionalPayment.DiscardUnknown(m)
}

var xxx_messageInfo_FractionalPayment proto.InternalMessageInfo

func (m *FractionalPayment) GetAccountID() AccountID {
	if m != nil {
		return m.AccountID
	}
	return AccountID{}
}

func (m *FractionalPayment) GetPaymentID() string {
	if m != nil {
		return m.PaymentID
	}
	return ""
}

func (m *FractionalPayment) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *FractionalPayment) GetState() FractionalPayment_State {
	if m != nil {
		return m.State
	}
	return PaymentStateInvalid
}

func (m *FractionalPayment) GetRate() types.DecCoin {
	if m != nil {
		return m.Rate
	}
	return types.DecCoin{}
}

func (m *FractionalPayment) GetBalance() types.DecCoin {
	if m != nil {
		return m.Balance
	}
	return types.DecCoin{}
}

func (m *FractionalPayment) GetWithdrawn() types.Coin {
	if m != nil {
		return m.Withdrawn
	}
	return types.Coin{}
}

func init() {
	proto.RegisterEnum("akash.escrow.v1.AccountState", AccountState_name, AccountState_value)
	proto.RegisterEnum("akash.escrow.v1.FractionalPayment_State", FractionalPayment_State_name, FractionalPayment_State_value)
	proto.RegisterType((*AccountID)(nil), "akash.escrow.v1.AccountID")
	proto.RegisterType((*Account)(nil), "akash.escrow.v1.Account")
	proto.RegisterType((*FractionalPayment)(nil), "akash.escrow.v1.FractionalPayment")
}

func init() { proto.RegisterFile("akash/escrow/v1/types.proto", fileDescriptor_50b386ecf43f1d04) }

var fileDescriptor_50b386ecf43f1d04 = []byte{
	// 830 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xbf, 0x8f, 0xe3, 0x44,
	0x14, 0x8e, 0xf3, 0x13, 0x4f, 0x8e, 0xbb, 0xdc, 0xb0, 0x12, 0x8e, 0xef, 0x62, 0x07, 0x03, 0x52,
	0x40, 0xc2, 0x56, 0x42, 0x81, 0xa0, 0xdb, 0xdc, 0x0a, 0x94, 0xe2, 0xf8, 0xe1, 0x2d, 0xe0, 0x28,
	0x88, 0x26, 0xf6, 0x5c, 0xce, 0xda, 0xc4, 0x63, 0x79, 0xe6, 0x92, 0xdb, 0xff, 0x00, 0xa5, 0x42,
	0x88, 0x86, 0x22, 0x15, 0x3d, 0x15, 0x7f, 0xc4, 0x95, 0x27, 0x2a, 0x2a, 0x0b, 0x65, 0xbb, 0x94,
	0xf9, 0x0b, 0x90, 0x67, 0xc6, 0x3f, 0x60, 0x77, 0x51, 0x24, 0xa8, 0x36, 0xef, 0xfb, 0xde, 0xfb,
	0xe6, 0xcd, 0xf7, 0xe6, 0xad, 0xc1, 0x03, 0x74, 0x81, 0xe8, 0x33, 0x07, 0x53, 0x2f, 0x26, 0x6b,
	0x67, 0x35, 0x74, 0xd8, 0x65, 0x84, 0xa9, 0x1d, 0xc5, 0x84, 0x11, 0x78, 0x8f, 0x93, 0xb6, 0x20,
	0xed, 0xd5, 0x50, 0x3f, 0x99, 0x93, 0x39, 0xe1, 0x9c, 0x93, 0xfe, 0x12, 0x69, 0x7a, 0xd7, 0x23,
	0x74, 0x49, 0xe8, 0x54, 0x10, 0x22, 0x90, 0x94, 0x21, 0x22, 0x67, 0x86, 0x28, 0x76, 0x56, 0xc3,
	0x19, 0x66, 0x68, 0xe8, 0x78, 0x24, 0x08, 0x05, 0x6f, 0x2d, 0x80, 0x7a, 0xea, 0x79, 0xe4, 0x79,
	0xc8, 0x26, 0x67, 0xd0, 0x01, 0x0d, 0xea, 0x91, 0x08, 0x6b, 0x4a, 0x5f, 0x19, 0xa8, 0xe3, 0xee,
	0x3e, 0x31, 0x05, 0x70, 0x48, 0xcc, 0x3b, 0x97, 0x68, 0xb9, 0xf8, 0xc4, 0xe2, 0xa1, 0xe5, 0x0a,
	0x18, 0xda, 0xa0, 0xf6, 0x22, 0xf0, 0xb5, 0x2a, 0x4f, 0x7f, 0xb8, 0x4b, 0xcc, 0xda, 0x37, 0x93,
	0xb3, 0x7d, 0x62, 0xa6, 0xe8, 0x21, 0x31, 0x81, 0xa8, 0x79, 0x11, 0xf8, 0x96, 0x9b, 0x42, 0xd6,
	0xaf, 0x0d, 0xd0, 0x92, 0xc7, 0xc1, 0x09, 0xa8, 0x06, 0x3e, 0x3f, 0xa9, 0x3d, 0xd2, 0xed, 0x7f,
	0x5c, 0xd4, 0xce, 0x9b, 0x1a, 0xf7, 0x5e, 0x26, 0x66, 0x65, 0x97, 0x98, 0x55, 0xae, 0x5c, 0xe5,
	0xc2, 0xaa, 0x10, 0x4e, 0x75, 0xab, 0x81, 0x0f, 0x3f, 0x03, 0x0d, 0xb2, 0x0e, 0x71, 0x2c, 0x1b,
	0x19, 0xa6, 0x7d, 0x73, 0xa0, 0xe8, 0x9b, 0x87, 0xd6, 0xef, 0xbf, 0x7d, 0x70, 0x22, 0xed, 0x39,
	0xf5, 0xfd, 0x18, 0x53, 0x7a, 0xce, 0xe2, 0x20, 0x9c, 0xbb, 0x22, 0x1d, 0x7e, 0x0e, 0x1a, 0x94,
	0x21, 0x86, 0xb5, 0x5a, 0x5f, 0x19, 0xdc, 0x1d, 0xf5, 0x6e, 0x6b, 0xeb, 0x3c, 0x4d, 0x92, 0xfe,
	0xa4, 0x3f, 0x4b, 0xfe, 0xa4, 0x61, 0xea, 0x4f, 0xfa, 0x17, 0x3e, 0x01, 0xad, 0x19, 0x5a, 0xa0,
	0xd0, 0xc3, 0x5a, 0x9d, 0x5f, 0xf4, 0xa1, 0x2d, 0x8f, 0x4f, 0xe7, 0x61, 0xcb, 0x79, 0xd8, 0x67,
	0xd8, 0x7b, 0x44, 0x82, 0x70, 0xfc, 0x56, 0x7a, 0xd5, 0x7d, 0x62, 0x66, 0x45, 0x87, 0xc4, 0xbc,
	0x2b, 0x64, 0x25, 0x60, 0xb9, 0x19, 0x05, 0x03, 0xd0, 0x66, 0x31, 0x0a, 0xe9, 0x53, 0x1c, 0xc7,
	0xd8, 0xd7, 0x1a, 0x47, 0xc8, 0xbf, 0x27, 0xe5, 0xcb, 0x85, 0x87, 0xc4, 0x84, 0xe2, 0x88, 0x12,
	0x68, 0xb9, 0xe5, 0x14, 0xf8, 0x18, 0x00, 0x8a, 0x19, 0x5b, 0x60, 0x7f, 0x8a, 0x98, 0xd6, 0xec,
	0x2b, 0x83, 0xda, 0xd8, 0xde, 0x25, 0xa6, 0x7a, 0x2e, 0xd0, 0x53, 0xb6, 0x4f, 0x4c, 0x95, 0x66,
	0xc1, 0x21, 0x31, 0x3b, 0xd2, 0x8c, 0x0c, 0xb2, 0xdc, 0x82, 0x86, 0x5f, 0x03, 0xd5, 0xc7, 0x11,
	0xa1, 0x01, 0x23, 0xb1, 0xd6, 0xe2, 0x13, 0xfb, 0x38, 0x15, 0xc8, 0xc1, 0x42, 0x20, 0x87, 0x6e,
	0x9f, 0x5c, 0x51, 0x06, 0xbf, 0x02, 0x8d, 0xa7, 0xcf, 0x43, 0x9f, 0x6a, 0xaf, 0x1d, 0x61, 0x46,
	0x4f, 0x9a, 0x21, 0x4a, 0x8a, 0x01, 0xf2, 0xd0, 0x72, 0x05, 0x6c, 0xfd, 0xd4, 0x04, 0xf7, 0x3f,
	0x8d, 0x91, 0xc7, 0x02, 0x12, 0xa2, 0xc5, 0x97, 0xe8, 0x72, 0x89, 0x43, 0x06, 0x97, 0x00, 0x20,
	0xf1, 0x10, 0xa6, 0x47, 0x3d, 0xe1, 0x91, 0x7c, 0xc2, 0xc5, 0xaa, 0xa5, 0xf7, 0x45, 0x59, 0x50,
	0xdc, 0x37, 0x87, 0x2c, 0x37, 0xa7, 0xb9, 0xff, 0x91, 0x38, 0x79, 0x9a, 0x2f, 0x1b, 0xf7, 0x5f,
	0xf6, 0x23, 0xe4, 0xa2, 0x2c, 0x28, 0xe4, 0x72, 0xc8, 0x72, 0x73, 0xba, 0xb4, 0x2d, 0xb5, 0xff,
	0xb8, 0x2d, 0x4f, 0xb2, 0x6d, 0xa9, 0xf3, 0x6d, 0x19, 0x5c, 0x73, 0xe0, 0x9a, 0x73, 0xf6, 0xb1,
	0x8b, 0xf3, 0x18, 0xd4, 0xe3, 0x54, 0xf9, 0x98, 0x67, 0xfd, 0x40, 0x4e, 0x92, 0x57, 0x1c, 0x12,
	0xb3, 0x2d, 0x04, 0x63, 0xae, 0xc7, 0xc1, 0xf2, 0x1e, 0x36, 0xff, 0xe7, 0x3d, 0xfc, 0x0e, 0xa8,
	0xeb, 0x80, 0x3d, 0xf3, 0x63, 0xb4, 0x0e, 0xf9, 0x6b, 0x6e, 0x8f, 0xba, 0x37, 0x8a, 0x73, 0xe5,
	0x77, 0xa5, 0x72, 0x51, 0x53, 0x4c, 0x2b, 0x87, 0x2c, 0xb7, 0xa0, 0xad, 0x1f, 0x15, 0xd0, 0xe0,
	0xae, 0xc1, 0x77, 0x40, 0x2b, 0x08, 0x57, 0x68, 0x11, 0xf8, 0x9d, 0x8a, 0xfe, 0xe6, 0x66, 0xdb,
	0x7f, 0x43, 0xba, 0xca, 0xe9, 0x89, 0xa0, 0x60, 0x17, 0xd4, 0x49, 0x84, 0xc3, 0x8e, 0xa2, 0xdf,
	0xdb, 0x6c, 0xfb, 0x6d, 0x99, 0xf2, 0x45, 0x84, 0x43, 0xd8, 0x03, 0x4d, 0x6f, 0x41, 0x28, 0xf6,
	0x3b, 0x55, 0xfd, 0xfe, 0x66, 0xdb, 0x7f, 0x5d, 0x92, 0x8f, 0x38, 0x08, 0xdf, 0x06, 0x2a, 0x59,
	0xe1, 0x98, 0x1f, 0xdb, 0xa9, 0xe9, 0x27, 0x9b, 0x6d, 0xbf, 0x93, 0x95, 0x67, 0xb8, 0x5e, 0xff,
	0xfe, 0x17, 0xa3, 0xf2, 0xfe, 0xcf, 0x0a, 0xb8, 0x53, 0xfe, 0x57, 0x78, 0x43, 0x6f, 0x65, 0xfa,
	0x96, 0xde, 0x64, 0xca, 0xcd, 0xbd, 0x49, 0xf2, 0x5f, 0x7a, 0xcb, 0xca, 0xff, 0xde, 0xdb, 0xf8,
	0xa3, 0x97, 0x3b, 0x43, 0x79, 0xb5, 0x33, 0x94, 0x3f, 0x77, 0x86, 0xf2, 0xc3, 0x95, 0x51, 0x79,
	0x75, 0x65, 0x54, 0xfe, 0xb8, 0x32, 0x2a, 0xdf, 0xf6, 0xa2, 0x8b, 0xb9, 0x8d, 0x2e, 0x98, 0xed,
	0xe3, 0x95, 0x33, 0x27, 0x4e, 0x48, 0x7c, 0x5c, 0x7c, 0x78, 0x67, 0x4d, 0xfe, 0x45, 0xfc, 0xf0,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0xd8, 0x74, 0x70, 0x92, 0x07, 0x00, 0x00,
}

func (m *AccountID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.XID) > 0 {
		i -= len(m.XID)
		copy(dAtA[i:], m.XID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.XID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Scope) > 0 {
		i -= len(m.Scope)
		copy(dAtA[i:], m.Scope)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Scope)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Funds.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x3a
	}
	if m.SettledAt != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.SettledAt))
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.Transferred.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.State != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *FractionalPayment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FractionalPayment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FractionalPayment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Withdrawn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Rate.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.State != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PaymentID) > 0 {
		i -= len(m.PaymentID)
		copy(dAtA[i:], m.PaymentID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.PaymentID)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.AccountID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Scope)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.XID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovTypes(uint64(m.State))
	}
	l = m.Balance.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Transferred.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.SettledAt != 0 {
		n += 1 + sovTypes(uint64(m.SettledAt))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Funds.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *FractionalPayment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AccountID.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.PaymentID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovTypes(uint64(m.State))
	}
	l = m.Rate.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Balance.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Withdrawn.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: AccountID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scope", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scope = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.XID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= AccountState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transferred", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Transferred.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SettledAt", wireType)
			}
			m.SettledAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SettledAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Funds.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *FractionalPayment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: FractionalPayment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FractionalPayment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AccountID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PaymentID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= FractionalPayment_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdrawn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Withdrawn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
