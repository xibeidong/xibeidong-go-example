package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"os"
	"os/signal"
	"time"
	nats2 "xibeidong-go-example/14-nats"
)

var nc *nats.Conn

func main() {

	_nc, err := nats2.NewConn()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	nc = _nc
	defer nc.Close()

	i := 0
	for i < 10000 {
		i++
		time.Sleep(time.Millisecond * 20)
		publish("test", fmt.Sprintf("%v %v", time.Now().UnixMicro(), i))

	}
	fmt.Println("---------end----------------")

	sg := make(chan os.Signal)
	signal.Notify(sg, os.Kill, os.Interrupt)
	<-sg
}

func publish(subject, val string) {
	if err := nc.Publish(subject, []byte(val)); err != nil {
		fmt.Println(err)
	}
	fmt.Println(" publish ok => ", val)
}
