// Code generated by protoc-gen-go. DO NOT EDIT.
// source: levels.proto

package pbLevels

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

type LevelResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LevelResponse) Reset()         { *m = LevelResponse{} }
func (m *LevelResponse) String() string { return proto.CompactTextString(m) }
func (*LevelResponse) ProtoMessage()    {}
func (*LevelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dd466c55a41d32c, []int{0}
}

func (m *LevelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LevelResponse.Unmarshal(m, b)
}
func (m *LevelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LevelResponse.Marshal(b, m, deterministic)
}
func (m *LevelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LevelResponse.Merge(m, src)
}
func (m *LevelResponse) XXX_Size() int {
	return xxx_messageInfo_LevelResponse.Size(m)
}
func (m *LevelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LevelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LevelResponse proto.InternalMessageInfo

func (m *LevelResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type LevelRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LevelRequest) Reset()         { *m = LevelRequest{} }
func (m *LevelRequest) String() string { return proto.CompactTextString(m) }
func (*LevelRequest) ProtoMessage()    {}
func (*LevelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dd466c55a41d32c, []int{1}
}

func (m *LevelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LevelRequest.Unmarshal(m, b)
}
func (m *LevelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LevelRequest.Marshal(b, m, deterministic)
}
func (m *LevelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LevelRequest.Merge(m, src)
}
func (m *LevelRequest) XXX_Size() int {
	return xxx_messageInfo_LevelRequest.Size(m)
}
func (m *LevelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LevelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LevelRequest proto.InternalMessageInfo

func (m *LevelRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*LevelResponse)(nil), "levels.LevelResponse")
	proto.RegisterType((*LevelRequest)(nil), "levels.LevelRequest")
}

func init() { proto.RegisterFile("levels.proto", fileDescriptor_2dd466c55a41d32c) }

var fileDescriptor_2dd466c55a41d32c = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x49, 0x2d, 0x4b,
	0xcd, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xe4, 0xb9, 0x78,
	0x7d, 0x40, 0xac, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x3e, 0x2e, 0xa6, 0xcc,
	0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0xa6, 0xcc, 0x14, 0x25, 0x39, 0x2e, 0x1e, 0xa8,
	0x82, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x74, 0x79, 0x23, 0x67, 0x2e, 0x56, 0xb0, 0xbc, 0x90, 0x15,
	0x17, 0xa7, 0x7b, 0x6a, 0x09, 0x98, 0x5d, 0x2c, 0x24, 0xa2, 0x07, 0xb5, 0x0d, 0x59, 0xaf, 0x94,
	0x28, 0x9a, 0x28, 0xc4, 0x4a, 0x25, 0x06, 0x27, 0xae, 0x28, 0x8e, 0x82, 0x24, 0x88, 0xd6, 0x24,
	0x36, 0xb0, 0x03, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xa0, 0x9b, 0x1f, 0xb0, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LevelClient is the client API for Level service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LevelClient interface {
	GetLevels(ctx context.Context, in *LevelRequest, opts ...grpc.CallOption) (*LevelResponse, error)
}

type levelClient struct {
	cc *grpc.ClientConn
}

func NewLevelClient(cc *grpc.ClientConn) LevelClient {
	return &levelClient{cc}
}

func (c *levelClient) GetLevels(ctx context.Context, in *LevelRequest, opts ...grpc.CallOption) (*LevelResponse, error) {
	out := new(LevelResponse)
	err := c.cc.Invoke(ctx, "/levels.Level/GetLevels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LevelServer is the server API for Level service.
type LevelServer interface {
	GetLevels(context.Context, *LevelRequest) (*LevelResponse, error)
}

// UnimplementedLevelServer can be embedded to have forward compatible implementations.
type UnimplementedLevelServer struct {
}

func (*UnimplementedLevelServer) GetLevels(ctx context.Context, req *LevelRequest) (*LevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLevels not implemented")
}

func RegisterLevelServer(s *grpc.Server, srv LevelServer) {
	s.RegisterService(&_Level_serviceDesc, srv)
}

func _Level_GetLevels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LevelServer).GetLevels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/levels.Level/GetLevels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LevelServer).GetLevels(ctx, req.(*LevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Level_serviceDesc = grpc.ServiceDesc{
	ServiceName: "levels.Level",
	HandlerType: (*LevelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLevels",
			Handler:    _Level_GetLevels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "levels.proto",
}