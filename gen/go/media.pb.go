// Code generated by protoc-gen-go. DO NOT EDIT.
// source: media.proto

package protobufpb

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

type GetAllMediaRes struct {
	Medias               []*Media `protobuf:"bytes,1,rep,name=Medias,proto3" json:"Medias"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllMediaRes) Reset()         { *m = GetAllMediaRes{} }
func (m *GetAllMediaRes) String() string { return proto.CompactTextString(m) }
func (*GetAllMediaRes) ProtoMessage()    {}
func (*GetAllMediaRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb54b56db72a97, []int{0}
}

func (m *GetAllMediaRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllMediaRes.Unmarshal(m, b)
}
func (m *GetAllMediaRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllMediaRes.Marshal(b, m, deterministic)
}
func (m *GetAllMediaRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllMediaRes.Merge(m, src)
}
func (m *GetAllMediaRes) XXX_Size() int {
	return xxx_messageInfo_GetAllMediaRes.Size(m)
}
func (m *GetAllMediaRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllMediaRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllMediaRes proto.InternalMessageInfo

func (m *GetAllMediaRes) GetMedias() []*Media {
	if m != nil {
		return m.Medias
	}
	return nil
}

type GetAllMediaReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllMediaReq) Reset()         { *m = GetAllMediaReq{} }
func (m *GetAllMediaReq) String() string { return proto.CompactTextString(m) }
func (*GetAllMediaReq) ProtoMessage()    {}
func (*GetAllMediaReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb54b56db72a97, []int{1}
}

func (m *GetAllMediaReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllMediaReq.Unmarshal(m, b)
}
func (m *GetAllMediaReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllMediaReq.Marshal(b, m, deterministic)
}
func (m *GetAllMediaReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllMediaReq.Merge(m, src)
}
func (m *GetAllMediaReq) XXX_Size() int {
	return xxx_messageInfo_GetAllMediaReq.Size(m)
}
func (m *GetAllMediaReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllMediaReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllMediaReq proto.InternalMessageInfo

type GetDetailMediaReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDetailMediaReq) Reset()         { *m = GetDetailMediaReq{} }
func (m *GetDetailMediaReq) String() string { return proto.CompactTextString(m) }
func (*GetDetailMediaReq) ProtoMessage()    {}
func (*GetDetailMediaReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb54b56db72a97, []int{2}
}

func (m *GetDetailMediaReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDetailMediaReq.Unmarshal(m, b)
}
func (m *GetDetailMediaReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDetailMediaReq.Marshal(b, m, deterministic)
}
func (m *GetDetailMediaReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDetailMediaReq.Merge(m, src)
}
func (m *GetDetailMediaReq) XXX_Size() int {
	return xxx_messageInfo_GetDetailMediaReq.Size(m)
}
func (m *GetDetailMediaReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDetailMediaReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetDetailMediaReq proto.InternalMessageInfo

func (m *GetDetailMediaReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetDetailMediaRes struct {
	Media                *Media   `protobuf:"bytes,1,opt,name=Media,proto3" json:"Media"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDetailMediaRes) Reset()         { *m = GetDetailMediaRes{} }
func (m *GetDetailMediaRes) String() string { return proto.CompactTextString(m) }
func (*GetDetailMediaRes) ProtoMessage()    {}
func (*GetDetailMediaRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb54b56db72a97, []int{3}
}

