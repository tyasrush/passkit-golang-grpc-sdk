// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/common/distribution.proto

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

type DistributionChannel int32

const (
	DistributionChannel_NO_DISTRIBUTION DistributionChannel = 0
	// Unique pass link is distributed via email (will use default PK email template if no custom Email Template is provided)
	DistributionChannel_CHANNEL_EMAIL DistributionChannel = 1
	// Unique pass link is distributed via SMS (will use default PK SMS template if no custom SMS Template is provided)
	DistributionChannel_CHANNEL_SMS DistributionChannel = 2
)

var DistributionChannel_name = map[int32]string{
	0: "NO_DISTRIBUTION",
	1: "CHANNEL_EMAIL",
	2: "CHANNEL_SMS",
}

var DistributionChannel_value = map[string]int32{
	"NO_DISTRIBUTION": 0,
	"CHANNEL_EMAIL":   1,
	"CHANNEL_SMS":     2,
}

func (x DistributionChannel) String() string {
	return proto.EnumName(DistributionChannel_name, int32(x))
}

func (DistributionChannel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_737c63c5d3d253ea, []int{0}
}

type EmailDistributionRequest struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Protocol             PassProtocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=io.PassProtocol" json:"protocol,omitempty"`
	AlternativeEmail     string       `protobuf:"bytes,3,opt,name=alternativeEmail,proto3" json:"alternativeEmail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *EmailDistributionRequest) Reset()         { *m = EmailDistributionRequest{} }
func (m *EmailDistributionRequest) String() string { return proto.CompactTextString(m) }
func (*EmailDistributionRequest) ProtoMessage()    {}
func (*EmailDistributionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_737c63c5d3d253ea, []int{0}
}

func (m *EmailDistributionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailDistributionRequest.Unmarshal(m, b)
}
func (m *EmailDistributionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailDistributionRequest.Marshal(b, m, deterministic)
}
func (m *EmailDistributionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailDistributionRequest.Merge(m, src)
}
func (m *EmailDistributionRequest) XXX_Size() int {
	return xxx_messageInfo_EmailDistributionRequest.Size(m)
}
func (m *EmailDistributionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailDistributionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailDistributionRequest proto.InternalMessageInfo

func (m *EmailDistributionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EmailDistributionRequest) GetProtocol() PassProtocol {
	if m != nil {
		return m.Protocol
	}
	return PassProtocol_RAW_PROTOCOL
}

func (m *EmailDistributionRequest) GetAlternativeEmail() string {
	if m != nil {
		return m.AlternativeEmail
	}
	return ""
}

// DistributionSettings can be used by protocol top level elements to define the email / sms-es
type DistributionSettings struct {
	DistributionChannels []DistributionChannel `protobuf:"varint,1,rep,packed,name=distributionChannels,proto3,enum=io.DistributionChannel" json:"distributionChannels,omitempty"`
	// Welcome email settings
	WelcomeEmail *EmailTemplate `protobuf:"bytes,2,opt,name=welcomeEmail,proto3" json:"welcomeEmail,omitempty"`
	// Welcome SMS settings
	WelcomeSms           *SmsTemplate `protobuf:"bytes,3,opt,name=welcomeSms,proto3" json:"welcomeSms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DistributionSettings) Reset()         { *m = DistributionSettings{} }
func (m *DistributionSettings) String() string { return proto.CompactTextString(m) }
func (*DistributionSettings) ProtoMessage()    {}
func (*DistributionSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_737c63c5d3d253ea, []int{1}
}

