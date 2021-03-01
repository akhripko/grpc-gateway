// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: echo/echo.proto

package echo

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data1 []string       `protobuf:"bytes,2,rep,name=data1,proto3" json:"data1,omitempty"`
	Data2 []int32        `protobuf:"varint,3,rep,packed,name=data2,proto3" json:"data2,omitempty"`
	EmId  *EchoMessageId `protobuf:"bytes,4,opt,name=em_id,json=emId,proto3" json:"em_id,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_echo_echo_proto_rawDescGZIP(), []int{0}
}

func (x *EchoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EchoRequest) GetData1() []string {
	if x != nil {
		return x.Data1
	}
	return nil
}

func (x *EchoRequest) GetData2() []int32 {
	if x != nil {
		return x.Data2
	}
	return nil
}

func (x *EchoRequest) GetEmId() *EchoMessageId {
	if x != nil {
		return x.EmId
	}
	return nil
}

type EchoMessageId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EchoMessageId) Reset() {
	*x = EchoMessageId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoMessageId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoMessageId) ProtoMessage() {}

func (x *EchoMessageId) ProtoReflect() protoreflect.Message {
	mi := &file_echo_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoMessageId.ProtoReflect.Descriptor instead.
func (*EchoMessageId) Descriptor() ([]byte, []int) {
	return file_echo_echo_proto_rawDescGZIP(), []int{1}
}

func (x *EchoMessageId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data1 []string       `protobuf:"bytes,2,rep,name=data1,proto3" json:"data1,omitempty"`
	Data2 []int32        `protobuf:"varint,3,rep,packed,name=data2,proto3" json:"data2,omitempty"`
	EmId  *EchoMessageId `protobuf:"bytes,4,opt,name=em_id,json=emId,proto3" json:"em_id,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_echo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_echo_echo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_echo_echo_proto_rawDescGZIP(), []int{2}
}

func (x *EchoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EchoResponse) GetData1() []string {
	if x != nil {
		return x.Data1
	}
	return nil
}

func (x *EchoResponse) GetData2() []int32 {
	if x != nil {
		return x.Data2
	}
	return nil
}

func (x *EchoResponse) GetEmId() *EchoMessageId {
	if x != nil {
		return x.EmId
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_echo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_echo_echo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_echo_echo_proto_rawDescGZIP(), []int{3}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_echo_echo_proto protoreflect.FileDescriptor

var file_echo_echo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x61,
	0x31, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x74, 0x61, 0x31, 0x12, 0x14,
	0x0a, 0x05, 0x64, 0x61, 0x74, 0x61, 0x32, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x64,
	0x61, 0x74, 0x61, 0x32, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x52, 0x04, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x1f,
	0x0a, 0x0d, 0x45, 0x63, 0x68, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x78, 0x0a, 0x0c, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x61, 0x31, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x74, 0x61, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x61, 0x74,
	0x61, 0x32, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x64, 0x61, 0x74, 0x61, 0x32, 0x12,
	0x28, 0x0a, 0x05, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x52, 0x04, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0xa7, 0x01, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4d, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x11, 0x2e, 0x65,
	0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x01, 0x2a, 0x12,
	0x49, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x11, 0x2e, 0x65, 0x63, 0x68,
	0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x63, 0x68, 0x6f, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x0a, 0x5a, 0x08, 0x61, 0x70,
	0x69, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_echo_echo_proto_rawDescOnce sync.Once
	file_echo_echo_proto_rawDescData = file_echo_echo_proto_rawDesc
)

func file_echo_echo_proto_rawDescGZIP() []byte {
	file_echo_echo_proto_rawDescOnce.Do(func() {
		file_echo_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_echo_echo_proto_rawDescData)
	})
	return file_echo_echo_proto_rawDescData
}

var file_echo_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_echo_echo_proto_goTypes = []interface{}{
	(*EchoRequest)(nil),   // 0: echo.EchoRequest
	(*EchoMessageId)(nil), // 1: echo.EchoMessageId
	(*EchoResponse)(nil),  // 2: echo.EchoResponse
	(*Error)(nil),         // 3: echo.Error
}
var file_echo_echo_proto_depIdxs = []int32{
	1, // 0: echo.EchoRequest.em_id:type_name -> echo.EchoMessageId
	1, // 1: echo.EchoResponse.em_id:type_name -> echo.EchoMessageId
	0, // 2: echo.EchoService.PostEcho:input_type -> echo.EchoRequest
	0, // 3: echo.EchoService.GetEcho:input_type -> echo.EchoRequest
	2, // 4: echo.EchoService.PostEcho:output_type -> echo.EchoResponse
	2, // 5: echo.EchoService.GetEcho:output_type -> echo.EchoResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_echo_echo_proto_init() }
func file_echo_echo_proto_init() {
	if File_echo_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echo_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
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
		file_echo_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoMessageId); i {
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
		file_echo_echo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResponse); i {
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
		file_echo_echo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_echo_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_echo_echo_proto_goTypes,
		DependencyIndexes: file_echo_echo_proto_depIdxs,
		MessageInfos:      file_echo_echo_proto_msgTypes,
	}.Build()
	File_echo_echo_proto = out.File
	file_echo_echo_proto_rawDesc = nil
	file_echo_echo_proto_goTypes = nil
	file_echo_echo_proto_depIdxs = nil
}
