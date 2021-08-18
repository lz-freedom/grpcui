// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.tl.handshake_service.proto

package mtproto

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

// /////////////////////////////////////////////////////////////////////////////
// req_pq#60469778 nonce:int128 = ResPQ;
type TLReqPq struct {
	Nonce                []byte   `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLReqPq) Reset()         { *m = TLReqPq{} }
func (m *TLReqPq) String() string { return proto.CompactTextString(m) }
func (*TLReqPq) ProtoMessage()    {}
func (*TLReqPq) Descriptor() ([]byte, []int) {
	return fileDescriptor_schema_tl_handshake_service_6871ce8821faa38e, []int{0}
}
func (m *TLReqPq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLReqPq.Unmarshal(m, b)
}
func (m *TLReqPq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLReqPq.Marshal(b, m, deterministic)
}
func (dst *TLReqPq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLReqPq.Merge(dst, src)
}
func (m *TLReqPq) XXX_Size() int {
	return xxx_messageInfo_TLReqPq.Size(m)
}
func (m *TLReqPq) XXX_DiscardUnknown() {
	xxx_messageInfo_TLReqPq.DiscardUnknown(m)
}

var xxx_messageInfo_TLReqPq proto.InternalMessageInfo

func (m *TLReqPq) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// req_pq_multi#be7e8ef1 nonce:int128 = ResPQ;
type TLReqPqMulti struct {
	Nonce                []byte   `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLReqPqMulti) Reset()         { *m = TLReqPqMulti{} }
func (m *TLReqPqMulti) String() string { return proto.CompactTextString(m) }
func (*TLReqPqMulti) ProtoMessage()    {}
func (*TLReqPqMulti) Descriptor() ([]byte, []int) {
	return fileDescriptor_schema_tl_handshake_service_6871ce8821faa38e, []int{1}
}
func (m *TLReqPqMulti) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLReqPqMulti.Unmarshal(m, b)
}
func (m *TLReqPqMulti) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLReqPqMulti.Marshal(b, m, deterministic)
}
func (dst *TLReqPqMulti) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLReqPqMulti.Merge(dst, src)
}
func (m *TLReqPqMulti) XXX_Size() int {
	return xxx_messageInfo_TLReqPqMulti.Size(m)
}
func (m *TLReqPqMulti) XXX_DiscardUnknown() {
	xxx_messageInfo_TLReqPqMulti.DiscardUnknown(m)
}

var xxx_messageInfo_TLReqPqMulti proto.InternalMessageInfo

