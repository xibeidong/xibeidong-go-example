package Builder

/*
标准 builder有四部分 ==>
• 抽象建造者角色（Builder）：
	为创建一个Product对象的各个部件指定抽象接口，以规范产品对象的各个组成成分的建造。
	一般而言，此角色规定要实现复杂对象的哪些部分的创建，并不涉及具体的对象部件的创建。

• 具体建造者（ConcreteBuilder）
	1）实现Builder的接口以构造和装配该产品的各个部件。即实现抽象建造者角色Builder的方法。
	2）定义并明确它所创建的表示，即针对不同的商业逻辑，具体化复杂对象的各部分的创建
	3）提供一个检索产品的接口
	4）构造一个使用Builder接口的对象即在指导者的调用下创建产品实例

• 指导者（Director）：
	调用具体建造者角色以创建产品对象的各个部分。指导者并没有涉及具体产品类的信息，真正拥有具体产品的信息是具体建造者对象。它只负责保证对象各部分完整创建或按某种顺序创建。

• 产品角色（Product）：
	建造中的复杂对象。它要包含那些定义组件的类，包括将这些组件装配成产品的接口。

https://blog.csdn.net/hguisu/article/details/7518060
*/

// Builder 抽象建造者
type Builder interface {
	GetProduct() CarProduct
	BuildPart1(engine string)
	BuildPart2(rack string)
}

// CarProduct 产品
type CarProduct struct {
	Engine string
	Rack   string
}

// ConcreteBuilder1 具体建造者1
type ConcreteBuilder1 struct {
	CarProduct
}

func (b *ConcreteBuilder1) BuildPart1(engine string) {
	b.Engine = engine
}

func (b *ConcreteBuilder1) BuildPart2(rack string) {
	b.Rack = rack
}

func (b *ConcreteBuilder1) GetProduct() CarProduct {
	return CarProduct{
		Engine: b.Engine,
		Rack:   b.Rack,
	}
}
func NewConcreteBuilder1() *ConcreteBuilder1 {
	return &ConcreteBuilder1{}
}

// Director 指导者
type Director struct {
}

func (d *Director) BuildCar(b Builder) {
	b.BuildPart1("engine2")
	b.BuildPart2("rack2")
}

func NewDirector() *Director {
	return &Director{}
}
