// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: transport.proto

package main

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Direction is from Client to Server
type DeviceConnected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialNumber string `protobuf:"bytes,1,opt,name=serialNumber,proto3" json:"serialNumber,omitempty"`
	VendorId     int32  `protobuf:"varint,2,opt,name=vendorId,proto3" json:"vendorId,omitempty"`
	ProductId    int32  `protobuf:"varint,3,opt,name=productId,proto3" json:"productId,omitempty"`
}

func (x *DeviceConnected) Reset() {
	*x = DeviceConnected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceConnected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceConnected) ProtoMessage() {}

func (x *DeviceConnected) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceConnected.ProtoReflect.Descriptor instead.
func (*DeviceConnected) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceConnected) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *DeviceConnected) GetVendorId() int32 {
	if x != nil {
		return x.VendorId
	}
	return 0
}

func (x *DeviceConnected) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type DataFromDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialNumber string `protobuf:"bytes,1,opt,name=serialNumber,proto3" json:"serialNumber,omitempty"`
	Data         []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DataFromDevice) Reset() {
	*x = DataFromDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataFromDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataFromDevice) ProtoMessage() {}

func (x *DataFromDevice) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataFromDevice.ProtoReflect.Descriptor instead.
func (*DataFromDevice) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{1}
}

func (x *DataFromDevice) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *DataFromDevice) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type DataToDeviceResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlationId,proto3" json:"correlationId,omitempty"`
	Success       bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DataToDeviceResult) Reset() {
	*x = DataToDeviceResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataToDeviceResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataToDeviceResult) ProtoMessage() {}

func (x *DataToDeviceResult) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataToDeviceResult.ProtoReflect.Descriptor instead.
func (*DataToDeviceResult) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{2}
}

func (x *DataToDeviceResult) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *DataToDeviceResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DataToDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialNumber  string `protobuf:"bytes,1,opt,name=serialNumber,proto3" json:"serialNumber,omitempty"`
	CorrelationId string `protobuf:"bytes,2,opt,name=correlationId,proto3" json:"correlationId,omitempty"`
	Data          []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DataToDevice) Reset() {
	*x = DataToDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataToDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataToDevice) ProtoMessage() {}

func (x *DataToDevice) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataToDevice.ProtoReflect.Descriptor instead.
func (*DataToDevice) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{3}
}

func (x *DataToDevice) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *DataToDevice) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *DataToDevice) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ServerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*ServerMessage_DeviceConnected
	//	*ServerMessage_FromDevice
	//	*ServerMessage_ToDeviceResult
	Message isServerMessage_Message `protobuf_oneof:"message"`
}

func (x *ServerMessage) Reset() {
	*x = ServerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerMessage) ProtoMessage() {}

func (x *ServerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerMessage.ProtoReflect.Descriptor instead.
func (*ServerMessage) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{4}
}

func (m *ServerMessage) GetMessage() isServerMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ServerMessage) GetDeviceConnected() *DeviceConnected {
	if x, ok := x.GetMessage().(*ServerMessage_DeviceConnected); ok {
		return x.DeviceConnected
	}
	return nil
}

func (x *ServerMessage) GetFromDevice() *DataFromDevice {
	if x, ok := x.GetMessage().(*ServerMessage_FromDevice); ok {
		return x.FromDevice
	}
	return nil
}

func (x *ServerMessage) GetToDeviceResult() *DataToDeviceResult {
	if x, ok := x.GetMessage().(*ServerMessage_ToDeviceResult); ok {
		return x.ToDeviceResult
	}
	return nil
}

type isServerMessage_Message interface {
	isServerMessage_Message()
}

type ServerMessage_DeviceConnected struct {
	DeviceConnected *DeviceConnected `protobuf:"bytes,1,opt,name=deviceConnected,proto3,oneof"`
}

