package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"time"
)

func main() {
	rlog.SetLogLevel("error")
	go NewBroadcastConsumer("home1")
	time.Sleep(time.Hour)
}

func NewBroadcastConsumer(group string) {
	pushConsumer, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(group),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"192.168.17.129:9876"})),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromLastOffset),
		consumer.WithConsumerModel(consumer.BroadCasting),
	)
	if err != nil {
		panic(err)
	}

	err = pushConsumer.Subscribe("test1",
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
