package main

import (
	"bytes"
	"fmt"
	"strings"
)

//string类型本质也是一个结构体，定义如下：
//type stringStruct struct {
//	str unsafe.Pointer
//	len int
//}

const base = "123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASFGHJKLZXCVBNM"

var baseSlice []string

func init() {
	for i := 0; i < 200; i++ {
		baseSlice = append(baseSlice, base)
	}
}
func main() {
	//综合对比性能排序：
	//strings.join ≈ strings.builder > bytes.buffer > []byte转换string > "+" > fmt.sprintf
	JoinString()
}

// SumString 原生拼接方式"+"
// 使用+操作符进行拼接时，会对字符串进行遍历，
// 计算并开辟一个新的空间来存储原来的两个字符串。
func SumString() string {
	res := ""
	for _, val := range baseSlice {
		res += val
	}
	return res
}

// SprintfString fmt.Sprintf实现原理主要是使用到了反射,性能最差
func SprintfString() string {
	res := ""
	for _, val := range baseSlice {
		res = fmt.Sprintf("%s%s", res, val)
	}
	return res
}

// BuilderString 官方推荐strings.Builder
func BuilderString() string {
	var builder strings.Builder

	//Grow方法提前分配slice大小，避免了slice扩容
	builder.Grow(200 * len(baseSlice))

	for _, val := range baseSlice {
		builder.WriteString(val)
	}

	//String方法就是将[]byte转换为string类型，
	//这里为了避免内存拷贝的问题，使用了强制转换来避免内存拷贝：*(*string)(unsafe.Pointer(&b.buf))
	return builder.String()
}

// bytesString bytes.Buffer
func bytesString() string {
	buf := new(bytes.Buffer)
	for _, val := range baseSlice {
		buf.WriteString(val)
	}
	// String方法没有使用强制转换，导致内存拷贝 string(b.buf[b.off:])
	return buf.String()
}

// byteSliceString append []byte 转 string
func byteSliceString() string {
	buf := make([]byte, 0)
	for _, val := range baseSlice {
		buf = append(buf, val...)
	}
	// 如果想减少内存分配，在将[]byte转换为string类型时可以考虑使用强制转换。
	//return *(*string)(unsafe.Pointer(&buf))
	return string(buf)
}

// JoinString strings.join将一个string类型的切片拼接成一个字符串，可以定义连接操作符
// 是基于strings.builder来实现的
// 因为提前进行容量分配可以减少内存分配，所以很高效
func JoinString() string {
	return strings.Join(baseSlice, "")
}