type ServerMessage_FromDevice struct {
	FromDevice *DataFromDevice `protobuf:"bytes,2,opt,name=fromDevice,proto3,oneof"`
}

type ServerMessage_ToDeviceResult struct {
	ToDeviceResult *DataToDeviceResult `protobuf:"bytes,3,opt,name=toDeviceResult,proto3,oneof"`
}

func (*ServerMessage_DeviceConnected) isServerMessage_Message() {}

func (*ServerMessage_FromDevice) isServerMessage_Message() {}

func (*ServerMessage_ToDeviceResult) isServerMessage_Message() {}

type ClientMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*ClientMessage_ToDevice
	Message isClientMessage_Message `protobuf_oneof:"message"`
}

func (x *ClientMessage) Reset() {
	*x = ClientMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMessage) ProtoMessage() {}

func (x *ClientMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMessage.ProtoReflect.Descriptor instead.
func (*ClientMessage) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{5}
}

func (m *ClientMessage) GetMessage() isClientMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ClientMessage) GetToDevice() *DataToDevice {
	if x, ok := x.GetMessage().(*ClientMessage_ToDevice); ok {
		return x.ToDevice
	}
	return nil
}

type isClientMessage_Message interface {
	isClientMessage_Message()
}

type ClientMessage_ToDevice struct {
	ToDevice *DataToDevice `protobuf:"bytes,1,opt,name=toDevice,proto3,oneof"`
}

func (*ClientMessage_ToDevice) isClientMessage_Message() {}

var File_transport_proto protoreflect.FileDescriptor

var file_transport_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6f, 0x0a, 0x0f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x22, 0x48, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x54, 0x0a, 0x12,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x6c, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xca, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52,
	0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x12, 0x31, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x46, 0x72, 0x6f, 0x6d, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x74, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x54, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x48, 0x00, 0x52, 0x0e, 0x74, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x47, 0x0a,
	0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b,
	0x0a, 0x08, 0x74, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x48,
	0x00, 0x52, 0x08, 0x74, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6d, 0x61, 0x69, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_proto_rawDescOnce sync.Once
	file_transport_proto_rawDescData = file_transport_proto_rawDesc
)

func file_transport_proto_rawDescGZIP() []byte {
	file_transport_proto_rawDescOnce.Do(func() {
		file_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_proto_rawDescData)
	})
	return file_transport_proto_rawDescData
}

var file_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_transport_proto_goTypes = []interface{}{
	(*DeviceConnected)(nil),    // 0: DeviceConnected
	(*DataFromDevice)(nil),     // 1: DataFromDevice
	(*DataToDeviceResult)(nil), // 2: DataToDeviceResult
	(*DataToDevice)(nil),       // 3: DataToDevice
	(*ServerMessage)(nil),      // 4: ServerMessage
	(*ClientMessage)(nil),      // 5: ClientMessage
}
var file_transport_proto_depIdxs = []int32{
	0, // 0: ServerMessage.deviceConnected:type_name -> DeviceConnected
	1, // 1: ServerMessage.fromDevice:type_name -> DataFromDevice
	2, // 2: ServerMessage.toDeviceResult:type_name -> DataToDeviceResult
	3, // 3: ClientMessage.toDevice:type_name -> DataToDevice
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_transport_proto_init() }
func file_transport_proto_init() {
	if File_transport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceConnected); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataFromDevice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataToDeviceResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transport_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataToDevice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transport_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transport_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_transport_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ServerMessage_DeviceConnected)(nil),
		(*ServerMessage_FromDevice)(nil),
		(*ServerMessage_ToDeviceResult)(nil),
	}
	file_transport_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ClientMessage_ToDevice)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_proto_goTypes,
		DependencyIndexes: file_transport_proto_depIdxs,
		MessageInfos:      file_transport_proto_msgTypes,
	}.Build()
	File_transport_proto = out.File
	file_transport_proto_rawDesc = nil
	file_transport_proto_goTypes = nil
	file_transport_proto_depIdxs = nil
}
