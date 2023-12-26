// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: v1/blueprint.proto

package power

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

type Blueprint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec *Spec  `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Blueprint) Reset() {
	*x = Blueprint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_blueprint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blueprint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blueprint) ProtoMessage() {}

func (x *Blueprint) ProtoReflect() protoreflect.Message {
	mi := &file_v1_blueprint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blueprint.ProtoReflect.Descriptor instead.
func (*Blueprint) Descriptor() ([]byte, []int) {
	return file_v1_blueprint_proto_rawDescGZIP(), []int{0}
}

func (x *Blueprint) GetSpec() *Spec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Blueprint) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Options     []*Option `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_blueprint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_v1_blueprint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec.ProtoReflect.Descriptor instead.
func (*Spec) Descriptor() ([]byte, []int) {
	return file_v1_blueprint_proto_rawDescGZIP(), []int{1}
}

func (x *Spec) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Spec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Spec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Spec) GetOptions() []*Option {
	if x != nil {
		return x.Options
	}
	return nil
}

type Option struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Default     string   `protobuf:"bytes,5,opt,name=default,proto3" json:"default,omitempty"`
	Choices     []string `protobuf:"bytes,6,rep,name=choices,proto3" json:"choices,omitempty"`
}

func (x *Option) Reset() {
	*x = Option{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_blueprint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Option) ProtoMessage() {}

func (x *Option) ProtoReflect() protoreflect.Message {
	mi := &file_v1_blueprint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Option.ProtoReflect.Descriptor instead.
func (*Option) Descriptor() ([]byte, []int) {
	return file_v1_blueprint_proto_rawDescGZIP(), []int{2}
}

func (x *Option) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Option) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Option) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Option) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Option) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *Option) GetChoices() []string {
	if x != nil {
		return x.Choices
	}
	return nil
}

var File_v1_blueprint_proto protoreflect.FileDescriptor

var file_v1_blueprint_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x22, 0x40, 0x0a, 0x09, 0x42,
	0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x75, 0x0a,
	0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x42, 0x19, 0x5a,
	0x17, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x63, 0x75, 0x62,
	0x62, 0x73, 0x2f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_blueprint_proto_rawDescOnce sync.Once
	file_v1_blueprint_proto_rawDescData = file_v1_blueprint_proto_rawDesc
)

func file_v1_blueprint_proto_rawDescGZIP() []byte {
	file_v1_blueprint_proto_rawDescOnce.Do(func() {
		file_v1_blueprint_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_blueprint_proto_rawDescData)
	})
	return file_v1_blueprint_proto_rawDescData
}

var file_v1_blueprint_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_blueprint_proto_goTypes = []interface{}{
	(*Blueprint)(nil), // 0: pb.v1.Blueprint
	(*Spec)(nil),      // 1: pb.v1.Spec
	(*Option)(nil),    // 2: pb.v1.Option
}
var file_v1_blueprint_proto_depIdxs = []int32{
	1, // 0: pb.v1.Blueprint.spec:type_name -> pb.v1.Spec
	2, // 1: pb.v1.Spec.options:type_name -> pb.v1.Option
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_blueprint_proto_init() }
func file_v1_blueprint_proto_init() {
	if File_v1_blueprint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_blueprint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blueprint); i {
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
		file_v1_blueprint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spec); i {
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
		file_v1_blueprint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Option); i {
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
			RawDescriptor: file_v1_blueprint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_blueprint_proto_goTypes,
		DependencyIndexes: file_v1_blueprint_proto_depIdxs,
		MessageInfos:      file_v1_blueprint_proto_msgTypes,
	}.Build()
	File_v1_blueprint_proto = out.File
	file_v1_blueprint_proto_rawDesc = nil
	file_v1_blueprint_proto_goTypes = nil
	file_v1_blueprint_proto_depIdxs = nil
}
