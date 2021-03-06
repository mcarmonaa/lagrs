// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	User
	ListRequest
	ListReply
	AddRequest
	AddReply
	GetRequest
	GetReply
	UpdateRequest
	UpdateReply
	RemoveRequest
	RemoveReply
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Role int32

const (
	Role_USER  Role = 0
	Role_ADMIN Role = 1
)

var Role_name = map[int32]string{
	0: "USER",
	1: "ADMIN",
}
var Role_value = map[string]int32{
	"USER":  0,
	"ADMIN": 1,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}
func (Role) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Status int32

const (
	Status_OK  Status = 0
	Status_ERR Status = 1
)

var Status_name = map[int32]string{
	0: "OK",
	1: "ERR",
}
var Status_value = map[string]int32{
	"OK":  0,
	"ERR": 1,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type User struct {
	Id   int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,3,opt,name=age" json:"age,omitempty"`
	Mail string `protobuf:"bytes,4,opt,name=mail" json:"mail,omitempty"`
	Role Role   `protobuf:"varint,5,opt,name=role,enum=user.Role" json:"role,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetMail() string {
	if m != nil {
		return m.Mail
	}
	return ""
}

func (m *User) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_USER
}

// /users GET
type ListRequest struct {
	UsersPerPage int32 `protobuf:"varint,1,opt,name=users_per_page,json=usersPerPage" json:"users_per_page,omitempty"`
	Page         int32 `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListRequest) GetUsersPerPage() int32 {
	if m != nil {
		return m.UsersPerPage
	}
	return 0
}

func (m *ListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type ListReply struct {
	Status         Status  `protobuf:"varint,1,opt,name=status,enum=user.Status" json:"status,omitempty"`
	UsersRequested int32   `protobuf:"varint,2,opt,name=users_requested,json=usersRequested" json:"users_requested,omitempty"`
	PageRequested  int32   `protobuf:"varint,3,opt,name=page_requested,json=pageRequested" json:"page_requested,omitempty"`
	UsersReplied   string  `protobuf:"bytes,4,opt,name=users_replied,json=usersReplied" json:"users_replied,omitempty"`
	PageReplied    string  `protobuf:"bytes,5,opt,name=page_replied,json=pageReplied" json:"page_replied,omitempty"`
	Pages          string  `protobuf:"bytes,6,opt,name=pages" json:"pages,omitempty"`
	Users          []*User `protobuf:"bytes,7,rep,name=users" json:"users,omitempty"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListReply) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func (m *ListReply) GetUsersRequested() int32 {
	if m != nil {
		return m.UsersRequested
	}
	return 0
}

func (m *ListReply) GetPageRequested() int32 {
	if m != nil {
		return m.PageRequested
	}
	return 0
}

func (m *ListReply) GetUsersReplied() string {
	if m != nil {
		return m.UsersReplied
	}
	return ""
}

func (m *ListReply) GetPageReplied() string {
	if m != nil {
		return m.PageReplied
	}
	return ""
}

func (m *ListReply) GetPages() string {
	if m != nil {
		return m.Pages
	}
	return ""
}

func (m *ListReply) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

// /users POST
type AddRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Mail string `protobuf:"bytes,3,opt,name=mail" json:"mail,omitempty"`
	Role Role   `protobuf:"varint,4,opt,name=role,enum=user.Role" json:"role,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AddRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddRequest) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *AddRequest) GetMail() string {
	if m != nil {
		return m.Mail
	}
	return ""
}

func (m *AddRequest) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_USER
}

type AddReply struct {
	Status Status `protobuf:"varint,1,opt,name=status,enum=user.Status" json:"status,omitempty"`
	User   *User  `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *AddReply) Reset()                    { *m = AddReply{} }
func (m *AddReply) String() string            { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()               {}
func (*AddReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AddReply) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func (m *AddReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// /users/id GET
type GetRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetReply struct {
	Status Status `protobuf:"varint,1,opt,name=status,enum=user.Status" json:"status,omitempty"`
	User   *User  `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *GetReply) Reset()                    { *m = GetReply{} }
func (m *GetReply) String() string            { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()               {}
func (*GetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetReply) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func (m *GetReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// /users/id PUT
type UpdateRequest struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UpdateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateReply struct {
	Status Status `protobuf:"varint,1,opt,name=status,enum=user.Status" json:"status,omitempty"`
	User   *User  `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *UpdateReply) Reset()                    { *m = UpdateReply{} }
func (m *UpdateReply) String() string            { return proto.CompactTextString(m) }
func (*UpdateReply) ProtoMessage()               {}
func (*UpdateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateReply) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func (m *UpdateReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// /users/id DELETE
type RemoveRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *RemoveRequest) Reset()                    { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()               {}
func (*RemoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RemoveRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RemoveReply struct {
	Status Status `protobuf:"varint,1,opt,name=status,enum=user.Status" json:"status,omitempty"`
	User   *User  `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *RemoveReply) Reset()                    { *m = RemoveReply{} }
func (m *RemoveReply) String() string            { return proto.CompactTextString(m) }
func (*RemoveReply) ProtoMessage()               {}
func (*RemoveReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RemoveReply) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_OK
}

func (m *RemoveReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*ListRequest)(nil), "user.ListRequest")
	proto.RegisterType((*ListReply)(nil), "user.ListReply")
	proto.RegisterType((*AddRequest)(nil), "user.AddRequest")
	proto.RegisterType((*AddReply)(nil), "user.AddReply")
	proto.RegisterType((*GetRequest)(nil), "user.GetRequest")
	proto.RegisterType((*GetReply)(nil), "user.GetReply")
	proto.RegisterType((*UpdateRequest)(nil), "user.UpdateRequest")
	proto.RegisterType((*UpdateReply)(nil), "user.UpdateReply")
	proto.RegisterType((*RemoveRequest)(nil), "user.RemoveRequest")
	proto.RegisterType((*RemoveReply)(nil), "user.RemoveReply")
	proto.RegisterEnum("user.Role", Role_name, Role_value)
	proto.RegisterEnum("user.Status", Status_name, Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Users service

type UsersClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/user.Users/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := grpc.Invoke(ctx, "/user.Users/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := grpc.Invoke(ctx, "/user.Users/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error) {
	out := new(UpdateReply)
	err := grpc.Invoke(ctx, "/user.Users/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error) {
	out := new(RemoveReply)
	err := grpc.Invoke(ctx, "/user.Users/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersServer interface {
	List(context.Context, *ListRequest) (*ListReply, error)
	Add(context.Context, *AddRequest) (*AddReply, error)
	Get(context.Context, *GetRequest) (*GetReply, error)
	Update(context.Context, *UpdateRequest) (*UpdateReply, error)
	Remove(context.Context, *RemoveRequest) (*RemoveReply, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Users_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Users_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Users_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Users_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Users_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x3f, 0xdb, 0xbe, 0x24, 0xae, 0x3b, 0x70, 0x30, 0x01, 0x95, 0x10, 0x8a, 0x1a, 0xe5,
	0x50, 0x50, 0xf8, 0x05, 0x41, 0x54, 0x11, 0xe2, 0x2b, 0xda, 0x28, 0x17, 0x2e, 0x95, 0x91, 0x17,
	0x64, 0xc9, 0xc1, 0xc6, 0x76, 0x90, 0x38, 0xf2, 0x73, 0xf9, 0x17, 0x68, 0x67, 0xbd, 0x89, 0x43,
	0x89, 0xc4, 0x21, 0xb7, 0xcd, 0x9b, 0x37, 0xef, 0xed, 0xce, 0x1b, 0x07, 0x58, 0x57, 0xb2, 0xbc,
	0x2e, 0xca, 0xbc, 0xce, 0xc9, 0x55, 0xe7, 0x61, 0x01, 0x77, 0x59, 0xc9, 0x92, 0x02, 0xd8, 0x69,
	0x12, 0x59, 0x03, 0x6b, 0xe4, 0x09, 0x3b, 0x4d, 0x88, 0xe0, 0x7e, 0x8b, 0x57, 0x32, 0xb2, 0x07,
	0xd6, 0xe8, 0x54, 0xf0, 0x99, 0x42, 0x38, 0xf1, 0x57, 0x19, 0x39, 0x4c, 0x52, 0x47, 0xc5, 0x5a,
	0xc5, 0x69, 0x16, 0xb9, 0x9a, 0xa5, 0xce, 0x74, 0x01, 0xb7, 0xcc, 0x33, 0x19, 0x79, 0x03, 0x6b,
	0x14, 0x4c, 0x70, 0xcd, 0x96, 0x22, 0xcf, 0xa4, 0x60, 0x7c, 0x38, 0x43, 0xe7, 0x5d, 0x5a, 0xd5,
	0x42, 0x7e, 0x5f, 0xcb, 0xaa, 0xa6, 0x4b, 0x04, 0x8a, 0x51, 0xdd, 0x16, 0xb2, 0xbc, 0x2d, 0x94,
	0xbe, 0xbe, 0x44, 0x97, 0xd1, 0xb9, 0x2c, 0xe7, 0x8d, 0x11, 0xd7, 0x6c, 0xae, 0xf1, 0x79, 0xf8,
	0xcb, 0xc6, 0xa9, 0x56, 0x2a, 0xb2, 0x9f, 0x74, 0x09, 0xbf, 0xaa, 0xe3, 0x7a, 0x5d, 0x71, 0x7f,
	0x30, 0xe9, 0x6a, 0xe3, 0x05, 0x63, 0xa2, 0xa9, 0xd1, 0x15, 0xce, 0xb4, 0x5b, 0xa9, 0xed, 0x65,
	0xd2, 0x48, 0xea, 0x4b, 0x08, 0x83, 0xd2, 0x33, 0x04, 0xca, 0xa4, 0xc5, 0xd3, 0xcf, 0xee, 0x29,
	0x74, 0x4b, 0x7b, 0x8a, 0x9e, 0xd1, 0x2b, 0xb2, 0x54, 0x26, 0xcd, 0x24, 0xba, 0x8d, 0x1a, 0x63,
	0xf4, 0x04, 0xdd, 0x46, 0x4b, 0x73, 0x3c, 0xe6, 0x74, 0xb4, 0x92, 0xa6, 0xdc, 0x87, 0xa7, 0x7e,
	0x56, 0x91, 0xcf, 0x35, 0xfd, 0x83, 0x06, 0xf0, 0x58, 0x28, 0x3a, 0x1e, 0x38, 0xa3, 0x8e, 0x99,
	0xa5, 0xca, 0x4b, 0xe8, 0xc2, 0xf0, 0x0b, 0x30, 0x4d, 0x12, 0x33, 0x4b, 0x13, 0x9a, 0x75, 0x37,
	0x34, 0xfb, 0x6e, 0x68, 0xce, 0x3f, 0x42, 0x73, 0xf7, 0x84, 0x36, 0xc7, 0x09, 0xfb, 0xfc, 0xff,
	0xa4, 0x2f, 0xc0, 0x0b, 0xc6, 0xc6, 0xbb, 0x57, 0xd7, 0x8b, 0xf7, 0x08, 0x98, 0xc9, 0xcd, 0x16,
	0xfc, 0xb5, 0x7e, 0xca, 0x8f, 0xab, 0x87, 0xf3, 0x7b, 0x8e, 0xde, 0xb2, 0x48, 0xe2, 0xda, 0x84,
	0xb7, 0x69, 0xb0, 0xf6, 0x34, 0x2c, 0xd0, 0x31, 0x0d, 0x87, 0xbb, 0xc5, 0x63, 0xf4, 0x84, 0x5c,
	0xe5, 0x3f, 0xe4, 0xbe, 0x87, 0x2f, 0xd0, 0x31, 0x84, 0x83, 0xb9, 0x8e, 0x1f, 0xc2, 0x55, 0x59,
	0xd2, 0x09, 0xdc, 0xe5, 0xe2, 0x46, 0x84, 0x47, 0x74, 0x0a, 0x6f, 0xfa, 0xfa, 0xfd, 0x9b, 0x0f,
	0xa1, 0x35, 0x7e, 0x00, 0x5f, 0xcb, 0x91, 0x0f, 0xfb, 0xe3, 0xdb, 0xf0, 0x88, 0x8e, 0xe1, 0xdc,
	0x08, 0x11, 0x5a, 0x93, 0xdf, 0x16, 0x3c, 0x25, 0x53, 0xd1, 0x18, 0xae, 0xfa, 0xd4, 0xe8, 0x5c,
	0x6b, 0xb7, 0x3e, 0xe0, 0xfe, 0x59, 0x1b, 0x52, 0x77, 0xbe, 0x82, 0x33, 0x4d, 0x12, 0x0a, 0x35,
	0xbe, 0x5d, 0xcf, 0x7e, 0xd0, 0x42, 0x1a, 0xe2, 0x4c, 0xd6, 0x86, 0xb8, 0xdd, 0x06, 0x43, 0xdc,
	0x6c, 0xc0, 0x0b, 0xf8, 0x3a, 0x0a, 0xba, 0xd7, 0xbc, 0xad, 0x9d, 0x64, 0xff, 0x7c, 0x17, 0x6c,
	0x3a, 0xf4, 0x18, 0x4d, 0xc7, 0xce, 0xd4, 0x4d, 0x47, 0x6b, 0xd2, 0xaf, 0xfc, 0x4f, 0x3c, 0xab,
	0xcf, 0x3e, 0xff, 0x3b, 0xbe, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x90, 0x77, 0xa5, 0xd3, 0x2b,
	0x05, 0x00, 0x00,
}
