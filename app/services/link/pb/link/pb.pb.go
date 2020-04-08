// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/link/pb.proto

package link

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type LinkDomainServiceErrorCode int32

const (
	LinkDomainServiceErrorCode_Unknown        LinkDomainServiceErrorCode = 0
	LinkDomainServiceErrorCode_Internal       LinkDomainServiceErrorCode = 5000
	LinkDomainServiceErrorCode_InvalidRequest LinkDomainServiceErrorCode = 4000
)

var LinkDomainServiceErrorCode_name = map[int32]string{
	0:    "Unknown",
	5000: "Internal",
	4000: "InvalidRequest",
}

var LinkDomainServiceErrorCode_value = map[string]int32{
	"Unknown":        0,
	"Internal":       5000,
	"InvalidRequest": 4000,
}

func (x LinkDomainServiceErrorCode) String() string {
	return proto.EnumName(LinkDomainServiceErrorCode_name, int32(x))
}

func (LinkDomainServiceErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17f110ab7087989c, []int{0}
}

type CreateLinkRequest struct {
	Link                 *LinkEntity `protobuf:"bytes,1,opt,name=Link,proto3" json:"Link,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateLinkRequest) Reset()         { *m = CreateLinkRequest{} }
func (m *CreateLinkRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLinkRequest) ProtoMessage()    {}
func (*CreateLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17f110ab7087989c, []int{0}
}

func (m *CreateLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLinkRequest.Unmarshal(m, b)
}
func (m *CreateLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLinkRequest.Marshal(b, m, deterministic)
}
func (m *CreateLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLinkRequest.Merge(m, src)
}
func (m *CreateLinkRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLinkRequest.Size(m)
}
func (m *CreateLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLinkRequest proto.InternalMessageInfo

func (m *CreateLinkRequest) GetLink() *LinkEntity {
	if m != nil {
		return m.Link
	}
	return nil
}

type CreateLinkResponse struct {
	Link                 *LinkEntity `protobuf:"bytes,1,opt,name=Link,proto3" json:"Link,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateLinkResponse) Reset()         { *m = CreateLinkResponse{} }
func (m *CreateLinkResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLinkResponse) ProtoMessage()    {}
func (*CreateLinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17f110ab7087989c, []int{1}
}

func (m *CreateLinkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLinkResponse.Unmarshal(m, b)
}
func (m *CreateLinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLinkResponse.Marshal(b, m, deterministic)
}
func (m *CreateLinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLinkResponse.Merge(m, src)
}
func (m *CreateLinkResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLinkResponse.Size(m)
}
func (m *CreateLinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLinkResponse proto.InternalMessageInfo

func (m *CreateLinkResponse) GetLink() *LinkEntity {
	if m != nil {
		return m.Link
	}
	return nil
}

type DeleteLinkRequest struct {
	LinkID               string   `protobuf:"bytes,1,opt,name=LinkID,proto3" json:"LinkID,omitempty"`
	UserID               string   `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteLinkRequest) Reset()         { *m = DeleteLinkRequest{} }
func (m *DeleteLinkRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteLinkRequest) ProtoMessage()    {}
func (*DeleteLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17f110ab7087989c, []int{2}
}

func (m *DeleteLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteLinkRequest.Unmarshal(m, b)
}
func (m *DeleteLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteLinkRequest.Marshal(b, m, deterministic)
}
func (m *DeleteLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteLinkRequest.Merge(m, src)
}
func (m *DeleteLinkRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteLinkRequest.Size(m)
}
func (m *DeleteLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteLinkRequest proto.InternalMessageInfo

func (m *DeleteLinkRequest) GetLinkID() string {
	if m != nil {
		return m.LinkID
	}
	return ""
}

func (m *DeleteLinkRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type DeleteLinkResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteLinkResponse) Reset()         { *m = DeleteLinkResponse{} }
func (m *DeleteLinkResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteLinkResponse) ProtoMessage()    {}
func (*DeleteLinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17f110ab7087989c, []int{3}
}

func (m *DeleteLinkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteLinkResponse.Unmarshal(m, b)
}
func (m *DeleteLinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteLinkResponse.Marshal(b, m, deterministic)
}
func (m *DeleteLinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteLinkResponse.Merge(m, src)
}
func (m *DeleteLinkResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteLinkResponse.Size(m)
}
func (m *DeleteLinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteLinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteLinkResponse proto.InternalMessageInfo

type FindLinksRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Limit                uint64   `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset               uint64   `protobuf:"varint,3,opt,name=Offset,proto3" json:"Offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindLinksRequest) Reset()         { *m = FindLinksRequest{} }
func (m *FindLinksRequest) String() string { return proto.CompactTextString(m) }
func (*FindLinksRequest) ProtoMessage()    {}
func (*FindLinksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17f110ab7087989c, []int{4}
}

func (m *FindLinksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindLinksRequest.Unmarshal(m, b)
}
func (m *FindLinksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindLinksRequest.Marshal(b, m, deterministic)
}
func (m *FindLinksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindLinksRequest.Merge(m, src)
}
func (m *FindLinksRequest) XXX_Size() int {
	return xxx_messageInfo_FindLinksRequest.Size(m)
}
func (m *FindLinksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindLinksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindLinksRequest proto.InternalMessageInfo

func (m *FindLinksRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *FindLinksRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *FindLinksRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type FindLinksResponse struct {
	Items                []*LinkEntity `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	TotalLinksNumber     uint64        `protobuf:"varint,2,opt,name=TotalLinksNumber,proto3" json:"TotalLinksNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FindLinksResponse) Reset()         { *m = FindLinksResponse{} }
func (m *FindLinksResponse) String() string { return proto.CompactTextString(m) }
func (*FindLinksResponse) ProtoMessage()    {}
func (*FindLinksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17f110ab7087989c, []int{5}
}

func (m *FindLinksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindLinksResponse.Unmarshal(m, b)
}
func (m *FindLinksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindLinksResponse.Marshal(b, m, deterministic)
}
func (m *FindLinksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindLinksResponse.Merge(m, src)
}
func (m *FindLinksResponse) XXX_Size() int {
	return xxx_messageInfo_FindLinksResponse.Size(m)
}
func (m *FindLinksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindLinksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindLinksResponse proto.InternalMessageInfo

func (m *FindLinksResponse) GetItems() []*LinkEntity {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *FindLinksResponse) GetTotalLinksNumber() uint64 {
	if m != nil {
		return m.TotalLinksNumber
	}
	return 0
}

type LinkEntity struct {
	LinkID               string   `protobuf:"bytes,1,opt,name=LinkID,proto3" json:"LinkID,omitempty"`
	UserID               string   `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=Address,proto3" json:"Address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkEntity) Reset()         { *m = LinkEntity{} }
func (m *LinkEntity) String() string { return proto.CompactTextString(m) }
func (*LinkEntity) ProtoMessage()    {}
func (*LinkEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_17f110ab7087989c, []int{6}
}

func (m *LinkEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkEntity.Unmarshal(m, b)
}
func (m *LinkEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkEntity.Marshal(b, m, deterministic)
}
func (m *LinkEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkEntity.Merge(m, src)
}
func (m *LinkEntity) XXX_Size() int {
	return xxx_messageInfo_LinkEntity.Size(m)
}
func (m *LinkEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkEntity.DiscardUnknown(m)
}

var xxx_messageInfo_LinkEntity proto.InternalMessageInfo

func (m *LinkEntity) GetLinkID() string {
	if m != nil {
		return m.LinkID
	}
	return ""
}

func (m *LinkEntity) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *LinkEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LinkEntity) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterEnum("link.LinkDomainServiceErrorCode", LinkDomainServiceErrorCode_name, LinkDomainServiceErrorCode_value)
	proto.RegisterType((*CreateLinkRequest)(nil), "link.CreateLinkRequest")
	proto.RegisterType((*CreateLinkResponse)(nil), "link.CreateLinkResponse")
	proto.RegisterType((*DeleteLinkRequest)(nil), "link.DeleteLinkRequest")
	proto.RegisterType((*DeleteLinkResponse)(nil), "link.DeleteLinkResponse")
	proto.RegisterType((*FindLinksRequest)(nil), "link.FindLinksRequest")
	proto.RegisterType((*FindLinksResponse)(nil), "link.FindLinksResponse")
	proto.RegisterType((*LinkEntity)(nil), "link.LinkEntity")
}

