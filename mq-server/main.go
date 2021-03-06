// package main
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-10-26 16:10:56
 * @LastEditors: congz
 * @LastEditTime: 2020-10-27 14:17:01
 */
package main

import (
	"mq-server/conf"
	"mq-server/services"
)

func main() {
	conf.Init()

	forever := make(chan bool)
	services.CreateProduct()
	<-forever
}
