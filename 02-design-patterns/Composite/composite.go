package Composite

import "fmt"

/*
组合（Composite Pattern）模式的定义：
有时又叫作整体-部分（Part-Whole）模式，
它是一种将对象组合成树状的层次结构的模式，
用来表示“整体-部分”的关系，使用户对单个对象和组合对象具有一致的访问性，
属于结构型设计模式。
组合模式一般用来描述整体与部分的关系，它将对象组织到树形结构中，顶层的节点被称为根节点，
根节点下面可以包含树枝节点和叶子节点，树枝节点下面又可以包含树枝节点和叶子节点

组合模式包含以下主要角色。
抽象构件（Component）角色：它的主要作用是为树叶构件和树枝构件声明公共接口，并实现它们的默认行为。在透明式的组合模式中抽象构件还声明访问和管理子类的接口；在安全式的组合模式中不声明访问和管理子类的接口，管理工作由树枝构件完成。（总的抽象类或接口，定义一些通用的方法，比如新增、删除）
树叶构件（Leaf）角色：是组合中的叶节点对象，它没有子节点，用于继承或实现抽象构件。
树枝构件（Composite）角色 / 中间构件：是组合中的分支节点对象，它有子节点，用于继承和实现抽象构件。它的主要作用是存储和管理子部件，通常包含 Add()、Remove()、GetChild() 等方法。

*/
//举例，小包可以放商品，大包可以放小包和商品， 计算商品总价。
// 商品是leaf，bag是树枝，树枝上可以有leaf和新的树枝。

// Article 商品 抽象构件
type Article interface {
	Calculation() float32 // 计算
	Show()
}

//Goods 商品 树叶构件
type Goods struct {
	name      string
	unitPrice float32 //单价
	quantity  int     //数量
}

func (g *Goods) Calculation() float32 {
	return g.unitPrice * float32(g.quantity)
}
func (g *Goods) Show() {
	fmt.Printf("%+v \n", *g)
}

//Bags 包 树枝构件
type Bags struct {
	name string
	bags []Article
}

func (b *Bags) Calculation() float32 {
	var p float32 = 0
	for _, v := range b.bags {
		p += v.Calculation()
	}
	return p
}
func (b *Bags) Show() {
	fmt.Printf("%+v \n", *b)
}

func (b *Bags) Add(art Article) {
	b.bags = append(b.bags, art)
}
