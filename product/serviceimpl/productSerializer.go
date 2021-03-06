// Package serviceimpl ...
/*
 * @Descripttion:商品服务
 * @Author: congz
 * @Date: 2020-09-15 10:57:26
 * @LastEditors: congz
 * @LastEditTime: 2020-10-14 12:25:38
 */
package serviceimpl

import (
	"product/model"
	"product/services"
)

//BuildProduct 序列化商品
func BuildProduct(item model.Product) *services.ProductModel {
	productModel := services.ProductModel{
		ID:            uint32(item.ID),
		Name:          item.Name,
		CategoryID:    item.CategoryID,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		CreatedAt:     item.CreatedAt.Unix(),
		UpdatedAt:     item.UpdatedAt.Unix(),
	}
	return &productModel
}

//BuildProductImg 序列化商品图片
func BuildProductImg(item model.ProductImg) *services.ProductImgModel {
	productImgModel := services.ProductImgModel{
		ID:        uint32(item.ID),
		ProductID: item.ProductID,
		ImgPath:   item.ImgPath,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &productImgModel
}

//BuildProductInfoImg 序列化商品介绍图片
func BuildProductInfoImg(item model.ProductInfoImg) *services.ProductImgModel {
	productImgModel := services.ProductImgModel{
		ID:        uint32(item.ID),
		ProductID: item.ProductID,
		ImgPath:   item.ImgPath,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &productImgModel
}

//BuildProductParamImg 序列化商品参数图片
func BuildProductParamImg(item model.ProductParamImg) *services.ProductImgModel {
	productImgModel := services.ProductImgModel{
		ID:        uint32(item.ID),
		ProductID: item.ProductID,
		ImgPath:   item.ImgPath,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &productImgModel
}
