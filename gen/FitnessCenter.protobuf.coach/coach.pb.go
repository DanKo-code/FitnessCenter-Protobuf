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

type CoachObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Photo       string `protobuf:"bytes,4,opt,name=photo,proto3" json:"photo,omitempty"`
	CreatedTime string `protobuf:"bytes,5,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	UpdatedTime string `protobuf:"bytes,6,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
}

func (x *CoachObject) Reset() {
	*x = CoachObject{}
	mi := &file_coach_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CoachObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoachObject) ProtoMessage() {}

func (x *CoachObject) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CoachObject.ProtoReflect.Descriptor instead.
func (*CoachObject) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{0}
}

func (x *CoachObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CoachObject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CoachObject) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CoachObject) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *CoachObject) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *CoachObject) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

type CoachDataForCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description     string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CoachServiceIds []string `protobuf:"bytes,3,rep,name=coach_service_ids,json=coachServiceIds,proto3" json:"coach_service_ids,omitempty"`
}

func (x *CoachDataForCreate) Reset() {
	*x = CoachDataForCreate{}
	mi := &file_coach_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CoachDataForCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoachDataForCreate) ProtoMessage() {}

func (x *CoachDataForCreate) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CoachDataForCreate.ProtoReflect.Descriptor instead.
func (*CoachDataForCreate) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{1}
}

func (x *CoachDataForCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CoachDataForCreate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CoachDataForCreate) GetCoachServiceIds() []string {
	if x != nil {
		return x.CoachServiceIds
	}
	return nil
}

type CoachDataForUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description     string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CoachServiceIds []string `protobuf:"bytes,4,rep,name=coach_service_ids,json=coachServiceIds,proto3" json:"coach_service_ids,omitempty"`
}

func (x *CoachDataForUpdate) Reset() {
	*x = CoachDataForUpdate{}
	mi := &file_coach_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CoachDataForUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoachDataForUpdate) ProtoMessage() {}

func (x *CoachDataForUpdate) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CoachDataForUpdate.ProtoReflect.Descriptor instead.
func (*CoachDataForUpdate) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{2}
}

func (x *CoachDataForUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CoachDataForUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CoachDataForUpdate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CoachDataForUpdate) GetCoachServiceIds() []string {
	if x != nil {
		return x.CoachServiceIds
	}
	return nil
}

type CreateCoachRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*CreateCoachRequest_CoachDataForCreate
	//	*CreateCoachRequest_CoachPhoto
	Payload isCreateCoachRequest_Payload `protobuf_oneof:"payload"`
}

func (x *CreateCoachRequest) Reset() {
	*x = CreateCoachRequest{}
	mi := &file_coach_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCoachRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoachRequest) ProtoMessage() {}

func (x *CreateCoachRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateCoachRequest.ProtoReflect.Descriptor instead.
func (*CreateCoachRequest) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{3}
}

func (m *CreateCoachRequest) GetPayload() isCreateCoachRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *CreateCoachRequest) GetCoachDataForCreate() *CoachDataForCreate {
	if x, ok := x.GetPayload().(*CreateCoachRequest_CoachDataForCreate); ok {
		return x.CoachDataForCreate
	}
	return nil
}

func (x *CreateCoachRequest) GetCoachPhoto() []byte {
	if x, ok := x.GetPayload().(*CreateCoachRequest_CoachPhoto); ok {
		return x.CoachPhoto
	}
	return nil
}

type isCreateCoachRequest_Payload interface {
	isCreateCoachRequest_Payload()
}

type CreateCoachRequest_CoachDataForCreate struct {
	CoachDataForCreate *CoachDataForCreate `protobuf:"bytes,1,opt,name=coachDataForCreate,proto3,oneof"`
}

type CreateCoachRequest_CoachPhoto struct {
	CoachPhoto []byte `protobuf:"bytes,2,opt,name=coachPhoto,proto3,oneof"`
}

func (*CreateCoachRequest_CoachDataForCreate) isCreateCoachRequest_Payload() {}

