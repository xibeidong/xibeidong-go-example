package Decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	f := WithFoo(Foo)
	f("Hello")
}

func TestPi(t *testing.T) {
	fmt.Println(Pi(5000))
	fmt.Println(Pi(10000))
	fmt.Println(Pi(50000))
}
