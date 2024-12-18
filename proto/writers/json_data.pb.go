// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: json_data.proto

package writers

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

type JsonData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Latitude  float64           `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64           `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Tags      map[string]string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *JsonData) Reset() {
	*x = JsonData{}
	mi := &file_json_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JsonData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonData) ProtoMessage() {}

func (x *JsonData) ProtoReflect() protoreflect.Message {
	mi := &file_json_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonData.ProtoReflect.Descriptor instead.
func (*JsonData) Descriptor() ([]byte, []int) {
	return file_json_data_proto_rawDescGZIP(), []int{0}
}

func (x *JsonData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JsonData) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *JsonData) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *JsonData) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_json_data_proto protoreflect.FileDescriptor

var file_json_data_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xbb, 0x01, 0x0a, 0x08, 0x4a, 0x73, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x2c,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x61, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09,
	0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x1f, 0x5a, 0x1d, 0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_json_data_proto_rawDescOnce sync.Once
	file_json_data_proto_rawDescData = file_json_data_proto_rawDesc
)

func file_json_data_proto_rawDescGZIP() []byte {
	file_json_data_proto_rawDescOnce.Do(func() {
		file_json_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_json_data_proto_rawDescData)
	})
	return file_json_data_proto_rawDescData
}

var file_json_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_json_data_proto_goTypes = []any{
	(*JsonData)(nil), // 0: main.JsonData
	nil,              // 1: main.JsonData.TagsEntry
}
var file_json_data_proto_depIdxs = []int32{
	1, // 0: main.JsonData.tags:type_name -> main.JsonData.TagsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_json_data_proto_init() }
func file_json_data_proto_init() {
	if File_json_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_json_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_json_data_proto_goTypes,
		DependencyIndexes: file_json_data_proto_depIdxs,
		MessageInfos:      file_json_data_proto_msgTypes,
	}.Build()
	File_json_data_proto = out.File
	file_json_data_proto_rawDesc = nil
	file_json_data_proto_goTypes = nil
	file_json_data_proto_depIdxs = nil
}