func (*CreateCoachRequest_CoachPhoto) isCreateCoachRequest_Payload() {}

type CreateCoachResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoachObject *CoachObject `protobuf:"bytes,1,opt,name=coachObject,proto3" json:"coachObject,omitempty"`
}

func (x *CreateCoachResponse) Reset() {
	*x = CreateCoachResponse{}
	mi := &file_coach_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCoachResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoachResponse) ProtoMessage() {}

func (x *CreateCoachResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateCoachResponse.ProtoReflect.Descriptor instead.
func (*CreateCoachResponse) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCoachResponse) GetCoachObject() *CoachObject {
	if x != nil {
		return x.CoachObject
	}
	return nil
}

type GetCoachByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCoachByIdRequest) Reset() {
	*x = GetCoachByIdRequest{}
	mi := &file_coach_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCoachByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoachByIdRequest) ProtoMessage() {}

func (x *GetCoachByIdRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCoachByIdRequest.ProtoReflect.Descriptor instead.
func (*GetCoachByIdRequest) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{5}
}

func (x *GetCoachByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCoachByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoachObject *CoachObject `protobuf:"bytes,1,opt,name=coachObject,proto3" json:"coachObject,omitempty"`
}

func (x *GetCoachByIdResponse) Reset() {
	*x = GetCoachByIdResponse{}
	mi := &file_coach_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCoachByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoachByIdResponse) ProtoMessage() {}

func (x *GetCoachByIdResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCoachByIdResponse.ProtoReflect.Descriptor instead.
func (*GetCoachByIdResponse) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{6}
}

func (x *GetCoachByIdResponse) GetCoachObject() *CoachObject {
	if x != nil {
		return x.CoachObject
	}
	return nil
}

type UpdateCoachRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*UpdateCoachRequest_CoachDataForUpdate
	//	*UpdateCoachRequest_CoachPhoto
	Payload isUpdateCoachRequest_Payload `protobuf_oneof:"payload"`
}

func (x *UpdateCoachRequest) Reset() {
	*x = UpdateCoachRequest{}
	mi := &file_coach_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCoachRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCoachRequest) ProtoMessage() {}

func (x *UpdateCoachRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateCoachRequest.ProtoReflect.Descriptor instead.
func (*UpdateCoachRequest) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{7}
}

func (m *UpdateCoachRequest) GetPayload() isUpdateCoachRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *UpdateCoachRequest) GetCoachDataForUpdate() *CoachDataForUpdate {
	if x, ok := x.GetPayload().(*UpdateCoachRequest_CoachDataForUpdate); ok {
		return x.CoachDataForUpdate
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

type UpdateCoachRequest_CoachDataForUpdate struct {
	CoachDataForUpdate *CoachDataForUpdate `protobuf:"bytes,1,opt,name=coachDataForUpdate,proto3,oneof"`
}

type UpdateCoachRequest_CoachPhoto struct {
	CoachPhoto []byte `protobuf:"bytes,2,opt,name=coachPhoto,proto3,oneof"`
}

func (*UpdateCoachRequest_CoachDataForUpdate) isUpdateCoachRequest_Payload() {}

func (*UpdateCoachRequest_CoachPhoto) isUpdateCoachRequest_Payload() {}

type UpdateCoachResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoachObject *CoachObject `protobuf:"bytes,1,opt,name=coachObject,proto3" json:"coachObject,omitempty"`
}

func (x *UpdateCoachResponse) Reset() {
	*x = UpdateCoachResponse{}
	mi := &file_coach_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCoachResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCoachResponse) ProtoMessage() {}

func (x *UpdateCoachResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coach_proto_msgTypes[8]
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
	return file_coach_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateCoachResponse) GetCoachObject() *CoachObject {
	if x != nil {
		return x.CoachObject
	}
	return nil
}

type DeleteCoachByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCoachByIdRequest) Reset() {
	*x = DeleteCoachByIdRequest{}
	mi := &file_coach_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCoachByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCoachByIdRequest) ProtoMessage() {}

