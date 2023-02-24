package _04_map_question

import (
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
