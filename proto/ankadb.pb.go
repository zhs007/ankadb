// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ankadb.proto

package ankadbpb

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

type CODE int32

const (
	CODE_OK                  CODE = 0
	CODE_VAR_PARSE_ERR       CODE = 1
	CODE_LOGIC_ONQUERY_ERR   CODE = 2
	CODE_CTX_CURDB_ERR       CODE = 3
	CODE_PROTOBUF_ENCODE_ERR CODE = 4
	CODE_DB_PUT_ERR          CODE = 5
	CODE_RESULT_NO_DATA      CODE = 6
	CODE_RESULT_DATA_INVALID CODE = 7
	CODE_INPUTOBJ_PARSE_ERR  CODE = 8
	CODE_ONQUERYSTREAM_ERR   CODE = 9
	CODE_CTX_ANKADB_ERR      CODE = 10
	CODE_TIMEZONE_ERR        CODE = 11
	CODE_CTX_SNAPSHOTMGR_ERR CODE = 12
	CODE_MAKE_SNAPSHOT_ERR   CODE = 13
	CODE_QUERY_PARAM_ERR     CODE = 14
	CODE_HTTP_BODY_PARSE_ERR CODE = 100
	CODE_HTTP_NO_QUERY       CODE = 101
	CODE_HTTP_VARIABLE_ERR   CODE = 102
	CODE_CLIENT_NO_CONN      CODE = 10000
	CODE_INVALID_CODE        CODE = 20000
)

var CODE_name = map[int32]string{
	0:     "OK",
	1:     "VAR_PARSE_ERR",
	2:     "LOGIC_ONQUERY_ERR",
	3:     "CTX_CURDB_ERR",
	4:     "PROTOBUF_ENCODE_ERR",
	5:     "DB_PUT_ERR",
	6:     "RESULT_NO_DATA",
	7:     "RESULT_DATA_INVALID",
	8:     "INPUTOBJ_PARSE_ERR",
	9:     "ONQUERYSTREAM_ERR",
	10:    "CTX_ANKADB_ERR",
	11:    "TIMEZONE_ERR",
	12:    "CTX_SNAPSHOTMGR_ERR",
	13:    "MAKE_SNAPSHOT_ERR",
	14:    "QUERY_PARAM_ERR",
	100:   "HTTP_BODY_PARSE_ERR",
	101:   "HTTP_NO_QUERY",
	102:   "HTTP_VARIABLE_ERR",
	10000: "CLIENT_NO_CONN",
	20000: "INVALID_CODE",
}
var CODE_value = map[string]int32{
	"OK":                  0,
	"VAR_PARSE_ERR":       1,
	"LOGIC_ONQUERY_ERR":   2,
	"CTX_CURDB_ERR":       3,
	"PROTOBUF_ENCODE_ERR": 4,
	"DB_PUT_ERR":          5,
	"RESULT_NO_DATA":      6,
	"RESULT_DATA_INVALID": 7,
	"INPUTOBJ_PARSE_ERR":  8,
	"ONQUERYSTREAM_ERR":   9,
	"CTX_ANKADB_ERR":      10,
	"TIMEZONE_ERR":        11,
	"CTX_SNAPSHOTMGR_ERR": 12,
	"MAKE_SNAPSHOT_ERR":   13,
	"QUERY_PARAM_ERR":     14,
	"HTTP_BODY_PARSE_ERR": 100,
	"HTTP_NO_QUERY":       101,
	"HTTP_VARIABLE_ERR":   102,
	"CLIENT_NO_CONN":      10000,
	"INVALID_CODE":        20000,
}

func (x CODE) String() string {
	return proto.EnumName(CODE_name, int32(x))
}
func (CODE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ankadb_e6683bd58904a052, []int{0}
}