func (x *DeleteCoachByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coach_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCoachByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteCoachByIdRequest) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteCoachByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteCoachByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoachObject *CoachObject `protobuf:"bytes,1,opt,name=coachObject,proto3" json:"coachObject,omitempty"`
}

func (x *DeleteCoachByIdResponse) Reset() {
	*x = DeleteCoachByIdResponse{}
	mi := &file_coach_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCoachByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCoachByIdResponse) ProtoMessage() {}

func (x *DeleteCoachByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coach_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCoachByIdResponse.ProtoReflect.Descriptor instead.
func (*DeleteCoachByIdResponse) Descriptor() ([]byte, []int) {
	return file_coach_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteCoachByIdResponse) GetCoachObject() *CoachObject {
	if x != nil {
		return x.CoachObject
	}
	return nil
}

type GetCoachesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoachObjects []*CoachObject `protobuf:"bytes,1,rep,name=coachObjects,proto3" json:"coachObjects,omitempty"`
}

func (x *GetCoachesResponse) Reset() {
	*x = GetCoachesResponse{}
	mi := &file_coach_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCoachesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoachesResponse) ProtoMessage() {}

func (x *GetCoachesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coach_proto_msgTypes[11]
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
	return file_coach_proto_rawDescGZIP(), []int{11}
}

func (x *GetCoachesResponse) GetCoachObjects() []*CoachObject {
	if x != nil {
		return x.CoachObjects
	}
	return nil
}

var File_coach_proto protoreflect.FileDescriptor

