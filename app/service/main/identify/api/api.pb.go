// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// TokenReq request param for rpc TokenInfo
type GetTokenInfoReq struct {
	// user access token
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" form:"access_key" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenInfoReq) Reset()         { *m = GetTokenInfoReq{} }
func (m *GetTokenInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetTokenInfoReq) ProtoMessage()    {}
func (*GetTokenInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *GetTokenInfoReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTokenInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTokenInfoReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTokenInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenInfoReq.Merge(m, src)
}
func (m *GetTokenInfoReq) XXX_Size() int {
	return m.Size()
}
func (m *GetTokenInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenInfoReq proto.InternalMessageInfo

func (m *GetTokenInfoReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// TokenReply reply val for rpc TokenInfo
type GetTokenInfoReply struct {
	// 用户是否登录
	IsLogin bool `protobuf:"varint,1,opt,name=is_login,json=isLogin,proto3" json:"is_login"`
	// user mid
	Mid int64 `protobuf:"varint,2,opt,name=mid,proto3" json:"mid"`
	//username
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username"`
	// cookie csrf
	// when token reqest this field is empty
	Csrf string `protobuf:"bytes,4,opt,name=csrf,proto3" json:"csrfToken"`
	// expire time(unix timestamp)
	Expires              int32    `protobuf:"varint,5,opt,name=expires,proto3" json:"expires"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenInfoReply) Reset()         { *m = GetTokenInfoReply{} }
func (m *GetTokenInfoReply) String() string { return proto.CompactTextString(m) }
func (*GetTokenInfoReply) ProtoMessage()    {}
func (*GetTokenInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *GetTokenInfoReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTokenInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTokenInfoReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTokenInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenInfoReply.Merge(m, src)
}
func (m *GetTokenInfoReply) XXX_Size() int {
	return m.Size()
}
func (m *GetTokenInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenInfoReply proto.InternalMessageInfo

func (m *GetTokenInfoReply) GetIsLogin() bool {
	if m != nil {
		return m.IsLogin
	}
	return false
}

func (m *GetTokenInfoReply) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *GetTokenInfoReply) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetTokenInfoReply) GetCsrf() string {
	if m != nil {
		return m.Csrf
	}
	return ""
}

func (m *GetTokenInfoReply) GetExpires() int32 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func init() {
	proto.RegisterType((*GetTokenInfoReq)(nil), "passport.service.identify.v1.GetTokenInfoReq")
	proto.RegisterType((*GetTokenInfoReply)(nil), "passport.service.identify.v1.GetTokenInfoReply")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0xef, 0xf4, 0xcf, 0x6d, 0x3a, 0xb7, 0x97, 0xcb, 0x1d, 0xee, 0x22, 0xb7, 0x48, 0x12,
	0x03, 0xc5, 0xb8, 0x68, 0x4a, 0x75, 0xd7, 0x8d, 0x90, 0x8d, 0x14, 0x5c, 0x0d, 0xae, 0xdc, 0x94,
	0x34, 0x99, 0xc4, 0xa1, 0x49, 0x26, 0x9d, 0x49, 0x82, 0xf1, 0x49, 0x7c, 0x24, 0xc1, 0x8d, 0x4f,
	0x10, 0xa4, 0xee, 0xb2, 0xf4, 0x09, 0x24, 0x53, 0x5b, 0xd4, 0x85, 0xe0, 0x26, 0x7c, 0xbf, 0x2f,
	0xe7, 0x7c, 0x9c, 0x39, 0x07, 0xf6, 0xdd, 0x94, 0xda, 0x29, 0x67, 0x19, 0x43, 0x07, 0xa9, 0x2b,
	0x44, 0xca, 0x78, 0x66, 0x0b, 0xc2, 0x0b, 0xea, 0x11, 0x9b, 0xfa, 0x24, 0xc9, 0x68, 0x50, 0xda,
	0xc5, 0x74, 0x38, 0x0e, 0x69, 0x76, 0x9d, 0x2f, 0x6d, 0x8f, 0xc5, 0x93, 0x90, 0x85, 0x6c, 0x22,
	0x9b, 0x96, 0x79, 0x20, 0x49, 0x82, 0x54, 0xdb, 0x30, 0x13, 0xc3, 0x3f, 0xe7, 0x24, 0xbb, 0x64,
	0x2b, 0x92, 0xcc, 0x93, 0x80, 0x61, 0xb2, 0x46, 0x67, 0xb0, 0x9b, 0x35, 0xac, 0x02, 0x03, 0x58,
	0x7d, 0xe7, 0xf8, 0xa5, 0xd2, 0x47, 0x01, 0xe3, 0xf1, 0xcc, 0x74, 0x3d, 0x8f, 0x08, 0xb1, 0x58,
	0x91, 0xd2, 0x34, 0x0a, 0x37, 0xa2, 0xbe, 0x9b, 0x91, 0x99, 0xc9, 0xc9, 0x3a, 0xa7, 0x9c, 0xf8,
	0x26, 0xde, 0xf6, 0x99, 0x0f, 0x00, 0xfe, 0xfd, 0x18, 0x9a, 0x46, 0x25, 0x3a, 0x82, 0x0a, 0x15,
	0x8b, 0x88, 0x85, 0x74, 0x9b, 0xac, 0x38, 0x83, 0xba, 0xd2, 0xf7, 0x1e, 0xee, 0x51, 0x71, 0xd1,
	0x08, 0xf4, 0x1f, 0xb6, 0x63, 0xea, 0xab, 0x2d, 0x03, 0x58, 0x6d, 0xa7, 0x57, 0x57, 0x7a, 0x83,
	0xb8, 0xf9, 0x20, 0x0b, 0x2a, 0xb9, 0x20, 0x3c, 0x71, 0x63, 0xa2, 0xb6, 0xe5, 0x74, 0x32, 0x63,
	0xe7, 0xe1, 0xbd, 0x42, 0x87, 0xb0, 0xe3, 0x09, 0x1e, 0xa8, 0x1d, 0x59, 0xf5, 0xbb, 0xae, 0xf4,
	0x7e, 0xc3, 0x72, 0x26, 0x2c, 0x7f, 0xa1, 0x11, 0xec, 0x91, 0x9b, 0x94, 0x72, 0x22, 0xd4, 0xae,
	0x01, 0xac, 0xae, 0xf3, 0xab, 0xae, 0xf4, 0x9d, 0x85, 0x77, 0xe2, 0xe4, 0x16, 0x2a, 0xf3, 0xb7,
	0xfd, 0xa2, 0x04, 0x0e, 0xde, 0x3f, 0x0c, 0x8d, 0xed, 0xaf, 0x6e, 0x61, 0x7f, 0xda, 0xec, 0x70,
	0xf2, 0x9d, 0xf2, 0x34, 0x2a, 0x9d, 0x7f, 0xf7, 0x1b, 0x0d, 0x3c, 0x6e, 0x34, 0xf0, 0xb4, 0xd1,
	0xc0, 0xdd, 0xb3, 0xf6, 0xe3, 0xaa, 0x55, 0x4c, 0x97, 0x3f, 0xe5, 0xe9, 0x4e, 0x5f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x93, 0x27, 0xc6, 0x0a, 0x14, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentifyClient is the client API for Identify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentifyClient interface {
	// TokenInfo identify info by token.
	GetTokenInfo(ctx context.Context, in *GetTokenInfoReq, opts ...grpc.CallOption) (*GetTokenInfoReply, error)
}

type identifyClient struct {
	cc *grpc.ClientConn
}

func NewIdentifyClient(cc *grpc.ClientConn) IdentifyClient {
	return &identifyClient{cc}
}

func (c *identifyClient) GetTokenInfo(ctx context.Context, in *GetTokenInfoReq, opts ...grpc.CallOption) (*GetTokenInfoReply, error) {
	out := new(GetTokenInfoReply)
	err := c.cc.Invoke(ctx, "/passport.service.identify.v1.Identify/GetTokenInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentifyServer is the server API for Identify service.
type IdentifyServer interface {
	// TokenInfo identify info by token.
	GetTokenInfo(context.Context, *GetTokenInfoReq) (*GetTokenInfoReply, error)
}

// UnimplementedIdentifyServer can be embedded to have forward compatible implementations.
type UnimplementedIdentifyServer struct {
}

func (*UnimplementedIdentifyServer) GetTokenInfo(ctx context.Context, req *GetTokenInfoReq) (*GetTokenInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenInfo not implemented")
}

func RegisterIdentifyServer(s *grpc.Server, srv IdentifyServer) {
	s.RegisterService(&_Identify_serviceDesc, srv)
}

func _Identify_GetTokenInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentifyServer).GetTokenInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passport.service.identify.v1.Identify/GetTokenInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentifyServer).GetTokenInfo(ctx, req.(*GetTokenInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Identify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "passport.service.identify.v1.Identify",
	HandlerType: (*IdentifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTokenInfo",
			Handler:    _Identify_GetTokenInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *GetTokenInfoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTokenInfoReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTokenInfoReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetTokenInfoReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTokenInfoReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTokenInfoReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Expires != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Expires))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Csrf) > 0 {
		i -= len(m.Csrf)
		copy(dAtA[i:], m.Csrf)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Csrf)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Mid != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Mid))
		i--
		dAtA[i] = 0x10
	}
	if m.IsLogin {
		i--
		if m.IsLogin {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetTokenInfoReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetTokenInfoReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsLogin {
		n += 2
	}
	if m.Mid != 0 {
		n += 1 + sovApi(uint64(m.Mid))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Csrf)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Expires != 0 {
		n += 1 + sovApi(uint64(m.Expires))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetTokenInfoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: GetTokenInfoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTokenInfoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
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
func (m *GetTokenInfoReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: GetTokenInfoReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTokenInfoReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsLogin", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsLogin = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
			}
			m.Mid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mid |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Csrf", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Csrf = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			m.Expires = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expires |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
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
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
				return 0, ErrInvalidLengthApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApi = fmt.Errorf("proto: unexpected end of group")
)
