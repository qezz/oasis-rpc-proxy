// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: grpc/transaction/transactionpb/transaction.proto

package transactionpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success           bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	SignatureVerified bool   `protobuf:"varint,2,opt,name=signature_verified,json=signatureVerified,proto3" json:"signature_verified,omitempty"`
	SanityChecked     bool   `protobuf:"varint,3,opt,name=sanity_checked,json=sanityChecked,proto3" json:"sanity_checked,omitempty"`
	Hash              string `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	PublicKey         string `protobuf:"bytes,5,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Signature         string `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	Nonce             uint64 `protobuf:"varint,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Fee               []byte `protobuf:"bytes,8,opt,name=fee,proto3" json:"fee,omitempty"`
	GasLimit          uint64 `protobuf:"varint,9,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasPrice          []byte `protobuf:"bytes,10,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	Method            string `protobuf:"bytes,11,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_grpc_transaction_transactionpb_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Transaction) GetSignatureVerified() bool {
	if x != nil {
		return x.SignatureVerified
	}
	return false
}

func (x *Transaction) GetSanityChecked() bool {
	if x != nil {
		return x.SanityChecked
	}
	return false
}

func (x *Transaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Transaction) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *Transaction) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Transaction) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Transaction) GetFee() []byte {
	if x != nil {
		return x.Fee
	}
	return nil
}

func (x *Transaction) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *Transaction) GetGasPrice() []byte {
	if x != nil {
		return x.GasPrice
	}
	return nil
}

func (x *Transaction) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type GetByHeightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *GetByHeightRequest) Reset() {
	*x = GetByHeightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByHeightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByHeightRequest) ProtoMessage() {}

func (x *GetByHeightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByHeightRequest.ProtoReflect.Descriptor instead.
func (*GetByHeightRequest) Descriptor() ([]byte, []int) {
	return file_grpc_transaction_transactionpb_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *GetByHeightRequest) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type GetByHeightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *GetByHeightResponse) Reset() {
	*x = GetByHeightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByHeightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByHeightResponse) ProtoMessage() {}

func (x *GetByHeightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByHeightResponse.ProtoReflect.Descriptor instead.
func (*GetByHeightResponse) Descriptor() ([]byte, []int) {
	return file_grpc_transaction_transactionpb_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *GetByHeightResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type BroadcastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxRaw string `protobuf:"bytes,1,opt,name=txRaw,proto3" json:"txRaw,omitempty"`
}

func (x *BroadcastRequest) Reset() {
	*x = BroadcastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRequest) ProtoMessage() {}

func (x *BroadcastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRequest.ProtoReflect.Descriptor instead.
func (*BroadcastRequest) Descriptor() ([]byte, []int) {
	return file_grpc_transaction_transactionpb_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *BroadcastRequest) GetTxRaw() string {
	if x != nil {
		return x.TxRaw
	}
	return ""
}

type BroadcastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submitted bool `protobuf:"varint,1,opt,name=submitted,proto3" json:"submitted,omitempty"`
}

func (x *BroadcastResponse) Reset() {
	*x = BroadcastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastResponse) ProtoMessage() {}

func (x *BroadcastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_transaction_transactionpb_transaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastResponse.ProtoReflect.Descriptor instead.
func (*BroadcastResponse) Descriptor() ([]byte, []int) {
	return file_grpc_transaction_transactionpb_transaction_proto_rawDescGZIP(), []int{4}
}

func (x *BroadcastResponse) GetSubmitted() bool {
	if x != nil {
		return x.Submitted
	}
	return false
}

var File_grpc_transaction_transactionpb_transaction_proto protoreflect.FileDescriptor

var file_grpc_transaction_transactionpb_transaction_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xc8, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x61, 0x6e, 0x69,
	0x74, 0x79, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x73, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x61, 0x73,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x53, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3c, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x28, 0x0a,
	0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x78, 0x52, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x78, 0x52, 0x61, 0x77, 0x22, 0x31, 0x0a, 0x11, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x32, 0xb6, 0x01, 0x0a, 0x12, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4c, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x1d,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x52, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x69, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x2f, 0x6f, 0x61, 0x73, 0x69, 0x73, 0x2d, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_transaction_transactionpb_transaction_proto_rawDescOnce sync.Once
	file_grpc_transaction_transactionpb_transaction_proto_rawDescData = file_grpc_transaction_transactionpb_transaction_proto_rawDesc
)

func file_grpc_transaction_transactionpb_transaction_proto_rawDescGZIP() []byte {
	file_grpc_transaction_transactionpb_transaction_proto_rawDescOnce.Do(func() {
		file_grpc_transaction_transactionpb_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_transaction_transactionpb_transaction_proto_rawDescData)
	})
	return file_grpc_transaction_transactionpb_transaction_proto_rawDescData
}

