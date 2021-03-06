// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/common/common_objects.proto

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
	return fileDescriptor_0ded2dfbe8bca787, []int{0}
}

// PassBundleFormat is used to specify which format(s) to provide the pass content.
type PassBundleFormat int32

const (
	// The URL to the web landing page.
	PassBundleFormat_PASS_URL PassBundleFormat = 0
	// A URL that can be used in an Android app or email. Note that this is not recommended for web use.
	PassBundleFormat_GOOGLE_URL PassBundleFormat = 2
	// The Base64 encoded bytes of the Apple Wallet pass bundle.
	PassBundleFormat_APPLE_PASS_BUNDLE PassBundleFormat = 8
	// The URL to a landing page for multiple passes containing a common identifier.
	PassBundleFormat_MULTI_LINK PassBundleFormat = 16
)

var PassBundleFormat_name = map[int32]string{
	0:  "PASS_URL",
	2:  "GOOGLE_URL",
	8:  "APPLE_PASS_BUNDLE",
	16: "MULTI_LINK",
}

var PassBundleFormat_value = map[string]int32{
	"PASS_URL":          0,
	"GOOGLE_URL":        2,
	"APPLE_PASS_BUNDLE": 8,
	"MULTI_LINK":        16,
}

func (x PassBundleFormat) String() string {
	return proto.EnumName(PassBundleFormat_name, int32(x))
}

func (PassBundleFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0ded2dfbe8bca787, []int{1}
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
	return fileDescriptor_0ded2dfbe8bca787, []int{2}
}

// An Id is used to access an unique object or record. Eg. a PassKit pass id.
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
	return fileDescriptor_0ded2dfbe8bca787, []int{0}
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
	return fileDescriptor_0ded2dfbe8bca787, []int{1}
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
	return fileDescriptor_0ded2dfbe8bca787, []int{2}
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
	return fileDescriptor_0ded2dfbe8bca787, []int{3}
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
	Response             bool     `protobuf:"varint,1,opt,name=response,proto3" json:"response"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Boolean) Reset()         { *m = Boolean{} }
func (m *Boolean) String() string { return proto.CompactTextString(m) }
func (*Boolean) ProtoMessage()    {}
func (*Boolean) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ded2dfbe8bca787, []int{4}
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
	return fileDescriptor_0ded2dfbe8bca787, []int{5}
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
	// @tag: validateGeneric:"max=9999"
	Year int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty" validateGeneric:"max=9999"`
	// Month.
	// @tag: validateGeneric:"min=1,max=12"
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty" validateGeneric:"min=1,max=12"`
	// Day.
	// @tag: validateGeneric:"min=1,max=31"
	Day                  int32    `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty" validateGeneric:"min=1,max=31"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Date) Reset()         { *m = Date{} }
func (m *Date) String() string { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()    {}
func (*Date) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ded2dfbe8bca787, []int{6}
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

// A fixed time without timezone information.
type Time struct {
	// Hour.
	// @tag: validateGeneric:"min=0,max=23"
	Hour int32 `protobuf:"varint,1,opt,name=hour,proto3" json:"hour,omitempty" validateGeneric:"min=0,max=23"`
	// Minute.
	// @tag: validateGeneric:"min=0,max=60"
	Minute int32 `protobuf:"varint,2,opt,name=minute,proto3" json:"minute,omitempty" validateGeneric:"min=0,max=60"`
	// Second.
	// @tag: validateGeneric:"min=0,max=60"
	Second               int32    `protobuf:"varint,3,opt,name=second,proto3" json:"second,omitempty" validateGeneric:"min=0,max=60"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Time) Reset()         { *m = Time{} }
func (m *Time) String() string { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()    {}
func (*Time) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ded2dfbe8bca787, []int{7}
}

