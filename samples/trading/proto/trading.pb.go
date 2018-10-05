// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trading.proto

package tradingdb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Candle struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Open                 int64    `protobuf:"varint,3,opt,name=open,proto3" json:"open,omitempty"`
	Close                int64    `protobuf:"varint,4,opt,name=close,proto3" json:"close,omitempty"`
	High                 int64    `protobuf:"varint,5,opt,name=high,proto3" json:"high,omitempty"`
	Low                  int64    `protobuf:"varint,6,opt,name=low,proto3" json:"low,omitempty"`
	Curtime              int64    `protobuf:"varint,7,opt,name=curtime,proto3" json:"curtime,omitempty"`
	Volume               int64    `protobuf:"varint,8,opt,name=volume,proto3" json:"volume,omitempty"`
	OpenInterest         int64    `protobuf:"varint,9,opt,name=openInterest,proto3" json:"openInterest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Candle) Reset()         { *m = Candle{} }
func (m *Candle) String() string { return proto.CompactTextString(m) }
func (*Candle) ProtoMessage()    {}
func (*Candle) Descriptor() ([]byte, []int) {
	return fileDescriptor_trading_7c30d75e9c92fbd4, []int{0}
}
func (m *Candle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candle.Unmarshal(m, b)
}
func (m *Candle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candle.Marshal(b, m, deterministic)
}
func (dst *Candle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candle.Merge(dst, src)
}
func (m *Candle) XXX_Size() int {
	return xxx_messageInfo_Candle.Size(m)
}
func (m *Candle) XXX_DiscardUnknown() {
	xxx_messageInfo_Candle.DiscardUnknown(m)
}

var xxx_messageInfo_Candle proto.InternalMessageInfo

func (m *Candle) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Candle) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Candle) GetOpen() int64 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *Candle) GetClose() int64 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *Candle) GetHigh() int64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *Candle) GetLow() int64 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *Candle) GetCurtime() int64 {
	if m != nil {
		return m.Curtime
	}
	return 0
}

func (m *Candle) GetVolume() int64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Candle) GetOpenInterest() int64 {
	if m != nil {
		return m.OpenInterest
	}
	return 0
}

type CandleChunk struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StartTime            int64    `protobuf:"varint,3,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              int64    `protobuf:"varint,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CandleChunk) Reset()         { *m = CandleChunk{} }
func (m *CandleChunk) String() string { return proto.CompactTextString(m) }
func (*CandleChunk) ProtoMessage()    {}
func (*CandleChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_trading_7c30d75e9c92fbd4, []int{1}
}
func (m *CandleChunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CandleChunk.Unmarshal(m, b)
}
func (m *CandleChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CandleChunk.Marshal(b, m, deterministic)
}
func (dst *CandleChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CandleChunk.Merge(dst, src)
}
func (m *CandleChunk) XXX_Size() int {
	return xxx_messageInfo_CandleChunk.Size(m)
}
func (m *CandleChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_CandleChunk.DiscardUnknown(m)
}

var xxx_messageInfo_CandleChunk proto.InternalMessageInfo

func (m *CandleChunk) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CandleChunk) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CandleChunk) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *CandleChunk) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Candle)(nil), "tradingdb.Candle")
	proto.RegisterType((*CandleChunk)(nil), "tradingdb.CandleChunk")
}

func init() { proto.RegisterFile("trading.proto", fileDescriptor_trading_7c30d75e9c92fbd4) }

var fileDescriptor_trading_7c30d75e9c92fbd4 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbf, 0x4e, 0xc3, 0x40,
	0x0c, 0xc6, 0x15, 0xd2, 0xa6, 0xc4, 0x80, 0x84, 0x2c, 0x84, 0x3c, 0x30, 0x54, 0x99, 0x3a, 0xb1,
	0xf0, 0x08, 0x9d, 0x58, 0x2b, 0x5e, 0xe0, 0x9a, 0xb3, 0x9a, 0x88, 0xfb, 0x53, 0x5d, 0x2e, 0xf0,
	0xa6, 0x3c, 0x0f, 0xb2, 0xaf, 0x11, 0x62, 0xeb, 0xf6, 0xfd, 0x7e, 0xb6, 0x74, 0xfe, 0x0e, 0x1e,
	0x72, 0x32, 0x76, 0x0c, 0xa7, 0xd7, 0x73, 0x8a, 0x39, 0x62, 0x7b, 0x41, 0x7b, 0xec, 0x7e, 0x2a,
	0x68, 0xf6, 0x26, 0x58, 0xc7, 0x88, 0xb0, 0xea, 0xa3, 0x65, 0xaa, 0xb6, 0xd5, 0xae, 0x3d, 0x68,
	0x16, 0x17, 0x8c, 0x67, 0xba, 0x29, 0x4e, 0xb2, 0xb8, 0x78, 0xe6, 0x40, 0xf5, 0xb6, 0xda, 0xd5,
	0x07, 0xcd, 0xf8, 0x04, 0xeb, 0xde, 0xc5, 0x89, 0x69, 0xa5, 0xb2, 0x80, 0x6c, 0x0e, 0xe3, 0x69,
	0xa0, 0x75, 0xd9, 0x94, 0x8c, 0x8f, 0x50, 0xbb, 0xf8, 0x4d, 0x8d, 0x2a, 0x89, 0x48, 0xb0, 0xe9,
	0xe7, 0x94, 0x47, 0xcf, 0xb4, 0x51, 0xbb, 0x20, 0x3e, 0x43, 0xf3, 0x15, 0xdd, 0xec, 0x99, 0x6e,
	0x75, 0x70, 0x21, 0xec, 0xe0, 0x5e, 0x5e, 0x7d, 0x0f, 0x99, 0x13, 0x4f, 0x99, 0x5a, 0x9d, 0xfe,
	0x73, 0x9d, 0x87, 0xbb, 0xd2, 0x6b, 0x3f, 0xcc, 0xe1, 0xf3, 0xea, 0x72, 0x2f, 0xd0, 0x4e, 0xd9,
	0xa4, 0xfc, 0x21, 0xe7, 0x94, 0x86, 0x7f, 0x42, 0x4e, 0xe5, 0x60, 0x75, 0x56, 0x8a, 0x2e, 0x78,
	0x6c, 0xf4, 0x67, 0xdf, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xeb, 0x4b, 0xa0, 0x6a, 0x01,
	0x00, 0x00,
}
