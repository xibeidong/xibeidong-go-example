package main

import (
	"fmt"
	"testing"
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
