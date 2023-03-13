package base128varints

// base 128 varints 压缩数字的方法，原理是小数用的多
// https://blog.csdn.net/mweibiao/article/details/83036427

//1byte = 8bit
//第一位表示后面是否还有值，1有值，0没有值，后面的7个bit是一组，代表数值
//byte要倒过来排列，低位放前面
//4个字节算的话，最大能表示32-4 = 28位，即最大表示数 2^28

//golang 默认是大端；TCP协议规定用大端，就是我们最习惯的方式

func EncodeUint32(n uint32) (data []byte) {
	var byteMax uint32 = 1 << 7 // 一个字节最大能表示128，不是256了
	for {
		if n>>7 != 0 {
			data = append(data, byte(n%byteMax+byteMax)) //第一位是 1
		} else {
			data = append(data, byte(n%byteMax)) // 第一位是 0
			break
		}
		n = n >> 7
	}
	return
}

// DecodeUint32
// start 开始位置，
// n 结果， l 占用字节长度
func DecodeUint32(data []byte, start int) (n uint32, l int) {

	for i := start; i < len(data); i++ {
		l++
		if data[i] < 128 {
			n += uint32(data[i]) << (7 * (l - 1))
			break
		}
		temp := data[i]
		temp &= 0x7F //  data[i] &= 0111 1111
		n += uint32(temp) << (7 * (l - 1))

	}
	return
}
