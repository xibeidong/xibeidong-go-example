package _09_reflect

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkSay(b *testing.B) {
	p := Person{name: "xiao"}
	for i := 0; i < b.N; i++ {
		p1 := &p
		p1.Say(context.Background(), "hello")
	}
}

func BenchmarkSayReflect(b *testing.B) {
	p := &Person{name: "xiao"}

	v := reflect.TypeOf(p)
	//num := v.NumMethod()
	//fmt.Println("method num = ", num)

	mt := v.Method(1)

	if mt.Name == "Say" {
		mt.Func.Call([]reflect.Value{reflect.ValueOf(p), reflect.ValueOf(context.Background()), reflect.ValueOf("hello")})
	} else {
		return
	}

	for i := 0; i < b.N; i++ {
		mt.Func.Call([]reflect.Value{reflect.ValueOf(p), reflect.ValueOf(context.Background()), reflect.ValueOf("hello")})
	}
}

func BenchmarkBye(b *testing.B) {
	p := Person{name: "xiao"}
	for i := 0; i < b.N; i++ {
		p1 := &p
		p1.Bye()
	}
}

func BenchmarkByeReflect(b *testing.B) {
	p := &Person{name: "xiao"}

	v := reflect.TypeOf(p)
	//num := v.NumMethod()
	//fmt.Println("method num = ", num)

	mt := v.Method(0)

	if mt.Name == "Bye" {
		mt.Func.Call([]reflect.Value{reflect.ValueOf(p)})
	}

	for i := 0; i < b.N; i++ {
		mt.Func.Call([]reflect.Value{reflect.ValueOf(p)})
	}
}

func TestSayOrBye(t *testing.T) {
	p := &Person{name: "xiao"}

	v := reflect.TypeOf(p)

	num := v.NumMethod()

	fmt.Println("method num = ", num)

	mt := v.Method(0)
	fmt.Println(mt.Name)
	if mt.Name == "Say" {
		mt.Func.Call([]reflect.Value{reflect.ValueOf(p), reflect.ValueOf(context.Background()), reflect.ValueOf("hello")})
	} else if mt.Name == "Bye" {
		mt.Func.Call([]reflect.Value{reflect.ValueOf(p)})
	}
}

//goos: windows
//goarch: amd64
//pkg: xibeidong-go-example/01-practise/109-reflect
//cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
//BenchmarkSay
//BenchmarkSay-8                  714830346                1.511 ns/op
//BenchmarkSayReflect
//BenchmarkSayReflect-8            1927776               619.8 ns/op
//BenchmarkBye
//BenchmarkBye-8                  751158661                1.571 ns/op
//BenchmarkByeReflect
//BenchmarkByeReflect-8            7198978               166.1 ns/op
//PASS
