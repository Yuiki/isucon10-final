// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: xsuportal/resources/benchmark_job.proto

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

type BenchmarkJob_Status int32

const (
	BenchmarkJob_PENDING   BenchmarkJob_Status = 0
	BenchmarkJob_SENT      BenchmarkJob_Status = 1
	BenchmarkJob_RUNNING   BenchmarkJob_Status = 2
	BenchmarkJob_ERRORED   BenchmarkJob_Status = 3
	BenchmarkJob_CANCELLED BenchmarkJob_Status = 4
	BenchmarkJob_FINISHED  BenchmarkJob_Status = 5
)

// Enum value maps for BenchmarkJob_Status.
var (
	BenchmarkJob_Status_name = map[int32]string{
		0: "PENDING",
		1: "SENT",
		2: "RUNNING",
		3: "ERRORED",
		4: "CANCELLED",
		5: "FINISHED",
	}
	BenchmarkJob_Status_value = map[string]int32{
		"PENDING":   0,
		"SENT":      1,
		"RUNNING":   2,
		"ERRORED":   3,
		"CANCELLED": 4,
		"FINISHED":  5,
	}
)

func (x BenchmarkJob_Status) Enum() *BenchmarkJob_Status {
	p := new(BenchmarkJob_Status)
	*p = x
	return p
}

func (x BenchmarkJob_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BenchmarkJob_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_xsuportal_resources_benchmark_job_proto_enumTypes[0].Descriptor()
}

func (BenchmarkJob_Status) Type() protoreflect.EnumType {
	return &file_xsuportal_resources_benchmark_job_proto_enumTypes[0]
}

func (x BenchmarkJob_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BenchmarkJob_Status.Descriptor instead.
func (BenchmarkJob_Status) EnumDescriptor() ([]byte, []int) {
	return file_xsuportal_resources_benchmark_job_proto_rawDescGZIP(), []int{0, 0}
}

type BenchmarkJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TeamId int64 `protobuf:"varint,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	// int64 target_id = 3;
	Status     BenchmarkJob_Status  `protobuf:"varint,4,opt,name=status,proto3,enum=xsuportal.proto.resources.BenchmarkJob_Status" json:"status,omitempty"`
	CreatedAt  *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	StartedAt  *timestamp.Timestamp `protobuf:"bytes,7,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	FinishedAt *timestamp.Timestamp `protobuf:"bytes,8,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	Team       *Team                `protobuf:"bytes,16,opt,name=team,proto3" json:"team,omitempty"`
	// target & result are only available at GetBenchmarkJobResponse
	// ContestantInstance target = 17;
	Result         *BenchmarkResult `protobuf:"bytes,18,opt,name=result,proto3" json:"result,omitempty"`
	TargetHostname string           `protobuf:"bytes,30,opt,name=target_hostname,json=targetHostname,proto3" json:"target_hostname,omitempty"`
}

func (x *BenchmarkJob) Reset() {
	*x = BenchmarkJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xsuportal_resources_benchmark_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BenchmarkJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenchmarkJob) ProtoMessage() {}

func (x *BenchmarkJob) ProtoReflect() protoreflect.Message {
	mi := &file_xsuportal_resources_benchmark_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BenchmarkJob.ProtoReflect.Descriptor instead.
func (*BenchmarkJob) Descriptor() ([]byte, []int) {
	return file_xsuportal_resources_benchmark_job_proto_rawDescGZIP(), []int{0}
}

func (x *BenchmarkJob) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BenchmarkJob) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *BenchmarkJob) GetStatus() BenchmarkJob_Status {
	if x != nil {
		return x.Status
	}
	return BenchmarkJob_PENDING
}

func (x *BenchmarkJob) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *BenchmarkJob) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *BenchmarkJob) GetStartedAt() *timestamp.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *BenchmarkJob) GetFinishedAt() *timestamp.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

func (x *BenchmarkJob) GetTeam() *Team {
	if x != nil {
		return x.Team
	}
	return nil
}

func (x *BenchmarkJob) GetResult() *BenchmarkResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *BenchmarkJob) GetTargetHostname() string {
	if x != nil {
		return x.TargetHostname
	}
	return ""
}

