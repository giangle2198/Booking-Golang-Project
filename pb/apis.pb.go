// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: apis.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("apis.proto", fileDescriptor_b480fce41873d8e6) }

var fileDescriptor_b480fce41873d8e6 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x69, 0x87, 0x28, 0x32, 0x74, 0x2a, 0x1e, 0xea, 0xba, 0x74, 0x93, 0xa2, 0x1d, 0x07,
	0x4d, 0x58, 0x27, 0x38, 0x0c, 0x09, 0x29, 0xeb, 0xa4, 0xaa, 0x02, 0xa1, 0x0a, 0x84, 0x10, 0xbb,
	0x39, 0xe9, 0xa7, 0x60, 0x91, 0xd9, 0x5e, 0xec, 0x14, 0x8d, 0x0b, 0x12, 0x77, 0x2e, 0xe1, 0x01,
	0x78, 0x28, 0x5e, 0x81, 0x03, 0x8f, 0x81, 0xe2, 0xd8, 0xdb, 0x92, 0x66, 0xa7, 0x9e, 0xaa, 0xfe,
	0xbf, 0x7f, 0xbe, 0xdf, 0x2f, 0x56, 0x12, 0x84, 0x88, 0xa0, 0xd2, 0x13, 0x29, 0x57, 0x1c, 0xb7,
	0x45, 0xe8, 0xec, 0xc6, 0x9c, 0xc7, 0x09, 0xf8, 0x44, 0x50, 0x9f, 0x30, 0xc6, 0x15, 0x51, 0x94,
	0x33, 0xd3, 0x70, 0x9e, 0xea, 0x9f, 0x68, 0x14, 0x03, 0x1b, 0xc9, 0xaf, 0x24, 0x8e, 0x21, 0xf5,
	0xb9, 0xd0, 0x8d, 0x86, 0x76, 0xef, 0x9c, 0x30, 0x12, 0xc3, 0x39, 0x30, 0x65, 0x92, 0xae, 0x48,
	0xf9, 0x22, 0x8b, 0xcc, 0xdf, 0xf1, 0xbf, 0x0e, 0xda, 0x08, 0xe6, 0x33, 0x9c, 0xa2, 0x07, 0x53,
	0x50, 0xef, 0x21, 0x5d, 0xd2, 0x08, 0x24, 0xee, 0x7b, 0x22, 0xf4, 0x6e, 0x04, 0xef, 0xe0, 0x22,
	0x03, 0xa9, 0x9c, 0xed, 0x95, 0x5c, 0x0a, 0xce, 0x24, 0xec, 0x8f, 0xf3, 0x60, 0x1b, 0x75, 0x66,
	0xec, 0x22, 0xa3, 0xe9, 0x25, 0x7e, 0x38, 0x05, 0xe5, 0xda, 0xce, 0x8f, 0x3f, 0x7f, 0x7f, 0xb5,
	0x31, 0xee, 0xf9, 0x61, 0x24, 0xfd, 0xe5, 0xa1, 0x2f, 0x2d, 0x64, 0x89, 0x8a, 0xde, 0x84, 0x67,
	0x4c, 0xa5, 0x14, 0x24, 0xb6, 0xcb, 0xaf, 0x12, 0x4b, 0x1d, 0xac, 0x0e, 0x0c, 0xf6, 0x68, 0x15,
	0x3b, 0xe1, 0x99, 0xee, 0x68, 0xec, 0x16, 0x7e, 0x64, 0xb1, 0xd1, 0x15, 0xe7, 0x3b, 0xda, 0x9c,
	0x82, 0x7a, 0x43, 0xa5, 0x9a, 0x97, 0x67, 0x81, 0x77, 0x0c, 0xe0, 0x46, 0x66, 0xd9, 0x4e, 0xd3,
	0xc8, 0xd0, 0x5f, 0xe4, 0xc1, 0x10, 0x75, 0xec, 0x92, 0x5e, 0x41, 0x4f, 0xa8, 0x54, 0xae, 0x39,
	0xe2, 0xfa, 0x8d, 0x9b, 0x58, 0xe2, 0x9f, 0x2d, 0x54, 0x94, 0xcd, 0x85, 0xa7, 0xa0, 0x08, 0x4d,
	0xf0, 0xd0, 0x80, 0x2a, 0xa9, 0xb5, 0xd8, 0x6d, 0x1e, 0x1a, 0x8f, 0x57, 0x79, 0xb0, 0x77, 0xed,
	0x81, 0x0b, 0x0f, 0xc3, 0x72, 0x17, 0xba, 0xa9, 0x4d, 0x06, 0xb8, 0x5f, 0x33, 0xf1, 0xcb, 0x29,
	0xfe, 0xdd, 0x42, 0xdd, 0x49, 0x0a, 0x44, 0x81, 0xdd, 0xa1, 0x4f, 0xbc, 0x12, 0x59, 0x93, 0x9d,
	0x86, 0x89, 0xd1, 0xf8, 0x94, 0x07, 0xcf, 0xaf, 0x35, 0x36, 0xcb, 0x92, 0x35, 0x09, 0x1d, 0x34,
	0x40, 0x8f, 0x83, 0xf9, 0xec, 0x35, 0x5c, 0x06, 0x99, 0xfa, 0x0c, 0x4c, 0xd1, 0x48, 0x3f, 0xab,
	0xf8, 0x8e, 0xd6, 0x1b, 0xee, 0xaf, 0xe8, 0x45, 0x7a, 0xc3, 0x71, 0xeb, 0x40, 0x1b, 0x7e, 0x10,
	0x8b, 0xba, 0x61, 0x25, 0xaa, 0x18, 0xd6, 0x26, 0xcd, 0x86, 0x65, 0x69, 0x1d, 0xc3, 0x4c, 0x6f,
	0xb0, 0x86, 0xa7, 0x90, 0x40, 0xcd, 0xb0, 0x12, 0x55, 0x0c, 0x6b, 0x93, 0x66, 0xc3, 0xb2, 0xb4,
	0x8e, 0xe1, 0x42, 0x6f, 0x38, 0x6e, 0x1d, 0x9c, 0x7c, 0xcc, 0x83, 0xb7, 0x78, 0x0f, 0xf5, 0x4f,
	0x38, 0xff, 0x42, 0x59, 0xec, 0x4e, 0x79, 0x42, 0x58, 0x6c, 0x5f, 0xd3, 0xf1, 0xc6, 0xa1, 0xf7,
	0xec, 0x6c, 0x84, 0x9e, 0xdc, 0x02, 0xd8, 0xba, 0xdf, 0x76, 0xba, 0x45, 0xc6, 0x53, 0xfa, 0x4d,
	0x47, 0x6e, 0xfb, 0xec, 0xae, 0xf7, 0x52, 0x84, 0xe1, 0x3d, 0xfd, 0x29, 0x39, 0xfa, 0x1f, 0x00,
	0x00, 0xff, 0xff, 0xed, 0xf5, 0x8d, 0xe3, 0xc9, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error)
	GetCountries(ctx context.Context, in *GetCountriesRequest, opts ...grpc.CallOption) (*GetCountriesResponse, error)
	GetListProduct(ctx context.Context, in *GetListProductRequest, opts ...grpc.CallOption) (*GetListProductResponse, error)
	GetProductDetail(ctx context.Context, in *GetProductDetailRequest, opts ...grpc.CallOption) (*GetProductDetailResponse, error)
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error) {
	out := new(GetServicesResponse)
	err := c.cc.Invoke(ctx, "/pb.API/GetServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetCountries(ctx context.Context, in *GetCountriesRequest, opts ...grpc.CallOption) (*GetCountriesResponse, error) {
	out := new(GetCountriesResponse)
	err := c.cc.Invoke(ctx, "/pb.API/GetCountries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetListProduct(ctx context.Context, in *GetListProductRequest, opts ...grpc.CallOption) (*GetListProductResponse, error) {
	out := new(GetListProductResponse)
	err := c.cc.Invoke(ctx, "/pb.API/GetListProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetProductDetail(ctx context.Context, in *GetProductDetailRequest, opts ...grpc.CallOption) (*GetProductDetailResponse, error) {
	out := new(GetProductDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.API/GetProductDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, "/pb.API/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error) {
	out := new(UpdateProductResponse)
	err := c.cc.Invoke(ctx, "/pb.API/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, "/pb.API/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error)
	GetCountries(context.Context, *GetCountriesRequest) (*GetCountriesResponse, error)
	GetListProduct(context.Context, *GetListProductRequest) (*GetListProductResponse, error)
	GetProductDetail(context.Context, *GetProductDetailRequest) (*GetProductDetailResponse, error)
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
}

// UnimplementedAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (*UnimplementedAPIServer) GetServices(ctx context.Context, req *GetServicesRequest) (*GetServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (*UnimplementedAPIServer) GetCountries(ctx context.Context, req *GetCountriesRequest) (*GetCountriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountries not implemented")
}
func (*UnimplementedAPIServer) GetListProduct(ctx context.Context, req *GetListProductRequest) (*GetListProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListProduct not implemented")
}
func (*UnimplementedAPIServer) GetProductDetail(ctx context.Context, req *GetProductDetailRequest) (*GetProductDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductDetail not implemented")
}
func (*UnimplementedAPIServer) CreateProduct(ctx context.Context, req *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (*UnimplementedAPIServer) UpdateProduct(ctx context.Context, req *UpdateProductRequest) (*UpdateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (*UnimplementedAPIServer) DeleteProduct(ctx context.Context, req *DeleteProductRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/GetServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetServices(ctx, req.(*GetServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetCountries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCountriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetCountries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/GetCountries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetCountries(ctx, req.(*GetCountriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetListProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetListProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/GetListProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetListProduct(ctx, req.(*GetListProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetProductDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetProductDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/GetProductDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetProductDetail(ctx, req.(*GetProductDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServices",
			Handler:    _API_GetServices_Handler,
		},
		{
			MethodName: "GetCountries",
			Handler:    _API_GetCountries_Handler,
		},
		{
			MethodName: "GetListProduct",
			Handler:    _API_GetListProduct_Handler,
		},
		{
			MethodName: "GetProductDetail",
			Handler:    _API_GetProductDetail_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _API_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _API_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _API_DeleteProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apis.proto",
}