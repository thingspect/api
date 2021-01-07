// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: api/device.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	common "github.com/thingspect/api/go/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Device represents a device as stored in the database.
type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device ID (UUID).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Organization ID (UUID).
	OrgId string `protobuf:"bytes,2,opt,name=org_id,json=orgID,proto3" json:"org_id,omitempty"`
	// Device unique ID.
	UniqId string `protobuf:"bytes,3,opt,name=uniq_id,json=uniqID,proto3" json:"uniq_id,omitempty"`
	// Set to true to disable the device.
	IsDisabled bool `protobuf:"varint,4,opt,name=is_disabled,json=isDisabled,proto3" json:"is_disabled,omitempty"`
	// Authentication token (UUID).
	Token string `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	// Device creation timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Device modification timestamp.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Device status.
	Status common.Status `protobuf:"varint,8,opt,name=status,proto3,enum=common.Status" json:"status,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_api_device_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Device) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Device) GetUniqId() string {
	if x != nil {
		return x.UniqId
	}
	return ""
}

func (x *Device) GetIsDisabled() bool {
	if x != nil {
		return x.IsDisabled
	}
	return false
}

func (x *Device) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Device) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Device) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Device) GetStatus() common.Status {
	if x != nil {
		return x.Status
	}
	return common.Status_STATUS_UNSPECIFIED
}

// CreateDeviceRequest is sent to create a device.
type CreateDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device message to create.
	Device *Device `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *CreateDeviceRequest) Reset() {
	*x = CreateDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeviceRequest) ProtoMessage() {}

func (x *CreateDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeviceRequest.ProtoReflect.Descriptor instead.
func (*CreateDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_device_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDeviceRequest) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

// CreateDeviceResponse is sent in response to a device creation.
type CreateDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Created device message.
	Device *Device `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *CreateDeviceResponse) Reset() {
	*x = CreateDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeviceResponse) ProtoMessage() {}

func (x *CreateDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeviceResponse.ProtoReflect.Descriptor instead.
func (*CreateDeviceResponse) Descriptor() ([]byte, []int) {
	return file_api_device_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDeviceResponse) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

// ReadDeviceRequest is sent to read a device.
type ReadDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device ID (UUID) to read.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadDeviceRequest) Reset() {
	*x = ReadDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadDeviceRequest) ProtoMessage() {}

func (x *ReadDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadDeviceRequest.ProtoReflect.Descriptor instead.
func (*ReadDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_device_proto_rawDescGZIP(), []int{3}
}

func (x *ReadDeviceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ReadDeviceResponse is sent in response to a device read.
type ReadDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Read device message.
	Device *Device `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *ReadDeviceResponse) Reset() {
	*x = ReadDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadDeviceResponse) ProtoMessage() {}

func (x *ReadDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadDeviceResponse.ProtoReflect.Descriptor instead.
func (*ReadDeviceResponse) Descriptor() ([]byte, []int) {
	return file_api_device_proto_rawDescGZIP(), []int{4}
}

func (x *ReadDeviceResponse) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

// UpdateDeviceRequest is sent to update a device.
type UpdateDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device message to update.
	Device *Device `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	// Fields to update. Automatically populated by a PATCH request. If not present, a full resource update is performed.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateDeviceRequest) Reset() {
	*x = UpdateDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeviceRequest) ProtoMessage() {}

func (x *UpdateDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeviceRequest.ProtoReflect.Descriptor instead.
func (*UpdateDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_device_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDeviceRequest) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *UpdateDeviceRequest) GetUpdateMask() *field_mask.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// UpdateDeviceResponse is sent in response to a device update.
type UpdateDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Updated device message.
	Device *Device `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *UpdateDeviceResponse) Reset() {
	*x = UpdateDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeviceResponse) ProtoMessage() {}

func (x *UpdateDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeviceResponse.ProtoReflect.Descriptor instead.
func (*UpdateDeviceResponse) Descriptor() ([]byte, []int) {
	return file_api_device_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateDeviceResponse) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

// DeleteDeviceRequest is sent to delete a device.
type DeleteDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device ID (UUID) to delete.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDeviceRequest) Reset() {
	*x = DeleteDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeviceRequest) ProtoMessage() {}

func (x *DeleteDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeviceRequest.ProtoReflect.Descriptor instead.
func (*DeleteDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_device_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteDeviceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of devices to return.
	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset to start the results.
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListDeviceRequest) Reset() {
	*x = ListDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeviceRequest) ProtoMessage() {}

func (x *ListDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeviceRequest.ProtoReflect.Descriptor instead.
func (*ListDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_device_proto_rawDescGZIP(), []int{8}
}

func (x *ListDeviceRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListDeviceRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total number of devices.
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// Device array.
	Devices []*Device `protobuf:"bytes,2,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *ListDeviceResponse) Reset() {
	*x = ListDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_device_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeviceResponse) ProtoMessage() {}

