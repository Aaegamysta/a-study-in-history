// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: a_study_in_history.proto

package gen

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

type TargetEvents int32

const (
	TargetEvents_TARGET_EVENTS_UNSPECIFIED TargetEvents = 0
	TargetEvents_TARGET_EVENTS_HISTORY     TargetEvents = 1
	TargetEvents_TARGET_EVENTS_BIRTH       TargetEvents = 2
	TargetEvents_TARGET_EVENTS_DEATH       TargetEvents = 3
	TargetEvents_TARGET_EVENTS_HOLIDAY     TargetEvents = 4
	TargetEvents_TARGET_EVENTS_ALL         TargetEvents = 5
)

// Enum value maps for TargetEvents.
var (
	TargetEvents_name = map[int32]string{
		0: "TARGET_EVENTS_UNSPECIFIED",
		1: "TARGET_EVENTS_HISTORY",
		2: "TARGET_EVENTS_BIRTH",
		3: "TARGET_EVENTS_DEATH",
		4: "TARGET_EVENTS_HOLIDAY",
		5: "TARGET_EVENTS_ALL",
	}
	TargetEvents_value = map[string]int32{
		"TARGET_EVENTS_UNSPECIFIED": 0,
		"TARGET_EVENTS_HISTORY":     1,
		"TARGET_EVENTS_BIRTH":       2,
		"TARGET_EVENTS_DEATH":       3,
		"TARGET_EVENTS_HOLIDAY":     4,
		"TARGET_EVENTS_ALL":         5,
	}
)

func (x TargetEvents) Enum() *TargetEvents {
	p := new(TargetEvents)
	*p = x
	return p
}

func (x TargetEvents) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TargetEvents) Descriptor() protoreflect.EnumDescriptor {
	return file_a_study_in_history_proto_enumTypes[0].Descriptor()
}

func (TargetEvents) Type() protoreflect.EnumType {
	return &file_a_study_in_history_proto_enumTypes[0]
}

func (x TargetEvents) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TargetEvents.Descriptor instead.
func (TargetEvents) EnumDescriptor() ([]byte, []int) {
	return file_a_study_in_history_proto_rawDescGZIP(), []int{0}
}

type ImportStatus int32

const (
	ImportStatus_IMPORT_STATUS_UNSPECIFIED     ImportStatus = 0
	ImportStatus_IMPORT_STATUS_SUCCESS         ImportStatus = 1
	ImportStatus_IMPORT_STATUS_PARTIAL_SUCCESS ImportStatus = 2
	ImportStatus_IMPORT_STATUS_FAILED          ImportStatus = 3
)

// Enum value maps for ImportStatus.
var (
	ImportStatus_name = map[int32]string{
		0: "IMPORT_STATUS_UNSPECIFIED",
		1: "IMPORT_STATUS_SUCCESS",
		2: "IMPORT_STATUS_PARTIAL_SUCCESS",
		3: "IMPORT_STATUS_FAILED",
	}
	ImportStatus_value = map[string]int32{
		"IMPORT_STATUS_UNSPECIFIED":     0,
		"IMPORT_STATUS_SUCCESS":         1,
		"IMPORT_STATUS_PARTIAL_SUCCESS": 2,
		"IMPORT_STATUS_FAILED":          3,
	}
)

func (x ImportStatus) Enum() *ImportStatus {
	p := new(ImportStatus)
	*p = x
	return p
}

func (x ImportStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImportStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_a_study_in_history_proto_enumTypes[1].Descriptor()
}

func (ImportStatus) Type() protoreflect.EnumType {
	return &file_a_study_in_history_proto_enumTypes[1]
}

