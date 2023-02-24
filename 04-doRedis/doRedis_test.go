package main

import "testing"

func TestSupportCommand(t *testing.T) {
	connectRedis()
	defer rdb.Close()
	supportCommand()
}

func TestUnSupportCommand(t *testing.T) {
	connectRedis()
	defer rdb.Close()
	unSupportCommand()
}

func TestSingleConn(t *testing.T) {
	connectRedis()
	defer rdb.Close()

	singleConn()
}

func TestConnCluster(t *testing.T) {
	client := connCluster()
	defer client.Close()
}
