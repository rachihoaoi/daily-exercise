// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package pbs

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

type Req struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (m *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(m, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

func (m *Req) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Resp struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resp) Reset()         { *m = Resp{} }
func (m *Resp) String() string { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()    {}
func (*Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resp.Unmarshal(m, b)
}
func (m *Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resp.Marshal(b, m, deterministic)
}
func (m *Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resp.Merge(m, src)
}
func (m *Resp) XXX_Size() int {
	return xxx_messageInfo_Resp.Size(m)
}
func (m *Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Resp proto.InternalMessageInfo

func (m *Resp) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Req)(nil), "search.Req")
	proto.RegisterType((*Resp)(nil), "search.Resp")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4a, 0xc5, 0x30,
	0x14, 0x86, 0x09, 0xd5, 0x88, 0xc7, 0xba, 0x44, 0x28, 0xc5, 0x41, 0xa5, 0x53, 0x5d, 0x4a, 0xa9,
	0x83, 0x7b, 0x7d, 0x83, 0x04, 0x17, 0xb7, 0xb4, 0x3d, 0xd4, 0x40, 0x9a, 0xa6, 0x27, 0x51, 0xf0,
	0x3d, 0x7d, 0x20, 0xa9, 0xb7, 0x97, 0xbb, 0x5d, 0x3a, 0x7e, 0x9c, 0xef, 0xe3, 0xc0, 0x0f, 0xb7,
	0x13, 0x86, 0xa0, 0x47, 0xac, 0x3c, 0xcd, 0x71, 0x16, 0x3c, 0xa0, 0xa6, 0xfe, 0xb3, 0x78, 0x84,
	0x44, 0xe2, 0x22, 0x72, 0xb8, 0xda, 0xee, 0x39, 0x7b, 0x62, 0xe5, 0xb5, 0x3c, 0x62, 0xf1, 0x00,
	0x17, 0x12, 0x83, 0x17, 0x19, 0x70, 0xc2, 0xf0, 0x65, 0xe3, 0x26, 0x6c, 0xd4, 0xfc, 0x32, 0xe0,
	0x4a, 0x4f, 0xde, 0xa2, 0x78, 0x86, 0xf4, 0xdd, 0x69, 0xfa, 0x51, 0x48, 0xdf, 0xa6, 0x47, 0x71,
	0x53, 0x1d, 0x9e, 0x54, 0x12, 0x97, 0xfb, 0xf4, 0x04, 0xc1, 0x8b, 0x06, 0xee, 0x56, 0x0b, 0x49,
	0x99, 0x01, 0x55, 0x24, 0xd4, 0x93, 0x71, 0xe3, 0x99, 0xa2, 0x66, 0x6b, 0xf3, 0x66, 0x0d, 0xba,
	0xb8, 0xb7, 0x29, 0x99, 0x78, 0x85, 0xac, 0x35, 0x83, 0x21, 0xec, 0xa3, 0x99, 0x9d, 0xb6, 0xfb,
	0xb2, 0x9a, 0xb5, 0x97, 0x1f, 0x89, 0xef, 0x42, 0xc7, 0xff, 0xd7, 0x7a, 0xf9, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0x1f, 0x0d, 0xc1, 0x9c, 0x3e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SampleClient is the client API for Sample service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SampleClient interface {
	UnaryService(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
	ServerSideStreaming(ctx context.Context, in *Req, opts ...grpc.CallOption) (Sample_ServerSideStreamingClient, error)
	ClientSideStreaming(ctx context.Context, opts ...grpc.CallOption) (Sample_ClientSideStreamingClient, error)
	BidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (Sample_BidirectionalStreamingClient, error)
}

type sampleClient struct {
	cc *grpc.ClientConn
}

func NewSampleClient(cc *grpc.ClientConn) SampleClient {
	return &sampleClient{cc}
}

func (c *sampleClient) UnaryService(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/search.Sample/UnaryService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleClient) ServerSideStreaming(ctx context.Context, in *Req, opts ...grpc.CallOption) (Sample_ServerSideStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sample_serviceDesc.Streams[0], "/search.Sample/ServerSideStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &sampleServerSideStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sample_ServerSideStreamingClient interface {
	Recv() (*Resp, error)
	grpc.ClientStream
}

type sampleServerSideStreamingClient struct {
	grpc.ClientStream
}

func (x *sampleServerSideStreamingClient) Recv() (*Resp, error) {
	m := new(Resp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sampleClient) ClientSideStreaming(ctx context.Context, opts ...grpc.CallOption) (Sample_ClientSideStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sample_serviceDesc.Streams[1], "/search.Sample/ClientSideStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &sampleClientSideStreamingClient{stream}
	return x, nil
}

type Sample_ClientSideStreamingClient interface {
	Send(*Req) error
	CloseAndRecv() (*Resp, error)
	grpc.ClientStream
}

type sampleClientSideStreamingClient struct {
	grpc.ClientStream
}

func (x *sampleClientSideStreamingClient) Send(m *Req) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sampleClientSideStreamingClient) CloseAndRecv() (*Resp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Resp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sampleClient) BidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (Sample_BidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sample_serviceDesc.Streams[2], "/search.Sample/BidirectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &sampleBidirectionalStreamingClient{stream}
	return x, nil
}

type Sample_BidirectionalStreamingClient interface {
	Send(*Req) error
	Recv() (*Resp, error)
	grpc.ClientStream
}

type sampleBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *sampleBidirectionalStreamingClient) Send(m *Req) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sampleBidirectionalStreamingClient) Recv() (*Resp, error) {
	m := new(Resp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SampleServer is the server API for Sample service.
type SampleServer interface {
	UnaryService(context.Context, *Req) (*Resp, error)
	ServerSideStreaming(*Req, Sample_ServerSideStreamingServer) error
	ClientSideStreaming(Sample_ClientSideStreamingServer) error
	BidirectionalStreaming(Sample_BidirectionalStreamingServer) error
}

// UnimplementedSampleServer can be embedded to have forward compatible implementations.
type UnimplementedSampleServer struct {
}

func (*UnimplementedSampleServer) UnaryService(ctx context.Context, req *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryService not implemented")
}
func (*UnimplementedSampleServer) ServerSideStreaming(req *Req, srv Sample_ServerSideStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerSideStreaming not implemented")
}
func (*UnimplementedSampleServer) ClientSideStreaming(srv Sample_ClientSideStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientSideStreaming not implemented")
}
func (*UnimplementedSampleServer) BidirectionalStreaming(srv Sample_BidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreaming not implemented")
}

func RegisterSampleServer(s *grpc.Server, srv SampleServer) {
	s.RegisterService(&_Sample_serviceDesc, srv)
}

func _Sample_UnaryService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).UnaryService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.Sample/UnaryService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).UnaryService(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sample_ServerSideStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Req)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SampleServer).ServerSideStreaming(m, &sampleServerSideStreamingServer{stream})
}

