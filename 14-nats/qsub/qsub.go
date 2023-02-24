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
	q := "q1"
	fmt.Println("输入queue name :")
	fmt.Scanln(&q)
	_nc, err := nats2.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	nc = _nc
	defer nc.Close()

	qsub("test", q)

	sg := make(chan os.Signal)
	signal.Notify(sg, os.Kill, os.Interrupt)
	<-sg
}

func qsub(subject, qName string) {

	//订阅同一个主题，如果 queue相同，每个订阅得到部分publish的消息，
	if _, err := nc.QueueSubscribe(subject, qName, func(msg *nats.Msg) {
		t := time.Now().UnixMicro()
		str := string(msg.Data)
		split := strings.Split(str, " ")
		i, _ := strconv.ParseInt(split[0], 10, 64)
		fmt.Println("delay = ", t-i, "  ", split[1])
		//msg.Respond([]byte("ok"))
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
