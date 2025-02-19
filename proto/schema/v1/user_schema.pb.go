// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: schema/v1/user_schema.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Nickname     string                 `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email        string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	IsVerified   bool                   `protobuf:"varint,4,opt,name=is_verified,json=isVerified,proto3" json:"is_verified,omitempty"`
	PremiumAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=premium_at,json=premiumAt,proto3" json:"premium_at,omitempty"`
	RegisteredAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=registered_at,json=registeredAt,proto3" json:"registered_at,omitempty"`
	FullName     string                 `protobuf:"bytes,7,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	ImageUrl     string                 `protobuf:"bytes,8,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	BirthAt      string                 `protobuf:"bytes,9,opt,name=birth_at,json=birthAt,proto3" json:"birth_at,omitempty"`
	Gender       Gender                 `protobuf:"varint,10,opt,name=gender,proto3,enum=v1.Gender" json:"gender,omitempty"`
	Company      string                 `protobuf:"bytes,11,opt,name=company,proto3" json:"company,omitempty"`
	JobTitle     string                 `protobuf:"bytes,12,opt,name=job_title,json=jobTitle,proto3" json:"job_title,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_v1_user_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_schema_v1_user_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_schema_v1_user_schema_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetIsVerified() bool {
	if x != nil {
		return x.IsVerified
	}
	return false
}

func (x *User) GetPremiumAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PremiumAt
	}
	return nil
}

func (x *User) GetRegisteredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RegisteredAt
	}
	return nil
}

func (x *User) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *User) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *User) GetBirthAt() string {
	if x != nil {
		return x.BirthAt
	}
	return ""
}

func (x *User) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_UNKNOWN_GENDER
}

func (x *User) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *User) GetJobTitle() string {
	if x != nil {
		return x.JobTitle
	}
	return ""
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current   uint32 `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	Size      uint32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Next      uint32 `protobuf:"varint,3,opt,name=next,proto3" json:"next,omitempty"`
	Prev      uint32 `protobuf:"varint,4,opt,name=prev,proto3" json:"prev,omitempty"`
	Count     uint32 `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	RowsCount uint32 `protobuf:"varint,6,opt,name=rows_count,json=rowsCount,proto3" json:"rows_count,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_v1_user_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_schema_v1_user_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_schema_v1_user_schema_proto_rawDescGZIP(), []int{1}
}

func (x *Page) GetCurrent() uint32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *Page) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Page) GetNext() uint32 {
	if x != nil {
		return x.Next
	}
	return 0
}

func (x *Page) GetPrev() uint32 {
	if x != nil {
		return x.Prev
	}
	return 0
}

func (x *Page) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Page) GetRowsCount() uint32 {
	if x != nil {
		return x.RowsCount
	}
	return 0
}

type ListUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize uint32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListUserRequest) Reset() {
	*x = ListUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_v1_user_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRequest) ProtoMessage() {}

