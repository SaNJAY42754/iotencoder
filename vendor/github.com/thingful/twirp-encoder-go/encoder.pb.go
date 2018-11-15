// Code generated by protoc-gen-go. DO NOT EDIT.
// source: encoder.proto

package encoder

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

// An enumeration which allows us to express whether the device will be
// located indoors or outdoors when deployed.
type CreateStreamRequest_Exposure int32

const (
	CreateStreamRequest_UNKNOWN CreateStreamRequest_Exposure = 0
	CreateStreamRequest_INDOOR  CreateStreamRequest_Exposure = 1
	CreateStreamRequest_OUTDOOR CreateStreamRequest_Exposure = 2
)

var CreateStreamRequest_Exposure_name = map[int32]string{
	0: "UNKNOWN",
	1: "INDOOR",
	2: "OUTDOOR",
}
var CreateStreamRequest_Exposure_value = map[string]int32{
	"UNKNOWN": 0,
	"INDOOR":  1,
	"OUTDOOR": 2,
}

func (x CreateStreamRequest_Exposure) String() string {
	return proto.EnumName(CreateStreamRequest_Exposure_name, int32(x))
}
func (CreateStreamRequest_Exposure) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_encoder_e884ac1c0a530b11, []int{0, 0}
}

// An enumeration which allows us to specify what type of sharing is to be
// defined for the specified sensor type. The default value is `SHARE` which
// implies sharing the data at full resolution. If this type is specified, it
// is an error if either of `buckets` or `interval` is also supplied.
type CreateStreamRequest_Operation_Action int32

const (
	CreateStreamRequest_Operation_UNKNOWN    CreateStreamRequest_Operation_Action = 0
	CreateStreamRequest_Operation_SHARE      CreateStreamRequest_Operation_Action = 1
	CreateStreamRequest_Operation_BIN        CreateStreamRequest_Operation_Action = 2
	CreateStreamRequest_Operation_MOVING_AVG CreateStreamRequest_Operation_Action = 3
)

var CreateStreamRequest_Operation_Action_name = map[int32]string{
	0: "UNKNOWN",
	1: "SHARE",
	2: "BIN",
	3: "MOVING_AVG",
}
var CreateStreamRequest_Operation_Action_value = map[string]int32{
	"UNKNOWN":    0,
	"SHARE":      1,
	"BIN":        2,
	"MOVING_AVG": 3,
}

func (x CreateStreamRequest_Operation_Action) String() string {
	return proto.EnumName(CreateStreamRequest_Operation_Action_name, int32(x))
}
func (CreateStreamRequest_Operation_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_encoder_e884ac1c0a530b11, []int{0, 1, 0}
}

