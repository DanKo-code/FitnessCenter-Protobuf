// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: order.proto

package FitnessCenter_protobuf_order

import (
	FitnessCenter_protobuf_abonement "github.com/DanKo-code/FitnessCenter-Protobuf/gen/FitnessCenter.protobuf.abonement"
	FitnessCenter_protobuf_service "github.com/DanKo-code/FitnessCenter-Protobuf/gen/FitnessCenter.protobuf.service"
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

type OrderObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AbonementId string `protobuf:"bytes,3,opt,name=abonement_id,json=abonementId,proto3" json:"abonement_id,omitempty"`
	Status      string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CreatedTime string `protobuf:"bytes,5,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	UpdatedTime string `protobuf:"bytes,6,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
}

func (x *OrderObject) Reset() {
	*x = OrderObject{}
	mi := &file_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderObject) ProtoMessage() {}

func (x *OrderObject) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderObject.ProtoReflect.Descriptor instead.
func (*OrderObject) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderObject) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderObject) GetAbonementId() string {
	if x != nil {
		return x.AbonementId
	}
	return ""
}

func (x *OrderObject) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderObject) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *OrderObject) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

type OrderObjectWithAbonementWithServices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderObject     *OrderObject                                      `protobuf:"bytes,1,opt,name=orderObject,proto3" json:"orderObject,omitempty"`
	AbonementObject *FitnessCenter_protobuf_abonement.AbonementObject `protobuf:"bytes,2,opt,name=abonementObject,proto3" json:"abonementObject,omitempty"`
	ServiceObjects  []*FitnessCenter_protobuf_service.ServiceObject   `protobuf:"bytes,3,rep,name=serviceObjects,proto3" json:"serviceObjects,omitempty"`
}

func (x *OrderObjectWithAbonementWithServices) Reset() {
	*x = OrderObjectWithAbonementWithServices{}
	mi := &file_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderObjectWithAbonementWithServices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderObjectWithAbonementWithServices) ProtoMessage() {}

func (x *OrderObjectWithAbonementWithServices) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderObjectWithAbonementWithServices.ProtoReflect.Descriptor instead.
func (*OrderObjectWithAbonementWithServices) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderObjectWithAbonementWithServices) GetOrderObject() *OrderObject {
	if x != nil {
		return x.OrderObject
	}
	return nil
}

func (x *OrderObjectWithAbonementWithServices) GetAbonementObject() *FitnessCenter_protobuf_abonement.AbonementObject {
	if x != nil {
		return x.AbonementObject
	}
	return nil
}

func (x *OrderObjectWithAbonementWithServices) GetServiceObjects() []*FitnessCenter_protobuf_service.ServiceObject {
	if x != nil {
		return x.ServiceObjects
	}
	return nil
}

type OrderDataForCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AbonementId string `protobuf:"bytes,2,opt,name=abonement_id,json=abonementId,proto3" json:"abonement_id,omitempty"`
}

func (x *OrderDataForCreate) Reset() {
	*x = OrderDataForCreate{}
	mi := &file_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderDataForCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDataForCreate) ProtoMessage() {}

func (x *OrderDataForCreate) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDataForCreate.ProtoReflect.Descriptor instead.
func (*OrderDataForCreate) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderDataForCreate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderDataForCreate) GetAbonementId() string {
	if x != nil {
		return x.AbonementId
	}
	return ""
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderDataForCreate *OrderDataForCreate `protobuf:"bytes,1,opt,name=orderDataForCreate,proto3" json:"orderDataForCreate,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrderRequest) GetOrderDataForCreate() *OrderDataForCreate {
	if x != nil {
		return x.OrderDataForCreate
	}
	return nil
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderObject *OrderObject `protobuf:"bytes,1,opt,name=orderObject,proto3" json:"orderObject,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *CreateOrderResponse) GetOrderObject() *OrderObject {
	if x != nil {
		return x.OrderObject
	}
	return nil
}

type GetUserOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserOrdersRequest) Reset() {
	*x = GetUserOrdersRequest{}
	mi := &file_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserOrdersRequest) ProtoMessage() {}

func (x *GetUserOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserOrdersRequest.ProtoReflect.Descriptor instead.
func (*GetUserOrdersRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserOrdersRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderObjectWithAbonementWithServices []*OrderObjectWithAbonementWithServices `protobuf:"bytes,1,rep,name=orderObjectWithAbonementWithServices,proto3" json:"orderObjectWithAbonementWithServices,omitempty"`
}

func (x *GetUserOrdersResponse) Reset() {
	*x = GetUserOrdersResponse{}
	mi := &file_order_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserOrdersResponse) ProtoMessage() {}

func (x *GetUserOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserOrdersResponse.ProtoReflect.Descriptor instead.
func (*GetUserOrdersResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserOrdersResponse) GetOrderObjectWithAbonementWithServices() []*OrderObjectWithAbonementWithServices {
	if x != nil {
		return x.OrderObjectWithAbonementWithServices
	}
	return nil
}

type CreateCheckoutSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId      string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	AbonementId   string `protobuf:"bytes,2,opt,name=abonement_id,json=abonementId,proto3" json:"abonement_id,omitempty"`
	StripePriceId string `protobuf:"bytes,3,opt,name=stripe_price_id,json=stripePriceId,proto3" json:"stripe_price_id,omitempty"`
}

func (x *CreateCheckoutSessionRequest) Reset() {
	*x = CreateCheckoutSessionRequest{}
	mi := &file_order_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCheckoutSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCheckoutSessionRequest) ProtoMessage() {}

func (x *CreateCheckoutSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCheckoutSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateCheckoutSessionRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{7}
}

func (x *CreateCheckoutSessionRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *CreateCheckoutSessionRequest) GetAbonementId() string {
	if x != nil {
		return x.AbonementId
	}
	return ""
}

func (x *CreateCheckoutSessionRequest) GetStripePriceId() string {
	if x != nil {
		return x.StripePriceId
	}
	return ""
}

type CreateCheckoutSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionUrl string `protobuf:"bytes,1,opt,name=session_url,json=sessionUrl,proto3" json:"session_url,omitempty"`
}

