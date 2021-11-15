//
//Copyright © 2021 Nikita Ivanovski info@slnt-opp.xyz
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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: pkg/drivers/instance/vanilla/driver.proto

package vanilla

import (
	proto "github.com/slntopp/nocloud/pkg/instances/proto"
	proto1 "github.com/slntopp/nocloud/pkg/services_providers/proto"
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

type GetTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTypeRequest) Reset() {
	*x = GetTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_drivers_instance_vanilla_driver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTypeRequest) ProtoMessage() {}

func (x *GetTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_drivers_instance_vanilla_driver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTypeRequest.ProtoReflect.Descriptor instead.
func (*GetTypeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_drivers_instance_vanilla_driver_proto_rawDescGZIP(), []int{0}
}

type GetTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GetTypeResponse) Reset() {
	*x = GetTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_drivers_instance_vanilla_driver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTypeResponse) ProtoMessage() {}

func (x *GetTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_drivers_instance_vanilla_driver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTypeResponse.ProtoReflect.Descriptor instead.
func (*GetTypeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_drivers_instance_vanilla_driver_proto_rawDescGZIP(), []int{1}
}

func (x *GetTypeResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type DeployRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group            *proto.InstancesGroup    `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	ServicesProvider *proto1.ServicesProvider `protobuf:"bytes,2,opt,name=services_provider,json=servicesProvider,proto3" json:"services_provider,omitempty"`
}

func (x *DeployRequest) Reset() {
	*x = DeployRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_drivers_instance_vanilla_driver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployRequest) ProtoMessage() {}

func (x *DeployRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_drivers_instance_vanilla_driver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployRequest.ProtoReflect.Descriptor instead.
func (*DeployRequest) Descriptor() ([]byte, []int) {
	return file_pkg_drivers_instance_vanilla_driver_proto_rawDescGZIP(), []int{2}
}

func (x *DeployRequest) GetGroup() *proto.InstancesGroup {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *DeployRequest) GetServicesProvider() *proto1.ServicesProvider {
	if x != nil {
		return x.ServicesProvider
	}
	return nil
}

var File_pkg_drivers_instance_vanilla_driver_proto protoreflect.FileDescriptor

var file_pkg_drivers_instance_vanilla_driver_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x61, 0x6e, 0x69, 0x6c, 0x6c, 0x61, 0x2f, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x6e, 0x6f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6e, 0x69, 0x6c, 0x6c, 0x61, 0x1a, 0x23, 0x70, 0x6b,
	0x67, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x35, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0xa3, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x59, 0x0a, 0x11,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x32, 0xe4, 0x02, 0x0a, 0x0d, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x14, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x79, 0x6e, 0x74,
	0x61, 0x78, 0x12, 0x36, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x6e, 0x6f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f,
	0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6e, 0x69, 0x6c, 0x6c, 0x61,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6e, 0x69, 0x6c, 0x6c,
	0x61, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5b, 0x0a, 0x06, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x2e, 0x2e, 0x6e, 0x6f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x61, 0x6e, 0x69, 0x6c, 0x6c, 0x61, 0x2e, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x6f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x8b,
	0x02, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x61, 0x6e, 0x69, 0x6c, 0x6c, 0x61, 0x42, 0x0b, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6c, 0x6e, 0x74, 0x6f, 0x70, 0x70, 0x2f, 0x6e, 0x6f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x61, 0x6e, 0x69, 0x6c, 0x6c, 0x61, 0xa2, 0x02,
	0x04, 0x4e, 0x49, 0x44, 0x56, 0xaa, 0x02, 0x1f, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e,
	0x56, 0x61, 0x6e, 0x69, 0x6c, 0x6c, 0x61, 0xca, 0x02, 0x1f, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x5c, 0x56, 0x61, 0x6e, 0x69, 0x6c, 0x6c, 0x61, 0xe2, 0x02, 0x2b, 0x4e, 0x6f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x5c, 0x56, 0x61, 0x6e, 0x69, 0x6c, 0x6c, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x22, 0x4e, 0x6f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x3a, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x61, 0x6e, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_drivers_instance_vanilla_driver_proto_rawDescOnce sync.Once
	file_pkg_drivers_instance_vanilla_driver_proto_rawDescData = file_pkg_drivers_instance_vanilla_driver_proto_rawDesc
)

func file_pkg_drivers_instance_vanilla_driver_proto_rawDescGZIP() []byte {
	file_pkg_drivers_instance_vanilla_driver_proto_rawDescOnce.Do(func() {
		file_pkg_drivers_instance_vanilla_driver_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_drivers_instance_vanilla_driver_proto_rawDescData)
	})
	return file_pkg_drivers_instance_vanilla_driver_proto_rawDescData
}

var file_pkg_drivers_instance_vanilla_driver_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_drivers_instance_vanilla_driver_proto_goTypes = []interface{}{
	(*GetTypeRequest)(nil),                             // 0: nocloud.instance.driver.vanilla.GetTypeRequest
	(*GetTypeResponse)(nil),                            // 1: nocloud.instance.driver.vanilla.GetTypeResponse
	(*DeployRequest)(nil),                              // 2: nocloud.instance.driver.vanilla.DeployRequest
	(*proto.InstancesGroup)(nil),                       // 3: nocloud.instances.InstancesGroup
	(*proto1.ServicesProvider)(nil),                    // 4: nocloud.services_providers.ServicesProvider
	(*proto.ValidateInstancesGroupConfigRequest)(nil),  // 5: nocloud.instances.ValidateInstancesGroupConfigRequest
	(*proto.ValidateInstancesGroupConfigResponse)(nil), // 6: nocloud.instances.ValidateInstancesGroupConfigResponse
}
var file_pkg_drivers_instance_vanilla_driver_proto_depIdxs = []int32{
	3, // 0: nocloud.instance.driver.vanilla.DeployRequest.group:type_name -> nocloud.instances.InstancesGroup
	4, // 1: nocloud.instance.driver.vanilla.DeployRequest.services_provider:type_name -> nocloud.services_providers.ServicesProvider
	5, // 2: nocloud.instance.driver.vanilla.DriverService.ValidateConfigSyntax:input_type -> nocloud.instances.ValidateInstancesGroupConfigRequest
	0, // 3: nocloud.instance.driver.vanilla.DriverService.GetType:input_type -> nocloud.instance.driver.vanilla.GetTypeRequest
	2, // 4: nocloud.instance.driver.vanilla.DriverService.Deploy:input_type -> nocloud.instance.driver.vanilla.DeployRequest
	6, // 5: nocloud.instance.driver.vanilla.DriverService.ValidateConfigSyntax:output_type -> nocloud.instances.ValidateInstancesGroupConfigResponse
	1, // 6: nocloud.instance.driver.vanilla.DriverService.GetType:output_type -> nocloud.instance.driver.vanilla.GetTypeResponse
	3, // 7: nocloud.instance.driver.vanilla.DriverService.Deploy:output_type -> nocloud.instances.InstancesGroup
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_drivers_instance_vanilla_driver_proto_init() }
func file_pkg_drivers_instance_vanilla_driver_proto_init() {
	if File_pkg_drivers_instance_vanilla_driver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_drivers_instance_vanilla_driver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTypeRequest); i {
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
		file_pkg_drivers_instance_vanilla_driver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTypeResponse); i {
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
		file_pkg_drivers_instance_vanilla_driver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployRequest); i {
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
			RawDescriptor: file_pkg_drivers_instance_vanilla_driver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_drivers_instance_vanilla_driver_proto_goTypes,
		DependencyIndexes: file_pkg_drivers_instance_vanilla_driver_proto_depIdxs,
		MessageInfos:      file_pkg_drivers_instance_vanilla_driver_proto_msgTypes,
	}.Build()
	File_pkg_drivers_instance_vanilla_driver_proto = out.File
	file_pkg_drivers_instance_vanilla_driver_proto_rawDesc = nil
	file_pkg_drivers_instance_vanilla_driver_proto_goTypes = nil
	file_pkg_drivers_instance_vanilla_driver_proto_depIdxs = nil
}
