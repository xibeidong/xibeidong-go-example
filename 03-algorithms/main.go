package main

import (
	"log"
	bitwise "xibeidong-go-example/03-algorithms/301-bitwise"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6}
	arr2 := arr[0:2]
	log.Println(arr2)

	bitwise.OPBitwise()

}
