package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
	nats2 "xibeidong-go-example/14-nats"
)

var nc *nats.Conn

func main() {

	_nc, err := nats2.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	nc = _nc
	defer nc.Close()

	sub("test")

	sg := make(chan os.Signal)
	signal.Notify(sg, os.Kill, os.Interrupt)
	<-sg
}

func sub(subject string) {
	if _, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		t := time.Now().UnixMicro()
		str := string(msg.Data)
		split := strings.Split(str, " ")
		i, _ := strconv.ParseInt(split[0], 10, 64)
		fmt.Println("delay = ", t-i, "  ", split[1])
	}); err != nil {
		fmt.Println(err)
		return
	}

	if err := nc.Flush(); err != nil {
		fmt.Println(err)
	}

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}
}

func syncSub(subject string) {
	sub, err := nc.SubscribeSync(subject)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg, err := sub.NextMsg(time.Second * 1)
	if err != nil {
		fmt.Println(err)
	}
	str := string(msg.Data)
	fmt.Println(str)
}
