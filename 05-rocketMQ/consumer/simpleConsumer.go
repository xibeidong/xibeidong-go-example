package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"
)

func main() {

	//a := 1
	//b := 1
	////a &^= b
	////a = a & (^b)
	//b = ^b
	fmt.Println(1&^1, 1&^0)
	//rlog.SetLogLevel("error")
	//go NewConsumer("test1_home1")
	//time.Sleep(time.Hour)
}

func NewConsumer(group string) {
	pushConsumer, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(group),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"192.168.17.129:9876"})),
	)
	if err != nil {
		panic(err)
	}

	err = pushConsumer.Subscribe("test2",
		consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i := range msgs {
				fmt.Printf("subscribe callback: %v \n", string(msgs[i].Body))
			}
			return consumer.ConsumeSuccess, nil
		})

	if err != nil {
		panic(err)
	}
	if err = pushConsumer.Start(); err != nil {
		panic(err)
	}
	defer pushConsumer.Shutdown()
	time.Sleep(time.Hour)
}