func (x ImportStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImportStatus.Descriptor instead.
func (ImportStatus) EnumDescriptor() ([]byte, []int) {
	return file_a_study_in_history_proto_rawDescGZIP(), []int{1}
}

type MissedEvents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  Type  `protobuf:"varint,1,opt,name=type,proto3,enum=daemon.Type" json:"type,omitempty"`
	Day   int64 `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	Month int64 `protobuf:"varint,3,opt,name=month,proto3" json:"month,omitempty"`
}

func (x *MissedEvents) Reset() {
	*x = MissedEvents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_a_study_in_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MissedEvents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MissedEvents) ProtoMessage() {}

func (x *MissedEvents) ProtoReflect() protoreflect.Message {
	mi := &file_a_study_in_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MissedEvents.ProtoReflect.Descriptor instead.
func (*MissedEvents) Descriptor() ([]byte, []int) {
	return file_a_study_in_history_proto_rawDescGZIP(), []int{0}
}

func (x *MissedEvents) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNSPECIFIED
}

func (x *MissedEvents) GetDay() int64 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *MissedEvents) GetMonth() int64 {
	if x != nil {
		return x.Month
	}
	return 0
}

type ImportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ImportRequest) Reset() {
	*x = ImportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_a_study_in_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportRequest) ProtoMessage() {}

func (x *ImportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_a_study_in_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportRequest.ProtoReflect.Descriptor instead.
func (*ImportRequest) Descriptor() ([]byte, []int) {
	return file_a_study_in_history_proto_rawDescGZIP(), []int{1}
}

type ImportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       ImportStatus    `protobuf:"varint,1,opt,name=status,proto3,enum=daemon.ImportStatus" json:"status,omitempty"`
	MissedEvents []*MissedEvents `protobuf:"bytes,2,rep,name=missedEvents,proto3" json:"missedEvents,omitempty"`
	ImportedOn   int64           `protobuf:"varint,3,opt,name=importedOn,proto3" json:"importedOn,omitempty"`
}

func (x *ImportResponse) Reset() {
	*x = ImportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_a_study_in_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportResponse) ProtoMessage() {}

func (x *ImportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_a_study_in_history_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportResponse.ProtoReflect.Descriptor instead.
func (*ImportResponse) Descriptor() ([]byte, []int) {
	return file_a_study_in_history_proto_rawDescGZIP(), []int{2}
}

func (x *ImportResponse) GetStatus() ImportStatus {
	if x != nil {
		return x.Status
	}
	return ImportStatus_IMPORT_STATUS_UNSPECIFIED
}

func (x *ImportResponse) GetMissedEvents() []*MissedEvents {
	if x != nil {
		return x.MissedEvents
	}
	return nil
}

func (x *ImportResponse) GetImportedOn() int64 {
	if x != nil {
		return x.ImportedOn
	}
	return 0
}

type ResynchronizeForRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetEvents TargetEvents `protobuf:"varint,1,opt,name=targetEvents,proto3,enum=daemon.TargetEvents" json:"targetEvents,omitempty"`
	Day          int64        `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	Month        int64        `protobuf:"varint,3,opt,name=month,proto3" json:"month,omitempty"`
}

func (x *ResynchronizeForRequest) Reset() {
	*x = ResynchronizeForRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_a_study_in_history_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResynchronizeForRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResynchronizeForRequest) ProtoMessage() {}

func (x *ResynchronizeForRequest) ProtoReflect() protoreflect.Message {
	mi := &file_a_study_in_history_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResynchronizeForRequest.ProtoReflect.Descriptor instead.
func (*ResynchronizeForRequest) Descriptor() ([]byte, []int) {
	return file_a_study_in_history_proto_rawDescGZIP(), []int{3}
}

func (x *ResynchronizeForRequest) GetTargetEvents() TargetEvents {
	if x != nil {
		return x.TargetEvents
	}
	return TargetEvents_TARGET_EVENTS_UNSPECIFIED
}

func (x *ResynchronizeForRequest) GetDay() int64 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *ResynchronizeForRequest) GetMonth() int64 {
	if x != nil {
		return x.Month
	}
	return 0
}

type ResynchronizeForResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResynchronizeForResponse) Reset() {
	*x = ResynchronizeForResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_a_study_in_history_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResynchronizeForResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResynchronizeForResponse) ProtoMessage() {}

