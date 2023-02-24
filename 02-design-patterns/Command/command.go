package Command

import "fmt"

/*
命令模式包含以下主要角色。
抽象命令类（Command）角色：声明执行命令的接口，拥有执行命令的抽象方法 execute()。
具体命令类（Concrete Command）角色：是抽象命令类的具体实现类，它拥有接收者对象，并通过调用接收者的功能来完成命令要执行的操作。
实现者/接收者（Receiver）角色：执行命令功能的相关操作，是具体命令对象业务的真正实现者。
调用者/请求者（Invoker）角色：是请求的发送者，它通常拥有很多的命令对象，并通过访问命令对象来执行相关请求，它不直接访问接收者。
*/

// Command 抽象命令
type Command interface {
	Execute()
}

// ConcreteCommand1 具体命令，拥有接收者，并通过调用接收者的功能来完成命令要执行的操作。
type ConcreteCommand1 struct {
	receiver *Receiver
}

func (c *ConcreteCommand1) Execute() {
	c.receiver.Action1()
}

type ConcreteCommand2 struct {
	receiver *Receiver
}

func (c *ConcreteCommand2) Execute() {
	c.receiver.Action2()
}

//Receiver 命令接收者，有具体处理方法
type Receiver struct {
}

func (r *Receiver) Action1() {
	fmt.Println("action1 execute.")
}

func (r *Receiver) Action2() {
	fmt.Println("action2 execute")
}

type Invoker struct {
	command []Command
}

func (i *Invoker) addCommand(c Command) {
	i.command = append(i.command, c)
}
func (i *Invoker) Call() {
	for _, v := range i.command {
		v.Execute()
	}

}
