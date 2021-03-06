//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 11:11:17
 * @LastEditors: congz
 * @LastEditTime: 2020-10-12 13:49:53
 */
package model

import (
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string `gorm:"unique"`
	PasswordDigest string
	Nickname       string `gorm:"unique"`
	Status         string
	Limit          uint32
	Avatar         string `gorm:"size:1000"`
}

// AvatarURL 封面地址
func (user *User) AvatarURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(user.Avatar, oss.HTTPGet, 1*60*60)
	//if strings.Contains(signedGetURL, "http://ailiaili-img-av.oss-cn-hangzhou.aliyuncs.com/?Exp") {
	//signedGetURL := "https://ailiaili-img-av.oss-cn-hangzhou.aliyuncs.com/img/noface.png"
	//}
	return signedGetURL
}
