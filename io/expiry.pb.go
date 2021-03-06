// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/common/expiry.proto

package io

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

// Options to manage the expiry date of the digital card.
type ExpiryType int32

const (
	// Please do not use this enum. This enum do not have any effect on expiry logic.
	ExpiryType_EXPIRE_NONE ExpiryType = 0
	// Expiry date is set with year, month and date. The expiry date will be based on your timezone. The digital card will expire at 23:59:59:59.99999 of set date in fixed timezone. The expiry date is the same for all cards.
	ExpiryType_EXPIRE_ON_FIXED_DATE ExpiryType = 1
	// The digital card expires after the number of days after the digital card issuing.
	ExpiryType_EXPIRE_AFTER_X_DAYS ExpiryType = 2
	// If you want to change expiry date for each digital card, you can use this expiry type. You can set expiry date and time in fixed timezone.
	ExpiryType_EXPIRE_ON_VARIABLE_DATE_TIME ExpiryType = 3
	// The digital card will set as NULL and the pass will not expire..
	ExpiryType_EXPIRE_SET_TO_NULL ExpiryType = 4
)

var ExpiryType_name = map[int32]string{
	0: "EXPIRE_NONE",
	1: "EXPIRE_ON_FIXED_DATE",
	2: "EXPIRE_AFTER_X_DAYS",
	3: "EXPIRE_ON_VARIABLE_DATE_TIME",
	4: "EXPIRE_SET_TO_NULL",
}

var ExpiryType_value = map[string]int32{
	"EXPIRE_NONE":                  0,
	"EXPIRE_ON_FIXED_DATE":         1,
	"EXPIRE_AFTER_X_DAYS":          2,
	"EXPIRE_ON_VARIABLE_DATE_TIME": 3,
	"EXPIRE_SET_TO_NULL":           4,
}

func (x ExpiryType) String() string {
	return proto.EnumName(ExpiryType_name, int32(x))
}

func (ExpiryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4e8d5f7792227f89, []int{0}
}

