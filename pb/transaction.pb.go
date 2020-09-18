// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/transaction.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PayloadType int32

const (
	PayloadType_COINBASE_TYPE       PayloadType = 0
	PayloadType_TRANSFER_ASSET_TYPE PayloadType = 1
	PayloadType_SIG_CHAIN_TXN_TYPE  PayloadType = 2
	PayloadType_REGISTER_NAME_TYPE  PayloadType = 3
	PayloadType_TRANSFER_NAME_TYPE  PayloadType = 4
	PayloadType_DELETE_NAME_TYPE    PayloadType = 5
	PayloadType_SUBSCRIBE_TYPE      PayloadType = 6
	PayloadType_UNSUBSCRIBE_TYPE    PayloadType = 7
	PayloadType_GENERATE_ID_TYPE    PayloadType = 8
	PayloadType_NANO_PAY_TYPE       PayloadType = 9
	PayloadType_ISSUE_ASSET_TYPE    PayloadType = 10
)

var PayloadType_name = map[int32]string{
	0:  "COINBASE_TYPE",
	1:  "TRANSFER_ASSET_TYPE",
	2:  "SIG_CHAIN_TXN_TYPE",
	3:  "REGISTER_NAME_TYPE",
	4:  "TRANSFER_NAME_TYPE",
	5:  "DELETE_NAME_TYPE",
	6:  "SUBSCRIBE_TYPE",
	7:  "UNSUBSCRIBE_TYPE",
	8:  "GENERATE_ID_TYPE",
	9:  "NANO_PAY_TYPE",
	10: "ISSUE_ASSET_TYPE",
}
var PayloadType_value = map[string]int32{
	"COINBASE_TYPE":       0,
	"TRANSFER_ASSET_TYPE": 1,
	"SIG_CHAIN_TXN_TYPE":  2,
	"REGISTER_NAME_TYPE":  3,
	"TRANSFER_NAME_TYPE":  4,
	"DELETE_NAME_TYPE":    5,
	"SUBSCRIBE_TYPE":      6,
	"UNSUBSCRIBE_TYPE":    7,
	"GENERATE_ID_TYPE":    8,
	"NANO_PAY_TYPE":       9,
	"ISSUE_ASSET_TYPE":    10,
}

func (x PayloadType) String() string {
	return proto.EnumName(PayloadType_name, int32(x))
}
func (PayloadType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{0}
}

type UnsignedTx struct {
	Payload              *Payload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Nonce                uint64   `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Fee                  int64    `protobuf:"varint,3,opt,name=fee,proto3" json:"fee,omitempty"`
	Attributes           []byte   `protobuf:"bytes,4,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnsignedTx) Reset()         { *m = UnsignedTx{} }
func (m *UnsignedTx) String() string { return proto.CompactTextString(m) }
func (*UnsignedTx) ProtoMessage()    {}
func (*UnsignedTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{0}
}
func (m *UnsignedTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnsignedTx.Unmarshal(m, b)
}
func (m *UnsignedTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnsignedTx.Marshal(b, m, deterministic)
}
func (dst *UnsignedTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsignedTx.Merge(dst, src)
}
func (m *UnsignedTx) XXX_Size() int {
	return xxx_messageInfo_UnsignedTx.Size(m)
}
func (m *UnsignedTx) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsignedTx.DiscardUnknown(m)
}

var xxx_messageInfo_UnsignedTx proto.InternalMessageInfo

