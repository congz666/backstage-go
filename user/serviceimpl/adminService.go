// Package serviceimpl ...
/*
 * @Descripttion:商品服务
 * @Author: congz
 * @Date: 2020-09-15 10:57:26
 * @LastEditors: congz
 * @LastEditTime: 2020-10-31 14:47:00
 */
package serviceimpl

import (
	"context"
	"errors"
	"os"
	"user/model"
	"user/services"

	"github.com/jinzhu/gorm"
)

//BuildAdmin 序列化管理员
func BuildAdmin(item model.Admin) *services.AdminModel {
	adminModel := services.AdminModel{
		ID:        uint32(item.ID),
		UserName:  item.UserName,
		Avatar:    item.Avatar,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &adminModel
}

//AdminLogin 实现用户服务接口 管理员登录
func (*UserService) AdminLogin(ctx context.Context, req *services.AdminRequest, res *services.AdminDetailResponse) error {
	var admin model.Admin
	res.Code = 200
	if err := model.DB.Where("user_name = ?", req.UserName).First(&admin).Error; err != nil {
		//如果查询不到，返回相应错误
		if gorm.IsRecordNotFoundError(err) {
			res.Code = 10003
			return nil
		}
		res.Code = 30001
		return nil
	}
	if admin.CheckPassword(req.Password) == false {
		res.Code = 10004
		return nil
	}
	adminRes := BuildAdmin(admin)
	res.AdminDetail = adminRes
	return nil
}

// AdminRegister 管理员注册
func (*UserService) AdminRegister(ctx context.Context, req *services.AdminRequest, res *services.AdminDetailResponse) error {
	if req.Identification != os.Getenv("Identification") {
		err := errors.New("识别码错误")
		return err
	}
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}
	count := 0
	if err := model.DB.Model(&model.Admin{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := errors.New("用户名已存在")
		return err
	}
	admin := model.Admin{
		UserName: req.UserName,
		Avatar:   "img/avatar/avatar1.jpg",
	}
	// 加密密码
	if err := admin.SetPassword(req.Password); err != nil {
		return err
	}
	// 创建用户
	if err := model.DB.Create(&admin).Error; err != nil {
		return err
	}
	adminRes := BuildAdmin(admin)
	res.AdminDetail = adminRes
	return nil
}
