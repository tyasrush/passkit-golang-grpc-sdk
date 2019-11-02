// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/api/common_objects.proto

package io

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

// List position is used to denote behavior for list items that may already have been populated bay a parent object.  Default is to append the item to the existing list.
type ListPosition int32

const (
	ListPosition_APPEND  ListPosition = 0
	ListPosition_PREPEND ListPosition = 1
	ListPosition_REPLACE ListPosition = 2
)

var ListPosition_name = map[int32]string{
	0: "APPEND",
	1: "PREPEND",
	2: "REPLACE",
}

var ListPosition_value = map[string]int32{
	"APPEND":  0,
	"PREPEND": 1,
	"REPLACE": 2,
}

func (x ListPosition) String() string {
	return proto.EnumName(ListPosition_name, int32(x))
}

func (ListPosition) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_91ad6ceef02bf15b, []int{0}
}

// PassBundleFormat is used to specify which format(s) to provide the pass content.
type PassBundleFormat int32

const (
	// The URL to the web landing page.
	PassBundleFormat_PASS_URL PassBundleFormat = 0
	// A signed JWT that can be used in a 'Save to Google Pay' button for web use.
	PassBundleFormat_GOOGLE_JWT PassBundleFormat = 1
	// A URL that can be used in an Android app or email. Note that this is not recommended for web use.
	PassBundleFormat_GOOGLE_URL PassBundleFormat = 2
	// A HTML snippet containing an 'Save to Google Pay' button for web use.
	PassBundleFormat_GOOGLE_LINK PassBundleFormat = 4
	// The Base64 encoded bytes of the Apple Wallet pass bundle.
	PassBundleFormat_APPLE_PAY_BUNDLE PassBundleFormat = 8
)

var PassBundleFormat_name = map[int32]string{
	0: "PASS_URL",
	1: "GOOGLE_JWT",
	2: "GOOGLE_URL",
	4: "GOOGLE_LINK",
	8: "APPLE_PAY_BUNDLE",
}

var PassBundleFormat_value = map[string]int32{
	"PASS_URL":         0,
	"GOOGLE_JWT":       1,
	"GOOGLE_URL":       2,
	"GOOGLE_LINK":      4,
	"APPLE_PAY_BUNDLE": 8,
}

func (x PassBundleFormat) String() string {
	return proto.EnumName(PassBundleFormat_name, int32(x))
}

func (PassBundleFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_91ad6ceef02bf15b, []int{1}
}

// UsageType is used to indicate where a field or link is used / rendered
type UsageType int32

const (
	UsageType_NO_USAGE UsageType = 0
	// Rendered on Apple Wallet
	UsageType_USAGE_APPLE_WALLET UsageType = 1
	// Rendered on Google Pay
	UsageType_USAGE_GOOGLE_PAY UsageType = 2
	// Rendered on data collection page
	UsageType_USAGE_DATA_COLLECTION_PAGE UsageType = 4
)

var UsageType_name = map[int32]string{
	0: "NO_USAGE",
	1: "USAGE_APPLE_WALLET",
	2: "USAGE_GOOGLE_PAY",
	4: "USAGE_DATA_COLLECTION_PAGE",
}

var UsageType_value = map[string]int32{
	"NO_USAGE":                   0,
	"USAGE_APPLE_WALLET":         1,
	"USAGE_GOOGLE_PAY":           2,
	"USAGE_DATA_COLLECTION_PAGE": 4,
}

func (x UsageType) String() string {
	return proto.EnumName(UsageType_name, int32(x))
}

func (UsageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_91ad6ceef02bf15b, []int{2}
}

// An Id is used to access an unique object or record. E.g. a PassKit pass id.
type Id struct {
	// The unique identifier to an object or record.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_91ad6ceef02bf15b, []int{0}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Url struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Url) Reset()         { *m = Url{} }
func (m *Url) String() string { return proto.CompactTextString(m) }
func (*Url) ProtoMessage()    {}
func (*Url) Descriptor() ([]byte, []int) {
	return fileDescriptor_91ad6ceef02bf15b, []int{1}
}

func (m *Url) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Url.Unmarshal(m, b)
}
func (m *Url) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Url.Marshal(b, m, deterministic)
}
func (m *Url) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Url.Merge(m, src)
}
func (m *Url) XXX_Size() int {
	return xxx_messageInfo_Url.Size(m)
}
func (m *Url) XXX_DiscardUnknown() {
	xxx_messageInfo_Url.DiscardUnknown(m)
}