func (m *Time) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Time.Unmarshal(m, b)
}
func (m *Time) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Time.Marshal(b, m, deterministic)
}
func (m *Time) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Time.Merge(m, src)
}
func (m *Time) XXX_Size() int {
	return xxx_messageInfo_Time.Size(m)
}
func (m *Time) XXX_DiscardUnknown() {
	xxx_messageInfo_Time.DiscardUnknown(m)
}

var xxx_messageInfo_Time proto.InternalMessageInfo

func (m *Time) GetHour() int32 {
	if m != nil {
		return m.Hour
	}
	return 0
}

func (m *Time) GetMinute() int32 {
	if m != nil {
		return m.Minute
	}
	return 0
}

func (m *Time) GetSecond() int32 {
	if m != nil {
		return m.Second
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
	return fileDescriptor_0ded2dfbe8bca787, []int{8}
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
	return fileDescriptor_0ded2dfbe8bca787, []int{9}
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
	// A URL that can be used in an Android app or email. Note that this is not recommended for web use.
	GooglePayURL string `protobuf:"bytes,5,opt,name=googlePayURL,proto3" json:"googlePayURL,omitempty"`
	// The Base64 encoded bytes of the Apple Wallet pass.
	ApplePassBytes []byte `protobuf:"bytes,6,opt,name=applePassBytes,proto3" json:"applePassBytes,omitempty"`
	// A URL for multiple passes, linked by a common index.
	MultiplePassesURL    string   `protobuf:"bytes,7,opt,name=multiplePassesURL,proto3" json:"multiplePassesURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassBundle) Reset()         { *m = PassBundle{} }
func (m *PassBundle) String() string { return proto.CompactTextString(m) }
func (*PassBundle) ProtoMessage()    {}
func (*PassBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ded2dfbe8bca787, []int{10}
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

func (m *PassBundle) GetGooglePayURL() string {
	if m != nil {
		return m.GooglePayURL
	}
	return ""
}

func (m *PassBundle) GetApplePassBytes() []byte {
	if m != nil {
		return m.ApplePassBytes
	}
	return nil
}

func (m *PassBundle) GetMultiplePassesURL() string {
	if m != nil {
		return m.MultiplePassesURL
	}
	return ""
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
	return fileDescriptor_0ded2dfbe8bca787, []int{11}
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

// [DEPRECATED: OR operator is not supported] Filter and list records.
type ListRequestDeprecated struct {
	// @tag: validateGeneric:"required" validateCreate:"required" validateUpdate:"required"
	ClassId string `protobuf:"bytes,1,opt,name=classId,proto3" json:"classId,omitempty" validateGeneric:"required" validateCreate:"required" validateUpdate:"required"`
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	Protocol PassProtocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=io.PassProtocol" json:"protocol,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	Pagination           *Pagination `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListRequestDeprecated) Reset()         { *m = ListRequestDeprecated{} }
func (m *ListRequestDeprecated) String() string { return proto.CompactTextString(m) }
func (*ListRequestDeprecated) ProtoMessage()    {}
func (*ListRequestDeprecated) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ded2dfbe8bca787, []int{12}
}

func (m *ListRequestDeprecated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequestDeprecated.Unmarshal(m, b)
}
func (m *ListRequestDeprecated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequestDeprecated.Marshal(b, m, deterministic)
}
func (m *ListRequestDeprecated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequestDeprecated.Merge(m, src)
}
func (m *ListRequestDeprecated) XXX_Size() int {
	return xxx_messageInfo_ListRequestDeprecated.Size(m)
}
func (m *ListRequestDeprecated) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequestDeprecated.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequestDeprecated proto.InternalMessageInfo

func (m *ListRequestDeprecated) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *ListRequestDeprecated) GetProtocol() PassProtocol {
	if m != nil {
		return m.Protocol
	}
	return PassProtocol_PASS_PROTOCOL_DO_NOT_USE
}