// The digital card will be expired on the expiry date. The barcode will not be rendered on digital card and the card itself will not be updated after it has been expired.
type ExpirySettings struct {
	ExpiryType ExpiryType `protobuf:"varint,1,opt,name=expiryType,proto3,enum=io.ExpiryType" json:"expiryType,omitempty"`
	// Types that are valid to be assigned to ExpiryOneof:
	//	*ExpirySettings_FixedExpiryDate
	//	*ExpirySettings_ExpireAfterXDays
	ExpiryOneof          isExpirySettings_ExpiryOneof `protobuf_oneof:"expiryOneof"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ExpirySettings) Reset()         { *m = ExpirySettings{} }
func (m *ExpirySettings) String() string { return proto.CompactTextString(m) }
func (*ExpirySettings) ProtoMessage()    {}
func (*ExpirySettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e8d5f7792227f89, []int{0}
}

func (m *ExpirySettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpirySettings.Unmarshal(m, b)
}
func (m *ExpirySettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpirySettings.Marshal(b, m, deterministic)
}
func (m *ExpirySettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpirySettings.Merge(m, src)
}
func (m *ExpirySettings) XXX_Size() int {
	return xxx_messageInfo_ExpirySettings.Size(m)
}
func (m *ExpirySettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpirySettings.DiscardUnknown(m)
}

var xxx_messageInfo_ExpirySettings proto.InternalMessageInfo

func (m *ExpirySettings) GetExpiryType() ExpiryType {
	if m != nil {
		return m.ExpiryType
	}
	return ExpiryType_EXPIRE_NONE
}

type isExpirySettings_ExpiryOneof interface {
	isExpirySettings_ExpiryOneof()
}

type ExpirySettings_FixedExpiryDate struct {
	FixedExpiryDate *Date `protobuf:"bytes,2,opt,name=fixedExpiryDate,proto3,oneof"`
}

type ExpirySettings_ExpireAfterXDays struct {
	ExpireAfterXDays uint32 `protobuf:"varint,3,opt,name=expireAfterXDays,proto3,oneof"`
}

func (*ExpirySettings_FixedExpiryDate) isExpirySettings_ExpiryOneof() {}

func (*ExpirySettings_ExpireAfterXDays) isExpirySettings_ExpiryOneof() {}

func (m *ExpirySettings) GetExpiryOneof() isExpirySettings_ExpiryOneof {
	if m != nil {
		return m.ExpiryOneof
	}
	return nil
}

func (m *ExpirySettings) GetFixedExpiryDate() *Date {
	if x, ok := m.GetExpiryOneof().(*ExpirySettings_FixedExpiryDate); ok {
		return x.FixedExpiryDate
	}
	return nil
}

func (m *ExpirySettings) GetExpireAfterXDays() uint32 {
	if x, ok := m.GetExpiryOneof().(*ExpirySettings_ExpireAfterXDays); ok {
		return x.ExpireAfterXDays
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExpirySettings) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExpirySettings_FixedExpiryDate)(nil),
		(*ExpirySettings_ExpireAfterXDays)(nil),
	}
}

func init() {
	proto.RegisterEnum("io.ExpiryType", ExpiryType_name, ExpiryType_value)
	proto.RegisterType((*ExpirySettings)(nil), "io.ExpirySettings")
}

func init() {
	proto.RegisterFile("io/common/expiry.proto", fileDescriptor_4e8d5f7792227f89)
}

var fileDescriptor_4e8d5f7792227f89 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xd1, 0x6a, 0xe2, 0x40,
	0x14, 0x86, 0x4d, 0x94, 0x65, 0x39, 0xa2, 0x86, 0xd9, 0xc5, 0x0d, 0xcb, 0xb2, 0x84, 0xd2, 0x0b,
	0x29, 0x25, 0x01, 0xdb, 0xeb, 0x42, 0x24, 0x23, 0x06, 0x6d, 0x22, 0x31, 0x2d, 0x69, 0x6f, 0x86,
	0xa8, 0xa3, 0x9d, 0xda, 0x78, 0x82, 0x33, 0x17, 0xfa, 0x0a, 0x7d, 0x94, 0xde, 0xf6, 0x05, 0x8b,
	0x89, 0x68, 0x69, 0xaf, 0x86, 0xf9, 0xff, 0xef, 0x3b, 0x33, 0x70, 0xa0, 0x2d, 0xd0, 0x99, 0x61,
	0x96, 0xe1, 0xda, 0xe1, 0xdb, 0x5c, 0x6c, 0x76, 0x76, 0xbe, 0x41, 0x85, 0x44, 0x17, 0xf8, 0xf7,
	0xff, 0xa9, 0x2b, 0x0f, 0x86, 0xd3, 0x67, 0x3e, 0x53, 0xb2, 0x64, 0xce, 0xde, 0x35, 0x68, 0xd2,
	0x42, 0x9a, 0x70, 0xa5, 0xc4, 0x7a, 0x29, 0x89, 0x0d, 0x50, 0x8e, 0x89, 0x77, 0x39, 0x37, 0x35,
	0x4b, 0xeb, 0x34, 0xbb, 0x4d, 0x5b, 0xa0, 0x4d, 0x8f, 0x69, 0xf4, 0x89, 0x20, 0xd7, 0xd0, 0x5a,
	0x88, 0x2d, 0x9f, 0x97, 0xb5, 0x97, 0x2a, 0x6e, 0xea, 0x96, 0xd6, 0xa9, 0x77, 0x7f, 0xee, 0xa5,
	0xfd, 0x7d, 0x50, 0x89, 0xbe, 0x22, 0xe4, 0x12, 0x8c, 0x62, 0x06, 0x77, 0x17, 0x8a, 0x6f, 0x12,
	0x2f, 0xdd, 0x49, 0xb3, 0x6a, 0x69, 0x9d, 0xc6, 0xa0, 0x12, 0x7d, 0x6b, 0x7a, 0x0d, 0xa8, 0x97,
	0x2f, 0x86, 0x6b, 0x8e, 0x8b, 0x8b, 0x57, 0x0d, 0xe0, 0xf4, 0x1b, 0xd2, 0x82, 0x3a, 0x4d, 0xc6,
	0x7e, 0x44, 0x59, 0x10, 0x06, 0xd4, 0xa8, 0x10, 0x13, 0x7e, 0x1f, 0x82, 0x30, 0x60, 0x7d, 0x3f,
	0xa1, 0x1e, 0xf3, 0xdc, 0x98, 0x1a, 0x1a, 0xf9, 0x03, 0xbf, 0x0e, 0x8d, 0xdb, 0x8f, 0x69, 0xc4,
	0x12, 0xe6, 0xb9, 0x0f, 0x13, 0x43, 0x27, 0x16, 0xfc, 0x3b, 0x29, 0xf7, 0x6e, 0xe4, 0xbb, 0xbd,
	0x11, 0x2d, 0x2c, 0x16, 0xfb, 0xb7, 0xd4, 0xa8, 0x92, 0x36, 0x90, 0x03, 0x31, 0xa1, 0x31, 0x8b,
	0x43, 0x16, 0xdc, 0x8d, 0x46, 0x46, 0xad, 0x77, 0x03, 0x2d, 0x81, 0x76, 0x9e, 0x4a, 0xb9, 0x12,
	0xca, 0x1e, 0x0f, 0x05, 0x3e, 0x9e, 0x4b, 0x95, 0xca, 0xa7, 0x63, 0x36, 0xc3, 0xcc, 0x11, 0xe8,
	0x64, 0x38, 0xe7, 0x2f, 0x8e, 0x9c, 0xaf, 0x9c, 0x25, 0x3a, 0x02, 0xdf, 0xf4, 0xda, 0x78, 0xe8,
	0xe3, 0xf4, 0x47, 0xb1, 0x89, 0xab, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0x14, 0x28, 0xfa,
	0xc7, 0x01, 0x00, 0x00,
}