type SnapshotMgr struct {
	MaxSnapshotID        int64    `protobuf:"varint,1,opt,name=maxSnapshotID,proto3" json:"maxSnapshotID,omitempty"`
	Snapshots            []int64  `protobuf:"varint,2,rep,packed,name=snapshots,proto3" json:"snapshots,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnapshotMgr) Reset()         { *m = SnapshotMgr{} }
func (m *SnapshotMgr) String() string { return proto.CompactTextString(m) }
func (*SnapshotMgr) ProtoMessage()    {}
func (*SnapshotMgr) Descriptor() ([]byte, []int) {
	return fileDescriptor_ankadb_e6683bd58904a052, []int{0}
}
func (m *SnapshotMgr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotMgr.Unmarshal(m, b)
}
func (m *SnapshotMgr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotMgr.Marshal(b, m, deterministic)
}
func (dst *SnapshotMgr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotMgr.Merge(dst, src)
}
func (m *SnapshotMgr) XXX_Size() int {
	return xxx_messageInfo_SnapshotMgr.Size(m)
}
func (m *SnapshotMgr) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotMgr.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotMgr proto.InternalMessageInfo

func (m *SnapshotMgr) GetMaxSnapshotID() int64 {
	if m != nil {
		return m.MaxSnapshotID
	}
	return 0
}

func (m *SnapshotMgr) GetSnapshots() []int64 {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

type Snapshot struct {
	SnapshotID           int64    `protobuf:"varint,1,opt,name=snapshotID,proto3" json:"snapshotID,omitempty"`
	Keys                 []string `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	CreateTime           int64    `protobuf:"varint,3,opt,name=createTime,proto3" json:"createTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_ankadb_e6683bd58904a052, []int{1}
}
func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Snapshot.Unmarshal(m, b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
}
func (dst *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(dst, src)
}
func (m *Snapshot) XXX_Size() int {
	return xxx_messageInfo_Snapshot.Size(m)
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetSnapshotID() int64 {
	if m != nil {
		return m.SnapshotID
	}
	return 0
}

func (m *Snapshot) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Snapshot) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

type Query struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // Deprecated: Do not use.
	QueryData            string   `protobuf:"bytes,2,opt,name=queryData,proto3" json:"queryData,omitempty"`
	VarData              string   `protobuf:"bytes,3,opt,name=varData,proto3" json:"varData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_ankadb_e6683bd58904a052, []int{2}
}
func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (dst *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(dst, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *Query) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Query) GetQueryData() string {
	if m != nil {
		return m.QueryData
	}
	return ""
}

func (m *Query) GetVarData() string {
	if m != nil {
		return m.VarData
	}
	return ""
}

type ReplyQuery struct {
	Code                 CODE     `protobuf:"varint,1,opt,name=code,proto3,enum=ankadbpb.CODE" json:"code,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Result               string   `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyQuery) Reset()         { *m = ReplyQuery{} }