func (x *ListUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_v1_user_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRequest.ProtoReflect.Descriptor instead.
func (*ListUserRequest) Descriptor() ([]byte, []int) {
	return file_schema_v1_user_schema_proto_rawDescGZIP(), []int{2}
}

func (x *ListUserRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListUserRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *Page   `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Data []*User `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListUserResponse) Reset() {
	*x = ListUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_v1_user_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserResponse) ProtoMessage() {}

func (x *ListUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_v1_user_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserResponse.ProtoReflect.Descriptor instead.
func (*ListUserResponse) Descriptor() ([]byte, []int) {
	return file_schema_v1_user_schema_proto_rawDescGZIP(), []int{3}
}

func (x *ListUserResponse) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListUserResponse) GetData() []*User {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateUserPreferenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreferenceUserId uint32         `protobuf:"varint,1,opt,name=preference_user_id,json=preferenceUserId,proto3" json:"preference_user_id,omitempty"`
	PreferenceType   PreferenceType `protobuf:"varint,2,opt,name=preference_type,json=preferenceType,proto3,enum=v1.PreferenceType" json:"preference_type,omitempty"`
}

func (x *CreateUserPreferenceRequest) Reset() {
	*x = CreateUserPreferenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_v1_user_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserPreferenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserPreferenceRequest) ProtoMessage() {}

func (x *CreateUserPreferenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_v1_user_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserPreferenceRequest.ProtoReflect.Descriptor instead.
func (*CreateUserPreferenceRequest) Descriptor() ([]byte, []int) {
	return file_schema_v1_user_schema_proto_rawDescGZIP(), []int{4}
}

func (x *CreateUserPreferenceRequest) GetPreferenceUserId() uint32 {
	if x != nil {
		return x.PreferenceUserId
	}
	return 0
}

func (x *CreateUserPreferenceRequest) GetPreferenceType() PreferenceType {
	if x != nil {
		return x.PreferenceType
	}
	return PreferenceType_UNKNOWN_PREFERENCE_TYPE
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *User `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_v1_user_schema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_v1_user_schema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_schema_v1_user_schema_proto_rawDescGZIP(), []int{5}
}

func (x *UserResponse) GetData() *User {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_schema_v1_user_schema_proto protoreflect.FileDescriptor

var file_schema_v1_user_schema_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76,
	0x31, 0x1a, 0x18, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x95, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x73, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x69, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0a,
	0x70, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x70, 0x72,
	0x65, 0x6d, 0x69, 0x75, 0x6d, 0x41, 0x74, 0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x69, 0x72, 0x74, 0x68, 0x41, 0x74, 0x12, 0x22, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6a,
	0x6f, 0x62, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6a, 0x6f, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x04, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6e,
	0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x72, 0x65, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x70, 0x72, 0x65, 0x76, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x6f, 0x77, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x72, 0x6f, 0x77, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x58, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x2a, 0x04, 0x20, 0x00, 0x40, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x26,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x20, 0x00, 0x40, 0x00, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x4e, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9f, 0x01, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x12, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x20, 0x00, 0x40, 0x00, 0x52, 0x10, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x47, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0a, 0xfa, 0x42,
	0x07, 0x82, 0x01, 0x04, 0x18, 0x01, 0x18, 0x02, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x94, 0x02, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x4c, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x74, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22,
	0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x02, 0x4d, 0x65, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0c, 0x12, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6e, 0x73, 0x32,
	0x30, 0x31, 0x32, 0x2f, 0x64, 0x65, 0x61, 0x6c, 0x6c, 0x73, 0x2d, 0x64, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_v1_user_schema_proto_rawDescOnce sync.Once
	file_schema_v1_user_schema_proto_rawDescData = file_schema_v1_user_schema_proto_rawDesc
)

func file_schema_v1_user_schema_proto_rawDescGZIP() []byte {
	file_schema_v1_user_schema_proto_rawDescOnce.Do(func() {
		file_schema_v1_user_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_v1_user_schema_proto_rawDescData)
	})
	return file_schema_v1_user_schema_proto_rawDescData
}

var file_schema_v1_user_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_schema_v1_user_schema_proto_goTypes = []interface{}{
	(*User)(nil),                        // 0: v1.User
	(*Page)(nil),                        // 1: v1.Page
	(*ListUserRequest)(nil),             // 2: v1.ListUserRequest
	(*ListUserResponse)(nil),            // 3: v1.ListUserResponse
	(*CreateUserPreferenceRequest)(nil), // 4: v1.CreateUserPreferenceRequest
	(*UserResponse)(nil),                // 5: v1.UserResponse
	(*timestamppb.Timestamp)(nil),       // 6: google.protobuf.Timestamp
	(Gender)(0),                         // 7: v1.Gender
	(PreferenceType)(0),                 // 8: v1.PreferenceType
	(*emptypb.Empty)(nil),               // 9: google.protobuf.Empty
}
var file_schema_v1_user_schema_proto_depIdxs = []int32{
	6,  // 0: v1.User.premium_at:type_name -> google.protobuf.Timestamp
	6,  // 1: v1.User.registered_at:type_name -> google.protobuf.Timestamp
	7,  // 2: v1.User.gender:type_name -> v1.Gender
	1,  // 3: v1.ListUserResponse.page:type_name -> v1.Page
	0,  // 4: v1.ListUserResponse.data:type_name -> v1.User
	8,  // 5: v1.CreateUserPreferenceRequest.preference_type:type_name -> v1.PreferenceType
	0,  // 6: v1.UserResponse.data:type_name -> v1.User
	2,  // 7: v1.UserSchema.ListUser:input_type -> v1.ListUserRequest
	4,  // 8: v1.UserSchema.CreateUserPreference:input_type -> v1.CreateUserPreferenceRequest
	9,  // 9: v1.UserSchema.Me:input_type -> google.protobuf.Empty
	3,  // 10: v1.UserSchema.ListUser:output_type -> v1.ListUserResponse
	9,  // 11: v1.UserSchema.CreateUserPreference:output_type -> google.protobuf.Empty
	5,  // 12: v1.UserSchema.Me:output_type -> v1.UserResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_schema_v1_user_schema_proto_init() }
func file_schema_v1_user_schema_proto_init() {
	if File_schema_v1_user_schema_proto != nil {
		return
	}
	file_schema_v1_constant_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schema_v1_user_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_schema_v1_user_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_schema_v1_user_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserRequest); i {
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
		file_schema_v1_user_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserResponse); i {
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
		file_schema_v1_user_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserPreferenceRequest); i {
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
		file_schema_v1_user_schema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
			RawDescriptor: file_schema_v1_user_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schema_v1_user_schema_proto_goTypes,
		DependencyIndexes: file_schema_v1_user_schema_proto_depIdxs,
		MessageInfos:      file_schema_v1_user_schema_proto_msgTypes,
	}.Build()
	File_schema_v1_user_schema_proto = out.File
	file_schema_v1_user_schema_proto_rawDesc = nil
	file_schema_v1_user_schema_proto_goTypes = nil
	file_schema_v1_user_schema_proto_depIdxs = nil
}