func (m *TLReqPqMulti) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// req_DH_params#d712e4be nonce:int128 server_nonce:int128 p:string q:string public_key_fingerprint:long encrypted_data:string = Server_DH_Params;
type TLReq_DHParams struct {
	Nonce                []byte   `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce          []byte   `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	P                    string   `protobuf:"bytes,3,opt,name=p,proto3" json:"p,omitempty"`
	Q                    string   `protobuf:"bytes,4,opt,name=q,proto3" json:"q,omitempty"`
	PublicKeyFingerprint int64    `protobuf:"varint,5,opt,name=public_key_fingerprint,json=publicKeyFingerprint,proto3" json:"public_key_fingerprint,omitempty"`
	EncryptedData        string   `protobuf:"bytes,6,opt,name=encrypted_data,json=encryptedData,proto3" json:"encrypted_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLReq_DHParams) Reset()         { *m = TLReq_DHParams{} }
func (m *TLReq_DHParams) String() string { return proto.CompactTextString(m) }
func (*TLReq_DHParams) ProtoMessage()    {}
func (*TLReq_DHParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_schema_tl_handshake_service_6871ce8821faa38e, []int{2}
}
func (m *TLReq_DHParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLReq_DHParams.Unmarshal(m, b)
}
func (m *TLReq_DHParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLReq_DHParams.Marshal(b, m, deterministic)
}
func (dst *TLReq_DHParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLReq_DHParams.Merge(dst, src)
}
func (m *TLReq_DHParams) XXX_Size() int {
	return xxx_messageInfo_TLReq_DHParams.Size(m)
}
func (m *TLReq_DHParams) XXX_DiscardUnknown() {
	xxx_messageInfo_TLReq_DHParams.DiscardUnknown(m)
}

var xxx_messageInfo_TLReq_DHParams proto.InternalMessageInfo

func (m *TLReq_DHParams) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *TLReq_DHParams) GetServerNonce() []byte {
	if m != nil {
		return m.ServerNonce
	}
	return nil
}

func (m *TLReq_DHParams) GetP() string {
	if m != nil {
		return m.P
	}
	return ""
}

func (m *TLReq_DHParams) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *TLReq_DHParams) GetPublicKeyFingerprint() int64 {
	if m != nil {
		return m.PublicKeyFingerprint
	}
	return 0
}

func (m *TLReq_DHParams) GetEncryptedData() string {
	if m != nil {
		return m.EncryptedData
	}
	return ""
}

// /////////////////////////////////////////////////////////////////////////////
// set_client_DH_params#f5045f1f nonce:int128 server_nonce:int128 encrypted_data:string = Set_client_DH_params_answer;
type TLSetClient_DHParams struct {
	Nonce                []byte   `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce          []byte   `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	EncryptedData        string   `protobuf:"bytes,3,opt,name=encrypted_data,json=encryptedData,proto3" json:"encrypted_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLSetClient_DHParams) Reset()         { *m = TLSetClient_DHParams{} }
func (m *TLSetClient_DHParams) String() string { return proto.CompactTextString(m) }
func (*TLSetClient_DHParams) ProtoMessage()    {}
func (*TLSetClient_DHParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_schema_tl_handshake_service_6871ce8821faa38e, []int{3}
}
func (m *TLSetClient_DHParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLSetClient_DHParams.Unmarshal(m, b)
}
func (m *TLSetClient_DHParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLSetClient_DHParams.Marshal(b, m, deterministic)
}
func (dst *TLSetClient_DHParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSetClient_DHParams.Merge(dst, src)
}
func (m *TLSetClient_DHParams) XXX_Size() int {
	return xxx_messageInfo_TLSetClient_DHParams.Size(m)
}
func (m *TLSetClient_DHParams) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSetClient_DHParams.DiscardUnknown(m)
}

var xxx_messageInfo_TLSetClient_DHParams proto.InternalMessageInfo

func (m *TLSetClient_DHParams) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *TLSetClient_DHParams) GetServerNonce() []byte {
	if m != nil {
		return m.ServerNonce
	}
	return nil
}

func (m *TLSetClient_DHParams) GetEncryptedData() string {
	if m != nil {
		return m.EncryptedData
	}
	return ""
}

// /////////////////////////////////////////////////////////////////////////////
// destroy_auth_key#d1435160 = DestroyAuthKeyRes;
type TLDestroyAuthKey struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLDestroyAuthKey) Reset()         { *m = TLDestroyAuthKey{} }
func (m *TLDestroyAuthKey) String() string { return proto.CompactTextString(m) }
func (*TLDestroyAuthKey) ProtoMessage()    {}
func (*TLDestroyAuthKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_schema_tl_handshake_service_6871ce8821faa38e, []int{4}
}
func (m *TLDestroyAuthKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLDestroyAuthKey.Unmarshal(m, b)
}
func (m *TLDestroyAuthKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLDestroyAuthKey.Marshal(b, m, deterministic)
}
func (dst *TLDestroyAuthKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLDestroyAuthKey.Merge(dst, src)
}
func (m *TLDestroyAuthKey) XXX_Size() int {
	return xxx_messageInfo_TLDestroyAuthKey.Size(m)
}
func (m *TLDestroyAuthKey) XXX_DiscardUnknown() {
	xxx_messageInfo_TLDestroyAuthKey.DiscardUnknown(m)
}

