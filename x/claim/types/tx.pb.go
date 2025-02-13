// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/claim/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgInitialClaim struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgInitialClaim) Reset()         { *m = MsgInitialClaim{} }
func (m *MsgInitialClaim) String() string { return proto.CompactTextString(m) }
func (*MsgInitialClaim) ProtoMessage()    {}
func (*MsgInitialClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ee4a19153cf6635, []int{0}
}
func (m *MsgInitialClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitialClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitialClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitialClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitialClaim.Merge(m, src)
}
func (m *MsgInitialClaim) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitialClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitialClaim.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitialClaim proto.InternalMessageInfo

func (m *MsgInitialClaim) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

type MsgInitialClaimResponse struct {
	// total initial claimable amount for the user
	ClaimedAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=claimed_amount,json=claimedAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"claimed_amount" yaml:"claimed_amount"`
}

func (m *MsgInitialClaimResponse) Reset()         { *m = MsgInitialClaimResponse{} }
func (m *MsgInitialClaimResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInitialClaimResponse) ProtoMessage()    {}
func (*MsgInitialClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ee4a19153cf6635, []int{1}
}
func (m *MsgInitialClaimResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitialClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitialClaimResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitialClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitialClaimResponse.Merge(m, src)
}
func (m *MsgInitialClaimResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitialClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitialClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitialClaimResponse proto.InternalMessageInfo

func (m *MsgInitialClaimResponse) GetClaimedAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ClaimedAmount
	}
	return nil
}

type MsgClaimFor struct {
	Sender  string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Action  Action `protobuf:"varint,3,opt,name=action,proto3,enum=publicawesome.stargaze.claim.v1beta1.Action" json:"action,omitempty"`
}

func (m *MsgClaimFor) Reset()         { *m = MsgClaimFor{} }
func (m *MsgClaimFor) String() string { return proto.CompactTextString(m) }
func (*MsgClaimFor) ProtoMessage()    {}
func (*MsgClaimFor) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ee4a19153cf6635, []int{2}
}
func (m *MsgClaimFor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimFor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimFor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimFor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimFor.Merge(m, src)
}
func (m *MsgClaimFor) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimFor) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimFor.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimFor proto.InternalMessageInfo

func (m *MsgClaimFor) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgClaimFor) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgClaimFor) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return ActionInitialClaim
}

type MsgClaimForResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// total initial claimable amount for the user
	ClaimedAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=claimed_amount,json=claimedAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"claimed_amount" yaml:"claimed_amount"`
}

func (m *MsgClaimForResponse) Reset()         { *m = MsgClaimForResponse{} }
func (m *MsgClaimForResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClaimForResponse) ProtoMessage()    {}
func (*MsgClaimForResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ee4a19153cf6635, []int{3}
}
func (m *MsgClaimForResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimForResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimForResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimForResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimForResponse.Merge(m, src)
}
func (m *MsgClaimForResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimForResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimForResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimForResponse proto.InternalMessageInfo

func (m *MsgClaimForResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgClaimForResponse) GetClaimedAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ClaimedAmount
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgInitialClaim)(nil), "publicawesome.stargaze.claim.v1beta1.MsgInitialClaim")
	proto.RegisterType((*MsgInitialClaimResponse)(nil), "publicawesome.stargaze.claim.v1beta1.MsgInitialClaimResponse")
	proto.RegisterType((*MsgClaimFor)(nil), "publicawesome.stargaze.claim.v1beta1.MsgClaimFor")
	proto.RegisterType((*MsgClaimForResponse)(nil), "publicawesome.stargaze.claim.v1beta1.MsgClaimForResponse")
}

func init() { proto.RegisterFile("stargaze/claim/v1beta1/tx.proto", fileDescriptor_9ee4a19153cf6635) }

