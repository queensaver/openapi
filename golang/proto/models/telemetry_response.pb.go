//
//Queensaver API
//
//Queensaver API to send in sensor data and retrieve it. It's also used for user management.
//
//The version of the OpenAPI document: 0.0.1
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: models/telemetry_response.proto

package models

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

type TelemetryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HTTP response code. Used for internal purposes, will be let out at the API level.
	HttpResponseCode int32 `protobuf:"varint,1,opt,name=httpResponseCode,proto3" json:"httpResponseCode,omitempty"`
	// The measurement responses
	Values []*Telemetry `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *TelemetryResponse) Reset() {
	*x = TelemetryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_telemetry_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryResponse) ProtoMessage() {}

func (x *TelemetryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_telemetry_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryResponse.ProtoReflect.Descriptor instead.
func (*TelemetryResponse) Descriptor() ([]byte, []int) {
	return file_models_telemetry_response_proto_rawDescGZIP(), []int{0}
}

func (x *TelemetryResponse) GetHttpResponseCode() int32 {
	if x != nil {
		return x.HttpResponseCode
	}
	return 0
}

func (x *TelemetryResponse) GetValues() []*Telemetry {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_models_telemetry_response_proto protoreflect.FileDescriptor

var file_models_telemetry_response_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x1a, 0x16, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x11, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x68, 0x74, 0x74, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x10, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x50,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_telemetry_response_proto_rawDescOnce sync.Once
	file_models_telemetry_response_proto_rawDescData = file_models_telemetry_response_proto_rawDesc
)

func file_models_telemetry_response_proto_rawDescGZIP() []byte {
	file_models_telemetry_response_proto_rawDescOnce.Do(func() {
		file_models_telemetry_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_telemetry_response_proto_rawDescData)
	})
	return file_models_telemetry_response_proto_rawDescData
}

var file_models_telemetry_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_telemetry_response_proto_goTypes = []interface{}{
	(*TelemetryResponse)(nil), // 0: openapi.TelemetryResponse
	(*Telemetry)(nil),         // 1: openapi.Telemetry
}
var file_models_telemetry_response_proto_depIdxs = []int32{
	1, // 0: openapi.TelemetryResponse.values:type_name -> openapi.Telemetry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_telemetry_response_proto_init() }
func file_models_telemetry_response_proto_init() {
	if File_models_telemetry_response_proto != nil {
		return
	}
	file_models_telemetry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_telemetry_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryResponse); i {
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
			RawDescriptor: file_models_telemetry_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_telemetry_response_proto_goTypes,
		DependencyIndexes: file_models_telemetry_response_proto_depIdxs,
		MessageInfos:      file_models_telemetry_response_proto_msgTypes,
	}.Build()
	File_models_telemetry_response_proto = out.File
	file_models_telemetry_response_proto_rawDesc = nil
	file_models_telemetry_response_proto_goTypes = nil
	file_models_telemetry_response_proto_depIdxs = nil
}