func (m *UnsignedTx) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *UnsignedTx) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *UnsignedTx) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *UnsignedTx) GetAttributes() []byte {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type Transaction struct {
	UnsignedTx           *UnsignedTx `protobuf:"bytes,1,opt,name=unsigned_tx,json=unsignedTx,proto3" json:"unsigned_tx,omitempty"`
	Programs             []*Program  `protobuf:"bytes,2,rep,name=programs,proto3" json:"programs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{1}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetUnsignedTx() *UnsignedTx {
	if m != nil {
		return m.UnsignedTx
	}
	return nil
}

func (m *Transaction) GetPrograms() []*Program {
	if m != nil {
		return m.Programs
	}
	return nil
}

type Program struct {
	Code                 []byte   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Parameter            []byte   `protobuf:"bytes,2,opt,name=parameter,proto3" json:"parameter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Program) Reset()         { *m = Program{} }
func (m *Program) String() string { return proto.CompactTextString(m) }
func (*Program) ProtoMessage()    {}
func (*Program) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{2}
}
func (m *Program) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Program.Unmarshal(m, b)
}
func (m *Program) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Program.Marshal(b, m, deterministic)
}
func (dst *Program) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Program.Merge(dst, src)
}
func (m *Program) XXX_Size() int {
	return xxx_messageInfo_Program.Size(m)
}
func (m *Program) XXX_DiscardUnknown() {
	xxx_messageInfo_Program.DiscardUnknown(m)
}

var xxx_messageInfo_Program proto.InternalMessageInfo

func (m *Program) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *Program) GetParameter() []byte {
	if m != nil {
		return m.Parameter
	}
	return nil
}

