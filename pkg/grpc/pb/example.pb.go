// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package pb

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

type Example struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ExampleId            string   `protobuf:"bytes,2,opt,name=example_id,json=exampleId,proto3" json:"example_id,omitempty"`
	ExampleName          string   `protobuf:"bytes,3,opt,name=example_name,json=exampleName,proto3" json:"example_name,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CreateTime           string   `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Example) Reset()         { *m = Example{} }
func (m *Example) String() string { return proto.CompactTextString(m) }
func (*Example) ProtoMessage()    {}
func (*Example) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *Example) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Example.Unmarshal(m, b)
}
func (m *Example) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Example.Marshal(b, m, deterministic)
}
func (m *Example) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Example.Merge(m, src)
}
func (m *Example) XXX_Size() int {
	return xxx_messageInfo_Example.Size(m)
}
func (m *Example) XXX_DiscardUnknown() {
	xxx_messageInfo_Example.DiscardUnknown(m)
}

var xxx_messageInfo_Example proto.InternalMessageInfo

func (m *Example) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Example) GetExampleId() string {
	if m != nil {
		return m.ExampleId
	}
	return ""
}

func (m *Example) GetExampleName() string {
	if m != nil {
		return m.ExampleName
	}
	return ""
}

func (m *Example) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Example) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Example)(nil), "pb.Example")
	proto.RegisterType((*GetRequest)(nil), "pb.GetRequest")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x4e, 0x87, 0x30,
	0x10, 0xc6, 0x6d, 0xff, 0x82, 0xe1, 0x50, 0x86, 0x1b, 0x4c, 0x63, 0x34, 0x22, 0x71, 0x60, 0x62,
	0xd0, 0xc4, 0x37, 0x30, 0xc4, 0xc5, 0x01, 0xdd, 0x49, 0xa1, 0x37, 0x34, 0xb1, 0x52, 0xe1, 0x30,
	0x3e, 0x88, 0x0f, 0x6c, 0x28, 0x35, 0x3a, 0xde, 0xef, 0xfb, 0xbe, 0xe4, 0x77, 0x70, 0x46, 0x5f,
	0xda, 0xf9, 0x37, 0x6a, 0xfc, 0x3c, 0xf1, 0x84, 0xd2, 0x0f, 0xd5, 0xb7, 0x80, 0x93, 0xc7, 0x9d,
	0x62, 0x01, 0xd2, 0x1a, 0x25, 0x4a, 0x51, 0x27, 0x9d, 0xb4, 0x06, 0xaf, 0x00, 0xe2, 0xa0, 0xb7,
	0x46, 0xc9, 0x52, 0xd4, 0x59, 0x97, 0x45, 0xf2, 0x64, 0xf0, 0x06, 0x4e, 0x7f, 0xe3, 0x77, 0xed,
	0x48, 0x1d, 0x42, 0x21, 0x8f, 0xec, 0x59, 0x3b, 0xc2, 0x73, 0x48, 0x17, 0xd6, 0xbc, 0x2e, 0xea,
	0x38, 0x84, 0xf1, 0xc2, 0x6b, 0xc8, 0xc7, 0x99, 0x34, 0x53, 0xcf, 0xd6, 0x91, 0x4a, 0x42, 0x08,
	0x3b, 0x7a, 0xb5, 0x8e, 0xaa, 0x4b, 0x80, 0x96, 0xb8, 0xa3, 0x8f, 0x95, 0x16, 0xfe, 0x27, 0x96,
	0x6d, 0x62, 0x77, 0x0f, 0x50, 0x44, 0xe7, 0x17, 0x9a, 0x3f, 0xed, 0x48, 0x78, 0x0b, 0x87, 0x96,
	0x18, 0x8b, 0xc6, 0x0f, 0xcd, 0xdf, 0xf0, 0x22, 0xdf, 0xee, 0x58, 0xad, 0x8e, 0x86, 0x34, 0xfc,
	0x7d, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x11, 0xe5, 0x03, 0xeb, 0x08, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Example, error)
}

type exampleServiceClient struct {
	cc *grpc.ClientConn
}

func NewExampleServiceClient(cc *grpc.ClientConn) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Example, error) {
	out := new(Example)
	err := c.cc.Invoke(ctx, "/pb.ExampleService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
type ExampleServiceServer interface {
	Get(context.Context, *GetRequest) (*Example, error)
}

// UnimplementedExampleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (*UnimplementedExampleServiceServer) Get(ctx context.Context, req *GetRequest) (*Example, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterExampleServiceServer(s *grpc.Server, srv ExampleServiceServer) {
	s.RegisterService(&_ExampleService_serviceDesc, srv)
}

func _ExampleService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ExampleService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ExampleService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}