func (x *CreateCheckoutSessionResponse) Reset() {
	*x = CreateCheckoutSessionResponse{}
	mi := &file_order_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCheckoutSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCheckoutSessionResponse) ProtoMessage() {}

func (x *CreateCheckoutSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCheckoutSessionResponse.ProtoReflect.Descriptor instead.
func (*CreateCheckoutSessionResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{8}
}

func (x *CreateCheckoutSessionResponse) GetSessionUrl() string {
	if x != nil {
		return x.SessionUrl
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x66,
	0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x1a, 0x0f, 0x61, 0x62, 0x6f, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x62, 0x6f, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x62, 0x6f, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x8f, 0x02,
	0x0a, 0x24, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x57, 0x69, 0x74,
	0x68, 0x41, 0x62, 0x6f, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x69,
	0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0b,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x53, 0x0a, 0x0f, 0x61,
	0x62, 0x6f, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x61, 0x62, 0x6f, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x41, 0x62, 0x6f, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x0f, 0x61, 0x62, 0x6f, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x4d, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65,
	0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22,
	0x50, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x62, 0x6f, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x62, 0x6f, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x6e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x58, 0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x12, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x22, 0x5a, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2f, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa8,
	0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x24, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x57, 0x69, 0x74, 0x68, 0x41, 0x62, 0x6f, 0x6e,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73,
	0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x57, 0x69, 0x74, 0x68, 0x41, 0x62,
	0x6f, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x24, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x57, 0x69, 0x74, 0x68, 0x41, 0x62, 0x6f, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x1c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x62, 0x6f, 0x6e, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x62, 0x6f, 0x6e, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x22, 0x40, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x55, 0x72, 0x6c, 0x32, 0xd8, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x62,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x28, 0x2e,
	0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73,
	0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x68, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x2a, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x80, 0x01, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x66, 0x69, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x1e, 0x5a, 0x1c, 0x46, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_order_proto_goTypes = []any{
	(*OrderObject)(nil),                                      // 0: fitness_center.order.OrderObject
	(*OrderObjectWithAbonementWithServices)(nil),             // 1: fitness_center.order.OrderObjectWithAbonementWithServices
	(*OrderDataForCreate)(nil),                               // 2: fitness_center.order.OrderDataForCreate
	(*CreateOrderRequest)(nil),                               // 3: fitness_center.order.CreateOrderRequest
	(*CreateOrderResponse)(nil),                              // 4: fitness_center.order.CreateOrderResponse
	(*GetUserOrdersRequest)(nil),                             // 5: fitness_center.order.GetUserOrdersRequest
	(*GetUserOrdersResponse)(nil),                            // 6: fitness_center.order.GetUserOrdersResponse
	(*CreateCheckoutSessionRequest)(nil),                     // 7: fitness_center.order.CreateCheckoutSessionRequest
	(*CreateCheckoutSessionResponse)(nil),                    // 8: fitness_center.order.CreateCheckoutSessionResponse
	(*FitnessCenter_protobuf_abonement.AbonementObject)(nil), // 9: fitness_center.abonement.AbonementObject
	(*FitnessCenter_protobuf_service.ServiceObject)(nil),     // 10: fitness_center.service.ServiceObject
}
var file_order_proto_depIdxs = []int32{
	0,  // 0: fitness_center.order.OrderObjectWithAbonementWithServices.orderObject:type_name -> fitness_center.order.OrderObject
	9,  // 1: fitness_center.order.OrderObjectWithAbonementWithServices.abonementObject:type_name -> fitness_center.abonement.AbonementObject
	10, // 2: fitness_center.order.OrderObjectWithAbonementWithServices.serviceObjects:type_name -> fitness_center.service.ServiceObject
	2,  // 3: fitness_center.order.CreateOrderRequest.orderDataForCreate:type_name -> fitness_center.order.OrderDataForCreate
	0,  // 4: fitness_center.order.CreateOrderResponse.orderObject:type_name -> fitness_center.order.OrderObject
	1,  // 5: fitness_center.order.GetUserOrdersResponse.orderObjectWithAbonementWithServices:type_name -> fitness_center.order.OrderObjectWithAbonementWithServices
	3,  // 6: fitness_center.order.Order.CreateOrder:input_type -> fitness_center.order.CreateOrderRequest
	5,  // 7: fitness_center.order.Order.GetUserOrders:input_type -> fitness_center.order.GetUserOrdersRequest
	7,  // 8: fitness_center.order.Order.CreateCheckoutSession:input_type -> fitness_center.order.CreateCheckoutSessionRequest
	4,  // 9: fitness_center.order.Order.CreateOrder:output_type -> fitness_center.order.CreateOrderResponse
	6,  // 10: fitness_center.order.Order.GetUserOrders:output_type -> fitness_center.order.GetUserOrdersResponse
	8,  // 11: fitness_center.order.Order.CreateCheckoutSession:output_type -> fitness_center.order.CreateCheckoutSessionResponse
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
