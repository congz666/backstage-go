// Package serviceimpl ...
/*
 * @Descripttion:商品服务
 * @Author: congz
 * @Date: 2020-09-15 10:57:26
 * @LastEditors: congz
 * @LastEditTime: 2020-10-27 13:58:59
 */
package serviceimpl

import (
	"context"
	"errors"
	"other/model"
	"other/services"
)

//BuildCategory 序列化分类
func BuildCategory(item model.Category) *services.CategoryModel {
	categoryModel := services.CategoryModel{
		ID:           uint32(item.ID),
		CategoryID:   item.CategoryID,
		CategoryName: item.CategoryName,
		Num:          item.Num,
		CreatedAt:    item.CreatedAt.Unix(),
		UpdatedAt:    item.UpdatedAt.Unix(),
	}
	return &categoryModel
}

//CreateCategory 实现其他服务接口 新建分类
func (*OtherService) CreateCategory(context context.Context, req *services.CategoryRequest, res *services.CategoryDetailResponse) error {
	categoryData := model.Category{
		CategoryID:   req.CategoryID,
		CategoryName: req.CategoryName,
		Num:          0,
	}
	err := model.DB.Create(&categoryData).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	res.CategoryDetail = BuildCategory(categoryData)
	return nil
}

//GetCategoriesList 实现其他服务接口 获取分类列表
func (*OtherService) GetCategoriesList(context context.Context, req *services.CategoryRequest, res *services.CategoriesListResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}
	categoryData := []model.Category{}
	var count uint32
	err := model.DB.Offset(req.Start).Limit(req.Limit).Find(&categoryData).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	err = model.DB.Model(&model.Category{}).Count(&count).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	categoryRes := []*services.CategoryModel{}
	for _, item := range categoryData {
		categoryRes = append(categoryRes, BuildCategory(item))
	}
	res.CategoriesList = categoryRes
	res.Count = count
	return nil
}

//GetCategory 实现其他服务接口 获取分类
func (*OtherService) GetCategory(context context.Context, req *services.CategoryRequest, res *services.CategoryDetailResponse) error {
	return nil
}

//UpdateCategory 实现其他服务接口 修改分类
func (*OtherService) UpdateCategory(context context.Context, req *services.CategoryRequest, res *services.CategoryDetailResponse) error {
	categoryData := model.Category{}
	err := model.DB.First(&categoryData, req.ID).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	categoryData.CategoryID = req.CategoryID
	categoryData.CategoryName = req.CategoryName
	categoryData.Num = req.Num
	err = model.DB.Save(&categoryData).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	return nil
}

//DeleteCategory 实现其他服务接口 删除分类
func (*OtherService) DeleteCategory(context context.Context, req *services.CategoryRequest, res *services.CategoryDetailResponse) error {
	err := model.DB.Delete(&model.Category{}, req.ID).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	return nil
}
