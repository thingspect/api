// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: common/thingspect_datapoint.proto

package common

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DataPoint represents a data point as stored in the database.
type DataPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device unique ID. Ignored during MQTT ingest if provided as part of the topic. Required for API publish.
	UniqId string `protobuf:"bytes,1,opt,name=uniq_id,json=uniqID,proto3" json:"uniq_id,omitempty"`
	// Device attribute.
	Attr string `protobuf:"bytes,2,opt,name=attr,proto3" json:"attr,omitempty"`
	// Attribute value.
	//
	// Types that are assignable to ValOneof:
	//
	//	*DataPoint_IntVal
	//	*DataPoint_Fl64Val
	//	*DataPoint_StrVal
	//	*DataPoint_BoolVal
	//	*DataPoint_BytesVal
	ValOneof isDataPoint_ValOneof `protobuf_oneof:"val_oneof"`
	// Timestamp. If not present during MQTT ingest or API publish, the current time will be used.
	Ts *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=ts,proto3" json:"ts,omitempty"`
	// Authentication token (UUID). Only used during MQTT ingest. Will be ignored if provided as part of the Payload message.
	Token string `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	// Trace ID (UUID). Assigned by the platform, will be ignored if provided during MQTT ingest or API publish.
	TraceId string `protobuf:"bytes,9,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
}

func (x *DataPoint) Reset() {
	*x = DataPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_thingspect_datapoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPoint) ProtoMessage() {}

func (x *DataPoint) ProtoReflect() protoreflect.Message {
	mi := &file_common_thingspect_datapoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPoint.ProtoReflect.Descriptor instead.
func (*DataPoint) Descriptor() ([]byte, []int) {
	return file_common_thingspect_datapoint_proto_rawDescGZIP(), []int{0}
}

func (x *DataPoint) GetUniqId() string {
	if x != nil {
		return x.UniqId
	}
	return ""
}

func (x *DataPoint) GetAttr() string {
	if x != nil {
		return x.Attr
	}
	return ""
}

func (m *DataPoint) GetValOneof() isDataPoint_ValOneof {
	if m != nil {
		return m.ValOneof
	}
	return nil
}

func (x *DataPoint) GetIntVal() int32 {
	if x, ok := x.GetValOneof().(*DataPoint_IntVal); ok {
		return x.IntVal
	}
	return 0
}

func (x *DataPoint) GetFl64Val() float64 {
	if x, ok := x.GetValOneof().(*DataPoint_Fl64Val); ok {
		return x.Fl64Val
	}
	return 0
}

func (x *DataPoint) GetStrVal() string {
	if x, ok := x.GetValOneof().(*DataPoint_StrVal); ok {
		return x.StrVal
	}
	return ""
}

func (x *DataPoint) GetBoolVal() bool {
	if x, ok := x.GetValOneof().(*DataPoint_BoolVal); ok {
		return x.BoolVal
	}
	return false
}

func (x *DataPoint) GetBytesVal() []byte {
	if x, ok := x.GetValOneof().(*DataPoint_BytesVal); ok {
		return x.BytesVal
	}
	return nil
}

func (x *DataPoint) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *DataPoint) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DataPoint) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

type isDataPoint_ValOneof interface {
	isDataPoint_ValOneof()
}

type DataPoint_IntVal struct {
	// Integer value, 32-bit.
	IntVal int32 `protobuf:"zigzag32,3,opt,name=int_val,json=intVal,proto3,oneof"`
}

type DataPoint_Fl64Val struct {
	// Float value, 64-bit.
	Fl64Val float64 `protobuf:"fixed64,4,opt,name=fl64_val,json=fl64Val,proto3,oneof"`
}

type DataPoint_StrVal struct {
	// String value.
	StrVal string `protobuf:"bytes,5,opt,name=str_val,json=strVal,proto3,oneof"`
}

type DataPoint_BoolVal struct {
	// Boolean value.
	BoolVal bool `protobuf:"varint,6,opt,name=bool_val,json=boolVal,proto3,oneof"`
}

type DataPoint_BytesVal struct {
	// Bytes value. When used in JSON, the value will be represented as a base64 string.
	BytesVal []byte `protobuf:"bytes,16,opt,name=bytes_val,json=bytesVal,proto3,oneof"`
}

func (*DataPoint_IntVal) isDataPoint_ValOneof() {}

func (*DataPoint_Fl64Val) isDataPoint_ValOneof() {}

func (*DataPoint_StrVal) isDataPoint_ValOneof() {}

func (*DataPoint_BoolVal) isDataPoint_ValOneof() {}

func (*DataPoint_BytesVal) isDataPoint_ValOneof() {}

var File_common_thingspect_datapoint_proto protoreflect.FileDescriptor

var file_common_thingspect_datapoint_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xda, 0x02, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x25, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0c, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x05, 0x18, 0x28, 0x52, 0x06,
	0x75, 0x6e, 0x69, 0x71, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x04, 0x61, 0x74, 0x74, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x28,
	0x52, 0x04, 0x61, 0x74, 0x74, 0x72, 0x12, 0x19, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x48, 0x00, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x56, 0x61,
	0x6c, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x6c, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x07, 0x66, 0x6c, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x12, 0x19,
	0x0a, 0x07, 0x73, 0x74, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x08, 0x62, 0x6f, 0x6f,
	0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x62,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f,
	0x76, 0x61, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x08, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x56, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74,
	0x73, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x42, 0x10, 0x0a, 0x09,
	0x76, 0x61, 0x6c, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x25,
	0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_thingspect_datapoint_proto_rawDescOnce sync.Once
	file_common_thingspect_datapoint_proto_rawDescData = file_common_thingspect_datapoint_proto_rawDesc
)

func file_common_thingspect_datapoint_proto_rawDescGZIP() []byte {
	file_common_thingspect_datapoint_proto_rawDescOnce.Do(func() {
		file_common_thingspect_datapoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_thingspect_datapoint_proto_rawDescData)
	})
	return file_common_thingspect_datapoint_proto_rawDescData
}

var file_common_thingspect_datapoint_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_thingspect_datapoint_proto_goTypes = []interface{}{
	(*DataPoint)(nil),             // 0: thingspect.common.DataPoint
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_common_thingspect_datapoint_proto_depIdxs = []int32{
	1, // 0: thingspect.common.DataPoint.ts:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_thingspect_datapoint_proto_init() }
func file_common_thingspect_datapoint_proto_init() {
	if File_common_thingspect_datapoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_thingspect_datapoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPoint); i {
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
	file_common_thingspect_datapoint_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DataPoint_IntVal)(nil),
		(*DataPoint_Fl64Val)(nil),
		(*DataPoint_StrVal)(nil),
		(*DataPoint_BoolVal)(nil),
		(*DataPoint_BytesVal)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_thingspect_datapoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_thingspect_datapoint_proto_goTypes,
		DependencyIndexes: file_common_thingspect_datapoint_proto_depIdxs,
		MessageInfos:      file_common_thingspect_datapoint_proto_msgTypes,
	}.Build()
	File_common_thingspect_datapoint_proto = out.File
	file_common_thingspect_datapoint_proto_rawDesc = nil
	file_common_thingspect_datapoint_proto_goTypes = nil
	file_common_thingspect_datapoint_proto_depIdxs = nil
}
