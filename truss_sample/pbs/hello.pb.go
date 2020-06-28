// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hello.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/metaverse/truss/deftree/googlethirdparty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SayHelloRequest struct {
	PersonName string `protobuf:"bytes,1,opt,name=person_name,json=personName,proto3" json:"person_name,omitempty"`
}

func (m *SayHelloRequest) Reset()         { *m = SayHelloRequest{} }
func (m *SayHelloRequest) String() string { return proto.CompactTextString(m) }
func (*SayHelloRequest) ProtoMessage()    {}
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}
func (m *SayHelloRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SayHelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SayHelloRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SayHelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloRequest.Merge(m, src)
}
func (m *SayHelloRequest) XXX_Size() int {
	return m.Size()
}
func (m *SayHelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloRequest proto.InternalMessageInfo

func (m *SayHelloRequest) GetPersonName() string {
	if m != nil {
		return m.PersonName
	}
	return ""
}

type SayHelloResponse struct {
	Sentence string `protobuf:"bytes,1,opt,name=sentence,proto3" json:"sentence,omitempty"`
}

func (m *SayHelloResponse) Reset()         { *m = SayHelloResponse{} }
func (m *SayHelloResponse) String() string { return proto.CompactTextString(m) }
func (*SayHelloResponse) ProtoMessage()    {}
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}
func (m *SayHelloResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SayHelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SayHelloResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SayHelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloResponse.Merge(m, src)
}
func (m *SayHelloResponse) XXX_Size() int {
	return m.Size()
}
func (m *SayHelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloResponse proto.InternalMessageInfo

func (m *SayHelloResponse) GetSentence() string {
	if m != nil {
		return m.Sentence
	}
	return ""
}

func init() {
	proto.RegisterType((*SayHelloRequest)(nil), "pb.SayHelloRequest")
	proto.RegisterType((*SayHelloResponse)(nil), "pb.SayHelloResponse")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x40, 0xb3, 0x29, 0xe4, 0xdc, 0x14, 0xca, 0x9a, 0x22, 0x5c, 0xb1, 0xca, 0x55, 0x56, 0xbb,
	0x18, 0xff, 0x40, 0x10, 0xad, 0x04, 0x63, 0x2d, 0xb2, 0x97, 0x8c, 0x7b, 0x07, 0x77, 0x3b, 0xeb,
	0xee, 0x5c, 0x20, 0xad, 0x5f, 0x20, 0xf8, 0x53, 0x96, 0x01, 0x1b, 0x4b, 0xb9, 0xf3, 0x43, 0xc4,
	0x9c, 0x51, 0xb4, 0x9c, 0x07, 0xef, 0x0d, 0x33, 0x7c, 0x54, 0x40, 0x55, 0xa1, 0xf2, 0x01, 0x09,
	0xc5, 0xd0, 0xe7, 0xe9, 0xb9, 0x2d, 0xa9, 0x68, 0x72, 0x35, 0xc7, 0x5a, 0xd7, 0x40, 0x66, 0x09,
	0x21, 0x82, 0xa6, 0xd0, 0xc4, 0xa8, 0x17, 0x70, 0x4f, 0x01, 0x40, 0x5b, 0x44, 0x5b, 0x01, 0x15,
	0x65, 0x58, 0x78, 0x13, 0x68, 0xa5, 0x8d, 0x73, 0x48, 0x86, 0x4a, 0x74, 0xb1, 0x4f, 0x65, 0x53,
	0xbe, 0x77, 0x63, 0x56, 0x97, 0x5f, 0xf1, 0x19, 0x3c, 0x34, 0x10, 0x49, 0x1c, 0xf2, 0x91, 0x87,
	0x10, 0xd1, 0xdd, 0x39, 0x53, 0xc3, 0x84, 0x1d, 0xb1, 0xe3, 0xdd, 0x19, 0xef, 0xd1, 0x95, 0xa9,
	0x21, 0x53, 0x7c, 0xff, 0xd7, 0x89, 0x1e, 0x5d, 0x04, 0x91, 0xf2, 0x24, 0x82, 0x23, 0x70, 0xf3,
	0xad, 0xf1, 0x33, 0x4f, 0x6f, 0x79, 0x72, 0x11, 0x00, 0xa8, 0x74, 0x56, 0x5c, 0xf3, 0x64, 0xeb,
	0x8a, 0x03, 0xe5, 0x73, 0xf5, 0x6f, 0x7b, 0x3a, 0xfe, 0x0b, 0xfb, 0x7c, 0x96, 0x3e, 0xbe, 0x7e,
	0x3c, 0x0f, 0xc7, 0x42, 0xe8, 0xe5, 0x89, 0xb6, 0xdf, 0x31, 0xbd, 0xf9, 0xc9, 0xd9, 0xe4, 0xa5,
	0x95, 0x6c, 0xdd, 0x4a, 0xf6, 0xde, 0x4a, 0xf6, 0xd4, 0xc9, 0xc1, 0xba, 0x93, 0x83, 0xb7, 0x4e,
	0x0e, 0xf2, 0x9d, 0xcd, 0x8d, 0xa7, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xca, 0x5e, 0x65,
	0x3d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetingClient is the client API for Greeting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetingClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
}

type greetingClient struct {
	cc *grpc.ClientConn
}

func NewGreetingClient(cc *grpc.ClientConn) GreetingClient {
	return &greetingClient{cc}
}

func (c *greetingClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, "/pb.Greeting/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetingServer is the server API for Greeting service.
type GreetingServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
}

// UnimplementedGreetingServer can be embedded to have forward compatible implementations.
type UnimplementedGreetingServer struct {
}

func (*UnimplementedGreetingServer) SayHello(ctx context.Context, req *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterGreetingServer(s *grpc.Server, srv GreetingServer) {
	s.RegisterService(&_Greeting_serviceDesc, srv)
}

func _Greeting_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Greeting/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Greeting",
	HandlerType: (*GreetingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeting_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

func (m *SayHelloRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SayHelloRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PersonName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHello(dAtA, i, uint64(len(m.PersonName)))
		i += copy(dAtA[i:], m.PersonName)
	}
	return i, nil
}

func (m *SayHelloResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SayHelloResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sentence) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHello(dAtA, i, uint64(len(m.Sentence)))
		i += copy(dAtA[i:], m.Sentence)
	}
	return i, nil
}

func encodeVarintHello(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SayHelloRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PersonName)
	if l > 0 {
		n += 1 + l + sovHello(uint64(l))
	}
	return n
}

func (m *SayHelloResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sentence)
	if l > 0 {
		n += 1 + l + sovHello(uint64(l))
	}
	return n
}

func sovHello(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHello(x uint64) (n int) {
	return sovHello(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SayHelloRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHello
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
			return fmt.Errorf("proto: SayHelloRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SayHelloRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PersonName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHello
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
				return ErrInvalidLengthHello
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHello
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PersonName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHello(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHello
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHello
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
func (m *SayHelloResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHello
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
			return fmt.Errorf("proto: SayHelloResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SayHelloResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sentence", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHello
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
				return ErrInvalidLengthHello
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHello
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sentence = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHello(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHello
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHello
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
func skipHello(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHello
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
					return 0, ErrIntOverflowHello
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHello
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
				return 0, ErrInvalidLengthHello
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthHello
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHello
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHello(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthHello
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHello = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHello   = fmt.Errorf("proto: integer overflow")
)