var xxx_messageInfo_Url proto.InternalMessageInfo

func (m *Url) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Count struct {
	Total                int32    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Count) Reset()         { *m = Count{} }
func (m *Count) String() string { return proto.CompactTextString(m) }
func (*Count) ProtoMessage()    {}
func (*Count) Descriptor() ([]byte, []int) {
	return fileDescriptor_91ad6ceef02bf15b, []int{2}
}

func (m *Count) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Count.Unmarshal(m, b)
}
func (m *Count) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Count.Marshal(b, m, deterministic)
}
func (m *Count) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Count.Merge(m, src)
}
func (m *Count) XXX_Size() int {
	return xxx_messageInfo_Count.Size(m)
}
func (m *Count) XXX_DiscardUnknown() {
	xxx_messageInfo_Count.DiscardUnknown(m)
}

var xxx_messageInfo_Count proto.InternalMessageInfo

func (m *Count) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type FileBytes struct {
	FileBytes            []byte   `protobuf:"bytes,1,opt,name=fileBytes,proto3" json:"fileBytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileBytes) Reset()         { *m = FileBytes{} }
func (m *FileBytes) String() string { return proto.CompactTextString(m) }
func (*FileBytes) ProtoMessage()    {}
func (*FileBytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_91ad6ceef02bf15b, []int{3}
}

func (m *FileBytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileBytes.Unmarshal(m, b)
}
func (m *FileBytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileBytes.Marshal(b, m, deterministic)
}
func (m *FileBytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileBytes.Merge(m, src)
}
func (m *FileBytes) XXX_Size() int {
	return xxx_messageInfo_FileBytes.Size(m)
}
func (m *FileBytes) XXX_DiscardUnknown() {
	xxx_messageInfo_FileBytes.DiscardUnknown(m)
}

var xxx_messageInfo_FileBytes proto.InternalMessageInfo

func (m *FileBytes) GetFileBytes() []byte {
	if m != nil {
		return m.FileBytes
	}
	return nil
}

type Boolean struct {
	// @tag: json:"response"
	Response             bool     `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Boolean) Reset()         { *m = Boolean{} }
func (m *Boolean) String() string { return proto.CompactTextString(m) }
func (*Boolean) ProtoMessage()    {}
func (*Boolean) Descriptor() ([]byte, []int) {
	return fileDescriptor_91ad6ceef02bf15b, []int{4}
}

func (m *Boolean) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Boolean.Unmarshal(m, b)
}
func (m *Boolean) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Boolean.Marshal(b, m, deterministic)
}
func (m *Boolean) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Boolean.Merge(m, src)
}
func (m *Boolean) XXX_Size() int {
	return xxx_messageInfo_Boolean.Size(m)
}
func (m *Boolean) XXX_DiscardUnknown() {
	xxx_messageInfo_Boolean.DiscardUnknown(m)
}

var xxx_messageInfo_Boolean proto.InternalMessageInfo

func (m *Boolean) GetResponse() bool {
	if m != nil {
		return m.Response
	}
	return false
}

// Repeated field of dynamically typed values (including string, bool, option, syntax)
type Strings struct {
	Response             []string `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Strings) Reset()         { *m = Strings{} }
func (m *Strings) String() string { return proto.CompactTextString(m) }
func (*Strings) ProtoMessage()    {}
func (*Strings) Descriptor() ([]byte, []int) {
	return fileDescriptor_91ad6ceef02bf15b, []int{5}
}

func (m *Strings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Strings.Unmarshal(m, b)
}
func (m *Strings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Strings.Marshal(b, m, deterministic)
}
func (m *Strings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Strings.Merge(m, src)
}
func (m *Strings) XXX_Size() int {
	return xxx_messageInfo_Strings.Size(m)
}
func (m *Strings) XXX_DiscardUnknown() {
	xxx_messageInfo_Strings.DiscardUnknown(m)
}

var xxx_messageInfo_Strings proto.InternalMessageInfo

func (m *Strings) GetResponse() []string {
	if m != nil {
		return m.Response
	}
	return nil
}

// A fixed date without timezone information.
type Date struct {
	// Year.
	Year int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// Month.
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	// Date.
	Day                  int32    `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Date) Reset()         { *m = Date{} }
