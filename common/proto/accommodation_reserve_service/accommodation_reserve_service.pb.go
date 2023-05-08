// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: accommodation_reserve_service.proto

package accommodation_reserve

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CancelReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationID *ReservationID `protobuf:"bytes,1,opt,name=reservationID,proto3" json:"reservationID,omitempty"`
}

func (x *CancelReservationRequest) Reset() {
	*x = CancelReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_reserve_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelReservationRequest) ProtoMessage() {}

func (x *CancelReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_reserve_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelReservationRequest.ProtoReflect.Descriptor instead.
func (*CancelReservationRequest) Descriptor() ([]byte, []int) {
	return file_accommodation_reserve_service_proto_rawDescGZIP(), []int{0}
}

func (x *CancelReservationRequest) GetReservationID() *ReservationID {
	if x != nil {
		return x.ReservationID
	}
	return nil
}

type ReservationID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReservationID) Reset() {
	*x = ReservationID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_reserve_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationID) ProtoMessage() {}

func (x *ReservationID) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_reserve_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationID.ProtoReflect.Descriptor instead.
func (*ReservationID) Descriptor() ([]byte, []int) {
	return file_accommodation_reserve_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReservationID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type IdMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdMessageRequest) Reset() {
	*x = IdMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_reserve_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdMessageRequest) ProtoMessage() {}

func (x *IdMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_reserve_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdMessageRequest.ProtoReflect.Descriptor instead.
func (*IdMessageRequest) Descriptor() ([]byte, []int) {
	return file_accommodation_reserve_service_proto_rawDescGZIP(), []int{2}
}

func (x *IdMessageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetReservationsByUserIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservations []*Reservation `protobuf:"bytes,1,rep,name=reservations,proto3" json:"reservations,omitempty"`
}

func (x *GetReservationsByUserIDResponse) Reset() {
	*x = GetReservationsByUserIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_reserve_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReservationsByUserIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReservationsByUserIDResponse) ProtoMessage() {}

func (x *GetReservationsByUserIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_reserve_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReservationsByUserIDResponse.ProtoReflect.Descriptor instead.
func (*GetReservationsByUserIDResponse) Descriptor() ([]byte, []int) {
	return file_accommodation_reserve_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetReservationsByUserIDResponse) GetReservations() []*Reservation {
	if x != nil {
		return x.Reservations
	}
	return nil
}

type GetAllForConfirmationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllForConfirmationRequest) Reset() {
	*x = GetAllForConfirmationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_reserve_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllForConfirmationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllForConfirmationRequest) ProtoMessage() {}

func (x *GetAllForConfirmationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_reserve_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllForConfirmationRequest.ProtoReflect.Descriptor instead.
func (*GetAllForConfirmationRequest) Descriptor() ([]byte, []int) {
	return file_accommodation_reserve_service_proto_rawDescGZIP(), []int{4}
}

type GetAllForConfirmationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservations []*Reservation `protobuf:"bytes,1,rep,name=reservations,proto3" json:"reservations,omitempty"`
}

func (x *GetAllForConfirmationResponse) Reset() {
	*x = GetAllForConfirmationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_reserve_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllForConfirmationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllForConfirmationResponse) ProtoMessage() {}

func (x *GetAllForConfirmationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_reserve_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllForConfirmationResponse.ProtoReflect.Descriptor instead.
func (*GetAllForConfirmationResponse) Descriptor() ([]byte, []int) {
	return file_accommodation_reserve_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllForConfirmationResponse) GetReservations() []*Reservation {
	if x != nil {
		return x.Reservations
	}
	return nil
}

type CreateReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewReservation *NewReservation `protobuf:"bytes,1,opt,name=newReservation,proto3" json:"newReservation,omitempty"`
}

func (x *CreateReservationRequest) Reset() {
	*x = CreateReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_reserve_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReservationRequest) ProtoMessage() {}

func (x *CreateReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_reserve_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReservationRequest.ProtoReflect.Descriptor instead.
func (*CreateReservationRequest) Descriptor() ([]byte, []int) {
	return file_accommodation_reserve_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreateReservationRequest) GetNewReservation() *NewReservation {
	if x != nil {
		return x.NewReservation
	}
	return nil
}

type CreateReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservation *Reservation `protobuf:"bytes,1,opt,name=reservation,proto3" json:"reservation,omitempty"`
}

func (x *CreateReservationResponse) Reset() {
	*x = CreateReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_reserve_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReservationResponse) ProtoMessage() {}

func (x *CreateReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_reserve_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReservationResponse.ProtoReflect.Descriptor instead.
func (*CreateReservationResponse) Descriptor() ([]byte, []int) {
	return file_accommodation_reserve_service_proto_rawDescGZIP(), []int{7}
}

func (x *CreateReservationResponse) GetReservation() *Reservation {
	if x != nil {
		return x.Reservation
	}
	return nil
}

type CancelReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelReservationResponse) Reset() {
	*x = CancelReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_reserve_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelReservationResponse) ProtoMessage() {}

func (x *CancelReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_reserve_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelReservationResponse.ProtoReflect.Descriptor instead.
func (*CancelReservationResponse) Descriptor() ([]byte, []int) {
	return file_accommodation_reserve_service_proto_rawDescGZIP(), []int{8}
}

type NewReservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate       string `protobuf:"bytes,1,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate         string `protobuf:"bytes,2,opt,name=endDate,proto3" json:"endDate,omitempty"`
	AccommodationID string `protobuf:"bytes,3,opt,name=accommodationID,proto3" json:"accommodationID,omitempty"`
	GuestNum        string `protobuf:"bytes,4,opt,name=guestNum,proto3" json:"guestNum,omitempty"`
	UserID          string `protobuf:"bytes,5,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *NewReservation) Reset() {
	*x = NewReservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_reserve_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewReservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewReservation) ProtoMessage() {}

func (x *NewReservation) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_reserve_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewReservation.ProtoReflect.Descriptor instead.
func (*NewReservation) Descriptor() ([]byte, []int) {
	return file_accommodation_reserve_service_proto_rawDescGZIP(), []int{9}
}

func (x *NewReservation) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *NewReservation) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *NewReservation) GetAccommodationID() string {
	if x != nil {
		return x.AccommodationID
	}
	return ""
}

