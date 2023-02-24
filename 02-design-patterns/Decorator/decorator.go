package Decorator

import (
	"fmt"
	"math"
)

/*
装饰器（Decorator）模式的定义：指在不改变现有对象结构的情况下，
动态地给该对象增加一些职责（即增加其额外功能）的模式。

装饰器模式主要包含以下角色。
抽象构件（Component）角色：定义一个抽象接口以规范准备接收附加责任的对象。
具体构件（ConcreteComponent）角色：实现抽象构件，通过装饰角色为其添加一些职责。
抽象装饰（Decorator）角色：继承抽象构件，并包含具体构件的实例，可以通过其子类扩展具体构件的功能。
具体装饰（ConcreteDecorator）角色：实现抽象装饰的相关方法，并给具体构件对象添加附加的责任。
*/

// 举例，把Foo方法装饰起来，添加新功能，不破坏Foo的源代码

// Foo 具体构件
func Foo(msg string) bool {
	fmt.Println(msg)
	return true
}

type FooDecorator func(msg string) bool

// WithFoo 具体装饰
func WithFoo(fun FooDecorator) FooDecorator {
	return func(msg string) bool {
		fmt.Println("pre ")
		fun(msg)
		fmt.Println("after")
		return true
	}
}

// Pi 计算圆周率
func Pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k < n; k++ {
		go func(ch chan float64, k float64) {
			ch <- 4 * math.Pow(-1, k) / (2*k + 1)
		}(ch, float64(k))
	}
	result := 0.0
	for i := 0; i < n; i++ {
		result += <-ch
	}
	return result
}