func (m *Date) String() string { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()    {}
func (*Date) Descriptor() ([]byte, []int) {
	return fileDescriptor_91ad6ceef02bf15b, []int{6}
}

func (m *Date) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Date.Unmarshal(m, b)
}
func (m *Date) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Date.Marshal(b, m, deterministic)
}
func (m *Date) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Date.Merge(m, src)
}
func (m *Date) XXX_Size() int {
	return xxx_messageInfo_Date.Size(m)
}
func (m *Date) XXX_DiscardUnknown() {
	xxx_messageInfo_Date.DiscardUnknown(m)
}

var xxx_messageInfo_Date proto.InternalMessageInfo

func (m *Date) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Date) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Date) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

// Local Date are specified in ISO 8601 format date without a time. The date should be local to where the pass will be at the time of use.
type LocalDate struct {
	// ISO 8601 format date without a time. E.g. 2019-08-07.
	DateTime             string   `protobuf:"bytes,1,opt,name=dateTime,proto3" json:"dateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalDate) Reset()         { *m = LocalDate{} }
func (m *LocalDate) String() string { return proto.CompactTextString(m) }
func (*LocalDate) ProtoMessage()    {}
func (*LocalDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_91ad6ceef02bf15b, []int{7}
}

func (m *LocalDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalDate.Unmarshal(m, b)
}
func (m *LocalDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalDate.Marshal(b, m, deterministic)
}
func (m *LocalDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalDate.Merge(m, src)
}
func (m *LocalDate) XXX_Size() int {
	return xxx_messageInfo_LocalDate.Size(m)
}
func (m *LocalDate) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalDate.DiscardUnknown(m)
}

var xxx_messageInfo_LocalDate proto.InternalMessageInfo

func (m *LocalDate) GetDateTime() string {
	if m != nil {
		return m.DateTime
	}
	return ""
}

// Local Date Times are specified in ISO 8601 extended format date/time without a timezone offset. The time should be local to where the pass will be at the time of use.
type LocalDateTime struct {
	// ISO 8601 extended format date/time without an offset E.g. 2019-08-07T18:00:00. Time can optionally be stated to millisecond precision. E.g. 2019-08-07T18:00:00.123.
	DateTime             string   `protobuf:"bytes,1,opt,name=dateTime,proto3" json:"dateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalDateTime) Reset()         { *m = LocalDateTime{} }
func (m *LocalDateTime) String() string { return proto.CompactTextString(m) }
func (*LocalDateTime) ProtoMessage()    {}
func (*LocalDateTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_91ad6ceef02bf15b, []int{8}
}

func (m *LocalDateTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalDateTime.Unmarshal(m, b)
}
func (m *LocalDateTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalDateTime.Marshal(b, m, deterministic)
}
func (m *LocalDateTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalDateTime.Merge(m, src)
}
func (m *LocalDateTime) XXX_Size() int {
	return xxx_messageInfo_LocalDateTime.Size(m)
}
func (m *LocalDateTime) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalDateTime.DiscardUnknown(m)
}

var xxx_messageInfo_LocalDateTime proto.InternalMessageInfo

func (m *LocalDateTime) GetDateTime() string {
	if m != nil {
		return m.DateTime
	}
	return ""
}

// A Pass Bundle contains a landing page URL and/or a compiled pass in multiple formats
type PassBundle struct {
	// The PassKit Id that uniquely references the pass.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The URL to the web landing page.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// A signed JWT that can be used in a 'Save to Google Pay' button for web use.
	GooglePayJWT string `protobuf:"bytes,3,opt,name=googlePayJWT,proto3" json:"googlePayJWT,omitempty"`
	// A HTML snippet containing an 'Save to Google Pay' button for web use.
	GooglePayHTML string `protobuf:"bytes,4,opt,name=googlePayHTML,proto3" json:"googlePayHTML,omitempty"`
	// A URL that can be used in an Android app or email. Note that this is not recommended for web use.
	GooglePayURL string `protobuf:"bytes,5,opt,name=googlePayURL,proto3" json:"googlePayURL,omitempty"`
	// The Base64 encoded bytes of the Apple Wallet pass.
	ApplePayBytes        []byte   `protobuf:"bytes,6,opt,name=applePayBytes,proto3" json:"applePayBytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassBundle) Reset()         { *m = PassBundle{} }
func (m *PassBundle) String() string { return proto.CompactTextString(m) }
func (*PassBundle) ProtoMessage()    {}
func (*PassBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_91ad6ceef02bf15b, []int{9}
}

func (m *PassBundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassBundle.Unmarshal(m, b)
}
func (m *PassBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassBundle.Marshal(b, m, deterministic)
}
func (m *PassBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassBundle.Merge(m, src)
}
func (m *PassBundle) XXX_Size() int {
	return xxx_messageInfo_PassBundle.Size(m)
}
func (m *PassBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_PassBundle.DiscardUnknown(m)
}

var xxx_messageInfo_PassBundle proto.InternalMessageInfo

func (m *PassBundle) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PassBundle) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PassBundle) GetGooglePayJWT() string {
	if m != nil {
		return m.GooglePayJWT
	}
	return ""
}

func (m *PassBundle) GetGooglePayHTML() string {
	if m != nil {
		return m.GooglePayHTML
	}
	return ""
}

func (m *PassBundle) GetGooglePayURL() string {
	if m != nil {
		return m.GooglePayURL
	}
	return ""
}

func (m *PassBundle) GetApplePayBytes() []byte {
	if m != nil {
		return m.ApplePayBytes
	}
	return nil
}

type PassBundleRequest struct {
	// The PassKit pass id that uniquely identifies the pass.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The format(s) of the pass required.  If omitted, the URL to the web landing page will be returned.
	Format               []PassBundleFormat `protobuf:"varint,2,rep,packed,name=format,proto3,enum=io.PassBundleFormat" json:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PassBundleRequest) Reset()         { *m = PassBundleRequest{} }
