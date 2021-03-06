//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 10:58:11
 * @LastEditors: congz
 * @LastEditTime: 2020-10-26 15:34:27
 */
package model

import (
	"github.com/streadway/amqp"
)

// MQ rabbitMQ链接单例
var MQ *amqp.Connection

// RabbitMQ 在中间件中初始化rabbitMQ链接
func RabbitMQ(connString string) {
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	MQ = conn
}