var File_xsuportal_resources_benchmark_job_proto protoreflect.FileDescriptor

var file_xsuportal_resources_benchmark_job_proto_rawDesc = []byte{
	0x0a, 0x27, 0x78, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x5f,
	0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x78, 0x73, 0x75, 0x70, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x1a, 0x2a, 0x78, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d,
	0x61, 0x72, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x78, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe7, 0x04, 0x0a, 0x0c, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x4a,
	0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x78, 0x73,
	0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72,
	0x6b, 0x4a, 0x6f, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x78, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x42, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x78, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45,
	0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0c, 0x0a,
	0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x05, 0x42, 0x4a, 0x5a, 0x48, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e,
	0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x31, 0x30, 0x2d, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2f,
	0x77, 0x65, 0x62, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xsuportal_resources_benchmark_job_proto_rawDescOnce sync.Once
	file_xsuportal_resources_benchmark_job_proto_rawDescData = file_xsuportal_resources_benchmark_job_proto_rawDesc
)

func file_xsuportal_resources_benchmark_job_proto_rawDescGZIP() []byte {
	file_xsuportal_resources_benchmark_job_proto_rawDescOnce.Do(func() {
		file_xsuportal_resources_benchmark_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_xsuportal_resources_benchmark_job_proto_rawDescData)
	})
	return file_xsuportal_resources_benchmark_job_proto_rawDescData
}

var file_xsuportal_resources_benchmark_job_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_xsuportal_resources_benchmark_job_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_xsuportal_resources_benchmark_job_proto_goTypes = []interface{}{
	(BenchmarkJob_Status)(0),    // 0: xsuportal.proto.resources.BenchmarkJob.Status
	(*BenchmarkJob)(nil),        // 1: xsuportal.proto.resources.BenchmarkJob
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*Team)(nil),                // 3: xsuportal.proto.resources.Team
	(*BenchmarkResult)(nil),     // 4: xsuportal.proto.resources.BenchmarkResult
}
var file_xsuportal_resources_benchmark_job_proto_depIdxs = []int32{
	0, // 0: xsuportal.proto.resources.BenchmarkJob.status:type_name -> xsuportal.proto.resources.BenchmarkJob.Status
	2, // 1: xsuportal.proto.resources.BenchmarkJob.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: xsuportal.proto.resources.BenchmarkJob.updated_at:type_name -> google.protobuf.Timestamp
	2, // 3: xsuportal.proto.resources.BenchmarkJob.started_at:type_name -> google.protobuf.Timestamp
	2, // 4: xsuportal.proto.resources.BenchmarkJob.finished_at:type_name -> google.protobuf.Timestamp
	3, // 5: xsuportal.proto.resources.BenchmarkJob.team:type_name -> xsuportal.proto.resources.Team
	4, // 6: xsuportal.proto.resources.BenchmarkJob.result:type_name -> xsuportal.proto.resources.BenchmarkResult
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_xsuportal_resources_benchmark_job_proto_init() }
func file_xsuportal_resources_benchmark_job_proto_init() {
	if File_xsuportal_resources_benchmark_job_proto != nil {
		return
	}
	file_xsuportal_resources_benchmark_result_proto_init()
	file_xsuportal_resources_team_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_xsuportal_resources_benchmark_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenchmarkJob); i {
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
			RawDescriptor: file_xsuportal_resources_benchmark_job_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xsuportal_resources_benchmark_job_proto_goTypes,
		DependencyIndexes: file_xsuportal_resources_benchmark_job_proto_depIdxs,
		EnumInfos:         file_xsuportal_resources_benchmark_job_proto_enumTypes,
		MessageInfos:      file_xsuportal_resources_benchmark_job_proto_msgTypes,
	}.Build()
	File_xsuportal_resources_benchmark_job_proto = out.File
	file_xsuportal_resources_benchmark_job_proto_rawDesc = nil
	file_xsuportal_resources_benchmark_job_proto_goTypes = nil
	file_xsuportal_resources_benchmark_job_proto_depIdxs = nil
}
