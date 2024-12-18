// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: contract.proto

package genprotos

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

type CreateContractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerName           string `protobuf:"bytes,1,opt,name=consumer_name,json=consumerName,proto3" json:"consumer_name,omitempty"`
	ConsumerPassportSerial string `protobuf:"bytes,2,opt,name=consumer_passport_serial,json=consumerPassportSerial,proto3" json:"consumer_passport_serial,omitempty"`
	ConsumerAddress        string `protobuf:"bytes,3,opt,name=consumer_address,json=consumerAddress,proto3" json:"consumer_address,omitempty"`
	ConsumerPhoneNumber    string `protobuf:"bytes,4,opt,name=consumer_phone_number,json=consumerPhoneNumber,proto3" json:"consumer_phone_number,omitempty"`
	PassportImage          string `protobuf:"bytes,5,opt,name=passport_image,json=passportImage,proto3" json:"passport_image,omitempty"`
	Status                 string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"` // Enum: pending, finished, canceled
	Duration               int32  `protobuf:"varint,7,opt,name=duration,proto3" json:"duration,omitempty"`
	StorageId              string `protobuf:"bytes,8,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
}

func (x *CreateContractRequest) Reset() {
	*x = CreateContractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContractRequest) ProtoMessage() {}

func (x *CreateContractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContractRequest.ProtoReflect.Descriptor instead.
func (*CreateContractRequest) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{0}
}

func (x *CreateContractRequest) GetConsumerName() string {
	if x != nil {
		return x.ConsumerName
	}
	return ""
}

func (x *CreateContractRequest) GetConsumerPassportSerial() string {
	if x != nil {
		return x.ConsumerPassportSerial
	}
	return ""
}

func (x *CreateContractRequest) GetConsumerAddress() string {
	if x != nil {
		return x.ConsumerAddress
	}
	return ""
}

func (x *CreateContractRequest) GetConsumerPhoneNumber() string {
	if x != nil {
		return x.ConsumerPhoneNumber
	}
	return ""
}

func (x *CreateContractRequest) GetPassportImage() string {
	if x != nil {
		return x.PassportImage
	}
	return ""
}

func (x *CreateContractRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateContractRequest) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *CreateContractRequest) GetStorageId() string {
	if x != nil {
		return x.StorageId
	}
	return ""
}

type CreateContractRequestSwagger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerName           string `protobuf:"bytes,1,opt,name=consumer_name,json=consumerName,proto3" json:"consumer_name,omitempty"`
	ConsumerPassportSerial string `protobuf:"bytes,2,opt,name=consumer_passport_serial,json=consumerPassportSerial,proto3" json:"consumer_passport_serial,omitempty"`
	ConsumerAddress        string `protobuf:"bytes,3,opt,name=consumer_address,json=consumerAddress,proto3" json:"consumer_address,omitempty"`
	ConsumerPhoneNumber    string `protobuf:"bytes,4,opt,name=consumer_phone_number,json=consumerPhoneNumber,proto3" json:"consumer_phone_number,omitempty"`
	Duration               int32  `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
	PassportImage          string `protobuf:"bytes,6,opt,name=passport_image,json=passportImage,proto3" json:"passport_image,omitempty"`
}

func (x *CreateContractRequestSwagger) Reset() {
	*x = CreateContractRequestSwagger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContractRequestSwagger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContractRequestSwagger) ProtoMessage() {}

func (x *CreateContractRequestSwagger) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContractRequestSwagger.ProtoReflect.Descriptor instead.
func (*CreateContractRequestSwagger) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{1}
}

func (x *CreateContractRequestSwagger) GetConsumerName() string {
	if x != nil {
		return x.ConsumerName
	}
	return ""
}

func (x *CreateContractRequestSwagger) GetConsumerPassportSerial() string {
	if x != nil {
		return x.ConsumerPassportSerial
	}
	return ""
}

func (x *CreateContractRequestSwagger) GetConsumerAddress() string {
	if x != nil {
		return x.ConsumerAddress
	}
	return ""
}

func (x *CreateContractRequestSwagger) GetConsumerPhoneNumber() string {
	if x != nil {
		return x.ConsumerPhoneNumber
	}
	return ""
}

func (x *CreateContractRequestSwagger) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *CreateContractRequestSwagger) GetPassportImage() string {
	if x != nil {
		return x.PassportImage
	}
	return ""
}

type ContractIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ContractIdRequest) Reset() {
	*x = ContractIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractIdRequest) ProtoMessage() {}

func (x *ContractIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractIdRequest.ProtoReflect.Descriptor instead.
func (*ContractIdRequest) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{2}
}

