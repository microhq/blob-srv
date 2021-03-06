// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/blob/blob.proto

package go_micro_srv_blob

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CreateBucketReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBucketReq) Reset()         { *m = CreateBucketReq{} }
func (m *CreateBucketReq) String() string { return proto.CompactTextString(m) }
func (*CreateBucketReq) ProtoMessage()    {}
func (*CreateBucketReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bd5654a139a5293, []int{0}
}

func (m *CreateBucketReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBucketReq.Unmarshal(m, b)
}
func (m *CreateBucketReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBucketReq.Marshal(b, m, deterministic)
}
func (m *CreateBucketReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBucketReq.Merge(m, src)
}
func (m *CreateBucketReq) XXX_Size() int {
	return xxx_messageInfo_CreateBucketReq.Size(m)
}
func (m *CreateBucketReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBucketReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBucketReq proto.InternalMessageInfo

func (m *CreateBucketReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateBucketResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBucketResp) Reset()         { *m = CreateBucketResp{} }
func (m *CreateBucketResp) String() string { return proto.CompactTextString(m) }
func (*CreateBucketResp) ProtoMessage()    {}
func (*CreateBucketResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bd5654a139a5293, []int{1}
}

func (m *CreateBucketResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBucketResp.Unmarshal(m, b)
}
func (m *CreateBucketResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBucketResp.Marshal(b, m, deterministic)
}
func (m *CreateBucketResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBucketResp.Merge(m, src)
}
func (m *CreateBucketResp) XXX_Size() int {
	return xxx_messageInfo_CreateBucketResp.Size(m)
}
func (m *CreateBucketResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBucketResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBucketResp proto.InternalMessageInfo

type DeleteBucketReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBucketReq) Reset()         { *m = DeleteBucketReq{} }
func (m *DeleteBucketReq) String() string { return proto.CompactTextString(m) }
func (*DeleteBucketReq) ProtoMessage()    {}
func (*DeleteBucketReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bd5654a139a5293, []int{2}
}

func (m *DeleteBucketReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBucketReq.Unmarshal(m, b)
}
func (m *DeleteBucketReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBucketReq.Marshal(b, m, deterministic)
}
func (m *DeleteBucketReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBucketReq.Merge(m, src)
}
func (m *DeleteBucketReq) XXX_Size() int {
	return xxx_messageInfo_DeleteBucketReq.Size(m)
}
func (m *DeleteBucketReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBucketReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBucketReq proto.InternalMessageInfo

func (m *DeleteBucketReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteBucketResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBucketResp) Reset()         { *m = DeleteBucketResp{} }
func (m *DeleteBucketResp) String() string { return proto.CompactTextString(m) }
func (*DeleteBucketResp) ProtoMessage()    {}
func (*DeleteBucketResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bd5654a139a5293, []int{3}
}

func (m *DeleteBucketResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBucketResp.Unmarshal(m, b)
}
func (m *DeleteBucketResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBucketResp.Marshal(b, m, deterministic)
}
func (m *DeleteBucketResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBucketResp.Merge(m, src)
}
func (m *DeleteBucketResp) XXX_Size() int {
	return xxx_messageInfo_DeleteBucketResp.Size(m)
}
func (m *DeleteBucketResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBucketResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBucketResp proto.InternalMessageInfo

type PutReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BucketId             string   `protobuf:"bytes,2,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutReq) Reset()         { *m = PutReq{} }
func (m *PutReq) String() string { return proto.CompactTextString(m) }
func (*PutReq) ProtoMessage()    {}
func (*PutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bd5654a139a5293, []int{4}
}

func (m *PutReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutReq.Unmarshal(m, b)
}
func (m *PutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutReq.Marshal(b, m, deterministic)
}
func (m *PutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutReq.Merge(m, src)
}
func (m *PutReq) XXX_Size() int {
	return xxx_messageInfo_PutReq.Size(m)
}
func (m *PutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PutReq.DiscardUnknown(m)
}

var xxx_messageInfo_PutReq proto.InternalMessageInfo

func (m *PutReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PutReq) GetBucketId() string {
	if m != nil {
		return m.BucketId
	}
	return ""
}

func (m *PutReq) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PutResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutResp) Reset()         { *m = PutResp{} }
func (m *PutResp) String() string { return proto.CompactTextString(m) }
func (*PutResp) ProtoMessage()    {}
func (*PutResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bd5654a139a5293, []int{5}
}

func (m *PutResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutResp.Unmarshal(m, b)
}
func (m *PutResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutResp.Marshal(b, m, deterministic)
}
func (m *PutResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutResp.Merge(m, src)
}
func (m *PutResp) XXX_Size() int {
	return xxx_messageInfo_PutResp.Size(m)
}
func (m *PutResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PutResp.DiscardUnknown(m)
}

var xxx_messageInfo_PutResp proto.InternalMessageInfo

type GetReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BucketId             string   `protobuf:"bytes,2,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReq) Reset()         { *m = GetReq{} }
func (m *GetReq) String() string { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()    {}
func (*GetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bd5654a139a5293, []int{6}
}

func (m *GetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReq.Unmarshal(m, b)
}
func (m *GetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReq.Marshal(b, m, deterministic)
}
func (m *GetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReq.Merge(m, src)
}
func (m *GetReq) XXX_Size() int {
	return xxx_messageInfo_GetReq.Size(m)
}
func (m *GetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetReq proto.InternalMessageInfo

func (m *GetReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetReq) GetBucketId() string {
	if m != nil {
		return m.BucketId
	}
	return ""
}

type GetResp struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResp) Reset()         { *m = GetResp{} }
func (m *GetResp) String() string { return proto.CompactTextString(m) }
func (*GetResp) ProtoMessage()    {}
func (*GetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bd5654a139a5293, []int{7}
}

func (m *GetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResp.Unmarshal(m, b)
}
func (m *GetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResp.Marshal(b, m, deterministic)
}
func (m *GetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResp.Merge(m, src)
}
func (m *GetResp) XXX_Size() int {
	return xxx_messageInfo_GetResp.Size(m)
}
func (m *GetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetResp proto.InternalMessageInfo

func (m *GetResp) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BucketId             string   `protobuf:"bytes,2,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteReq) Reset()         { *m = DeleteReq{} }
func (m *DeleteReq) String() string { return proto.CompactTextString(m) }
func (*DeleteReq) ProtoMessage()    {}
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bd5654a139a5293, []int{8}
}

func (m *DeleteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReq.Unmarshal(m, b)
}
func (m *DeleteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReq.Marshal(b, m, deterministic)
}
func (m *DeleteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReq.Merge(m, src)
}
func (m *DeleteReq) XXX_Size() int {
	return xxx_messageInfo_DeleteReq.Size(m)
}
func (m *DeleteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReq proto.InternalMessageInfo

func (m *DeleteReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteReq) GetBucketId() string {
	if m != nil {
		return m.BucketId
	}
	return ""
}

type DeleteResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResp) Reset()         { *m = DeleteResp{} }
func (m *DeleteResp) String() string { return proto.CompactTextString(m) }
func (*DeleteResp) ProtoMessage()    {}
func (*DeleteResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bd5654a139a5293, []int{9}
}

func (m *DeleteResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResp.Unmarshal(m, b)
}
func (m *DeleteResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResp.Marshal(b, m, deterministic)
}
func (m *DeleteResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResp.Merge(m, src)
}
func (m *DeleteResp) XXX_Size() int {
	return xxx_messageInfo_DeleteResp.Size(m)
}
func (m *DeleteResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResp proto.InternalMessageInfo

type ListReq struct {
	BucketId             string   `protobuf:"bytes,1,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListReq) Reset()         { *m = ListReq{} }
func (m *ListReq) String() string { return proto.CompactTextString(m) }
func (*ListReq) ProtoMessage()    {}
func (*ListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bd5654a139a5293, []int{10}
}

func (m *ListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReq.Unmarshal(m, b)
}
func (m *ListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReq.Marshal(b, m, deterministic)
}
func (m *ListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReq.Merge(m, src)
}
func (m *ListReq) XXX_Size() int {
	return xxx_messageInfo_ListReq.Size(m)
}
func (m *ListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListReq proto.InternalMessageInfo

func (m *ListReq) GetBucketId() string {
	if m != nil {
		return m.BucketId
	}
	return ""
}

type ListResp struct {
	Id                   []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResp) Reset()         { *m = ListResp{} }
func (m *ListResp) String() string { return proto.CompactTextString(m) }
func (*ListResp) ProtoMessage()    {}
func (*ListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bd5654a139a5293, []int{11}
}

func (m *ListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResp.Unmarshal(m, b)
}
func (m *ListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResp.Marshal(b, m, deterministic)
}
func (m *ListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResp.Merge(m, src)
}
func (m *ListResp) XXX_Size() int {
	return xxx_messageInfo_ListResp.Size(m)
}
func (m *ListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListResp proto.InternalMessageInfo

func (m *ListResp) GetId() []string {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateBucketReq)(nil), "go.micro.srv.blob.CreateBucketReq")
	proto.RegisterType((*CreateBucketResp)(nil), "go.micro.srv.blob.CreateBucketResp")
	proto.RegisterType((*DeleteBucketReq)(nil), "go.micro.srv.blob.DeleteBucketReq")
	proto.RegisterType((*DeleteBucketResp)(nil), "go.micro.srv.blob.DeleteBucketResp")
	proto.RegisterType((*PutReq)(nil), "go.micro.srv.blob.PutReq")
	proto.RegisterType((*PutResp)(nil), "go.micro.srv.blob.PutResp")
	proto.RegisterType((*GetReq)(nil), "go.micro.srv.blob.GetReq")
	proto.RegisterType((*GetResp)(nil), "go.micro.srv.blob.GetResp")
	proto.RegisterType((*DeleteReq)(nil), "go.micro.srv.blob.DeleteReq")
	proto.RegisterType((*DeleteResp)(nil), "go.micro.srv.blob.DeleteResp")
	proto.RegisterType((*ListReq)(nil), "go.micro.srv.blob.ListReq")
	proto.RegisterType((*ListResp)(nil), "go.micro.srv.blob.ListResp")
}

func init() { proto.RegisterFile("proto/blob/blob.proto", fileDescriptor_3bd5654a139a5293) }

var fileDescriptor_3bd5654a139a5293 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x97, 0x75, 0x74, 0xeb, 0x65, 0xf8, 0xe7, 0x82, 0x30, 0x3b, 0x07, 0x33, 0x82, 0xec,
	0x29, 0x8a, 0x22, 0xf8, 0xa8, 0x53, 0x28, 0x03, 0x1f, 0xc6, 0xde, 0x7c, 0x92, 0x75, 0x0d, 0x52,
	0x9c, 0x24, 0x36, 0x9d, 0x9f, 0xd2, 0x0f, 0x25, 0x49, 0xdc, 0x68, 0xbb, 0xd4, 0xb1, 0x97, 0xd2,
	0xe6, 0x9e, 0xf3, 0xbb, 0xa7, 0x9c, 0x16, 0x4e, 0x64, 0x26, 0x72, 0x71, 0x15, 0x2f, 0x45, 0x6c,
	0x2e, 0xcc, 0x3c, 0xe3, 0xf1, 0xbb, 0x60, 0x9f, 0xe9, 0x22, 0x13, 0x4c, 0x65, 0xdf, 0x4c, 0x0f,
	0xe8, 0x39, 0x1c, 0x3e, 0x65, 0x7c, 0x9e, 0xf3, 0xf1, 0x6a, 0xf1, 0xc1, 0xf3, 0x19, 0xff, 0xc2,
	0x03, 0x68, 0xa6, 0x49, 0x8f, 0x0c, 0xc9, 0x28, 0x98, 0x35, 0xd3, 0x84, 0x22, 0x1c, 0x95, 0x25,
	0x4a, 0x6a, 0xdb, 0x33, 0x5f, 0xf2, 0x1d, 0xb6, 0xb2, 0x44, 0x49, 0x3a, 0x01, 0x7f, 0xba, 0x72,
	0xa9, 0xb1, 0x0f, 0x41, 0x6c, 0x74, 0x6f, 0x69, 0xd2, 0x6b, 0x9a, 0xe3, 0x8e, 0x3d, 0x98, 0x24,
	0x88, 0xd0, 0x4a, 0xe6, 0xf9, 0xbc, 0xe7, 0x0d, 0xc9, 0xa8, 0x3b, 0x33, 0xf7, 0x34, 0x80, 0xb6,
	0x41, 0x29, 0x49, 0xef, 0xc0, 0x8f, 0xf8, 0xde, 0x54, 0x3a, 0x80, 0x76, 0x64, 0x73, 0x6d, 0x16,
	0x90, 0xc2, 0x82, 0x7b, 0x08, 0x6c, 0xfe, 0xbd, 0xc1, 0x5d, 0x80, 0xb5, 0x53, 0x49, 0x7a, 0x09,
	0xed, 0x97, 0x54, 0x99, 0x78, 0x25, 0x17, 0xa9, 0xb8, 0x42, 0xe8, 0x58, 0x9d, 0x92, 0x9b, 0x75,
	0x9e, 0x5d, 0x77, 0xf3, 0xe3, 0x41, 0x6b, 0xbc, 0x14, 0x31, 0xbe, 0x42, 0xb7, 0xd8, 0x05, 0x52,
	0xb6, 0x55, 0x29, 0xab, 0xf4, 0x19, 0x5e, 0xec, 0xd4, 0x28, 0x49, 0x1b, 0x1a, 0x5d, 0xec, 0xcb,
	0x89, 0xae, 0x74, 0xee, 0x44, 0x6f, 0x95, 0xde, 0xc0, 0x07, 0xf0, 0xa6, 0xab, 0x1c, 0x4f, 0x1d,
	0x6a, 0xfb, 0x39, 0x84, 0x61, 0xdd, 0x48, 0xfb, 0x47, 0x44, 0x13, 0x22, 0xee, 0x26, 0xd8, 0xea,
	0x9d, 0x84, 0x68, 0x9d, 0xe0, 0x9a, 0x60, 0x04, 0xbe, 0x4d, 0x86, 0x67, 0xb5, 0xa1, 0x35, 0x67,
	0xf0, 0xcf, 0xd4, 0xbc, 0xcc, 0x23, 0xb4, 0x74, 0x4f, 0xe8, 0x5a, 0xf8, 0x57, 0x74, 0xd8, 0xaf,
	0x9d, 0x69, 0x44, 0xec, 0x9b, 0xdf, 0xf1, 0xf6, 0x37, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x91, 0x9c,
	0x06, 0xa7, 0x03, 0x00, 0x00,
}