func (x *NewReservation) GetGuestNum() string {
	if x != nil {
		return x.GuestNum
	}
	return ""
}

func (x *NewReservation) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StartDate       string `protobuf:"bytes,2,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate         string `protobuf:"bytes,3,opt,name=endDate,proto3" json:"endDate,omitempty"`
	AccommodationID string `protobuf:"bytes,4,opt,name=accommodationID,proto3" json:"accommodationID,omitempty"`
	GuestNum        string `protobuf:"bytes,5,opt,name=guestNum,proto3" json:"guestNum,omitempty"`
	UserID          string `protobuf:"bytes,6,opt,name=userID,proto3" json:"userID,omitempty"`
	Status          string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_reserve_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservation) ProtoMessage() {}

func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_reserve_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservation.ProtoReflect.Descriptor instead.
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_accommodation_reserve_service_proto_rawDescGZIP(), []int{10}
}

func (x *Reservation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Reservation) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Reservation) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Reservation) GetAccommodationID() string {
	if x != nil {
		return x.AccommodationID
	}
	return ""
}

func (x *Reservation) GetGuestNum() string {
	if x != nil {
		return x.GuestNum
	}
	return ""
}

func (x *Reservation) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Reservation) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_accommodation_reserve_service_proto protoreflect.FileDescriptor

var file_accommodation_reserve_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x18, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x22, 0x1f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x49, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x69, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0c, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x6f, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x67, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x6f, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x69, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x61, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1b, 0x0a, 0x19, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x67,
	0x75, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67,
	0x75, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22,
	0xcb, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xdb, 0x05,
	0x0a, 0x1d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xa7, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29,
	0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x0e, 0x6e, 0x65, 0x77, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xb2, 0x01, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x33, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46,
	0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xb1,
	0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x27, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x2e, 0x49, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2f, 0x12, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0xa6, 0x01, 0x0a, 0x11, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x28, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x3a, 0x0d, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x42, 0x57, 0x5a, 0x55, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2d, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2d, 0x77, 0x65, 0x62, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accommodation_reserve_service_proto_rawDescOnce sync.Once
	file_accommodation_reserve_service_proto_rawDescData = file_accommodation_reserve_service_proto_rawDesc
)

func file_accommodation_reserve_service_proto_rawDescGZIP() []byte {
	file_accommodation_reserve_service_proto_rawDescOnce.Do(func() {
		file_accommodation_reserve_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_accommodation_reserve_service_proto_rawDescData)
	})
	return file_accommodation_reserve_service_proto_rawDescData
}

var file_accommodation_reserve_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_accommodation_reserve_service_proto_goTypes = []interface{}{
	(*CancelReservationRequest)(nil),        // 0: accommodation_reserve.CancelReservationRequest
	(*ReservationID)(nil),                   // 1: accommodation_reserve.reservationID
	(*IdMessageRequest)(nil),                // 2: accommodation_reserve.IdMessageRequest
	(*GetReservationsByUserIDResponse)(nil), // 3: accommodation_reserve.GetReservationsByUserIDResponse
	(*GetAllForConfirmationRequest)(nil),    // 4: accommodation_reserve.GetAllForConfirmationRequest
	(*GetAllForConfirmationResponse)(nil),   // 5: accommodation_reserve.GetAllForConfirmationResponse
	(*CreateReservationRequest)(nil),        // 6: accommodation_reserve.CreateReservationRequest
	(*CreateReservationResponse)(nil),       // 7: accommodation_reserve.CreateReservationResponse
	(*CancelReservationResponse)(nil),       // 8: accommodation_reserve.CancelReservationResponse
	(*NewReservation)(nil),                  // 9: accommodation_reserve.NewReservation
	(*Reservation)(nil),                     // 10: accommodation_reserve.Reservation
}
var file_accommodation_reserve_service_proto_depIdxs = []int32{
	1,  // 0: accommodation_reserve.CancelReservationRequest.reservationID:type_name -> accommodation_reserve.reservationID
	10, // 1: accommodation_reserve.GetReservationsByUserIDResponse.reservations:type_name -> accommodation_reserve.Reservation
	10, // 2: accommodation_reserve.GetAllForConfirmationResponse.reservations:type_name -> accommodation_reserve.Reservation
	9,  // 3: accommodation_reserve.CreateReservationRequest.newReservation:type_name -> accommodation_reserve.NewReservation
	10, // 4: accommodation_reserve.CreateReservationResponse.reservation:type_name -> accommodation_reserve.Reservation
	6,  // 5: accommodation_reserve.accommodation_reserve_service.CreateReservation:input_type -> accommodation_reserve.CreateReservationRequest
	4,  // 6: accommodation_reserve.accommodation_reserve_service.GetAllForConfirmation:input_type -> accommodation_reserve.GetAllForConfirmationRequest
	2,  // 7: accommodation_reserve.accommodation_reserve_service.GetReservationsByUserID:input_type -> accommodation_reserve.IdMessageRequest
	0,  // 8: accommodation_reserve.accommodation_reserve_service.CancelReservation:input_type -> accommodation_reserve.CancelReservationRequest
	7,  // 9: accommodation_reserve.accommodation_reserve_service.CreateReservation:output_type -> accommodation_reserve.CreateReservationResponse
	5,  // 10: accommodation_reserve.accommodation_reserve_service.GetAllForConfirmation:output_type -> accommodation_reserve.GetAllForConfirmationResponse
	3,  // 11: accommodation_reserve.accommodation_reserve_service.GetReservationsByUserID:output_type -> accommodation_reserve.GetReservationsByUserIDResponse
	8,  // 12: accommodation_reserve.accommodation_reserve_service.CancelReservation:output_type -> accommodation_reserve.CancelReservationResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_accommodation_reserve_service_proto_init() }
func file_accommodation_reserve_service_proto_init() {
	if File_accommodation_reserve_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accommodation_reserve_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelReservationRequest); i {
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
		file_accommodation_reserve_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationID); i {
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
		file_accommodation_reserve_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdMessageRequest); i {
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
		file_accommodation_reserve_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReservationsByUserIDResponse); i {
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
		file_accommodation_reserve_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllForConfirmationRequest); i {
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
		file_accommodation_reserve_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllForConfirmationResponse); i {
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
		file_accommodation_reserve_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReservationRequest); i {
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
		file_accommodation_reserve_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReservationResponse); i {
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
		file_accommodation_reserve_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelReservationResponse); i {
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
		file_accommodation_reserve_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewReservation); i {
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
		file_accommodation_reserve_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reservation); i {
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
			RawDescriptor: file_accommodation_reserve_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accommodation_reserve_service_proto_goTypes,
		DependencyIndexes: file_accommodation_reserve_service_proto_depIdxs,
		MessageInfos:      file_accommodation_reserve_service_proto_msgTypes,
	}.Build()
	File_accommodation_reserve_service_proto = out.File
	file_accommodation_reserve_service_proto_rawDesc = nil
	file_accommodation_reserve_service_proto_goTypes = nil
	file_accommodation_reserve_service_proto_depIdxs = nil
}
