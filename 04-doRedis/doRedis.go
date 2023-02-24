package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// for more:
//	https://redis.uptrace.dev/guide/
var rdb *redis.Client
var ctx context.Context

func main() {

	fmt.Println("doRedis")
	connectRedis()
	defer rdb.Close()
	cmd := rdb.Set(ctx, "k1", "v11", 0)
	result, err := cmd.Result()
	if err != nil {
		fmt.Println("err ", err)
	}
	fmt.Println(result)

	s, err := rdb.Get(ctx, "k2").Result()
	if err != nil {
		fmt.Println("err ", err)
	}
	fmt.Println(s)

}

func connectRedis() {
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.17.129:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	result, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

// 执行go-redis定义的command
func supportCommand() {
	//expire = 0 表示不过期
	result, _ := rdb.Set(ctx, "u1", "jack", 0).Result()
	fmt.Println("set u1 result: ", result)
	s, _ := rdb.Get(ctx, "u1").Result()
	fmt.Println("get u1: ", s)
}

// 执行go-redis未定义的command
func unSupportCommand() {
	val, err := rdb.Do(ctx, "get", "u1").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("u1 : ", "key does not exist")
	case err != nil:
		fmt.Println("Get failed", err)
	case val == "":
		fmt.Println("value is empty")
	}
	fmt.Println(val)
}

// 获取一个单独的连接，去执行commands
func singleConn() {
	cn := rdb.Conn(ctx)
	defer cn.Close()

	if err := cn.ClientSetName(ctx, "myclient").Err(); err != nil {
		panic(err)
	}

	name, err := cn.ClientGetName(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("client name", name)
}

// cluster模式,最低3主3从
func connCluster() *redis.ClusterClient {
	ctx = context.Background()
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"192.168.17.129:6379",
			"192.168.17.129:6380",
			"192.168.17.129:6381",
			"192.168.17.129:6382",
			"192.168.17.129:6383",
			"192.168.17.129:6384",
		},
	})

	err := rdb.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		fmt.Println(err)
	}
	//To iterate over master nodes, use ForEachMaster.
	//To iterate over slave nodes, use ForEachSlave.
	return rdb
}

// 哨兵模式
func connSentinel() *redis.Client {
	//通过哨兵集群获取redis-master连接
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master-name",
		SentinelAddrs: []string{":9126", ":9127", ":9128"},
	})

	// 连接一个哨兵
	//To connect to a Redis Sentinel itself:
	sentinel := redis.NewSentinelClient(&redis.Options{
		Addr: ":9126",
	})
	addr, err := sentinel.GetMasterAddrByName(ctx, "master-name").Result()
	fmt.Println(addr, err)

	return rdb
}

//Ring模式，？？？
func connRing() *redis.Ring {
	rdb := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"shard1": ":7000",
			"shard2": ":7001",
			"shard3": ":7002",
		},
	})
	if err := rdb.Set(ctx, "foo", "bar", 0).Err(); err != nil {
		fmt.Println(err)
	}
	return rdb
}

//通用连接，根据不通配置返回
//NewUniversalClient returns a new multi client. The type of the returned client depends on the following conditions:
//
//	1.If the MasterName option is specified, a sentinel-backed FailoverClient is returned.
//	2.if the number of Addrs is two or more, a ClusterClient is returned.
//	3.Otherwise, a single-node Client is returned.
func connUniversal() redis.UniversalClient {
	// rdb is *redis.Client.
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{":6379"},
	})

	// rdb is *redis.ClusterClient.
	rdb = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{":6379", ":6380"},
	})

	// rdb is *redis.FailoverClient.
	rdb = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      []string{":6379"},
		MasterName: "mymaster",
	})

	return rdb
}
