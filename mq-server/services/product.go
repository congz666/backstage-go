// Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-10-26 16:15:42
 * @LastEditors: congz
 * @LastEditTime: 2020-10-27 14:16:43
 */
package services

import (
	"encoding/json"
	"log"
	"mq-server/model"
)

//CreateProduct 从RabbitMQ接收信息再写入数据库
func CreateProduct() {
	ch, err := model.MQ.Channel()
	if err != nil {
		panic(err)
	}
	q, err := ch.QueueDeclare("product_queue", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	err = ch.Qos(1, 0, false)
	if err != nil {
		panic(err)
	}
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		for d := range msgs {
			var p model.Product
			err := json.Unmarshal(d.Body, &p)
			if err != nil {
				panic(err)
			}
			model.DB.Create(&p)
			log.Printf("Done")
			d.Ack(false)
		}
	}()
}
