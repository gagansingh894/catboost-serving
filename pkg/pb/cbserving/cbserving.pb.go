// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/cbserving/v1/cbserving.proto

package cbserving

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

type GetPredictionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelName string                             `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	InputData []*GetPredictionsRequest_InputData `protobuf:"bytes,3,rep,name=input_data,json=inputData,proto3" json:"input_data,omitempty"`
}

func (x *GetPredictionsRequest) Reset() {
	*x = GetPredictionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cbserving_v1_cbserving_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPredictionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPredictionsRequest) ProtoMessage() {}

func (x *GetPredictionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cbserving_v1_cbserving_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPredictionsRequest.ProtoReflect.Descriptor instead.
func (*GetPredictionsRequest) Descriptor() ([]byte, []int) {
	return file_proto_cbserving_v1_cbserving_proto_rawDescGZIP(), []int{0}
}

func (x *GetPredictionsRequest) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *GetPredictionsRequest) GetInputData() []*GetPredictionsRequest_InputData {
	if x != nil {
		return x.InputData
	}
	return nil
}

type GetPredictionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelName   string    `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	Predictions []float64 `protobuf:"fixed64,3,rep,packed,name=predictions,proto3" json:"predictions,omitempty"`
}

func (x *GetPredictionsResponse) Reset() {
	*x = GetPredictionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cbserving_v1_cbserving_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPredictionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPredictionsResponse) ProtoMessage() {}

func (x *GetPredictionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cbserving_v1_cbserving_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPredictionsResponse.ProtoReflect.Descriptor instead.
func (*GetPredictionsResponse) Descriptor() ([]byte, []int) {
	return file_proto_cbserving_v1_cbserving_proto_rawDescGZIP(), []int{1}
}

func (x *GetPredictionsResponse) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *GetPredictionsResponse) GetPredictions() []float64 {
	if x != nil {
		return x.Predictions
	}
	return nil
}

type GetPredictionsRequest_InputData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input map[string]float32 `protobuf:"bytes,1,rep,name=input,proto3" json:"input,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
}

func (x *GetPredictionsRequest_InputData) Reset() {
	*x = GetPredictionsRequest_InputData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cbserving_v1_cbserving_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPredictionsRequest_InputData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPredictionsRequest_InputData) ProtoMessage() {}

func (x *GetPredictionsRequest_InputData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cbserving_v1_cbserving_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPredictionsRequest_InputData.ProtoReflect.Descriptor instead.
func (*GetPredictionsRequest_InputData) Descriptor() ([]byte, []int) {
	return file_proto_cbserving_v1_cbserving_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GetPredictionsRequest_InputData) GetInput() map[string]float32 {
	if x != nil {
		return x.Input
	}
	return nil
}

var File_proto_cbserving_v1_cbserving_proto protoreflect.FileDescriptor

var file_proto_cbserving_v1_cbserving_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x62, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x62, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x62, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x22, 0x9c, 0x02, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x63, 0x62, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x95, 0x01, 0x0a, 0x09, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4e, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x62, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x38, 0x0a, 0x0a, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x59, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x01, 0x52,
	0x0b, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x70, 0x0a, 0x11,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x23, 0x2e, 0x63, 0x62, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d,
	0x5a, 0x1b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x3b, 0x63, 0x62, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cbserving_v1_cbserving_proto_rawDescOnce sync.Once
	file_proto_cbserving_v1_cbserving_proto_rawDescData = file_proto_cbserving_v1_cbserving_proto_rawDesc
)

func file_proto_cbserving_v1_cbserving_proto_rawDescGZIP() []byte {
	file_proto_cbserving_v1_cbserving_proto_rawDescOnce.Do(func() {
		file_proto_cbserving_v1_cbserving_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cbserving_v1_cbserving_proto_rawDescData)
	})
	return file_proto_cbserving_v1_cbserving_proto_rawDescData
}

var file_proto_cbserving_v1_cbserving_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_cbserving_v1_cbserving_proto_goTypes = []interface{}{
	(*GetPredictionsRequest)(nil),           // 0: cbserving.v1.GetPredictionsRequest
	(*GetPredictionsResponse)(nil),          // 1: cbserving.v1.GetPredictionsResponse
	(*GetPredictionsRequest_InputData)(nil), // 2: cbserving.v1.GetPredictionsRequest.InputData
	nil,                                     // 3: cbserving.v1.GetPredictionsRequest.InputData.InputEntry
}
var file_proto_cbserving_v1_cbserving_proto_depIdxs = []int32{
	2, // 0: cbserving.v1.GetPredictionsRequest.input_data:type_name -> cbserving.v1.GetPredictionsRequest.InputData
	3, // 1: cbserving.v1.GetPredictionsRequest.InputData.input:type_name -> cbserving.v1.GetPredictionsRequest.InputData.InputEntry
	0, // 2: cbserving.v1.DeploymentService.GetPredictions:input_type -> cbserving.v1.GetPredictionsRequest
	1, // 3: cbserving.v1.DeploymentService.GetPredictions:output_type -> cbserving.v1.GetPredictionsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_cbserving_v1_cbserving_proto_init() }
func file_proto_cbserving_v1_cbserving_proto_init() {
	if File_proto_cbserving_v1_cbserving_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cbserving_v1_cbserving_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPredictionsRequest); i {
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
		file_proto_cbserving_v1_cbserving_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPredictionsResponse); i {
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
		file_proto_cbserving_v1_cbserving_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPredictionsRequest_InputData); i {
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
			RawDescriptor: file_proto_cbserving_v1_cbserving_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cbserving_v1_cbserving_proto_goTypes,
		DependencyIndexes: file_proto_cbserving_v1_cbserving_proto_depIdxs,
		MessageInfos:      file_proto_cbserving_v1_cbserving_proto_msgTypes,
	}.Build()
	File_proto_cbserving_v1_cbserving_proto = out.File
	file_proto_cbserving_v1_cbserving_proto_rawDesc = nil
	file_proto_cbserving_v1_cbserving_proto_goTypes = nil
	file_proto_cbserving_v1_cbserving_proto_depIdxs = nil
}
