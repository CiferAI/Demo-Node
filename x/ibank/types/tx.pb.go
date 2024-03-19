// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibank/tx.proto

package types

import (
	context "context"
	fmt "fmt"
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

type MsgSend struct {
	FromAddress  string     `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ToAddress    string     `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount       types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
	PasswordHash string     `protobuf:"bytes,4,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
}

func (m *MsgSend) Reset()         { *m = MsgSend{} }
func (m *MsgSend) String() string { return proto.CompactTextString(m) }
func (*MsgSend) ProtoMessage()    {}
func (*MsgSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73d3d253c1616b9, []int{0}
}
func (m *MsgSend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSend.Merge(m, src)
}
func (m *MsgSend) XXX_Size() int {
	return m.Size()
}
func (m *MsgSend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSend proto.InternalMessageInfo

func (m *MsgSend) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgSend) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *MsgSend) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *MsgSend) GetPasswordHash() string {
	if m != nil {
		return m.PasswordHash
	}
	return ""
}

type MsgSendResponse struct {
}

func (m *MsgSendResponse) Reset()         { *m = MsgSendResponse{} }
func (m *MsgSendResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendResponse) ProtoMessage()    {}
func (*MsgSendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73d3d253c1616b9, []int{1}
}
func (m *MsgSendResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendResponse.Merge(m, src)
}
func (m *MsgSendResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendResponse proto.InternalMessageInfo

type MsgReceive struct {
	Receiver      string `protobuf:"bytes,1,opt,name=receiver,proto3" json:"receiver,omitempty"`
	TransactionId int64  `protobuf:"varint,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Password      string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *MsgReceive) Reset()         { *m = MsgReceive{} }
func (m *MsgReceive) String() string { return proto.CompactTextString(m) }
func (*MsgReceive) ProtoMessage()    {}
func (*MsgReceive) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73d3d253c1616b9, []int{2}
}
func (m *MsgReceive) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgReceive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgReceive.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgReceive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReceive.Merge(m, src)
}
func (m *MsgReceive) XXX_Size() int {
	return m.Size()
}
func (m *MsgReceive) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReceive.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReceive proto.InternalMessageInfo

func (m *MsgReceive) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgReceive) GetTransactionId() int64 {
	if m != nil {
		return m.TransactionId
	}
	return 0
}

func (m *MsgReceive) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type MsgReceiveResponse struct {
}

func (m *MsgReceiveResponse) Reset()         { *m = MsgReceiveResponse{} }
func (m *MsgReceiveResponse) String() string { return proto.CompactTextString(m) }
func (*MsgReceiveResponse) ProtoMessage()    {}
func (*MsgReceiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73d3d253c1616b9, []int{3}
}
func (m *MsgReceiveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgReceiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgReceiveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgReceiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReceiveResponse.Merge(m, src)
}
func (m *MsgReceiveResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgReceiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReceiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReceiveResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSend)(nil), "mun.ibank.MsgSend")
	proto.RegisterType((*MsgSendResponse)(nil), "mun.ibank.MsgSendResponse")
	proto.RegisterType((*MsgReceive)(nil), "mun.ibank.MsgReceive")
	proto.RegisterType((*MsgReceiveResponse)(nil), "mun.ibank.MsgReceiveResponse")
}

func init() { proto.RegisterFile("ibank/tx.proto", fileDescriptor_f73d3d253c1616b9) }

