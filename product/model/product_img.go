//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-10-13 16:42:08
 * @LastEditors: congz
 * @LastEditTime: 2020-10-13 23:54:04
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// ProductImg 商品图片模型
type ProductImg struct {
	gorm.Model
	ProductID uint32
	ImgPath   string
}

// ProductInfoImg 商品图片模型
type ProductInfoImg struct {
	gorm.Model
	ProductID uint32
	ImgPath   string
}

// ProductParamImg 商品图片模型
type ProductParamImg struct {
	gorm.Model
	ProductID uint32
	ImgPath   string
}