type Payload struct {
	Type                 PayloadType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.PayloadType" json:"type,omitempty"`
	Data                 []byte      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{3}
}
func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (dst *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(dst, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetType() PayloadType {
	if m != nil {
		return m.Type
	}
	return PayloadType_COINBASE_TYPE
}

func (m *Payload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Coinbase struct {
	Sender               []byte   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            []byte   `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount               int64    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coinbase) Reset()         { *m = Coinbase{} }
func (m *Coinbase) String() string { return proto.CompactTextString(m) }
func (*Coinbase) ProtoMessage()    {}
func (*Coinbase) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{4}
}
func (m *Coinbase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coinbase.Unmarshal(m, b)
}
func (m *Coinbase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coinbase.Marshal(b, m, deterministic)
}
func (dst *Coinbase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coinbase.Merge(dst, src)
}
func (m *Coinbase) XXX_Size() int {
	return xxx_messageInfo_Coinbase.Size(m)
}
func (m *Coinbase) XXX_DiscardUnknown() {
	xxx_messageInfo_Coinbase.DiscardUnknown(m)
}

var xxx_messageInfo_Coinbase proto.InternalMessageInfo

func (m *Coinbase) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Coinbase) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *Coinbase) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type SigChainTxn struct {
	SigChain             []byte   `protobuf:"bytes,1,opt,name=sig_chain,json=sigChain,proto3" json:"sig_chain,omitempty"`
	Submitter            []byte   `protobuf:"bytes,2,opt,name=submitter,proto3" json:"submitter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SigChainTxn) Reset()         { *m = SigChainTxn{} }
func (m *SigChainTxn) String() string { return proto.CompactTextString(m) }
func (*SigChainTxn) ProtoMessage()    {}
func (*SigChainTxn) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{5}
}
func (m *SigChainTxn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SigChainTxn.Unmarshal(m, b)
}
func (m *SigChainTxn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SigChainTxn.Marshal(b, m, deterministic)
}
func (dst *SigChainTxn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigChainTxn.Merge(dst, src)
}
func (m *SigChainTxn) XXX_Size() int {
	return xxx_messageInfo_SigChainTxn.Size(m)
}
func (m *SigChainTxn) XXX_DiscardUnknown() {
	xxx_messageInfo_SigChainTxn.DiscardUnknown(m)
}

var xxx_messageInfo_SigChainTxn proto.InternalMessageInfo

func (m *SigChainTxn) GetSigChain() []byte {
	if m != nil {
		return m.SigChain
	}
	return nil
}

func (m *SigChainTxn) GetSubmitter() []byte {
	if m != nil {
		return m.Submitter
	}
	return nil
}

type RegisterName struct {
	Registrant           []byte   `protobuf:"bytes,1,opt,name=registrant,proto3" json:"registrant,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RegistrationFee      int64    `protobuf:"varint,3,opt,name=registration_fee,json=registrationFee,proto3" json:"registration_fee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterName) Reset()         { *m = RegisterName{} }
func (m *RegisterName) String() string { return proto.CompactTextString(m) }
func (*RegisterName) ProtoMessage()    {}
func (*RegisterName) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{6}
}
func (m *RegisterName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterName.Unmarshal(m, b)
}
func (m *RegisterName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterName.Marshal(b, m, deterministic)
}
func (dst *RegisterName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterName.Merge(dst, src)
}
func (m *RegisterName) XXX_Size() int {
	return xxx_messageInfo_RegisterName.Size(m)
}
func (m *RegisterName) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterName.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterName proto.InternalMessageInfo

func (m *RegisterName) GetRegistrant() []byte {
	if m != nil {
		return m.Registrant
	}
	return nil
}

func (m *RegisterName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterName) GetRegistrationFee() int64 {
	if m != nil {
		return m.RegistrationFee
	}
	return 0
}

type TransferName struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Registrant           []byte   `protobuf:"bytes,2,opt,name=registrant,proto3" json:"registrant,omitempty"`
	Recipient            []byte   `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferName) Reset()         { *m = TransferName{} }
func (m *TransferName) String() string { return proto.CompactTextString(m) }
func (*TransferName) ProtoMessage()    {}
func (*TransferName) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{7}
}
func (m *TransferName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferName.Unmarshal(m, b)
}
func (m *TransferName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferName.Marshal(b, m, deterministic)
}
func (dst *TransferName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferName.Merge(dst, src)
}
func (m *TransferName) XXX_Size() int {
	return xxx_messageInfo_TransferName.Size(m)
}
func (m *TransferName) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferName.DiscardUnknown(m)
}

var xxx_messageInfo_TransferName proto.InternalMessageInfo

func (m *TransferName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TransferName) GetRegistrant() []byte {
	if m != nil {
		return m.Registrant
	}
	return nil
}

func (m *TransferName) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

type DeleteName struct {
	Registrant           []byte   `protobuf:"bytes,1,opt,name=registrant,proto3" json:"registrant,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteName) Reset()         { *m = DeleteName{} }
func (m *DeleteName) String() string { return proto.CompactTextString(m) }
func (*DeleteName) ProtoMessage()    {}
func (*DeleteName) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{8}
}
func (m *DeleteName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteName.Unmarshal(m, b)
}
func (m *DeleteName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteName.Marshal(b, m, deterministic)
}
func (dst *DeleteName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteName.Merge(dst, src)
}
func (m *DeleteName) XXX_Size() int {
	return xxx_messageInfo_DeleteName.Size(m)
}
func (m *DeleteName) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteName.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteName proto.InternalMessageInfo

func (m *DeleteName) GetRegistrant() []byte {
	if m != nil {
		return m.Registrant
	}
	return nil
}

func (m *DeleteName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Subscribe struct {
	Subscriber           []byte   `protobuf:"bytes,1,opt,name=subscriber,proto3" json:"subscriber,omitempty"`
	Identifier           string   `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Topic                string   `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Bucket               uint32   `protobuf:"varint,4,opt,name=bucket,proto3" json:"bucket,omitempty"` // Deprecated: Do not use.
	Duration             uint32   `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
	Meta                 string   `protobuf:"bytes,6,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subscribe) Reset()         { *m = Subscribe{} }
func (m *Subscribe) String() string { return proto.CompactTextString(m) }
func (*Subscribe) ProtoMessage()    {}
func (*Subscribe) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{9}
}
func (m *Subscribe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscribe.Unmarshal(m, b)
}
func (m *Subscribe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscribe.Marshal(b, m, deterministic)
}
func (dst *Subscribe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscribe.Merge(dst, src)
}
func (m *Subscribe) XXX_Size() int {
	return xxx_messageInfo_Subscribe.Size(m)
}
func (m *Subscribe) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscribe.DiscardUnknown(m)
}

var xxx_messageInfo_Subscribe proto.InternalMessageInfo

func (m *Subscribe) GetSubscriber() []byte {
	if m != nil {
		return m.Subscriber
	}
	return nil
}

func (m *Subscribe) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Subscribe) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

// Deprecated: Do not use.
func (m *Subscribe) GetBucket() uint32 {
	if m != nil {
		return m.Bucket
	}
	return 0
}

func (m *Subscribe) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Subscribe) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

