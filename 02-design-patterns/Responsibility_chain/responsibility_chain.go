package Responsibility_chain

import "fmt"

/*
责任链（Chain of Responsibility）模式的定义：
为了避免请求发送者与多个请求处理者耦合在一起，
于是将所有请求的处理者通过前一对象记住其下一个对象的引用而连成一条链；
当有请求发生时，可将请求沿着这条链传递，直到有对象处理它为止。

模式的结构与实现
通常情况下，可以通过数据链表来实现职责链模式的数据结构。
1. 模式的结构
职责链模式主要包含以下角色。
抽象处理者（Handler）角色：定义一个处理请求的接口，包含抽象处理方法和一个后继连接。
具体处理者（Concrete Handler）角色：实现抽象处理者的处理方法，判断能否处理本次请求，如果可以处理请求则处理，否则将该请求转给它的后继者。
客户类（Client）角色：创建处理链，并向链头的具体处理者对象提交请求，它不关心处理细节和请求的传递过程。

责任链模式的本质是解耦请求与处理，让请求在处理链中能进行传递与被处理；理解责任链模式应当理解其模式，而不是其具体实现。责任链模式的独到之处是将其节点处理者组合成了链式结构，并允许节点自身决定是否进行请求处理或转发，相当于让请求流动起来。
*/

type Handle interface {
	HandleMsg(id int)
}

type Teacher struct {
	name string //老师
	next Handle
}

func (h *Teacher) HandleMsg(id int) {
	fmt.Println("i am teacher. msg id is ", id)
	if h.next != nil {
		h.next.HandleMsg(id)
	}
}

type Leader struct {
	name string //领导（系主任）
	next Handle
}

func (h *Leader) HandleMsg(id int) {
	fmt.Println("i am Leader. msg id is ", id)
	if h.next != nil {
		h.next.HandleMsg(id)
	}
}

type Dean struct {
	name string //院长
	next Handle
}

func (h *Dean) HandleMsg(id int) {
	fmt.Println("i am Dean. msg id is ", id)
	if h.next != nil {
		h.next.HandleMsg(id)
	}
}
