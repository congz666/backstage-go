// Package serviceimpl ...
/*
 * @Descripttion:商品服务
 * @Author: congz
 * @Date: 2020-09-15 10:57:26
 * @LastEditors: congz
 * @LastEditTime: 2020-10-27 13:59:16
 */
package serviceimpl

import (
	"context"
	"errors"
	"other/model"
	"other/services"
)

//BuildNotice 序列化公告
func BuildNotice(item model.Notice) *services.NoticeModel {
	noticeModel := services.NoticeModel{
		ID:        uint32(item.ID),
		Text:      item.Text,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &noticeModel
}

//CreateNotice 实现其他服务接口 新建公告
func (*OtherService) CreateNotice(ctx context.Context, req *services.NoticeRequest, res *services.NoticeDetailResponse) error {
	noticeData := model.Notice{
		Text: req.Text,
	}
	err := model.DB.Create(&noticeData).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	res.NoticeDetail = BuildNotice(noticeData)
	return nil
}

//GetNotice 实现其他服务接口 获取公告
func (*OtherService) GetNotice(ctx context.Context, req *services.NoticeRequest, res *services.NoticeDetailResponse) error {
	noticeData := model.Notice{}
	err := model.DB.First(&noticeData, req.NoticeID).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	res.NoticeDetail = BuildNotice(noticeData)
	return nil
}

//UpdateNotice 实现其他服务接口 修改公告
func (*OtherService) UpdateNotice(ctx context.Context, req *services.NoticeRequest, res *services.NoticeDetailResponse) error {
	noticeData := model.Notice{}
	err := model.DB.First(&noticeData, req.NoticeID).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	noticeData.Text = req.Text
	err = model.DB.Save(&noticeData).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	return nil
}

//DeleteNotice 实现其他服务接口 删除公告
func (*OtherService) DeleteNotice(ctx context.Context, req *services.NoticeRequest, res *services.NoticeDetailResponse) error {
	err := model.DB.Delete(&model.Notice{}, req.NoticeID).Error
	if err != nil {
		err = errors.New("mysql err:" + err.Error())
		return err
	}
	return nil
}
