package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"strconv"
	"time"
	nats2 "xibeidong-go-example/14-nats"
)

func main() {
	// 发布-订阅模式
	nc, err := nats2.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer nc.Close()

	// 异步订阅test01 , 同步订阅可使用 SubscribeSync
	if _, err = nc.Subscribe("test01", func(msg *nats.Msg) {
		fmt.Println("user1 recv: ", string(msg.Data))
	}); err != nil {
		fmt.Println(err)
		return
	}

	// 异步订阅test01 , 两个订阅收到相同的消息
	if _, err = nc.Subscribe("test01", func(msg *nats.Msg) {
		fmt.Println("user2 recv: ", string(msg.Data))
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
