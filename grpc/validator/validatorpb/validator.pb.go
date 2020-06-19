// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/validator/validatorpb/validator.proto

package validatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type P2PInfo struct {
	// ID is the unique identifier of the node on the P2P transport.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Addresses is the list of addresses at which the node can be reached.
	Addresses            []string `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *P2PInfo) Reset()         { *m = P2PInfo{} }
func (m *P2PInfo) String() string { return proto.CompactTextString(m) }
func (*P2PInfo) ProtoMessage()    {}
func (*P2PInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ee7217cd93dc09, []int{0}
}

func (m *P2PInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_P2PInfo.Unmarshal(m, b)
}
func (m *P2PInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_P2PInfo.Marshal(b, m, deterministic)
}
func (m *P2PInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_P2PInfo.Merge(m, src)
}
func (m *P2PInfo) XXX_Size() int {
	return xxx_messageInfo_P2PInfo.Size(m)
}
func (m *P2PInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_P2PInfo.DiscardUnknown(m)
}

var xxx_messageInfo_P2PInfo proto.InternalMessageInfo

func (m *P2PInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *P2PInfo) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type ConsensusAddress struct {
	// ID is public key identifying the node.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Address is the address at which the node can be reached.
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsensusAddress) Reset()         { *m = ConsensusAddress{} }
func (m *ConsensusAddress) String() string { return proto.CompactTextString(m) }
func (*ConsensusAddress) ProtoMessage()    {}
func (*ConsensusAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ee7217cd93dc09, []int{1}
}

func (m *ConsensusAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusAddress.Unmarshal(m, b)
}
func (m *ConsensusAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusAddress.Marshal(b, m, deterministic)
}
func (m *ConsensusAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusAddress.Merge(m, src)
}
func (m *ConsensusAddress) XXX_Size() int {
	return xxx_messageInfo_ConsensusAddress.Size(m)
}
func (m *ConsensusAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusAddress.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusAddress proto.InternalMessageInfo

func (m *ConsensusAddress) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ConsensusAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ConsensusInfo struct {
	// ID is the unique identifier of the node as a consensus member.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Addresses is the list of addresses at which the node can be reached.
	Addresses            []*ConsensusAddress `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ConsensusInfo) Reset()         { *m = ConsensusInfo{} }
func (m *ConsensusInfo) String() string { return proto.CompactTextString(m) }
func (*ConsensusInfo) ProtoMessage()    {}
func (*ConsensusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ee7217cd93dc09, []int{2}
}

