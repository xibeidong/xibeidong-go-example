package main

import (
	"testing"
)

func BenchmarkSumString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumString()
	}
	b.StopTimer()
}

func BenchmarkSprintfString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SprintfString()
	}
	b.StopTimer()
}

func BenchmarkBuilderString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BuilderString()
	}
	b.StopTimer()
}

func BenchmarkBytesBufferString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytesString()
	}
	b.StopTimer()
}

func BenchmarkJoinstring(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JoinString()
	}
	b.StopTimer()
}

func BenchmarkByteSliceString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		byteSliceString()
	}
	b.StopTimer()
}

/*
结果：strings.join > strings.builder > bytes.buffer > []byte转换string > "+" > fmt.sprintf

BenchmarkSumString
BenchmarkSumString-8                2050            507542 ns/op
BenchmarkSprintfString
BenchmarkSprintfString-8            1716            614132 ns/op
BenchmarkBuilderString
BenchmarkBuilderString-8           85262             14359 ns/op
BenchmarkBytesBufferString
BenchmarkBytesBufferString-8       56900             21345 ns/op
BenchmarkJoinstring
BenchmarkJoinstring-8             143372              8080 ns/op
BenchmarkByteSliceString
BenchmarkByteSliceString-8         49581             24326 ns/op
PASS

Process finished with the exit code 0

*/
