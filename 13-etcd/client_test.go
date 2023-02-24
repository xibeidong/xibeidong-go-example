package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"testing"
	"time"
)

var cli *clientv3.Client

func TestMain(m *testing.M) {
	fmt.Println("---------------begin-------------------")

	//client, err := NewClient([]string{"192.168.254.128:2379"})

	client, err := NewClient([]string{
		"192.168.254.128:2479",
		"192.168.254.128:2579",
		"192.168.254.128:2679"})

	if err != nil {
		fmt.Println(err)
	}
	cli = client
	defer client.Close()

	m.Run()
	fmt.Println("---------------end-------------------")

}

func TestPut(t *testing.T) {
	err := put(cli, "addr01", "shanghai")
	if err != nil {
		return
	}
}

func TestGet(t *testing.T) {
	err := get(cli, "addr01")
	if err != nil {
		return
	}
}

func TestWatch(t *testing.T) {

	go watch(cli, "addr01")

	time.Sleep(time.Second * 1)
	put(cli, "addr01", "天津")

	time.Sleep(time.Second * 1)
	put(cli, "addr01", "保定")

	time.Sleep(time.Second * 1)
	cli.Delete(context.Background(), "addr01")

	time.Sleep(time.Second * 3)

}

func TestLease(t *testing.T) {
	// 创建一个5秒的租约
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	// 5秒钟之后, /lmh/ 这个key就会被移除
	_, err = cli.Put(context.TODO(), "/lmh/", "lmh", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
}

func TestKeepalive(t *testing.T) {
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	_, err = cli.Put(context.TODO(), "/lmh/", "lmh", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}

	// the key 'foo' will be kept forever
	ch, kaerr := cli.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		log.Fatal(kaerr)
	}
	for {
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}

//https://juejin.cn/post/7062900835038003208

func TestLock(t *testing.T) {
	// 创建两个单独的会话用来演示分布式锁竞争
	s1, err := concurrency.NewSession(cli, concurrency.WithTTL(3))
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "my-lock")

	s2, err := concurrency.NewSession(cli, concurrency.WithTTL(3))
	if err != nil {
		log.Fatal(err)
	}
	defer s2.Close()
	m2 := concurrency.NewMutex(s2, "my-lock")

	// 会话s1获取锁
	if err := m1.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("acquired lock for s1")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// 等待直到会话s1释放了/my-lock/的锁
		if err := m2.Lock(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	if err := m1.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("released lock for s1")

	<-m2Locked
	fmt.Println("acquired lock for s2")
}
