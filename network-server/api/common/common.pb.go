// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package common

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

type Modulation int32

const (
	// LoRa
	Modulation_LORA Modulation = 0
	// FSK
	Modulation_FSK Modulation = 1
)

var Modulation_name = map[int32]string{
	0: "LORA",
	1: "FSK",
}

var Modulation_value = map[string]int32{
	"LORA": 0,
	"FSK":  1,
}

func (x Modulation) String() string {
	return proto.EnumName(Modulation_name, int32(x))
}

func (Modulation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type Region int32

const (
	// EU868
	Region_EU868 Region = 0
	// US915
	Region_US915 Region = 2
	// CN779
	Region_CN779 Region = 3
	// EU433
	Region_EU433 Region = 4
	// AU915
	Region_AU915 Region = 5
	// CN470
	Region_CN470 Region = 6
	// AS923
	Region_AS923 Region = 7
	// KR920
	Region_KR920 Region = 8
	// IN865
	Region_IN865 Region = 9
	// RU864
	Region_RU864 Region = 10
)

var Region_name = map[int32]string{
	0:  "EU868",
	2:  "US915",
	3:  "CN779",
	4:  "EU433",
	5:  "AU915",
	6:  "CN470",
	7:  "AS923",
	8:  "KR920",
	9:  "IN865",
	10: "RU864",
}

var Region_value = map[string]int32{
	"EU868": 0,
	"US915": 2,
	"CN779": 3,
	"EU433": 4,
	"AU915": 5,
	"CN470": 6,
	"AS923": 7,
	"KR920": 8,
	"IN865": 9,
	"RU864": 10,
}

func (x Region) String() string {
	return proto.EnumName(Region_name, int32(x))
}

func (Region) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

type LocationSource int32

const (
	// Unknown.
	LocationSource_UNKNOWN LocationSource = 0
	// GPS.
	LocationSource_GPS LocationSource = 1
	// Manually configured.
	LocationSource_CONFIG LocationSource = 2
	// Geo resolver.
	LocationSource_GEO_RESOLVER LocationSource = 3
)

var LocationSource_name = map[int32]string{
	0: "UNKNOWN",
	1: "GPS",
	2: "CONFIG",
	3: "GEO_RESOLVER",
}

var LocationSource_value = map[string]int32{
	"UNKNOWN":      0,
	"GPS":          1,
	"CONFIG":       2,
	"GEO_RESOLVER": 3,
}

func (x LocationSource) String() string {
	return proto.EnumName(LocationSource_name, int32(x))
}

func (LocationSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

type KeyEnvelope struct {
	// KEK label.
	KekLabel string `protobuf:"bytes,1,opt,name=kek_label,json=kekLabel,proto3" json:"kek_label,omitempty"`
	// AES key (when the kek_label is set, this key is encrypted using a key
	// known to the join-server and application-server.
	// For more information please refer to the LoRaWAN Backend Interface
	// 'Key Transport Security' section.
	AesKey               []byte   `protobuf:"bytes,2,opt,name=aes_key,json=aesKey,proto3" json:"aes_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyEnvelope) Reset()         { *m = KeyEnvelope{} }
func (m *KeyEnvelope) String() string { return proto.CompactTextString(m) }
func (*KeyEnvelope) ProtoMessage()    {}
func (*KeyEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *KeyEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyEnvelope.Unmarshal(m, b)
}
func (m *KeyEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyEnvelope.Marshal(b, m, deterministic)
}
func (m *KeyEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyEnvelope.Merge(m, src)
}
func (m *KeyEnvelope) XXX_Size() int {
	return xxx_messageInfo_KeyEnvelope.Size(m)
}
func (m *KeyEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_KeyEnvelope proto.InternalMessageInfo

func (m *KeyEnvelope) GetKekLabel() string {
	if m != nil {
		return m.KekLabel
	}
	return ""
}

func (m *KeyEnvelope) GetAesKey() []byte {
	if m != nil {
		return m.AesKey
	}
	return nil
}

type Location struct {
	// Latitude.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Longitude.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// Altitude.
	Altitude float64 `protobuf:"fixed64,3,opt,name=altitude,proto3" json:"altitude,omitempty"`
	// Location source.
	Source LocationSource `protobuf:"varint,4,opt,name=source,proto3,enum=common.LocationSource" json:"source,omitempty"`
	// Accuracy (in meters).
	Accuracy             uint32   `protobuf:"varint,5,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Location) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Location) GetSource() LocationSource {
	if m != nil {
		return m.Source
	}
	return LocationSource_UNKNOWN
}

func (m *Location) GetAccuracy() uint32 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

type ContextDetails struct {
	// Context ID (UUID).
	ContextId            []byte   `protobuf:"bytes,1,opt,name=context_id,json=contextId,proto3" json:"context_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContextDetails) Reset()         { *m = ContextDetails{} }
func (m *ContextDetails) String() string { return proto.CompactTextString(m) }
func (*ContextDetails) ProtoMessage()    {}
func (*ContextDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *ContextDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContextDetails.Unmarshal(m, b)
}
func (m *ContextDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContextDetails.Marshal(b, m, deterministic)
}
func (m *ContextDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContextDetails.Merge(m, src)
}
func (m *ContextDetails) XXX_Size() int {
	return xxx_messageInfo_ContextDetails.Size(m)
}
func (m *ContextDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ContextDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ContextDetails proto.InternalMessageInfo

func (m *ContextDetails) GetContextId() []byte {
	if m != nil {
		return m.ContextId
	}
	return nil
}

func init() {
	proto.RegisterEnum("common.Modulation", Modulation_name, Modulation_value)
	proto.RegisterEnum("common.Region", Region_name, Region_value)
	proto.RegisterEnum("common.LocationSource", LocationSource_name, LocationSource_value)
	proto.RegisterType((*KeyEnvelope)(nil), "common.KeyEnvelope")
	proto.RegisterType((*Location)(nil), "common.Location")
	proto.RegisterType((*ContextDetails)(nil), "common.ContextDetails")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0x4b, 0x6f, 0xda, 0x40,
	0x10, 0x8e, 0x79, 0x18, 0x7b, 0x42, 0xd1, 0x6a, 0x0f, 0x2d, 0xea, 0x43, 0x45, 0x9c, 0x10, 0x52,
	0x70, 0x8a, 0x79, 0x1e, 0x53, 0xe2, 0x20, 0x64, 0x6a, 0x57, 0x6b, 0xb9, 0x95, 0x7a, 0x41, 0x8b,
	0xd9, 0x52, 0xcb, 0x8e, 0x17, 0xf9, 0x41, 0xeb, 0xff, 0xd4, 0x1f, 0x59, 0xad, 0x97, 0xa4, 0xca,
	0xed, 0x7b, 0xcd, 0xec, 0xcc, 0x68, 0xa1, 0x1d, 0xf0, 0xc7, 0x47, 0x9e, 0x8c, 0x4e, 0x29, 0xcf,
	0x39, 0x56, 0x25, 0xeb, 0xaf, 0xe0, 0xda, 0x66, 0xa5, 0x95, 0x9c, 0x59, 0xcc, 0x4f, 0x0c, 0xbf,
	0x03, 0x3d, 0x62, 0xd1, 0x2e, 0xa6, 0x7b, 0x16, 0x77, 0x95, 0x9e, 0x32, 0xd0, 0x89, 0x16, 0xb1,
	0x68, 0x2b, 0x38, 0x7e, 0x03, 0x2d, 0xca, 0xb2, 0x5d, 0xc4, 0xca, 0x6e, 0xad, 0xa7, 0x0c, 0xda,
	0x44, 0xa5, 0x2c, 0xb3, 0x59, 0xd9, 0xff, 0xab, 0x80, 0xb6, 0xe5, 0x01, 0xcd, 0x43, 0x9e, 0xe0,
	0xb7, 0xa0, 0xc5, 0x34, 0x0f, 0xf3, 0xe2, 0xc0, 0xaa, 0x0e, 0x0a, 0x79, 0xe6, 0xf8, 0x3d, 0xe8,
	0x31, 0x4f, 0x8e, 0xd2, 0xac, 0x55, 0xe6, 0x7f, 0x41, 0x54, 0xd2, 0xf8, 0x52, 0x59, 0x97, 0x95,
	0x4f, 0x1c, 0x8f, 0x40, 0xcd, 0x78, 0x91, 0x06, 0xac, 0xdb, 0xe8, 0x29, 0x83, 0xce, 0xf8, 0xf5,
	0xe8, 0xb2, 0xce, 0xd3, 0xbb, 0x5e, 0xe5, 0x92, 0x4b, 0xaa, 0xea, 0x15, 0x04, 0x45, 0x4a, 0x83,
	0xb2, 0xdb, 0xec, 0x29, 0x83, 0x57, 0xe4, 0x99, 0xf7, 0x0d, 0xe8, 0xac, 0x78, 0x92, 0xb3, 0x3f,
	0xf9, 0x3d, 0xcb, 0x69, 0x18, 0x67, 0xf8, 0x03, 0x40, 0x20, 0x95, 0x5d, 0x78, 0xa8, 0xa6, 0x6e,
	0x13, 0xfd, 0xa2, 0x6c, 0x0e, 0xc3, 0x8f, 0x00, 0x5f, 0xf8, 0xa1, 0x88, 0xe5, 0x82, 0x1a, 0x34,
	0xb6, 0x2e, 0xb9, 0x43, 0x57, 0xb8, 0x05, 0xf5, 0x07, 0xcf, 0x46, 0xca, 0xf0, 0x0c, 0x2a, 0x61,
	0x47, 0x61, 0xea, 0xd0, 0xb4, 0xfc, 0xc5, 0x6c, 0x81, 0xae, 0x04, 0xf4, 0xbd, 0xe5, 0xa7, 0x29,
	0xaa, 0x09, 0xb8, 0x72, 0xe6, 0xf3, 0x25, 0xaa, 0xcb, 0xc0, 0xc4, 0x34, 0x51, 0x43, 0xc0, 0x3b,
	0x5f, 0x04, 0x9a, 0x32, 0x30, 0x99, 0xdf, 0x22, 0xb5, 0x52, 0xbd, 0xe5, 0xd8, 0x44, 0x2d, 0x01,
	0x6d, 0xb2, 0x1c, 0xdf, 0x22, 0x4d, 0xc0, 0x8d, 0xb3, 0x98, 0x4d, 0x91, 0x2e, 0x20, 0xf1, 0x17,
	0xb3, 0x09, 0x82, 0xe1, 0x3d, 0x74, 0x5e, 0xee, 0x8f, 0xaf, 0xa1, 0xe5, 0x3b, 0xb6, 0xe3, 0x7e,
	0x77, 0xe4, 0x7c, 0xeb, 0xaf, 0x1e, 0x52, 0x30, 0x80, 0xba, 0x72, 0x9d, 0x87, 0xcd, 0x1a, 0xd5,
	0x30, 0x82, 0xf6, 0xda, 0x72, 0x77, 0xc4, 0xf2, 0xdc, 0xed, 0x37, 0x8b, 0xa0, 0xfa, 0xe7, 0xe9,
	0x0f, 0xf3, 0x18, 0xe6, 0xbf, 0x8a, 0xbd, 0xb8, 0xa9, 0xb1, 0x4f, 0x79, 0x40, 0x69, 0x6a, 0xc4,
	0x3c, 0x0d, 0x7f, 0x96, 0x37, 0x09, 0xcb, 0x7f, 0xf3, 0x34, 0xba, 0xc9, 0x58, 0x7a, 0x66, 0xa9,
	0x41, 0x4f, 0xa1, 0x21, 0x2f, 0xbf, 0x57, 0xab, 0x9f, 0x64, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0x3c, 0xc5, 0x7d, 0x3d, 0x59, 0x02, 0x00, 0x00,
}
