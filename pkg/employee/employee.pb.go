// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.2
// source: employee/employee.proto

package employee

import (
	comm "./comm"
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

type JobType int32

const (
	JobType_LONG        JobType = 0 //长工
	JobType_SHORT       JobType = 1 //短工
	JobType_PARTTIMEJOB JobType = 2 //兼职
)

// Enum value maps for JobType.
var (
	JobType_name = map[int32]string{
		0: "LONG",
		1: "SHORT",
		2: "PARTTIMEJOB",
	}
	JobType_value = map[string]int32{
		"LONG":        0,
		"SHORT":       1,
		"PARTTIMEJOB": 2,
	}
)

func (x JobType) Enum() *JobType {
	p := new(JobType)
	*p = x
	return p
}

func (x JobType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobType) Descriptor() protoreflect.EnumDescriptor {
	return file_employee_employee_proto_enumTypes[0].Descriptor()
}

func (JobType) Type() protoreflect.EnumType {
	return &file_employee_employee_proto_enumTypes[0]
}

func (x JobType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobType.Descriptor instead.
func (JobType) EnumDescriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{0}
}

// 请求
type OperateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	JobId int32 `protobuf:"varint,2,opt,name=jobId,proto3" json:"jobId,omitempty"`
}

func (x *OperateRequest) Reset() {
	*x = OperateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateRequest) ProtoMessage() {}

func (x *OperateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateRequest.ProtoReflect.Descriptor instead.
func (*OperateRequest) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{0}
}

func (x *OperateRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OperateRequest) GetJobId() int32 {
	if x != nil {
		return x.JobId
	}
	return 0
}

type FindRecordsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindRecordsReq) Reset() {
	*x = FindRecordsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRecordsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRecordsReq) ProtoMessage() {}

func (x *FindRecordsReq) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRecordsReq.ProtoReflect.Descriptor instead.
func (*FindRecordsReq) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{1}
}

func (x *FindRecordsReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SelfInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SimpleInfo  *SimpleInformation `protobuf:"bytes,1,opt,name=simpleInfo,proto3" json:"simpleInfo,omitempty"`
	IsMarry     int32              `protobuf:"varint,2,opt,name=isMarry,proto3" json:"isMarry,omitempty"`         //是否结婚
	Education   int32              `protobuf:"varint,3,opt,name=education,proto3" json:"education,omitempty"`     //学历，可以分几种；1表示小学，2表示初中，3表示高中，以此类推
	JobStatus   int32              `protobuf:"varint,4,opt,name=jobStatus,proto3" json:"jobStatus,omitempty"`     //求职状态，可以定几种类类型，例如1为在职，2为离职不到一个月，3为离职一个月以上
	Home        string             `protobuf:"bytes,5,opt,name=home,proto3" json:"home,omitempty"`                //家庭地址
	ExpectMoney int32              `protobuf:"varint,6,opt,name=expectMoney,proto3" json:"expectMoney,omitempty"` //期望薪资
}

func (x *SelfInformation) Reset() {
	*x = SelfInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelfInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelfInformation) ProtoMessage() {}

func (x *SelfInformation) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelfInformation.ProtoReflect.Descriptor instead.
func (*SelfInformation) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{2}
}

func (x *SelfInformation) GetSimpleInfo() *SimpleInformation {
	if x != nil {
		return x.SimpleInfo
	}
	return nil
}

func (x *SelfInformation) GetIsMarry() int32 {
	if x != nil {
		return x.IsMarry
	}
	return 0
}

func (x *SelfInformation) GetEducation() int32 {
	if x != nil {
		return x.Education
	}
	return 0
}

func (x *SelfInformation) GetJobStatus() int32 {
	if x != nil {
		return x.JobStatus
	}
	return 0
}

func (x *SelfInformation) GetHome() string {
	if x != nil {
		return x.Home
	}
	return ""
}

func (x *SelfInformation) GetExpectMoney() int32 {
	if x != nil {
		return x.ExpectMoney
	}
	return 0
}

type SimpleInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                 //用户id
	RealName         string  `protobuf:"bytes,2,opt,name=realName,proto3" json:"realName,omitempty"`                      //真实名字
	Sex              int32   `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`                               //性别，0男，1女
	WorkYear         int32   `protobuf:"varint,4,opt,name=workYear,proto3" json:"workYear,omitempty"`                     //工作年限
	ExpectedLocation string  `protobuf:"bytes,5,opt,name=expectedLocation,proto3" json:"expectedLocation,omitempty"`      //期望工作地点
	Industry         string  `protobuf:"bytes,6,opt,name=industry,proto3" json:"industry,omitempty"`                      //行业
	JobType          JobType `protobuf:"varint,7,opt,name=jobType,proto3,enum=employee.JobType" json:"jobType,omitempty"` //
}

