// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package ethereum_sharding_p2p_v1

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

// TODO: Split the topics into p2p for beacon chain and p2p for sharding.
type Topic int32

const (
	Topic_UNKNOWN                          Topic = 0
	Topic_COLLATION_BODY_REQUEST           Topic = 1
	Topic_COLLATION_BODY_RESPONSE          Topic = 2
	Topic_TRANSACTIONS                     Topic = 3
	Topic_BEACON_BLOCK_HASH_ANNOUNCE       Topic = 4
	Topic_BEACON_BLOCK_REQUEST             Topic = 5
	Topic_BEACON_BLOCK_RESPONSE            Topic = 6
	Topic_CRYSTALLIZED_STATE_HASH_ANNOUNCE Topic = 7
	Topic_CRYSTALLIZED_STATE_REQUEST       Topic = 8
	Topic_CRYSTALLIZED_STATE_RESPONSE      Topic = 9
	Topic_ACTIVE_STATE_HASH_ANNOUNCE       Topic = 10
	Topic_ACTIVE_STATE_REQUEST             Topic = 11
	Topic_ACTIVE_STATE_RESPONSE            Topic = 12
)

var Topic_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "COLLATION_BODY_REQUEST",
	2:  "COLLATION_BODY_RESPONSE",
	3:  "TRANSACTIONS",
	4:  "BEACON_BLOCK_HASH_ANNOUNCE",
	5:  "BEACON_BLOCK_REQUEST",
	6:  "BEACON_BLOCK_RESPONSE",
	7:  "CRYSTALLIZED_STATE_HASH_ANNOUNCE",
	8:  "CRYSTALLIZED_STATE_REQUEST",
	9:  "CRYSTALLIZED_STATE_RESPONSE",
	10: "ACTIVE_STATE_HASH_ANNOUNCE",
	11: "ACTIVE_STATE_REQUEST",
	12: "ACTIVE_STATE_RESPONSE",
}
var Topic_value = map[string]int32{
	"UNKNOWN":                          0,
	"COLLATION_BODY_REQUEST":           1,
	"COLLATION_BODY_RESPONSE":          2,
	"TRANSACTIONS":                     3,
	"BEACON_BLOCK_HASH_ANNOUNCE":       4,
	"BEACON_BLOCK_REQUEST":             5,
	"BEACON_BLOCK_RESPONSE":            6,
	"CRYSTALLIZED_STATE_HASH_ANNOUNCE": 7,
	"CRYSTALLIZED_STATE_REQUEST":       8,
	"CRYSTALLIZED_STATE_RESPONSE":      9,
	"ACTIVE_STATE_HASH_ANNOUNCE":       10,
	"ACTIVE_STATE_REQUEST":             11,
	"ACTIVE_STATE_RESPONSE":            12,
}

func (x Topic) String() string {
	return proto.EnumName(Topic_name, int32(x))
}
func (Topic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_messages_923b4f96a23fc196, []int{0}
}