var file_grpc_transaction_transactionpb_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_transaction_transactionpb_transaction_proto_goTypes = []interface{}{
	(*Transaction)(nil),         // 0: transaction.Transaction
	(*GetByHeightRequest)(nil),  // 1: transaction.GetByHeightRequest
	(*GetByHeightResponse)(nil), // 2: transaction.GetByHeightResponse
	(*BroadcastRequest)(nil),    // 3: transaction.BroadcastRequest
	(*BroadcastResponse)(nil),   // 4: transaction.BroadcastResponse
}
var file_grpc_transaction_transactionpb_transaction_proto_depIdxs = []int32{
	0, // 0: transaction.GetByHeightResponse.transactions:type_name -> transaction.Transaction
	3, // 1: transaction.TransactionService.Broadcast:input_type -> transaction.BroadcastRequest
	1, // 2: transaction.TransactionService.GetByHeight:input_type -> transaction.GetByHeightRequest
	4, // 3: transaction.TransactionService.Broadcast:output_type -> transaction.BroadcastResponse
	2, // 4: transaction.TransactionService.GetByHeight:output_type -> transaction.GetByHeightResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_transaction_transactionpb_transaction_proto_init() }
func file_grpc_transaction_transactionpb_transaction_proto_init() {
	if File_grpc_transaction_transactionpb_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_transaction_transactionpb_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_transaction_transactionpb_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByHeightRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_transaction_transactionpb_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByHeightResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_transaction_transactionpb_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_transaction_transactionpb_transaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_transaction_transactionpb_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_transaction_transactionpb_transaction_proto_goTypes,
		DependencyIndexes: file_grpc_transaction_transactionpb_transaction_proto_depIdxs,
		MessageInfos:      file_grpc_transaction_transactionpb_transaction_proto_msgTypes,
	}.Build()
	File_grpc_transaction_transactionpb_transaction_proto = out.File
	file_grpc_transaction_transactionpb_transaction_proto_rawDesc = nil
	file_grpc_transaction_transactionpb_transaction_proto_goTypes = nil
	file_grpc_transaction_transactionpb_transaction_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionServiceClient interface {
	Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error)
	GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error) {
	out := new(BroadcastResponse)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error) {
	out := new(GetByHeightResponse)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/GetByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
type TransactionServiceServer interface {
	Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error)
	GetByHeight(context.Context, *GetByHeightRequest) (*GetByHeightResponse, error)
}

// UnimplementedTransactionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (*UnimplementedTransactionServiceServer) Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (*UnimplementedTransactionServiceServer) GetByHeight(context.Context, *GetByHeightRequest) (*GetByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByHeight not implemented")
}

func RegisterTransactionServiceServer(s *grpc.Server, srv TransactionServiceServer) {
	s.RegisterService(&_TransactionService_serviceDesc, srv)
}

func _TransactionService_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Broadcast(ctx, req.(*BroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/GetByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetByHeight(ctx, req.(*GetByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransactionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Broadcast",
			Handler:    _TransactionService_Broadcast_Handler,
		},
		{
			MethodName: "GetByHeight",
			Handler:    _TransactionService_GetByHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/transaction/transactionpb/transaction.proto",
}