func (x *SimpleInformation) Reset() {
	*x = SimpleInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleInformation) ProtoMessage() {}

func (x *SimpleInformation) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleInformation.ProtoReflect.Descriptor instead.
func (*SimpleInformation) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{3}
}

func (x *SimpleInformation) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SimpleInformation) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *SimpleInformation) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *SimpleInformation) GetWorkYear() int32 {
	if x != nil {
		return x.WorkYear
	}
	return 0
}

func (x *SimpleInformation) GetExpectedLocation() string {
	if x != nil {
		return x.ExpectedLocation
	}
	return ""
}

func (x *SimpleInformation) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *SimpleInformation) GetJobType() JobType {
	if x != nil {
		return x.JobType
	}
	return JobType_LONG
}

type FindEmployeeInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindEmployeeInfoReq) Reset() {
	*x = FindEmployeeInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEmployeeInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEmployeeInfoReq) ProtoMessage() {}

func (x *FindEmployeeInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEmployeeInfoReq.ProtoReflect.Descriptor instead.
func (*FindEmployeeInfoReq) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{4}
}

func (x *FindEmployeeInfoReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SplitPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Size  int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *SplitPage) Reset() {
	*x = SplitPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplitPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitPage) ProtoMessage() {}

func (x *SplitPage) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitPage.ProtoReflect.Descriptor instead.
func (*SplitPage) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{5}
}

func (x *SplitPage) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *SplitPage) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FindEmployeeSimpleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *comm.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Employees []*SimpleInformation `protobuf:"bytes,2,rep,name=employees,proto3" json:"employees,omitempty"`
}

func (x *FindEmployeeSimpleResponse) Reset() {
	*x = FindEmployeeSimpleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEmployeeSimpleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEmployeeSimpleResponse) ProtoMessage() {}

func (x *FindEmployeeSimpleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEmployeeSimpleResponse.ProtoReflect.Descriptor instead.
func (*FindEmployeeSimpleResponse) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{6}
}

func (x *FindEmployeeSimpleResponse) GetHeader() *comm.ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *FindEmployeeSimpleResponse) GetEmployees() []*SimpleInformation {
	if x != nil {
		return x.Employees
	}
	return nil
}

type FindEmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header   *comm.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	SelfInfo *SelfInformation     `protobuf:"bytes,2,opt,name=selfInfo,proto3" json:"selfInfo,omitempty"`
}

func (x *FindEmployeeResponse) Reset() {
	*x = FindEmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEmployeeResponse) ProtoMessage() {}

func (x *FindEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEmployeeResponse.ProtoReflect.Descriptor instead.
func (*FindEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{7}
}

func (x *FindEmployeeResponse) GetHeader() *comm.ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *FindEmployeeResponse) GetSelfInfo() *SelfInformation {
	if x != nil {
		return x.SelfInfo
	}
	return nil
}

type OperateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *comm.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	JobIds []int32              `protobuf:"varint,2,rep,packed,name=jobIds,proto3" json:"jobIds,omitempty"`
}

func (x *OperateResponse) Reset() {
	*x = OperateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_employee_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateResponse) ProtoMessage() {}

func (x *OperateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_employee_employee_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateResponse.ProtoReflect.Descriptor instead.
func (*OperateResponse) Descriptor() ([]byte, []int) {
	return file_employee_employee_proto_rawDescGZIP(), []int{8}
}

func (x *OperateResponse) GetHeader() *comm.ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *OperateResponse) GetJobIds() []int32 {
	if x != nil {
		return x.JobIds
	}
	return nil
}

var File_employee_employee_proto protoreflect.FileDescriptor

var file_employee_employee_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22,
	0x20, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xda, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x6c, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x4d, 0x61, 0x72, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x73, 0x4d, 0x61, 0x72, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6a, 0x6f,
	0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6a,
	0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x22, 0xe2,
	0x01, 0x0a, 0x11, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73,
	0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x59, 0x65, 0x61, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x59, 0x65, 0x61, 0x72, 0x12, 0x2a,
	0x0a, 0x10, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e,
	0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e,
	0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x12, 0x2b, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x09, 0x53, 0x70,
	0x6c, 0x69, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x22, 0x85, 0x01, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x39,
	0x0a, 0x09, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x22, 0x7b, 0x0a, 0x14, 0x46, 0x69, 0x6e,
	0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x35, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x53, 0x65, 0x6c,
	0x66, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65,
	0x6c, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x57, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x49, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x73, 0x2a,
	0x2f, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f,
	0x4e, 0x47, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x48, 0x4f, 0x52, 0x54, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x50, 0x41, 0x52, 0x54, 0x54, 0x49, 0x4d, 0x45, 0x4a, 0x4f, 0x42, 0x10, 0x02,
	0x32, 0xdf, 0x03, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x40, 0x0a,
	0x07, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x07, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x47, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x12, 0x18, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x2e, 0x53, 0x65, 0x6c, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x17, 0x66, 0x69, 0x6e, 0x64,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x1d, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x2e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x50, 0x61, 0x67, 0x65, 0x1a, 0x24, 0x2e, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x27, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x70, 0x74, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x50, 0x01, 0x5a,
	0x0a, 0x2e, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_employee_employee_proto_rawDescOnce sync.Once
	file_employee_employee_proto_rawDescData = file_employee_employee_proto_rawDesc
)

