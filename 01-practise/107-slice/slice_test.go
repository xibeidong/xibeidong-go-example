package main

import "testing"

type Packet struct {
	typ    byte
	length int
	body   []byte
}

func NewDefaultPacket() *Packet {
	l := 1024
	data := make([]byte, l)
	for i := 0; i < l; i++ {
		data[i] = 0xf1
	}
	return &Packet{
		typ:    0x04,
		length: l,
		body:   data,
	}
}

func (p *Packet) encode0() []byte {
	data := make([]byte, 1)
	data[0] = p.typ
	data = append(data, []byte{byte((p.length >> 16) & 0xff), byte((p.length >> 8) & 0xff), byte(p.length)}...)
	data = append(data, p.body...)
	return data
}

func (p *Packet) encode1() []byte {
	data := make([]byte, 1, p.length+4)
	data[0] = p.typ
	data = append(data, []byte{byte((p.length >> 16) & 0xff), byte((p.length >> 8) & 0xff), byte(p.length)}...)
	data = append(data, p.body...)
	return data
}

func (p *Packet) encode2() []byte {
	data := make([]byte, p.length+4)
	data[0] = p.typ
	copy(data[1:4], []byte{byte((p.length >> 16) & 0xff), byte((p.length >> 8) & 0xff), byte(p.length)})
	copy(data[4:], p.body)
	return data
}

func BenchmarkSlice0(b *testing.B) {
	p := NewDefaultPacket()
	for i := 0; i < b.N; i++ {
		p.encode0()
	}
}

func BenchmarkSlice1(b *testing.B) {
	p := NewDefaultPacket()
	for i := 0; i < b.N; i++ {
		p.encode1()
	}
}
func BenchmarkSlice2(b *testing.B) {
	p := NewDefaultPacket()
	for i := 0; i < b.N; i++ {
		p.encode2()
	}
}

// l = 277
//BenchmarkSlice0
//BenchmarkSlice0-8        7036789               152.9 ns/op
//BenchmarkSlice1
//BenchmarkSlice1-8        9684854               132.2 ns/op
//BenchmarkSlice2
//BenchmarkSlice2-8       10026837               122.0 ns/op

// l = 1024
//cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
//BenchmarkSlice0
//BenchmarkSlice0-8        2540698               458.4 ns/op
//BenchmarkSlice1
//BenchmarkSlice1-8        2999148               434.7 ns/op
//BenchmarkSlice2
//BenchmarkSlice2-8        2682819               439.3 ns/op

// l = 277011
//BenchmarkSlice0
//BenchmarkSlice0-8          23157             55486 ns/op
//BenchmarkSlice1
//BenchmarkSlice1-8          20959             58964 ns/op
//BenchmarkSlice2
//BenchmarkSlice2-8          17768             61400 ns/op
