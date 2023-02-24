package main

import "fmt"

type Dog struct {
	name string
	age  int
}

func main() {
	one()
	println(two1())  // 20
	println(two2())  // 15
	println(*two3()) // 15
}

func one() {
	dog1 := &Dog{
		name: "dog1",
		age:  1,
	}
	dog2 := *dog1 //复制一份给dog2
	dog2.name = "dog2"
	fmt.Printf("dog1: %p %#v\n", &dog1, dog1)
	fmt.Printf("dog2: %p %#v\n", &dog2, dog2)

}

// 声明了返回值变量，在defer中修改成功
func two1() (i int64) {
	i = 10
	defer func() {
		i = 20
	}()
	i = 15
	return
}

// 没声明返回值变量，在defer中修改失败，
func two2() int {
	i := 10
	defer func() {
		i = 20
	}()
	i = 15
	return i
}

// 指针在任何地方修改都成功
func two3() *int {
	i := 10
	defer func() {
		i = 20
	}()
	i = 15
	return &i
}

//
//对于 defer 一种比较好的理解方式，
//就是假设 Golang 的函数需有一个“真正”的返回值变量。

//如果在函数的返回值列表中做了声明，
//那么这个“真正”的返回值变量就等同于我们定义的返回值变量，
//一切对于返回值的修改都是“有效”的；

//如果没有声明，
//那么 Golang 会自动创建一个“匿名”的返回值变量，
//在“主函数”执行完成后，将返回值赋值给“匿名”的返回值变量，
//此时再对返回值变量修改就是无效的了，
//因为已经不会再次赋值。
