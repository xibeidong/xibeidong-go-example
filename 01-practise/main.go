package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var a uint8 = 255
	//fmt.Println(a + 200)

	//Push(false)

	typeOfError := reflect.TypeOf((*error)(nil)).Elem()
	typeOfBytes := reflect.TypeOf(([]byte)(nil))
	fmt.Println(typeOfError)
	fmt.Println(typeOfBytes)

}

func testSwitch() {
	s := 1
	switch s {
	case 1, 2:
		fmt.Println("it is 1 or 2.")
	case 3:
		fmt.Println("it is 3")
	default:
		fmt.Println("it is nothing")
	}
}

func Push(isError ...bool) {
	fmt.Println(isError)
}