func (m *ConsensusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusInfo.Unmarshal(m, b)
}
func (m *ConsensusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusInfo.Marshal(b, m, deterministic)
}
func (m *ConsensusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusInfo.Merge(m, src)
}
func (m *ConsensusInfo) XXX_Size() int {
	return xxx_messageInfo_ConsensusInfo.Size(m)
}
func (m *ConsensusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusInfo proto.InternalMessageInfo

func (m *ConsensusInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ConsensusInfo) GetAddresses() []*ConsensusAddress {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type Version struct {
	Major                uint32   `protobuf:"varint,1,opt,name=Major,proto3" json:"Major,omitempty"`
	Minor                uint32   `protobuf:"varint,2,opt,name=Minor,proto3" json:"Minor,omitempty"`
	Patch                uint32   `protobuf:"varint,3,opt,name=Patch,proto3" json:"Patch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ee7217cd93dc09, []int{3}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetMajor() uint32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *Version) GetMinor() uint32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *Version) GetPatch() uint32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

type Runtime struct {
	// ID is the public key identifying the runtime.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Version is the version of the runtime.
	Version              *Version `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Runtime) Reset()         { *m = Runtime{} }
func (m *Runtime) String() string { return proto.CompactTextString(m) }
func (*Runtime) ProtoMessage()    {}
func (*Runtime) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ee7217cd93dc09, []int{4}
}

func (m *Runtime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Runtime.Unmarshal(m, b)
}
func (m *Runtime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Runtime.Marshal(b, m, deterministic)
}
func (m *Runtime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Runtime.Merge(m, src)
}
func (m *Runtime) XXX_Size() int {
	return xxx_messageInfo_Runtime.Size(m)
}
func (m *Runtime) XXX_DiscardUnknown() {
	xxx_messageInfo_Runtime.DiscardUnknown(m)
}

var xxx_messageInfo_Runtime proto.InternalMessageInfo

func (m *Runtime) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Runtime) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type Node struct {
	// ID is the public key identifying the node.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// EntityID is the public key identifying the Entity controlling
	// the node.
	EntityId string `protobuf:"bytes,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	// Expiration is the epoch in which this node's commitment expires.
	Expiration uint64 `protobuf:"varint,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// P2P contains information for connecting to this node via P2P transport.
	P2P *P2PInfo `protobuf:"bytes,5,opt,name=p2p,proto3" json:"p2p,omitempty"`
	// Consensus contains information for connecting to this node as a
	// consensus member.
	Consensus *ConsensusInfo `protobuf:"bytes,6,opt,name=consensus,proto3" json:"consensus,omitempty"`
	// Runtimes are the node's runtimes.
	Runtimes []*Runtime `protobuf:"bytes,7,rep,name=runtimes,proto3" json:"runtimes,omitempty"`
	// Roles is a bitmask representing the node roles.
	Roles                uint32   `protobuf:"varint,8,opt,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ee7217cd93dc09, []int{5}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetEntityId() string {
	if m != nil {
		return m.EntityId
	}
	return ""
}

func (m *Node) GetExpiration() uint64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

func (m *Node) GetP2P() *P2PInfo {
	if m != nil {
		return m.P2P
	}
	return nil
}

func (m *Node) GetConsensus() *ConsensusInfo {
	if m != nil {
		return m.Consensus
	}
	return nil
}

func (m *Node) GetRuntimes() []*Runtime {
	if m != nil {
		return m.Runtimes
	}
	return nil
}

func (m *Node) GetRoles() uint32 {
	if m != nil {
		return m.Roles
	}
	return 0
}

type Validator struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VotingPower          int64    `protobuf:"varint,2,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	TendermintAddress    string   `protobuf:"bytes,4,opt,name=tendermint_address,json=tendermintAddress,proto3" json:"tendermint_address,omitempty"`
	Node                 *Node    `protobuf:"bytes,5,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ee7217cd93dc09, []int{6}
}

func (m *Validator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Validator.Unmarshal(m, b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return xxx_messageInfo_Validator.Size(m)
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Validator) GetVotingPower() int64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

func (m *Validator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Validator) GetTendermintAddress() string {
	if m != nil {
		return m.TendermintAddress
	}
	return ""
}

func (m *Validator) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type GetByHeightRequest struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByHeightRequest) Reset()         { *m = GetByHeightRequest{} }
func (m *GetByHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetByHeightRequest) ProtoMessage()    {}
func (*GetByHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ee7217cd93dc09, []int{7}
}

func (m *GetByHeightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByHeightRequest.Unmarshal(m, b)
}
func (m *GetByHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByHeightRequest.Marshal(b, m, deterministic)
}
func (m *GetByHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByHeightRequest.Merge(m, src)
}
func (m *GetByHeightRequest) XXX_Size() int {
	return xxx_messageInfo_GetByHeightRequest.Size(m)
}
func (m *GetByHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByHeightRequest proto.InternalMessageInfo

func (m *GetByHeightRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetByHeightResponse struct {
	Validators           []*Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetByHeightResponse) Reset()         { *m = GetByHeightResponse{} }
func (m *GetByHeightResponse) String() string { return proto.CompactTextString(m) }
func (*GetByHeightResponse) ProtoMessage()    {}
func (*GetByHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ee7217cd93dc09, []int{8}
}

func (m *GetByHeightResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByHeightResponse.Unmarshal(m, b)
}
func (m *GetByHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByHeightResponse.Marshal(b, m, deterministic)
}
func (m *GetByHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByHeightResponse.Merge(m, src)
}
func (m *GetByHeightResponse) XXX_Size() int {
	return xxx_messageInfo_GetByHeightResponse.Size(m)
}
func (m *GetByHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetByHeightResponse proto.InternalMessageInfo

func (m *GetByHeightResponse) GetValidators() []*Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func init() {
	proto.RegisterType((*P2PInfo)(nil), "validator.P2PInfo")
	proto.RegisterType((*ConsensusAddress)(nil), "validator.ConsensusAddress")
	proto.RegisterType((*ConsensusInfo)(nil), "validator.ConsensusInfo")
	proto.RegisterType((*Version)(nil), "validator.Version")
	proto.RegisterType((*Runtime)(nil), "validator.Runtime")
	proto.RegisterType((*Node)(nil), "validator.Node")
	proto.RegisterType((*Validator)(nil), "validator.Validator")
	proto.RegisterType((*GetByHeightRequest)(nil), "validator.GetByHeightRequest")
	proto.RegisterType((*GetByHeightResponse)(nil), "validator.GetByHeightResponse")
}

func init() {
	proto.RegisterFile("grpc/validator/validatorpb/validator.proto", fileDescriptor_a8ee7217cd93dc09)
}

var fileDescriptor_a8ee7217cd93dc09 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x51, 0x6f, 0xd3, 0x3c,
	0x14, 0xfd, 0xda, 0x74, 0x4d, 0x73, 0xfb, 0x0d, 0x86, 0x99, 0x90, 0xc5, 0x60, 0x1a, 0x81, 0x87,
	0x09, 0xad, 0xad, 0x54, 0x10, 0x08, 0x89, 0x17, 0x86, 0xc4, 0x36, 0x4d, 0x4c, 0x95, 0x91, 0xf6,
	0xb0, 0x97, 0x2a, 0x4d, 0xee, 0x5a, 0xc3, 0x6a, 0x07, 0xdb, 0xed, 0xd6, 0xff, 0xc4, 0xff, 0xe3,
	0x15, 0xc5, 0x49, 0x1a, 0x6f, 0xd9, 0xde, 0x7c, 0xcf, 0x3d, 0xbe, 0xe7, 0xf8, 0xd8, 0x32, 0xbc,
	0x9d, 0xaa, 0x34, 0x1e, 0x2c, 0xa3, 0x2b, 0x9e, 0x44, 0x46, 0xaa, 0x6a, 0x95, 0x4e, 0xaa, 0x75,
	0x3f, 0x55, 0xd2, 0x48, 0x12, 0xac, 0x81, 0xf0, 0x23, 0xf8, 0xa3, 0xe1, 0xe8, 0x44, 0x5c, 0x4a,
	0xf2, 0x08, 0x9a, 0x3c, 0xa1, 0x8d, 0xbd, 0xc6, 0x7e, 0xc0, 0x9a, 0x3c, 0x21, 0x2f, 0x20, 0x88,
	0x92, 0x44, 0xa1, 0xd6, 0xa8, 0x69, 0x73, 0xcf, 0xdb, 0x0f, 0x58, 0x05, 0x84, 0x9f, 0x61, 0xeb,
	0xab, 0x14, 0x1a, 0x85, 0x5e, 0xe8, 0x2f, 0x39, 0x5a, 0x9b, 0x40, 0xc1, 0x2f, 0x36, 0xd0, 0xa6,
	0x05, 0xcb, 0x32, 0xbc, 0x80, 0xcd, 0xf5, 0xee, 0x7b, 0xc5, 0x3f, 0xdd, 0x15, 0xef, 0x0e, 0x77,
	0xfa, 0xd5, 0x39, 0xee, 0x4a, 0xbb, 0xce, 0x4e, 0xc1, 0x3f, 0x47, 0xa5, 0xb9, 0x14, 0x64, 0x1b,
	0x36, 0xbe, 0x47, 0x3f, 0xa5, 0xb2, 0x83, 0x37, 0x59, 0x5e, 0x58, 0x94, 0x0b, 0xa9, 0xac, 0xa9,
	0x0c, 0xcd, 0x8a, 0x0c, 0x1d, 0x45, 0x26, 0x9e, 0x51, 0x2f, 0x47, 0x6d, 0x11, 0x1e, 0x81, 0xcf,
	0x16, 0xc2, 0xf0, 0x39, 0xd6, 0x2c, 0x1e, 0x80, 0xbf, 0xcc, 0x75, 0xec, 0xa0, 0xee, 0x90, 0x38,
	0x06, 0x0b, 0x07, 0xac, 0xa4, 0x84, 0x7f, 0x1b, 0xd0, 0x3a, 0x93, 0x49, 0x7d, 0xcc, 0x0e, 0x04,
	0x28, 0x0c, 0x37, 0xab, 0x31, 0x4f, 0x8a, 0x98, 0x3a, 0x39, 0x70, 0x92, 0x90, 0x5d, 0x00, 0xbc,
	0x49, 0xb9, 0x8a, 0x4c, 0x26, 0x93, 0x39, 0x6b, 0x31, 0x07, 0x21, 0x6f, 0xc0, 0x4b, 0x87, 0x29,
	0xdd, 0xa8, 0xe9, 0x17, 0x97, 0xca, 0xb2, 0x36, 0xf9, 0x00, 0x41, 0x5c, 0x06, 0x46, 0xdb, 0x96,
	0x4b, 0xef, 0x0b, 0xd3, 0xee, 0xa8, 0xa8, 0xa4, 0x0f, 0x1d, 0x95, 0x1f, 0x5e, 0x53, 0xdf, 0xde,
	0x81, 0x2b, 0x51, 0xe4, 0xc2, 0xd6, 0x9c, 0x2c, 0x42, 0x25, 0xaf, 0x50, 0xd3, 0x4e, 0x1e, 0xa1,
	0x2d, 0xc2, 0x3f, 0x0d, 0x08, 0xce, 0xcb, 0x5d, 0xb5, 0xe3, 0xbf, 0x82, 0xff, 0x97, 0xd2, 0x70,
	0x31, 0x1d, 0xa7, 0xf2, 0x1a, 0xf3, 0x3b, 0xf1, 0x58, 0x37, 0xc7, 0x46, 0x19, 0xe4, 0x3e, 0x23,
	0xef, 0xd6, 0x33, 0x22, 0x3d, 0x20, 0x06, 0x45, 0x82, 0x6a, 0xce, 0x85, 0x19, 0x97, 0xa4, 0x96,
	0x25, 0x3d, 0xa9, 0x3a, 0xe5, 0xfb, 0x7c, 0x0d, 0x2d, 0x21, 0x13, 0x2c, 0xe2, 0x7a, 0xec, 0x9c,
	0x25, 0xbb, 0x19, 0x66, 0x9b, 0xe1, 0x01, 0x90, 0x23, 0x34, 0x87, 0xab, 0x63, 0xe4, 0xd3, 0x99,
	0x61, 0xf8, 0x7b, 0x81, 0xda, 0x90, 0x67, 0xd0, 0x9e, 0x59, 0xc0, 0x5a, 0xf7, 0x58, 0x51, 0x85,
	0xa7, 0xf0, 0xf4, 0x16, 0x5b, 0xa7, 0x59, 0x7a, 0xe4, 0x3d, 0xc0, 0x7a, 0xb8, 0xa6, 0x0d, 0x9b,
	0xdd, 0xb6, 0xfb, 0x3c, 0xca, 0x15, 0x73, 0x78, 0xc3, 0x09, 0x6c, 0xad, 0x1b, 0x3f, 0x50, 0x2d,
	0x79, 0x8c, 0xe4, 0x0c, 0xba, 0x8e, 0x00, 0x79, 0xe9, 0x0c, 0xa9, 0xdb, 0x7c, 0xbe, 0xfb, 0x50,
	0x3b, 0xf7, 0x15, 0xfe, 0x77, 0x78, 0x7c, 0xf1, 0x6d, 0xca, 0xcd, 0x6c, 0x31, 0xe9, 0xc7, 0x72,
	0x3e, 0xb8, 0xe4, 0xd3, 0x39, 0x0a, 0xd3, 0x13, 0x68, 0xae, 0xa5, 0xfa, 0xa5, 0x07, 0x32, 0xd2,
	0x5c, 0xf7, 0x54, 0x1a, 0xf7, 0x52, 0x25, 0x6f, 0x56, 0x83, 0x87, 0x7f, 0x95, 0x49, 0xdb, 0x7e,
	0x26, 0xef, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x00, 0x98, 0xb7, 0x05, 0x7a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ValidatorServiceClient is the client API for ValidatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ValidatorServiceClient interface {
	GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error)
}

type validatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewValidatorServiceClient(cc grpc.ClientConnInterface) ValidatorServiceClient {
	return &validatorServiceClient{cc}
}

func (c *validatorServiceClient) GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error) {
	out := new(GetByHeightResponse)
	err := c.cc.Invoke(ctx, "/validator.ValidatorService/GetByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorServiceServer is the server API for ValidatorService service.
type ValidatorServiceServer interface {
	GetByHeight(context.Context, *GetByHeightRequest) (*GetByHeightResponse, error)
}

// UnimplementedValidatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedValidatorServiceServer struct {
}

func (*UnimplementedValidatorServiceServer) GetByHeight(ctx context.Context, req *GetByHeightRequest) (*GetByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByHeight not implemented")
}

func RegisterValidatorServiceServer(s *grpc.Server, srv ValidatorServiceServer) {
	s.RegisterService(&_ValidatorService_serviceDesc, srv)
}

func _ValidatorService_GetByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).GetByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/validator.ValidatorService/GetByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).GetByHeight(ctx, req.(*GetByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ValidatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "validator.ValidatorService",
	HandlerType: (*ValidatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByHeight",
			Handler:    _ValidatorService_GetByHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/validator/validatorpb/validator.proto",
}
