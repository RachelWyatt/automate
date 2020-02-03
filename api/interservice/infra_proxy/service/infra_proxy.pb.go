// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/infra_proxy/service/infra_proxy.proto

package service

import (
	context "context"
	fmt "fmt"
	version "github.com/chef/automate/api/external/common/version"
	request "github.com/chef/automate/api/interservice/infra_proxy/request"
	response "github.com/chef/automate/api/interservice/infra_proxy/response"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("api/interservice/infra_proxy/service/infra_proxy.proto", fileDescriptor_59a1fd4cbd8b9945)
}

var fileDescriptor_59a1fd4cbd8b9945 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xd3, 0xcb, 0xd0, 0xcc, 0x24, 0x90, 0x8f, 0x96, 0xb8, 0xe4, 0x04, 0x07, 0xec, 0xad,
	0xed, 0x8a, 0x38, 0x20, 0xc4, 0x36, 0x34, 0xed, 0xc2, 0x10, 0x08, 0x0e, 0x5c, 0x90, 0xdb, 0xbd,
	0x76, 0xd6, 0xda, 0x38, 0xd8, 0x6e, 0xd5, 0x89, 0x13, 0x82, 0xcf, 0xc3, 0x67, 0x44, 0x89, 0x9d,
	0xd4, 0x5b, 0xd1, 0x96, 0xe7, 0x9d, 0x5a, 0x47, 0xfe, 0xbd, 0xff, 0x4f, 0xcf, 0x2f, 0x72, 0xc8,
	0x48, 0x96, 0x4a, 0xa8, 0xc2, 0x81, 0xb1, 0x60, 0x56, 0x6a, 0x02, 0x42, 0x15, 0x53, 0x23, 0xbf,
	0x97, 0x46, 0xaf, 0xaf, 0xc5, 0x7f, 0x9e, 0xf1, 0xd2, 0x68, 0xa7, 0xe9, 0xf3, 0xc9, 0x25, 0x4c,
	0xb9, 0x5c, 0x3a, 0xbd, 0x90, 0x0e, 0xf8, 0x85, 0x5e, 0x48, 0x55, 0xf0, 0x78, 0x5f, 0x60, 0xd9,
	0x8b, 0x2a, 0x01, 0xd6, 0x0e, 0x4c, 0x21, 0xe7, 0x62, 0xa2, 0x17, 0x0b, 0x5d, 0x88, 0x15, 0x18,
	0xab, 0x36, 0xbf, 0xbe, 0x28, 0xeb, 0xdf, 0x29, 0x63, 0xe0, 0xc7, 0x12, 0xac, 0xab, 0xa5, 0xc0,
	0xd8, 0xc0, 0x88, 0x4e, 0x8c, 0x36, 0xb3, 0x06, 0x18, 0x76, 0x02, 0x26, 0x5a, 0x5f, 0x8d, 0xb5,
	0xbe, 0x6a, 0xa8, 0xc1, 0x3d, 0x94, 0x2d, 0x75, 0x61, 0xe1, 0x96, 0xdb, 0x7e, 0x37, 0x28, 0x92,
	0x3b, 0xec, 0x46, 0xdc, 0xb2, 0xeb, 0xff, 0x7e, 0x4a, 0xc8, 0x59, 0xb5, 0xf1, 0x63, 0xb5, 0x8f,
	0x5a, 0x42, 0x4e, 0xc1, 0x7d, 0xf5, 0xbd, 0xa5, 0x43, 0x7e, 0xf3, 0xac, 0x64, 0xa9, 0xb8, 0x3f,
	0x06, 0xde, 0xb4, 0x3f, 0x6c, 0x3d, 0x2b, 0xa6, 0xfa, 0x93, 0x6f, 0x00, 0x7b, 0x89, 0xa2, 0xf2,
	0x8c, 0xfe, 0xea, 0x91, 0xbd, 0x63, 0x03, 0xd2, 0xc1, 0xe7, 0xba, 0x09, 0x74, 0xc4, 0xef, 0x9d,
	0x91, 0xd0, 0x6d, 0x1e, 0x73, 0xec, 0x55, 0x17, 0xce, 0x37, 0xe2, 0x06, 0x18, 0x1c, 0x4e, 0x60,
	0x0e, 0x29, 0x0e, 0x31, 0x87, 0x72, 0x88, 0xc1, 0xe0, 0xf0, 0xa5, 0xbc, 0x48, 0xea, 0x43, 0xcc,
	0xa1, 0x1c, 0x62, 0x30, 0xcf, 0xe8, 0xcf, 0x7a, 0x00, 0xfc, 0xd2, 0x6e, 0x0d, 0xc0, 0x1d, 0x02,
	0x1b, 0x8a, 0x1d, 0x22, 0xe2, 0x37, 0x58, 0x9e, 0xd1, 0x35, 0xd9, 0x6d, 0xd7, 0x74, 0x90, 0x90,
	0xcd, 0x86, 0x29, 0xd1, 0x79, 0x46, 0xff, 0xf4, 0xc8, 0x93, 0x76, 0x7d, 0x74, 0xfd, 0x41, 0x2e,
	0x80, 0xbe, 0x4e, 0x10, 0xf0, 0x68, 0xb2, 0xc6, 0x9a, 0xec, 0xfa, 0xb9, 0x3c, 0x37, 0x33, 0x4c,
	0x03, 0x5a, 0x08, 0x95, 0xdc, 0x52, 0x3e, 0xd9, 0x4f, 0x23, 0x32, 0xb9, 0x85, 0x50, 0xc9, 0x2d,
	0xe5, 0x93, 0xfd, 0x0c, 0x22, 0x93, 0x5b, 0x08, 0x95, 0xdc, 0x52, 0x79, 0x46, 0x0d, 0x79, 0x74,
	0x0a, 0xee, 0xdc, 0xcc, 0x2c, 0x3d, 0x40, 0x9d, 0x75, 0x85, 0xb0, 0x3e, 0xee, 0x8c, 0x2b, 0x26,
	0xcf, 0xa8, 0x26, 0x3b, 0x7e, 0x41, 0xf7, 0xb1, 0x91, 0xec, 0x00, 0x9d, 0x58, 0xbf, 0xd0, 0x7b,
	0xfe, 0x7f, 0x98, 0xea, 0x11, 0x36, 0x36, 0x8c, 0xf4, 0x03, 0xc2, 0x8f, 0x9b, 0x3b, 0x07, 0x35,
	0xd2, 0x0d, 0x84, 0x1b, 0xe9, 0x86, 0xca, 0x33, 0xfa, 0xb7, 0x47, 0x9e, 0xc5, 0xe9, 0xef, 0x56,
	0x52, 0xcd, 0xe5, 0x78, 0x0e, 0xe1, 0xf6, 0xb1, 0xf4, 0x24, 0x41, 0x67, 0xab, 0x0a, 0x7b, 0x9f,
	0xe2, 0xb7, 0x55, 0xa6, 0x7e, 0x13, 0x1e, 0x47, 0xbe, 0xb4, 0x8f, 0xb7, 0x63, 0x83, 0x04, 0x97,
	0x3c, 0x3b, 0x7a, 0xfb, 0xed, 0xcd, 0x4c, 0xb9, 0xcb, 0xe5, 0xb8, 0xba, 0xa8, 0x45, 0x55, 0x42,
	0x34, 0x25, 0x44, 0x97, 0xcf, 0xbc, 0xf1, 0x4e, 0xfd, 0x35, 0x31, 0xf8, 0x17, 0x00, 0x00, 0xff,
	0xff, 0x80, 0x7f, 0x79, 0xbd, 0x15, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InfraProxyClient is the client API for InfraProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InfraProxyClient interface {
	GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error)
	CreateServer(ctx context.Context, in *request.CreateServer, opts ...grpc.CallOption) (*response.CreateServer, error)
	DeleteServer(ctx context.Context, in *request.DeleteServer, opts ...grpc.CallOption) (*response.DeleteServer, error)
	UpdateServer(ctx context.Context, in *request.UpdateServer, opts ...grpc.CallOption) (*response.UpdateServer, error)
	GetServers(ctx context.Context, in *request.GetServers, opts ...grpc.CallOption) (*response.GetServers, error)
	GetServer(ctx context.Context, in *request.GetServer, opts ...grpc.CallOption) (*response.GetServer, error)
	GetServerByName(ctx context.Context, in *request.GetServerByName, opts ...grpc.CallOption) (*response.GetServer, error)
	CreateOrg(ctx context.Context, in *request.CreateOrg, opts ...grpc.CallOption) (*response.CreateOrg, error)
	DeleteOrg(ctx context.Context, in *request.DeleteOrg, opts ...grpc.CallOption) (*response.DeleteOrg, error)
	UpdateOrg(ctx context.Context, in *request.UpdateOrg, opts ...grpc.CallOption) (*response.UpdateOrg, error)
	GetOrgs(ctx context.Context, in *request.GetOrgs, opts ...grpc.CallOption) (*response.GetOrgs, error)
	GetOrg(ctx context.Context, in *request.GetOrg, opts ...grpc.CallOption) (*response.GetOrg, error)
	GetOrgByName(ctx context.Context, in *request.GetOrgByName, opts ...grpc.CallOption) (*response.GetOrg, error)
	GetCookbooks(ctx context.Context, in *request.Cookbooks, opts ...grpc.CallOption) (*response.Cookbooks, error)
	GetCookbooksAvailableVersions(ctx context.Context, in *request.CookbooksAvailableVersions, opts ...grpc.CallOption) (*response.CookbooksAvailableVersions, error)
	GetCookbook(ctx context.Context, in *request.Cookbook, opts ...grpc.CallOption) (*response.Cookbook, error)
}