func (m *GetDetailMediaRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDetailMediaRes.Unmarshal(m, b)
}
func (m *GetDetailMediaRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDetailMediaRes.Marshal(b, m, deterministic)
}
func (m *GetDetailMediaRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDetailMediaRes.Merge(m, src)
}
func (m *GetDetailMediaRes) XXX_Size() int {
	return xxx_messageInfo_GetDetailMediaRes.Size(m)
}
func (m *GetDetailMediaRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDetailMediaRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetDetailMediaRes proto.InternalMessageInfo

func (m *GetDetailMediaRes) GetMedia() *Media {
	if m != nil {
		return m.Media
	}
	return nil
}

type CreateCategoryRes struct {
	Category             *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateCategoryRes) Reset()         { *m = CreateCategoryRes{} }
func (m *CreateCategoryRes) String() string { return proto.CompactTextString(m) }
func (*CreateCategoryRes) ProtoMessage()    {}
func (*CreateCategoryRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb54b56db72a97, []int{4}
}

func (m *CreateCategoryRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCategoryRes.Unmarshal(m, b)
}
func (m *CreateCategoryRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCategoryRes.Marshal(b, m, deterministic)
}
func (m *CreateCategoryRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCategoryRes.Merge(m, src)
}
func (m *CreateCategoryRes) XXX_Size() int {
	return xxx_messageInfo_CreateCategoryRes.Size(m)
}
func (m *CreateCategoryRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCategoryRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCategoryRes proto.InternalMessageInfo

func (m *CreateCategoryRes) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

type GetListCategoryRes struct {
	Categories           []*Category `protobuf:"bytes,1,rep,name=Categories,proto3" json:"Categories"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetListCategoryRes) Reset()         { *m = GetListCategoryRes{} }
func (m *GetListCategoryRes) String() string { return proto.CompactTextString(m) }
func (*GetListCategoryRes) ProtoMessage()    {}
func (*GetListCategoryRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb54b56db72a97, []int{5}
}

func (m *GetListCategoryRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListCategoryRes.Unmarshal(m, b)
}
func (m *GetListCategoryRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListCategoryRes.Marshal(b, m, deterministic)
}
func (m *GetListCategoryRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListCategoryRes.Merge(m, src)
}
func (m *GetListCategoryRes) XXX_Size() int {
	return xxx_messageInfo_GetListCategoryRes.Size(m)
}
func (m *GetListCategoryRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListCategoryRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetListCategoryRes proto.InternalMessageInfo

func (m *GetListCategoryRes) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type CreateMediaRes struct {
	Media                *Media   `protobuf:"bytes,1,opt,name=media,proto3" json:"media"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMediaRes) Reset()         { *m = CreateMediaRes{} }
func (m *CreateMediaRes) String() string { return proto.CompactTextString(m) }
func (*CreateMediaRes) ProtoMessage()    {}
func (*CreateMediaRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb54b56db72a97, []int{6}
}

func (m *CreateMediaRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMediaRes.Unmarshal(m, b)
}
func (m *CreateMediaRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMediaRes.Marshal(b, m, deterministic)
}
func (m *CreateMediaRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMediaRes.Merge(m, src)
}
func (m *CreateMediaRes) XXX_Size() int {
	return xxx_messageInfo_CreateMediaRes.Size(m)
}
func (m *CreateMediaRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMediaRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMediaRes proto.InternalMessageInfo

func (m *CreateMediaRes) GetMedia() *Media {
	if m != nil {
		return m.Media
	}
	return nil
}

type CreateMediaReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	Creator              string   `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator"`
	Image                string   `protobuf:"bytes,5,opt,name=image,proto3" json:"image"`
	CategoryId           string   `protobuf:"bytes,6,opt,name=category_id,json=categoryId,proto3" json:"category_id"`
	Tags                 []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags"`
	Signature            string   `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMediaReq) Reset()         { *m = CreateMediaReq{} }
func (m *CreateMediaReq) String() string { return proto.CompactTextString(m) }
func (*CreateMediaReq) ProtoMessage()    {}
func (*CreateMediaReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb54b56db72a97, []int{7}
}

func (m *CreateMediaReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMediaReq.Unmarshal(m, b)
}
func (m *CreateMediaReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMediaReq.Marshal(b, m, deterministic)
}
func (m *CreateMediaReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMediaReq.Merge(m, src)
}
func (m *CreateMediaReq) XXX_Size() int {
	return xxx_messageInfo_CreateMediaReq.Size(m)
}
func (m *CreateMediaReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMediaReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMediaReq proto.InternalMessageInfo

func (m *CreateMediaReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateMediaReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateMediaReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateMediaReq) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *CreateMediaReq) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *CreateMediaReq) GetCategoryId() string {
	if m != nil {
		return m.CategoryId
	}
	return ""
}

func (m *CreateMediaReq) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *CreateMediaReq) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type CreateMediaAgainReq struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMediaAgainReq) Reset()         { *m = CreateMediaAgainReq{} }
func (m *CreateMediaAgainReq) String() string { return proto.CompactTextString(m) }
func (*CreateMediaAgainReq) ProtoMessage()    {}
func (*CreateMediaAgainReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb54b56db72a97, []int{8}
}

func (m *CreateMediaAgainReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMediaAgainReq.Unmarshal(m, b)
}
func (m *CreateMediaAgainReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMediaAgainReq.Marshal(b, m, deterministic)
}
func (m *CreateMediaAgainReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMediaAgainReq.Merge(m, src)
}
func (m *CreateMediaAgainReq) XXX_Size() int {
	return xxx_messageInfo_CreateMediaAgainReq.Size(m)
}
func (m *CreateMediaAgainReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMediaAgainReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMediaAgainReq proto.InternalMessageInfo

func (m *CreateMediaAgainReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateCategoryReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description"`
	AccessKey            string   `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCategoryReq) Reset()         { *m = CreateCategoryReq{} }
func (m *CreateCategoryReq) String() string { return proto.CompactTextString(m) }
func (*CreateCategoryReq) ProtoMessage()    {}
func (*CreateCategoryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb54b56db72a97, []int{9}
}

func (m *CreateCategoryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCategoryReq.Unmarshal(m, b)
}
func (m *CreateCategoryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCategoryReq.Marshal(b, m, deterministic)
}
func (m *CreateCategoryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCategoryReq.Merge(m, src)
}
func (m *CreateCategoryReq) XXX_Size() int {
	return xxx_messageInfo_CreateCategoryReq.Size(m)
}
func (m *CreateCategoryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCategoryReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCategoryReq proto.InternalMessageInfo

func (m *CreateCategoryReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCategoryReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateCategoryReq) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

type GetListCategoryReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListCategoryReq) Reset()         { *m = GetListCategoryReq{} }
func (m *GetListCategoryReq) String() string { return proto.CompactTextString(m) }
func (*GetListCategoryReq) ProtoMessage()    {}
func (*GetListCategoryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb54b56db72a97, []int{10}
}

func (m *GetListCategoryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListCategoryReq.Unmarshal(m, b)
}
func (m *GetListCategoryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListCategoryReq.Marshal(b, m, deterministic)
}
func (m *GetListCategoryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListCategoryReq.Merge(m, src)
}
func (m *GetListCategoryReq) XXX_Size() int {
	return xxx_messageInfo_GetListCategoryReq.Size(m)
}
func (m *GetListCategoryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListCategoryReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetListCategoryReq proto.InternalMessageInfo

type AddSignatureReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Creator              string   `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSignatureReq) Reset()         { *m = AddSignatureReq{} }
func (m *AddSignatureReq) String() string { return proto.CompactTextString(m) }
func (*AddSignatureReq) ProtoMessage()    {}
func (*AddSignatureReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb54b56db72a97, []int{11}
}

func (m *AddSignatureReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSignatureReq.Unmarshal(m, b)
}
func (m *AddSignatureReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSignatureReq.Marshal(b, m, deterministic)
}
func (m *AddSignatureReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSignatureReq.Merge(m, src)
}
func (m *AddSignatureReq) XXX_Size() int {
	return xxx_messageInfo_AddSignatureReq.Size(m)
}
func (m *AddSignatureReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSignatureReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddSignatureReq proto.InternalMessageInfo

func (m *AddSignatureReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddSignatureReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddSignatureReq) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type AddSignatureRes struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSignatureRes) Reset()         { *m = AddSignatureRes{} }
func (m *AddSignatureRes) String() string { return proto.CompactTextString(m) }
func (*AddSignatureRes) ProtoMessage()    {}
func (*AddSignatureRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb54b56db72a97, []int{12}
}

func (m *AddSignatureRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSignatureRes.Unmarshal(m, b)
}
func (m *AddSignatureRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSignatureRes.Marshal(b, m, deterministic)
}
func (m *AddSignatureRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSignatureRes.Merge(m, src)
}
func (m *AddSignatureRes) XXX_Size() int {
	return xxx_messageInfo_AddSignatureRes.Size(m)
}
func (m *AddSignatureRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSignatureRes.DiscardUnknown(m)
}

var xxx_messageInfo_AddSignatureRes proto.InternalMessageInfo

func (m *AddSignatureRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*GetAllMediaRes)(nil), "protobuf.GetAllMediaRes")
	proto.RegisterType((*GetAllMediaReq)(nil), "protobuf.GetAllMediaReq")
	proto.RegisterType((*GetDetailMediaReq)(nil), "protobuf.GetDetailMediaReq")
	proto.RegisterType((*GetDetailMediaRes)(nil), "protobuf.GetDetailMediaRes")
	proto.RegisterType((*CreateCategoryRes)(nil), "protobuf.CreateCategoryRes")
	proto.RegisterType((*GetListCategoryRes)(nil), "protobuf.GetListCategoryRes")
	proto.RegisterType((*CreateMediaRes)(nil), "protobuf.CreateMediaRes")
	proto.RegisterType((*CreateMediaReq)(nil), "protobuf.CreateMediaReq")
	proto.RegisterType((*CreateMediaAgainReq)(nil), "protobuf.CreateMediaAgainReq")
	proto.RegisterType((*CreateCategoryReq)(nil), "protobuf.CreateCategoryReq")
	proto.RegisterType((*GetListCategoryReq)(nil), "protobuf.GetListCategoryReq")
	proto.RegisterType((*AddSignatureReq)(nil), "protobuf.AddSignatureReq")
	proto.RegisterType((*AddSignatureRes)(nil), "protobuf.AddSignatureRes")
}

