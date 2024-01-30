// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: audit/v1/audit.proto

package auditv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "go.datalift.io/admiral/common/api/api/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_audit_v1_audit_proto_rawDescGZIP(), []int{0}
}

func (x *TimeRange) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *TimeRange) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type GetEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Window:
	//
	//	*GetEventsRequest_Range
	//	*GetEventsRequest_Since
	Window isGetEventsRequest_Window `protobuf_oneof:"window"`
	// Currently page_token specifies the page number you wish to request.
	// The rationale behind the naming is we might changes this to a cursor implementation
	// in the future and did not want to break existing implementations of the API.
	// https://cloud.google.com/apis/design/design_patterns#list_pagination
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Limit     uint64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_proto_msgTypes[1]
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
	return file_audit_v1_audit_proto_rawDescGZIP(), []int{1}
}

func (m *GetEventsRequest) GetWindow() isGetEventsRequest_Window {
	if m != nil {
		return m.Window
	}
	return nil
}

func (x *GetEventsRequest) GetRange() *TimeRange {
	if x, ok := x.GetWindow().(*GetEventsRequest_Range); ok {
		return x.Range
	}
	return nil
}

func (x *GetEventsRequest) GetSince() *durationpb.Duration {
	if x, ok := x.GetWindow().(*GetEventsRequest_Since); ok {
		return x.Since
	}
	return nil
}

func (x *GetEventsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *GetEventsRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type isGetEventsRequest_Window interface {
	isGetEventsRequest_Window()
}

type GetEventsRequest_Range struct {
	Range *TimeRange `protobuf:"bytes,1,opt,name=range,proto3,oneof"`
}

type GetEventsRequest_Since struct {
	Since *durationpb.Duration `protobuf:"bytes,2,opt,name=since,proto3,oneof"`
}

func (*GetEventsRequest_Range) isGetEventsRequest_Window() {}

func (*GetEventsRequest_Since) isGetEventsRequest_Window() {}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_audit_v1_audit_proto_rawDescGZIP(), []int{2}
}

func (x *Resource) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *Resource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RequestMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body *anypb.Any `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *RequestMetadata) Reset() {
	*x = RequestMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMetadata) ProtoMessage() {}

func (x *RequestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMetadata.ProtoReflect.Descriptor instead.
func (*RequestMetadata) Descriptor() ([]byte, []int) {
	return file_audit_v1_audit_proto_rawDescGZIP(), []int{3}
}

func (x *RequestMetadata) GetBody() *anypb.Any {
	if x != nil {
		return x.Body
	}
	return nil
}

type ResponseMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body *anypb.Any `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ResponseMetadata) Reset() {
	*x = ResponseMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMetadata) ProtoMessage() {}

func (x *ResponseMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMetadata.ProtoReflect.Descriptor instead.
func (*ResponseMetadata) Descriptor() ([]byte, []int) {
	return file_audit_v1_audit_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseMetadata) GetBody() *anypb.Any {
	if x != nil {
		return x.Body
	}
	return nil
}

type RequestEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// What attempted the action.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// The service performing the operation.
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The method being called.
	MethodName string `protobuf:"bytes,3,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	// The type of operation being performed.
	Type v1.ActionType `protobuf:"varint,4,opt,name=type,proto3,enum=admiral.api.v1.ActionType" json:"type,omitempty"`
	// The status of the overall operation.
	Status *status.Status `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	// The resources touched during the event.
	Resources []*Resource `protobuf:"bytes,6,rep,name=resources,proto3" json:"resources,omitempty"`
	// The API request saved as metadata for the event.
	RequestMetadata *RequestMetadata `protobuf:"bytes,7,opt,name=request_metadata,json=requestMetadata,proto3" json:"request_metadata,omitempty"`
	// The API response saved as metadata for the event.
	ResponseMetadata *ResponseMetadata `protobuf:"bytes,8,opt,name=response_metadata,json=responseMetadata,proto3" json:"response_metadata,omitempty"`
}

func (x *RequestEvent) Reset() {
	*x = RequestEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestEvent) ProtoMessage() {}

func (x *RequestEvent) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestEvent.ProtoReflect.Descriptor instead.
func (*RequestEvent) Descriptor() ([]byte, []int) {
	return file_audit_v1_audit_proto_rawDescGZIP(), []int{5}
}

func (x *RequestEvent) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RequestEvent) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *RequestEvent) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *RequestEvent) GetType() v1.ActionType {
	if x != nil {
		return x.Type
	}
	return v1.ActionType(0)
}

func (x *RequestEvent) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *RequestEvent) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *RequestEvent) GetRequestMetadata() *RequestMetadata {
	if x != nil {
		return x.RequestMetadata
	}
	return nil
}

