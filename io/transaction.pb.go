// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/common/transaction.proto

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

// Some point based loyalty programs will pass on transaction information (can be used for segmenting, additional reporting, and/or showing transactions on back of the pass).
type Transaction struct {
	// Reference ID is the ID used in the system where the transaction is coming from. Needs to be unique within the program.
	ReferenceId string `protobuf:"bytes,1,opt,name=referenceId,proto3" json:"referenceId,omitempty"`
	// The total amount of all order items. Based on POS setting, the totalPrice can already include the tax amount
	TotalPrice float32 `protobuf:"fixed32,2,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	// List of order items in the  transaction
	OrderItems []*OrderItem `protobuf:"bytes,3,rep,name=orderItems,proto3" json:"orderItems,omitempty"`
	// The total discount amount
	Discount float32 `protobuf:"fixed32,4,opt,name=discount,proto3" json:"discount,omitempty"`
	// List of discount items
	DiscountItems []*DiscountItem `protobuf:"bytes,5,rep,name=discountItems,proto3" json:"discountItems,omitempty"`
	// The service charge amount
	ServiceCharge float32 `protobuf:"fixed32,6,opt,name=serviceCharge,proto3" json:"serviceCharge,omitempty"`
	// The tax amount.
	TotalTax float32 `protobuf:"fixed32,7,opt,name=totalTax,proto3" json:"totalTax,omitempty"`
	// The final price of the transaction: (total price + service charge + (totalTax: if POS settings set to tax exclusive)) - discount
	FinalPrice float32 `protobuf:"fixed32,8,opt,name=finalPrice,proto3" json:"finalPrice,omitempty"`
	// Rounding difference
	RoundingDifference float32 `protobuf:"fixed32,9,opt,name=roundingDifference,proto3" json:"roundingDifference,omitempty"`
	// Indicates if this transaction is a refund
	IsRefunded bool `protobuf:"varint,10,opt,name=isRefunded,proto3" json:"isRefunded,omitempty"`
	// The transaction timestamp
	Timestamp *Date `protobuf:"bytes,11,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The transaction currency
	Currency string `protobuf:"bytes,12,opt,name=currency,proto3" json:"currency,omitempty"`
	// GPS details of check in
	Location *GPSLocation `protobuf:"bytes,13,opt,name=location,proto3" json:"location,omitempty"`
	// 17 is reserved for transaction source (online, or restaurant code)
	TransactionSource    string   `protobuf:"bytes,14,opt,name=transactionSource,proto3" json:"transactionSource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_1531df266e8e06cf, []int{0}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetReferenceId() string {
	if m != nil {
		return m.ReferenceId
	}
	return ""
}

func (m *Transaction) GetTotalPrice() float32 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func (m *Transaction) GetOrderItems() []*OrderItem {
	if m != nil {
		return m.OrderItems
	}
	return nil
}

func (m *Transaction) GetDiscount() float32 {
	if m != nil {
		return m.Discount
	}
	return 0
}

func (m *Transaction) GetDiscountItems() []*DiscountItem {
	if m != nil {
		return m.DiscountItems
	}
	return nil
}

func (m *Transaction) GetServiceCharge() float32 {
	if m != nil {
		return m.ServiceCharge
	}
	return 0
}

func (m *Transaction) GetTotalTax() float32 {
	if m != nil {
		return m.TotalTax
	}
	return 0
}

func (m *Transaction) GetFinalPrice() float32 {
	if m != nil {
		return m.FinalPrice
	}
	return 0
}

func (m *Transaction) GetRoundingDifference() float32 {
	if m != nil {
		return m.RoundingDifference
	}
	return 0
}

func (m *Transaction) GetIsRefunded() bool {
	if m != nil {
		return m.IsRefunded
	}
	return false
}

func (m *Transaction) GetTimestamp() *Date {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Transaction) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Transaction) GetLocation() *GPSLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Transaction) GetTransactionSource() string {
	if m != nil {
		return m.TransactionSource
	}
	return ""
}

type DiscountItem struct {
	// Discount code
	DiscountCode string `protobuf:"bytes,1,opt,name=discountCode,proto3" json:"discountCode,omitempty"`
	// Specific voucher code
	VoucherCode string `protobuf:"bytes,2,opt,name=voucherCode,proto3" json:"voucherCode,omitempty"`
	// The discount amount
	Amount float32 `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	// The discount item name
	ItemName             string   `protobuf:"bytes,5,opt,name=itemName,proto3" json:"itemName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscountItem) Reset()         { *m = DiscountItem{} }
func (m *DiscountItem) String() string { return proto.CompactTextString(m) }
func (*DiscountItem) ProtoMessage()    {}
func (*DiscountItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_1531df266e8e06cf, []int{1}
}

func (m *DiscountItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscountItem.Unmarshal(m, b)
}
func (m *DiscountItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscountItem.Marshal(b, m, deterministic)
}
func (m *DiscountItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscountItem.Merge(m, src)
}
func (m *DiscountItem) XXX_Size() int {
	return xxx_messageInfo_DiscountItem.Size(m)
}
func (m *DiscountItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscountItem.DiscardUnknown(m)
}

var xxx_messageInfo_DiscountItem proto.InternalMessageInfo

func (m *DiscountItem) GetDiscountCode() string {
	if m != nil {
		return m.DiscountCode
	}
	return ""
}

func (m *DiscountItem) GetVoucherCode() string {
	if m != nil {
		return m.VoucherCode
	}
	return ""
}

func (m *DiscountItem) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *DiscountItem) GetItemName() string {
	if m != nil {
		return m.ItemName
	}
	return ""
}

type OrderItem struct {
	// The item price
	Amount float32 `protobuf:"fixed32,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// Tax on the item
	Tax float32 `protobuf:"fixed32,2,opt,name=tax,proto3" json:"tax,omitempty"`
	// The item name
	ItemName string `protobuf:"bytes,3,opt,name=itemName,proto3" json:"itemName,omitempty"`
	// The quantity of items ordered
	Quantity int32 `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// The SKU as used in the system generating the transaction
	Sku                  string   `protobuf:"bytes,5,opt,name=sku,proto3" json:"sku,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderItem) Reset()         { *m = OrderItem{} }
func (m *OrderItem) String() string { return proto.CompactTextString(m) }
func (*OrderItem) ProtoMessage()    {}
func (*OrderItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_1531df266e8e06cf, []int{2}
}

func (m *OrderItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderItem.Unmarshal(m, b)
}
func (m *OrderItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderItem.Marshal(b, m, deterministic)
}
func (m *OrderItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderItem.Merge(m, src)
}
func (m *OrderItem) XXX_Size() int {
	return xxx_messageInfo_OrderItem.Size(m)
}
func (m *OrderItem) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderItem.DiscardUnknown(m)
}

