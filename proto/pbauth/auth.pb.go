// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: proto/pbauth/auth.proto

package pbauth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_proto_pbauth_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbauth_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_proto_pbauth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        int64                  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_proto_pbauth_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbauth_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_proto_pbauth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RegisterResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RegisterResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type AdminRegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdminRegisterRequest) Reset() {
	*x = AdminRegisterRequest{}
	mi := &file_proto_pbauth_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdminRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRegisterRequest) ProtoMessage() {}

func (x *AdminRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbauth_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRegisterRequest.ProtoReflect.Descriptor instead.
func (*AdminRegisterRequest) Descriptor() ([]byte, []int) {
	return file_proto_pbauth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *AdminRegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AdminRegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AdminRegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdminRegisterResponse) Reset() {
	*x = AdminRegisterResponse{}
	mi := &file_proto_pbauth_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdminRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRegisterResponse) ProtoMessage() {}

func (x *AdminRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbauth_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRegisterResponse.ProtoReflect.Descriptor instead.
func (*AdminRegisterResponse) Descriptor() ([]byte, []int) {
	return file_proto_pbauth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *AdminRegisterResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AdminRegisterResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_proto_pbauth_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbauth_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_pbauth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AdminLoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdminLoginRequest) Reset() {
	*x = AdminLoginRequest{}
	mi := &file_proto_pbauth_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdminLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminLoginRequest) ProtoMessage() {}

func (x *AdminLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbauth_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminLoginRequest.ProtoReflect.Descriptor instead.
func (*AdminLoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_pbauth_auth_proto_rawDescGZIP(), []int{5}
}

func (x *AdminLoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AdminLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Status        int64                  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_proto_pbauth_auth_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbauth_auth_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_pbauth_auth_proto_rawDescGZIP(), []int{6}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *LoginResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ValidateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	mi := &file_proto_pbauth_auth_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbauth_auth_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_proto_pbauth_auth_proto_rawDescGZIP(), []int{7}
}

func (x *ValidateRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ValidateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int64                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	UserID        int64                  `protobuf:"varint,3,opt,name=userID,proto3" json:"userID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateResponse) Reset() {
	*x = ValidateResponse{}
	mi := &file_proto_pbauth_auth_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResponse) ProtoMessage() {}

func (x *ValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbauth_auth_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateResponse.ProtoReflect.Descriptor instead.
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return file_proto_pbauth_auth_proto_rawDescGZIP(), []int{8}
}

func (x *ValidateResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ValidateResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ValidateResponse) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

var File_proto_pbauth_auth_proto protoreflect.FileDescriptor

const file_proto_pbauth_auth_proto_rawDesc = "" +
	"\n" +
	"\x17proto/pbauth/auth.proto\x12\x06pbauth\"_\n" +
	"\x0fRegisterRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\"P\n" +
	"\x10RegisterResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x16\n" +
	"\x06status\x18\x02 \x01(\x03R\x06status\x12\x14\n" +
	"\x05error\x18\x03 \x01(\tR\x05error\"N\n" +
	"\x14AdminRegisterRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"O\n" +
	"\x15AdminRegisterResponse\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"@\n" +
	"\fLoginRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"K\n" +
	"\x11AdminLoginRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"S\n" +
	"\rLoginResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12\x16\n" +
	"\x06status\x18\x02 \x01(\x03R\x06status\x12\x14\n" +
	"\x05error\x18\x03 \x01(\tR\x05error\"'\n" +
	"\x0fValidateRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"X\n" +
	"\x10ValidateResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\x03R\x06status\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error\x12\x16\n" +
	"\x06userID\x18\x03 \x01(\x03R\x06userID2\xca\x02\n" +
	"\vAuthService\x12=\n" +
	"\bRegister\x12\x17.pbauth.RegisterRequest\x1a\x18.pbauth.RegisterResponse\x12G\n" +
	"\rAdminRegister\x12\x1c.pbauth.AdminRegisterRequest\x1a\x18.pbauth.RegisterResponse\x124\n" +
	"\x05Login\x12\x14.pbauth.LoginRequest\x1a\x15.pbauth.LoginResponse\x12>\n" +
	"\n" +
	"AdminLogin\x12\x19.pbauth.AdminLoginRequest\x1a\x15.pbauth.LoginResponse\x12=\n" +
	"\bValidate\x12\x17.pbauth.ValidateRequest\x1a\x18.pbauth.ValidateResponseB;Z9github.com/shoksin/marketplace-protos/proto/pbauth;pbauthb\x06proto3"

var (
	file_proto_pbauth_auth_proto_rawDescOnce sync.Once
	file_proto_pbauth_auth_proto_rawDescData []byte
)

func file_proto_pbauth_auth_proto_rawDescGZIP() []byte {
	file_proto_pbauth_auth_proto_rawDescOnce.Do(func() {
		file_proto_pbauth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_pbauth_auth_proto_rawDesc), len(file_proto_pbauth_auth_proto_rawDesc)))
	})
	return file_proto_pbauth_auth_proto_rawDescData
}

var file_proto_pbauth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_pbauth_auth_proto_goTypes = []any{
	(*RegisterRequest)(nil),       // 0: pbauth.RegisterRequest
	(*RegisterResponse)(nil),      // 1: pbauth.RegisterResponse
	(*AdminRegisterRequest)(nil),  // 2: pbauth.AdminRegisterRequest
	(*AdminRegisterResponse)(nil), // 3: pbauth.AdminRegisterResponse
	(*LoginRequest)(nil),          // 4: pbauth.LoginRequest
	(*AdminLoginRequest)(nil),     // 5: pbauth.AdminLoginRequest
	(*LoginResponse)(nil),         // 6: pbauth.LoginResponse
	(*ValidateRequest)(nil),       // 7: pbauth.ValidateRequest
	(*ValidateResponse)(nil),      // 8: pbauth.ValidateResponse
}
var file_proto_pbauth_auth_proto_depIdxs = []int32{
	0, // 0: pbauth.AuthService.Register:input_type -> pbauth.RegisterRequest
	2, // 1: pbauth.AuthService.AdminRegister:input_type -> pbauth.AdminRegisterRequest
	4, // 2: pbauth.AuthService.Login:input_type -> pbauth.LoginRequest
	5, // 3: pbauth.AuthService.AdminLogin:input_type -> pbauth.AdminLoginRequest
	7, // 4: pbauth.AuthService.Validate:input_type -> pbauth.ValidateRequest
	1, // 5: pbauth.AuthService.Register:output_type -> pbauth.RegisterResponse
	1, // 6: pbauth.AuthService.AdminRegister:output_type -> pbauth.RegisterResponse
	6, // 7: pbauth.AuthService.Login:output_type -> pbauth.LoginResponse
	6, // 8: pbauth.AuthService.AdminLogin:output_type -> pbauth.LoginResponse
	8, // 9: pbauth.AuthService.Validate:output_type -> pbauth.ValidateResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_pbauth_auth_proto_init() }
func file_proto_pbauth_auth_proto_init() {
	if File_proto_pbauth_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_pbauth_auth_proto_rawDesc), len(file_proto_pbauth_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pbauth_auth_proto_goTypes,
		DependencyIndexes: file_proto_pbauth_auth_proto_depIdxs,
		MessageInfos:      file_proto_pbauth_auth_proto_msgTypes,
	}.Build()
	File_proto_pbauth_auth_proto = out.File
	file_proto_pbauth_auth_proto_goTypes = nil
	file_proto_pbauth_auth_proto_depIdxs = nil
}
