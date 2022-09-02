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
// source: models/alerts.proto

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

// Various alerts a bee hive can have.
type Alerts_AlertTypeEnum int32

const (
	Alerts_AlertTypeEnum_VARROAHIGH Alerts_AlertTypeEnum = 0
	Alerts_AlertTypeEnum_WEIGHTLOW  Alerts_AlertTypeEnum = 1
	Alerts_AlertTypeEnum_QUEENLESS  Alerts_AlertTypeEnum = 2
)

// Enum value maps for Alerts_AlertTypeEnum.
var (
	Alerts_AlertTypeEnum_name = map[int32]string{
		0: "AlertTypeEnum_VARROAHIGH",
		1: "AlertTypeEnum_WEIGHTLOW",
		2: "AlertTypeEnum_QUEENLESS",
	}
	Alerts_AlertTypeEnum_value = map[string]int32{
		"AlertTypeEnum_VARROAHIGH": 0,
		"AlertTypeEnum_WEIGHTLOW":  1,
		"AlertTypeEnum_QUEENLESS":  2,
	}
)

func (x Alerts_AlertTypeEnum) Enum() *Alerts_AlertTypeEnum {
	p := new(Alerts_AlertTypeEnum)
	*p = x
	return p
}

func (x Alerts_AlertTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Alerts_AlertTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_models_alerts_proto_enumTypes[0].Descriptor()
}

func (Alerts_AlertTypeEnum) Type() protoreflect.EnumType {
	return &file_models_alerts_proto_enumTypes[0]
}

func (x Alerts_AlertTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Alerts_AlertTypeEnum.Descriptor instead.
func (Alerts_AlertTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_models_alerts_proto_rawDescGZIP(), []int{0, 0}
}

type Alerts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlertType Alerts_AlertTypeEnum `protobuf:"varint,1,opt,name=alertType,proto3,enum=openapi.Alerts_AlertTypeEnum" json:"alertType,omitempty"`
}

func (x *Alerts) Reset() {
	*x = Alerts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_alerts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alerts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alerts) ProtoMessage() {}

func (x *Alerts) ProtoReflect() protoreflect.Message {
	mi := &file_models_alerts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alerts.ProtoReflect.Descriptor instead.
func (*Alerts) Descriptor() ([]byte, []int) {
	return file_models_alerts_proto_rawDescGZIP(), []int{0}
}

func (x *Alerts) GetAlertType() Alerts_AlertTypeEnum {
	if x != nil {
		return x.AlertType
	}
	return Alerts_AlertTypeEnum_VARROAHIGH
}

var File_models_alerts_proto protoreflect.FileDescriptor

var file_models_alerts_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x22, 0xae,
	0x01, 0x0a, 0x06, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x09, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x67, 0x0a, 0x0d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x56, 0x41, 0x52, 0x52, 0x4f, 0x41, 0x48,
	0x49, 0x47, 0x48, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x57, 0x45, 0x49, 0x47, 0x48, 0x54, 0x4c, 0x4f, 0x57,
	0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x5f, 0x51, 0x55, 0x45, 0x45, 0x4e, 0x4c, 0x45, 0x53, 0x53, 0x10, 0x02, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_alerts_proto_rawDescOnce sync.Once
	file_models_alerts_proto_rawDescData = file_models_alerts_proto_rawDesc
)

func file_models_alerts_proto_rawDescGZIP() []byte {
	file_models_alerts_proto_rawDescOnce.Do(func() {
		file_models_alerts_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_alerts_proto_rawDescData)
	})
	return file_models_alerts_proto_rawDescData
}

var file_models_alerts_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_models_alerts_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_alerts_proto_goTypes = []interface{}{
	(Alerts_AlertTypeEnum)(0), // 0: openapi.Alerts.AlertTypeEnum
	(*Alerts)(nil),            // 1: openapi.Alerts
}
var file_models_alerts_proto_depIdxs = []int32{
	0, // 0: openapi.Alerts.alertType:type_name -> openapi.Alerts.AlertTypeEnum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_alerts_proto_init() }
func file_models_alerts_proto_init() {
	if File_models_alerts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_alerts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alerts); i {
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
			RawDescriptor: file_models_alerts_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_alerts_proto_goTypes,
		DependencyIndexes: file_models_alerts_proto_depIdxs,
		EnumInfos:         file_models_alerts_proto_enumTypes,
		MessageInfos:      file_models_alerts_proto_msgTypes,
	}.Build()
	File_models_alerts_proto = out.File
	file_models_alerts_proto_rawDesc = nil
	file_models_alerts_proto_goTypes = nil
	file_models_alerts_proto_depIdxs = nil
}