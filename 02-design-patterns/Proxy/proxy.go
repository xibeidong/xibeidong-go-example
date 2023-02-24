package Proxy

import "fmt"

/*
代理模式的定义：由于某些原因需要给某对象提供一个代理以控制对该对象的访问。
这时，访问对象不适合或者不能直接引用目标对象，代理对象作为访问对象和目标对象之间的中介。

代理模式的主要角色如下。
抽象主题（Subject）类：通过接口或抽象类声明真实主题和代理对象实现的业务方法。
真实主题（Real Subject）类：实现了抽象主题中的具体业务，是代理对象所代表的真实对象，是最终要引用的对象。
代理（Proxy）类：提供了与真实主题相同的接口，其内部含有对真实主题的引用，它可以访问、控制或扩展真实主题的功能。
*/

//RentHouse 抽象主题
type RentHouse interface {
	Rent(price int)
}

// RentHouseCity 真实主题
type RentHouseCity struct {
}

func (r *RentHouseCity) Rent(price int) {
	fmt.Printf("花费%v元可以在城市租房子.\n", price)
}

// RentHouseProxy 代理
type RentHouseProxy struct {
	rentInCity *RentHouseCity
}

func (r *RentHouseProxy) Rent(price int) {
	r.PreRent()
	r.rentInCity.Rent(price)
	r.AfterRent()
}

func (r *RentHouseProxy) PreRent() {
	fmt.Println("先打扫一遍。")
}
func (r *RentHouseProxy) AfterRent() {
	fmt.Println("定期收租。")
}