func (x *RequestEvent) GetResponseMetadata() *ResponseMetadata {
	if x != nil {
		return x.ResponseMetadata
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// When the event happened.
	OccurredAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	// Types that are assignable to EventType:
	//
	//	*Event_Event
	EventType isEvent_EventType `protobuf_oneof:"event_type"`
	// The event id.
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_audit_v1_audit_proto_rawDescGZIP(), []int{6}
}

func (x *Event) GetOccurredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.OccurredAt
	}
	return nil
}

func (m *Event) GetEventType() isEvent_EventType {
	if m != nil {
		return m.EventType
	}
	return nil
}

func (x *Event) GetEvent() *RequestEvent {
	if x, ok := x.GetEventType().(*Event_Event); ok {
		return x.Event
	}
	return nil
}

func (x *Event) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type isEvent_EventType interface {
	isEvent_EventType()
}

type Event_Event struct {
	Event *RequestEvent `protobuf:"bytes,2,opt,name=event,proto3,oneof"`
}

func (*Event_Event) isEvent_EventType() {}

type GetEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events        []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	NextPageToken string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *GetEventsResponse) Reset() {
	*x = GetEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsResponse) ProtoMessage() {}

func (x *GetEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_proto_msgTypes[7]
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
	return file_audit_v1_audit_proto_rawDescGZIP(), []int{7}
}