func (x *ContractIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetContractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ConsumerName           string  `protobuf:"bytes,2,opt,name=consumer_name,json=consumerName,proto3" json:"consumer_name,omitempty"`
	ConsumerPassportSerial string  `protobuf:"bytes,3,opt,name=consumer_passport_serial,json=consumerPassportSerial,proto3" json:"consumer_passport_serial,omitempty"`
	ConsumerAddress        string  `protobuf:"bytes,4,opt,name=consumer_address,json=consumerAddress,proto3" json:"consumer_address,omitempty"`
	ConsumerPhoneNumber    string  `protobuf:"bytes,5,opt,name=consumer_phone_number,json=consumerPhoneNumber,proto3" json:"consumer_phone_number,omitempty"`
	PassportImage          string  `protobuf:"bytes,6,opt,name=passport_image,json=passportImage,proto3" json:"passport_image,omitempty"`
	Status                 string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Duration               int32   `protobuf:"varint,8,opt,name=duration,proto3" json:"duration,omitempty"`
	CreatedAt              string  `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DeletedAt              string  `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Price                  float64 `protobuf:"fixed64,11,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *GetContractResponse) Reset() {
	*x = GetContractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContractResponse) ProtoMessage() {}

func (x *GetContractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContractResponse.ProtoReflect.Descriptor instead.
func (*GetContractResponse) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{3}
}

func (x *GetContractResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetContractResponse) GetConsumerName() string {
	if x != nil {
		return x.ConsumerName
	}
	return ""
}

func (x *GetContractResponse) GetConsumerPassportSerial() string {
	if x != nil {
		return x.ConsumerPassportSerial
	}
	return ""
}

func (x *GetContractResponse) GetConsumerAddress() string {
	if x != nil {
		return x.ConsumerAddress
	}
	return ""
}

func (x *GetContractResponse) GetConsumerPhoneNumber() string {
	if x != nil {
		return x.ConsumerPhoneNumber
	}
	return ""
}

func (x *GetContractResponse) GetPassportImage() string {
	if x != nil {
		return x.PassportImage
	}
	return ""
}

func (x *GetContractResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetContractResponse) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *GetContractResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetContractResponse) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *GetContractResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type UpdateContractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ConsumerName           string `protobuf:"bytes,2,opt,name=consumer_name,json=consumerName,proto3" json:"consumer_name,omitempty"`
	ConsumerPassportSerial string `protobuf:"bytes,3,opt,name=consumer_passport_serial,json=consumerPassportSerial,proto3" json:"consumer_passport_serial,omitempty"`
	ConsumerAddress        string `protobuf:"bytes,4,opt,name=consumer_address,json=consumerAddress,proto3" json:"consumer_address,omitempty"`
	ConsumerPhoneNumber    string `protobuf:"bytes,5,opt,name=consumer_phone_number,json=consumerPhoneNumber,proto3" json:"consumer_phone_number,omitempty"`
	PassportImage          string `protobuf:"bytes,6,opt,name=passport_image,json=passportImage,proto3" json:"passport_image,omitempty"`
	Status                 string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Duration               int32  `protobuf:"varint,8,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *UpdateContractRequest) Reset() {
	*x = UpdateContractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateContractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateContractRequest) ProtoMessage() {}

func (x *UpdateContractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateContractRequest.ProtoReflect.Descriptor instead.
func (*UpdateContractRequest) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateContractRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateContractRequest) GetConsumerName() string {
	if x != nil {
		return x.ConsumerName
	}
	return ""
}

func (x *UpdateContractRequest) GetConsumerPassportSerial() string {
	if x != nil {
		return x.ConsumerPassportSerial
	}
	return ""
}

func (x *UpdateContractRequest) GetConsumerAddress() string {
	if x != nil {
		return x.ConsumerAddress
	}
	return ""
}

func (x *UpdateContractRequest) GetConsumerPhoneNumber() string {
	if x != nil {
		return x.ConsumerPhoneNumber
	}
	return ""
}

func (x *UpdateContractRequest) GetPassportImage() string {
	if x != nil {
		return x.PassportImage
	}
	return ""
}

func (x *UpdateContractRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateContractRequest) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type ContractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ContractResponse) Reset() {
	*x = ContractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractResponse) ProtoMessage() {}

func (x *ContractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractResponse.ProtoReflect.Descriptor instead.
func (*ContractResponse) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{5}
}

func (x *ContractResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ContractResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetAllContractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerName string `protobuf:"bytes,1,opt,name=consumer_name,json=consumerName,proto3" json:"consumer_name,omitempty"`
	Status       string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	PasportSeria string `protobuf:"bytes,3,opt,name=pasport_seria,json=pasportSeria,proto3" json:"pasport_seria,omitempty"`
	StorageId    string `protobuf:"bytes,4,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
	Limit        int32  `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Page         int32  `protobuf:"varint,6,opt,name=Page,proto3" json:"Page,omitempty"`
}

