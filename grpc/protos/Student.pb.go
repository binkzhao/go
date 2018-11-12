// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Student.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the message.
type Request struct {
	Sid                  int32    `protobuf:"varint,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_Student_c47dcbe45bfce483, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSid() int32 {
	if m != nil {
		return m.Sid
	}
	return 0
}

func (m *Request) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the ServiceController
type Response struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_Student_c47dcbe45bfce483, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterType((*Response)(nil), "proto.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StudentServiceClient interface {
	StudentAdd(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type studentServiceClient struct {
	cc *grpc.ClientConn
}

func NewStudentServiceClient(cc *grpc.ClientConn) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) StudentAdd(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.StudentService/StudentAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
type StudentServiceServer interface {
	StudentAdd(context.Context, *Request) (*Response, error)
}

func RegisterStudentServiceServer(s *grpc.Server, srv StudentServiceServer) {
	s.RegisterService(&_StudentService_serviceDesc, srv)
}

func _StudentService_StudentAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).StudentAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StudentService/StudentAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).StudentAdd(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _StudentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StudentAdd",
			Handler:    _StudentService_StudentAdd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Student.proto",
}

func init() { proto.RegisterFile("Student.proto", fileDescriptor_Student_c47dcbe45bfce483) }

var fileDescriptor_Student_c47dcbe45bfce483 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0x41, 0xae, 0x82, 0x30,
	0x10, 0x86, 0x1f, 0x0f, 0x51, 0x99, 0x44, 0x34, 0xb3, 0x22, 0xae, 0x08, 0x2b, 0x56, 0x68, 0xf4,
	0x04, 0x5c, 0xa1, 0x9c, 0x00, 0xe9, 0xa4, 0x61, 0x41, 0x8b, 0xb4, 0x78, 0x7e, 0xd3, 0x69, 0xe3,
	0x6a, 0xbe, 0xf9, 0x26, 0xf9, 0xe7, 0x87, 0x53, 0xef, 0x36, 0x49, 0xda, 0xb5, 0xcb, 0x6a, 0x9c,
	0xc1, 0x8c, 0x47, 0xdd, 0xc1, 0x41, 0xd0, 0x7b, 0x23, 0xeb, 0xf0, 0x02, 0xa9, 0x9d, 0x64, 0x99,
	0x54, 0x49, 0x93, 0x09, 0x8f, 0xde, 0x0c, 0x8a, 0xca, 0xff, 0x60, 0x06, 0x45, 0x88, 0xb0, 0xd3,
	0xc3, 0x4c, 0x65, 0x5a, 0x25, 0x4d, 0x2e, 0x98, 0xeb, 0x3b, 0x1c, 0x05, 0xd9, 0xc5, 0x68, 0xcb,
	0xf7, 0xd1, 0x48, 0x8a, 0x21, 0xcc, 0x3e, 0x65, 0xb6, 0x8a, 0x53, 0x72, 0xe1, 0xf1, 0xd1, 0x41,
	0x11, 0xcb, 0xf4, 0xb4, 0x7e, 0xa6, 0x91, 0xf0, 0x06, 0x10, 0x4d, 0x27, 0x25, 0x16, 0xa1, 0x63,
	0x1b, 0x9b, 0x5d, 0xcf, 0xbf, 0x3d, 0xbc, 0xa9, 0xff, 0x5e, 0x7b, 0x36, 0xcf, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc8, 0x35, 0xba, 0x8d, 0xd6, 0x00, 0x00, 0x00,
}