// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: grpc.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ExampleGrpcReqVO struct {
	// 姓
	FirstName string `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	// 名
	LastName             string   `protobuf:"bytes,2,opt,name=LastName,proto3" json:"LastName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExampleGrpcReqVO) Reset()         { *m = ExampleGrpcReqVO{} }
func (m *ExampleGrpcReqVO) String() string { return proto.CompactTextString(m) }
func (*ExampleGrpcReqVO) ProtoMessage()    {}
func (*ExampleGrpcReqVO) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}
func (m *ExampleGrpcReqVO) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExampleGrpcReqVO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExampleGrpcReqVO.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExampleGrpcReqVO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExampleGrpcReqVO.Merge(m, src)
}
func (m *ExampleGrpcReqVO) XXX_Size() int {
	return m.Size()
}
func (m *ExampleGrpcReqVO) XXX_DiscardUnknown() {
	xxx_messageInfo_ExampleGrpcReqVO.DiscardUnknown(m)
}

var xxx_messageInfo_ExampleGrpcReqVO proto.InternalMessageInfo

func (m *ExampleGrpcReqVO) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *ExampleGrpcReqVO) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type ExampleGrpcRspVO struct {
	// 问候语
	Greetings            string   `protobuf:"bytes,1,opt,name=Greetings,proto3" json:"Greetings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExampleGrpcRspVO) Reset()         { *m = ExampleGrpcRspVO{} }
func (m *ExampleGrpcRspVO) String() string { return proto.CompactTextString(m) }
func (*ExampleGrpcRspVO) ProtoMessage()    {}
func (*ExampleGrpcRspVO) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}
func (m *ExampleGrpcRspVO) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExampleGrpcRspVO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExampleGrpcRspVO.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExampleGrpcRspVO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExampleGrpcRspVO.Merge(m, src)
}
func (m *ExampleGrpcRspVO) XXX_Size() int {
	return m.Size()
}
func (m *ExampleGrpcRspVO) XXX_DiscardUnknown() {
	xxx_messageInfo_ExampleGrpcRspVO.DiscardUnknown(m)
}

var xxx_messageInfo_ExampleGrpcRspVO proto.InternalMessageInfo

func (m *ExampleGrpcRspVO) GetGreetings() string {
	if m != nil {
		return m.Greetings
	}
	return ""
}

func init() {
	proto.RegisterType((*ExampleGrpcReqVO)(nil), "protobuf.ExampleGrpcReqVO")
	proto.RegisterType((*ExampleGrpcRspVO)(nil), "protobuf.ExampleGrpcRspVO")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5, 0x69, 0x4a, 0x3e, 0x5c,
	0x02, 0xae, 0x15, 0x89, 0xb9, 0x05, 0x39, 0xa9, 0xee, 0x45, 0x05, 0xc9, 0x41, 0xa9, 0x85, 0x61,
	0xfe, 0x42, 0x32, 0x5c, 0x9c, 0x6e, 0x99, 0x45, 0xc5, 0x25, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x08, 0x01, 0x21, 0x29, 0x2e, 0x0e, 0x9f, 0x44, 0xa8, 0x24, 0x13,
	0x58, 0x12, 0xce, 0x57, 0x32, 0x40, 0x35, 0xad, 0xb8, 0x00, 0x62, 0x9a, 0x7b, 0x51, 0x6a, 0x6a,
	0x49, 0x66, 0x5e, 0x7a, 0x31, 0xcc, 0x34, 0xb8, 0x80, 0x51, 0x18, 0x17, 0x1f, 0x54, 0x47, 0x70,
	0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x0b, 0x17, 0x97, 0x47, 0x6a, 0x4e, 0x4e, 0x7e, 0x78,
	0x7e, 0x51, 0x4e, 0x8a, 0x90, 0x94, 0x1e, 0xcc, 0xa9, 0x7a, 0xe8, 0xee, 0x94, 0xc2, 0x21, 0x07,
	0xb2, 0xd5, 0x49, 0xe0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63,
	0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0xac, 0xd8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x35,
	0x42, 0xa3, 0x62, 0x08, 0x01, 0x00, 0x00,
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
	HelloWorld(ctx context.Context, in *ExampleGrpcReqVO, opts ...grpc.CallOption) (*ExampleGrpcRspVO, error)
}

type exampleServiceClient struct {
	cc *grpc.ClientConn
}

func NewExampleServiceClient(cc *grpc.ClientConn) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) HelloWorld(ctx context.Context, in *ExampleGrpcReqVO, opts ...grpc.CallOption) (*ExampleGrpcRspVO, error) {
	out := new(ExampleGrpcRspVO)
	err := c.cc.Invoke(ctx, "/protobuf.ExampleService/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
type ExampleServiceServer interface {
	HelloWorld(context.Context, *ExampleGrpcReqVO) (*ExampleGrpcRspVO, error)
}

// UnimplementedExampleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (*UnimplementedExampleServiceServer) HelloWorld(ctx context.Context, req *ExampleGrpcReqVO) (*ExampleGrpcRspVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}

func RegisterExampleServiceServer(s *grpc.Server, srv ExampleServiceServer) {
	s.RegisterService(&_ExampleService_serviceDesc, srv)
}

func _ExampleService_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleGrpcReqVO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.ExampleService/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).HelloWorld(ctx, req.(*ExampleGrpcReqVO))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _ExampleService_HelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

func (m *ExampleGrpcReqVO) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExampleGrpcReqVO) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExampleGrpcReqVO) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.LastName) > 0 {
		i -= len(m.LastName)
		copy(dAtA[i:], m.LastName)
		i = encodeVarintGrpc(dAtA, i, uint64(len(m.LastName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FirstName) > 0 {
		i -= len(m.FirstName)
		copy(dAtA[i:], m.FirstName)
		i = encodeVarintGrpc(dAtA, i, uint64(len(m.FirstName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ExampleGrpcRspVO) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExampleGrpcRspVO) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExampleGrpcRspVO) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Greetings) > 0 {
		i -= len(m.Greetings)
		copy(dAtA[i:], m.Greetings)
		i = encodeVarintGrpc(dAtA, i, uint64(len(m.Greetings)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGrpc(dAtA []byte, offset int, v uint64) int {
	offset -= sovGrpc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExampleGrpcReqVO) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FirstName)
	if l > 0 {
		n += 1 + l + sovGrpc(uint64(l))
	}
	l = len(m.LastName)
	if l > 0 {
		n += 1 + l + sovGrpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ExampleGrpcRspVO) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Greetings)
	if l > 0 {
		n += 1 + l + sovGrpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGrpc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGrpc(x uint64) (n int) {
	return sovGrpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExampleGrpcReqVO) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpc
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
			return fmt.Errorf("proto: ExampleGrpcReqVO: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExampleGrpcReqVO: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
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
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
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
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExampleGrpcRspVO) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpc
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
			return fmt.Errorf("proto: ExampleGrpcRspVO: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExampleGrpcRspVO: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Greetings", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
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
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Greetings = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGrpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGrpc
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
					return 0, ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGrpc
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
				return 0, ErrInvalidLengthGrpc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGrpc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGrpc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGrpc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGrpc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGrpc = fmt.Errorf("proto: unexpected end of group")
)