func file_employee_employee_proto_rawDescGZIP() []byte {
	file_employee_employee_proto_rawDescOnce.Do(func() {
		file_employee_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_employee_employee_proto_rawDescData)
	})
	return file_employee_employee_proto_rawDescData
}

var file_employee_employee_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_employee_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_employee_employee_proto_goTypes = []interface{}{
	(JobType)(0),                       // 0: employee.JobType
	(*OperateRequest)(nil),             // 1: employee.OperateRequest
	(*FindRecordsReq)(nil),             // 2: employee.FindRecordsReq
	(*SelfInformation)(nil),            // 3: employee.SelfInformation
	(*SimpleInformation)(nil),          // 4: employee.SimpleInformation
	(*FindEmployeeInfoReq)(nil),        // 5: employee.FindEmployeeInfoReq
	(*SplitPage)(nil),                  // 6: employee.SplitPage
	(*FindEmployeeSimpleResponse)(nil), // 7: employee.FindEmployeeSimpleResponse
	(*FindEmployeeResponse)(nil),       // 8: employee.FindEmployeeResponse
	(*OperateResponse)(nil),            // 9: employee.OperateResponse
	(*comm.ResponseHeader)(nil),        // 10: comm.ResponseHeader
}
var file_employee_employee_proto_depIdxs = []int32{
	4,  // 0: employee.SelfInformation.simpleInfo:type_name -> employee.SimpleInformation
	0,  // 1: employee.SimpleInformation.jobType:type_name -> employee.JobType
	10, // 2: employee.FindEmployeeSimpleResponse.header:type_name -> comm.ResponseHeader
	4,  // 3: employee.FindEmployeeSimpleResponse.employees:type_name -> employee.SimpleInformation
	10, // 4: employee.FindEmployeeResponse.header:type_name -> comm.ResponseHeader
	3,  // 5: employee.FindEmployeeResponse.selfInfo:type_name -> employee.SelfInformation
	10, // 6: employee.OperateResponse.header:type_name -> comm.ResponseHeader
	1,  // 7: employee.Employee.deliver:input_type -> employee.OperateRequest
	1,  // 8: employee.Employee.browses:input_type -> employee.OperateRequest
	2,  // 9: employee.Employee.deliverRecords:input_type -> employee.FindRecordsReq
	3,  // 10: employee.Employee.completeSelfInformation:input_type -> employee.SelfInformation
	5,  // 11: employee.Employee.findEmployeeInformation:input_type -> employee.FindEmployeeInfoReq
	6,  // 12: employee.Employee.findEmployeeSimpleInformation:input_type -> employee.SplitPage
	9,  // 13: employee.Employee.deliver:output_type -> employee.OperateResponse
	9,  // 14: employee.Employee.browses:output_type -> employee.OperateResponse
	9,  // 15: employee.Employee.deliverRecords:output_type -> employee.OperateResponse
	10, // 16: employee.Employee.completeSelfInformation:output_type -> comm.ResponseHeader
	8,  // 17: employee.Employee.findEmployeeInformation:output_type -> employee.FindEmployeeResponse
	7,  // 18: employee.Employee.findEmployeeSimpleInformation:output_type -> employee.FindEmployeeSimpleResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_employee_employee_proto_init() }
func file_employee_employee_proto_init() {
	if File_employee_employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_employee_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateRequest); i {
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
		file_employee_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRecordsReq); i {
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
		file_employee_employee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelfInformation); i {
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
		file_employee_employee_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleInformation); i {
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
		file_employee_employee_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEmployeeInfoReq); i {
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
		file_employee_employee_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplitPage); i {
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
		file_employee_employee_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEmployeeSimpleResponse); i {
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
		file_employee_employee_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEmployeeResponse); i {
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
		file_employee_employee_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateResponse); i {
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
			RawDescriptor: file_employee_employee_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_employee_employee_proto_goTypes,
		DependencyIndexes: file_employee_employee_proto_depIdxs,
		EnumInfos:         file_employee_employee_proto_enumTypes,
		MessageInfos:      file_employee_employee_proto_msgTypes,
	}.Build()
	File_employee_employee_proto = out.File
	file_employee_employee_proto_rawDesc = nil
	file_employee_employee_proto_goTypes = nil
	file_employee_employee_proto_depIdxs = nil
}
