package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func NewClient(endPoints []string) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   endPoints,
		DialTimeout: time.Second * 5,
	})
}

func put(client *clientv3.Client, key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.Put(ctx, key, value)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("put ok, cluster_id: ", response.Header.ClusterId)

	//done := ctx.Done()
	//<-done
	return nil
}

func get(client *clientv3.Client, key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.Get(ctx, key)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, v := range response.Kvs {
		fmt.Println(string(v.Value))
	}
	fmt.Println(key, " len = ", len(response.Kvs))
	return nil
}
func watch(cli *clientv3.Client, key string) {
	watchChan := cli.Watch(context.Background(), key)
	for resp := range watchChan {
		for _, ev := range resp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
