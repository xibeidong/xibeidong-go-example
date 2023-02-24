package _06_my_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkPool(b *testing.B) {
	var num int32 = 0
	pool := &sync.Pool{New: func() interface{} {
		// New 方法里存在数据race，要用原子操作，统计调用次数
		atomic.AddInt32(&num, 1)
		buffer := make([]byte, 1024)
		// 用&buffer（速度更快），不用buffer（速度慢）
		return &buffer
	}}
	// 速度降低几十倍，一次操作30ns，内存消耗变得很低，属于时间换空间
	for i := 0; i < b.N; i++ {
		data := pool.Get().(*[]byte)
		//d:=*data
		//pool.Put(d[:10]) 必须把原来的1024长度的slice放回去，不能是修改过len的
		pool.Put(data)
	}
	fmt.Println(" num= ", num)
}

func BenchmarkSliceBytes(b *testing.B) {
	//速度很快，一次操作不到 1ns 内存占用很高，
	for i := 0; i < b.N; i++ {
		data := make([]byte, 1024)
		data[500] = 0x02
	}
}
