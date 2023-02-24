package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"strconv"
	"time"
	nats2 "xibeidong-go-example/14-nats"
)

func main() {
	// 队列模式，同一个队列的多个订阅者，一条消息只能给一个订阅者，其他人收不到
	nc, err := nats2.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer nc.Close()

	// 异步订阅test01, queue = q1
	if _, err = nc.QueueSubscribe("test01", "q1", func(msg *nats.Msg) {
		fmt.Println("user1 queue=q1 recv: ", string(msg.Data))
	}); err != nil {
		fmt.Println(err)
		return
	}

	// 异步订阅test01 , queue = q1, 两个订阅者，queue相同，所以消息只能被一个收到
	if _, err = nc.QueueSubscribe("test01", "q1", func(msg *nats.Msg) {
		fmt.Println("user2 queue=q1 recv: ", string(msg.Data))
	}); err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		i := 0
		for i < 1000 {
			time.Sleep(time.Second)
			i++
			// 发布test01
			err3 := nc.Publish("test01", []byte(strconv.Itoa(i)))
			if err3 != nil {
				return
			}

		}
	}()

	select {}
}