func (m *ListRequestDeprecated) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type ListRequest struct {
	// @tag: validateGeneric:"required" validateCreate:"required" validateUpdate:"required"
	ClassId string `protobuf:"bytes,1,opt,name=classId,proto3" json:"classId,omitempty" validateGeneric:"required" validateCreate:"required" validateUpdate:"required"`
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	Protocol PassProtocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=io.PassProtocol" json:"protocol,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	Filters              *Filters `protobuf:"bytes,3,opt,name=filters,proto3" json:"filters,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ded2dfbe8bca787, []int{13}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *ListRequest) GetProtocol() PassProtocol {
	if m != nil {
		return m.Protocol
	}
	return PassProtocol_PASS_PROTOCOL_DO_NOT_USE
}

func (m *ListRequest) GetFilters() *Filters {
	if m != nil {
		return m.Filters
	}
	return nil
}

// DataItems are used to supply additional metadata when creating or updating passes.
type DataItems struct {
	// Key value pairs. All values must be strings, but formatting and conversion can be defined in the pass template. Keys should be alphanumeric, start with a lowercase level and use camel case.  They can be accessed in the pass template by prefixing with 'meta.'. E.g. 'meta.myCustomKey'.
	Items                map[string]string `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DataItems) Reset()         { *m = DataItems{} }
func (m *DataItems) String() string { return proto.CompactTextString(m) }
func (*DataItems) ProtoMessage()    {}
func (*DataItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ded2dfbe8bca787, []int{14}
}

func (m *DataItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataItems.Unmarshal(m, b)
}
func (m *DataItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataItems.Marshal(b, m, deterministic)
}
func (m *DataItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataItems.Merge(m, src)
}
func (m *DataItems) XXX_Size() int {
	return xxx_messageInfo_DataItems.Size(m)
}
func (m *DataItems) XXX_DiscardUnknown() {
	xxx_messageInfo_DataItems.DiscardUnknown(m)
}

var xxx_messageInfo_DataItems proto.InternalMessageInfo