func (m *DistributionSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributionSettings.Unmarshal(m, b)
}
func (m *DistributionSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributionSettings.Marshal(b, m, deterministic)
}
func (m *DistributionSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionSettings.Merge(m, src)
}
func (m *DistributionSettings) XXX_Size() int {
	return xxx_messageInfo_DistributionSettings.Size(m)
}
func (m *DistributionSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionSettings.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionSettings proto.InternalMessageInfo

func (m *DistributionSettings) GetDistributionChannels() []DistributionChannel {
	if m != nil {
		return m.DistributionChannels
	}
	return nil
}

func (m *DistributionSettings) GetWelcomeEmail() *EmailTemplate {
	if m != nil {
		return m.WelcomeEmail
	}
	return nil
}

func (m *DistributionSettings) GetWelcomeSms() *SmsTemplate {
	if m != nil {
		return m.WelcomeSms
	}
	return nil
}

// EmailTemplate contains details for sending an email to a customer.
// Uses the default PassKit email provider and sending credentials
type EmailTemplate struct {
	// Subject of the email. Can contain any of the fields from fields array: ${DISPLAY_NAME}, ${EMAIL_ADDRESS}, etc..
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	// Localized subject of the email. Can contain any of the fields from fields array: ${DISPLAY_NAME}, ${EMAIL_ADDRESS}, etc.
	LocalizedSubject *LocalizedString `protobuf:"bytes,2,opt,name=localizedSubject,proto3" json:"localizedSubject,omitempty"`
	// Text content of the email (for email clients that don't support HTML). The body text is rendered above pass install url text. Can contain any of the fields from fields array: ${DISPLAY_NAME}, ${EMAIL_ADDRESS}, etc.
	BodyTextContent string `protobuf:"bytes,3,opt,name=bodyTextContent,proto3" json:"bodyTextContent,omitempty"`
	// Localized text content of the email (for email clients that don't support HTML). The body text is rendered above pass install url text. Can contain any of the fields from fields array: ${DISPLAY_NAME}, ${EMAIL_ADDRESS}, etc.
	LocalizedBodyTextContent *LocalizedString `protobuf:"bytes,4,opt,name=localizedBodyTextContent,proto3" json:"localizedBodyTextContent,omitempty"`
	// Content for HTML email rendered above the pass install button. Can contain any of the fields from fields array: ${DISPLAY_NAME}, ${EMAIL_ADDRESS}, etc.
	BodyHtmlContent string `protobuf:"bytes,5,opt,name=bodyHtmlContent,proto3" json:"bodyHtmlContent,omitempty"`
	// Localized content for HTML email. Can contain any of the fields from fields array: ${DISPLAY_NAME}, ${EMAIL_ADDRESS}, etc.
	LocalizedBodyHtmlContent *LocalizedString `protobuf:"bytes,6,opt,name=localizedBodyHtmlContent,proto3" json:"localizedBodyHtmlContent,omitempty"`
	// 7 is reserved for custom email SMTP configuration
	Configuration *EmailConfiguration `protobuf:"bytes,7,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// Text of the pass install button.
	ButtonText string `protobuf:"bytes,8,opt,name=buttonText,proto3" json:"buttonText,omitempty"`
	// Text color of the pass install button.
	ButtonTextColor string `protobuf:"bytes,9,opt,name=buttonTextColor,proto3" json:"buttonTextColor,omitempty"`
	// Color of the pass install button.
	ButtonBackgroundColor string `protobuf:"bytes,10,opt,name=buttonBackgroundColor,proto3" json:"buttonBackgroundColor,omitempty"`
	// Radius of the pass install button. The unit can be px or %. Default is px.
	ButtonBorderRadius string `protobuf:"bytes,11,opt,name=buttonBorderRadius,proto3" json:"buttonBorderRadius,omitempty"`
	// Footer text content (for email clients that don't support HTML). The footer text is rendered below the pass install url text.
	FooterTextContent string `protobuf:"bytes,12,opt,name=footerTextContent,proto3" json:"footerTextContent,omitempty"`
	// Localized footer text content (for email clients that don't support HTML). The footer text is rendered below the pass install url text.
	LocalizedFooterTextContent *LocalizedString `protobuf:"bytes,13,opt,name=localizedFooterTextContent,proto3" json:"localizedFooterTextContent,omitempty"`
	// Footer HTML content rendered below the pass install button.
	FooterHtmlContent string `protobuf:"bytes,14,opt,name=footerHtmlContent,proto3" json:"footerHtmlContent,omitempty"`
	// Localized content for HTML footer content which is rendered below the pass install button.
	LocalizedFooterHtmlContent *LocalizedString `protobuf:"bytes,15,opt,name=localizedFooterHtmlContent,proto3" json:"localizedFooterHtmlContent,omitempty"`
	// Background color for the message content section.
	MessageBackgroundColor string `protobuf:"bytes,16,opt,name=messageBackgroundColor,proto3" json:"messageBackgroundColor,omitempty"`
	// Background color for the margin space surrounding the message content. This setting is relevant for PC viewers.
	PageBackgroundColor  string   `protobuf:"bytes,17,opt,name=pageBackgroundColor,proto3" json:"pageBackgroundColor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailTemplate) Reset()         { *m = EmailTemplate{} }
func (m *EmailTemplate) String() string { return proto.CompactTextString(m) }
func (*EmailTemplate) ProtoMessage()    {}
func (*EmailTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_737c63c5d3d253ea, []int{2}
}

func (m *EmailTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailTemplate.Unmarshal(m, b)
}
func (m *EmailTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailTemplate.Marshal(b, m, deterministic)
}
func (m *EmailTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailTemplate.Merge(m, src)
}
func (m *EmailTemplate) XXX_Size() int {
	return xxx_messageInfo_EmailTemplate.Size(m)
}
func (m *EmailTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_EmailTemplate proto.InternalMessageInfo

func (m *EmailTemplate) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *EmailTemplate) GetLocalizedSubject() *LocalizedString {
	if m != nil {
		return m.LocalizedSubject
	}
	return nil
}

func (m *EmailTemplate) GetBodyTextContent() string {
	if m != nil {
		return m.BodyTextContent
	}
	return ""
}

func (m *EmailTemplate) GetLocalizedBodyTextContent() *LocalizedString {
	if m != nil {
		return m.LocalizedBodyTextContent
	}
	return nil
}

func (m *EmailTemplate) GetBodyHtmlContent() string {
	if m != nil {
		return m.BodyHtmlContent
	}
	return ""
}

func (m *EmailTemplate) GetLocalizedBodyHtmlContent() *LocalizedString {
	if m != nil {
		return m.LocalizedBodyHtmlContent
	}
	return nil
}

func (m *EmailTemplate) GetConfiguration() *EmailConfiguration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *EmailTemplate) GetButtonText() string {
	if m != nil {
		return m.ButtonText
	}
	return ""
}

func (m *EmailTemplate) GetButtonTextColor() string {
	if m != nil {
		return m.ButtonTextColor
	}
	return ""
}

func (m *EmailTemplate) GetButtonBackgroundColor() string {
	if m != nil {
		return m.ButtonBackgroundColor
	}
	return ""
}

func (m *EmailTemplate) GetButtonBorderRadius() string {
	if m != nil {
		return m.ButtonBorderRadius
	}
	return ""
}

func (m *EmailTemplate) GetFooterTextContent() string {
	if m != nil {
		return m.FooterTextContent
	}
	return ""
}

func (m *EmailTemplate) GetLocalizedFooterTextContent() *LocalizedString {
	if m != nil {
		return m.LocalizedFooterTextContent
	}
	return nil
}

func (m *EmailTemplate) GetFooterHtmlContent() string {
	if m != nil {
		return m.FooterHtmlContent
	}
	return ""
}

func (m *EmailTemplate) GetLocalizedFooterHtmlContent() *LocalizedString {
	if m != nil {
		return m.LocalizedFooterHtmlContent
	}
	return nil
}

func (m *EmailTemplate) GetMessageBackgroundColor() string {
	if m != nil {
		return m.MessageBackgroundColor
	}
	return ""
}

func (m *EmailTemplate) GetPageBackgroundColor() string {
	if m != nil {
		return m.PageBackgroundColor
	}
	return ""
}

type EmailConfiguration struct {
	// Sender's email address.
	EmailFrom string `protobuf:"bytes,1,opt,name=emailFrom,proto3" json:"emailFrom,omitempty"`
	// Sender's company / organization name.
	EmailFromName string `protobuf:"bytes,2,opt,name=emailFromName,proto3" json:"emailFromName,omitempty"`
	// Indicates that PassKit is authorized to send from the emailFrom address (either via AWS or Google Cloud)
	EmailFromVerifiedForSending bool     `protobuf:"varint,3,opt,name=EmailFromVerifiedForSending,proto3" json:"EmailFromVerifiedForSending,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *EmailConfiguration) Reset()         { *m = EmailConfiguration{} }
func (m *EmailConfiguration) String() string { return proto.CompactTextString(m) }
func (*EmailConfiguration) ProtoMessage()    {}
func (*EmailConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_737c63c5d3d253ea, []int{3}
}

func (m *EmailConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailConfiguration.Unmarshal(m, b)
}
func (m *EmailConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailConfiguration.Marshal(b, m, deterministic)
}
func (m *EmailConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailConfiguration.Merge(m, src)
}
func (m *EmailConfiguration) XXX_Size() int {
	return xxx_messageInfo_EmailConfiguration.Size(m)
}
func (m *EmailConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_EmailConfiguration proto.InternalMessageInfo

func (m *EmailConfiguration) GetEmailFrom() string {
	if m != nil {
		return m.EmailFrom
	}
	return ""
}

func (m *EmailConfiguration) GetEmailFromName() string {
	if m != nil {
		return m.EmailFromName
	}
	return ""
}

func (m *EmailConfiguration) GetEmailFromVerifiedForSending() bool {
	if m != nil {
		return m.EmailFromVerifiedForSending
	}
	return false
}

// SMS contains details for a sending an sms to customers.
// Uses the default PassKit SMS provider and sending credentials.
type SmsTemplate struct {
	// content of the SMS; needs to be limited to 70 unicode characters. Can contain any of the fields from fields array: ${DISPLAY_NAME}, ${EMAIL_ADDRESS}, etc.
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// localized content of the SMS; needs to be limited to 70 unicode characters. Can contain any of the fields from fields array: ${DISPLAY_NAME}, ${EMAIL_ADDRESS}, etc.
	LocalizedContent     *LocalizedString `protobuf:"bytes,2,opt,name=localizedContent,proto3" json:"localizedContent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SmsTemplate) Reset()         { *m = SmsTemplate{} }
func (m *SmsTemplate) String() string { return proto.CompactTextString(m) }
func (*SmsTemplate) ProtoMessage()    {}
func (*SmsTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_737c63c5d3d253ea, []int{4}
}

func (m *SmsTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmsTemplate.Unmarshal(m, b)
}
func (m *SmsTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmsTemplate.Marshal(b, m, deterministic)
}
func (m *SmsTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmsTemplate.Merge(m, src)
}
func (m *SmsTemplate) XXX_Size() int {
	return xxx_messageInfo_SmsTemplate.Size(m)
}
func (m *SmsTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_SmsTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_SmsTemplate proto.InternalMessageInfo

func (m *SmsTemplate) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SmsTemplate) GetLocalizedContent() *LocalizedString {
	if m != nil {
		return m.LocalizedContent
	}
	return nil
}

func init() {
	proto.RegisterEnum("io.DistributionChannel", DistributionChannel_name, DistributionChannel_value)
	proto.RegisterType((*EmailDistributionRequest)(nil), "io.EmailDistributionRequest")
	proto.RegisterType((*DistributionSettings)(nil), "io.DistributionSettings")
	proto.RegisterType((*EmailTemplate)(nil), "io.EmailTemplate")
	proto.RegisterType((*EmailConfiguration)(nil), "io.EmailConfiguration")
	proto.RegisterType((*SmsTemplate)(nil), "io.SmsTemplate")
}

func init() { proto.RegisterFile("io/common/distribution.proto", fileDescriptor_737c63c5d3d253ea) }

var fileDescriptor_737c63c5d3d253ea = []byte{
	// 772 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0x46, 0x4e, 0xda, 0xc4, 0xc7, 0xb5, 0x2d, 0xaf, 0x4b, 0x11, 0xa1, 0x03, 0x19, 0x0f, 0x17,
	0x9e, 0x4e, 0x6c, 0x31, 0xe6, 0xe7, 0x8a, 0x19, 0x88, 0x1d, 0x7b, 0x6a, 0x9c, 0x3a, 0x1e, 0xc9,
	0x70, 0xc1, 0x4d, 0x67, 0x2d, 0x6d, 0xd4, 0x25, 0xd2, 0xae, 0xd1, 0xae, 0x68, 0x61, 0xb8, 0xe4,
	0x09, 0xb8, 0xe4, 0x11, 0x78, 0x08, 0x5e, 0x80, 0x97, 0xea, 0x68, 0x2d, 0xc9, 0xeb, 0xc8, 0xee,
	0xf4, 0x72, 0xbf, 0xf3, 0x7d, 0xe7, 0x7c, 0xe7, 0x1c, 0xed, 0x0a, 0x9e, 0x52, 0x6e, 0x7b, 0x3c,
	0x8a, 0x38, 0xb3, 0x7d, 0x2a, 0x64, 0x4c, 0x57, 0x89, 0xa4, 0x9c, 0xf5, 0xd7, 0x31, 0x97, 0x1c,
	0x55, 0x28, 0x3f, 0xd3, 0x18, 0x21, 0xf7, 0x70, 0x48, 0xff, 0xc0, 0x5b, 0xc6, 0xd9, 0xc7, 0xdb,
	0xa8, 0x02, 0x3c, 0x1e, 0x8a, 0x2c, 0x74, 0xb1, 0x01, 0x7a, 0x01, 0x61, 0x3d, 0xf1, 0x1a, 0x07,
	0x01, 0x89, 0x6d, 0xbe, 0x4e, 0xb5, 0xc2, 0xc6, 0x8c, 0x71, 0xa9, 0xf2, 0x64, 0xec, 0xce, 0x5f,
	0x06, 0x58, 0xe3, 0x08, 0xd3, 0xf0, 0x4a, 0xb3, 0xe1, 0x90, 0x5f, 0x13, 0x22, 0x24, 0x6a, 0x40,
	0x85, 0xfa, 0x96, 0x71, 0x6e, 0x74, 0xab, 0x4e, 0x85, 0xfa, 0xe8, 0x02, 0x4e, 0xf3, 0x6a, 0x56,
	0xe5, 0xdc, 0xe8, 0x36, 0x06, 0x66, 0x9f, 0xf2, 0xfe, 0x02, 0x0b, 0xb1, 0xc8, 0x70, 0xa7, 0x60,
	0xa0, 0x67, 0x60, 0xe2, 0x50, 0x92, 0x98, 0x61, 0x49, 0x7f, 0x23, 0xaa, 0x88, 0x75, 0xa4, 0x72,
	0x95, 0xf0, 0xce, 0xff, 0x06, 0x3c, 0xd6, 0x1d, 0xb8, 0x44, 0x4a, 0xca, 0x02, 0x81, 0x66, 0xf0,
	0x58, 0x1f, 0xd0, 0xe8, 0x15, 0x66, 0x8c, 0x84, 0xc2, 0x32, 0xce, 0x8f, 0xba, 0x8d, 0xc1, 0x47,
	0x69, 0xf9, 0xab, 0x72, 0xdc, 0xd9, 0x2b, 0x42, 0x5f, 0xc3, 0xa3, 0xd7, 0x24, 0xf4, 0x78, 0x94,
	0xb9, 0x49, 0x7b, 0xa8, 0x0d, 0x5a, 0x69, 0x12, 0x05, 0x2c, 0x49, 0xb4, 0x0e, 0xb1, 0x24, 0xce,
	0x0e, 0x0d, 0xd9, 0x00, 0xd9, 0xd9, 0x8d, 0x84, 0x6a, 0xa1, 0x36, 0x68, 0xa6, 0x22, 0x37, 0x12,
	0x85, 0x44, 0xa3, 0x74, 0xfe, 0x3b, 0x81, 0xfa, 0x4e, 0x42, 0x64, 0xc1, 0x89, 0x48, 0x56, 0xbf,
	0x10, 0x4f, 0x66, 0xe3, 0xcc, 0x8f, 0xe8, 0x3b, 0x30, 0xb3, 0xfd, 0x12, 0xdf, 0xcd, 0x28, 0x1b,
	0x5f, 0xed, 0xb4, 0xc4, 0x75, 0x11, 0x93, 0x31, 0x65, 0x81, 0x53, 0x22, 0xa3, 0x2e, 0x34, 0x57,
	0xdc, 0xff, 0x7d, 0x49, 0xde, 0xc8, 0x11, 0x67, 0x92, 0x30, 0x99, 0x4d, 0xf9, 0x3e, 0x8c, 0x6e,
	0xc0, 0x2a, 0xd4, 0xc3, 0x7b, 0x92, 0xe3, 0xc3, 0x25, 0x0f, 0x8a, 0xf2, 0xd2, 0xcf, 0x65, 0x14,
	0xe6, 0x79, 0x1e, 0x6c, 0x4b, 0x6b, 0x70, 0xa9, 0xb4, 0x2e, 0x79, 0xf8, 0xbe, 0xa5, 0xf5, 0x84,
	0xdf, 0x42, 0xdd, 0xe3, 0xec, 0x96, 0x06, 0x49, 0xac, 0xbe, 0x67, 0xeb, 0x44, 0x65, 0x79, 0x52,
	0xec, 0x72, 0xa4, 0x47, 0x9d, 0x5d, 0x32, 0xfa, 0x14, 0x60, 0x95, 0x48, 0xc9, 0x59, 0xda, 0x8d,
	0x75, 0xaa, 0x3c, 0x6b, 0x88, 0x6a, 0xac, 0x38, 0x8d, 0x78, 0xc8, 0x63, 0xab, 0x9a, 0x35, 0xb6,
	0x0b, 0xa3, 0xaf, 0xe0, 0xc3, 0x0d, 0x34, 0xc4, 0xde, 0x5d, 0x10, 0xf3, 0x84, 0xf9, 0x1b, 0x3e,
	0x28, 0xfe, 0xfe, 0x20, 0xea, 0x03, 0xca, 0x02, 0x3c, 0xf6, 0x49, 0xec, 0x60, 0x9f, 0x26, 0xc2,
	0xaa, 0x29, 0xc9, 0x9e, 0x08, 0xba, 0x80, 0xd6, 0x2d, 0xe7, 0x92, 0xc4, 0xfa, 0xca, 0x1e, 0x29,
	0x7a, 0x39, 0x80, 0x5c, 0x38, 0x2b, 0xe6, 0x36, 0x29, 0xc9, 0xea, 0x87, 0xc7, 0xfd, 0x0e, 0xd9,
	0xd6, 0x82, 0xbe, 0xba, 0x86, 0x6e, 0x41, 0x5f, 0x4f, 0xd9, 0x82, 0x2e, 0x6b, 0xbe, 0xbf, 0x05,
	0x3d, 0xe9, 0x37, 0xf0, 0x24, 0x22, 0x42, 0xe0, 0x80, 0xdc, 0x1f, 0xb6, 0xa9, 0x7c, 0x1c, 0x88,
	0xa2, 0x2f, 0xa0, 0xbd, 0xde, 0x23, 0x6a, 0x29, 0xd1, 0xbe, 0x50, 0xe7, 0x1f, 0x03, 0x50, 0xf9,
	0x2b, 0x42, 0x4f, 0xa1, 0x4a, 0x52, 0x74, 0x12, 0xf3, 0x28, 0xbb, 0xc7, 0x5b, 0x00, 0x7d, 0x0e,
	0xf5, 0xe2, 0x30, 0xc7, 0x11, 0x51, 0xd7, 0xb8, 0xea, 0xec, 0x82, 0xe8, 0x7b, 0xf8, 0x64, 0x9c,
	0x03, 0x3f, 0x91, 0x98, 0xde, 0xd2, 0xb4, 0xd5, 0xd8, 0x25, 0xcc, 0xa7, 0x2c, 0x50, 0x57, 0xf7,
	0xd4, 0x79, 0x17, 0xa5, 0xf3, 0x27, 0xd4, 0xb4, 0x87, 0x07, 0x7d, 0x06, 0x27, 0x5e, 0x36, 0x57,
	0x65, 0x69, 0xf8, 0xe0, 0xef, 0xcb, 0xca, 0x9b, 0x89, 0x93, 0xa3, 0x68, 0xa2, 0xbd, 0x30, 0xf9,
	0x06, 0x0e, 0xbf, 0x30, 0xb9, 0xbc, 0xa4, 0xf9, 0xe1, 0xf8, 0xf4, 0xc8, 0x3c, 0x7e, 0x36, 0x87,
	0xf6, 0x9e, 0x07, 0x17, 0xb5, 0xa1, 0x39, 0xbf, 0x79, 0x79, 0x35, 0x75, 0x97, 0xce, 0x74, 0xf8,
	0xe3, 0x72, 0x7a, 0x33, 0x37, 0x3f, 0x40, 0x2d, 0xa8, 0x8f, 0x9e, 0x5f, 0xce, 0xe7, 0xe3, 0xeb,
	0x97, 0xe3, 0x17, 0x97, 0xd3, 0x6b, 0xd3, 0x40, 0x4d, 0xa8, 0xe5, 0x90, 0xfb, 0xc2, 0x35, 0x2b,
	0xc3, 0x21, 0x34, 0x29, 0xef, 0xaf, 0xb1, 0x10, 0x77, 0x54, 0xf6, 0x17, 0x33, 0xca, 0x7f, 0xee,
	0x06, 0x54, 0xbe, 0x4a, 0x56, 0x7d, 0x8f, 0x47, 0x76, 0xfa, 0x6f, 0x99, 0x51, 0x69, 0x67, 0x84,
	0x5e, 0xc0, 0x43, 0xcc, 0x82, 0x9e, 0xf0, 0xef, 0x6c, 0xca, 0xff, 0xad, 0x1c, 0x2f, 0x66, 0x53,
	0xbe, 0x7a, 0xa8, 0xfe, 0x39, 0x5f, 0xbe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x45, 0x6e, 0x59, 0xee,
	0x56, 0x07, 0x00, 0x00,
}
