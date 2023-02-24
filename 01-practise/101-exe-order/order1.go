package main

import (
	"fmt"
	"sync/atomic"
)

type T struct {
	x, y int
}

// 多核处理中，代码不一定是顺序执行的
func test3() {
	a := T{
		x: 1,
		y: 2,
	}
	b := T{
		x: 3,
		y: 4,
	}
	t := T{}
	go func() {
		// 使用 go build -race
		// 不要使用编译优化，不然for{}里面的赋值会被优化掉
		for {
			t = b
		}
	}()
	go func() {
		for {
			t = a
		}
	}()
	for {
		fmt.Println(t)
		//最后结果 ｛1 2｝ ｛1 4｝ ｛3 2｝ ｛3 4｝
	}
}
func test4() {
	v := atomic.Value{}
	a := T{
		x: 1,
		y: 2,
	}
	b := T{
		x: 3,
		y: 4,
	}

	v.Store(T{})
	go func() {
		for {
			v.Store(a)
		}
	}()
	go func() {
		for {
			v.Store(b)
		}
	}()
	for {
		fmt.Println(v.Load())
		//最后结果｛1 2｝ ｛3 4｝
	}
}
