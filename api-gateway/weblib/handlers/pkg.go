// Package handlers ..
/*
 * @Descripttion:处理函数
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-10-27 13:38:48
 */
package handlers

import (
	"api-gateway/pkg/logging"
	"errors"
)

//PanicIfProductError 包装错误函数
func PanicIfProductError(err error) {
	if err != nil {
		err = errors.New("productService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}

//PanicIfUserError 包装错误函数
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}

//PanicIfOtherError 包装错误函数
func PanicIfOtherError(err error) {
	if err != nil {
		err = errors.New("otherService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}
