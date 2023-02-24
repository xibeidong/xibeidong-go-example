package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"time"
)

func main() {

	p, err := rocketmq.NewProducer(
		//producer.WithNameServer(endPoint),
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"192.168.17.129:9876"})),
		producer.WithRetry(2),
		producer.WithGroupName("home1"),
	)
	if err != nil {
		panic(err)
	}
	err = p.Start()
	if err != nil {
		panic(err)
	}
	defer p.Shutdown()

	t := time.NewTicker(time.Millisecond * 1000)
	defer t.Stop()

	topic := "test2"
	ctx := context.Background()
	for {
		select {
		case <-t.C:
			msg := primitive.Message{
				Topic: topic,
				Body:  []byte("Hi! " + time.Now().Format("2006-01-02 15:04:05")),
			}
			result, err := p.SendSync(ctx, &msg)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("send result = ", result.Status)
			}

		}
	}

}