func init() { proto.RegisterFile("media.proto", fileDescriptor_07eb54b56db72a97) }

var fileDescriptor_07eb54b56db72a97 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdd, 0x6e, 0x12, 0x41,
	0x14, 0xee, 0xf2, 0x57, 0x38, 0x5b, 0xa1, 0x1d, 0x7b, 0x31, 0xae, 0x6d, 0x24, 0x63, 0x1a, 0x9b,
	0x34, 0xc1, 0x04, 0x2f, 0x8c, 0x7a, 0x45, 0x31, 0xa2, 0x91, 0xa6, 0x0d, 0xbd, 0xf3, 0xa6, 0x19,
	0x76, 0x47, 0x9c, 0x58, 0x58, 0xd8, 0x19, 0x4c, 0x78, 0x13, 0x5f, 0xcd, 0xb7, 0x31, 0x73, 0x96,
	0x81, 0xd9, 0xe5, 0x27, 0xbd, 0xda, 0x39, 0xe7, 0xfb, 0xce, 0xb7, 0xe7, 0x17, 0xfc, 0xb1, 0x88,
	0x24, 0x6f, 0x4d, 0x93, 0x58, 0xc7, 0xa4, 0x8a, 0x9f, 0xe1, 0xfc, 0x67, 0x00, 0x43, 0xae, 0x44,
	0xea, 0x65, 0x1f, 0xa0, 0xde, 0x13, 0xba, 0xf3, 0xf8, 0x78, 0x63, 0xa8, 0x03, 0xa1, 0xc8, 0x1b,
	0xa8, 0xe0, 0x5b, 0x51, 0xaf, 0x59, 0xbc, 0xf4, 0xdb, 0x8d, 0x96, 0x0d, 0x6c, 0xa5, 0x9c, 0x25,
	0xcc, 0x8e, 0x73, 0xa1, 0x33, 0xf6, 0x1a, 0x4e, 0x7a, 0x42, 0x7f, 0x16, 0x9a, 0xcb, 0x95, 0x93,
	0xd4, 0xa1, 0x20, 0x23, 0xea, 0x35, 0xbd, 0xcb, 0xda, 0xa0, 0x20, 0x23, 0xf6, 0x71, 0x93, 0xa4,
	0xc8, 0x05, 0x94, 0xf1, 0x8d, 0xbc, 0x2d, 0xff, 0x4c, 0x51, 0xd6, 0x85, 0x93, 0x6e, 0x22, 0xb8,
	0x16, 0x5d, 0xae, 0xc5, 0x28, 0x4e, 0x16, 0x26, 0xb6, 0x05, 0xd5, 0x70, 0x69, 0x2e, 0xc3, 0xc9,
	0x3a, 0x7c, 0x45, 0x5c, 0x71, 0xd8, 0x57, 0x20, 0x3d, 0xa1, 0xfb, 0x52, 0x69, 0x57, 0xa5, 0x0d,
	0xb0, 0x34, 0xa5, 0xb0, 0xa5, 0x6f, 0xd3, 0x71, 0x58, 0xec, 0x3d, 0xd4, 0xd3, 0x74, 0xdc, 0x3a,
	0xc6, 0x7b, 0xeb, 0x40, 0x94, 0xfd, 0xf3, 0x72, 0x91, 0x1b, 0x6d, 0x22, 0x04, 0x4a, 0x13, 0x3e,
	0x16, 0xb4, 0x80, 0x1e, 0x7c, 0x93, 0x26, 0xf8, 0x91, 0x50, 0x61, 0x22, 0xa7, 0x5a, 0xc6, 0x13,
	0x5a, 0x44, 0xc8, 0x75, 0x11, 0x0a, 0x87, 0xa1, 0xd1, 0x8d, 0x13, 0x5a, 0x42, 0xd4, 0x9a, 0xe4,
	0x14, 0xca, 0x72, 0xcc, 0x47, 0x82, 0x96, 0xd1, 0x9f, 0x1a, 0xe4, 0x15, 0xf8, 0xb6, 0x2f, 0x0f,
	0x32, 0xa2, 0x15, 0xc4, 0xc0, 0xba, 0xbe, 0x61, 0x1a, 0x9a, 0x8f, 0x14, 0x3d, 0x6c, 0x16, 0x4d,
	0x1a, 0xe6, 0x4d, 0xce, 0xa0, 0xa6, 0xe4, 0x68, 0xc2, 0xf5, 0x3c, 0x11, 0xb4, 0x8a, 0x21, 0x6b,
	0x07, 0xbb, 0x80, 0xe7, 0x4e, 0x69, 0x9d, 0x11, 0x97, 0x93, 0x6c, 0x7d, 0x25, 0x5c, 0x83, 0x5f,
	0x9b, 0xa3, 0x9c, 0xad, 0x8a, 0xf6, 0x76, 0x17, 0x5d, 0xd8, 0x2c, 0xfa, 0x1c, 0x80, 0x87, 0xa1,
	0x50, 0xea, 0xe1, 0xb7, 0x58, 0x2c, 0xbb, 0x52, 0x4b, 0x3d, 0xdf, 0xc5, 0x82, 0x9d, 0x6e, 0x99,
	0xf7, 0x8c, 0xdd, 0x42, 0xa3, 0x13, 0x45, 0xf7, 0x36, 0xed, 0xa7, 0x8e, 0xc0, 0x69, 0x70, 0x31,
	0xd3, 0x60, 0x76, 0x95, 0x17, 0x54, 0x86, 0xac, 0xe6, 0x98, 0x07, 0xaa, 0x56, 0x07, 0xd6, 0x6c,
	0xff, 0x2d, 0x41, 0x15, 0xfb, 0x73, 0x9f, 0xfc, 0x21, 0x5d, 0xf0, 0x9d, 0x43, 0x22, 0x74, 0xbd,
	0x34, 0xd9, 0xfb, 0x0a, 0x76, 0x21, 0x8a, 0x1d, 0x90, 0x3e, 0x5e, 0xa3, 0x73, 0x56, 0xe4, 0x65,
	0x86, 0x9d, 0xbd, 0xca, 0x60, 0x0f, 0x68, 0xd4, 0xba, 0xe0, 0x3b, 0x43, 0x74, 0x53, 0xca, 0xae,
	0x6d, 0xb0, 0x0b, 0x31, 0x22, 0x37, 0x70, 0x9c, 0xdf, 0x04, 0x72, 0xbe, 0x95, 0x6f, 0xb7, 0x64,
	0xaf, 0x5c, 0xdf, 0xde, 0x8c, 0x1d, 0xa3, 0x5b, 0xe1, 0xc6, 0x2e, 0x05, 0x7b, 0x40, 0xa3, 0x76,
	0x0b, 0x8d, 0xdc, 0x56, 0x90, 0xb3, 0x4c, 0x4f, 0x72, 0x0b, 0x13, 0xec, 0x43, 0x8d, 0xe0, 0x17,
	0x38, 0x72, 0xe7, 0x4f, 0x5e, 0xac, 0xf9, 0xb9, 0x45, 0x0b, 0x76, 0x42, 0x8a, 0x1d, 0x5c, 0x5f,
	0xc1, 0x51, 0x18, 0x8f, 0x57, 0x8c, 0x6b, 0xc0, 0x16, 0xdc, 0x19, 0xf3, 0xce, 0xfb, 0xf1, 0xac,
	0xf5, 0xf6, 0x93, 0x85, 0xa6, 0xc3, 0x61, 0x05, 0xdf, 0xef, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xdc, 0xf8, 0xa1, 0xf9, 0xea, 0x05, 0x00, 0x00,
}
