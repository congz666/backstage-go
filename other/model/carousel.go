//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 11:11:17
 * @LastEditors: congz
 * @LastEditTime: 2020-09-20 15:59:40
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// Carousel 轮播图模型
type Carousel struct {
	gorm.Model
	ImgPath   string
	ProductID uint32
}
