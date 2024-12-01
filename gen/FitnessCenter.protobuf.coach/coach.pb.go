// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: coach.proto

package FitnessCenter_protobuf_coach

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CoachData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CoachData) Reset() {
	*x = CoachData{}
	mi := &file_coach_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CoachData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoachData) ProtoMessage() {}

func (x *CoachData) ProtoReflect() protoreflect.Message {
	mi := &file_coach_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoachData.ProtoReflect.Descriptor instead.
func (*CoachData) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{0}
}

func (x *CoachData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CoachData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CoachData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateCoachRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateCoachRequest) Reset() {
	*x = CreateCoachRequest{}
	mi := &file_coach_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCoachRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoachRequest) ProtoMessage() {}

func (x *CreateCoachRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coach_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoachRequest.ProtoReflect.Descriptor instead.
func (*CreateCoachRequest) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCoachRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCoachRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateCoachResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoachData *CoachData `protobuf:"bytes,1,opt,name=coachData,proto3" json:"coachData,omitempty"`
}

func (x *CreateCoachResponse) Reset() {
	*x = CreateCoachResponse{}
	mi := &file_coach_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCoachResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoachResponse) ProtoMessage() {}

func (x *CreateCoachResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coach_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoachResponse.ProtoReflect.Descriptor instead.
func (*CreateCoachResponse) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCoachResponse) GetCoachData() *CoachData {
	if x != nil {
		return x.CoachData
	}
	return nil
}

type GetCoachesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coaches []*CoachData `protobuf:"bytes,1,rep,name=coaches,proto3" json:"coaches,omitempty"`
}

func (x *GetCoachesResponse) Reset() {
	*x = GetCoachesResponse{}
	mi := &file_coach_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCoachesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoachesResponse) ProtoMessage() {}

func (x *GetCoachesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coach_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoachesResponse.ProtoReflect.Descriptor instead.
func (*GetCoachesResponse) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{3}
}

func (x *GetCoachesResponse) GetCoaches() []*CoachData {
	if x != nil {
		return x.Coaches
	}
	return nil
}

type DeleteCoachRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCoachRequest) Reset() {
	*x = DeleteCoachRequest{}
	mi := &file_coach_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCoachRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCoachRequest) ProtoMessage() {}

func (x *DeleteCoachRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coach_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCoachRequest.ProtoReflect.Descriptor instead.
func (*DeleteCoachRequest) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCoachRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteCoachResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coach *CoachData `protobuf:"bytes,1,opt,name=coach,proto3" json:"coach,omitempty"`
}

func (x *DeleteCoachResponse) Reset() {
	*x = DeleteCoachResponse{}
	mi := &file_coach_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCoachResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCoachResponse) ProtoMessage() {}

func (x *DeleteCoachResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coach_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCoachResponse.ProtoReflect.Descriptor instead.
func (*DeleteCoachResponse) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCoachResponse) GetCoach() *CoachData {
	if x != nil {
		return x.Coach
	}
	return nil
}

type UpdateCoachRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*UpdateCoachRequest_CoachData
	//	*UpdateCoachRequest_CoachPhoto
	Payload isUpdateCoachRequest_Payload `protobuf_oneof:"payload"`
}

func (x *UpdateCoachRequest) Reset() {
	*x = UpdateCoachRequest{}
	mi := &file_coach_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCoachRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCoachRequest) ProtoMessage() {}

func (x *UpdateCoachRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coach_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCoachRequest.ProtoReflect.Descriptor instead.
func (*UpdateCoachRequest) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{6}
}

func (m *UpdateCoachRequest) GetPayload() isUpdateCoachRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *UpdateCoachRequest) GetCoachData() *CoachData {
	if x, ok := x.GetPayload().(*UpdateCoachRequest_CoachData); ok {
		return x.CoachData
	}
	return nil
}

func (x *UpdateCoachRequest) GetCoachPhoto() []byte {
	if x, ok := x.GetPayload().(*UpdateCoachRequest_CoachPhoto); ok {
		return x.CoachPhoto
	}
	return nil
}

type isUpdateCoachRequest_Payload interface {
	isUpdateCoachRequest_Payload()
}

type UpdateCoachRequest_CoachData struct {
	CoachData *CoachData `protobuf:"bytes,1,opt,name=coachData,proto3,oneof"`
}

type UpdateCoachRequest_CoachPhoto struct {
	CoachPhoto []byte `protobuf:"bytes,2,opt,name=coachPhoto,proto3,oneof"`
}

func (*UpdateCoachRequest_CoachData) isUpdateCoachRequest_Payload() {}

func (*UpdateCoachRequest_CoachPhoto) isUpdateCoachRequest_Payload() {}

type UpdateCoachResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoachData *CoachData `protobuf:"bytes,1,opt,name=coachData,proto3" json:"coachData,omitempty"`
}

func (x *UpdateCoachResponse) Reset() {
	*x = UpdateCoachResponse{}
	mi := &file_coach_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCoachResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCoachResponse) ProtoMessage() {}

func (x *UpdateCoachResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coach_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCoachResponse.ProtoReflect.Descriptor instead.
func (*UpdateCoachResponse) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCoachResponse) GetCoachData() *CoachData {
	if x != nil {
		return x.CoachData
	}
	return nil
}

