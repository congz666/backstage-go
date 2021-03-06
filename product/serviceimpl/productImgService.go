// Package serviceimpl ...
/*
 * @Descripttion:商品服务
 * @Author: congz
 * @Date: 2020-09-15 10:57:26
 * @LastEditors: congz
 * @LastEditTime: 2020-10-27 13:48:48
 */
package serviceimpl

import (
	"context"
	"errors"
	"product/model"
	"product/services"
)

//CreateProductImg 实现商品服务接口 创建商品图片
func (*ProductService) CreateProductImg(ctx context.Context, req *services.ProductImgRequest, res *services.ProductImgDetailResponse) error {
	switch req.ImgType {
	case 1:
		productImgData := model.ProductImg{
			ProductID: req.ProductID,
			ImgPath:   req.ImgPath,
		}
		if err := model.DB.Create(&productImgData).Error; err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		res.ProductImgDetail = BuildProductImg(productImgData)
		return nil
	case 2:
		productImgData := model.ProductInfoImg{
			ProductID: req.ProductID,
			ImgPath:   req.ImgPath,
		}
		if err := model.DB.Create(&productImgData).Error; err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		res.ProductImgDetail = BuildProductInfoImg(productImgData)
		return nil
	case 3:
		productImgData := model.ProductParamImg{
			ProductID: req.ProductID,
			ImgPath:   req.ImgPath,
		}
		if err := model.DB.Create(&productImgData).Error; err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		res.ProductImgDetail = BuildProductParamImg(productImgData)
		return nil
	default:
		return errors.New("type err: 没有此类型")
	}
}

//GetProductImgsList 实现商品服务接口 获取商品图片列表
func (*ProductService) GetProductImgsList(ctx context.Context, req *services.ProductImgRequest, res *services.ProductImgsListResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}
	switch req.ImgType {
	case 1:
		var productImgData []model.ProductImg
		var count uint32
		if err := model.DB.Offset(req.Start).Limit(req.Limit).Where("product_id=?", req.ProductID).Find(&productImgData).Error; err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		if err := model.DB.Model(&model.ProductImg{}).Where("product_id=?", req.ProductID).Count(&count).Error; err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}

		//序类化商品图片列表
		productImgRes := []*services.ProductImgModel{}
		for _, item := range productImgData {
			productImgRes = append(productImgRes, BuildProductImg(item))
		}
		//序列化后的结果赋给response
		res.ProductImgsList = productImgRes
		res.Count = count
		return nil
	case 2:
		var productImgData []model.ProductInfoImg
		var count uint32
		if err := model.DB.Offset(req.Start).Limit(req.Limit).Where("product_id=?", req.ProductID).Find(&productImgData).Error; err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		if err := model.DB.Model(&model.ProductInfoImg{}).Where("product_id=?", req.ProductID).Count(&count).Error; err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		//序类化商品图片列表
		productImgRes := []*services.ProductImgModel{}
		for _, item := range productImgData {
			productImgRes = append(productImgRes, BuildProductInfoImg(item))
		}
		//序列化后的结果赋给response
		res.ProductImgsList = productImgRes
		res.Count = count
		return nil
	case 3:
		var productImgData []model.ProductParamImg
		var count uint32
		if err := model.DB.Offset(req.Start).Limit(req.Limit).Where("product_id=?", req.ProductID).Find(&productImgData).Error; err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		if err := model.DB.Model(&model.ProductParamImg{}).Where("product_id=?", req.ProductID).Count(&count).Error; err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		//序类化商品图片列表
		productImgRes := []*services.ProductImgModel{}
		for _, item := range productImgData {
			productImgRes = append(productImgRes, BuildProductParamImg(item))
		}
		//序列化后的结果赋给response
		res.ProductImgsList = productImgRes
		res.Count = count
		return nil
	default:
		return errors.New("type err: 没有此类型")
	}

}

//UpdateProductImg 实现商品服务接口 修改商品图片
func (*ProductService) UpdateProductImg(ctx context.Context, req *services.ProductImgRequest, res *services.ProductImgDetailResponse) error {
	switch req.ImgType {
	case 1:
		//在数据库查找值
		productImgData := model.ProductImg{}
		err := model.DB.First(&productImgData, req.ImgID).Error
		if err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		//将要更新的数据赋值给结构体
		productImgData.ImgPath = req.ImgPath
		//update
		err = model.DB.Save(&productImgData).Error
		if err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		//序类化商品
		res.ProductImgDetail = BuildProductImg(productImgData)
		return nil
	case 2:
		//在数据库查找值
		productImgData := model.ProductInfoImg{}
		err := model.DB.First(&productImgData, req.ImgID).Error
		if err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		//将要更新的数据赋值给结构体
		productImgData.ImgPath = req.ImgPath
		//update
		err = model.DB.Save(&productImgData).Error
		if err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		//序类化商品
		res.ProductImgDetail = BuildProductInfoImg(productImgData)
		return nil
	case 3:
		//在数据库查找值
		productImgData := model.ProductParamImg{}
		err := model.DB.First(&productImgData, req.ImgID).Error
		if err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		//将要更新的数据赋值给结构体
		productImgData.ImgPath = req.ImgPath
		//update
		err = model.DB.Save(&productImgData).Error
		if err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		//序类化商品
		res.ProductImgDetail = BuildProductParamImg(productImgData)
		return nil
	default:
		return errors.New("type err: 没有此类型")
	}
}

//DeleteProductImg 实现商品服务接口 删除商品图片
func (*ProductService) DeleteProductImg(ctx context.Context, req *services.ProductImgRequest, res *services.ProductImgDetailResponse) error {
	switch req.ImgType {
	case 1:
		//在数据库删除值
		err := model.DB.Delete(&model.ProductImg{}, req.ImgID).Error
		if err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		return nil
	case 2:
		//在数据库删除值
		err := model.DB.Delete(&model.ProductInfoImg{}, req.ImgID).Error
		if err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		return nil
	case 3:
		//在数据库删除值
		err := model.DB.Delete(&model.ProductParamImg{}, req.ImgID).Error
		if err != nil {
			err = errors.New("mysql err:" + err.Error())
			return err
		}
		return nil
	default:
		return errors.New("type err: 没有此类型")
	}
}