var file_coach_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x66,
	0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f,
	0x61, 0x63, 0x68, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xaf, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x76, 0x0a, 0x12, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46,
	0x6f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a,
	0x0a, 0x11, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x61, 0x63, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x12, 0x43,
	0x6f, 0x61, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x61, 0x63, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x61, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5a, 0x0a, 0x12, 0x63, 0x6f,
	0x61, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x43, 0x6f,
	0x61, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x48, 0x00, 0x52, 0x12, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0a, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6f,
	0x61, 0x63, 0x68, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x5a, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x63, 0x6f,
	0x61, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22,
	0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x61,
	0x63, 0x68, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43,
	0x0a, 0x0b, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x61, 0x63, 0x68,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x22, 0x9d, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x61, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5a, 0x0a, 0x12, 0x63, 0x6f,
	0x61, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x43, 0x6f,
	0x61, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x00, 0x52, 0x12, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0a, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6f,
	0x61, 0x63, 0x68, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x5a, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x63, 0x6f,
	0x61, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22,
	0x28, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5e, 0x0a, 0x17, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x69, 0x74, 0x6e,
	0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68,
	0x2e, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0b, 0x63, 0x6f,
	0x61, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x5b, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x61, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x0c, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x61,
	0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0c, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x32, 0xfa, 0x03, 0x0a, 0x05, 0x43, 0x6f, 0x61, 0x63, 0x68,
	0x12, 0x64, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x12,
	0x28, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x66, 0x69, 0x74, 0x6e,
	0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x65, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x61,
	0x63, 0x68, 0x42, 0x79, 0x49, 0x64, 0x12, 0x29, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x61, 0x63,
	0x68, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x12, 0x28, 0x2e, 0x66,
	0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f,
	0x61, 0x63, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x12, 0x6e, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x61,
	0x63, 0x68, 0x42, 0x79, 0x49, 0x64, 0x12, 0x2c, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x65,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e, 0x66, 0x69, 0x74, 0x6e,
	0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x61, 0x63, 0x68,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x46, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f,
	0x61, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_coach_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_coach_proto_goTypes = []any{
	(*CoachObject)(nil),             // 0: fitness_center.coach.CoachObject
	(*CoachDataForCreate)(nil),      // 1: fitness_center.coach.CoachDataForCreate
	(*CoachDataForUpdate)(nil),      // 2: fitness_center.coach.CoachDataForUpdate
	(*CreateCoachRequest)(nil),      // 3: fitness_center.coach.CreateCoachRequest
	(*CreateCoachResponse)(nil),     // 4: fitness_center.coach.CreateCoachResponse
	(*GetCoachByIdRequest)(nil),     // 5: fitness_center.coach.GetCoachByIdRequest
	(*GetCoachByIdResponse)(nil),    // 6: fitness_center.coach.GetCoachByIdResponse
	(*UpdateCoachRequest)(nil),      // 7: fitness_center.coach.UpdateCoachRequest
	(*UpdateCoachResponse)(nil),     // 8: fitness_center.coach.UpdateCoachResponse
	(*DeleteCoachByIdRequest)(nil),  // 9: fitness_center.coach.DeleteCoachByIdRequest
	(*DeleteCoachByIdResponse)(nil), // 10: fitness_center.coach.DeleteCoachByIdResponse
	(*GetCoachesResponse)(nil),      // 11: fitness_center.coach.GetCoachesResponse
	(*emptypb.Empty)(nil),           // 12: google.protobuf.Empty
}
var file_coach_proto_depIdxs = []int32{
	1,  // 0: fitness_center.coach.CreateCoachRequest.coachDataForCreate:type_name -> fitness_center.coach.CoachDataForCreate
	0,  // 1: fitness_center.coach.CreateCoachResponse.coachObject:type_name -> fitness_center.coach.CoachObject
	0,  // 2: fitness_center.coach.GetCoachByIdResponse.coachObject:type_name -> fitness_center.coach.CoachObject
	2,  // 3: fitness_center.coach.UpdateCoachRequest.coachDataForUpdate:type_name -> fitness_center.coach.CoachDataForUpdate
	0,  // 4: fitness_center.coach.UpdateCoachResponse.coachObject:type_name -> fitness_center.coach.CoachObject
	0,  // 5: fitness_center.coach.DeleteCoachByIdResponse.coachObject:type_name -> fitness_center.coach.CoachObject
	0,  // 6: fitness_center.coach.GetCoachesResponse.coachObjects:type_name -> fitness_center.coach.CoachObject
	3,  // 7: fitness_center.coach.Coach.CreateCoach:input_type -> fitness_center.coach.CreateCoachRequest
	5,  // 8: fitness_center.coach.Coach.GetCoachById:input_type -> fitness_center.coach.GetCoachByIdRequest
	7,  // 9: fitness_center.coach.Coach.UpdateCoach:input_type -> fitness_center.coach.UpdateCoachRequest
	9,  // 10: fitness_center.coach.Coach.DeleteCoachById:input_type -> fitness_center.coach.DeleteCoachByIdRequest
	12, // 11: fitness_center.coach.Coach.GetCoaches:input_type -> google.protobuf.Empty
	4,  // 12: fitness_center.coach.Coach.CreateCoach:output_type -> fitness_center.coach.CreateCoachResponse
	6,  // 13: fitness_center.coach.Coach.GetCoachById:output_type -> fitness_center.coach.GetCoachByIdResponse
	8,  // 14: fitness_center.coach.Coach.UpdateCoach:output_type -> fitness_center.coach.UpdateCoachResponse
	10, // 15: fitness_center.coach.Coach.DeleteCoachById:output_type -> fitness_center.coach.DeleteCoachByIdResponse
	11, // 16: fitness_center.coach.Coach.GetCoaches:output_type -> fitness_center.coach.GetCoachesResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_coach_proto_init() }
func file_coach_proto_init() {
	if File_coach_proto != nil {
		return
	}
	file_coach_proto_msgTypes[3].OneofWrappers = []any{
		(*CreateCoachRequest_CoachDataForCreate)(nil),
		(*CreateCoachRequest_CoachPhoto)(nil),
	}
	file_coach_proto_msgTypes[7].OneofWrappers = []any{
		(*UpdateCoachRequest_CoachDataForUpdate)(nil),
		(*UpdateCoachRequest_CoachPhoto)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coach_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
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
