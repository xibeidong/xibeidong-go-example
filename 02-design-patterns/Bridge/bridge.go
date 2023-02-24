package Bridge

import "fmt"

/*
桥接模式
可以将抽象化部分与实现化部分分开，取消二者的继承关系，改用组合关系。
1. 模式的结构
桥接（Bridge）模式包含以下主要角色。
抽象化（Abstraction）角色：定义抽象类，并包含一个对实现化对象的引用。
扩展抽象化（Refined Abstraction）角色：是抽象化角色的子类，实现父类中的业务方法，并通过组合关系调用实现化角色中的业务方法。
实现化（Implementor）角色：定义实现化角色的接口，供扩展抽象化角色调用。
具体实现化（Concrete Implementor）角色：给出实现化角色接口的具体实现。
*/
// 咖啡为例，大中小杯，加糖加奶

type Coffee interface {
	Volume() //容量
}

type LargeCoffee struct {
	additives Additives
}

func (c *LargeCoffee) Volume() {
	fmt.Println("it is large")
}

type MediumCoffee struct {
	additives Additives
}

func (c *MediumCoffee) Volume() {
	fmt.Println("it is medium")
}

type SmallCoffee struct {
	additives Additives
}

func (c *SmallCoffee) Volume() {
	fmt.Println("it is small")
}

type Additives interface {
	AddSomething()
}
type Sugar struct {
}

func (s *Sugar) AddSomething() {
	fmt.Println("加糖")
}

type Milk struct {
}

func (m *Milk) AddSomething() {
	fmt.Println("加奶")
}