var fileDescriptor_f73d3d253c1616b9 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xbd, 0x6e, 0xd4, 0x40,
	0x10, 0xc7, 0x6d, 0xee, 0x94, 0xe0, 0xc9, 0x07, 0x62, 0x15, 0xa4, 0xc3, 0x52, 0x4c, 0x30, 0x42,
	0x4a, 0xc3, 0xae, 0x12, 0x24, 0x28, 0x11, 0xa1, 0x21, 0xc5, 0x35, 0xa6, 0xa3, 0x39, 0xad, 0xed,
	0xc5, 0x5e, 0x1d, 0xde, 0xb1, 0x3c, 0xeb, 0x10, 0x1a, 0x9e, 0x81, 0xd7, 0xe0, 0x4d, 0x52, 0xa6,
	0xa4, 0x42, 0xe8, 0xee, 0x45, 0x90, 0xd7, 0x1f, 0x80, 0xa0, 0x9b, 0xf9, 0xf9, 0xef, 0xd9, 0x9f,
	0xbd, 0x03, 0x87, 0x3a, 0x95, 0x66, 0x2d, 0xec, 0x35, 0xaf, 0x1b, 0xb4, 0xc8, 0x82, 0xaa, 0x35,
	0xdc, 0xb1, 0x30, 0xca, 0x90, 0x2a, 0x24, 0x91, 0x4a, 0x52, 0xe2, 0xea, 0x2c, 0x55, 0x56, 0x9e,
	0x89, 0x0c, 0xb5, 0xe9, 0xa3, 0xe1, 0x51, 0x81, 0x05, 0xba, 0x52, 0x74, 0x55, 0x4f, 0xe3, 0x6f,
	0x3e, 0xec, 0x2e, 0xa9, 0x78, 0xa7, 0x4c, 0xce, 0x1e, 0xc3, 0xfe, 0x87, 0x06, 0xab, 0x95, 0xcc,
	0xf3, 0x46, 0x11, 0x2d, 0xfc, 0x13, 0xff, 0x34, 0x48, 0xf6, 0x3a, 0xf6, 0xba, 0x47, 0xec, 0x18,
	0xc0, 0xe2, 0x14, 0xb8, 0xe3, 0x02, 0x81, 0xc5, 0xf1, 0xf1, 0x4b, 0xd8, 0x91, 0x15, 0xb6, 0xc6,
	0x2e, 0x66, 0x27, 0xfe, 0xe9, 0xde, 0xf9, 0x43, 0xde, 0x4b, 0xf1, 0x4e, 0x8a, 0x0f, 0x52, 0xfc,
	0x0d, 0x6a, 0x73, 0x31, 0xbf, 0xf9, 0xf1, 0xc8, 0x4b, 0x86, 0x38, 0x7b, 0x02, 0x07, 0xb5, 0x24,
	0xfa, 0x84, 0x4d, 0xbe, 0x2a, 0x25, 0x95, 0x8b, 0xb9, 0x1b, 0xbd, 0x3f, 0xc2, 0xb7, 0x92, 0xca,
	0xf8, 0x3e, 0xdc, 0x1b, 0x54, 0x13, 0x45, 0x35, 0x1a, 0x52, 0xf1, 0x1a, 0x60, 0x49, 0x45, 0xa2,
	0x32, 0xa5, 0xaf, 0x14, 0x0b, 0xe1, 0x6e, 0xd3, 0x97, 0xcd, 0x20, 0x3f, 0xf5, 0xec, 0x29, 0x1c,
	0xda, 0x46, 0x1a, 0x92, 0x99, 0xd5, 0x68, 0x56, 0x3a, 0x77, 0xf6, 0xb3, 0xe4, 0xe0, 0x0f, 0x7a,
	0x99, 0x77, 0x23, 0xc6, 0x33, 0xdd, 0x37, 0x04, 0xc9, 0xd4, 0xc7, 0x47, 0xc0, 0x7e, 0x1f, 0x36,
	0x2a, 0x9c, 0x7f, 0x81, 0xd9, 0x92, 0x0a, 0xf6, 0x02, 0xe6, 0xee, 0x27, 0x32, 0x3e, 0x5d, 0x09,
	0x1f, 0x6c, 0xc3, 0xf0, 0x5f, 0x36, 0xbe, 0xce, 0x5e, 0xc1, 0xee, 0xa8, 0xff, 0xe0, 0xef, 0xd8,
	0x80, 0xc3, 0xe3, 0xff, 0xe2, 0x71, 0xc0, 0xc5, 0xe5, 0xcd, 0x26, 0xf2, 0x6f, 0x37, 0x91, 0xff,
	0x73, 0x13, 0xf9, 0x5f, 0xb7, 0x91, 0x77, 0xbb, 0x8d, 0xbc, 0xef, 0xdb, 0xc8, 0x7b, 0x2f, 0x0a,
	0x6d, 0xcb, 0x36, 0xe5, 0x19, 0x56, 0xa2, 0x6a, 0x4d, 0xfa, 0x11, 0xb3, 0x75, 0x56, 0x4a, 0x6d,
	0xba, 0xee, 0x99, 0xc1, 0x5c, 0x89, 0x6b, 0x31, 0x2c, 0xd4, 0xe7, 0x5a, 0x51, 0xba, 0xe3, 0x76,
	0xe2, 0xf9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x31, 0x62, 0x3a, 0x66, 0x02, 0x00, 0x00,
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
	Send(ctx context.Context, in *MsgSend, opts ...grpc.CallOption) (*MsgSendResponse, error)
	Receive(ctx context.Context, in *MsgReceive, opts ...grpc.CallOption) (*MsgReceiveResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Send(ctx context.Context, in *MsgSend, opts ...grpc.CallOption) (*MsgSendResponse, error) {
	out := new(MsgSendResponse)
	err := c.cc.Invoke(ctx, "/mun.ibank.Msg/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Receive(ctx context.Context, in *MsgReceive, opts ...grpc.CallOption) (*MsgReceiveResponse, error) {
	out := new(MsgReceiveResponse)
	err := c.cc.Invoke(ctx, "/mun.ibank.Msg/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Send(context.Context, *MsgSend) (*MsgSendResponse, error)
	Receive(context.Context, *MsgReceive) (*MsgReceiveResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Send(ctx context.Context, req *MsgSend) (*MsgSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedMsgServer) Receive(ctx context.Context, req *MsgReceive) (*MsgReceiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mun.ibank.Msg/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Send(ctx, req.(*MsgSend))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgReceive)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mun.ibank.Msg/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Receive(ctx, req.(*MsgReceive))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mun.ibank.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Msg_Send_Handler,
		},
		{
			MethodName: "Receive",
			Handler:    _Msg_Receive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibank/tx.proto",
}

func (m *MsgSend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PasswordHash) > 0 {
		i -= len(m.PasswordHash)
		copy(dAtA[i:], m.PasswordHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PasswordHash)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgReceive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgReceive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgReceive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TransactionId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TransactionId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgReceiveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgReceiveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgReceiveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgSend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.PasswordHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSendResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgReceive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TransactionId != 0 {
		n += 1 + sovTx(uint64(m.TransactionId))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgReceiveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSend) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
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
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
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
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PasswordHash", wireType)
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
			m.PasswordHash = string(dAtA[iNdEx:postIndex])
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
func (m *MsgSendResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSendResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *MsgReceive) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgReceive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgReceive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
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
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionId", wireType)
			}
			m.TransactionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransactionId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
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
			m.Password = string(dAtA[iNdEx:postIndex])
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
func (m *MsgReceiveResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgReceiveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgReceiveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