func (m *DataItems) GetItems() map[string]string {
	if m != nil {
		return m.Items
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
	proto.RegisterType((*Time)(nil), "io.Time")
	proto.RegisterType((*LocalDate)(nil), "io.LocalDate")
	proto.RegisterType((*LocalDateTime)(nil), "io.LocalDateTime")
	proto.RegisterType((*PassBundle)(nil), "io.PassBundle")
	proto.RegisterType((*PassBundleRequest)(nil), "io.PassBundleRequest")
	proto.RegisterType((*ListRequestDeprecated)(nil), "io.ListRequestDeprecated")
	proto.RegisterType((*ListRequest)(nil), "io.ListRequest")
	proto.RegisterType((*DataItems)(nil), "io.DataItems")
	proto.RegisterMapType((map[string]string)(nil), "io.DataItems.ItemsEntry")
}

func init() {
	proto.RegisterFile("io/common/common_objects.proto", fileDescriptor_0ded2dfbe8bca787)
}

var fileDescriptor_0ded2dfbe8bca787 = []byte{
	// 927 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xed, 0x6e, 0xe3, 0x44,
	0x14, 0x5d, 0xbb, 0x5f, 0xc9, 0x6d, 0x09, 0xee, 0xa8, 0x5b, 0x42, 0x04, 0xab, 0x95, 0xb5, 0x8b,
	0xd2, 0xaa, 0x49, 0xda, 0x74, 0x41, 0xdb, 0x15, 0x62, 0xe5, 0x34, 0x6e, 0x15, 0x6a, 0x5a, 0xe3,
	0x26, 0x5a, 0x20, 0x2a, 0xd1, 0x6c, 0x3c, 0x75, 0x87, 0xda, 0x9e, 0xe0, 0x19, 0x03, 0x11, 0xdd,
	0x17, 0xd8, 0x3f, 0x48, 0x2b, 0x9e, 0x01, 0x21, 0x9e, 0x12, 0xcd, 0xd8, 0x4e, 0x4a, 0xcb, 0x2f,
	0xc4, 0x9f, 0xe4, 0x9e, 0x73, 0xcf, 0x9c, 0x7b, 0x6c, 0x5f, 0x1b, 0x1e, 0x51, 0xd6, 0x1a, 0xb3,
	0x28, 0x62, 0x71, 0xfe, 0x37, 0x62, 0xaf, 0x7f, 0x20, 0x63, 0xc1, 0x9b, 0x93, 0x84, 0x09, 0x86,
	0x74, 0xca, 0x6a, 0x3b, 0xaa, 0x1c, 0x37, 0x02, 0x12, 0x37, 0xf8, 0xcf, 0x38, 0x08, 0x48, 0xd2,
	0x62, 0x13, 0x41, 0x59, 0xcc, 0x5b, 0x38, 0x8e, 0x99, 0xc0, 0xaa, 0xce, 0x4e, 0xd4, 0x6a, 0x73,
	0xc7, 0x09, 0x0e, 0x68, 0xac, 0x9a, 0x79, 0x6f, 0x73, 0xde, 0xbb, 0xa4, 0xa1, 0x20, 0x49, 0xce,
	0x7f, 0x78, 0xeb, 0x8c, 0x9a, 0xc5, 0xc2, 0xdc, 0xce, 0xdc, 0x00, 0xbd, 0xe7, 0xa3, 0x0a, 0xe8,
	0xd4, 0xaf, 0x6a, 0x8f, 0xb5, 0x7a, 0xd9, 0xd3, 0xa9, 0x6f, 0x7e, 0x00, 0x0b, 0x83, 0x24, 0x44,
	0x06, 0x2c, 0xa4, 0x49, 0x98, 0xf3, 0xb2, 0x34, 0x3f, 0x86, 0xa5, 0x43, 0x96, 0xc6, 0x02, 0x6d,
	0xc0, 0x92, 0x60, 0x02, 0x67, 0xcd, 0x25, 0x2f, 0x03, 0xe6, 0x16, 0x94, 0x8f, 0x68, 0x48, 0x3a,
	0x53, 0x41, 0x38, 0xfa, 0x08, 0xca, 0x97, 0x05, 0x50, 0xb2, 0x35, 0x6f, 0x4e, 0x98, 0x4f, 0x61,
	0xa5, 0xc3, 0x58, 0x48, 0x70, 0x8c, 0x6a, 0x50, 0x4a, 0x08, 0x9f, 0xb0, 0x98, 0x13, 0xa5, 0x2b,
	0x79, 0x33, 0x2c, 0x65, 0xe7, 0x22, 0xa1, 0x71, 0xc0, 0xef, 0xc8, 0x16, 0xea, 0xe5, 0x5b, 0xb2,
	0x0e, 0x2c, 0x76, 0xb1, 0x20, 0x08, 0xc1, 0xe2, 0x94, 0xe0, 0x24, 0x4f, 0xa5, 0x6a, 0x19, 0x35,
	0x62, 0xb1, 0xb8, 0xaa, 0xea, 0x59, 0x54, 0x05, 0xe4, 0xb5, 0xf9, 0x78, 0x5a, 0x5d, 0x50, 0x9c,
	0x2c, 0xcd, 0x2f, 0x61, 0xb1, 0x4f, 0x23, 0xe5, 0x71, 0xc5, 0xd2, 0x99, 0x87, 0xac, 0xd1, 0x26,
	0x2c, 0x47, 0x34, 0x4e, 0x05, 0xc9, 0x4d, 0x72, 0x24, 0x79, 0x4e, 0xc6, 0x2c, 0xf6, 0x73, 0xa3,
	0x1c, 0x99, 0x04, 0xca, 0x0e, 0x1b, 0xe3, 0x50, 0x85, 0xfa, 0x06, 0x4a, 0x3e, 0x16, 0x44, 0x9a,
	0x67, 0xf7, 0xb2, 0xf3, 0xf9, 0x3b, 0xeb, 0xe0, 0xad, 0xf6, 0xd9, 0xf7, 0xf5, 0xe1, 0x6e, 0xe3,
	0xe0, 0xe2, 0xd7, 0x67, 0x6f, 0xb6, 0x1a, 0x2f, 0xeb, 0x7b, 0xc3, 0xdd, 0x46, 0xfb, 0xe2, 0x66,
	0x77, 0xb8, 0xd7, 0x38, 0xb8, 0x90, 0xc4, 0xfe, 0x70, 0x77, 0xaf, 0xc0, 0x37, 0xc3, 0xbd, 0xf6,
	0x85, 0x52, 0x6f, 0x79, 0x33, 0x37, 0xf3, 0x0f, 0x0d, 0xde, 0x9b, 0xcd, 0x51, 0xe1, 0x7f, 0xd7,
	0xee, 0x0d, 0xfb, 0xe5, 0x9d, 0x95, 0xbe, 0xd5, 0x92, 0xff, 0x36, 0xac, 0x5f, 0x6f, 0x0f, 0x77,
	0x1b, 0xfb, 0x17, 0x37, 0xb2, 0x9d, 0x51, 0x2f, 0xa4, 0xd1, 0xa7, 0xff, 0x06, 0xea, 0xcd, 0x6c,
	0xc4, 0xde, 0xce, 0xfe, 0x9b, 0xad, 0x97, 0x4f, 0x6e, 0x05, 0xfd, 0x53, 0x03, 0x70, 0x31, 0xe7,
	0x9d, 0x34, 0xf6, 0x43, 0x72, 0x77, 0xdf, 0x8a, 0x45, 0xd3, 0x67, 0x8b, 0x86, 0x4c, 0x58, 0x0b,
	0x18, 0x0b, 0x42, 0xe2, 0xe2, 0xe9, 0xc0, 0x73, 0xaa, 0x4b, 0xaa, 0xf5, 0x0f, 0x0e, 0x7d, 0x02,
	0x15, 0x3c, 0x99, 0x48, 0xc8, 0x79, 0xb6, 0x65, 0xcb, 0x6a, 0xcb, 0xee, 0xb0, 0x68, 0x07, 0xd6,
	0xa3, 0x34, 0x14, 0x34, 0x27, 0x09, 0x97, 0x86, 0x2b, 0xca, 0xf0, 0x7e, 0xc3, 0xfc, 0x1a, 0xd6,
	0xe7, 0x49, 0x3d, 0xf2, 0x63, 0x4a, 0xb8, 0xb8, 0x17, 0x78, 0x07, 0x96, 0x2f, 0x59, 0x12, 0x61,
	0x51, 0xd5, 0x1f, 0x2f, 0xd4, 0x2b, 0xed, 0x8d, 0x26, 0x65, 0xcd, 0xf9, 0xb1, 0x23, 0xd5, 0xf3,
	0x72, 0x8d, 0xf9, 0x9b, 0x06, 0x0f, 0x1d, 0xca, 0x45, 0xee, 0xd6, 0x25, 0x93, 0x84, 0x8c, 0xb1,
	0x20, 0x3e, 0xaa, 0xc2, 0xca, 0x38, 0xc4, 0x9c, 0xf7, 0x0a, 0xf3, 0x02, 0xa2, 0x1d, 0x28, 0x15,
	0xef, 0xaa, 0xba, 0x2f, 0x95, 0xb6, 0x51, 0xcc, 0x70, 0x73, 0xde, 0x9b, 0x29, 0x50, 0x13, 0x60,
	0xfe, 0x35, 0x50, 0xbb, 0xb8, 0xda, 0xae, 0x64, 0xfa, 0x82, 0xf5, 0x6e, 0x29, 0xcc, 0x1b, 0x58,
	0xbd, 0x15, 0xe8, 0x7f, 0x8b, 0xf1, 0x14, 0x56, 0xb2, 0x0f, 0x0f, 0xcf, 0x33, 0xac, 0x4a, 0xf1,
	0x51, 0x46, 0x79, 0x45, 0xcf, 0x4c, 0xa1, 0xdc, 0xc5, 0x02, 0xf7, 0x04, 0x89, 0x38, 0x6a, 0xc2,
	0x12, 0x95, 0x85, 0x7a, 0xa7, 0x57, 0xdb, 0x55, 0x79, 0x62, 0xd6, 0x6d, 0xaa, 0x5f, 0x3b, 0x16,
	0xc9, 0xd4, 0xcb, 0x64, 0xb5, 0xe7, 0x00, 0x73, 0x52, 0x6e, 0xce, 0x35, 0x99, 0x16, 0x9f, 0xa8,
	0x6b, 0x32, 0x95, 0xaf, 0xfb, 0x4f, 0x38, 0x4c, 0x49, 0xbe, 0x4d, 0x19, 0x78, 0xa1, 0x3f, 0xd7,
	0xb6, 0x9f, 0xc1, 0x9a, 0xbc, 0x68, 0x97, 0x71, 0x2a, 0x6f, 0x02, 0x02, 0x58, 0xb6, 0x5c, 0xd7,
	0x3e, 0xed, 0x1a, 0x0f, 0xd0, 0x2a, 0xac, 0xb8, 0x9e, 0xad, 0x80, 0x26, 0x81, 0x67, 0xbb, 0x8e,
	0x75, 0x68, 0x1b, 0xfa, 0xf6, 0x2b, 0x30, 0xee, 0x3e, 0x58, 0xb4, 0x06, 0x25, 0xd7, 0x3a, 0x3f,
	0x1f, 0x0d, 0x3c, 0xc7, 0x78, 0x80, 0x2a, 0x00, 0xc7, 0x67, 0x67, 0xc7, 0x8e, 0xad, 0xb0, 0x8e,
	0x1e, 0xc2, 0xba, 0xe5, 0xba, 0x8e, 0x3d, 0x52, 0x9a, 0xce, 0xe0, 0xb4, 0xeb, 0xd8, 0x46, 0x49,
	0xca, 0xbe, 0x1a, 0x38, 0xfd, 0xde, 0xc8, 0xe9, 0x9d, 0x9e, 0x18, 0xc6, 0x76, 0x00, 0xe5, 0x01,
	0xc7, 0x01, 0xe9, 0x4f, 0x27, 0x44, 0x3a, 0x9e, 0x9e, 0x8d, 0x06, 0xe7, 0xd6, 0xb1, 0x6d, 0x3c,
	0x40, 0x9b, 0x80, 0x54, 0x39, 0xca, 0x7c, 0x5e, 0x59, 0x8e, 0x63, 0xf7, 0x0d, 0x0d, 0x6d, 0x80,
	0x91, 0xf1, 0xf9, 0x3c, 0xd7, 0xfa, 0xd6, 0xd0, 0xd1, 0x23, 0xa8, 0x65, 0x6c, 0xd7, 0xea, 0x5b,
	0xa3, 0xc3, 0x33, 0xc7, 0xb1, 0x0f, 0xfb, 0xbd, 0xb3, 0xd3, 0x91, 0x2b, 0xdd, 0x16, 0x3b, 0x5f,
	0xc0, 0xfb, 0x94, 0x35, 0x27, 0x98, 0xf3, 0x6b, 0x2a, 0x9a, 0xee, 0x09, 0x65, 0xdf, 0x3d, 0xe1,
	0x02, 0xf3, 0xab, 0x19, 0x37, 0x66, 0x51, 0x8b, 0xb2, 0x56, 0xc4, 0x7c, 0x12, 0xb6, 0xb8, 0x7f,
	0xdd, 0x0a, 0x58, 0x8b, 0xb2, 0xbf, 0xf4, 0x45, 0xf7, 0xa4, 0xc7, 0x5e, 0x2f, 0xab, 0xe7, 0xbb,
	0xff, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0x3a, 0xa2, 0xc1, 0xcd, 0x06, 0x00, 0x00,
}
