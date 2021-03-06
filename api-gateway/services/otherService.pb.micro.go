// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: otherService.proto

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

// Client API for OtherService service

type OtherService interface {
	//轮播图服务
	GetCarouselsList(ctx context.Context, in *CarouselRequest, opts ...client.CallOption) (*CarouselsListResponse, error)
	GetCarousel(ctx context.Context, in *CarouselRequest, opts ...client.CallOption) (*CarouselDetailResponse, error)
	UpdateCarousel(ctx context.Context, in *CarouselRequest, opts ...client.CallOption) (*CarouselDetailResponse, error)
	//公告服务
	CreateNotice(ctx context.Context, in *NoticeRequest, opts ...client.CallOption) (*NoticeDetailResponse, error)
	GetNotice(ctx context.Context, in *NoticeRequest, opts ...client.CallOption) (*NoticeDetailResponse, error)
	UpdateNotice(ctx context.Context, in *NoticeRequest, opts ...client.CallOption) (*NoticeDetailResponse, error)
	DeleteNotice(ctx context.Context, in *NoticeRequest, opts ...client.CallOption) (*NoticeDetailResponse, error)
	//分类服务
	CreateCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoryDetailResponse, error)
	GetCategoriesList(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoriesListResponse, error)
	GetCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoryDetailResponse, error)
	UpdateCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoryDetailResponse, error)
	DeleteCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoryDetailResponse, error)
}

type otherService struct {
	c    client.Client
	name string
}

func NewOtherService(name string, c client.Client) OtherService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "services"
	}
	return &otherService{
		c:    c,
		name: name,
	}
}

