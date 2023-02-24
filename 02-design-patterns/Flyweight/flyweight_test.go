package Flyweight

import (
	"fmt"
	"testing"
)

func TestPiece(t *testing.T) {
	factory := NewPieceFactory()

	b := factory.GetBlackPiece()
	b.DownPiece(&Point{
		X: 46,
		Y: 0,
	})

	b2 := factory.GetBlackPiece()
	fmt.Println(b == b2) //true

	w := factory.GetWhitePiece()
	w.DownPiece(&Point{X: 12, Y: 25})
	w2 := factory.GetWhitePiece()
	fmt.Println(w == w2) //true
	fmt.Printf("%p,%p \n", w, w2)
}
