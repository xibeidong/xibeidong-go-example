package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"math/rand"
	"sync"
	"time"
	"xibeidong-go-example/07-grpc/protos/common"
)

func main() {
	conn, err := grpc.Dial(":8888", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := common.NewPersonGreeterClient(conn)
	testChatPoints(client)
}
func testSay(client common.PersonGreeterClient) {
	replySay, err := client.Say(context.Background(), &common.RequestSay{Str: "nn1"})
	if err != nil {
		panic(err)
	}
	fmt.Println(replySay.Str)
}
func testGetPoints(client common.PersonGreeterClient) {
	points, err := client.GetPoints(context.Background(), &common.Person{
		Name: "nn",
		Id:   102,
	})
	if err != nil {
		panic(err)
	}
	for {
		point, err := points.Recv()
		if err != nil {
			fmt.Println("err: ", err)
			if err == io.EOF {
				//
			}
			break
		}
		fmt.Println(*point)
	}
}
func testSavePoints(client common.PersonGreeterClient) {
	points, err := client.SavePoints(context.Background())
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		points.Send(&common.Point{
			X: float32(rand.Intn(1000) + 1000),
			Y: float32(rand.Intn(100) + 100),
			Z: float32(rand.Intn(10) + 10),
		})
	}
	//关闭写入流，并接收返回
	reply, err := points.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	fmt.Println(reply.Data)
}
func testChatPoints(client common.PersonGreeterClient) {
	points, err := client.ChatPoints(context.Background())
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			point, err2 := points.Recv()
			if err2 != nil {
				break
			}
			fmt.Println(*point)
		}
		fmt.Println("recv over!")
	}()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		points.Send(&common.Person{
			Name: "nn",
			Id:   102,
		})
	}
	//关闭发送流，如果不关闭，server端的接收会一直阻塞
	err = points.CloseSend()
	if err != nil {
		panic(err)
	}
	wg.Wait()
}