type Unsubscribe struct {
	Subscriber           []byte   `protobuf:"bytes,1,opt,name=subscriber,proto3" json:"subscriber,omitempty"`
	Identifier           string   `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Topic                string   `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Unsubscribe) Reset()         { *m = Unsubscribe{} }
func (m *Unsubscribe) String() string { return proto.CompactTextString(m) }
func (*Unsubscribe) ProtoMessage()    {}
func (*Unsubscribe) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{10}
}
func (m *Unsubscribe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unsubscribe.Unmarshal(m, b)
}
func (m *Unsubscribe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unsubscribe.Marshal(b, m, deterministic)
}
func (dst *Unsubscribe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unsubscribe.Merge(dst, src)
}
func (m *Unsubscribe) XXX_Size() int {
	return xxx_messageInfo_Unsubscribe.Size(m)
}
func (m *Unsubscribe) XXX_DiscardUnknown() {
	xxx_messageInfo_Unsubscribe.DiscardUnknown(m)
}

var xxx_messageInfo_Unsubscribe proto.InternalMessageInfo

func (m *Unsubscribe) GetSubscriber() []byte {
	if m != nil {
		return m.Subscriber
	}
	return nil
}

func (m *Unsubscribe) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Unsubscribe) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type TransferAsset struct {
	Sender               []byte   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            []byte   `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount               int64    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferAsset) Reset()         { *m = TransferAsset{} }
func (m *TransferAsset) String() string { return proto.CompactTextString(m) }
func (*TransferAsset) ProtoMessage()    {}
func (*TransferAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{11}
}
func (m *TransferAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferAsset.Unmarshal(m, b)
}
func (m *TransferAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferAsset.Marshal(b, m, deterministic)
}
func (dst *TransferAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferAsset.Merge(dst, src)
}
func (m *TransferAsset) XXX_Size() int {
	return xxx_messageInfo_TransferAsset.Size(m)
}
func (m *TransferAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferAsset.DiscardUnknown(m)
}

var xxx_messageInfo_TransferAsset proto.InternalMessageInfo

func (m *TransferAsset) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *TransferAsset) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *TransferAsset) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type GenerateID struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	RegistrationFee      int64    `protobuf:"varint,2,opt,name=registration_fee,json=registrationFee,proto3" json:"registration_fee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateID) Reset()         { *m = GenerateID{} }
func (m *GenerateID) String() string { return proto.CompactTextString(m) }
func (*GenerateID) ProtoMessage()    {}
func (*GenerateID) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{12}
}
func (m *GenerateID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateID.Unmarshal(m, b)
}
func (m *GenerateID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateID.Marshal(b, m, deterministic)
}
func (dst *GenerateID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateID.Merge(dst, src)
}
func (m *GenerateID) XXX_Size() int {
	return xxx_messageInfo_GenerateID.Size(m)
}
func (m *GenerateID) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateID.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateID proto.InternalMessageInfo

func (m *GenerateID) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *GenerateID) GetRegistrationFee() int64 {
	if m != nil {
		return m.RegistrationFee
	}
	return 0
}

