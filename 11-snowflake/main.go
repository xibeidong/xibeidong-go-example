package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"log"
	"time"
)

func main() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 100; i++ {
		id := node.Generate()
		fmt.Println(id.Int64())
		time.Sleep(time.Millisecond * 15)
	}

}
