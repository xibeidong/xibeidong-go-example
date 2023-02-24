package discovery

import (
	"log"
	"testing"
	"time"
)

func TestServiceDiscovery(t *testing.T) {
	var endpoints = []string{"192.168.254.128:2379"}
	ser := NewServiceDiscovery(endpoints)
	defer ser.Close()
	ser.WatchService("/web/")
	ser.WatchService("/gRPC/")
	for {
		select {
		case <-time.Tick(10 * time.Second):
			log.Println(ser.GetServices())
		}
	}
}