type NanoPay struct {
	Sender               []byte   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            []byte   `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Id                   uint64   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Amount               int64    `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	TxnExpiration        uint32   `protobuf:"varint,5,opt,name=txn_expiration,json=txnExpiration,proto3" json:"txn_expiration,omitempty"`
	NanoPayExpiration    uint32   `protobuf:"varint,6,opt,name=nano_pay_expiration,json=nanoPayExpiration,proto3" json:"nano_pay_expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NanoPay) Reset()         { *m = NanoPay{} }
func (m *NanoPay) String() string { return proto.CompactTextString(m) }
func (*NanoPay) ProtoMessage()    {}
func (*NanoPay) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{13}
}
func (m *NanoPay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NanoPay.Unmarshal(m, b)
}
func (m *NanoPay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NanoPay.Marshal(b, m, deterministic)
}
func (dst *NanoPay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NanoPay.Merge(dst, src)
}
func (m *NanoPay) XXX_Size() int {
	return xxx_messageInfo_NanoPay.Size(m)
}
func (m *NanoPay) XXX_DiscardUnknown() {
	xxx_messageInfo_NanoPay.DiscardUnknown(m)
}

var xxx_messageInfo_NanoPay proto.InternalMessageInfo

func (m *NanoPay) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *NanoPay) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *NanoPay) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NanoPay) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *NanoPay) GetTxnExpiration() uint32 {
	if m != nil {
		return m.TxnExpiration
	}
	return 0
}

func (m *NanoPay) GetNanoPayExpiration() uint32 {
	if m != nil {
		return m.NanoPayExpiration
	}
	return 0
}

type IssueAsset struct {
	Sender               []byte   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol               string   `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	TotalSupply          int64    `protobuf:"varint,4,opt,name=total_supply,json=totalSupply,proto3" json:"total_supply,omitempty"`
	Precision            uint32   `protobuf:"varint,5,opt,name=precision,proto3" json:"precision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueAsset) Reset()         { *m = IssueAsset{} }
func (m *IssueAsset) String() string { return proto.CompactTextString(m) }
func (*IssueAsset) ProtoMessage()    {}
func (*IssueAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_dee487b0d08e39df, []int{14}
}
func (m *IssueAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueAsset.Unmarshal(m, b)
}
func (m *IssueAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueAsset.Marshal(b, m, deterministic)
}
func (dst *IssueAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueAsset.Merge(dst, src)
}
func (m *IssueAsset) XXX_Size() int {
	return xxx_messageInfo_IssueAsset.Size(m)
}
func (m *IssueAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueAsset.DiscardUnknown(m)
}

var xxx_messageInfo_IssueAsset proto.InternalMessageInfo

func (m *IssueAsset) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *IssueAsset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IssueAsset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *IssueAsset) GetTotalSupply() int64 {
	if m != nil {
		return m.TotalSupply
	}
	return 0
}

func (m *IssueAsset) GetPrecision() uint32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func init() {
	proto.RegisterType((*UnsignedTx)(nil), "pb.UnsignedTx")
	proto.RegisterType((*Transaction)(nil), "pb.Transaction")
	proto.RegisterType((*Program)(nil), "pb.Program")
	proto.RegisterType((*Payload)(nil), "pb.Payload")
	proto.RegisterType((*Coinbase)(nil), "pb.Coinbase")
	proto.RegisterType((*SigChainTxn)(nil), "pb.SigChainTxn")
	proto.RegisterType((*RegisterName)(nil), "pb.RegisterName")
	proto.RegisterType((*TransferName)(nil), "pb.TransferName")
	proto.RegisterType((*DeleteName)(nil), "pb.DeleteName")
	proto.RegisterType((*Subscribe)(nil), "pb.Subscribe")
	proto.RegisterType((*Unsubscribe)(nil), "pb.Unsubscribe")
	proto.RegisterType((*TransferAsset)(nil), "pb.TransferAsset")
	proto.RegisterType((*GenerateID)(nil), "pb.GenerateID")
	proto.RegisterType((*NanoPay)(nil), "pb.NanoPay")
	proto.RegisterType((*IssueAsset)(nil), "pb.IssueAsset")
	proto.RegisterEnum("pb.PayloadType", PayloadType_name, PayloadType_value)
}

func init() { proto.RegisterFile("pb/transaction.proto", fileDescriptor_transaction_dee487b0d08e39df) }

var fileDescriptor_transaction_dee487b0d08e39df = []byte{
	// 833 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xc7, 0xb1, 0x93, 0xcd, 0xc7, 0x49, 0x36, 0xdd, 0x4e, 0x57, 0x25, 0x2a, 0x1f, 0x0a, 0x46,
	0x15, 0x0b, 0x17, 0x5b, 0xa9, 0x5c, 0x72, 0x43, 0x92, 0x75, 0xb7, 0x16, 0xe0, 0xae, 0xc6, 0x0e,
	0x6a, 0x2f, 0x90, 0x19, 0xdb, 0xb3, 0x61, 0xd4, 0x78, 0x6c, 0x79, 0xc6, 0x52, 0x2c, 0x6e, 0x78,
	0x05, 0x1e, 0x84, 0x57, 0xe0, 0xd9, 0xd0, 0x8c, 0x27, 0xb1, 0xb7, 0xac, 0xb8, 0xa8, 0xe8, 0xdd,
	0x9c, 0xdf, 0x39, 0x39, 0xf3, 0x3f, 0x1f, 0xe3, 0xc0, 0x79, 0x11, 0x3f, 0x93, 0x25, 0xe1, 0x82,
	0x24, 0x92, 0xe5, 0xfc, 0xb2, 0x28, 0x73, 0x99, 0x23, 0xbb, 0x88, 0x9d, 0xdf, 0x01, 0x36, 0x5c,
	0xb0, 0x2d, 0xa7, 0x69, 0xb8, 0x47, 0x4f, 0x61, 0x58, 0x90, 0x7a, 0x97, 0x93, 0x74, 0x6e, 0x2d,
	0xac, 0x8b, 0xc9, 0xf3, 0xc9, 0x65, 0x11, 0x5f, 0xde, 0x34, 0x08, 0x1f, 0x7c, 0xe8, 0x1c, 0x4e,
	0x78, 0xce, 0x13, 0x3a, 0xb7, 0x17, 0xd6, 0x45, 0x1f, 0x37, 0x06, 0x3a, 0x83, 0xde, 0x2d, 0xa5,
	0xf3, 0xde, 0xc2, 0xba, 0xe8, 0x61, 0x75, 0x44, 0x9f, 0x03, 0x10, 0x29, 0x4b, 0x16, 0x57, 0x92,
	0x8a, 0x79, 0x7f, 0x61, 0x5d, 0x4c, 0x71, 0x87, 0x38, 0x5b, 0x98, 0x84, 0xad, 0x2a, 0xf4, 0x0c,
	0x26, 0x95, 0xd1, 0x12, 0xc9, 0xbd, 0x51, 0x30, 0x53, 0x0a, 0x5a, 0x89, 0x18, 0xaa, 0x56, 0xee,
	0x57, 0x30, 0x2a, 0xca, 0x7c, 0x5b, 0x92, 0x4c, 0xcc, 0xed, 0x45, 0xef, 0xa8, 0xb7, 0x61, 0xf8,
	0xe8, 0x74, 0xbe, 0x83, 0xa1, 0x81, 0x08, 0x41, 0x3f, 0xc9, 0x53, 0xaa, 0xb3, 0x4f, 0xb1, 0x3e,
	0xa3, 0x4f, 0x61, 0x5c, 0x90, 0x92, 0x64, 0x54, 0xd2, 0x52, 0xd7, 0x34, 0xc5, 0x2d, 0x70, 0x56,
	0x30, 0x34, 0x1d, 0x40, 0x5f, 0x42, 0x5f, 0xd6, 0x45, 0xf3, 0xe3, 0xd9, 0xf3, 0x07, 0x9d, 0xe6,
	0x84, 0x75, 0x41, 0xb1, 0x76, 0xaa, 0x1b, 0x52, 0x22, 0x89, 0x49, 0xa4, 0xcf, 0xce, 0x6b, 0x18,
	0xad, 0x73, 0xc6, 0x63, 0x22, 0x28, 0x7a, 0x0c, 0x03, 0x41, 0x79, 0x4a, 0x4b, 0xa3, 0xc1, 0x58,
	0x4a, 0x45, 0x49, 0x13, 0x56, 0x30, 0xca, 0xe5, 0x41, 0xc5, 0x11, 0xa8, 0x5f, 0x91, 0x2c, 0xaf,
	0xb8, 0x34, 0x0d, 0x36, 0x96, 0xf3, 0x12, 0x26, 0x01, 0xdb, 0xae, 0x7f, 0x23, 0x8c, 0x87, 0x7b,
	0x8e, 0x3e, 0x81, 0xb1, 0x60, 0xdb, 0x28, 0x51, 0xb6, 0xc9, 0x3f, 0x12, 0xc6, 0xaf, 0x6e, 0x10,
	0x55, 0x9c, 0x31, 0xd9, 0xa9, 0xf3, 0x08, 0x9c, 0x0c, 0xa6, 0x98, 0x6e, 0x99, 0x90, 0xb4, 0xf4,
	0x49, 0xa6, 0xa7, 0x57, 0x6a, 0xbb, 0x24, 0x5c, 0x9a, 0x5c, 0x1d, 0xa2, 0xea, 0xe4, 0x24, 0x6b,
	0x96, 0x60, 0x8c, 0xf5, 0x19, 0x7d, 0x0d, 0x67, 0x87, 0x08, 0x35, 0xd2, 0xa8, 0x5d, 0x88, 0x07,
	0x5d, 0xfe, 0x82, 0x52, 0xe7, 0x57, 0x98, 0xea, 0xe1, 0xdf, 0x9a, 0xeb, 0x0e, 0xe9, 0xac, 0x4e,
	0xba, 0xbb, 0x12, 0xec, 0x7f, 0x49, 0xb8, 0xd3, 0xb2, 0xde, 0x3b, 0x2d, 0x73, 0xbe, 0x07, 0xb8,
	0xa2, 0x3b, 0x2a, 0xe9, 0xfb, 0x96, 0xe3, 0xfc, 0x65, 0xc1, 0x38, 0xa8, 0x62, 0x91, 0x94, 0x2c,
	0xd6, 0x19, 0xc4, 0xc1, 0x38, 0x0c, 0xaf, 0x43, 0x94, 0x9f, 0xa5, 0x94, 0x4b, 0x76, 0xcb, 0x4c,
	0x7f, 0xc7, 0xb8, 0x43, 0xd4, 0xb3, 0x91, 0x79, 0xc1, 0x12, 0xad, 0x74, 0x8c, 0x1b, 0x03, 0x3d,
	0x81, 0x41, 0x5c, 0x25, 0x6f, 0xa9, 0xd4, 0x0f, 0xe4, 0x74, 0x65, 0xcf, 0x2d, 0x6c, 0x08, 0x7a,
	0x02, 0xa3, 0xb4, 0x6a, 0x5a, 0x36, 0x3f, 0x51, 0x5e, 0x7c, 0xb4, 0x95, 0xde, 0x8c, 0x4a, 0x32,
	0x1f, 0x34, 0x7a, 0xd5, 0xd9, 0x49, 0x60, 0xb2, 0xe1, 0xe2, 0xc3, 0x0a, 0x76, 0x7e, 0x81, 0xd3,
	0xc3, 0xe0, 0x96, 0x42, 0x50, 0xf9, 0x3f, 0x2f, 0xf4, 0xcf, 0x00, 0xd7, 0x94, 0xd3, 0x92, 0x48,
	0xea, 0x5d, 0xa1, 0xcf, 0x00, 0x8a, 0x2a, 0xde, 0xb1, 0x24, 0x7a, 0x4b, 0x6b, 0x93, 0x7f, 0xdc,
	0x90, 0x1f, 0x68, 0x7d, 0xef, 0xbe, 0xd9, 0xf7, 0xef, 0xdb, 0xdf, 0x16, 0x0c, 0x7d, 0xc2, 0xf3,
	0x1b, 0x52, 0xbf, 0xa7, 0xe2, 0x19, 0xd8, 0x2c, 0xd5, 0x6a, 0xfb, 0xd8, 0x66, 0x69, 0xa7, 0x82,
	0x7e, 0xb7, 0x02, 0xf4, 0x14, 0x66, 0x72, 0xcf, 0x23, 0xba, 0x2f, 0xd8, 0x9d, 0xd9, 0x9d, 0xca,
	0x3d, 0x77, 0x8f, 0x10, 0x5d, 0xc2, 0x23, 0x4e, 0x78, 0x1e, 0x15, 0xa4, 0xee, 0xc6, 0x0e, 0x74,
	0xec, 0x43, 0xde, 0x48, 0x6d, 0xe3, 0x9d, 0x3f, 0x2d, 0x00, 0x4f, 0x88, 0x8a, 0xfe, 0x77, 0xd7,
	0xef, 0x7b, 0x96, 0x2a, 0xb6, 0xce, 0xe2, 0x7c, 0x67, 0x26, 0x69, 0x2c, 0xf4, 0x05, 0x4c, 0x65,
	0x2e, 0xc9, 0x2e, 0x12, 0x55, 0x51, 0xec, 0x6a, 0x53, 0xc7, 0x44, 0xb3, 0x40, 0x23, 0xfd, 0x6d,
	0x54, 0x2d, 0x10, 0x6d, 0x1d, 0x2d, 0xf8, 0xe6, 0x0f, 0x1b, 0x26, 0x9d, 0x2f, 0x20, 0x7a, 0x08,
	0xa7, 0xeb, 0x57, 0x9e, 0xbf, 0x5a, 0x06, 0x6e, 0x14, 0xbe, 0xb9, 0x71, 0xcf, 0x3e, 0x42, 0x1f,
	0xc3, 0xa3, 0x10, 0x2f, 0xfd, 0xe0, 0x85, 0x8b, 0xa3, 0x65, 0x10, 0xb8, 0x61, 0xe3, 0xb0, 0xd0,
	0x63, 0x40, 0x81, 0x77, 0x1d, 0xad, 0x5f, 0x2e, 0x3d, 0x3f, 0x0a, 0x5f, 0xfb, 0x0d, 0xb7, 0x15,
	0xc7, 0xee, 0xb5, 0x17, 0x84, 0x2e, 0x8e, 0xfc, 0xe5, 0x4f, 0x26, 0x51, 0x4f, 0xf1, 0x63, 0xa2,
	0x96, 0xf7, 0xd1, 0x39, 0x9c, 0x5d, 0xb9, 0x3f, 0xba, 0xa1, 0xdb, 0xa1, 0x27, 0x08, 0xc1, 0x2c,
	0xd8, 0xac, 0x82, 0x35, 0xf6, 0x56, 0x86, 0x0d, 0x54, 0xe4, 0xc6, 0x7f, 0x87, 0x0e, 0x15, 0xbd,
	0x76, 0x7d, 0x17, 0x2f, 0x43, 0x37, 0xf2, 0xae, 0x1a, 0x3a, 0x52, 0x95, 0xf8, 0x4b, 0xff, 0x55,
	0x74, 0xb3, 0x7c, 0xd3, 0xa0, 0xb1, 0x0a, 0xf4, 0x82, 0x60, 0xe3, 0x76, 0xcb, 0x80, 0x78, 0xa0,
	0xff, 0x4c, 0xbf, 0xfd, 0x27, 0x00, 0x00, 0xff, 0xff, 0x38, 0x07, 0xb4, 0xac, 0x64, 0x07, 0x00,
	0x00,
}
