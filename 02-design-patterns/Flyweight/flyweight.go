package Flyweight

import "fmt"

/*
享元（Flyweight）模式的定义：运用共享技术来有效地支持大量细粒度对象的复用。
它通过共享已经存在的对象来大幅度减少需要创建的对象数量、避免大量相似类的开销，
从而提高系统资源的利用率。

享元模式的主要优点是：
相同对象只要保存一份，这降低了系统中对象的数量，
从而降低了系统中细粒度对象给内存带来的压力。

享元模式的本质是缓存共享对象，降低内存消耗。
1. 模式的结构
享元模式的主要角色有如下。
抽象享元角色（Flyweight）：是所有的具体享元类的基类，为具体享元规范需要实现的公共接口，非享元的外部状态以参数的形式通过方法传入。
具体享元（Concrete Flyweight）角色：实现抽象享元角色中所规定的接口。
非享元（Unsharable Flyweight)角色：是不可以共享的外部状态，它以参数的形式注入具体享元的相关方法中。
享元工厂（Flyweight Factory）角色：负责创建和管理享元角色。当客户对象请求一个享元对象时，享元工厂检査系统中是否存在符合要求的享元对象，如果存在则提供给客户；如果不存在的话，则创建一个新的享元对象。

连接池，对象池 使用享元模式
*/

//举例 下棋，棋子可以享元，棋子的位置是非享元

// ChessPiece 棋子 抽象享元角色
type ChessPiece interface {
	DownPiece(p *Point)
}

// Point 非享元角色
type Point struct {
	X, Y int
}

// WhitePiece 具体享元角色1
type WhitePiece struct {
	Color string
}

func (w *WhitePiece) DownPiece(p *Point) {
	fmt.Printf("white point is %v\n", *p)
}

// BlackPiece 具体享元角色2
type BlackPiece struct {
	Color string
}

func (b *BlackPiece) DownPiece(p *Point) {
	fmt.Printf("black point is %v\n", *p)
}

// PieceFactory 享元工厂
type PieceFactory struct {
	pieceMap map[string]ChessPiece
}

func NewPieceFactory() *PieceFactory {
	return &PieceFactory{pieceMap: map[string]ChessPiece{
		"white": &WhitePiece{Color: "white"},
		"black": &BlackPiece{Color: "black"},
	}}
}
func (f *PieceFactory) GetWhitePiece() *WhitePiece {
	if v, ok := f.pieceMap["white"]; ok {
		return v.(*WhitePiece)
	}
	return nil
}

func (f *PieceFactory) GetBlackPiece() *BlackPiece {
	if v, ok := f.pieceMap["black"]; ok {
		return v.(*BlackPiece)
	}
	return nil
}