func (m *ReplyQuery) String() string { return proto.CompactTextString(m) }
func (*ReplyQuery) ProtoMessage()    {}
func (*ReplyQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_ankadb_e6683bd58904a052, []int{3}
}
func (m *ReplyQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyQuery.Unmarshal(m, b)
}
func (m *ReplyQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyQuery.Marshal(b, m, deterministic)
}
func (dst *ReplyQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyQuery.Merge(dst, src)
}
func (m *ReplyQuery) XXX_Size() int {
	return xxx_messageInfo_ReplyQuery.Size(m)
}
func (m *ReplyQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyQuery proto.InternalMessageInfo

func (m *ReplyQuery) GetCode() CODE {
	if m != nil {
		return m.Code
	}
	return CODE_OK
}

func (m *ReplyQuery) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *ReplyQuery) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*SnapshotMgr)(nil), "ankadbpb.SnapshotMgr")
	proto.RegisterType((*Snapshot)(nil), "ankadbpb.Snapshot")
	proto.RegisterType((*Query)(nil), "ankadbpb.Query")
	proto.RegisterType((*ReplyQuery)(nil), "ankadbpb.ReplyQuery")
	proto.RegisterEnum("ankadbpb.CODE", CODE_name, CODE_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AnkaDBServClient is the client API for AnkaDBServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnkaDBServClient interface {
	Query(ctx context.Context, in *Query, opts ...grpc.CallOption) (*ReplyQuery, error)
	QueryStream(ctx context.Context, in *Query, opts ...grpc.CallOption) (AnkaDBServ_QueryStreamClient, error)
}

type ankaDBServClient struct {
	cc *grpc.ClientConn
}

func NewAnkaDBServClient(cc *grpc.ClientConn) AnkaDBServClient {
	return &ankaDBServClient{cc}
}

func (c *ankaDBServClient) Query(ctx context.Context, in *Query, opts ...grpc.CallOption) (*ReplyQuery, error) {
	out := new(ReplyQuery)
	err := c.cc.Invoke(ctx, "/ankadbpb.AnkaDBServ/query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ankaDBServClient) QueryStream(ctx context.Context, in *Query, opts ...grpc.CallOption) (AnkaDBServ_QueryStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AnkaDBServ_serviceDesc.Streams[0], "/ankadbpb.AnkaDBServ/queryStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &ankaDBServQueryStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AnkaDBServ_QueryStreamClient interface {
	Recv() (*ReplyQuery, error)
	grpc.ClientStream
}

type ankaDBServQueryStreamClient struct {
	grpc.ClientStream
}

func (x *ankaDBServQueryStreamClient) Recv() (*ReplyQuery, error) {
	m := new(ReplyQuery)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AnkaDBServServer is the server API for AnkaDBServ service.
type AnkaDBServServer interface {
	Query(context.Context, *Query) (*ReplyQuery, error)
	QueryStream(*Query, AnkaDBServ_QueryStreamServer) error
}

func RegisterAnkaDBServServer(s *grpc.Server, srv AnkaDBServServer) {
	s.RegisterService(&_AnkaDBServ_serviceDesc, srv)
}

func _AnkaDBServ_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnkaDBServServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankadbpb.AnkaDBServ/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnkaDBServServer).Query(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnkaDBServ_QueryStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Query)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AnkaDBServServer).QueryStream(m, &ankaDBServQueryStreamServer{stream})
}

type AnkaDBServ_QueryStreamServer interface {
	Send(*ReplyQuery) error
	grpc.ServerStream
}

type ankaDBServQueryStreamServer struct {
	grpc.ServerStream
}

func (x *ankaDBServQueryStreamServer) Send(m *ReplyQuery) error {
	return x.ServerStream.SendMsg(m)
}

var _AnkaDBServ_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ankadbpb.AnkaDBServ",
	HandlerType: (*AnkaDBServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "query",
			Handler:    _AnkaDBServ_Query_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "queryStream",
			Handler:       _AnkaDBServ_QueryStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ankadb.proto",
}

func init() { proto.RegisterFile("ankadb.proto", fileDescriptor_ankadb_e6683bd58904a052) }

var fileDescriptor_ankadb_e6683bd58904a052 = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4d, 0x4f, 0xdc, 0x3e,
	0x10, 0xc6, 0xc9, 0x66, 0x59, 0xd8, 0x61, 0x09, 0x61, 0xf8, 0xff, 0x69, 0x54, 0x55, 0x15, 0x8a,
	0x7a, 0x40, 0x3d, 0x20, 0x44, 0x2f, 0xbd, 0x3a, 0x2f, 0x85, 0x94, 0x5d, 0x3b, 0x38, 0x0e, 0x2d,
	0x1c, 0x6a, 0x05, 0x70, 0x5f, 0x04, 0xfb, 0xd2, 0x6c, 0x40, 0xf0, 0x2d, 0xfa, 0x01, 0x7a, 0xe8,
	0xad, 0x5f, 0xb3, 0xb2, 0x93, 0x6d, 0x68, 0x4f, 0xbd, 0x65, 0x7e, 0xcf, 0x3c, 0x8f, 0xc7, 0xd6,
	0x04, 0x06, 0xc5, 0xe4, 0xba, 0xb8, 0xba, 0xd8, 0x9b, 0x95, 0xd3, 0x6a, 0x8a, 0xab, 0x75, 0x35,
	0xbb, 0xf0, 0x4f, 0x60, 0x2d, 0x9b, 0x14, 0xb3, 0xf9, 0xe7, 0x69, 0x35, 0xfa, 0x54, 0xe2, 0x0b,
	0x58, 0x1f, 0x17, 0xf7, 0x0b, 0x92, 0x44, 0x9e, 0xb5, 0x63, 0xed, 0xda, 0xfc, 0x4f, 0x88, 0xcf,
	0xa0, 0x3f, 0x6f, 0xaa, 0xb9, 0xd7, 0xd9, 0xb1, 0x77, 0x6d, 0xde, 0x02, 0xff, 0x03, 0xac, 0x2e,
	0x7a, 0xf1, 0x39, 0xc0, 0xfc, 0xef, 0xb0, 0x47, 0x04, 0x11, 0xba, 0xd7, 0xea, 0xa1, 0x0e, 0xe9,
	0x73, 0xf3, 0xad, 0x3d, 0x97, 0xa5, 0x2a, 0x2a, 0x25, 0xbe, 0x8c, 0x95, 0x67, 0xd7, 0x9e, 0x96,
	0xf8, 0xef, 0x60, 0xf9, 0xe4, 0x56, 0x95, 0x0f, 0xb8, 0x0d, 0xdd, 0x49, 0x31, 0x56, 0x26, 0xb6,
	0x1f, 0x74, 0x3c, 0x8b, 0x9b, 0x5a, 0x8f, 0xf7, 0x55, 0x37, 0x44, 0x45, 0x55, 0x78, 0x1d, 0x2d,
	0xf2, 0x16, 0xa0, 0x07, 0x2b, 0x77, 0x45, 0x69, 0x34, 0xdb, 0x68, 0x8b, 0xd2, 0x3f, 0x07, 0xe0,
	0x6a, 0x76, 0xf3, 0x50, 0xa7, 0xfb, 0xd0, 0xbd, 0x9c, 0x5e, 0xd5, 0xe9, 0xce, 0x81, 0xb3, 0xb7,
	0x78, 0xb2, 0xbd, 0x90, 0x45, 0x31, 0x37, 0x1a, 0xba, 0x60, 0xab, 0xb2, 0x6c, 0xce, 0xd0, 0x9f,
	0xb8, 0x0d, 0xbd, 0x52, 0xcd, 0x6f, 0x6f, 0xaa, 0x26, 0xbc, 0xa9, 0x5e, 0xfe, 0xb4, 0xa1, 0xab,
	0x8d, 0xd8, 0x83, 0x0e, 0x3b, 0x76, 0x97, 0x70, 0x13, 0xd6, 0x4f, 0x09, 0x97, 0x29, 0xe1, 0x59,
	0x2c, 0x63, 0xce, 0x5d, 0x0b, 0xff, 0x87, 0xcd, 0x21, 0x3b, 0x4c, 0x42, 0xc9, 0xe8, 0x49, 0x1e,
	0xf3, 0x33, 0x83, 0x3b, 0xba, 0x33, 0x14, 0xef, 0x65, 0x98, 0xf3, 0x28, 0x30, 0xc8, 0xc6, 0x27,
	0xb0, 0x95, 0x72, 0x26, 0x58, 0x90, 0xbf, 0x91, 0x31, 0xd5, 0xb9, 0x46, 0xe8, 0xa2, 0x03, 0x10,
	0x05, 0x32, 0xcd, 0x85, 0xa9, 0x97, 0x11, 0xc1, 0xe1, 0x71, 0x96, 0x0f, 0x85, 0xa4, 0x4c, 0x46,
	0x44, 0x10, 0xb7, 0xa7, 0xcd, 0x0d, 0xd3, 0x40, 0x26, 0xf4, 0x94, 0x0c, 0x93, 0xc8, 0x5d, 0xc1,
	0x6d, 0xc0, 0x84, 0xa6, 0xb9, 0x60, 0xc1, 0xdb, 0x47, 0x73, 0xad, 0xea, 0xb9, 0x9a, 0x89, 0x32,
	0xc1, 0x63, 0x32, 0x32, 0xb8, 0xaf, 0xb3, 0xf5, 0x5c, 0x84, 0x1e, 0x93, 0x66, 0x30, 0x40, 0x17,
	0x06, 0x22, 0x19, 0xc5, 0xe7, 0x8c, 0xd6, 0xe6, 0x35, 0x7d, 0x9a, 0xee, 0xca, 0x28, 0x49, 0xb3,
	0x23, 0x26, 0x46, 0x87, 0xdc, 0x08, 0x03, 0x9d, 0x3a, 0x22, 0xc7, 0xf1, 0x6f, 0xc5, 0xe0, 0x75,
	0xdc, 0x82, 0x8d, 0xfa, 0xf2, 0x29, 0xe1, 0xcd, 0x51, 0x8e, 0x0e, 0x39, 0x12, 0x22, 0x95, 0x01,
	0x8b, 0xce, 0x1e, 0x8d, 0x76, 0xa5, 0xdf, 0xc6, 0x08, 0x94, 0x49, 0xe3, 0x72, 0x95, 0xce, 0x35,
	0xe8, 0x94, 0xf0, 0x84, 0x04, 0xc3, 0xba, 0xf3, 0x23, 0x6e, 0x81, 0x13, 0x0e, 0x93, 0x98, 0x9a,
	0x97, 0x08, 0x19, 0xa5, 0xee, 0x37, 0x8a, 0x08, 0x83, 0xe6, 0xfa, 0x52, 0x3f, 0xa2, 0xfb, 0xe3,
	0xbb, 0x75, 0x70, 0x0f, 0x40, 0x26, 0xd7, 0x45, 0x14, 0x64, 0xaa, 0xbc, 0xc3, 0x7d, 0x58, 0x36,
	0xab, 0x83, 0x1b, 0xed, 0x02, 0x98, 0xfd, 0x78, 0xfa, 0x5f, 0x0b, 0xda, 0xad, 0xf1, 0x97, 0xf0,
	0x35, 0xac, 0x19, 0x47, 0x56, 0x95, 0xaa, 0x18, 0xff, 0xb3, 0x6f, 0xdf, 0xba, 0xe8, 0x99, 0x9f,
	0xf3, 0xd5, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0xcf, 0xed, 0x1d, 0xac, 0x03, 0x00, 0x00,
}
