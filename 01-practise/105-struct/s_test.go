package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestAddr(t *testing.T) {
	d := Dog{
		name: "xx",
		age:  12,
	}
	d1 := d
	fmt.Printf("d addr = %p\n", &d)
	fmt.Printf("d1 addr = %p\n", &d1)

}

type Cat struct {
	age  int
	name string
}

func (c *Cat) Eat() {
	fmt.Println("eat...")
}

func TestFunc(t *testing.T) {
	dog := Dog{
		name: "1",
		age:  2,
	}
	cat := Cat{
		name: "1",
		age:  2,
	}
	fmt.Println(unsafe.Sizeof(cat)) //24
	fmt.Println(unsafe.Sizeof(dog)) //24

	fmt.Println(unsafe.Alignof(cat)) //8
	fmt.Println(unsafe.Alignof(dog)) //8

}

func TestEQ(t *testing.T) {
	d1 := Dog{
		name: "1",
		age:  2,
	}
	d2 := Dog{
		name: "1",
		age:  2,
	}
	if d1 == d2 {
		fmt.Println(" it is equal") // yes
	}
}
