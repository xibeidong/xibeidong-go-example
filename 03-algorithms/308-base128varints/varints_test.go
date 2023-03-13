package base128varints

import (
	"fmt"
	"testing"
)

func TestUint(t *testing.T) {
	var n uint32 = 13
	n = n >> 7
	fmt.Println(n) // 0
}

func TestEncodeUint32(t *testing.T) {
	// 45689 => f9e402
	n := uint32(45689)
	data := EncodeUint32(n)
	fmt.Printf("%x \n", data)

}

func TestDecodeUint32(t *testing.T) {
	data := []byte{0xff, 0xf9, 0xe4, 0x02, 0xf1}

	n, l := DecodeUint32(data, 1)
	fmt.Println(n, l)
}
