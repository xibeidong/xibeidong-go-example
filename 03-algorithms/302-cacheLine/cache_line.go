package main

import (
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type T struct {
	d1 int64
	_  [7]int64 // 内存对齐
	d2 int64
	_  [7]int64

	//pad [128 - 8%128]byte
}

var count int64 = 10_0000_0000

var data T

func main() {
	log.Println(runtime.GOMAXPROCS(0))

	log.Println(unsafe.Sizeof(T{}))
	log.Println(unsafe.Alignof(T{}))
	w := sync.WaitGroup{}
	w.Add(2)
	t1 := time.Now().UnixNano()
	go func() {
		var i = count
		for ; i > 0; i-- {
			atomic.AddInt64(&data.d1, 1)
		}
		w.Done()
	}()
	go func() {
		var i = count
		for ; i > 0; i-- {
			atomic.AddInt64(&data.d2, 1)
		}
		w.Done()
	}()
	w.Wait()
	t2 := time.Now().UnixNano()
	log.Println(data)
	log.Println(t2 - t1)

	// 同样执行10亿次
	// 内存对齐后，11s
	// 不对齐， 50-70s
}
