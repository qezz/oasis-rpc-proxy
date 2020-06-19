// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/chain/chainpb/chain.proto

package chainpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GetStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatusRequest) Reset()         { *m = GetStatusRequest{} }
func (m *GetStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatusRequest) ProtoMessage()    {}
func (*GetStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce694d07bea9bb01, []int{0}
}

func (m *GetStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatusRequest.Unmarshal(m, b)
}
func (m *GetStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatusRequest.Marshal(b, m, deterministic)
}
func (m *GetStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatusRequest.Merge(m, src)
}
func (m *GetStatusRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatusRequest.Size(m)
}
func (m *GetStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatusRequest proto.InternalMessageInfo

type GetStatusResponse struct {
	Id                   string               `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	GenesisTime          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=GenesisTime,proto3" json:"GenesisTime,omitempty"`
	GenesisHeight        int64                `protobuf:"varint,4,opt,name=GenesisHeight,proto3" json:"GenesisHeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetStatusResponse) Reset()         { *m = GetStatusResponse{} }
func (m *GetStatusResponse) String() string { return proto.CompactTextString(m) }
func (*GetStatusResponse) ProtoMessage()    {}
func (*GetStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce694d07bea9bb01, []int{1}
}

func (m *GetStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatusResponse.Unmarshal(m, b)
}
func (m *GetStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatusResponse.Marshal(b, m, deterministic)
}
func (m *GetStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatusResponse.Merge(m, src)
}
func (m *GetStatusResponse) XXX_Size() int {
	return xxx_messageInfo_GetStatusResponse.Size(m)
}
func (m *GetStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatusResponse proto.InternalMessageInfo

func (m *GetStatusResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetStatusResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetStatusResponse) GetGenesisTime() *timestamp.Timestamp {
	if m != nil {
		return m.GenesisTime
	}
	return nil
}

func (m *GetStatusResponse) GetGenesisHeight() int64 {
	if m != nil {
		return m.GenesisHeight
	}
	return 0
}

type GetHeadRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHeadRequest) Reset()         { *m = GetHeadRequest{} }
func (m *GetHeadRequest) String() string { return proto.CompactTextString(m) }
func (*GetHeadRequest) ProtoMessage()    {}
func (*GetHeadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce694d07bea9bb01, []int{2}
}

func (m *GetHeadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHeadRequest.Unmarshal(m, b)
}
func (m *GetHeadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHeadRequest.Marshal(b, m, deterministic)
}
func (m *GetHeadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHeadRequest.Merge(m, src)
}
func (m *GetHeadRequest) XXX_Size() int {
	return xxx_messageInfo_GetHeadRequest.Size(m)
}
func (m *GetHeadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHeadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHeadRequest proto.InternalMessageInfo

type GetHeadResponse struct {
	Height               int64                `protobuf:"varint,1,opt,name=Height,proto3" json:"Height,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,2,opt,name=Time,proto3" json:"Time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetHeadResponse) Reset()         { *m = GetHeadResponse{} }
func (m *GetHeadResponse) String() string { return proto.CompactTextString(m) }
func (*GetHeadResponse) ProtoMessage()    {}
func (*GetHeadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce694d07bea9bb01, []int{3}
}

func (m *GetHeadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHeadResponse.Unmarshal(m, b)
}
func (m *GetHeadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHeadResponse.Marshal(b, m, deterministic)
}
func (m *GetHeadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHeadResponse.Merge(m, src)
}
func (m *GetHeadResponse) XXX_Size() int {
	return xxx_messageInfo_GetHeadResponse.Size(m)
}
func (m *GetHeadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHeadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHeadResponse proto.InternalMessageInfo

func (m *GetHeadResponse) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *GetHeadResponse) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type GetMetaByHeightRequest struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMetaByHeightRequest) Reset()         { *m = GetMetaByHeightRequest{} }
func (m *GetMetaByHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetMetaByHeightRequest) ProtoMessage()    {}
func (*GetMetaByHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce694d07bea9bb01, []int{4}
}

func (m *GetMetaByHeightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMetaByHeightRequest.Unmarshal(m, b)
}
func (m *GetMetaByHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMetaByHeightRequest.Marshal(b, m, deterministic)
}
func (m *GetMetaByHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMetaByHeightRequest.Merge(m, src)
}
func (m *GetMetaByHeightRequest) XXX_Size() int {
	return xxx_messageInfo_GetMetaByHeightRequest.Size(m)
}
func (m *GetMetaByHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMetaByHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMetaByHeightRequest proto.InternalMessageInfo

func (m *GetMetaByHeightRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetMetaByHeightResponse struct {
	Height               int64                `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	AppVersion           uint64               `protobuf:"varint,3,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	BlockVersion         uint64               `protobuf:"varint,4,opt,name=block_version,json=blockVersion,proto3" json:"block_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetMetaByHeightResponse) Reset()         { *m = GetMetaByHeightResponse{} }
func (m *GetMetaByHeightResponse) String() string { return proto.CompactTextString(m) }
func (*GetMetaByHeightResponse) ProtoMessage()    {}
func (*GetMetaByHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce694d07bea9bb01, []int{5}
}

func (m *GetMetaByHeightResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMetaByHeightResponse.Unmarshal(m, b)
}
func (m *GetMetaByHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMetaByHeightResponse.Marshal(b, m, deterministic)
}
func (m *GetMetaByHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMetaByHeightResponse.Merge(m, src)
}
func (m *GetMetaByHeightResponse) XXX_Size() int {
	return xxx_messageInfo_GetMetaByHeightResponse.Size(m)
}
func (m *GetMetaByHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMetaByHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMetaByHeightResponse proto.InternalMessageInfo

func (m *GetMetaByHeightResponse) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *GetMetaByHeightResponse) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *GetMetaByHeightResponse) GetAppVersion() uint64 {
	if m != nil {
		return m.AppVersion
	}
	return 0
}

func (m *GetMetaByHeightResponse) GetBlockVersion() uint64 {
	if m != nil {
		return m.BlockVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*GetStatusRequest)(nil), "chain.GetStatusRequest")
	proto.RegisterType((*GetStatusResponse)(nil), "chain.GetStatusResponse")
	proto.RegisterType((*GetHeadRequest)(nil), "chain.GetHeadRequest")
	proto.RegisterType((*GetHeadResponse)(nil), "chain.GetHeadResponse")
	proto.RegisterType((*GetMetaByHeightRequest)(nil), "chain.GetMetaByHeightRequest")
	proto.RegisterType((*GetMetaByHeightResponse)(nil), "chain.GetMetaByHeightResponse")
}

func init() { proto.RegisterFile("grpc/chain/chainpb/chain.proto", fileDescriptor_ce694d07bea9bb01) }

var fileDescriptor_ce694d07bea9bb01 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xd9, 0xd4, 0x14, 0x75, 0xd2, 0x96, 0xb2, 0x12, 0xa9, 0x65, 0x89, 0x36, 0x32, 0x1c,
	0x72, 0x89, 0x8d, 0xca, 0x0d, 0x21, 0x54, 0x95, 0x83, 0xdb, 0x03, 0x1c, 0xdc, 0x0a, 0x09, 0x2e,
	0x68, 0xed, 0x4c, 0xed, 0x55, 0x6b, 0xef, 0xe2, 0x5d, 0x17, 0xfa, 0x2c, 0x3c, 0x00, 0xaf, 0xc5,
	0xa3, 0x20, 0xef, 0xda, 0x4e, 0x4c, 0x82, 0xe0, 0x92, 0x78, 0xfe, 0xfd, 0x77, 0xe6, 0xf3, 0xcc,
	0x18, 0x8e, 0xb2, 0x4a, 0xa6, 0x61, 0x9a, 0x33, 0x5e, 0xda, 0x5f, 0x99, 0xd8, 0xff, 0x40, 0x56,
	0x42, 0x0b, 0xfa, 0xd0, 0x04, 0xde, 0x71, 0x26, 0x44, 0x76, 0x8b, 0xa1, 0x11, 0x93, 0xfa, 0x3a,
	0xd4, 0xbc, 0x40, 0xa5, 0x59, 0x21, 0xad, 0xcf, 0xa7, 0x70, 0x10, 0xa1, 0xbe, 0xd4, 0x4c, 0xd7,
	0x2a, 0xc6, 0xaf, 0x35, 0x2a, 0xed, 0xff, 0x20, 0xf0, 0x64, 0x45, 0x54, 0x52, 0x94, 0x0a, 0xe9,
	0x3e, 0x8c, 0x2e, 0x16, 0x2e, 0x99, 0x92, 0xd9, 0x4e, 0x3c, 0xba, 0x58, 0x50, 0x0a, 0xce, 0x07,
	0x56, 0xa0, 0x3b, 0x32, 0x8a, 0x79, 0xa6, 0x6f, 0x60, 0x1c, 0x61, 0x89, 0x8a, 0xab, 0x2b, 0x5e,
	0xa0, 0xbb, 0x35, 0x25, 0xb3, 0xf1, 0x89, 0x17, 0x58, 0x88, 0xa0, 0x83, 0x08, 0xae, 0x3a, 0x88,
	0x78, 0xd5, 0x4e, 0x5f, 0xc0, 0x5e, 0x1b, 0x9e, 0x23, 0xcf, 0x72, 0xed, 0x3a, 0x53, 0x32, 0xdb,
	0x8a, 0x87, 0xa2, 0x7f, 0x00, 0xfb, 0x11, 0xea, 0x73, 0x64, 0x8b, 0x8e, 0xf7, 0x13, 0x3c, 0xee,
	0x95, 0x16, 0x76, 0x02, 0xdb, 0x6d, 0x0e, 0x62, 0x72, 0xb4, 0x11, 0x0d, 0xc0, 0x31, 0x64, 0xa3,
	0x7f, 0x92, 0x19, 0x9f, 0xff, 0x12, 0x26, 0x11, 0xea, 0xf7, 0xa8, 0xd9, 0xd9, 0xbd, 0x4d, 0xd1,
	0x16, 0x6d, 0x2a, 0xe4, 0x83, 0x0a, 0x36, 0xf2, 0x7f, 0x12, 0x38, 0x5c, 0xbb, 0xb2, 0xa4, 0xda,
	0x74, 0xa7, 0xa1, 0xd2, 0xff, 0x49, 0xd5, 0xf8, 0xe8, 0x31, 0x8c, 0x99, 0x94, 0x5f, 0xee, 0xb0,
	0x52, 0x5c, 0x94, 0xa6, 0xcd, 0x4e, 0x0c, 0x4c, 0xca, 0x8f, 0x56, 0xa1, 0xcf, 0x61, 0x2f, 0xb9,
	0x15, 0xe9, 0x4d, 0x6f, 0x71, 0x8c, 0x65, 0xd7, 0x88, 0xad, 0xe9, 0xe4, 0x17, 0x81, 0xdd, 0x77,
	0xcd, 0x96, 0x5c, 0x62, 0x75, 0xc7, 0x53, 0xa4, 0xaf, 0xe1, 0x51, 0xdb, 0x47, 0xfa, 0x34, 0xb0,
	0xcb, 0x34, 0xec, 0xb4, 0x37, 0xf9, 0x53, 0xb6, 0x2f, 0xe6, 0x3f, 0xa0, 0xa7, 0xb0, 0xd3, 0xaf,
	0x0c, 0x3d, 0x5c, 0xda, 0x06, 0x9b, 0xe5, 0xb9, 0xeb, 0x07, 0x7d, 0x86, 0xd8, 0x4c, 0x71, 0xb5,
	0x6f, 0xf4, 0xd9, 0xd2, 0xbe, 0x61, 0x04, 0xde, 0xd1, 0xdf, 0x8e, 0xbb, 0x9c, 0x67, 0xa7, 0x9f,
	0xdf, 0x66, 0x5c, 0xe7, 0x75, 0x12, 0xa4, 0xa2, 0x08, 0xaf, 0x79, 0x56, 0x60, 0xa9, 0xe7, 0x25,
	0xea, 0x6f, 0xa2, 0xba, 0x51, 0xa1, 0x60, 0x8a, 0xab, 0x79, 0x25, 0xd3, 0xb9, 0xac, 0xc4, 0xf7,
	0xfb, 0x70, 0xfd, 0x9b, 0x4a, 0xb6, 0xcd, 0x10, 0x5e, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x58,
	0xdb, 0xd3, 0x54, 0x70, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChainServiceClient is the client API for ChainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChainServiceClient interface {
	GetHead(ctx context.Context, in *GetHeadRequest, opts ...grpc.CallOption) (*GetHeadResponse, error)
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
	GetMetaByHeight(ctx context.Context, in *GetMetaByHeightRequest, opts ...grpc.CallOption) (*GetMetaByHeightResponse, error)
}

type chainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChainServiceClient(cc grpc.ClientConnInterface) ChainServiceClient {
	return &chainServiceClient{cc}
}

func (c *chainServiceClient) GetHead(ctx context.Context, in *GetHeadRequest, opts ...grpc.CallOption) (*GetHeadResponse, error) {
	out := new(GetHeadResponse)
	err := c.cc.Invoke(ctx, "/chain.ChainService/GetHead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/chain.ChainService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) GetMetaByHeight(ctx context.Context, in *GetMetaByHeightRequest, opts ...grpc.CallOption) (*GetMetaByHeightResponse, error) {
	out := new(GetMetaByHeightResponse)
	err := c.cc.Invoke(ctx, "/chain.ChainService/GetMetaByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChainServiceServer is the server API for ChainService service.
type ChainServiceServer interface {
	GetHead(context.Context, *GetHeadRequest) (*GetHeadResponse, error)
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	GetMetaByHeight(context.Context, *GetMetaByHeightRequest) (*GetMetaByHeightResponse, error)
}

// UnimplementedChainServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChainServiceServer struct {
}

func (*UnimplementedChainServiceServer) GetHead(ctx context.Context, req *GetHeadRequest) (*GetHeadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHead not implemented")
}
func (*UnimplementedChainServiceServer) GetStatus(ctx context.Context, req *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (*UnimplementedChainServiceServer) GetMetaByHeight(ctx context.Context, req *GetMetaByHeightRequest) (*GetMetaByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetaByHeight not implemented")
}

func RegisterChainServiceServer(s *grpc.Server, srv ChainServiceServer) {
	s.RegisterService(&_ChainService_serviceDesc, srv)
}

func _ChainService_GetHead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).GetHead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.ChainService/GetHead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).GetHead(ctx, req.(*GetHeadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.ChainService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_GetMetaByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetaByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).GetMetaByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.ChainService/GetMetaByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).GetMetaByHeight(ctx, req.(*GetMetaByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chain.ChainService",
	HandlerType: (*ChainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHead",
			Handler:    _ChainService_GetHead_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _ChainService_GetStatus_Handler,
		},
		{
			MethodName: "GetMetaByHeight",
			Handler:    _ChainService_GetMetaByHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/chain/chainpb/chain.proto",
}
