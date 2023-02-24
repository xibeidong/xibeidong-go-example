package bitwise

import (
	"log"
)

//  位运算

func OPBitwise() {
	maxUint32 := ^uint32(0) //对0取反得到该类型的最大值
	log.Println(1 << 32)    // 4294967296
	log.Println(maxUint32)  // 4294967295
}

// 取大于参数v的最小的且是2的幂的数 如 v=60 返回64； v=1025 返回2048
func roundUpToPower2(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v

	/*	输入1026
			10000000001  	1026-1 = 1025
		 	 1000000000
			11000000001
		 	  110000000
			11110000001
				1111000
			11111111001
					111
			11111111111 +1
		   100000000000
	*/

}

// 有一组数，其中有两种数出现了奇数次，其它的数出现偶数次，求这两种数a和b；
func getTwoOdd(arr []int) (a, b int) {
	eor := 0
	for i := 0; i < len(arr); i++ {
		eor ^= arr[i]
	}
	// eor = a^b
	// eor != 0
	// eor 必然有一个位置上是 1
	rightIsOne := getRightOne(eor) // 提取最右侧的1
	//两种奇数在rightIsOne处是不一样的,所以我们按照rightOne把所有数分成两组，一组有a,另一组有b
	for i := 0; i < len(arr); i++ {
		if arr[i]&rightIsOne == 0 {
			a ^= arr[i]
		}
	}
	b = a ^ eor
	return
}

// 返回最右一位是1的数，如10110，则返回 00010
func getRightOne(num int) int {
	// ^num 表示取反，golang中重用^符号，即可表示异或，也可表示取反
	//例如：
	// 10110 原数
	// 01001 取反
	// 01010 取反+1
	// 00010 原数&(取反+1)
	return num & (^num + 1)
}
