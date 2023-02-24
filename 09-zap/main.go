package main

import (
	"fmt"
	"xibeidong-go-example/09-zap/zlog"
)

func main() {
	zlog.Ready("debug")

	zlog.Infof("info")
	zlog.Errorf("error")
	zlog.Warnf("warn")
	//add()
	fmt.Println("   end ")
	//zlog.Info("this is debug", zap.String("key1", "v2"))
}

func add() {
	for {
		zlog.Errorf("hello")
	}

}
