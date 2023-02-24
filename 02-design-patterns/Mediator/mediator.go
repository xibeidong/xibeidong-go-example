package Mediator

import "fmt"

/*
中介者（Mediator）模式的定义：
定义一个中介对象来封装一系列对象之间的交互，使原有对象之间的耦合松散，且可以独立地改变它们之间的交互。
中介者模式又叫调停模式，它是迪米特法则的典型应用。

中介者模式实现的关键是找出“中介者”，下面对它的结构和实现进行分析。
1. 模式的结构
中介者模式包含以下主要角色。
抽象中介者（Mediator）角色：它是中介者的接口，提供了同事对象注册与转发同事对象信息的抽象方法。
具体中介者（Concrete Mediator）角色：实现中介者接口，定义一个 List 来管理同事对象，协调各个同事角色之间的交互关系，因此它依赖于同事角色。
抽象同事类（Colleague）角色：定义同事类的接口，保存中介者对象，提供同事对象交互的抽象方法，实现所有相互影响的同事类的公共功能。
具体同事类（Concrete Colleague）角色：是抽象同事类的实现者，当需要与其他同事对象交互时，由中介者对象负责后续的交互。
*/

// 举例 房产中介平台上的买卖双方通过中介交流

type Medium interface {
	Register(u ICustomer)  //注册
	Relay(uid, mag string) //中继
}

// EstateMedium 房产中介
type EstateMedium struct {
	users []ICustomer
}

func (e *EstateMedium) Register(u ICustomer) {
	e.users = append(e.users, u)
}
func (e *EstateMedium) Relay(uid, msg string) {
	for _, v := range e.users {
		if v.GetUid() != uid {
			v.Receive(msg)
		}

	}
}

type ICustomer interface {
	Send(msg string)
	Receive(msg string)
	GetUid() string
}

type Customer struct {
	name   string
	id     string
	medium Medium
}

type Buyer struct {
	Customer
	money int
}

func (b *Buyer) Send(msg string) {
	b.medium.Relay(b.id, msg)
}
func (b *Buyer) Receive(msg string) {

	fmt.Printf(" %v receive a msg : %v\n\n", b.name, msg)
}
func (b *Buyer) GetUid() string {
	return b.id
}

type Seller struct {
	Customer
	price     float32
	houseArea float32
}

func (b *Seller) Send(msg string) {
	b.medium.Relay(b.id, msg)
}
func (b *Seller) Receive(msg string) {
	fmt.Printf("%v receive a msg : %v\n\n", b.name, msg)
}
func (b *Seller) GetUid() string {
	return b.id
}