func (x *ResynchronizeForResponse) ProtoReflect() protoreflect.Message {
	mi := &file_a_study_in_history_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResynchronizeForResponse.ProtoReflect.Descriptor instead.
func (*ResynchronizeForResponse) Descriptor() ([]byte, []int) {
	return file_a_study_in_history_proto_rawDescGZIP(), []int{4}
}

func (x *ResynchronizeForResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type GetEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetEvents TargetEvents `protobuf:"varint,1,opt,name=targetEvents,proto3,enum=daemon.TargetEvents" json:"targetEvents,omitempty"`
	Day          int64        `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	Month        int64        `protobuf:"varint,3,opt,name=month,proto3" json:"month,omitempty"`
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_a_study_in_history_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_a_study_in_history_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return file_a_study_in_history_proto_rawDescGZIP(), []int{5}
}

func (x *GetEventsRequest) GetTargetEvents() TargetEvents {
	if x != nil {
		return x.TargetEvents
	}
	return TargetEvents_TARGET_EVENTS_UNSPECIFIED
}

func (x *GetEventsRequest) GetDay() int64 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *GetEventsRequest) GetMonth() int64 {
	if x != nil {
		return x.Month
	}
	return 0
}

type GetEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventsCollection *EventsCollection `protobuf:"bytes,1,opt,name=eventsCollection,proto3" json:"eventsCollection,omitempty"`
	Day              int64             `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	Month            int64             `protobuf:"varint,3,opt,name=month,proto3" json:"month,omitempty"`
	Type             Type              `protobuf:"varint,4,opt,name=type,proto3,enum=daemon.Type" json:"type,omitempty"`
}

func (x *GetEventsResponse) Reset() {
	*x = GetEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_a_study_in_history_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsResponse) ProtoMessage() {}

func (x *GetEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_a_study_in_history_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsResponse.ProtoReflect.Descriptor instead.
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return file_a_study_in_history_proto_rawDescGZIP(), []int{6}
}

func (x *GetEventsResponse) GetEventsCollection() *EventsCollection {
	if x != nil {
		return x.EventsCollection
	}
	return nil
}

func (x *GetEventsResponse) GetDay() int64 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *GetEventsResponse) GetMonth() int64 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *GetEventsResponse) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNSPECIFIED
}

var File_a_study_in_history_proto protoreflect.FileDescriptor

var file_a_study_in_history_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x69, 0x6e, 0x5f, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x61, 0x65, 0x6d,
	0x6f, 0x6e, 0x1a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x58, 0x0a, 0x0c, 0x4d, 0x69, 0x73, 0x73, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x64, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x22, 0x0f, 0x0a, 0x0d, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x98, 0x01, 0x0a, 0x0e,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x0c,
	0x6d, 0x69, 0x73, 0x73, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x69, 0x73, 0x73,
	0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0c, 0x6d, 0x69, 0x73, 0x73, 0x65, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x4f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x22, 0x7b, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x79, 0x6e, 0x63,
	0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x38, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e,
	0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0c, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x22, 0x32, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f,
	0x6e, 0x69, 0x7a, 0x65, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x74, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x22, 0xa3, 0x01,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x2a, 0xac, 0x01, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x53, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x01, 0x12, 0x17,
	0x0a, 0x13, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x53, 0x5f,
	0x42, 0x49, 0x52, 0x54, 0x48, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x41, 0x52, 0x47, 0x45,
	0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x44, 0x45, 0x41, 0x54, 0x48, 0x10, 0x03,
	0x12, 0x19, 0x0a, 0x15, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x53, 0x5f, 0x48, 0x4f, 0x4c, 0x49, 0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x54,
	0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x41, 0x4c, 0x4c,
	0x10, 0x05, 0x2a, 0x85, 0x01, 0x0a, 0x0c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x21, 0x0a,
	0x1d, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50,
	0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02,
	0x12, 0x18, 0x0a, 0x14, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x32, 0xe6, 0x01, 0x0a, 0x0f, 0x41,
	0x53, 0x74, 0x75, 0x64, 0x79, 0x49, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x37,
	0x0a, 0x06, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x79, 0x6e,
	0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x46, 0x6f, 0x72, 0x12, 0x1f, 0x2e, 0x64, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69,
	0x7a, 0x65, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64,
	0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e,
	0x69, 0x7a, 0x65, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x12, 0x18,
	0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_a_study_in_history_proto_rawDescOnce sync.Once
	file_a_study_in_history_proto_rawDescData = file_a_study_in_history_proto_rawDesc
)