func (x *ListDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_device_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeviceResponse.ProtoReflect.Descriptor instead.
func (*ListDeviceResponse) Descriptor() ([]byte, []int) {
	return file_api_device_proto_rawDescGZIP(), []int{9}
}

func (x *ListDeviceResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListDeviceResponse) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

var File_api_device_proto protoreflect.FileDescriptor

var file_api_device_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x02, 0x0a, 0x06,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x6f, 0x72, 0x67,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x49, 0x44,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x3f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x3b, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x28, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x12,
	0x52, 0x65, 0x61, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x7c, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x3b, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x2a, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x22, 0x51, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x25, 0x0a,
	0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x32, 0xe9, 0x03, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x51, 0x0a,
	0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x89, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x4a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x44, 0x1a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x64,
	0x7d, 0x3a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5a, 0x21, 0x32, 0x17, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x69, 0x64, 0x7d, 0x3a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x2a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x4c, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_device_proto_rawDescOnce sync.Once
	file_api_device_proto_rawDescData = file_api_device_proto_rawDesc
)

func file_api_device_proto_rawDescGZIP() []byte {
	file_api_device_proto_rawDescOnce.Do(func() {
		file_api_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_device_proto_rawDescData)
	})
	return file_api_device_proto_rawDescData
}

var file_api_device_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_device_proto_goTypes = []interface{}{
	(*Device)(nil),               // 0: api.Device
	(*CreateDeviceRequest)(nil),  // 1: api.CreateDeviceRequest
	(*CreateDeviceResponse)(nil), // 2: api.CreateDeviceResponse
	(*ReadDeviceRequest)(nil),    // 3: api.ReadDeviceRequest
	(*ReadDeviceResponse)(nil),   // 4: api.ReadDeviceResponse
	(*UpdateDeviceRequest)(nil),  // 5: api.UpdateDeviceRequest
	(*UpdateDeviceResponse)(nil), // 6: api.UpdateDeviceResponse
	(*DeleteDeviceRequest)(nil),  // 7: api.DeleteDeviceRequest
	(*ListDeviceRequest)(nil),    // 8: api.ListDeviceRequest
	(*ListDeviceResponse)(nil),   // 9: api.ListDeviceResponse
	(*timestamp.Timestamp)(nil),  // 10: google.protobuf.Timestamp
	(common.Status)(0),           // 11: common.Status
	(*field_mask.FieldMask)(nil), // 12: google.protobuf.FieldMask
	(*empty.Empty)(nil),          // 13: google.protobuf.Empty
}
var file_api_device_proto_depIdxs = []int32{
	10, // 0: api.Device.created_at:type_name -> google.protobuf.Timestamp
	10, // 1: api.Device.updated_at:type_name -> google.protobuf.Timestamp
	11, // 2: api.Device.status:type_name -> common.Status
	0,  // 3: api.CreateDeviceRequest.device:type_name -> api.Device
	0,  // 4: api.CreateDeviceResponse.device:type_name -> api.Device
	0,  // 5: api.ReadDeviceResponse.device:type_name -> api.Device
	0,  // 6: api.UpdateDeviceRequest.device:type_name -> api.Device
	12, // 7: api.UpdateDeviceRequest.update_mask:type_name -> google.protobuf.FieldMask
	0,  // 8: api.UpdateDeviceResponse.device:type_name -> api.Device
	0,  // 9: api.ListDeviceResponse.devices:type_name -> api.Device
	1,  // 10: api.DeviceService.Create:input_type -> api.CreateDeviceRequest
	3,  // 11: api.DeviceService.Read:input_type -> api.ReadDeviceRequest
	5,  // 12: api.DeviceService.Update:input_type -> api.UpdateDeviceRequest
	7,  // 13: api.DeviceService.Delete:input_type -> api.DeleteDeviceRequest
	8,  // 14: api.DeviceService.List:input_type -> api.ListDeviceRequest
	2,  // 15: api.DeviceService.Create:output_type -> api.CreateDeviceResponse
	4,  // 16: api.DeviceService.Read:output_type -> api.ReadDeviceResponse
	6,  // 17: api.DeviceService.Update:output_type -> api.UpdateDeviceResponse
	13, // 18: api.DeviceService.Delete:output_type -> google.protobuf.Empty
	9,  // 19: api.DeviceService.List:output_type -> api.ListDeviceResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_api_device_proto_init() }
func file_api_device_proto_init() {
	if File_api_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_api_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeviceRequest); i {
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
		file_api_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeviceResponse); i {
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
		file_api_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadDeviceRequest); i {
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
		file_api_device_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadDeviceResponse); i {
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
		file_api_device_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeviceRequest); i {
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
		file_api_device_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeviceResponse); i {
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
		file_api_device_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeviceRequest); i {
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
		file_api_device_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeviceRequest); i {
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
		file_api_device_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeviceResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_device_proto_goTypes,
		DependencyIndexes: file_api_device_proto_depIdxs,
		MessageInfos:      file_api_device_proto_msgTypes,
	}.Build()
	File_api_device_proto = out.File
	file_api_device_proto_rawDesc = nil
	file_api_device_proto_goTypes = nil
	file_api_device_proto_depIdxs = nil
}
