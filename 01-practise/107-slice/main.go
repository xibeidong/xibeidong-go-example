package main

import "fmt"

func main() {
	var s1, s2 []byte
	s1 = append(s1, s2...)
	s1 = append(s1, 0x01)

	n := byte(uint16(511)) // 运行时报错，511超出byte范围
	fmt.Println(n)
}
