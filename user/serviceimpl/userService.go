// Package serviceimpl ...
/*
 * @Descripttion:商品服务
 * @Author: congz
 * @Date: 2020-09-15 10:57:26
 * @LastEditors: congz
 * @LastEditTime: 2020-10-27 13:55:44
 */
package serviceimpl

import (
	"context"
	"errors"
	"user/model"
	"user/services"
)

//BuildUser 序列化用户
func BuildUser(item model.User) *services.UserModel {
	userModel := services.UserModel{
		ID:        uint32(item.ID),
		UserName:  item.UserName,
		Avatar:    item.AvatarURL(),
		Email:     item.Email,
		NickName:  item.Nickname,
		Status:    item.Status,
		Limit:     item.Limit,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}

//GetUsersList 实现用户服务接口 获取用户列表
func (*UserService) GetUsersList(ctx context.Context, req *services.UserRequest, res *services.UsersListResponse) error {
	if req.Limit == 0 {
		req.Limit = 6
	}
	//在数据库查找值
	userData := []model.User{}
	var count uint32
	err := model.DB.Offset(req.Start).Limit(req.Limit).Find(&userData).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	err = model.DB.Model(&model.User{}).Count(&count).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	//序列化用户列表
	userRes := []*services.UserModel{}
	for _, item := range userData {
		userRes = append(userRes, BuildUser(item))
	}
	//序列化后的结果赋给response
	res.UsersList = userRes
	res.Count = count
	return nil
}

//GetUser 实现用户服务接口 获取用户详情
func (*UserService) GetUser(ctx context.Context, req *services.UserRequest, res *services.UserDetailResponse) error {
	userData := model.User{}
	err := model.DB.First(&userData, req.UserID).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	userRes := BuildUser(userData)
	res.UserDetail = userRes
	return nil
}
