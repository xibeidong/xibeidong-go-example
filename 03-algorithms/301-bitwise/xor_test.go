package bitwise

import (
	"log"
	"testing"
)

func TestOdd(t *testing.T) {
	//定义一个数组，其中6和7出现奇数次,其它数出现偶数次，
	arr := []int{1, 1, 1, 1, 2, 2, 6, 7, 7, 7, 11, 11}
	a, b := getTwoOdd(arr)
	log.Println(a, b)
}