type CollationBodyRequest struct {
	ShardId              uint64   `protobuf:"varint,1,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	Period               uint64   `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	ChunkRoot            []byte   `protobuf:"bytes,3,opt,name=chunk_root,json=chunkRoot,proto3" json:"chunk_root,omitempty"`
	ProposerAddress      []byte   `protobuf:"bytes,4,opt,name=proposer_address,json=proposerAddress,proto3" json:"proposer_address,omitempty"`
	Signature            []byte   `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollationBodyRequest) Reset()         { *m = CollationBodyRequest{} }
func (m *CollationBodyRequest) String() string { return proto.CompactTextString(m) }
func (*CollationBodyRequest) ProtoMessage()    {}
func (*CollationBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_923b4f96a23fc196, []int{0}
}
func (m *CollationBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollationBodyRequest.Unmarshal(m, b)
}
func (m *CollationBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollationBodyRequest.Marshal(b, m, deterministic)
}
func (dst *CollationBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollationBodyRequest.Merge(dst, src)
}
func (m *CollationBodyRequest) XXX_Size() int {
	return xxx_messageInfo_CollationBodyRequest.Size(m)
}
func (m *CollationBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CollationBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CollationBodyRequest proto.InternalMessageInfo

func (m *CollationBodyRequest) GetShardId() uint64 {
	if m != nil {
		return m.ShardId
	}
	return 0
}

func (m *CollationBodyRequest) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *CollationBodyRequest) GetChunkRoot() []byte {
	if m != nil {
		return m.ChunkRoot
	}
	return nil
}

func (m *CollationBodyRequest) GetProposerAddress() []byte {
	if m != nil {
		return m.ProposerAddress
	}
	return nil
}

func (m *CollationBodyRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type CollationBodyResponse struct {
	HeaderHash           []byte   `protobuf:"bytes,1,opt,name=header_hash,json=headerHash,proto3" json:"header_hash,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollationBodyResponse) Reset()         { *m = CollationBodyResponse{} }
func (m *CollationBodyResponse) String() string { return proto.CompactTextString(m) }
func (*CollationBodyResponse) ProtoMessage()    {}
func (*CollationBodyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_923b4f96a23fc196, []int{1}
}
func (m *CollationBodyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollationBodyResponse.Unmarshal(m, b)
}
func (m *CollationBodyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollationBodyResponse.Marshal(b, m, deterministic)
}
func (dst *CollationBodyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollationBodyResponse.Merge(dst, src)
}
func (m *CollationBodyResponse) XXX_Size() int {
	return xxx_messageInfo_CollationBodyResponse.Size(m)
}
func (m *CollationBodyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CollationBodyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CollationBodyResponse proto.InternalMessageInfo

func (m *CollationBodyResponse) GetHeaderHash() []byte {
	if m != nil {
		return m.HeaderHash
	}
	return nil
}

func (m *CollationBodyResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Transaction struct {
	Nonce                uint64     `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	GasPrice             uint64     `protobuf:"varint,2,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	GasLimit             uint64     `protobuf:"varint,3,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	Recipient            []byte     `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Value                uint64     `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
	Input                []byte     `protobuf:"bytes,6,opt,name=input,proto3" json:"input,omitempty"`
	Signature            *Signature `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_923b4f96a23fc196, []int{2}
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

func (m *Transaction) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetGasPrice() uint64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *Transaction) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *Transaction) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *Transaction) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Transaction) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *Transaction) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Signature struct {
	V                    uint64   `protobuf:"varint,1,opt,name=v,proto3" json:"v,omitempty"`
	R                    uint64   `protobuf:"varint,2,opt,name=r,proto3" json:"r,omitempty"`
	S                    uint64   `protobuf:"varint,3,opt,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_923b4f96a23fc196, []int{3}
}
func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (dst *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(dst, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetV() uint64 {
	if m != nil {
		return m.V
	}
	return 0
}

func (m *Signature) GetR() uint64 {
	if m != nil {
		return m.R
	}
	return 0
}

func (m *Signature) GetS() uint64 {
	if m != nil {
		return m.S
	}
	return 0
}

func init() {
	proto.RegisterType((*CollationBodyRequest)(nil), "ethereum.sharding.p2p.v1.CollationBodyRequest")
	proto.RegisterType((*CollationBodyResponse)(nil), "ethereum.sharding.p2p.v1.CollationBodyResponse")
	proto.RegisterType((*Transaction)(nil), "ethereum.sharding.p2p.v1.Transaction")
	proto.RegisterType((*Signature)(nil), "ethereum.sharding.p2p.v1.Signature")
	proto.RegisterEnum("ethereum.sharding.p2p.v1.Topic", Topic_name, Topic_value)
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_messages_923b4f96a23fc196) }

var fileDescriptor_messages_923b4f96a23fc196 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xcb, 0x6e, 0xd3, 0x4c,
	0x18, 0xfd, 0xdd, 0xe6, 0xd2, 0x7c, 0xb1, 0x7e, 0xac, 0x51, 0x5b, 0xdc, 0x06, 0x68, 0x14, 0x58,
	0x04, 0x16, 0x96, 0x28, 0xe2, 0x01, 0x1c, 0xd7, 0x52, 0xa2, 0x5a, 0x76, 0xb1, 0x1d, 0x50, 0xd9,
	0x58, 0x53, 0x7b, 0x14, 0x8f, 0x48, 0x3c, 0x66, 0xc6, 0x8e, 0xd4, 0xc7, 0xe2, 0xb5, 0x58, 0xf1,
	0x08, 0x68, 0x7c, 0x29, 0x6d, 0x08, 0xbb, 0x9c, 0x8b, 0xce, 0xf9, 0xce, 0x44, 0x86, 0xff, 0x37,
	0x44, 0x08, 0xbc, 0x22, 0xc2, 0xc8, 0x39, 0x2b, 0x18, 0xd2, 0x49, 0x91, 0x12, 0x4e, 0xca, 0x8d,
	0x21, 0x52, 0xcc, 0x13, 0x9a, 0xad, 0x8c, 0xfc, 0x32, 0x37, 0xb6, 0xef, 0x27, 0x3f, 0x14, 0x38,
	0xb6, 0xd8, 0x7a, 0x8d, 0x0b, 0xca, 0xb2, 0x19, 0x4b, 0xee, 0x7d, 0xf2, 0xbd, 0x24, 0xa2, 0x40,
	0x67, 0x70, 0x54, 0x79, 0x23, 0x9a, 0xe8, 0xca, 0x58, 0x99, 0x76, 0xfc, 0x7e, 0x85, 0x17, 0x09,
	0x3a, 0x85, 0x5e, 0x4e, 0x38, 0x65, 0x89, 0x7e, 0x50, 0x09, 0x0d, 0x42, 0x2f, 0x01, 0xe2, 0xb4,
	0xcc, 0xbe, 0x45, 0x9c, 0xb1, 0x42, 0x3f, 0x1c, 0x2b, 0x53, 0xd5, 0x1f, 0x54, 0x8c, 0xcf, 0x58,
	0x81, 0xde, 0x82, 0x96, 0x73, 0x96, 0x33, 0x41, 0x78, 0x84, 0x93, 0x84, 0x13, 0x21, 0xf4, 0x4e,
	0x65, 0x7a, 0xd6, 0xf2, 0x66, 0x4d, 0xa3, 0x17, 0x30, 0x10, 0x74, 0x95, 0xe1, 0xa2, 0xe4, 0x44,
	0xef, 0xd6, 0x41, 0x0f, 0xc4, 0xc4, 0x81, 0x93, 0x9d, 0x93, 0x45, 0xce, 0x32, 0x41, 0xd0, 0x05,
	0x0c, 0x53, 0x82, 0x13, 0xc2, 0xa3, 0x14, 0x8b, 0xb4, 0x3a, 0x5b, 0xf5, 0xa1, 0xa6, 0xe6, 0x58,
	0xa4, 0x08, 0x41, 0xe7, 0x8e, 0x25, 0xf7, 0xd5, 0xdd, 0xaa, 0x5f, 0xfd, 0x9e, 0xfc, 0x54, 0x60,
	0x18, 0x72, 0x9c, 0x09, 0x1c, 0xcb, 0x40, 0x74, 0x0c, 0xdd, 0x8c, 0x65, 0x31, 0x69, 0x56, 0xd7,
	0x00, 0x8d, 0x60, 0xb0, 0xc2, 0x22, 0xca, 0x39, 0x8d, 0x49, 0x33, 0xfb, 0x68, 0x85, 0xc5, 0x8d,
	0xc4, 0xad, 0xb8, 0xa6, 0x1b, 0x5a, 0xef, 0xae, 0x45, 0x47, 0x62, 0xb9, 0x85, 0x93, 0x98, 0xe6,
	0x94, 0x64, 0x45, 0xb3, 0xf7, 0x0f, 0x21, 0xdb, 0xb6, 0x78, 0x5d, 0xd6, 0x2b, 0x3b, 0x7e, 0x0d,
	0x24, 0x4b, 0xb3, 0xbc, 0x2c, 0xf4, 0x5e, 0xe5, 0xaf, 0x01, 0x32, 0x1f, 0xbf, 0x4a, 0x7f, 0xac,
	0x4c, 0x87, 0x97, 0xaf, 0x8d, 0x7f, 0xfd, 0xb3, 0x46, 0xd0, 0x5a, 0x1f, 0x3f, 0xdd, 0x47, 0x18,
	0x3c, 0xf0, 0x48, 0x05, 0x65, 0xdb, 0xac, 0x54, 0xb6, 0x12, 0xf1, 0x66, 0x99, 0xc2, 0x25, 0x12,
	0xcd, 0x14, 0x45, 0xbc, 0xfb, 0x75, 0x00, 0xdd, 0x90, 0xe5, 0x34, 0x46, 0x43, 0xe8, 0x2f, 0xdd,
	0x6b, 0xd7, 0xfb, 0xe2, 0x6a, 0xff, 0xa1, 0x73, 0x38, 0xb5, 0x3c, 0xc7, 0x31, 0xc3, 0x85, 0xe7,
	0x46, 0x33, 0xef, 0xea, 0x36, 0xf2, 0xed, 0x4f, 0x4b, 0x3b, 0x08, 0x35, 0x05, 0x8d, 0xe0, 0xf9,
	0x5f, 0x5a, 0x70, 0xe3, 0xb9, 0x81, 0xad, 0x1d, 0x20, 0x0d, 0xd4, 0xd0, 0x37, 0xdd, 0xc0, 0xb4,
	0xa4, 0x1c, 0x68, 0x87, 0xe8, 0x15, 0x9c, 0xcf, 0x6c, 0xd3, 0x92, 0x5e, 0xc7, 0xb3, 0xae, 0xa3,
	0xb9, 0x19, 0xcc, 0x23, 0xd3, 0x75, 0xbd, 0xa5, 0x6b, 0xd9, 0x5a, 0x07, 0xe9, 0x70, 0xfc, 0x44,
	0x6f, 0x8b, 0xba, 0xe8, 0x0c, 0x4e, 0x76, 0x94, 0xa6, 0xa6, 0x87, 0xde, 0xc0, 0xd8, 0xf2, 0x6f,
	0x83, 0xd0, 0x74, 0x9c, 0xc5, 0x57, 0xfb, 0x2a, 0x0a, 0x42, 0x33, 0xb4, 0x77, 0xa2, 0xfb, 0xb2,
	0x7a, 0x8f, 0xab, 0x2d, 0x38, 0x42, 0x17, 0x30, 0xda, 0xab, 0x37, 0x35, 0x03, 0x19, 0x20, 0x87,
	0x7c, 0xb6, 0xf7, 0x16, 0x80, 0xbc, 0xfd, 0x89, 0xde, 0x46, 0x0f, 0xe5, 0xed, 0x3b, 0x4a, 0x13,
	0xaa, 0xde, 0xf5, 0xaa, 0x2f, 0xf7, 0xc3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0xf0, 0x6f,
	0x03, 0xcb, 0x03, 0x00, 0x00,
}