func file_a_study_in_history_proto_rawDescGZIP() []byte {
	file_a_study_in_history_proto_rawDescOnce.Do(func() {
		file_a_study_in_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_a_study_in_history_proto_rawDescData)
	})
	return file_a_study_in_history_proto_rawDescData
}

var file_a_study_in_history_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_a_study_in_history_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_a_study_in_history_proto_goTypes = []any{
	(TargetEvents)(0),                // 0: daemon.TargetEvents
	(ImportStatus)(0),                // 1: daemon.ImportStatus
	(*MissedEvents)(nil),             // 2: daemon.MissedEvents
	(*ImportRequest)(nil),            // 3: daemon.ImportRequest
	(*ImportResponse)(nil),           // 4: daemon.ImportResponse
	(*ResynchronizeForRequest)(nil),  // 5: daemon.ResynchronizeForRequest
	(*ResynchronizeForResponse)(nil), // 6: daemon.ResynchronizeForResponse
	(*GetEventsRequest)(nil),         // 7: daemon.GetEventsRequest
	(*GetEventsResponse)(nil),        // 8: daemon.GetEventsResponse
	(Type)(0),                        // 9: daemon.Type
	(*EventsCollection)(nil),         // 10: daemon.EventsCollection
}
var file_a_study_in_history_proto_depIdxs = []int32{
	9,  // 0: daemon.MissedEvents.type:type_name -> daemon.Type
	1,  // 1: daemon.ImportResponse.status:type_name -> daemon.ImportStatus
	2,  // 2: daemon.ImportResponse.missedEvents:type_name -> daemon.MissedEvents
	0,  // 3: daemon.ResynchronizeForRequest.targetEvents:type_name -> daemon.TargetEvents
	0,  // 4: daemon.GetEventsRequest.targetEvents:type_name -> daemon.TargetEvents
	10, // 5: daemon.GetEventsResponse.eventsCollection:type_name -> daemon.EventsCollection
	9,  // 6: daemon.GetEventsResponse.type:type_name -> daemon.Type
	3,  // 7: daemon.AStudyInHistory.Import:input_type -> daemon.ImportRequest
	5,  // 8: daemon.AStudyInHistory.ResynchronizeFor:input_type -> daemon.ResynchronizeForRequest
	7,  // 9: daemon.AStudyInHistory.GetEventsFor:input_type -> daemon.GetEventsRequest
	4,  // 10: daemon.AStudyInHistory.Import:output_type -> daemon.ImportResponse
	6,  // 11: daemon.AStudyInHistory.ResynchronizeFor:output_type -> daemon.ResynchronizeForResponse
	8,  // 12: daemon.AStudyInHistory.GetEventsFor:output_type -> daemon.GetEventsResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_a_study_in_history_proto_init() }
func file_a_study_in_history_proto_init() {
	if File_a_study_in_history_proto != nil {
		return
	}
	file_events_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_a_study_in_history_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MissedEvents); i {
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
		file_a_study_in_history_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ImportRequest); i {
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
		file_a_study_in_history_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ImportResponse); i {
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
		file_a_study_in_history_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ResynchronizeForRequest); i {
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
		file_a_study_in_history_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ResynchronizeForResponse); i {
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
		file_a_study_in_history_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetEventsRequest); i {
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
		file_a_study_in_history_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetEventsResponse); i {
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
			RawDescriptor: file_a_study_in_history_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_a_study_in_history_proto_goTypes,
		DependencyIndexes: file_a_study_in_history_proto_depIdxs,
		EnumInfos:         file_a_study_in_history_proto_enumTypes,
		MessageInfos:      file_a_study_in_history_proto_msgTypes,
	}.Build()
	File_a_study_in_history_proto = out.File
	file_a_study_in_history_proto_rawDesc = nil
	file_a_study_in_history_proto_goTypes = nil
	file_a_study_in_history_proto_depIdxs = nil
}