func (x *GetEventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *GetEventsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId int64 `protobuf:"varint,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *GetEventRequest) Reset() {
	*x = GetEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventRequest) ProtoMessage() {}

func (x *GetEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventRequest.ProtoReflect.Descriptor instead.
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return file_audit_v1_audit_proto_rawDescGZIP(), []int{8}
}

func (x *GetEventRequest) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

type GetEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *GetEventResponse) Reset() {
	*x = GetEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventResponse) ProtoMessage() {}

func (x *GetEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventResponse.ProtoReflect.Descriptor instead.
func (*GetEventResponse) Descriptor() ([]byte, []int) {
	return file_audit_v1_audit_proto_rawDescGZIP(), []int{9}
}

func (x *GetEventResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_audit_v1_audit_proto protoreflect.FileDescriptor

var file_audit_v1_audit_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x43, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xb2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb9, 0x01,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00,
	0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42,
	0x08, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x22, 0x64, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x3a, 0x2d, 0xb2, 0xe1, 0x1c, 0x29, 0x0a, 0x27, 0x0a, 0x19, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x0a, 0x7b, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x7d, 0x22,
	0x3b, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x3c, 0x0a, 0x10,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x28, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0xd9, 0x03, 0x0a, 0x0c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2a, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0b,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x4f, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x0f, 0xaa, 0xe1, 0x1c, 0x0b, 0x0a, 0x09, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x3b, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x36, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x6c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72,
	0x61, 0x6c, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x41, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x32, 0xfe, 0x01, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69, 0x74, 0x41, 0x50, 0x49, 0x12,
	0x7a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0xaa, 0xe1, 0x1c, 0x02, 0x08, 0x02, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2f, 0x67, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x76, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23,
	0xaa, 0xe1, 0x1c, 0x02, 0x08, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22,
	0x12, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x42, 0xb8, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x72, 0x61, 0x6c, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x6f, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x72,
	0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x41, 0x41, 0x58, 0xaa, 0x02, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x72, 0x61, 0x6c, 0x2e, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x72, 0x61,
	0x6c, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x41, 0x64, 0x6d,
	0x69, 0x72, 0x61, 0x6c, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x41, 0x64, 0x6d, 0x69,
	0x72, 0x61, 0x6c, 0x3a, 0x3a, 0x41, 0x75, 0x64, 0x69, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_audit_v1_audit_proto_rawDescOnce sync.Once
	file_audit_v1_audit_proto_rawDescData = file_audit_v1_audit_proto_rawDesc
)

func file_audit_v1_audit_proto_rawDescGZIP() []byte {
	file_audit_v1_audit_proto_rawDescOnce.Do(func() {
		file_audit_v1_audit_proto_rawDescData = protoimpl.X.CompressGZIP(file_audit_v1_audit_proto_rawDescData)
	})
	return file_audit_v1_audit_proto_rawDescData
}

var file_audit_v1_audit_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_audit_v1_audit_proto_goTypes = []interface{}{
	(*TimeRange)(nil),             // 0: admiral.audit.v1.TimeRange
	(*GetEventsRequest)(nil),      // 1: admiral.audit.v1.GetEventsRequest
	(*Resource)(nil),              // 2: admiral.audit.v1.Resource
	(*RequestMetadata)(nil),       // 3: admiral.audit.v1.RequestMetadata
	(*ResponseMetadata)(nil),      // 4: admiral.audit.v1.ResponseMetadata
	(*RequestEvent)(nil),          // 5: admiral.audit.v1.RequestEvent
	(*Event)(nil),                 // 6: admiral.audit.v1.Event
	(*GetEventsResponse)(nil),     // 7: admiral.audit.v1.GetEventsResponse
	(*GetEventRequest)(nil),       // 8: admiral.audit.v1.GetEventRequest
	(*GetEventResponse)(nil),      // 9: admiral.audit.v1.GetEventResponse
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 11: google.protobuf.Duration
	(*anypb.Any)(nil),             // 12: google.protobuf.Any
	(v1.ActionType)(0),            // 13: admiral.api.v1.ActionType
	(*status.Status)(nil),         // 14: google.rpc.Status
}
var file_audit_v1_audit_proto_depIdxs = []int32{
	10, // 0: admiral.audit.v1.TimeRange.start_time:type_name -> google.protobuf.Timestamp
	10, // 1: admiral.audit.v1.TimeRange.end_time:type_name -> google.protobuf.Timestamp
	0,  // 2: admiral.audit.v1.GetEventsRequest.range:type_name -> admiral.audit.v1.TimeRange
	11, // 3: admiral.audit.v1.GetEventsRequest.since:type_name -> google.protobuf.Duration
	12, // 4: admiral.audit.v1.RequestMetadata.body:type_name -> google.protobuf.Any
	12, // 5: admiral.audit.v1.ResponseMetadata.body:type_name -> google.protobuf.Any
	13, // 6: admiral.audit.v1.RequestEvent.type:type_name -> admiral.api.v1.ActionType
	14, // 7: admiral.audit.v1.RequestEvent.status:type_name -> google.rpc.Status
	2,  // 8: admiral.audit.v1.RequestEvent.resources:type_name -> admiral.audit.v1.Resource
	3,  // 9: admiral.audit.v1.RequestEvent.request_metadata:type_name -> admiral.audit.v1.RequestMetadata
	4,  // 10: admiral.audit.v1.RequestEvent.response_metadata:type_name -> admiral.audit.v1.ResponseMetadata
	10, // 11: admiral.audit.v1.Event.occurred_at:type_name -> google.protobuf.Timestamp
	5,  // 12: admiral.audit.v1.Event.event:type_name -> admiral.audit.v1.RequestEvent
	6,  // 13: admiral.audit.v1.GetEventsResponse.events:type_name -> admiral.audit.v1.Event
	6,  // 14: admiral.audit.v1.GetEventResponse.event:type_name -> admiral.audit.v1.Event
	1,  // 15: admiral.audit.v1.AuditAPI.GetEvents:input_type -> admiral.audit.v1.GetEventsRequest
	8,  // 16: admiral.audit.v1.AuditAPI.GetEvent:input_type -> admiral.audit.v1.GetEventRequest
	7,  // 17: admiral.audit.v1.AuditAPI.GetEvents:output_type -> admiral.audit.v1.GetEventsResponse
	9,  // 18: admiral.audit.v1.AuditAPI.GetEvent:output_type -> admiral.audit.v1.GetEventResponse
	17, // [17:19] is the sub-list for method output_type
	15, // [15:17] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_audit_v1_audit_proto_init() }
func file_audit_v1_audit_proto_init() {
	if File_audit_v1_audit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_audit_v1_audit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRange); i {
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
		file_audit_v1_audit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_audit_v1_audit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_audit_v1_audit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMetadata); i {
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
		file_audit_v1_audit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMetadata); i {
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
		file_audit_v1_audit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestEvent); i {
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
		file_audit_v1_audit_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_audit_v1_audit_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_audit_v1_audit_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventRequest); i {
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
		file_audit_v1_audit_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventResponse); i {
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
	file_audit_v1_audit_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*GetEventsRequest_Range)(nil),
		(*GetEventsRequest_Since)(nil),
	}
	file_audit_v1_audit_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Event_Event)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_audit_v1_audit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_audit_v1_audit_proto_goTypes,
		DependencyIndexes: file_audit_v1_audit_proto_depIdxs,
		MessageInfos:      file_audit_v1_audit_proto_msgTypes,
	}.Build()
	File_audit_v1_audit_proto = out.File
	file_audit_v1_audit_proto_rawDesc = nil
	file_audit_v1_audit_proto_goTypes = nil
	file_audit_v1_audit_proto_depIdxs = nil
}
