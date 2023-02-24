package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"math/rand"
	"net"
	"sync"
	"time"
	"xibeidong-go-example/07-grpc/protos/common"
)

type server struct {
	//common.UnimplementedPersonGreeterServer
}

func (s *server) Say(ctx context.Context, in *common.RequestSay) (*common.ReplySay, error) {
	str := "nothing"
	if in.Str == "nn" {
		str = "hello,nn"
	}
	reply := common.ReplySay{Str: str}
	return &reply, nil
}

func (s *server) GetPoints(p *common.Person, server2 common.PersonGreeter_GetPointsServer) error {

	ctx, _ := context.WithTimeout(context.Background(), time.Second*7)

	go func(c context.Context) {
		tick := time.NewTicker(time.Second)
		defer tick.Stop()
		for {
			select {
			case <-c.Done():
				return
			case <-tick.C:
				server2.Send(&common.Point{
					X:        float32(rand.Intn(100)),
					Y:        0,
					Z:        0,
					PersonId: p.Id,
				})

			}
		}
	}(ctx)
	//必须暂停一会，如果直接退出，将会无法发送数据
	time.Sleep(time.Second * 15)
	//cancel()
	return nil
}
func (s *server) SavePoints(server2 common.PersonGreeter_SavePointsServer) error {
	for {
		point, err := server2.Recv()
		if err != nil {
			if err == io.EOF {
				//发送响应后关闭
				server2.SendAndClose(&common.ReplyCommon{Data: "服务端接收完毕！"})
			}
			break
		}
		fmt.Println(*point)
	}
	return nil
}
func (s *server) ChatPoints(server2 common.PersonGreeter_ChatPointsServer) error {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			person, err := server2.Recv()
			if err != nil {
				break
			}
			fmt.Println(*person)
		}
	}()
	for i := 0; i < 10; i++ {
		server2.Send(&common.Point{
			X: float32(rand.Intn(1000) + 1000),
			Y: float32(rand.Intn(100) + 100),
			Z: float32(rand.Intn(10) + 10),
		})
	}

	wg.Wait()
	fmt.Println("chat over!")
	return nil
}
func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	common.RegisterPersonGreeterServer(s, &server{})
	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}

}
