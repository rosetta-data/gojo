// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: v1/uspb/service_user.proto

package uspbv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_v1_uspb_service_user_proto protoreflect.FileDescriptor

var file_v1_uspb_service_user_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x31,
	0x2e, 0x75, 0x73, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x70, 0x62, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70,
	0x63, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x72,
	0x65, 0x6e, 0x65, 0x77, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67,
	0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcd, 0x07, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x70, 0x62, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3d, 0x92, 0x41, 0x23, 0x1a, 0x21, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61,
	0x20, 0x6e, 0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a,
	0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0xa2, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b,
	0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x76, 0x31,
	0x2e, 0x75, 0x73, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0x92, 0x41, 0x41, 0x1a, 0x3f,
	0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x67,
	0x65, 0x74, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20,
	0x26, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0xb3, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x70, 0x62, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x68, 0x92, 0x41, 0x4e, 0x1a, 0x4c, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x6f, 0x72, 0x20, 0x66, 0x75, 0x6c,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x6f, 0x72, 0x20, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x20, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x32, 0x0c, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x0b,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x2e, 0x76, 0x31,
	0x2e, 0x75, 0x73, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x31, 0x2e,
	0x75, 0x73, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x43, 0x92, 0x41, 0x2d, 0x1a,
	0x2b, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f,
	0x20, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0xa6, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12,
	0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6e, 0x65,
	0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6e, 0x65, 0x77,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58,
	0x92, 0x41, 0x3b, 0x1a, 0x39, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27,
	0x73, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x26,
	0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x22, 0x12, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x6e, 0x65,
	0x77, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x9a, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x76, 0x31,
	0x2e, 0x75, 0x73, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x76, 0x31, 0x2e, 0x75, 0x73, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x43, 0x92, 0x41, 0x22, 0x1a, 0x20, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20,
	0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x7d, 0x42, 0x81, 0x01, 0x92, 0x41, 0x49, 0x12, 0x47, 0x0a, 0x08, 0x47,
	0x6f, 0x6a, 0x6f, 0x20, 0x41, 0x50, 0x49, 0x22, 0x36, 0x0a, 0x09, 0x44, 0x6a, 0x20, 0x59, 0x61,
	0x63, 0x69, 0x6e, 0x65, 0x12, 0x29, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69,
	0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x32,
	0x03, 0x34, 0x2e, 0x30, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74,
	0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x70, 0x62, 0x3b, 0x75, 0x73, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_v1_uspb_service_user_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),      // 0: v1.uspbv1.CreateUserRequest
	(*LoginUserRequest)(nil),       // 1: v1.uspbv1.LoginUserRequest
	(*UpdateUserRequest)(nil),      // 2: v1.uspbv1.UpdateUserRequest
	(*VerifyEmailRequest)(nil),     // 3: v1.uspbv1.VerifyEmailRequest
	(*RenewTokensRequest)(nil),     // 4: v1.uspbv1.RenewTokensRequest
	(*GetUserDevicesRequest)(nil),  // 5: v1.uspbv1.GetUserDevicesRequest
	(*CreateUserResponse)(nil),     // 6: v1.uspbv1.CreateUserResponse
	(*LoginUserResponse)(nil),      // 7: v1.uspbv1.LoginUserResponse
	(*UpdateUserResponse)(nil),     // 8: v1.uspbv1.UpdateUserResponse
	(*VerifyEmailResponse)(nil),    // 9: v1.uspbv1.VerifyEmailResponse
	(*RenewTokensResponse)(nil),    // 10: v1.uspbv1.RenewTokensResponse
	(*GetUserDevicesResponse)(nil), // 11: v1.uspbv1.GetUserDevicesResponse
}
var file_v1_uspb_service_user_proto_depIdxs = []int32{
	0,  // 0: v1.uspbv1.UserService.CreateUser:input_type -> v1.uspbv1.CreateUserRequest
	1,  // 1: v1.uspbv1.UserService.LoginUser:input_type -> v1.uspbv1.LoginUserRequest
	2,  // 2: v1.uspbv1.UserService.UpdateUser:input_type -> v1.uspbv1.UpdateUserRequest
	3,  // 3: v1.uspbv1.UserService.VerifyEmail:input_type -> v1.uspbv1.VerifyEmailRequest
	4,  // 4: v1.uspbv1.UserService.RenewTokens:input_type -> v1.uspbv1.RenewTokensRequest
	5,  // 5: v1.uspbv1.UserService.GetUserDevices:input_type -> v1.uspbv1.GetUserDevicesRequest
	6,  // 6: v1.uspbv1.UserService.CreateUser:output_type -> v1.uspbv1.CreateUserResponse
	7,  // 7: v1.uspbv1.UserService.LoginUser:output_type -> v1.uspbv1.LoginUserResponse
	8,  // 8: v1.uspbv1.UserService.UpdateUser:output_type -> v1.uspbv1.UpdateUserResponse
	9,  // 9: v1.uspbv1.UserService.VerifyEmail:output_type -> v1.uspbv1.VerifyEmailResponse
	10, // 10: v1.uspbv1.UserService.RenewTokens:output_type -> v1.uspbv1.RenewTokensResponse
	11, // 11: v1.uspbv1.UserService.GetUserDevices:output_type -> v1.uspbv1.GetUserDevicesResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_v1_uspb_service_user_proto_init() }
func file_v1_uspb_service_user_proto_init() {
	if File_v1_uspb_service_user_proto != nil {
		return
	}
	file_v1_uspb_rpc_create_user_proto_init()
	file_v1_uspb_rpc_login_user_proto_init()
	file_v1_uspb_rpc_update_user_proto_init()
	file_v1_uspb_rpc_verify_email_proto_init()
	file_v1_uspb_rpc_renew_tokens_proto_init()
	file_v1_uspb_rpc_get_user_devices_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_uspb_service_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_uspb_service_user_proto_goTypes,
		DependencyIndexes: file_v1_uspb_service_user_proto_depIdxs,
	}.Build()
	File_v1_uspb_service_user_proto = out.File
	file_v1_uspb_service_user_proto_rawDesc = nil
	file_v1_uspb_service_user_proto_goTypes = nil
	file_v1_uspb_service_user_proto_depIdxs = nil
}
