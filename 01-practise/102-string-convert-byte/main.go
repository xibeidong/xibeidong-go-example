package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//一般方法
	byte2Str()
	str2Byte()

	// 高效方法，类型转换，避免内存拷贝，
	// 注意：[]byte被修改，返回的string也会变
	StringToBytes("hello")
	var str string
	str = BytesToString([]byte{104, 101, 108, 108, 111})
	fmt.Println(str)
}
func byte2Str() {
	bt := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	fmt.Println(string(bt)) // 输出 hello world
}

func str2Byte() {
	str := "hello world"
	bt := []byte(str)
	fmt.Println(bt) // 输出 [104 101-exe-order 108 108 111 32 119 111 114 108 100]
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
