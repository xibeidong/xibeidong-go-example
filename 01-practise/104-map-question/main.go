package main

import (
	"fmt"
	"reflect"
	"sync"
)

//下面代码有什么问题

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
func run03() {
	//：map线程安全
	//：可能会出现fatal error: concurrent map read and map write.
	//有可能发生同一个key，同时发生读和写，所以读也要加锁,修改如下,
	//或使用读写锁
	//func (ua *UserAges) Get(name string) int {
	//    ua.Lock()
	//    defer ua.Unlock()
	//    if age, ok := ua.ages[name]; ok {
	//        return age
	//    }
	//    return -1
	//}

}

type User struct {
}

func main() {
	dict := make(map[int]string)
	m := dict[12]
	if m == "" {
		fmt.Println("m是空string")
	}
	fmt.Println("m的类型是：", reflect.TypeOf(m))

	dict2 := make(map[int]*User)
	n := dict2[12]
	if n == nil {
		fmt.Println("n是nil")
	}
	fmt.Println("n的类型是：", reflect.TypeOf(n))

	fmt.Println(byte(511 & 0x00ff))

	var data2, data3 []byte
	data2 = append(data3, data2...)

	n2 := byte(uint16(257)) // 运行时报错，257超出byte范围
	fmt.Println(n2)
}