type Sample_ServerSideStreamingServer interface {
	Send(*Resp) error
	grpc.ServerStream
}

type sampleServerSideStreamingServer struct {
	grpc.ServerStream
}

func (x *sampleServerSideStreamingServer) Send(m *Resp) error {
	return x.ServerStream.SendMsg(m)
}

func _Sample_ClientSideStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SampleServer).ClientSideStreaming(&sampleClientSideStreamingServer{stream})
}

type Sample_ClientSideStreamingServer interface {
	SendAndClose(*Resp) error
	Recv() (*Req, error)
	grpc.ServerStream
}

type sampleClientSideStreamingServer struct {
	grpc.ServerStream
}

func (x *sampleClientSideStreamingServer) SendAndClose(m *Resp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sampleClientSideStreamingServer) Recv() (*Req, error) {
	m := new(Req)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Sample_BidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SampleServer).BidirectionalStreaming(&sampleBidirectionalStreamingServer{stream})
}

type Sample_BidirectionalStreamingServer interface {
	Send(*Resp) error
	Recv() (*Req, error)
	grpc.ServerStream
}

type sampleBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *sampleBidirectionalStreamingServer) Send(m *Resp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sampleBidirectionalStreamingServer) Recv() (*Req, error) {
	m := new(Req)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Sample_serviceDesc = grpc.ServiceDesc{
	ServiceName: "search.Sample",
	HandlerType: (*SampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryService",
			Handler:    _Sample_UnaryService_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerSideStreaming",
			Handler:       _Sample_ServerSideStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientSideStreaming",
			Handler:       _Sample_ClientSideStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStreaming",
			Handler:       _Sample_BidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "message.proto",
}
