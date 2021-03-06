// Package serviceimpl ...
/*
 * @Descripttion:商品服务
 * @Author: congz
 * @Date: 2020-09-15 10:57:26
 * @LastEditors: congz
 * @LastEditTime: 2020-10-27 13:58:33
 */
package serviceimpl

import (
	"context"
	"errors"
	"other/model"
	"other/services"
)

//BuildCarousel 序列化轮播图
func BuildCarousel(item model.Carousel) *services.CarouselModel {
	carouselModel := services.CarouselModel{
		ID:        uint32(item.ID),
		ProductID: item.ProductID,
		ImgPath:   item.ImgPath,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &carouselModel
}

//GetCarouselsList 实现其他服务接口 获取轮播图列表
func (*OtherService) GetCarouselsList(context context.Context, req *services.CarouselRequest, res *services.CarouselsListResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}
	carouselData := []model.Carousel{}
	var count uint32
	err := model.DB.Offset(req.Start).Limit(req.Limit).Find(&carouselData).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	err = model.DB.Model(&model.Carousel{}).Count(&count).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	carouselRes := []*services.CarouselModel{}
	for _, item := range carouselData {
		carouselRes = append(carouselRes, BuildCarousel(item))
	}
	res.CarouselsList = carouselRes
	res.Count = count
	return nil
}

//GetCarousel 实现其他服务接口 获取轮播图
func (*OtherService) GetCarousel(context context.Context, req *services.CarouselRequest, res *services.CarouselDetailResponse) error {
	carouselData := model.Carousel{}
	err := model.DB.First(&carouselData, req.CarouselID).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	res.CarouselDetail = BuildCarousel(carouselData)
	return nil
}

//UpdateCarousel 实现其他服务接口 修改轮播图
func (*OtherService) UpdateCarousel(context context.Context, req *services.CarouselRequest, res *services.CarouselDetailResponse) error {
	carouselData := model.Carousel{}
	err := model.DB.First(&carouselData, req.CarouselID).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	carouselData.ImgPath = req.ImgPath
	carouselData.ProductID = req.ProductID
	err = model.DB.Save(&carouselData).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	return nil
}