func (x *GetAllContractRequest) Reset() {
	*x = GetAllContractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllContractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllContractRequest) ProtoMessage() {}

func (x *GetAllContractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllContractRequest.ProtoReflect.Descriptor instead.
func (*GetAllContractRequest) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllContractRequest) GetConsumerName() string {
	if x != nil {
		return x.ConsumerName
	}
	return ""
}

func (x *GetAllContractRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetAllContractRequest) GetPasportSeria() string {
	if x != nil {
		return x.PasportSeria
	}
	return ""
}

func (x *GetAllContractRequest) GetStorageId() string {
	if x != nil {
		return x.StorageId
	}
	return ""
}

func (x *GetAllContractRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllContractRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetAllContractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllContracts []*GetContractResponse `protobuf:"bytes,1,rep,name=all_contracts,json=allContracts,proto3" json:"all_contracts,omitempty"`
	Count        int32                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Message      string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetAllContractResponse) Reset() {
	*x = GetAllContractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllContractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllContractResponse) ProtoMessage() {}

func (x *GetAllContractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllContractResponse.ProtoReflect.Descriptor instead.
func (*GetAllContractResponse) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllContractResponse) GetAllContracts() []*GetContractResponse {
	if x != nil {
		return x.AllContracts
	}
	return nil
}

func (x *GetAllContractResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetAllContractResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_contract_proto protoreflect.FileDescriptor

var file_contract_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0xcf, 0x02, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x15,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x9f, 0x02, 0x0a, 0x1c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x53, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x38, 0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x16, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x23, 0x0a, 0x11,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x92, 0x03, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38,
	0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x16, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0xc0, 0x02, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12,
	0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25,
	0x0a, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x46, 0x0a, 0x10, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0xc2, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x61, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x40, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x85, 0x03, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_contract_proto_rawDescOnce sync.Once
	file_contract_proto_rawDescData = file_contract_proto_rawDesc
)

func file_contract_proto_rawDescGZIP() []byte {
	file_contract_proto_rawDescOnce.Do(func() {
		file_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_proto_rawDescData)
	})
	return file_contract_proto_rawDescData
}

var file_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_contract_proto_goTypes = []interface{}{
	(*CreateContractRequest)(nil),        // 0: protos.CreateContractRequest
	(*CreateContractRequestSwagger)(nil), // 1: protos.CreateContractRequestSwagger
	(*ContractIdRequest)(nil),            // 2: protos.ContractIdRequest
	(*GetContractResponse)(nil),          // 3: protos.GetContractResponse
	(*UpdateContractRequest)(nil),        // 4: protos.UpdateContractRequest
	(*ContractResponse)(nil),             // 5: protos.ContractResponse
	(*GetAllContractRequest)(nil),        // 6: protos.GetAllContractRequest
	(*GetAllContractResponse)(nil),       // 7: protos.GetAllContractResponse
}
var file_contract_proto_depIdxs = []int32{
	3, // 0: protos.GetAllContractResponse.all_contracts:type_name -> protos.GetContractResponse
	0, // 1: protos.ContractService.CreateContract:input_type -> protos.CreateContractRequest
	2, // 2: protos.ContractService.GetContract:input_type -> protos.ContractIdRequest
	4, // 3: protos.ContractService.UpdateContract:input_type -> protos.UpdateContractRequest
	2, // 4: protos.ContractService.DeleteContract:input_type -> protos.ContractIdRequest
	6, // 5: protos.ContractService.ListContracts:input_type -> protos.GetAllContractRequest
	5, // 6: protos.ContractService.CreateContract:output_type -> protos.ContractResponse
	3, // 7: protos.ContractService.GetContract:output_type -> protos.GetContractResponse
	5, // 8: protos.ContractService.UpdateContract:output_type -> protos.ContractResponse
	5, // 9: protos.ContractService.DeleteContract:output_type -> protos.ContractResponse
	7, // 10: protos.ContractService.ListContracts:output_type -> protos.GetAllContractResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contract_proto_init() }
func file_contract_proto_init() {
	if File_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContractRequest); i {
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
		file_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContractRequestSwagger); i {
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
		file_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractIdRequest); i {
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
		file_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContractResponse); i {
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
		file_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateContractRequest); i {
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
		file_contract_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractResponse); i {
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
		file_contract_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllContractRequest); i {
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
		file_contract_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllContractResponse); i {
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
			RawDescriptor: file_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_proto_goTypes,
		DependencyIndexes: file_contract_proto_depIdxs,
		MessageInfos:      file_contract_proto_msgTypes,
	}.Build()
	File_contract_proto = out.File
	file_contract_proto_rawDesc = nil
	file_contract_proto_goTypes = nil
	file_contract_proto_depIdxs = nil
}