func (m *PassBundleRequest) String() string { return proto.CompactTextString(m) }
func (*PassBundleRequest) ProtoMessage()    {}
func (*PassBundleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91ad6ceef02bf15b, []int{10}
}

func (m *PassBundleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassBundleRequest.Unmarshal(m, b)
}
func (m *PassBundleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassBundleRequest.Marshal(b, m, deterministic)
}
func (m *PassBundleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassBundleRequest.Merge(m, src)
}
func (m *PassBundleRequest) XXX_Size() int {
	return xxx_messageInfo_PassBundleRequest.Size(m)
}
func (m *PassBundleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PassBundleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PassBundleRequest proto.InternalMessageInfo

func (m *PassBundleRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PassBundleRequest) GetFormat() []PassBundleFormat {
	if m != nil {
		return m.Format
	}
	return nil
}

func init() {
	proto.RegisterEnum("io.ListPosition", ListPosition_name, ListPosition_value)
	proto.RegisterEnum("io.PassBundleFormat", PassBundleFormat_name, PassBundleFormat_value)
	proto.RegisterEnum("io.UsageType", UsageType_name, UsageType_value)
	proto.RegisterType((*Id)(nil), "io.Id")
	proto.RegisterType((*Url)(nil), "io.Url")
	proto.RegisterType((*Count)(nil), "io.Count")
	proto.RegisterType((*FileBytes)(nil), "io.FileBytes")
	proto.RegisterType((*Boolean)(nil), "io.Boolean")
	proto.RegisterType((*Strings)(nil), "io.Strings")
	proto.RegisterType((*Date)(nil), "io.Date")
	proto.RegisterType((*LocalDate)(nil), "io.LocalDate")
	proto.RegisterType((*LocalDateTime)(nil), "io.LocalDateTime")
	proto.RegisterType((*PassBundle)(nil), "io.PassBundle")
	proto.RegisterType((*PassBundleRequest)(nil), "io.PassBundleRequest")
}

func init() { proto.RegisterFile("io/api/common_objects.proto", fileDescriptor_91ad6ceef02bf15b) }

var fileDescriptor_91ad6ceef02bf15b = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x6e, 0xdb, 0x46,
	0x10, 0x35, 0x29, 0xd9, 0x91, 0x26, 0x8e, 0xc3, 0x2e, 0x84, 0x56, 0x70, 0x2f, 0x08, 0x88, 0x14,
	0xa0, 0x0d, 0x53, 0x17, 0x3b, 0x2d, 0xe0, 0xa2, 0x40, 0x40, 0xda, 0x8c, 0xeb, 0x98, 0x95, 0x58,
	0x9a, 0x82, 0x9b, 0x0a, 0x2a, 0xb1, 0x96, 0x36, 0xcc, 0x36, 0x14, 0x97, 0xe5, 0xae, 0xd0, 0x0a,
	0x75, 0x7e, 0x20, 0x8f, 0x41, 0xbf, 0xa1, 0x0f, 0xfd, 0x8c, 0x7e, 0x59, 0xb1, 0x4b, 0x59, 0xb6,
	0xe5, 0x3e, 0xf5, 0x6d, 0xce, 0x99, 0xb3, 0xe7, 0xcc, 0x08, 0x23, 0xc2, 0xa7, 0x94, 0xb5, 0x71,
	0x4e, 0xdb, 0x63, 0x36, 0x9d, 0xb2, 0x2c, 0x66, 0x97, 0xbf, 0x90, 0xb1, 0xe0, 0xad, 0xbc, 0x60,
	0x82, 0x21, 0x9d, 0xb2, 0xed, 0x3d, 0x55, 0x8e, 0xed, 0x84, 0x64, 0x36, 0xff, 0x0d, 0x27, 0x09,
	0x29, 0xda, 0x2c, 0x17, 0x94, 0x65, 0xbc, 0x8d, 0xb3, 0x8c, 0x09, 0xac, 0xea, 0xf2, 0x85, 0xd9,
	0x00, 0xfd, 0x74, 0x82, 0xb6, 0x40, 0xa7, 0x93, 0xa6, 0xf6, 0x44, 0xb3, 0xea, 0xa1, 0x4e, 0x27,
	0xe6, 0x27, 0x50, 0x19, 0x14, 0x29, 0x32, 0xa0, 0x32, 0x2b, 0xd2, 0x05, 0x2f, 0x4b, 0xf3, 0x73,
	0x58, 0x3f, 0x62, 0xb3, 0x4c, 0xa0, 0x06, 0xac, 0x0b, 0x26, 0x70, 0xd9, 0x5c, 0x0f, 0x4b, 0x60,
	0xee, 0x40, 0xfd, 0x05, 0x4d, 0x89, 0x3b, 0x17, 0x84, 0xa3, 0xcf, 0xa0, 0xfe, 0xfa, 0x1a, 0x28,
	0xd9, 0x66, 0x78, 0x43, 0x98, 0x5f, 0xc2, 0x03, 0x97, 0xb1, 0x94, 0xe0, 0x0c, 0x6d, 0x43, 0xad,
	0x20, 0x3c, 0x67, 0x19, 0x27, 0x4a, 0x57, 0x0b, 0x97, 0x58, 0xca, 0xce, 0x45, 0x41, 0xb3, 0x84,
	0xaf, 0xc8, 0x2a, 0x56, 0xfd, 0x96, 0xcc, 0x85, 0xea, 0x31, 0x16, 0x04, 0x21, 0xa8, 0xce, 0x09,
	0x2e, 0x16, 0x53, 0xa9, 0x5a, 0x8e, 0x3a, 0x65, 0x99, 0x78, 0xd3, 0xd4, 0xcb, 0x51, 0x15, 0x90,
	0xbb, 0x4d, 0xf0, 0xbc, 0x59, 0x51, 0x9c, 0x2c, 0x4d, 0x02, 0x75, 0x9f, 0x8d, 0x71, 0xaa, 0x8c,
	0x7e, 0x84, 0xda, 0x04, 0x0b, 0x12, 0xd1, 0x69, 0x39, 0x53, 0xdd, 0xfd, 0xf6, 0x83, 0x73, 0xf8,
	0x5e, 0xfb, 0xfa, 0x67, 0x6b, 0xd8, 0xb1, 0x0f, 0x47, 0x7f, 0x3c, 0x7b, 0xb7, 0x63, 0x3f, 0xb7,
	0xba, 0xc3, 0x8e, 0xbd, 0x3f, 0xba, 0xea, 0x0c, 0xbb, 0xf6, 0xe1, 0x48, 0x12, 0x07, 0xc3, 0x4e,
	0xf7, 0x1a, 0x5f, 0x0d, 0xbb, 0xfb, 0x23, 0xa5, 0xde, 0x09, 0x97, 0x6e, 0xe6, 0x5f, 0x1a, 0x3c,
	0x5a, 0xe6, 0x48, 0x06, 0xfd, 0xa9, 0xdd, 0x0b, 0xfb, 0xfd, 0x83, 0x33, 0x7b, 0xaf, 0x15, 0xff,
	0x2f, 0x2c, 0xb2, 0xf6, 0x87, 0x1d, 0xfb, 0x60, 0x74, 0x25, 0xdb, 0x25, 0xf5, 0x8d, 0x34, 0xfa,
	0xea, 0xbf, 0x80, 0xd5, 0x2a, 0x23, 0xba, 0x7b, 0x07, 0xef, 0x76, 0x9e, 0x3f, 0xbd, 0x35, 0xe8,
	0x3f, 0x1a, 0x40, 0x80, 0x39, 0x77, 0x67, 0xd9, 0x24, 0x25, 0xab, 0x37, 0x72, 0x7d, 0x1c, 0xfa,
	0xf2, 0x38, 0x90, 0x09, 0x9b, 0x09, 0x63, 0x49, 0x4a, 0x02, 0x3c, 0x7f, 0x79, 0x11, 0xa9, 0xdf,
	0xb6, 0x1e, 0xde, 0xe1, 0xd0, 0x53, 0x78, 0xb4, 0xc4, 0xdf, 0x45, 0xdf, 0xfb, 0xcd, 0xaa, 0x12,
	0xdd, 0x25, 0xef, 0x38, 0x0d, 0x42, 0xbf, 0xb9, 0xbe, 0xe2, 0x34, 0x08, 0x7d, 0xe9, 0x84, 0xf3,
	0x5c, 0xc1, 0xf2, 0xc4, 0x36, 0xd4, 0x89, 0xdd, 0x25, 0xcd, 0x1f, 0xe0, 0xa3, 0x9b, 0x1d, 0x42,
	0xf2, 0xeb, 0x8c, 0x70, 0x71, 0x6f, 0x95, 0x3d, 0xd8, 0x78, 0xcd, 0x8a, 0x29, 0x16, 0x4d, 0xfd,
	0x49, 0xc5, 0xda, 0xda, 0x6f, 0xb4, 0x28, 0x6b, 0xdd, 0x3c, 0x7b, 0xa1, 0x7a, 0xe1, 0x42, 0xb3,
	0xfb, 0x0c, 0x36, 0x7d, 0xca, 0x45, 0xc0, 0x38, 0x95, 0xff, 0x24, 0x04, 0xb0, 0xe1, 0x04, 0x81,
	0xd7, 0x3b, 0x36, 0xd6, 0xd0, 0x43, 0x78, 0x10, 0x84, 0x9e, 0x02, 0x9a, 0x04, 0xa1, 0x17, 0xf8,
	0xce, 0x91, 0x67, 0xe8, 0xbb, 0x09, 0x18, 0xab, 0x8e, 0x68, 0x13, 0x6a, 0x81, 0x73, 0x7e, 0x1e,
	0x0f, 0x42, 0xdf, 0x58, 0x43, 0x5b, 0x00, 0x27, 0xfd, 0xfe, 0x89, 0xef, 0xc5, 0x2f, 0x2f, 0x22,
	0x43, 0xbb, 0x85, 0x65, 0x5f, 0x47, 0x8f, 0xe1, 0xe1, 0x02, 0xfb, 0xa7, 0xbd, 0x33, 0xa3, 0x8a,
	0x1a, 0x60, 0x38, 0x41, 0xe0, 0x7b, 0x71, 0xe0, 0xbc, 0x8a, 0xdd, 0x41, 0xef, 0xd8, 0xf7, 0x8c,
	0xda, 0x6e, 0x02, 0xf5, 0x01, 0xc7, 0x09, 0x89, 0xe6, 0x39, 0x91, 0x09, 0xbd, 0x7e, 0x3c, 0x38,
	0x77, 0x4e, 0x3c, 0x63, 0x0d, 0x7d, 0x0c, 0x48, 0x95, 0x71, 0xf9, 0xec, 0xc2, 0xf1, 0x7d, 0x4f,
	0x26, 0x35, 0xc0, 0x28, 0xf9, 0x85, 0x7f, 0xe0, 0xbc, 0x32, 0x74, 0xf4, 0x05, 0x6c, 0x97, 0xec,
	0xb1, 0x13, 0x39, 0xf1, 0x51, 0xdf, 0xf7, 0xbd, 0xa3, 0xe8, 0xb4, 0xdf, 0x8b, 0x03, 0xe9, 0x56,
	0x75, 0x5d, 0x78, 0x4c, 0x59, 0x2b, 0xc7, 0x9c, 0xbf, 0xa5, 0xa2, 0x15, 0x9c, 0x51, 0xf6, 0x93,
	0x95, 0x50, 0xf1, 0x66, 0x76, 0xd9, 0x1a, 0xb3, 0x69, 0x5b, 0x6e, 0x7b, 0x46, 0x45, 0x7b, 0x21,
	0xb0, 0x13, 0x96, 0xe2, 0x2c, 0xb1, 0xf9, 0xe4, 0x6d, 0x9b, 0xb2, 0xbf, 0xf5, 0x6a, 0x70, 0x76,
	0xca, 0x2e, 0x37, 0xd4, 0x57, 0xe8, 0xe0, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x46, 0x62,
	0x39, 0xd6, 0x04, 0x00, 0x00,
}