type infraProxyClient struct {
	cc *grpc.ClientConn
}

func NewInfraProxyClient(cc *grpc.ClientConn) InfraProxyClient {
	return &infraProxyClient{cc}
}

func (c *infraProxyClient) GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error) {
	out := new(version.VersionInfo)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) CreateServer(ctx context.Context, in *request.CreateServer, opts ...grpc.CallOption) (*response.CreateServer, error) {
	out := new(response.CreateServer)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/CreateServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) DeleteServer(ctx context.Context, in *request.DeleteServer, opts ...grpc.CallOption) (*response.DeleteServer, error) {
	out := new(response.DeleteServer)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/DeleteServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) UpdateServer(ctx context.Context, in *request.UpdateServer, opts ...grpc.CallOption) (*response.UpdateServer, error) {
	out := new(response.UpdateServer)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/UpdateServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetServers(ctx context.Context, in *request.GetServers, opts ...grpc.CallOption) (*response.GetServers, error) {
	out := new(response.GetServers)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/GetServers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetServer(ctx context.Context, in *request.GetServer, opts ...grpc.CallOption) (*response.GetServer, error) {
	out := new(response.GetServer)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/GetServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetServerByName(ctx context.Context, in *request.GetServerByName, opts ...grpc.CallOption) (*response.GetServer, error) {
	out := new(response.GetServer)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/GetServerByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) CreateOrg(ctx context.Context, in *request.CreateOrg, opts ...grpc.CallOption) (*response.CreateOrg, error) {
	out := new(response.CreateOrg)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/CreateOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) DeleteOrg(ctx context.Context, in *request.DeleteOrg, opts ...grpc.CallOption) (*response.DeleteOrg, error) {
	out := new(response.DeleteOrg)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/DeleteOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) UpdateOrg(ctx context.Context, in *request.UpdateOrg, opts ...grpc.CallOption) (*response.UpdateOrg, error) {
	out := new(response.UpdateOrg)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/UpdateOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetOrgs(ctx context.Context, in *request.GetOrgs, opts ...grpc.CallOption) (*response.GetOrgs, error) {
	out := new(response.GetOrgs)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/GetOrgs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetOrg(ctx context.Context, in *request.GetOrg, opts ...grpc.CallOption) (*response.GetOrg, error) {
	out := new(response.GetOrg)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/GetOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetOrgByName(ctx context.Context, in *request.GetOrgByName, opts ...grpc.CallOption) (*response.GetOrg, error) {
	out := new(response.GetOrg)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/GetOrgByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetCookbooks(ctx context.Context, in *request.Cookbooks, opts ...grpc.CallOption) (*response.Cookbooks, error) {
	out := new(response.Cookbooks)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/GetCookbooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetCookbooksAvailableVersions(ctx context.Context, in *request.CookbooksAvailableVersions, opts ...grpc.CallOption) (*response.CookbooksAvailableVersions, error) {
	out := new(response.CookbooksAvailableVersions)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/GetCookbooksAvailableVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infraProxyClient) GetCookbook(ctx context.Context, in *request.Cookbook, opts ...grpc.CallOption) (*response.Cookbook, error) {
	out := new(response.Cookbook)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.infra_proxy.service.InfraProxy/GetCookbook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfraProxyServer is the server API for InfraProxy service.
type InfraProxyServer interface {
	GetVersion(context.Context, *version.VersionInfoRequest) (*version.VersionInfo, error)
	CreateServer(context.Context, *request.CreateServer) (*response.CreateServer, error)
	DeleteServer(context.Context, *request.DeleteServer) (*response.DeleteServer, error)
	UpdateServer(context.Context, *request.UpdateServer) (*response.UpdateServer, error)
	GetServers(context.Context, *request.GetServers) (*response.GetServers, error)
	GetServer(context.Context, *request.GetServer) (*response.GetServer, error)
	GetServerByName(context.Context, *request.GetServerByName) (*response.GetServer, error)
	CreateOrg(context.Context, *request.CreateOrg) (*response.CreateOrg, error)
	DeleteOrg(context.Context, *request.DeleteOrg) (*response.DeleteOrg, error)
	UpdateOrg(context.Context, *request.UpdateOrg) (*response.UpdateOrg, error)
	GetOrgs(context.Context, *request.GetOrgs) (*response.GetOrgs, error)
	GetOrg(context.Context, *request.GetOrg) (*response.GetOrg, error)
	GetOrgByName(context.Context, *request.GetOrgByName) (*response.GetOrg, error)
	GetCookbooks(context.Context, *request.Cookbooks) (*response.Cookbooks, error)
	GetCookbooksAvailableVersions(context.Context, *request.CookbooksAvailableVersions) (*response.CookbooksAvailableVersions, error)
	GetCookbook(context.Context, *request.Cookbook) (*response.Cookbook, error)
}

// UnimplementedInfraProxyServer can be embedded to have forward compatible implementations.
type UnimplementedInfraProxyServer struct {
}

func (*UnimplementedInfraProxyServer) GetVersion(ctx context.Context, req *version.VersionInfoRequest) (*version.VersionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (*UnimplementedInfraProxyServer) CreateServer(ctx context.Context, req *request.CreateServer) (*response.CreateServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServer not implemented")
}
func (*UnimplementedInfraProxyServer) DeleteServer(ctx context.Context, req *request.DeleteServer) (*response.DeleteServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServer not implemented")
}
func (*UnimplementedInfraProxyServer) UpdateServer(ctx context.Context, req *request.UpdateServer) (*response.UpdateServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServer not implemented")
}
func (*UnimplementedInfraProxyServer) GetServers(ctx context.Context, req *request.GetServers) (*response.GetServers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServers not implemented")
}
func (*UnimplementedInfraProxyServer) GetServer(ctx context.Context, req *request.GetServer) (*response.GetServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServer not implemented")
}
func (*UnimplementedInfraProxyServer) GetServerByName(ctx context.Context, req *request.GetServerByName) (*response.GetServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerByName not implemented")
}
func (*UnimplementedInfraProxyServer) CreateOrg(ctx context.Context, req *request.CreateOrg) (*response.CreateOrg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrg not implemented")
}
func (*UnimplementedInfraProxyServer) DeleteOrg(ctx context.Context, req *request.DeleteOrg) (*response.DeleteOrg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrg not implemented")
}
func (*UnimplementedInfraProxyServer) UpdateOrg(ctx context.Context, req *request.UpdateOrg) (*response.UpdateOrg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrg not implemented")
}
func (*UnimplementedInfraProxyServer) GetOrgs(ctx context.Context, req *request.GetOrgs) (*response.GetOrgs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgs not implemented")
}
func (*UnimplementedInfraProxyServer) GetOrg(ctx context.Context, req *request.GetOrg) (*response.GetOrg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrg not implemented")
}
func (*UnimplementedInfraProxyServer) GetOrgByName(ctx context.Context, req *request.GetOrgByName) (*response.GetOrg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgByName not implemented")
}
func (*UnimplementedInfraProxyServer) GetCookbooks(ctx context.Context, req *request.Cookbooks) (*response.Cookbooks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCookbooks not implemented")
}
func (*UnimplementedInfraProxyServer) GetCookbooksAvailableVersions(ctx context.Context, req *request.CookbooksAvailableVersions) (*response.CookbooksAvailableVersions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCookbooksAvailableVersions not implemented")
}
func (*UnimplementedInfraProxyServer) GetCookbook(ctx context.Context, req *request.Cookbook) (*response.Cookbook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCookbook not implemented")
}

func RegisterInfraProxyServer(s *grpc.Server, srv InfraProxyServer) {
	s.RegisterService(&_InfraProxy_serviceDesc, srv)
}

func _InfraProxy_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(version.VersionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetVersion(ctx, req.(*version.VersionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_CreateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).CreateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/CreateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).CreateServer(ctx, req.(*request.CreateServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_DeleteServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).DeleteServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/DeleteServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).DeleteServer(ctx, req.(*request.DeleteServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_UpdateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).UpdateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/UpdateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).UpdateServer(ctx, req.(*request.UpdateServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetServers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/GetServers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetServers(ctx, req.(*request.GetServers))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/GetServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetServer(ctx, req.(*request.GetServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetServerByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetServerByName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetServerByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/GetServerByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetServerByName(ctx, req.(*request.GetServerByName))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_CreateOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateOrg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).CreateOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/CreateOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).CreateOrg(ctx, req.(*request.CreateOrg))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_DeleteOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteOrg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).DeleteOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/DeleteOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).DeleteOrg(ctx, req.(*request.DeleteOrg))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_UpdateOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateOrg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).UpdateOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/UpdateOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).UpdateOrg(ctx, req.(*request.UpdateOrg))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetOrgs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetOrgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetOrgs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/GetOrgs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetOrgs(ctx, req.(*request.GetOrgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetOrg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/GetOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetOrg(ctx, req.(*request.GetOrg))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetOrgByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetOrgByName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetOrgByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/GetOrgByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetOrgByName(ctx, req.(*request.GetOrgByName))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetCookbooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Cookbooks)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetCookbooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/GetCookbooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetCookbooks(ctx, req.(*request.Cookbooks))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetCookbooksAvailableVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CookbooksAvailableVersions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetCookbooksAvailableVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/GetCookbooksAvailableVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetCookbooksAvailableVersions(ctx, req.(*request.CookbooksAvailableVersions))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfraProxy_GetCookbook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Cookbook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraProxyServer).GetCookbook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.infra_proxy.service.InfraProxy/GetCookbook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraProxyServer).GetCookbook(ctx, req.(*request.Cookbook))
	}
	return interceptor(ctx, in, info, handler)
}

var _InfraProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.infra_proxy.service.InfraProxy",
	HandlerType: (*InfraProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _InfraProxy_GetVersion_Handler,
		},
		{
			MethodName: "CreateServer",
			Handler:    _InfraProxy_CreateServer_Handler,
		},
		{
			MethodName: "DeleteServer",
			Handler:    _InfraProxy_DeleteServer_Handler,
		},
		{
			MethodName: "UpdateServer",
			Handler:    _InfraProxy_UpdateServer_Handler,
		},
		{
			MethodName: "GetServers",
			Handler:    _InfraProxy_GetServers_Handler,
		},
		{
			MethodName: "GetServer",
			Handler:    _InfraProxy_GetServer_Handler,
		},
		{
			MethodName: "GetServerByName",
			Handler:    _InfraProxy_GetServerByName_Handler,
		},
		{
			MethodName: "CreateOrg",
			Handler:    _InfraProxy_CreateOrg_Handler,
		},
		{
			MethodName: "DeleteOrg",
			Handler:    _InfraProxy_DeleteOrg_Handler,
		},
		{
			MethodName: "UpdateOrg",
			Handler:    _InfraProxy_UpdateOrg_Handler,
		},
		{
			MethodName: "GetOrgs",
			Handler:    _InfraProxy_GetOrgs_Handler,
		},
		{
			MethodName: "GetOrg",
			Handler:    _InfraProxy_GetOrg_Handler,
		},
		{
			MethodName: "GetOrgByName",
			Handler:    _InfraProxy_GetOrgByName_Handler,
		},
		{
			MethodName: "GetCookbooks",
			Handler:    _InfraProxy_GetCookbooks_Handler,
		},
		{
			MethodName: "GetCookbooksAvailableVersions",
			Handler:    _InfraProxy_GetCookbooksAvailableVersions_Handler,
		},
		{
			MethodName: "GetCookbook",
			Handler:    _InfraProxy_GetCookbook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interservice/infra_proxy/service/infra_proxy.proto",
}