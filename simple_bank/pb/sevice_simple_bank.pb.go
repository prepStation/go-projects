// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: sevice_simple_bank.proto

package pb

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

var File_sevice_simple_bank_proto protoreflect.FileDescriptor

var file_sevice_simple_bank_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f,
	0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x70,
	0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x72, 0x70, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd5, 0x04, 0x0a, 0x0a, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x86, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0x92, 0x41, 0x2c, 0x12, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1d, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20,
	0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x89, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c,
	0x92, 0x41, 0x2f, 0x12, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x1a, 0x20, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x61, 0x70, 0x69, 0x20, 0x74,
	0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x69, 0x6e,
	0x66, 0x6f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x32, 0x0f, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x90, 0x01, 0x0a,
	0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x92, 0x41, 0x3a, 0x12, 0x0a, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x2c, 0x75, 0x73, 0x65, 0x20, 0x74, 0x68,
	0x69, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x20,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22,
	0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x9e, 0x01, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x16, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x5e, 0x92, 0x41, 0x43, 0x12, 0x14, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x2b, 0x75, 0x73, 0x65,
	0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x61, 0x70, 0x69, 0x20, 0x74, 0x6f, 0x20, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x42, 0x8d, 0x01, 0x92, 0x41, 0x63, 0x12, 0x61, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x42, 0x61, 0x6e, 0x6b, 0x41, 0x50, 0x49, 0x22, 0x4b, 0x0a, 0x07, 0x41, 0x73, 0x73, 0x69, 0x6d,
	0x20, 0x73, 0x12, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x1a, 0x10, 0x6e, 0x6f, 0x6e, 0x65, 0x40, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x32, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x65, 0x70, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_sevice_simple_bank_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),   // 0: pb.CreateUserRequest
	(*UpdateUserRequest)(nil),   // 1: pb.UpdateUserRequest
	(*LoginUserRequest)(nil),    // 2: pb.LoginUserRequest
	(*VerifyEmailRequest)(nil),  // 3: pb.VerifyEmailRequest
	(*CreateUserResponse)(nil),  // 4: pb.CreateUserResponse
	(*UpdateUserResponse)(nil),  // 5: pb.UpdateUserResponse
	(*LoginUserResponse)(nil),   // 6: pb.LoginUserResponse
	(*VerifyEmailResponse)(nil), // 7: pb.VerifyEmailResponse
}
var file_sevice_simple_bank_proto_depIdxs = []int32{
	0, // 0: pb.SimpleBank.CreateUser:input_type -> pb.CreateUserRequest
	1, // 1: pb.SimpleBank.UpdateUser:input_type -> pb.UpdateUserRequest
	2, // 2: pb.SimpleBank.LoginUser:input_type -> pb.LoginUserRequest
	3, // 3: pb.SimpleBank.VerifyEmail:input_type -> pb.VerifyEmailRequest
	4, // 4: pb.SimpleBank.CreateUser:output_type -> pb.CreateUserResponse
	5, // 5: pb.SimpleBank.UpdateUser:output_type -> pb.UpdateUserResponse
	6, // 6: pb.SimpleBank.LoginUser:output_type -> pb.LoginUserResponse
	7, // 7: pb.SimpleBank.VerifyEmail:output_type -> pb.VerifyEmailResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sevice_simple_bank_proto_init() }
func file_sevice_simple_bank_proto_init() {
	if File_sevice_simple_bank_proto != nil {
		return
	}
	file_rpc_create_user_proto_init()
	file_rpc_login_user_proto_init()
	file_rpc_update_user_proto_init()
	file_rpc_verify_email_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sevice_simple_bank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sevice_simple_bank_proto_goTypes,
		DependencyIndexes: file_sevice_simple_bank_proto_depIdxs,
	}.Build()
	File_sevice_simple_bank_proto = out.File
	file_sevice_simple_bank_proto_rawDesc = nil
	file_sevice_simple_bank_proto_goTypes = nil
	file_sevice_simple_bank_proto_depIdxs = nil
}