var xxx_messageInfo_OrderItem proto.InternalMessageInfo

func (m *OrderItem) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *OrderItem) GetTax() float32 {
	if m != nil {
		return m.Tax
	}
	return 0
}

func (m *OrderItem) GetItemName() string {
	if m != nil {
		return m.ItemName
	}
	return ""
}

func (m *OrderItem) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *OrderItem) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func init() {
	proto.RegisterType((*Transaction)(nil), "io.Transaction")
	proto.RegisterType((*DiscountItem)(nil), "io.DiscountItem")
	proto.RegisterType((*OrderItem)(nil), "io.OrderItem")
}

func init() {
	proto.RegisterFile("io/common/transaction.proto", fileDescriptor_1531df266e8e06cf)
}

var fileDescriptor_1531df266e8e06cf = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0x4d, 0x6f, 0x13, 0x3d,
	0x10, 0xc7, 0xb5, 0x79, 0xe9, 0x93, 0x38, 0xcd, 0xd3, 0xe2, 0x03, 0x5a, 0x8a, 0x54, 0x59, 0x51,
	0x85, 0x22, 0x68, 0xb3, 0x52, 0x91, 0x38, 0x70, 0x40, 0x82, 0x56, 0x42, 0xa5, 0x08, 0xa2, 0x6d,
	0x4f, 0x5c, 0x90, 0xe3, 0xf5, 0x6e, 0x4c, 0x62, 0x4f, 0xb0, 0xbd, 0x7d, 0x39, 0xf1, 0x0d, 0x38,
	0x70, 0x46, 0xe2, 0xc0, 0x8d, 0x6f, 0xc2, 0xb5, 0x9f, 0x08, 0xd9, 0xfb, 0x92, 0x8d, 0xe0, 0x94,
	0x99, 0xff, 0x4c, 0x7e, 0x1e, 0x7b, 0xfe, 0x09, 0x7a, 0x28, 0x20, 0x62, 0x20, 0x25, 0xa8, 0xc8,
	0x6a, 0xaa, 0x0c, 0x65, 0x56, 0x80, 0x9a, 0xac, 0x34, 0x58, 0xc0, 0x2d, 0x01, 0x7b, 0x87, 0x3e,
	0x64, 0x47, 0x19, 0x57, 0x47, 0xe6, 0x9a, 0x66, 0x19, 0xd7, 0x11, 0xac, 0x5c, 0x93, 0x89, 0xa8,
	0x52, 0x60, 0xa9, 0x8f, 0x8b, 0x6f, 0xec, 0xed, 0xaf, 0x71, 0xc5, 0xc7, 0x47, 0x98, 0x7d, 0xe2,
	0xcc, 0x56, 0xf5, 0x07, 0xeb, 0xfa, 0x4a, 0xc3, 0x8d, 0x90, 0xc2, 0xde, 0x16, 0xa5, 0xd1, 0xef,
	0x2e, 0x1a, 0x5c, 0xae, 0x47, 0xc0, 0x04, 0x0d, 0x34, 0x4f, 0xb9, 0xe6, 0x8a, 0xf1, 0xb3, 0x24,
	0x0c, 0x48, 0x30, 0xee, 0xc7, 0x4d, 0x09, 0xef, 0x23, 0x64, 0xc1, 0xd2, 0xe5, 0x54, 0x0b, 0xc6,
	0xc3, 0x16, 0x09, 0xc6, 0xad, 0xb8, 0xa1, 0xe0, 0x23, 0x84, 0x40, 0x27, 0x5c, 0x9f, 0x59, 0x2e,
	0x4d, 0xd8, 0x26, 0xed, 0xf1, 0xe0, 0x78, 0x38, 0x11, 0x30, 0x79, 0x5f, 0xa9, 0x71, 0xa3, 0x01,
	0xef, 0xa1, 0x5e, 0x22, 0x0c, 0x83, 0x5c, 0xd9, 0xb0, 0xe3, 0x61, 0x75, 0x8e, 0x9f, 0xa1, 0x61,
	0x15, 0x17, 0xb4, 0xae, 0xa7, 0xed, 0x3a, 0xda, 0x69, 0xa3, 0x10, 0x6f, 0xb6, 0xe1, 0x03, 0x34,
	0x34, 0x5c, 0x5f, 0x09, 0xc6, 0x4f, 0xe6, 0x54, 0x67, 0x3c, 0xdc, 0xf2, 0xe0, 0x4d, 0xd1, 0x9d,
	0xec, 0xc7, 0xbe, 0xa4, 0x37, 0xe1, 0x7f, 0xc5, 0xc9, 0x55, 0xee, 0x2e, 0x99, 0x0a, 0x55, 0x5d,
	0xb2, 0x57, 0x5c, 0x72, 0xad, 0xe0, 0x09, 0xc2, 0x1a, 0x72, 0x95, 0x08, 0x95, 0x9d, 0x8a, 0xb4,
	0x7c, 0x9c, 0xb0, 0xef, 0xfb, 0xfe, 0x51, 0x71, 0x3c, 0x61, 0x62, 0x9e, 0xe6, 0x2a, 0xe1, 0x49,
	0x88, 0x48, 0x30, 0xee, 0xc5, 0x0d, 0x05, 0x3f, 0x42, 0x7d, 0x2b, 0x24, 0x37, 0x96, 0xca, 0x55,
	0x38, 0x20, 0xc1, 0x78, 0x70, 0xdc, 0xf3, 0xb7, 0xa4, 0x96, 0xc7, 0xeb, 0x92, 0x9b, 0x99, 0xe5,
	0xda, 0x31, 0x6f, 0xc3, 0x6d, 0xbf, 0x9b, 0x3a, 0xc7, 0x4f, 0x50, 0x6f, 0x09, 0xcc, 0x1b, 0x23,
	0x1c, 0x7a, 0xc4, 0x8e, 0x43, 0xbc, 0x9e, 0x5e, 0xbc, 0x2d, 0xe5, 0xb8, 0x6e, 0xc0, 0x87, 0xe8,
	0x5e, 0xc3, 0x79, 0x17, 0x90, 0x6b, 0xc6, 0xc3, 0xff, 0x3d, 0xf1, 0xef, 0xc2, 0xf3, 0x9f, 0xc1,
	0xb7, 0x97, 0x3f, 0x02, 0xf4, 0x3d, 0x78, 0xdc, 0x34, 0xcb, 0xb1, 0x6c, 0x24, 0x44, 0xa8, 0x14,
	0xb4, 0xf4, 0x74, 0x92, 0x82, 0x26, 0x92, 0xcb, 0x19, 0xd7, 0x64, 0xa5, 0x21, 0xd3, 0x54, 0x1a,
	0x62, 0xe7, 0xd4, 0x92, 0x6b, 0xaa, 0x2c, 0xb1, 0x40, 0xcc, 0x1c, 0xae, 0xc9, 0x92, 0x5a, 0x6e,
	0x2c, 0x69, 0x1c, 0x67, 0x08, 0x28, 0x32, 0xa3, 0x6c, 0x41, 0x20, 0x25, 0x76, 0xce, 0x2b, 0x08,
	0xa3, 0x3a, 0x99, 0xdc, 0x05, 0xeb, 0x87, 0xb8, 0x0b, 0xb6, 0xa8, 0x74, 0x0b, 0xbf, 0x0b, 0xea,
	0x07, 0x18, 0x7d, 0x0d, 0xd0, 0x76, 0xd3, 0x16, 0x78, 0x84, 0xb6, 0x2b, 0x63, 0x9c, 0x40, 0xc2,
	0x4b, 0x37, 0x6f, 0x68, 0xce, 0xf0, 0x57, 0x90, 0xb3, 0x39, 0xd7, 0xbe, 0xa5, 0x55, 0x18, 0xbe,
	0x21, 0xe1, 0xfb, 0xa8, 0x3c, 0xab, 0xf4, 0x67, 0x99, 0xb9, 0x5d, 0x08, 0xcb, 0xe5, 0x3b, 0x2a,
	0x79, 0xd8, 0x2d, 0x76, 0x51, 0xe5, 0x6f, 0x3a, 0xbd, 0xf6, 0x6e, 0x67, 0xf4, 0x05, 0xf5, 0x6b,
	0xd3, 0x37, 0x30, 0xc1, 0x06, 0x66, 0x17, 0xb5, 0x2d, 0xbd, 0x29, 0x7f, 0x48, 0x2e, 0xdc, 0x00,
	0xb7, 0x37, 0xc1, 0xae, 0xf6, 0x39, 0xa7, 0xca, 0x0a, 0x7b, 0xeb, 0xc7, 0xe9, 0xc6, 0x75, 0xee,
	0x48, 0x66, 0x91, 0x97, 0xb3, 0xb8, 0xf0, 0xd5, 0x0b, 0xb4, 0x23, 0x60, 0xb2, 0xa2, 0xc6, 0x2c,
	0x84, 0x9d, 0x4c, 0xcf, 0x05, 0x7c, 0x38, 0x30, 0x96, 0x9a, 0x79, 0xad, 0x31, 0x90, 0x91, 0x80,
	0x48, 0x42, 0xc2, 0x97, 0x91, 0x49, 0x16, 0x51, 0x06, 0x91, 0x80, 0x5f, 0xad, 0xce, 0xf4, 0xfc,
	0x0c, 0x66, 0x5b, 0xfe, 0x4f, 0xe2, 0xe9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x58, 0x9e,
	0x2f, 0xb0, 0x04, 0x00, 0x00,
}
