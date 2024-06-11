// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: rabbit/push/config.proto

package push

import (
	api "github.com/aide-family/moon/api"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type NotifyObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 根据路由匹配具体的发送对象
	Receivers map[string]*api.Receiver `protobuf:"bytes,4,rep,name=receivers,proto3" json:"receivers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Templates map[string]string        `protobuf:"bytes,5,rep,name=templates,proto3" json:"templates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NotifyObjectRequest) Reset() {
	*x = NotifyObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rabbit_push_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyObjectRequest) ProtoMessage() {}

func (x *NotifyObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rabbit_push_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyObjectRequest.ProtoReflect.Descriptor instead.
func (*NotifyObjectRequest) Descriptor() ([]byte, []int) {
	return file_rabbit_push_config_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyObjectRequest) GetReceivers() map[string]*api.Receiver {
	if x != nil {
		return x.Receivers
	}
	return nil
}

func (x *NotifyObjectRequest) GetTemplates() map[string]string {
	if x != nil {
		return x.Templates
	}
	return nil
}

type NotifyObjectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 发送的结果
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	// 状态码
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	// 发送时间
	Time string `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *NotifyObjectReply) Reset() {
	*x = NotifyObjectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rabbit_push_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyObjectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyObjectReply) ProtoMessage() {}

func (x *NotifyObjectReply) ProtoReflect() protoreflect.Message {
	mi := &file_rabbit_push_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyObjectReply.ProtoReflect.Descriptor instead.
func (*NotifyObjectReply) Descriptor() ([]byte, []int) {
	return file_rabbit_push_config_proto_rawDescGZIP(), []int{1}
}

func (x *NotifyObjectReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *NotifyObjectReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *NotifyObjectReply) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

var File_rabbit_push_config_proto protoreflect.FileDescriptor

var file_rabbit_push_config_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x02, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x51, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x2e,
	0x70, 0x75, 0x73, 0x68, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x12, 0x51, 0x0a, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x62, 0x62,
	0x69, 0x74, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x4b, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x4d, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x32,
	0x85, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x7b, 0x0a, 0x0c, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x2e, 0x70, 0x75,
	0x73, 0x68, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22,
	0x16, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x2f, 0x70, 0x75, 0x73, 0x68,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x45, 0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x61, 0x62, 0x62, 0x69, 0x74, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x64, 0x65, 0x2d, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x2f, 0x6d, 0x6f, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x61,
	0x62, 0x62, 0x69, 0x74, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x3b, 0x70, 0x75, 0x73, 0x68, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rabbit_push_config_proto_rawDescOnce sync.Once
	file_rabbit_push_config_proto_rawDescData = file_rabbit_push_config_proto_rawDesc
)

func file_rabbit_push_config_proto_rawDescGZIP() []byte {
	file_rabbit_push_config_proto_rawDescOnce.Do(func() {
		file_rabbit_push_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_rabbit_push_config_proto_rawDescData)
	})
	return file_rabbit_push_config_proto_rawDescData
}

var file_rabbit_push_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rabbit_push_config_proto_goTypes = []interface{}{
	(*NotifyObjectRequest)(nil), // 0: api.rabbit.push.NotifyObjectRequest
	(*NotifyObjectReply)(nil),   // 1: api.rabbit.push.NotifyObjectReply
	nil,                         // 2: api.rabbit.push.NotifyObjectRequest.ReceiversEntry
	nil,                         // 3: api.rabbit.push.NotifyObjectRequest.TemplatesEntry
	(*api.Receiver)(nil),        // 4: api.Receiver
}
var file_rabbit_push_config_proto_depIdxs = []int32{
	2, // 0: api.rabbit.push.NotifyObjectRequest.receivers:type_name -> api.rabbit.push.NotifyObjectRequest.ReceiversEntry
	3, // 1: api.rabbit.push.NotifyObjectRequest.templates:type_name -> api.rabbit.push.NotifyObjectRequest.TemplatesEntry
	4, // 2: api.rabbit.push.NotifyObjectRequest.ReceiversEntry.value:type_name -> api.Receiver
	0, // 3: api.rabbit.push.Config.NotifyObject:input_type -> api.rabbit.push.NotifyObjectRequest
	1, // 4: api.rabbit.push.Config.NotifyObject:output_type -> api.rabbit.push.NotifyObjectReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rabbit_push_config_proto_init() }
func file_rabbit_push_config_proto_init() {
	if File_rabbit_push_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rabbit_push_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyObjectRequest); i {
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
		file_rabbit_push_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyObjectReply); i {
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
			RawDescriptor: file_rabbit_push_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rabbit_push_config_proto_goTypes,
		DependencyIndexes: file_rabbit_push_config_proto_depIdxs,
		MessageInfos:      file_rabbit_push_config_proto_msgTypes,
	}.Build()
	File_rabbit_push_config_proto = out.File
	file_rabbit_push_config_proto_rawDesc = nil
	file_rabbit_push_config_proto_goTypes = nil
	file_rabbit_push_config_proto_depIdxs = nil
}
