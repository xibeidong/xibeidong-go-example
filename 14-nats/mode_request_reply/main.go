package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
	nats2 "xibeidong-go-example/14-nats"
)

func main() {
	// 请求-响应模式，只接受第一个的响应
	nc, err := nats2.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer nc.Close()

	// 异步订阅test01 , 收到消息后 reply
	if _, err = nc.Subscribe("test01", func(msg *nats.Msg) {
		fmt.Println("user1 recv: ", string(msg.Data))
		nc.Publish(msg.Reply, []byte("this is user1,i can help you !"))
	}); err != nil {
		fmt.Println(err)
		return
	}

	// 异步订阅test01 , 收到消息后reply
	if _, err = nc.Subscribe("test01", func(msg *nats.Msg) {
		fmt.Println("user2 recv: ", string(msg.Data))
		fmt.Println("user2 recv: msg.Reply= ", msg.Reply)
		// msg.Respond([]byte("this is user2,i can help you !"))
		nc.Publish(msg.Reply, []byte("this is user2,i can help you !"))
	}); err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		i := 0
		for i < 3 {
			time.Sleep(time.Second)
			i++
			// request 只接受第一个reply，其他的不要
			msg, err2 := nc.Request("test01", []byte("this is user3. help me !"), time.Second)
			if err2 != nil {
				fmt.Println(err2)
				continue
			}
			fmt.Println("user3 recv reply : ", string(msg.Data))
		}
	}()

	select {}
}
