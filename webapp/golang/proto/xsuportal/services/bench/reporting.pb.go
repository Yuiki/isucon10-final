// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: xsuportal/services/bench/reporting.proto

package bench

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	resources "github.com/isucon/isucon10-final/webapp/golang/proto/xsuportal/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ReportBenchmarkResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId  int64                      `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Handle string                     `protobuf:"bytes,2,opt,name=handle,proto3" json:"handle,omitempty"`
	Nonce  int64                      `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Result *resources.BenchmarkResult `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ReportBenchmarkResultRequest) Reset() {
	*x = ReportBenchmarkResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xsuportal_services_bench_reporting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportBenchmarkResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportBenchmarkResultRequest) ProtoMessage() {}

func (x *ReportBenchmarkResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xsuportal_services_bench_reporting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportBenchmarkResultRequest.ProtoReflect.Descriptor instead.
func (*ReportBenchmarkResultRequest) Descriptor() ([]byte, []int) {
	return file_xsuportal_services_bench_reporting_proto_rawDescGZIP(), []int{0}
}

func (x *ReportBenchmarkResultRequest) GetJobId() int64 {
	if x != nil {
		return x.JobId
	}
	return 0
}

func (x *ReportBenchmarkResultRequest) GetHandle() string {
	if x != nil {
		return x.Handle
	}
	return ""
}

func (x *ReportBenchmarkResultRequest) GetNonce() int64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *ReportBenchmarkResultRequest) GetResult() *resources.BenchmarkResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type ReportBenchmarkResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AckedNonce int64 `protobuf:"varint,1,opt,name=acked_nonce,json=ackedNonce,proto3" json:"acked_nonce,omitempty"`
}

func (x *ReportBenchmarkResultResponse) Reset() {
	*x = ReportBenchmarkResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xsuportal_services_bench_reporting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportBenchmarkResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportBenchmarkResultResponse) ProtoMessage() {}

