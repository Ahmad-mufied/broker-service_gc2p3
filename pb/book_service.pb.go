// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: book_service.proto

package pb

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

type CreateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author      string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	PublishDate string `protobuf:"bytes,4,opt,name=publish_date,json=publishDate,proto3" json:"publish_date,omitempty"`
	Status      string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_book_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBookRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *CreateBookRequest) GetPublishDate() string {
	if x != nil {
		return x.PublishDate
	}
	return ""
}

func (x *CreateBookRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CreateBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateBookResponse) Reset() {
	*x = CreateBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookResponse) ProtoMessage() {}

func (x *CreateBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookResponse.ProtoReflect.Descriptor instead.
func (*CreateBookResponse) Descriptor() ([]byte, []int) {
	return file_book_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBookResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateBookResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBookByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBookByIdRequest) Reset() {
	*x = GetBookByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookByIdRequest) ProtoMessage() {}

func (x *GetBookByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookByIdRequest.ProtoReflect.Descriptor instead.
func (*GetBookByIdRequest) Descriptor() ([]byte, []int) {
	return file_book_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetBookByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBookByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author      string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	PublishDate string `protobuf:"bytes,4,opt,name=publish_date,json=publishDate,proto3" json:"publish_date,omitempty"`
	Status      string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetBookByIdResponse) Reset() {
	*x = GetBookByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookByIdResponse) ProtoMessage() {}

func (x *GetBookByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookByIdResponse.ProtoReflect.Descriptor instead.
func (*GetBookByIdResponse) Descriptor() ([]byte, []int) {
	return file_book_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetBookByIdResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetBookByIdResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetBookByIdResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *GetBookByIdResponse) GetPublishDate() string {
	if x != nil {
		return x.PublishDate
	}
	return ""
}

func (x *GetBookByIdResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetBookByTitleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *GetBookByTitleRequest) Reset() {
	*x = GetBookByTitleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookByTitleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookByTitleRequest) ProtoMessage() {}

func (x *GetBookByTitleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookByTitleRequest.ProtoReflect.Descriptor instead.
func (*GetBookByTitleRequest) Descriptor() ([]byte, []int) {
	return file_book_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetBookByTitleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GetBookByTitleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Author      string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	PublishDate string `protobuf:"bytes,3,opt,name=publish_date,json=publishDate,proto3" json:"publish_date,omitempty"`
	Status      string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetBookByTitleResponse) Reset() {
	*x = GetBookByTitleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookByTitleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookByTitleResponse) ProtoMessage() {}

func (x *GetBookByTitleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookByTitleResponse.ProtoReflect.Descriptor instead.
func (*GetBookByTitleResponse) Descriptor() ([]byte, []int) {
	return file_book_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetBookByTitleResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetBookByTitleResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *GetBookByTitleResponse) GetPublishDate() string {
	if x != nil {
		return x.PublishDate
	}
	return ""
}

func (x *GetBookByTitleResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author      string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	PublishDate string `protobuf:"bytes,4,opt,name=publish_date,json=publishDate,proto3" json:"publish_date,omitempty"`
	Status      string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateBookRequest) Reset() {
	*x = UpdateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRequest) ProtoMessage() {}

func (x *UpdateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return file_book_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBookRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *UpdateBookRequest) GetPublishDate() string {
	if x != nil {
		return x.PublishDate
	}
	return ""
}

func (x *UpdateBookRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UpdateBookResponse) Reset() {
	*x = UpdateBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookResponse) ProtoMessage() {}

func (x *UpdateBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookResponse.ProtoReflect.Descriptor instead.
func (*UpdateBookResponse) Descriptor() ([]byte, []int) {
	return file_book_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateBookResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBookRequest) Reset() {
	*x = DeleteBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequest) ProtoMessage() {}

func (x *DeleteBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return file_book_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteBookRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteBookResponse) Reset() {
	*x = DeleteBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookResponse) ProtoMessage() {}

func (x *DeleteBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookResponse.ProtoReflect.Descriptor instead.
func (*DeleteBookResponse) Descriptor() ([]byte, []int) {
	return file_book_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteBookResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CheckBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CheckBookRequest) Reset() {
	*x = CheckBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckBookRequest) ProtoMessage() {}

func (x *CheckBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckBookRequest.ProtoReflect.Descriptor instead.
func (*CheckBookRequest) Descriptor() ([]byte, []int) {
	return file_book_service_proto_rawDescGZIP(), []int{10}
}

func (x *CheckBookRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CheckBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist bool `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
}

func (x *CheckBookResponse) Reset() {
	*x = CheckBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckBookResponse) ProtoMessage() {}

func (x *CheckBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckBookResponse.ProtoReflect.Descriptor instead.
func (*CheckBookResponse) Descriptor() ([]byte, []int) {
	return file_book_service_proto_rawDescGZIP(), []int{11}
}

func (x *CheckBookResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

var File_book_service_proto protoreflect.FileDescriptor

var file_book_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x8c, 0x01, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3e, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x98, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x7b, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29,
	0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x32, 0x8f, 0x03, 0x0a, 0x0b, 0x42, 0x6f,
	0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x62,
	0x6f, 0x6f, 0x6b, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x63, 0x32, 0x70,
	0x33, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_book_service_proto_rawDescOnce sync.Once
	file_book_service_proto_rawDescData = file_book_service_proto_rawDesc
)

func file_book_service_proto_rawDescGZIP() []byte {
	file_book_service_proto_rawDescOnce.Do(func() {
		file_book_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_book_service_proto_rawDescData)
	})
	return file_book_service_proto_rawDescData
}

var file_book_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_book_service_proto_goTypes = []any{
	(*CreateBookRequest)(nil),      // 0: book.CreateBookRequest
	(*CreateBookResponse)(nil),     // 1: book.CreateBookResponse
	(*GetBookByIdRequest)(nil),     // 2: book.GetBookByIdRequest
	(*GetBookByIdResponse)(nil),    // 3: book.GetBookByIdResponse
	(*GetBookByTitleRequest)(nil),  // 4: book.GetBookByTitleRequest
	(*GetBookByTitleResponse)(nil), // 5: book.GetBookByTitleResponse
	(*UpdateBookRequest)(nil),      // 6: book.UpdateBookRequest
	(*UpdateBookResponse)(nil),     // 7: book.UpdateBookResponse
	(*DeleteBookRequest)(nil),      // 8: book.DeleteBookRequest
	(*DeleteBookResponse)(nil),     // 9: book.DeleteBookResponse
	(*CheckBookRequest)(nil),       // 10: book.CheckBookRequest
	(*CheckBookResponse)(nil),      // 11: book.CheckBookResponse
}
var file_book_service_proto_depIdxs = []int32{
	0,  // 0: book.BookService.Create:input_type -> book.CreateBookRequest
	2,  // 1: book.BookService.GetBookById:input_type -> book.GetBookByIdRequest
	4,  // 2: book.BookService.GetBookByTitle:input_type -> book.GetBookByTitleRequest
	6,  // 3: book.BookService.Update:input_type -> book.UpdateBookRequest
	8,  // 4: book.BookService.Delete:input_type -> book.DeleteBookRequest
	10, // 5: book.BookService.Check:input_type -> book.CheckBookRequest
	1,  // 6: book.BookService.Create:output_type -> book.CreateBookResponse
	3,  // 7: book.BookService.GetBookById:output_type -> book.GetBookByIdResponse
	5,  // 8: book.BookService.GetBookByTitle:output_type -> book.GetBookByTitleResponse
	7,  // 9: book.BookService.Update:output_type -> book.UpdateBookResponse
	9,  // 10: book.BookService.Delete:output_type -> book.DeleteBookResponse
	11, // 11: book.BookService.Check:output_type -> book.CheckBookResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_book_service_proto_init() }
func file_book_service_proto_init() {
	if File_book_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_book_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateBookRequest); i {
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
		file_book_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateBookResponse); i {
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
		file_book_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetBookByIdRequest); i {
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
		file_book_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetBookByIdResponse); i {
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
		file_book_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetBookByTitleRequest); i {
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
		file_book_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetBookByTitleResponse); i {
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
		file_book_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateBookRequest); i {
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
		file_book_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateBookResponse); i {
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
		file_book_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteBookRequest); i {
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
		file_book_service_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteBookResponse); i {
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
		file_book_service_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*CheckBookRequest); i {
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
		file_book_service_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*CheckBookResponse); i {
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
			RawDescriptor: file_book_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_book_service_proto_goTypes,
		DependencyIndexes: file_book_service_proto_depIdxs,
		MessageInfos:      file_book_service_proto_msgTypes,
	}.Build()
	File_book_service_proto = out.File
	file_book_service_proto_rawDesc = nil
	file_book_service_proto_goTypes = nil
	file_book_service_proto_depIdxs = nil
}
