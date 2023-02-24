package main

import (
	"fmt"
	"github.com/dop251/goja"
	"log"
)

func main() {
	//getValueFromJS()
	//fmt.Println("-----------------")
	//sendValueToJS()
	//doArray()
	doArray2()
}
func getValueFromJS() {
	vm := goja.New()
	vm.RunString(`
var name = "hjn";
var age = 30;
function f1(param){
	return +param+2
}
`)
	var name string
	err := vm.ExportTo(vm.Get("name"), &name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	var age int
	err = vm.ExportTo(vm.Get("age"), &age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(age)

	var f1 func(int) int
	err = vm.ExportTo(vm.Get("f1"), &f1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f1(40))
}

func getValueFromJS2() {
	vm := goja.New()
	vm.RunString("")
}

func sendValueToJS() {

	vm := goja.New()
	vm.RunString(`
	let x,y;
	function f1(){
return x+y
}
`)
	if err := vm.Set("x", 10); err != nil {
		log.Fatal(err)
	}
	if err := vm.Set("y", 70); err != nil {
		log.Fatal(err)
	}

	var f1 func() int
	vm.ExportTo(vm.Get("f1"), &f1)
	fmt.Println(f1())
}

func doArray() {
	vm := goja.New()
	array := make([]int, 10)

	for i := 0; i < 10; i++ {
		array[i] = i
	}

	vm.RunString(`
	let arr;
	function f1(){
let sum = 0;
		for(let i=0;i<10;i++){
sum+=arr[i]
	arr[i]=i+1000
}
return sum
	}
`)
	if err := vm.Set("arr", &array); err != nil {
		log.Fatal(err)
	}

	var f1 func() int
	err := vm.ExportTo(vm.Get("f1"), &f1)
	if err != nil {
		log.Fatal(err)
	}
	a := f1()
	fmt.Println(a)

	err = vm.ExportTo(vm.Get("arr"), &array)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(array)
}

func doArray2() {
	vm := goja.New()
	array := make([][]int, 10)
	num := 0
	for i := 0; i < 10; i++ {
		array[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			array[i][j] = num
			num++
		}
	}

	vm.RunString(`
	let arr;
	function f1(){
let sum = 0;
		for(let i=0;i<arr.length;i++){
sum+=arr[i][0]
	arr[i][0]=i+1000
}
return sum
	}
`)
	if err := vm.Set("arr", &array); err != nil {
		log.Fatal(err)
	}

	var f1 func() int
	err := vm.ExportTo(vm.Get("f1"), &f1)
	if err != nil {
		log.Fatal(err)
	}
	a := f1()
	fmt.Println(a)

	err = vm.ExportTo(vm.Get("arr"), &array)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(array)
}