func init() {
	proto.RegisterFile("pb/link/pb.proto", fileDescriptor_17f110ab7087989c)
}

var fileDescriptor_17f110ab7087989c = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x5b, 0x27, 0x69, 0xa7, 0x02, 0x39, 0x43, 0x95, 0x58, 0xe1, 0xd0, 0x62, 0x21, 0x40,
	0x3d, 0x24, 0x52, 0xb9, 0x71, 0x40, 0xaa, 0x93, 0x22, 0x59, 0x2a, 0x45, 0x0a, 0xe4, 0xc2, 0xcd,
	0xc1, 0x53, 0xb4, 0x8a, 0xbd, 0x1b, 0x76, 0xb7, 0x45, 0x48, 0x7c, 0x00, 0xe2, 0x2b, 0xf8, 0x2d,
	0x7e, 0xa5, 0x27, 0xb4, 0xeb, 0x75, 0xe2, 0xd4, 0x11, 0x37, 0xcf, 0x9b, 0x99, 0x37, 0xe3, 0x37,
	0x6f, 0x21, 0x58, 0x2d, 0xc6, 0x39, 0xe3, 0xcb, 0xf1, 0x6a, 0x31, 0x5a, 0x49, 0xa1, 0x05, 0xfa,
	0x26, 0x1c, 0x0e, 0xee, 0xd2, 0x9c, 0x65, 0xa9, 0xa6, 0x71, 0xf5, 0x51, 0xa6, 0xa3, 0x09, 0xf4,
	0x26, 0x92, 0x52, 0x4d, 0x57, 0x8c, 0x2f, 0x67, 0xf4, 0xed, 0x96, 0x94, 0xc6, 0x11, 0xf8, 0x26,
	0x0c, 0xbd, 0x53, 0xef, 0xd5, 0xd1, 0x79, 0x30, 0x32, 0x14, 0x23, 0x83, 0x5c, 0x72, 0xcd, 0xf4,
	0x8f, 0xf8, 0xe0, 0x3e, 0x6e, 0xff, 0xf6, 0xf6, 0x02, 0x6f, 0x66, 0xeb, 0xa2, 0x37, 0x80, 0x75,
	0x12, 0xb5, 0x12, 0x5c, 0x11, 0x3e, 0xff, 0x3f, 0x8b, 0xeb, 0x9d, 0x43, 0x6f, 0x4a, 0x39, 0x6d,
	0x2f, 0x70, 0x02, 0x1d, 0x13, 0x26, 0x53, 0xdb, 0x7c, 0x18, 0x77, 0xef, 0x63, 0x5f, 0x9a, 0x79,
	0x0e, 0x36, 0x05, 0x73, 0x45, 0x32, 0x99, 0x86, 0x7b, 0x0f, 0x0a, 0x4a, 0x38, 0x3a, 0x06, 0xac,
	0xd3, 0x96, 0x2b, 0x45, 0x29, 0x04, 0xef, 0x18, 0xcf, 0x0c, 0xa6, 0x6a, 0xb3, 0x1c, 0x95, 0xb7,
	0x93, 0x0a, 0x8f, 0xa1, 0x7d, 0xc5, 0x0a, 0xa6, 0xed, 0x28, 0x7f, 0x56, 0x06, 0xd8, 0x87, 0xce,
	0x87, 0x9b, 0x1b, 0x45, 0x3a, 0xdc, 0xb7, 0xb0, 0x8b, 0xa2, 0xaf, 0xd0, 0xab, 0x8d, 0x70, 0x52,
	0xbc, 0x80, 0x76, 0xa2, 0xa9, 0x50, 0xa1, 0x77, 0xba, 0xbf, 0x53, 0x8b, 0x32, 0x8d, 0x67, 0x10,
	0x7c, 0x12, 0x3a, 0xcd, 0x6d, 0xf7, 0xf5, 0x6d, 0xb1, 0x20, 0xe9, 0xa6, 0x36, 0xf0, 0xe8, 0x27,
	0xc0, 0x86, 0xc0, 0xac, 0x53, 0x57, 0x6c, 0x2d, 0x54, 0x7f, 0x5b, 0xa8, 0xf5, 0x4f, 0x3d, 0x05,
	0xff, 0x3a, 0x2d, 0xc8, 0x2e, 0x5f, 0xfb, 0x67, 0x0b, 0xe2, 0x33, 0xe8, 0x5e, 0x64, 0x99, 0x24,
	0xa5, 0x42, 0x7f, 0x3b, 0x5f, 0xe1, 0x67, 0xef, 0x61, 0x68, 0x26, 0x4c, 0x45, 0x91, 0x32, 0xfe,
	0x91, 0xe4, 0x1d, 0xfb, 0x42, 0x97, 0x52, 0x0a, 0x39, 0x11, 0x19, 0xe1, 0x11, 0x74, 0xe7, 0x7c,
	0xc9, 0xc5, 0x77, 0x1e, 0xb4, 0xf0, 0x11, 0x1c, 0x24, 0x5c, 0x93, 0xe4, 0x69, 0x1e, 0xfc, 0x7a,
	0x89, 0x4f, 0xe0, 0x71, 0xc2, 0xad, 0x0b, 0xdd, 0x05, 0x82, 0x3f, 0x27, 0xe7, 0x7f, 0x3d, 0xe8,
	0x35, 0xf8, 0xf0, 0x02, 0x60, 0xe3, 0x2b, 0x1c, 0x94, 0xaa, 0x35, 0xec, 0x3a, 0x0c, 0x9b, 0x09,
	0x77, 0xef, 0x96, 0xa1, 0xd8, 0xf8, 0xa0, 0xa2, 0x68, 0x18, 0xae, 0xa2, 0xd8, 0x61, 0x99, 0x16,
	0xbe, 0x85, 0xc3, 0xf5, 0x45, 0xb1, 0x5f, 0x16, 0x3e, 0x74, 0xd1, 0x70, 0xd0, 0xc0, 0xab, 0xfe,
	0xb8, 0xf3, 0xd9, 0xbe, 0xc1, 0x45, 0xc7, 0xbe, 0xb8, 0xd7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x67, 0x14, 0x41, 0x9a, 0xa4, 0x03, 0x00, 0x00,
}