var xxx_messageInfo_TLDestroyAuthKey proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TLReqPq)(nil), "mtproto.TL_req_pq")
	proto.RegisterType((*TLReqPqMulti)(nil), "mtproto.TL_req_pq_multi")
	proto.RegisterType((*TLReq_DHParams)(nil), "mtproto.TL_req_DH_params")
	proto.RegisterType((*TLSetClient_DHParams)(nil), "mtproto.TL_set_client_DH_params")
	proto.RegisterType((*TLDestroyAuthKey)(nil), "mtproto.TL_destroy_auth_key")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCAuthKeyClient is the client API for RPCAuthKey service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCAuthKeyClient interface {
	// req_pq#60469778 nonce:int128 = ResPQ;
	ReqPq(ctx context.Context, in *TLReqPq, opts ...grpc.CallOption) (*ResPQ, error)
	// req_pq_multi#be7e8ef1 nonce:int128 = ResPQ;
	ReqPqMulti(ctx context.Context, in *TLReqPqMulti, opts ...grpc.CallOption) (*ResPQ, error)
	// req_DH_params#d712e4be nonce:int128 server_nonce:int128 p:string q:string public_key_fingerprint:long encrypted_data:string = Server_DH_Params;
	Req_DHParams(ctx context.Context, in *TLReq_DHParams, opts ...grpc.CallOption) (*Server_DH_Params, error)
	// set_client_DH_params#f5045f1f nonce:int128 server_nonce:int128 encrypted_data:string = Set_client_DH_params_answer;
	SetClient_DHParams(ctx context.Context, in *TLSetClient_DHParams, opts ...grpc.CallOption) (*SetClient_DHParamsAnswer, error)
	// destroy_auth_key#d1435160 = DestroyAuthKeyRes;
	DestroyAuthKey(ctx context.Context, in *TLDestroyAuthKey, opts ...grpc.CallOption) (*DestroyAuthKeyRes, error)
}

type rPCAuthKeyClient struct {
	cc *grpc.ClientConn
}

func NewRPCAuthKeyClient(cc *grpc.ClientConn) RPCAuthKeyClient {
	return &rPCAuthKeyClient{cc}
}

