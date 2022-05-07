//
//Copyright © 2022 Infinite Devices GmbH
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: pkg/node/proto/devices/devices.proto

package devices

import (
	access "github.com/infinimesh/infinimesh/pkg/node/proto/access"
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

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid         string         `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title        string         `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Enabled      bool           `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Certificate  *Certificate   `protobuf:"bytes,4,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Tags         []string       `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	BasicEnabled bool           `protobuf:"varint,6,opt,name=basic_enabled,json=basicEnabled,proto3" json:"basic_enabled,omitempty"`
	Token        string         `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	Access       *access.Access `protobuf:"bytes,8,opt,name=access,proto3,oneof" json:"access,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_node_proto_devices_devices_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_node_proto_devices_devices_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_pkg_node_proto_devices_devices_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Device) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Device) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Device) GetCertificate() *Certificate {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *Device) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Device) GetBasicEnabled() bool {
	if x != nil {
		return x.BasicEnabled
	}
	return false
}

func (x *Device) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Device) GetAccess() *access.Access {
	if x != nil {
		return x.Access
	}
	return nil
}

type Certificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PemData     string `protobuf:"bytes,1,opt,name=pem_data,json=pemData,proto3" json:"pem_data,omitempty"`
	Algorithm   string `protobuf:"bytes,2,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Fingerprint []byte `protobuf:"bytes,3,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
}

func (x *Certificate) Reset() {
	*x = Certificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_node_proto_devices_devices_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificate) ProtoMessage() {}

func (x *Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_node_proto_devices_devices_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificate.ProtoReflect.Descriptor instead.
func (*Certificate) Descriptor() ([]byte, []int) {
	return file_pkg_node_proto_devices_devices_proto_rawDescGZIP(), []int{1}
}

func (x *Certificate) GetPemData() string {
	if x != nil {
		return x.PemData
	}
	return ""
}

func (x *Certificate) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

func (x *Certificate) GetFingerprint() []byte {
	if x != nil {
		return x.Fingerprint
	}
	return nil
}

type Devices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *Devices) Reset() {
	*x = Devices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_node_proto_devices_devices_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Devices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Devices) ProtoMessage() {}

func (x *Devices) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_node_proto_devices_devices_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Devices.ProtoReflect.Descriptor instead.
func (*Devices) Descriptor() ([]byte, []int) {
	return file_pkg_node_proto_devices_devices_proto_rawDescGZIP(), []int{2}
}

func (x *Devices) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device    *Device `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Namespace string  `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"` // Namespace to put device under
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_node_proto_devices_devices_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_node_proto_devices_devices_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_node_proto_devices_devices_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRequest) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *CreateRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device *Device `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_node_proto_devices_devices_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_node_proto_devices_devices_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_pkg_node_proto_devices_devices_proto_rawDescGZIP(), []int{4}
}

func (x *CreateResponse) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

type GetByFingerprintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fingerprint []byte `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
}

func (x *GetByFingerprintRequest) Reset() {
	*x = GetByFingerprintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_node_proto_devices_devices_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByFingerprintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByFingerprintRequest) ProtoMessage() {}

func (x *GetByFingerprintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_node_proto_devices_devices_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByFingerprintRequest.ProtoReflect.Descriptor instead.
func (*GetByFingerprintRequest) Descriptor() ([]byte, []int) {
	return file_pkg_node_proto_devices_devices_proto_rawDescGZIP(), []int{5}
}

func (x *GetByFingerprintRequest) GetFingerprint() []byte {
	if x != nil {
		return x.Fingerprint
	}
	return nil
}

var File_pkg_node_proto_devices_devices_proto protoreflect.FileDescriptor

var file_pkg_node_proto_devices_devices_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x22, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x46, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x69, 0x63, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3b, 0x0a, 0x06, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x69,
	0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x06, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x68, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x65, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6e,
	0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x22, 0x48, 0x0a, 0x0b, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x39, 0x0a, 0x07, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e,
	0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x66, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x49, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x37, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x3b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x42, 0xe2, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e,
	0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x0c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x69, 0x6e, 0x66,
	0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02,
	0x03, 0x49, 0x4e, 0x44, 0xaa, 0x02, 0x17, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02,
	0x17, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x4e, 0x6f, 0x64, 0x65,
	0x5c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0xe2, 0x02, 0x23, 0x49, 0x6e, 0x66, 0x69, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x4e, 0x6f, 0x64, 0x65, 0x5c, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x19, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x4e, 0x6f, 0x64,
	0x65, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_node_proto_devices_devices_proto_rawDescOnce sync.Once
	file_pkg_node_proto_devices_devices_proto_rawDescData = file_pkg_node_proto_devices_devices_proto_rawDesc
)

func file_pkg_node_proto_devices_devices_proto_rawDescGZIP() []byte {
	file_pkg_node_proto_devices_devices_proto_rawDescOnce.Do(func() {
		file_pkg_node_proto_devices_devices_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_node_proto_devices_devices_proto_rawDescData)
	})
	return file_pkg_node_proto_devices_devices_proto_rawDescData
}

var file_pkg_node_proto_devices_devices_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_node_proto_devices_devices_proto_goTypes = []interface{}{
	(*Device)(nil),                  // 0: infinimesh.node.devices.Device
	(*Certificate)(nil),             // 1: infinimesh.node.devices.Certificate
	(*Devices)(nil),             // 2: infinimesh.node.devices.Devices
	(*CreateRequest)(nil),           // 3: infinimesh.node.devices.CreateRequest
	(*CreateResponse)(nil),          // 4: infinimesh.node.devices.CreateResponse
	(*GetByFingerprintRequest)(nil), // 5: infinimesh.node.devices.GetByFingerprintRequest
	(*access.Access)(nil),           // 6: infinimesh.node.access.Access
}
var file_pkg_node_proto_devices_devices_proto_depIdxs = []int32{
	1, // 0: infinimesh.node.devices.Device.certificate:type_name -> infinimesh.node.devices.Certificate
	6, // 1: infinimesh.node.devices.Device.access:type_name -> infinimesh.node.access.Access
	0, // 2: infinimesh.node.devices.Devices.devices:type_name -> infinimesh.node.devices.Device
	0, // 3: infinimesh.node.devices.CreateRequest.device:type_name -> infinimesh.node.devices.Device
	0, // 4: infinimesh.node.devices.CreateResponse.device:type_name -> infinimesh.node.devices.Device
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_node_proto_devices_devices_proto_init() }
func file_pkg_node_proto_devices_devices_proto_init() {
	if File_pkg_node_proto_devices_devices_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_node_proto_devices_devices_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_pkg_node_proto_devices_devices_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificate); i {
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
		file_pkg_node_proto_devices_devices_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Devices); i {
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
		file_pkg_node_proto_devices_devices_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_pkg_node_proto_devices_devices_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_pkg_node_proto_devices_devices_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByFingerprintRequest); i {
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
	file_pkg_node_proto_devices_devices_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_node_proto_devices_devices_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_node_proto_devices_devices_proto_goTypes,
		DependencyIndexes: file_pkg_node_proto_devices_devices_proto_depIdxs,
		MessageInfos:      file_pkg_node_proto_devices_devices_proto_msgTypes,
	}.Build()
	File_pkg_node_proto_devices_devices_proto = out.File
	file_pkg_node_proto_devices_devices_proto_rawDesc = nil
	file_pkg_node_proto_devices_devices_proto_goTypes = nil
	file_pkg_node_proto_devices_devices_proto_depIdxs = nil
}
