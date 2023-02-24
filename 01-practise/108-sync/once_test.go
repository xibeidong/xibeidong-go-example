package _08_sync

import (
	"fmt"
	"sync"
	"testing"
)

var num = 0

func add() {
	num++
}
func decrease() {
	num--
}
func TestOnce(t *testing.T) {
	once := sync.Once{}
	once.Do(add)
	once.Do(add)
	fmt.Println(num) // 1
	once.Do(decrease)
	once.Do(decrease)
	fmt.Println(num) // 1
}
