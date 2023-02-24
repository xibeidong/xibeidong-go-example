package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	count := runtime.GOMAXPROCS(0)
	fmt.Println("max process = ", count)

	test1()

	//test3()

	//使用atomic，效率降低，但是安全
	//test4()
}

// 1、go只保证同一个协程内的执行顺序不受CPU指令重排的影响。
func test1() {
	var count int64 = 0
	w := sync.WaitGroup{}
	for {
		a, b, x, y := 0, 0, 0, 0
		count++

		w.Add(2)
		go func() {
			a = 1
			y = b
			w.Done()
		}()
		go func() {
			b = 1
			x = a
			w.Done()
		}()
		w.Wait()
		//有可能发生，虽然概率不高，
		//go只保证了一个goroutine 内部重排对读写的顺序没有影响；
		//其它goroutine看到的顺序是不一定的
		if x == 0 && y == 0 {
			fmt.Println(count)
			break
		}

	}

}