// CreateStreamRequest is the message sent in order to create a new encoded
// stream. As a result of this method call, the stream encoder will have
// configured a stream that receives messages, applies all defined entitlement
// operations, then encrypts the data and sends it on to the configured
// datastore.
type CreateStreamRequest struct {
	// The token that uniquely identifies the device. This is a required field.
	DeviceToken string `protobuf:"bytes,2,opt,name=device_token,json=deviceToken,proto3" json:"device_token,omitempty"`
	// The public key of the recipient, again this is used in order to encrypt
	// outgoing data, as well as being used to signify to the datastore the bucket
	// in which data should be stored. This is a required field.
	RecipientPublicKey string `protobuf:"bytes,3,opt,name=recipient_public_key,json=recipientPublicKey,proto3" json:"recipient_public_key,omitempty"`
	// The location of the device to be claimed.
	Location *CreateStreamRequest_Location `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	// The specific exposure of the device, i.e. is this instance indoors or
	// outdoors.
	Exposure CreateStreamRequest_Exposure `protobuf:"varint,6,opt,name=exposure,proto3,enum=decode.iot.encoder.CreateStreamRequest_Exposure" json:"exposure,omitempty"`
	// The entitlements field holds a repeated list of Operations which each
	// define a transformational function for a specific sensor id. If no
	// operations are submitted, we currently create a stream that writes
	// through all received channels without applying any processing transformations
	// to the data, but if this field contains any elements, the resulting stream
	// will only contain the specified sensor type.
	Operations           []*CreateStreamRequest_Operation `protobuf:"bytes,7,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CreateStreamRequest) Reset()         { *m = CreateStreamRequest{} }
func (m *CreateStreamRequest) String() string { return proto.CompactTextString(m) }
func (*CreateStreamRequest) ProtoMessage()    {}
func (*CreateStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_encoder_e884ac1c0a530b11, []int{0}
}
func (m *CreateStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStreamRequest.Unmarshal(m, b)
}
func (m *CreateStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStreamRequest.Marshal(b, m, deterministic)
}
func (dst *CreateStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStreamRequest.Merge(dst, src)
}
func (m *CreateStreamRequest) XXX_Size() int {
	return xxx_messageInfo_CreateStreamRequest.Size(m)
}
func (m *CreateStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStreamRequest proto.InternalMessageInfo

func (m *CreateStreamRequest) GetDeviceToken() string {
	if m != nil {
		return m.DeviceToken
	}
	return ""
}

func (m *CreateStreamRequest) GetRecipientPublicKey() string {
	if m != nil {
		return m.RecipientPublicKey
	}
	return ""
}

func (m *CreateStreamRequest) GetLocation() *CreateStreamRequest_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *CreateStreamRequest) GetExposure() CreateStreamRequest_Exposure {
	if m != nil {
		return m.Exposure
	}
	return CreateStreamRequest_UNKNOWN
}

func (m *CreateStreamRequest) GetOperations() []*CreateStreamRequest_Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// A nested type capturing the location of the device expressed via decimal
// long/lat pair.
type CreateStreamRequest_Location struct {
	// The longitude expressed as a decimal.
	Longitude float64 `protobuf:"fixed64,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// The latitude expressed as a decimal.
	Latitude             float64  `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateStreamRequest_Location) Reset()         { *m = CreateStreamRequest_Location{} }
func (m *CreateStreamRequest_Location) String() string { return proto.CompactTextString(m) }
func (*CreateStreamRequest_Location) ProtoMessage()    {}
func (*CreateStreamRequest_Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_encoder_e884ac1c0a530b11, []int{0, 0}
}
func (m *CreateStreamRequest_Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStreamRequest_Location.Unmarshal(m, b)
}
func (m *CreateStreamRequest_Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStreamRequest_Location.Marshal(b, m, deterministic)
}
func (dst *CreateStreamRequest_Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStreamRequest_Location.Merge(dst, src)
}
func (m *CreateStreamRequest_Location) XXX_Size() int {
	return xxx_messageInfo_CreateStreamRequest_Location.Size(m)
}
func (m *CreateStreamRequest_Location) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStreamRequest_Location.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStreamRequest_Location proto.InternalMessageInfo

func (m *CreateStreamRequest_Location) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *CreateStreamRequest_Location) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