func (c *rPCAuthKeyClient) ReqPq(ctx context.Context, in *TLReqPq, opts ...grpc.CallOption) (*ResPQ, error) {
	out := new(ResPQ)
	err := c.cc.Invoke(ctx, "/mtproto.RPCAuthKey/req_pq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthKeyClient) ReqPqMulti(ctx context.Context, in *TLReqPqMulti, opts ...grpc.CallOption) (*ResPQ, error) {
	out := new(ResPQ)
	err := c.cc.Invoke(ctx, "/mtproto.RPCAuthKey/req_pq_multi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthKeyClient) Req_DHParams(ctx context.Context, in *TLReq_DHParams, opts ...grpc.CallOption) (*Server_DH_Params, error) {
	out := new(Server_DH_Params)
	err := c.cc.Invoke(ctx, "/mtproto.RPCAuthKey/req_DH_params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthKeyClient) SetClient_DHParams(ctx context.Context, in *TLSetClient_DHParams, opts ...grpc.CallOption) (*SetClient_DHParamsAnswer, error) {
	out := new(SetClient_DHParamsAnswer)
	err := c.cc.Invoke(ctx, "/mtproto.RPCAuthKey/set_client_DH_params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthKeyClient) DestroyAuthKey(ctx context.Context, in *TLDestroyAuthKey, opts ...grpc.CallOption) (*DestroyAuthKeyRes, error) {
	out := new(DestroyAuthKeyRes)
	err := c.cc.Invoke(ctx, "/mtproto.RPCAuthKey/destroy_auth_key", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCAuthKeyServer is the server API for RPCAuthKey service.
type RPCAuthKeyServer interface {
	// req_pq#60469778 nonce:int128 = ResPQ;
	ReqPq(context.Context, *TLReqPq) (*ResPQ, error)
	// req_pq_multi#be7e8ef1 nonce:int128 = ResPQ;
	ReqPqMulti(context.Context, *TLReqPqMulti) (*ResPQ, error)
	// req_DH_params#d712e4be nonce:int128 server_nonce:int128 p:string q:string public_key_fingerprint:long encrypted_data:string = Server_DH_Params;
	Req_DHParams(context.Context, *TLReq_DHParams) (*Server_DH_Params, error)
	// set_client_DH_params#f5045f1f nonce:int128 server_nonce:int128 encrypted_data:string = Set_client_DH_params_answer;
	SetClient_DHParams(context.Context, *TLSetClient_DHParams) (*SetClient_DHParamsAnswer, error)
	// destroy_auth_key#d1435160 = DestroyAuthKeyRes;
	DestroyAuthKey(context.Context, *TLDestroyAuthKey) (*DestroyAuthKeyRes, error)
}

func RegisterRPCAuthKeyServer(s *grpc.Server, srv RPCAuthKeyServer) {
	s.RegisterService(&_RPCAuthKey_serviceDesc, srv)
}

func _RPCAuthKey_ReqPq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLReqPq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthKeyServer).ReqPq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCAuthKey/ReqPq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthKeyServer).ReqPq(ctx, req.(*TLReqPq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthKey_ReqPqMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLReqPqMulti)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthKeyServer).ReqPqMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCAuthKey/ReqPqMulti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthKeyServer).ReqPqMulti(ctx, req.(*TLReqPqMulti))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthKey_Req_DHParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLReq_DHParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthKeyServer).Req_DHParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCAuthKey/Req_DHParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthKeyServer).Req_DHParams(ctx, req.(*TLReq_DHParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthKey_SetClient_DHParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSetClient_DHParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthKeyServer).SetClient_DHParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCAuthKey/SetClient_DHParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthKeyServer).SetClient_DHParams(ctx, req.(*TLSetClient_DHParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthKey_DestroyAuthKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLDestroyAuthKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthKeyServer).DestroyAuthKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCAuthKey/DestroyAuthKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthKeyServer).DestroyAuthKey(ctx, req.(*TLDestroyAuthKey))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCAuthKey_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mtproto.RPCAuthKey",
	HandlerType: (*RPCAuthKeyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "req_pq",
			Handler:    _RPCAuthKey_ReqPq_Handler,
		},
		{
			MethodName: "req_pq_multi",
			Handler:    _RPCAuthKey_ReqPqMulti_Handler,
		},
		{
			MethodName: "req_DH_params",
			Handler:    _RPCAuthKey_Req_DHParams_Handler,
		},
		{
			MethodName: "set_client_DH_params",
			Handler:    _RPCAuthKey_SetClient_DHParams_Handler,
		},
		{
			MethodName: "destroy_auth_key",
			Handler:    _RPCAuthKey_DestroyAuthKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema.tl.handshake_service.proto",
}

func init() {
	proto.RegisterFile("schema.tl.handshake_service.proto", fileDescriptor_schema_tl_handshake_service_6871ce8821faa38e)
}

var fileDescriptor_schema_tl_handshake_service_6871ce8821faa38e = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe3, 0x84, 0xa6, 0xea, 0x90, 0x96, 0x6a, 0x09, 0xe0, 0x5a, 0x08, 0x25, 0x16, 0x88,
	0x48, 0x48, 0x7b, 0x00, 0x4e, 0xdc, 0x28, 0x11, 0x44, 0x6a, 0xa8, 0x82, 0xc9, 0x7d, 0xd9, 0x38,
	0x43, 0x6d, 0xd5, 0x59, 0xaf, 0x77, 0xd7, 0x20, 0xbf, 0x1f, 0x6f, 0xc4, 0x0b, 0x20, 0x7b, 0x8d,
	0x9b, 0x3f, 0xe6, 0xc6, 0x6d, 0x66, 0x7e, 0xdf, 0xcc, 0xee, 0xcc, 0x07, 0x63, 0x1d, 0x46, 0xb8,
	0xe1, 0xd4, 0x24, 0x34, 0xe2, 0x62, 0xad, 0x23, 0x7e, 0x8b, 0x4c, 0xa3, 0xfa, 0x11, 0x87, 0x48,
	0xa5, 0x4a, 0x4d, 0x4a, 0x8e, 0x37, 0xa6, 0x0a, 0xbc, 0x8b, 0x16, 0xad, 0xd5, 0xf8, 0x63, 0x38,
	0x59, 0xce, 0x99, 0xc2, 0x8c, 0xc9, 0x8c, 0x0c, 0xe1, 0x48, 0xa4, 0x22, 0x44, 0xd7, 0x19, 0x39,
	0x93, 0x41, 0x60, 0x13, 0xff, 0x25, 0x3c, 0x68, 0x24, 0x6c, 0x93, 0x27, 0x26, 0xfe, 0x87, 0xf0,
	0x97, 0x03, 0xe7, 0xb5, 0x72, 0x3a, 0x63, 0x92, 0x2b, 0xbe, 0xd1, 0xed, 0x52, 0x32, 0x86, 0x41,
	0xf9, 0x57, 0x54, 0xcc, 0xc2, 0x6e, 0x05, 0xef, 0xdb, 0xda, 0x75, 0x25, 0x19, 0x80, 0x23, 0xdd,
	0xde, 0xc8, 0x99, 0x9c, 0x04, 0x8e, 0x2c, 0xb3, 0xcc, 0xbd, 0x67, 0xb3, 0x8c, 0xbc, 0x85, 0xc7,
	0x32, 0x5f, 0x25, 0x71, 0xc8, 0x6e, 0xb1, 0x60, 0xdf, 0x63, 0x71, 0x83, 0x4a, 0xaa, 0x58, 0x18,
	0xf7, 0x68, 0xe4, 0x4c, 0x7a, 0xc1, 0xd0, 0xd2, 0x2b, 0x2c, 0x3e, 0xde, 0x31, 0xf2, 0x02, 0xce,
	0x50, 0x84, 0xaa, 0x90, 0x06, 0xd7, 0x6c, 0xcd, 0x0d, 0x77, 0xfb, 0xd5, 0xc0, 0xd3, 0xa6, 0x3a,
	0xe5, 0x86, 0xfb, 0x05, 0x3c, 0x59, 0xce, 0x99, 0x46, 0xc3, 0xc2, 0x24, 0x46, 0x61, 0xfe, 0xc7,
	0x32, 0x87, 0x4f, 0xf7, 0xda, 0x9e, 0x7e, 0x04, 0x0f, 0x97, 0x73, 0xb6, 0x46, 0x6d, 0x54, 0x5a,
	0x30, 0x9e, 0x9b, 0xa8, 0x5c, 0xf0, 0xf5, 0xef, 0x2e, 0x40, 0xb0, 0xf8, 0xf0, 0x3e, 0x37, 0xd1,
	0x15, 0x16, 0x84, 0x42, 0xbf, 0x36, 0x8c, 0xd0, 0xda, 0x62, 0xda, 0x38, 0xe4, 0x9d, 0x35, 0xb5,
	0x00, 0xf5, 0xe2, 0x8b, 0xdf, 0x21, 0xef, 0x60, 0xb0, 0xe3, 0x9e, 0x7b, 0xd8, 0x65, 0x49, 0x4b,
	0xef, 0x27, 0x38, 0xdd, 0xf5, 0xf3, 0x62, 0xbf, 0xb9, 0x41, 0xde, 0x1d, 0xfa, 0x6a, 0xcf, 0x31,
	0x9d, 0xb1, 0x45, 0x85, 0xfc, 0x0e, 0xf9, 0x06, 0xc3, 0xd6, 0x93, 0x8e, 0xb6, 0xe7, 0xb5, 0x29,
	0xbc, 0xe7, 0x5b, 0x63, 0x0f, 0x31, 0xe3, 0x42, 0xff, 0x44, 0xe5, 0x77, 0xc8, 0x35, 0x9c, 0xef,
	0x5f, 0x8e, 0x3c, 0xdd, 0x9e, 0xbe, 0x4f, 0x3d, 0xaf, 0xa1, 0x53, 0x8b, 0xea, 0x0b, 0x07, 0xa8,
	0xfd, 0xce, 0xe5, 0x2b, 0x78, 0x16, 0x46, 0xdc, 0x50, 0x81, 0xab, 0x3c, 0xe1, 0xb4, 0x8c, 0x51,
	0xdc, 0xc4, 0x02, 0xff, 0x76, 0x5d, 0x1e, 0x7f, 0x5e, 0x2e, 0xca, 0x60, 0xd6, 0x5d, 0xf5, 0xab,
	0xca, 0x9b, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x34, 0xbb, 0x7a, 0x95, 0x97, 0x03, 0x00, 0x00,
}
