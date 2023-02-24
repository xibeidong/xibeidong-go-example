package Adapter

import "fmt"

/*
适配器模式（Adapter）的定义如下：
将一个类的接口转换成客户希望的另外一个接口，
使得原本由于接口不兼容而不能一起工作的那些类能一起工作。

适配器模式（Adapter）包含以下主要角色。
目标（Target）接口：当前系统业务所期待的接口，它可以是抽象类或接口。
适配者（Adaptee）类：它是被访问和适配的现存组件库中的组件接口。
适配器（Adapter）类：它是一个转换器，通过继承或引用适配者的对象，把适配者接口转换成目标接口，让客户按目标接口的格式访问适配者。
*/

// ElectricEngine 电机 适配者1
type ElectricEngine struct {
}

func (e *ElectricEngine) ElectricDrive() {
	fmt.Println("电力发动机...")
}

// OilEngine 油机 适配者2
type OilEngine struct {
}

func (o *OilEngine) OilDrive() {
	fmt.Println("燃油机...")
}

// Motor 目标接口
type Motor interface {
	Driver()
}

// ElectricAdapter 适配器1
type ElectricAdapter struct {
	elect ElectricEngine
}

func (e *ElectricAdapter) Driver() {
	e.elect.ElectricDrive()
}

// OilAdapter 适配器2
type OilAdapter struct {
	oil OilEngine
}

func (e *OilAdapter) Driver() {
	e.oil.OilDrive()
}
