// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: productService.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ProductService service

type ProductService interface {
	//商品服务
	CreateProduct(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductDetailResponse, error)
	GetProductsList(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductsListResponse, error)
	GetProduct(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductDetailResponse, error)
	UpdateProduct(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductDetailResponse, error)
	DeleteProduct(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductDetailResponse, error)
	//商品图片服务
	CreateProductImg(ctx context.Context, in *ProductImgRequest, opts ...client.CallOption) (*ProductImgDetailResponse, error)
	GetProductImgsList(ctx context.Context, in *ProductImgRequest, opts ...client.CallOption) (*ProductImgsListResponse, error)
	UpdateProductImg(ctx context.Context, in *ProductImgRequest, opts ...client.CallOption) (*ProductImgDetailResponse, error)
	DeleteProductImg(ctx context.Context, in *ProductImgRequest, opts ...client.CallOption) (*ProductImgDetailResponse, error)
}

type productService struct {
	c    client.Client
	name string
}

func NewProductService(name string, c client.Client) ProductService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "services"
	}
	return &productService{
		c:    c,
		name: name,
	}
}

func (c *productService) CreateProduct(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductDetailResponse, error) {
	req := c.c.NewRequest(c.name, "ProductService.CreateProduct", in)
	out := new(ProductDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) GetProductsList(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductsListResponse, error) {
	req := c.c.NewRequest(c.name, "ProductService.GetProductsList", in)
	out := new(ProductsListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) GetProduct(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductDetailResponse, error) {
	req := c.c.NewRequest(c.name, "ProductService.GetProduct", in)
	out := new(ProductDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) UpdateProduct(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductDetailResponse, error) {
	req := c.c.NewRequest(c.name, "ProductService.UpdateProduct", in)
	out := new(ProductDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) DeleteProduct(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductDetailResponse, error) {
	req := c.c.NewRequest(c.name, "ProductService.DeleteProduct", in)
	out := new(ProductDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) CreateProductImg(ctx context.Context, in *ProductImgRequest, opts ...client.CallOption) (*ProductImgDetailResponse, error) {
	req := c.c.NewRequest(c.name, "ProductService.CreateProductImg", in)
	out := new(ProductImgDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) GetProductImgsList(ctx context.Context, in *ProductImgRequest, opts ...client.CallOption) (*ProductImgsListResponse, error) {
	req := c.c.NewRequest(c.name, "ProductService.GetProductImgsList", in)
	out := new(ProductImgsListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) UpdateProductImg(ctx context.Context, in *ProductImgRequest, opts ...client.CallOption) (*ProductImgDetailResponse, error) {
	req := c.c.NewRequest(c.name, "ProductService.UpdateProductImg", in)
	out := new(ProductImgDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) DeleteProductImg(ctx context.Context, in *ProductImgRequest, opts ...client.CallOption) (*ProductImgDetailResponse, error) {
	req := c.c.NewRequest(c.name, "ProductService.DeleteProductImg", in)
	out := new(ProductImgDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductService service

type ProductServiceHandler interface {
	//商品服务
	CreateProduct(context.Context, *ProductRequest, *ProductDetailResponse) error
	GetProductsList(context.Context, *ProductRequest, *ProductsListResponse) error
	GetProduct(context.Context, *ProductRequest, *ProductDetailResponse) error
	UpdateProduct(context.Context, *ProductRequest, *ProductDetailResponse) error
	DeleteProduct(context.Context, *ProductRequest, *ProductDetailResponse) error
	//商品图片服务
	CreateProductImg(context.Context, *ProductImgRequest, *ProductImgDetailResponse) error
	GetProductImgsList(context.Context, *ProductImgRequest, *ProductImgsListResponse) error
	UpdateProductImg(context.Context, *ProductImgRequest, *ProductImgDetailResponse) error
	DeleteProductImg(context.Context, *ProductImgRequest, *ProductImgDetailResponse) error
}

func RegisterProductServiceHandler(s server.Server, hdlr ProductServiceHandler, opts ...server.HandlerOption) error {
	type productService interface {
		CreateProduct(ctx context.Context, in *ProductRequest, out *ProductDetailResponse) error
		GetProductsList(ctx context.Context, in *ProductRequest, out *ProductsListResponse) error
		GetProduct(ctx context.Context, in *ProductRequest, out *ProductDetailResponse) error
		UpdateProduct(ctx context.Context, in *ProductRequest, out *ProductDetailResponse) error
		DeleteProduct(ctx context.Context, in *ProductRequest, out *ProductDetailResponse) error
		CreateProductImg(ctx context.Context, in *ProductImgRequest, out *ProductImgDetailResponse) error
		GetProductImgsList(ctx context.Context, in *ProductImgRequest, out *ProductImgsListResponse) error
		UpdateProductImg(ctx context.Context, in *ProductImgRequest, out *ProductImgDetailResponse) error
		DeleteProductImg(ctx context.Context, in *ProductImgRequest, out *ProductImgDetailResponse) error
	}
	type ProductService struct {
		productService
	}
	h := &productServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProductService{h}, opts...))
}

type productServiceHandler struct {
	ProductServiceHandler
}

func (h *productServiceHandler) CreateProduct(ctx context.Context, in *ProductRequest, out *ProductDetailResponse) error {
	return h.ProductServiceHandler.CreateProduct(ctx, in, out)
}

func (h *productServiceHandler) GetProductsList(ctx context.Context, in *ProductRequest, out *ProductsListResponse) error {
	return h.ProductServiceHandler.GetProductsList(ctx, in, out)
}

func (h *productServiceHandler) GetProduct(ctx context.Context, in *ProductRequest, out *ProductDetailResponse) error {
	return h.ProductServiceHandler.GetProduct(ctx, in, out)
}

func (h *productServiceHandler) UpdateProduct(ctx context.Context, in *ProductRequest, out *ProductDetailResponse) error {
	return h.ProductServiceHandler.UpdateProduct(ctx, in, out)
}

func (h *productServiceHandler) DeleteProduct(ctx context.Context, in *ProductRequest, out *ProductDetailResponse) error {
	return h.ProductServiceHandler.DeleteProduct(ctx, in, out)
}

func (h *productServiceHandler) CreateProductImg(ctx context.Context, in *ProductImgRequest, out *ProductImgDetailResponse) error {
	return h.ProductServiceHandler.CreateProductImg(ctx, in, out)
}

func (h *productServiceHandler) GetProductImgsList(ctx context.Context, in *ProductImgRequest, out *ProductImgsListResponse) error {
	return h.ProductServiceHandler.GetProductImgsList(ctx, in, out)
}

func (h *productServiceHandler) UpdateProductImg(ctx context.Context, in *ProductImgRequest, out *ProductImgDetailResponse) error {
	return h.ProductServiceHandler.UpdateProductImg(ctx, in, out)
}

func (h *productServiceHandler) DeleteProductImg(ctx context.Context, in *ProductImgRequest, out *ProductImgDetailResponse) error {
	return h.ProductServiceHandler.DeleteProductImg(ctx, in, out)
}