var File_coach_proto protoreflect.FileDescriptor

var file_coach_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x66,
	0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f,
	0x61, 0x63, 0x68, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x51, 0x0a, 0x09, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x4a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x54, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x69, 0x74, 0x6e,
	0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68,
	0x2e, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x63, 0x6f, 0x61, 0x63,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x22, 0x4f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x61, 0x63,
	0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x63,
	0x6f, 0x61, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66,
	0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f,
	0x61, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x63,
	0x6f, 0x61, 0x63, 0x68, 0x65, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x61, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x05, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x22, 0x82, 0x01, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3f, 0x0a, 0x09, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x61, 0x63,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x20, 0x0a, 0x0a, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x50, 0x68, 0x6f, 0x74, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x54, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x69, 0x74, 0x6e,
	0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68,
	0x2e, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x63, 0x6f, 0x61, 0x63,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x32, 0x89, 0x03, 0x0a, 0x05, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x12,
	0x62, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x12, 0x28,
	0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65,
	0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x65,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e, 0x66, 0x69, 0x74, 0x6e,
	0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x61,
	0x63, 0x68, 0x42, 0x79, 0x49, 0x64, 0x12, 0x28, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x61, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x12, 0x28, 0x2e, 0x66, 0x69, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63,
	0x68, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28,
	0x01, 0x42, 0x1e, 0x5a, 0x1c, 0x46, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f, 0x61, 0x63,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coach_proto_rawDescOnce sync.Once
	file_coach_proto_rawDescData = file_coach_proto_rawDesc
)

func file_coach_proto_rawDescGZIP() []byte {
	file_coach_proto_rawDescOnce.Do(func() {
		file_coach_proto_rawDescData = protoimpl.X.CompressGZIP(file_coach_proto_rawDescData)
	})
	return file_coach_proto_rawDescData
}

var file_coach_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_coach_proto_goTypes = []any{
	(*CoachData)(nil),           // 0: fitness_center.coach.CoachData
	(*CreateCoachRequest)(nil),  // 1: fitness_center.coach.CreateCoachRequest
	(*CreateCoachResponse)(nil), // 2: fitness_center.coach.CreateCoachResponse
	(*GetCoachesResponse)(nil),  // 3: fitness_center.coach.GetCoachesResponse
	(*DeleteCoachRequest)(nil),  // 4: fitness_center.coach.DeleteCoachRequest
	(*DeleteCoachResponse)(nil), // 5: fitness_center.coach.DeleteCoachResponse
	(*UpdateCoachRequest)(nil),  // 6: fitness_center.coach.UpdateCoachRequest
	(*UpdateCoachResponse)(nil), // 7: fitness_center.coach.UpdateCoachResponse
	(*emptypb.Empty)(nil),       // 8: google.protobuf.Empty
}
var file_coach_proto_depIdxs = []int32{
	0, // 0: fitness_center.coach.CreateCoachResponse.coachData:type_name -> fitness_center.coach.CoachData
	0, // 1: fitness_center.coach.GetCoachesResponse.coaches:type_name -> fitness_center.coach.CoachData
	0, // 2: fitness_center.coach.DeleteCoachResponse.coach:type_name -> fitness_center.coach.CoachData
	0, // 3: fitness_center.coach.UpdateCoachRequest.coachData:type_name -> fitness_center.coach.CoachData
	0, // 4: fitness_center.coach.UpdateCoachResponse.coachData:type_name -> fitness_center.coach.CoachData
	1, // 5: fitness_center.coach.Coach.CreateCoach:input_type -> fitness_center.coach.CreateCoachRequest
	8, // 6: fitness_center.coach.Coach.GetCoaches:input_type -> google.protobuf.Empty
	4, // 7: fitness_center.coach.Coach.DeleteCoachById:input_type -> fitness_center.coach.DeleteCoachRequest
	6, // 8: fitness_center.coach.Coach.UpdateCoach:input_type -> fitness_center.coach.UpdateCoachRequest
	2, // 9: fitness_center.coach.Coach.CreateCoach:output_type -> fitness_center.coach.CreateCoachResponse
	3, // 10: fitness_center.coach.Coach.GetCoaches:output_type -> fitness_center.coach.GetCoachesResponse
	5, // 11: fitness_center.coach.Coach.DeleteCoachById:output_type -> fitness_center.coach.DeleteCoachResponse
	7, // 12: fitness_center.coach.Coach.UpdateCoach:output_type -> fitness_center.coach.UpdateCoachResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_coach_proto_init() }
func file_coach_proto_init() {
	if File_coach_proto != nil {
		return
	}
	file_coach_proto_msgTypes[6].OneofWrappers = []any{
		(*UpdateCoachRequest_CoachData)(nil),
		(*UpdateCoachRequest_CoachPhoto)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coach_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coach_proto_goTypes,
		DependencyIndexes: file_coach_proto_depIdxs,
		MessageInfos:      file_coach_proto_msgTypes,
	}.Build()
	File_coach_proto = out.File
	file_coach_proto_rawDesc = nil
	file_coach_proto_goTypes = nil
	file_coach_proto_depIdxs = nil
}