// A nested type which is used to capture a list of specific operations we
// perform the stream.
type CreateStreamRequest_Operation struct {
	// The unique id of the sensor type for which this specific configuration is
	// defined. This is a required field.
	SensorId uint32 `protobuf:"varint,1,opt,name=sensor_id,json=sensorId,proto3" json:"sensor_id,omitempty"`
	// The specific action this entitlement defines for the sensor type. This is a
	// required field.
	Action CreateStreamRequest_Operation_Action `protobuf:"varint,2,opt,name=action,proto3,enum=decode.iot.encoder.CreateStreamRequest_Operation_Action" json:"action,omitempty"`
	// The bins attribute is used to specify the the bins into which incoming
	// values should be classified. Each element in the list is the upper
	// inclusive bound of a bin. The values submitted must be sorted in strictly
	// increasing order. There is no need to add a highest bin with +Inf bound, it
	// will be added implicitly. This field is optional unless an Action of `BIN`
	// has been requested, in which case it is required. It is an error to send
	// values for this attribute unless the value of Action is `BIN`.
	Bins []float64 `protobuf:"fixed64,3,rep,packed,name=bins,proto3" json:"bins,omitempty"`
	// This attribute is used to control the entitlement in the case for which we
	// have specified an action type representing a moving average. It represents
	// the interval in seconds over which the moving average should be calculated,
	// e.g. for a 15 minute moving average the value supplied here would be 900.
	// This field is optional unless an Action of `MOVING_AVG` has been specified,
	// in which case it is required. It is an error to send a value for this
	// attribute unless the value of Action is `MOVING_AVG`.
	Interval             uint32   `protobuf:"varint,4,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateStreamRequest_Operation) Reset()         { *m = CreateStreamRequest_Operation{} }
func (m *CreateStreamRequest_Operation) String() string { return proto.CompactTextString(m) }
func (*CreateStreamRequest_Operation) ProtoMessage()    {}
func (*CreateStreamRequest_Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_encoder_e884ac1c0a530b11, []int{0, 1}
}
func (m *CreateStreamRequest_Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStreamRequest_Operation.Unmarshal(m, b)
}
func (m *CreateStreamRequest_Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStreamRequest_Operation.Marshal(b, m, deterministic)
}
func (dst *CreateStreamRequest_Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStreamRequest_Operation.Merge(dst, src)
}
func (m *CreateStreamRequest_Operation) XXX_Size() int {
	return xxx_messageInfo_CreateStreamRequest_Operation.Size(m)
}
func (m *CreateStreamRequest_Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStreamRequest_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStreamRequest_Operation proto.InternalMessageInfo

func (m *CreateStreamRequest_Operation) GetSensorId() uint32 {
	if m != nil {
		return m.SensorId
	}
	return 0
}

func (m *CreateStreamRequest_Operation) GetAction() CreateStreamRequest_Operation_Action {
	if m != nil {
		return m.Action
	}
	return CreateStreamRequest_Operation_UNKNOWN
}

func (m *CreateStreamRequest_Operation) GetBins() []float64 {
	if m != nil {
		return m.Bins
	}
	return nil
}

func (m *CreateStreamRequest_Operation) GetInterval() uint32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

// CreateStreamResponse is the message returned from the stream encoder after it
// successfully creates a stream. The device registration service should keep a
// record of this value so that it is able to delete the stream if required.
type CreateStreamResponse struct {
	// An identifier for the stream which can be used in order to delete a stream
	// when required.
	StreamUid string `protobuf:"bytes,1,opt,name=stream_uid,json=streamUid,proto3" json:"stream_uid,omitempty"`
	// A secret token passed back to the caller which it must keep secret, in
	// order to be permitted to delete the stream.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateStreamResponse) Reset()         { *m = CreateStreamResponse{} }
func (m *CreateStreamResponse) String() string { return proto.CompactTextString(m) }
func (*CreateStreamResponse) ProtoMessage()    {}
func (*CreateStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_encoder_e884ac1c0a530b11, []int{1}
}
func (m *CreateStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStreamResponse.Unmarshal(m, b)
}
func (m *CreateStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStreamResponse.Marshal(b, m, deterministic)
}
func (dst *CreateStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStreamResponse.Merge(dst, src)
}
func (m *CreateStreamResponse) XXX_Size() int {
	return xxx_messageInfo_CreateStreamResponse.Size(m)
}
func (m *CreateStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStreamResponse proto.InternalMessageInfo

func (m *CreateStreamResponse) GetStreamUid() string {
	if m != nil {
		return m.StreamUid
	}
	return ""
}

func (m *CreateStreamResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// DeleteStreamRequest is the message sent to the encoder in order to delete a
// configured stream. Sending this message must delete the MQTT subscription, as
// well as deleting all encryption credentials stored on the encoder.
type DeleteStreamRequest struct {
	// The identifier for the stream to be deleted. This is a required field.
	StreamUid string `protobuf:"bytes,1,opt,name=stream_uid,json=streamUid,proto3" json:"stream_uid,omitempty"`
	// The secret token that was returned to the caller when creating the stream.
	// This is a required field, and must match the value stored internally for
	// the stream.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteStreamRequest) Reset()         { *m = DeleteStreamRequest{} }
func (m *DeleteStreamRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteStreamRequest) ProtoMessage()    {}
func (*DeleteStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_encoder_e884ac1c0a530b11, []int{2}
}
func (m *DeleteStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteStreamRequest.Unmarshal(m, b)
}
func (m *DeleteStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteStreamRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteStreamRequest.Merge(dst, src)
}
func (m *DeleteStreamRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteStreamRequest.Size(m)
}
func (m *DeleteStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteStreamRequest proto.InternalMessageInfo

func (m *DeleteStreamRequest) GetStreamUid() string {
	if m != nil {
		return m.StreamUid
	}
	return ""
}

func (m *DeleteStreamRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// DeleteStreamResponse is a placeholder response message on a successful
// deletion of stream on the encoder.
type DeleteStreamResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteStreamResponse) Reset()         { *m = DeleteStreamResponse{} }
func (m *DeleteStreamResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteStreamResponse) ProtoMessage()    {}
func (*DeleteStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_encoder_e884ac1c0a530b11, []int{3}
}
func (m *DeleteStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteStreamResponse.Unmarshal(m, b)
}
func (m *DeleteStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteStreamResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteStreamResponse.Merge(dst, src)
}
func (m *DeleteStreamResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteStreamResponse.Size(m)
}
func (m *DeleteStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteStreamResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateStreamRequest)(nil), "decode.iot.encoder.CreateStreamRequest")
	proto.RegisterType((*CreateStreamRequest_Location)(nil), "decode.iot.encoder.CreateStreamRequest.Location")
	proto.RegisterType((*CreateStreamRequest_Operation)(nil), "decode.iot.encoder.CreateStreamRequest.Operation")
	proto.RegisterType((*CreateStreamResponse)(nil), "decode.iot.encoder.CreateStreamResponse")
	proto.RegisterType((*DeleteStreamRequest)(nil), "decode.iot.encoder.DeleteStreamRequest")
	proto.RegisterType((*DeleteStreamResponse)(nil), "decode.iot.encoder.DeleteStreamResponse")
	proto.RegisterEnum("decode.iot.encoder.CreateStreamRequest_Exposure", CreateStreamRequest_Exposure_name, CreateStreamRequest_Exposure_value)
	proto.RegisterEnum("decode.iot.encoder.CreateStreamRequest_Operation_Action", CreateStreamRequest_Operation_Action_name, CreateStreamRequest_Operation_Action_value)
}

func init() { proto.RegisterFile("encoder.proto", fileDescriptor_encoder_e884ac1c0a530b11) }

var fileDescriptor_encoder_e884ac1c0a530b11 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xd9, 0xb8, 0xf9, 0xe3, 0x49, 0x5b, 0x45, 0xd3, 0x08, 0x59, 0x01, 0x24, 0x93, 0x0b,
	0x3e, 0x59, 0x25, 0x5c, 0xe0, 0x98, 0x92, 0xa8, 0x84, 0x14, 0xbb, 0x6c, 0x9b, 0x22, 0x71, 0xb1,
	0x1c, 0x7b, 0x84, 0x56, 0x35, 0x5e, 0x63, 0x6f, 0x2a, 0xfa, 0x68, 0x3c, 0x05, 0x4f, 0xc1, 0x7b,
	0xa0, 0xac, 0x9d, 0x90, 0x40, 0x24, 0x02, 0x37, 0xcf, 0xf7, 0xed, 0xfc, 0xf6, 0xdb, 0x59, 0xdb,
	0x70, 0x44, 0x69, 0x24, 0x63, 0xca, 0xdd, 0x2c, 0x97, 0x4a, 0x22, 0xc6, 0xb4, 0x2c, 0x5d, 0x21,
	0x95, 0x5b, 0x39, 0xfd, 0x6f, 0x75, 0x38, 0x79, 0x9d, 0x53, 0xa8, 0xe8, 0x4a, 0xe5, 0x14, 0x7e,
	0xe6, 0xf4, 0x65, 0x41, 0x85, 0xc2, 0xa7, 0x70, 0x18, 0xd3, 0x9d, 0x88, 0x28, 0x50, 0xf2, 0x96,
	0x52, 0xab, 0x66, 0x33, 0xc7, 0xe4, 0xed, 0x52, 0xbb, 0x5e, 0x4a, 0x78, 0x0a, 0xdd, 0x9c, 0x22,
	0x91, 0x09, 0x4a, 0x55, 0x90, 0x2d, 0xe6, 0x89, 0x88, 0x82, 0x5b, 0xba, 0xb7, 0x0c, 0xbd, 0x14,
	0xd7, 0xde, 0xa5, 0xb6, 0xa6, 0x74, 0x8f, 0x17, 0xd0, 0x4a, 0x64, 0x14, 0x2a, 0x21, 0x53, 0xab,
	0x6e, 0x33, 0xa7, 0x3d, 0x38, 0x75, 0xff, 0xcc, 0xe4, 0xee, 0xc8, 0xe3, 0x5e, 0x54, 0x7d, 0x7c,
	0x4d, 0x58, 0xd2, 0xe8, 0x6b, 0x26, 0x8b, 0x45, 0x4e, 0x56, 0xc3, 0x66, 0xce, 0xf1, 0xfe, 0xb4,
	0x71, 0xd5, 0xc7, 0xd7, 0x04, 0x7c, 0x0f, 0x20, 0x33, 0xca, 0x35, 0xba, 0xb0, 0x9a, 0xb6, 0xe1,
	0xb4, 0x07, 0xcf, 0xf7, 0xe5, 0xf9, 0xab, 0x4e, 0xbe, 0x01, 0xe9, 0x8d, 0xa0, 0xb5, 0x8a, 0x8d,
	0x8f, 0xc1, 0x4c, 0x64, 0xfa, 0x49, 0xa8, 0x45, 0x4c, 0x16, 0xb3, 0x99, 0xc3, 0xf8, 0x2f, 0x01,
	0x7b, 0xd0, 0x4a, 0x42, 0x55, 0x9a, 0x35, 0x6d, 0xae, 0xeb, 0xde, 0x0f, 0x06, 0xe6, 0x9a, 0x8f,
	0x8f, 0xc0, 0x2c, 0x28, 0x2d, 0x64, 0x1e, 0x88, 0x58, 0x73, 0x8e, 0x78, 0xab, 0x14, 0x26, 0x31,
	0x5e, 0x42, 0x23, 0x8c, 0xf4, 0x74, 0x6b, 0x7a, 0x1e, 0x2f, 0xff, 0x39, 0xbf, 0x3b, 0xd4, 0xfd,
	0xbc, 0xe2, 0x20, 0xc2, 0xc1, 0x5c, 0xa4, 0x85, 0x65, 0xd8, 0x86, 0xc3, 0xb8, 0x7e, 0x5e, 0x86,
	0x15, 0xa9, 0xa2, 0xfc, 0x2e, 0x4c, 0xac, 0x83, 0x32, 0xc1, 0xaa, 0xee, 0xbf, 0x82, 0x46, 0x49,
	0xc0, 0x36, 0x34, 0x67, 0xde, 0xd4, 0xf3, 0x3f, 0x78, 0x9d, 0x07, 0x68, 0x42, 0xfd, 0xea, 0xcd,
	0x90, 0x8f, 0x3b, 0x0c, 0x9b, 0x60, 0x9c, 0x4d, 0xbc, 0x4e, 0x0d, 0x8f, 0x01, 0xde, 0xf9, 0x37,
	0x13, 0xef, 0x3c, 0x18, 0xde, 0x9c, 0x77, 0x8c, 0xfe, 0x29, 0xb4, 0x56, 0xd7, 0xb2, 0xdd, 0x0c,
	0xd0, 0x98, 0x78, 0x23, 0xdf, 0xe7, 0x1d, 0xb6, 0x34, 0xfc, 0xd9, 0xb5, 0x2e, 0x6a, 0xfd, 0x29,
	0x74, 0xb7, 0x0f, 0x53, 0x64, 0x32, 0x2d, 0x08, 0x9f, 0x00, 0x14, 0x5a, 0x09, 0x16, 0xd5, 0x90,
	0x4c, 0x6e, 0x96, 0xca, 0x4c, 0xc4, 0xd8, 0x85, 0xfa, 0xe6, 0x3b, 0x5d, 0x16, 0xfd, 0xb7, 0x70,
	0x32, 0xa2, 0x84, 0x7e, 0xff, 0x0e, 0xfe, 0x8b, 0xf5, 0x10, 0xba, 0xdb, 0xac, 0x32, 0xd8, 0xe0,
	0x3b, 0x83, 0xe6, 0xb8, 0xbc, 0x06, 0x0c, 0xe1, 0x70, 0x33, 0x3c, 0x3e, 0xdb, 0xf3, 0xae, 0x7a,
	0xce, 0xdf, 0x17, 0x56, 0x73, 0x08, 0xe1, 0x70, 0x33, 0xc6, 0xee, 0x2d, 0x76, 0x1c, 0x7a, 0xf7,
	0x16, 0xbb, 0x4e, 0x74, 0x66, 0x7e, 0x6c, 0x56, 0xfe, 0xbc, 0xa1, 0x7f, 0x32, 0x2f, 0x7e, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xf4, 0x25, 0x81, 0x9a, 0x75, 0x04, 0x00, 0x00,
}
