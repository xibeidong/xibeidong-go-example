package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"unsafe"
)

func TestSwapPointer(t *testing.T) {
	n1 := &node{
		size: 12,
		name: "12",
	}
	n2 := &node{
		size: 34,
		name: "34",
	}
	fmt.Printf("n1 %p %v\n", &n1, n1)
	fmt.Printf("n2 %p %v\n", &n2, n2)

	// 注意参数，第一个是指针的指针，第二个是指针
	old := atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&n1)), unsafe.Pointer(n2))

	fmt.Println("--------------------")

	fmt.Printf("n1 %p %v\n", &n1, n1)
	fmt.Printf("n2 %p %v\n", &n2, n2)
	n := (*node)(old) // 返回的是新地址，和n1、n2不同
	fmt.Println(n)
	fmt.Printf("old %p %v\n", &n, n)

	//=== RUN   TestSwapPointer
	//n1 0xc0000c4028 &{12 12}
	//n2 0xc0000c4030 &{34 34}
	//--------------------
	//n1 0xc0000c4028 &{34 34}
	//n2 0xc0000c4030 &{34 34}
	//&{12 12}
	//old 0xc0000c4038 &{12 12}
	//--- PASS: TestSwapPointer (0.00s)
	//PASS
}
