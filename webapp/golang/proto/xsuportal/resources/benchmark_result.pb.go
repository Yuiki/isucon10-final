// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: xsuportal/resources/benchmark_result.proto

package resources

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type BenchmarkResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Finished       bool                            `protobuf:"varint,1,opt,name=finished,proto3" json:"finished,omitempty"`
	Passed         bool                            `protobuf:"varint,2,opt,name=passed,proto3" json:"passed,omitempty"`
	Score          int64                           `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	ScoreBreakdown *BenchmarkResult_ScoreBreakdown `protobuf:"bytes,4,opt,name=score_breakdown,json=scoreBreakdown,proto3" json:"score_breakdown,omitempty"`
	Reason         string                          `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	MarkedAt       *timestamp.Timestamp            `protobuf:"bytes,6,opt,name=marked_at,json=markedAt,proto3" json:"marked_at,omitempty"`
}

func (x *BenchmarkResult) Reset() {
	*x = BenchmarkResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xsuportal_resources_benchmark_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BenchmarkResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenchmarkResult) ProtoMessage() {}

func (x *BenchmarkResult) ProtoReflect() protoreflect.Message {
	mi := &file_xsuportal_resources_benchmark_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BenchmarkResult.ProtoReflect.Descriptor instead.
func (*BenchmarkResult) Descriptor() ([]byte, []int) {
	return file_xsuportal_resources_benchmark_result_proto_rawDescGZIP(), []int{0}
}

func (x *BenchmarkResult) GetFinished() bool {
	if x != nil {
		return x.Finished
	}
	return false
}

func (x *BenchmarkResult) GetPassed() bool {
	if x != nil {
		return x.Passed
	}
	return false
}

func (x *BenchmarkResult) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *BenchmarkResult) GetScoreBreakdown() *BenchmarkResult_ScoreBreakdown {
	if x != nil {
		return x.ScoreBreakdown
	}
	return nil
}

func (x *BenchmarkResult) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *BenchmarkResult) GetMarkedAt() *timestamp.Timestamp {
	if x != nil {
		return x.MarkedAt
	}
	return nil
}

type BenchmarkResult_ScoreBreakdown struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Raw       int64 `protobuf:"varint,1,opt,name=raw,proto3" json:"raw,omitempty"`
	Deduction int64 `protobuf:"varint,2,opt,name=deduction,proto3" json:"deduction,omitempty"`
}

func (x *BenchmarkResult_ScoreBreakdown) Reset() {
	*x = BenchmarkResult_ScoreBreakdown{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xsuportal_resources_benchmark_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BenchmarkResult_ScoreBreakdown) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenchmarkResult_ScoreBreakdown) ProtoMessage() {}

func (x *BenchmarkResult_ScoreBreakdown) ProtoReflect() protoreflect.Message {
	mi := &file_xsuportal_resources_benchmark_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BenchmarkResult_ScoreBreakdown.ProtoReflect.Descriptor instead.
func (*BenchmarkResult_ScoreBreakdown) Descriptor() ([]byte, []int) {
	return file_xsuportal_resources_benchmark_result_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BenchmarkResult_ScoreBreakdown) GetRaw() int64 {
	if x != nil {
		return x.Raw
	}
	return 0
}

func (x *BenchmarkResult_ScoreBreakdown) GetDeduction() int64 {
	if x != nil {
		return x.Deduction
	}
	return 0
}

var File_xsuportal_resources_benchmark_result_proto protoreflect.FileDescriptor

var file_xsuportal_resources_benchmark_result_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x78, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x78, 0x73,
	0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x02, 0x0a, 0x0f, 0x42, 0x65, 0x6e,
	0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x73, 0x73,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x62, 0x0a, 0x0f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x62, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x39, 0x2e, 0x78, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x63,
	0x68, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x0e, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x40, 0x0a, 0x0e, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x4a, 0x5a,
	0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x73, 0x75, 0x63,
	0x6f, 0x6e, 0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x31, 0x30, 0x2d, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x2f, 0x77, 0x65, 0x62, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_xsuportal_resources_benchmark_result_proto_rawDescOnce sync.Once
	file_xsuportal_resources_benchmark_result_proto_rawDescData = file_xsuportal_resources_benchmark_result_proto_rawDesc
)

func file_xsuportal_resources_benchmark_result_proto_rawDescGZIP() []byte {
	file_xsuportal_resources_benchmark_result_proto_rawDescOnce.Do(func() {
		file_xsuportal_resources_benchmark_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_xsuportal_resources_benchmark_result_proto_rawDescData)
	})
	return file_xsuportal_resources_benchmark_result_proto_rawDescData
}

var file_xsuportal_resources_benchmark_result_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_xsuportal_resources_benchmark_result_proto_goTypes = []interface{}{
	(*BenchmarkResult)(nil),                // 0: xsuportal.proto.resources.BenchmarkResult
	(*BenchmarkResult_ScoreBreakdown)(nil), // 1: xsuportal.proto.resources.BenchmarkResult.ScoreBreakdown
	(*timestamp.Timestamp)(nil),            // 2: google.protobuf.Timestamp
}
var file_xsuportal_resources_benchmark_result_proto_depIdxs = []int32{
	1, // 0: xsuportal.proto.resources.BenchmarkResult.score_breakdown:type_name -> xsuportal.proto.resources.BenchmarkResult.ScoreBreakdown
	2, // 1: xsuportal.proto.resources.BenchmarkResult.marked_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_xsuportal_resources_benchmark_result_proto_init() }
func file_xsuportal_resources_benchmark_result_proto_init() {
	if File_xsuportal_resources_benchmark_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xsuportal_resources_benchmark_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenchmarkResult); i {
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
		file_xsuportal_resources_benchmark_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenchmarkResult_ScoreBreakdown); i {
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
			RawDescriptor: file_xsuportal_resources_benchmark_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xsuportal_resources_benchmark_result_proto_goTypes,
		DependencyIndexes: file_xsuportal_resources_benchmark_result_proto_depIdxs,
		MessageInfos:      file_xsuportal_resources_benchmark_result_proto_msgTypes,
	}.Build()
	File_xsuportal_resources_benchmark_result_proto = out.File
	file_xsuportal_resources_benchmark_result_proto_rawDesc = nil
	file_xsuportal_resources_benchmark_result_proto_goTypes = nil
	file_xsuportal_resources_benchmark_result_proto_depIdxs = nil
}