var fileDescriptor_9ee4a19153cf6635 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0xbf, 0x6e, 0xd4, 0x40,
	0x10, 0xc6, 0x6f, 0xef, 0xa4, 0x03, 0x36, 0x10, 0x24, 0xf3, 0xcf, 0xb8, 0xb0, 0x4f, 0x16, 0x85,
	0x23, 0x91, 0x5d, 0xdd, 0x05, 0x0a, 0x90, 0x28, 0x92, 0x20, 0xa4, 0x14, 0x6e, 0x5c, 0xd2, 0x44,
	0x6b, 0x7b, 0x65, 0x56, 0xd8, 0x5e, 0xcb, 0xb3, 0x17, 0x2e, 0xd4, 0x40, 0x43, 0xc3, 0x5b, 0x20,
	0xf1, 0x0e, 0xf4, 0x29, 0x53, 0x52, 0x05, 0x74, 0xf7, 0x06, 0x3c, 0x01, 0xba, 0xf5, 0xda, 0xba,
	0x8b, 0x14, 0xe9, 0xa0, 0xa2, 0xb2, 0x57, 0x33, 0xbf, 0x6f, 0x66, 0xbe, 0xdd, 0xc1, 0x1e, 0x28,
	0x56, 0x67, 0xec, 0x3d, 0xa7, 0x49, 0xce, 0x44, 0x41, 0x4f, 0xc6, 0x31, 0x57, 0x6c, 0x4c, 0xd5,
	0x8c, 0x54, 0xb5, 0x54, 0xd2, 0x7a, 0x54, 0x4d, 0xe3, 0x5c, 0x24, 0xec, 0x1d, 0x07, 0x59, 0x70,
	0xd2, 0xa6, 0x13, 0x9d, 0x4e, 0x4c, 0xba, 0x73, 0x37, 0x93, 0x99, 0xd4, 0x00, 0x5d, 0xfe, 0x35,
	0xac, 0xe3, 0x26, 0x12, 0x0a, 0x09, 0x34, 0x66, 0xc0, 0x3b, 0xe5, 0x44, 0x8a, 0xd2, 0xc4, 0x77,
	0xae, 0x28, 0xae, 0x4f, 0xc7, 0x35, 0x4f, 0x64, 0x9d, 0x36, 0xa9, 0xfe, 0x0e, 0xbe, 0x1d, 0x42,
	0x76, 0x54, 0x0a, 0x25, 0x58, 0x7e, 0xb8, 0x8c, 0x5b, 0xf7, 0xf1, 0x10, 0x78, 0x99, 0xf2, 0xda,
	0x46, 0x23, 0x14, 0xdc, 0x88, 0xcc, 0xc9, 0xff, 0x8a, 0xf0, 0x83, 0x4b, 0xb9, 0x11, 0x87, 0x4a,
	0x96, 0xc0, 0xad, 0xcf, 0x08, 0x6f, 0x6b, 0x75, 0x9e, 0x1e, 0xb3, 0x42, 0x4e, 0x4b, 0x65, 0xf7,
	0x47, 0x83, 0x60, 0x6b, 0xf2, 0x90, 0x34, 0xbd, 0x92, 0x65, 0xaf, 0xed, 0x58, 0xe4, 0x50, 0x8a,
	0xf2, 0xe0, 0xe8, 0xec, 0xc2, 0xeb, 0xfd, 0xbe, 0xf0, 0xee, 0x9d, 0xb2, 0x22, 0x7f, 0xee, 0xaf,
	0xe3, 0xfe, 0xb7, 0x9f, 0x5e, 0x90, 0x09, 0xf5, 0x66, 0x1a, 0x93, 0x44, 0x16, 0xd4, 0x4c, 0xdc,
	0x7c, 0x76, 0x21, 0x7d, 0x4b, 0xd5, 0x69, 0xc5, 0x41, 0x2b, 0x41, 0x74, 0xcb, 0xc0, 0xfb, 0x0d,
	0xfb, 0x11, 0xe1, 0xad, 0x10, 0x32, 0xdd, 0xe2, 0x2b, 0x59, 0x5f, 0x35, 0x91, 0x65, 0xe3, 0x6b,
	0x2c, 0x4d, 0x6b, 0x0e, 0x60, 0xf7, 0x75, 0xa0, 0x3d, 0x5a, 0x2f, 0xf1, 0x90, 0x25, 0x4a, 0xc8,
	0xd2, 0x1e, 0x8c, 0x50, 0xb0, 0x3d, 0x79, 0x4c, 0x36, 0xb9, 0x2e, 0xb2, 0xaf, 0x99, 0xc8, 0xb0,
	0xfe, 0x77, 0x84, 0xef, 0xac, 0xf4, 0xd1, 0xb9, 0xb5, 0x52, 0x17, 0xad, 0xd7, 0xfd, 0xaf, 0x7c,
	0x9c, 0x7c, 0xea, 0xe3, 0x41, 0x08, 0x99, 0xf5, 0x01, 0xe1, 0x9b, 0x6b, 0x4f, 0xe4, 0xe9, 0x66,
	0x76, 0x5c, 0x7a, 0x2d, 0xce, 0x8b, 0x7f, 0xc2, 0x3a, 0xdb, 0x66, 0xf8, 0x7a, 0x77, 0xa5, 0xe3,
	0x8d, 0xa5, 0x5a, 0xc4, 0x79, 0xf6, 0xd7, 0x48, 0x5b, 0xf9, 0x20, 0x3c, 0x9b, 0xbb, 0xe8, 0x7c,
	0xee, 0xa2, 0x5f, 0x73, 0x17, 0x7d, 0x59, 0xb8, 0xbd, 0xf3, 0x85, 0xdb, 0xfb, 0xb1, 0x70, 0x7b,
	0xaf, 0xf7, 0x56, 0xbc, 0x6d, 0xe4, 0x77, 0x8d, 0x3e, 0xed, 0x96, 0xf0, 0xe4, 0x09, 0x9d, 0x99,
	0x4d, 0xd4, 0x66, 0xc7, 0x43, 0xbd, 0x7b, 0x7b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x7b,
	0xa3, 0xcd, 0x25, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	InitialClaim(ctx context.Context, in *MsgInitialClaim, opts ...grpc.CallOption) (*MsgInitialClaimResponse, error)
	// this line is used by starport scaffolding # proto/tx/rpc
	ClaimFor(ctx context.Context, in *MsgClaimFor, opts ...grpc.CallOption) (*MsgClaimForResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) InitialClaim(ctx context.Context, in *MsgInitialClaim, opts ...grpc.CallOption) (*MsgInitialClaimResponse, error) {
	out := new(MsgInitialClaimResponse)
	err := c.cc.Invoke(ctx, "/publicawesome.stargaze.claim.v1beta1.Msg/InitialClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClaimFor(ctx context.Context, in *MsgClaimFor, opts ...grpc.CallOption) (*MsgClaimForResponse, error) {
	out := new(MsgClaimForResponse)
	err := c.cc.Invoke(ctx, "/publicawesome.stargaze.claim.v1beta1.Msg/ClaimFor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	InitialClaim(context.Context, *MsgInitialClaim) (*MsgInitialClaimResponse, error)
	// this line is used by starport scaffolding # proto/tx/rpc
	ClaimFor(context.Context, *MsgClaimFor) (*MsgClaimForResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) InitialClaim(ctx context.Context, req *MsgInitialClaim) (*MsgInitialClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitialClaim not implemented")
}
func (*UnimplementedMsgServer) ClaimFor(ctx context.Context, req *MsgClaimFor) (*MsgClaimForResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimFor not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_InitialClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInitialClaim)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InitialClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicawesome.stargaze.claim.v1beta1.Msg/InitialClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InitialClaim(ctx, req.(*MsgInitialClaim))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClaimFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimFor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicawesome.stargaze.claim.v1beta1.Msg/ClaimFor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimFor(ctx, req.(*MsgClaimFor))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "publicawesome.stargaze.claim.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitialClaim",
			Handler:    _Msg_InitialClaim_Handler,
		},
		{
			MethodName: "ClaimFor",
			Handler:    _Msg_ClaimFor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stargaze/claim/v1beta1/tx.proto",
}

func (m *MsgInitialClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitialClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitialClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInitialClaimResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitialClaimResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitialClaimResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimedAmount) > 0 {
		for iNdEx := len(m.ClaimedAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimedAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimFor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimFor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimFor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Action != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Action))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimForResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimForResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimForResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimedAmount) > 0 {
		for iNdEx := len(m.ClaimedAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimedAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgInitialClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgInitialClaimResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ClaimedAmount) > 0 {
		for _, e := range m.ClaimedAmount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgClaimFor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Action != 0 {
		n += 1 + sovTx(uint64(m.Action))
	}
	return n
}

func (m *MsgClaimForResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.ClaimedAmount) > 0 {
		for _, e := range m.ClaimedAmount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInitialClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInitialClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitialClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgInitialClaimResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInitialClaimResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitialClaimResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimedAmount = append(m.ClaimedAmount, types.Coin{})
			if err := m.ClaimedAmount[len(m.ClaimedAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgClaimFor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClaimFor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimFor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= Action(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgClaimForResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgClaimForResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimForResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimedAmount = append(m.ClaimedAmount, types.Coin{})
			if err := m.ClaimedAmount[len(m.ClaimedAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