func (c *otherService) GetCarouselsList(ctx context.Context, in *CarouselRequest, opts ...client.CallOption) (*CarouselsListResponse, error) {
	req := c.c.NewRequest(c.name, "OtherService.GetCarouselsList", in)
	out := new(CarouselsListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherService) GetCarousel(ctx context.Context, in *CarouselRequest, opts ...client.CallOption) (*CarouselDetailResponse, error) {
	req := c.c.NewRequest(c.name, "OtherService.GetCarousel", in)
	out := new(CarouselDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherService) UpdateCarousel(ctx context.Context, in *CarouselRequest, opts ...client.CallOption) (*CarouselDetailResponse, error) {
	req := c.c.NewRequest(c.name, "OtherService.UpdateCarousel", in)
	out := new(CarouselDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherService) CreateNotice(ctx context.Context, in *NoticeRequest, opts ...client.CallOption) (*NoticeDetailResponse, error) {
	req := c.c.NewRequest(c.name, "OtherService.CreateNotice", in)
	out := new(NoticeDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherService) GetNotice(ctx context.Context, in *NoticeRequest, opts ...client.CallOption) (*NoticeDetailResponse, error) {
	req := c.c.NewRequest(c.name, "OtherService.GetNotice", in)
	out := new(NoticeDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherService) UpdateNotice(ctx context.Context, in *NoticeRequest, opts ...client.CallOption) (*NoticeDetailResponse, error) {
	req := c.c.NewRequest(c.name, "OtherService.UpdateNotice", in)
	out := new(NoticeDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherService) DeleteNotice(ctx context.Context, in *NoticeRequest, opts ...client.CallOption) (*NoticeDetailResponse, error) {
	req := c.c.NewRequest(c.name, "OtherService.DeleteNotice", in)
	out := new(NoticeDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherService) CreateCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoryDetailResponse, error) {
	req := c.c.NewRequest(c.name, "OtherService.CreateCategory", in)
	out := new(CategoryDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherService) GetCategoriesList(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoriesListResponse, error) {
	req := c.c.NewRequest(c.name, "OtherService.GetCategoriesList", in)
	out := new(CategoriesListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherService) GetCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoryDetailResponse, error) {
	req := c.c.NewRequest(c.name, "OtherService.GetCategory", in)
	out := new(CategoryDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherService) UpdateCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoryDetailResponse, error) {
	req := c.c.NewRequest(c.name, "OtherService.UpdateCategory", in)
	out := new(CategoryDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherService) DeleteCategory(ctx context.Context, in *CategoryRequest, opts ...client.CallOption) (*CategoryDetailResponse, error) {
	req := c.c.NewRequest(c.name, "OtherService.DeleteCategory", in)
	out := new(CategoryDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OtherService service

type OtherServiceHandler interface {
	//轮播图服务
	GetCarouselsList(context.Context, *CarouselRequest, *CarouselsListResponse) error
	GetCarousel(context.Context, *CarouselRequest, *CarouselDetailResponse) error
	UpdateCarousel(context.Context, *CarouselRequest, *CarouselDetailResponse) error
	//公告服务
	CreateNotice(context.Context, *NoticeRequest, *NoticeDetailResponse) error
	GetNotice(context.Context, *NoticeRequest, *NoticeDetailResponse) error
	UpdateNotice(context.Context, *NoticeRequest, *NoticeDetailResponse) error
	DeleteNotice(context.Context, *NoticeRequest, *NoticeDetailResponse) error
	//分类服务
	CreateCategory(context.Context, *CategoryRequest, *CategoryDetailResponse) error
	GetCategoriesList(context.Context, *CategoryRequest, *CategoriesListResponse) error
	GetCategory(context.Context, *CategoryRequest, *CategoryDetailResponse) error
	UpdateCategory(context.Context, *CategoryRequest, *CategoryDetailResponse) error
	DeleteCategory(context.Context, *CategoryRequest, *CategoryDetailResponse) error
}

func RegisterOtherServiceHandler(s server.Server, hdlr OtherServiceHandler, opts ...server.HandlerOption) error {
	type otherService interface {
		GetCarouselsList(ctx context.Context, in *CarouselRequest, out *CarouselsListResponse) error
		GetCarousel(ctx context.Context, in *CarouselRequest, out *CarouselDetailResponse) error
		UpdateCarousel(ctx context.Context, in *CarouselRequest, out *CarouselDetailResponse) error
		CreateNotice(ctx context.Context, in *NoticeRequest, out *NoticeDetailResponse) error
		GetNotice(ctx context.Context, in *NoticeRequest, out *NoticeDetailResponse) error
		UpdateNotice(ctx context.Context, in *NoticeRequest, out *NoticeDetailResponse) error
		DeleteNotice(ctx context.Context, in *NoticeRequest, out *NoticeDetailResponse) error
		CreateCategory(ctx context.Context, in *CategoryRequest, out *CategoryDetailResponse) error
		GetCategoriesList(ctx context.Context, in *CategoryRequest, out *CategoriesListResponse) error
		GetCategory(ctx context.Context, in *CategoryRequest, out *CategoryDetailResponse) error
		UpdateCategory(ctx context.Context, in *CategoryRequest, out *CategoryDetailResponse) error
		DeleteCategory(ctx context.Context, in *CategoryRequest, out *CategoryDetailResponse) error
	}
	type OtherService struct {
		otherService
	}
	h := &otherServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OtherService{h}, opts...))
}

type otherServiceHandler struct {
	OtherServiceHandler
}

func (h *otherServiceHandler) GetCarouselsList(ctx context.Context, in *CarouselRequest, out *CarouselsListResponse) error {
	return h.OtherServiceHandler.GetCarouselsList(ctx, in, out)
}

func (h *otherServiceHandler) GetCarousel(ctx context.Context, in *CarouselRequest, out *CarouselDetailResponse) error {
	return h.OtherServiceHandler.GetCarousel(ctx, in, out)
}

func (h *otherServiceHandler) UpdateCarousel(ctx context.Context, in *CarouselRequest, out *CarouselDetailResponse) error {
	return h.OtherServiceHandler.UpdateCarousel(ctx, in, out)
}

func (h *otherServiceHandler) CreateNotice(ctx context.Context, in *NoticeRequest, out *NoticeDetailResponse) error {
	return h.OtherServiceHandler.CreateNotice(ctx, in, out)
}

func (h *otherServiceHandler) GetNotice(ctx context.Context, in *NoticeRequest, out *NoticeDetailResponse) error {
	return h.OtherServiceHandler.GetNotice(ctx, in, out)
}

func (h *otherServiceHandler) UpdateNotice(ctx context.Context, in *NoticeRequest, out *NoticeDetailResponse) error {
	return h.OtherServiceHandler.UpdateNotice(ctx, in, out)
}

func (h *otherServiceHandler) DeleteNotice(ctx context.Context, in *NoticeRequest, out *NoticeDetailResponse) error {
	return h.OtherServiceHandler.DeleteNotice(ctx, in, out)
}

func (h *otherServiceHandler) CreateCategory(ctx context.Context, in *CategoryRequest, out *CategoryDetailResponse) error {
	return h.OtherServiceHandler.CreateCategory(ctx, in, out)
}

func (h *otherServiceHandler) GetCategoriesList(ctx context.Context, in *CategoryRequest, out *CategoriesListResponse) error {
	return h.OtherServiceHandler.GetCategoriesList(ctx, in, out)
}

func (h *otherServiceHandler) GetCategory(ctx context.Context, in *CategoryRequest, out *CategoryDetailResponse) error {
	return h.OtherServiceHandler.GetCategory(ctx, in, out)
}

func (h *otherServiceHandler) UpdateCategory(ctx context.Context, in *CategoryRequest, out *CategoryDetailResponse) error {
	return h.OtherServiceHandler.UpdateCategory(ctx, in, out)
}

func (h *otherServiceHandler) DeleteCategory(ctx context.Context, in *CategoryRequest, out *CategoryDetailResponse) error {
	return h.OtherServiceHandler.DeleteCategory(ctx, in, out)
}