func (x *ReportBenchmarkResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xsuportal_services_bench_reporting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportBenchmarkResultResponse.ProtoReflect.Descriptor instead.
func (*ReportBenchmarkResultResponse) Descriptor() ([]byte, []int) {
	return file_xsuportal_services_bench_reporting_proto_rawDescGZIP(), []int{1}
}

func (x *ReportBenchmarkResultResponse) GetAckedNonce() int64 {
	if x != nil {
		return x.AckedNonce
	}
	return 0
}

var File_xsuportal_services_bench_reporting_proto protoreflect.FileDescriptor

var file_xsuportal_services_bench_reporting_proto_rawDesc = []byte{
	0x0a, 0x28, 0x78, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x78, 0x73, 0x75, 0x70,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x1a, 0x2a, 0x78, 0x73, 0x75, 0x70,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x78,
	0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61,
	0x72, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x40, 0x0a, 0x1d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d,
	0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x4e, 0x6f, 0x6e,
	0x63, 0x65, 0x32, 0xac, 0x01, 0x0a, 0x0f, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x98, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x3c, 0x2e, 0x78, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x65, 0x6e, 0x63,
	0x68, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72,
	0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d,
	0x2e, 0x78, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x31, 0x30, 0x2d,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x77, 0x65, 0x62, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x73, 0x75, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x65, 0x6e,
	0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xsuportal_services_bench_reporting_proto_rawDescOnce sync.Once
	file_xsuportal_services_bench_reporting_proto_rawDescData = file_xsuportal_services_bench_reporting_proto_rawDesc
)

func file_xsuportal_services_bench_reporting_proto_rawDescGZIP() []byte {
	file_xsuportal_services_bench_reporting_proto_rawDescOnce.Do(func() {
		file_xsuportal_services_bench_reporting_proto_rawDescData = protoimpl.X.CompressGZIP(file_xsuportal_services_bench_reporting_proto_rawDescData)
	})
	return file_xsuportal_services_bench_reporting_proto_rawDescData
}

var file_xsuportal_services_bench_reporting_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_xsuportal_services_bench_reporting_proto_goTypes = []interface{}{
	(*ReportBenchmarkResultRequest)(nil),  // 0: xsuportal.proto.services.bench.ReportBenchmarkResultRequest
	(*ReportBenchmarkResultResponse)(nil), // 1: xsuportal.proto.services.bench.ReportBenchmarkResultResponse
	(*resources.BenchmarkResult)(nil),     // 2: xsuportal.proto.resources.BenchmarkResult
}
var file_xsuportal_services_bench_reporting_proto_depIdxs = []int32{
	2, // 0: xsuportal.proto.services.bench.ReportBenchmarkResultRequest.result:type_name -> xsuportal.proto.resources.BenchmarkResult
	0, // 1: xsuportal.proto.services.bench.BenchmarkReport.ReportBenchmarkResult:input_type -> xsuportal.proto.services.bench.ReportBenchmarkResultRequest
	1, // 2: xsuportal.proto.services.bench.BenchmarkReport.ReportBenchmarkResult:output_type -> xsuportal.proto.services.bench.ReportBenchmarkResultResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_xsuportal_services_bench_reporting_proto_init() }
func file_xsuportal_services_bench_reporting_proto_init() {
	if File_xsuportal_services_bench_reporting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xsuportal_services_bench_reporting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportBenchmarkResultRequest); i {
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
		file_xsuportal_services_bench_reporting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportBenchmarkResultResponse); i {
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
			RawDescriptor: file_xsuportal_services_bench_reporting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_xsuportal_services_bench_reporting_proto_goTypes,
		DependencyIndexes: file_xsuportal_services_bench_reporting_proto_depIdxs,
		MessageInfos:      file_xsuportal_services_bench_reporting_proto_msgTypes,
	}.Build()
	File_xsuportal_services_bench_reporting_proto = out.File
	file_xsuportal_services_bench_reporting_proto_rawDesc = nil
	file_xsuportal_services_bench_reporting_proto_goTypes = nil
	file_xsuportal_services_bench_reporting_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BenchmarkReportClient is the client API for BenchmarkReport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BenchmarkReportClient interface {
	ReportBenchmarkResult(ctx context.Context, opts ...grpc.CallOption) (BenchmarkReport_ReportBenchmarkResultClient, error)
}

type benchmarkReportClient struct {
	cc grpc.ClientConnInterface
}

func NewBenchmarkReportClient(cc grpc.ClientConnInterface) BenchmarkReportClient {
	return &benchmarkReportClient{cc}
}

func (c *benchmarkReportClient) ReportBenchmarkResult(ctx context.Context, opts ...grpc.CallOption) (BenchmarkReport_ReportBenchmarkResultClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BenchmarkReport_serviceDesc.Streams[0], "/xsuportal.proto.services.bench.BenchmarkReport/ReportBenchmarkResult", opts...)
	if err != nil {
		return nil, err
	}
	x := &benchmarkReportReportBenchmarkResultClient{stream}
	return x, nil
}

type BenchmarkReport_ReportBenchmarkResultClient interface {
	Send(*ReportBenchmarkResultRequest) error
	Recv() (*ReportBenchmarkResultResponse, error)
	grpc.ClientStream
}

type benchmarkReportReportBenchmarkResultClient struct {
	grpc.ClientStream
}

func (x *benchmarkReportReportBenchmarkResultClient) Send(m *ReportBenchmarkResultRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *benchmarkReportReportBenchmarkResultClient) Recv() (*ReportBenchmarkResultResponse, error) {
	m := new(ReportBenchmarkResultResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BenchmarkReportServer is the server API for BenchmarkReport service.
type BenchmarkReportServer interface {
	ReportBenchmarkResult(BenchmarkReport_ReportBenchmarkResultServer) error
}

// UnimplementedBenchmarkReportServer can be embedded to have forward compatible implementations.
type UnimplementedBenchmarkReportServer struct {
}

func (*UnimplementedBenchmarkReportServer) ReportBenchmarkResult(BenchmarkReport_ReportBenchmarkResultServer) error {
	return status.Errorf(codes.Unimplemented, "method ReportBenchmarkResult not implemented")
}

func RegisterBenchmarkReportServer(s *grpc.Server, srv BenchmarkReportServer) {
	s.RegisterService(&_BenchmarkReport_serviceDesc, srv)
}

func _BenchmarkReport_ReportBenchmarkResult_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BenchmarkReportServer).ReportBenchmarkResult(&benchmarkReportReportBenchmarkResultServer{stream})
}

type BenchmarkReport_ReportBenchmarkResultServer interface {
	Send(*ReportBenchmarkResultResponse) error
	Recv() (*ReportBenchmarkResultRequest, error)
	grpc.ServerStream
}

type benchmarkReportReportBenchmarkResultServer struct {
	grpc.ServerStream
}

func (x *benchmarkReportReportBenchmarkResultServer) Send(m *ReportBenchmarkResultResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *benchmarkReportReportBenchmarkResultServer) Recv() (*ReportBenchmarkResultRequest, error) {
	m := new(ReportBenchmarkResultRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BenchmarkReport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xsuportal.proto.services.bench.BenchmarkReport",
	HandlerType: (*BenchmarkReportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReportBenchmarkResult",
			Handler:       _BenchmarkReport_ReportBenchmarkResult_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "xsuportal/services/bench/reporting.proto",
}
