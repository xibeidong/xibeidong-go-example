package State

import "fmt"

/*
状态（State）模式的定义：
对有状态的对象，把复杂的“判断逻辑”提取到不同的状态对象中，
允许状态对象在其内部状态发生改变时改变其行为。

状态模式把受环境改变的对象行为包装在不同的状态对象里，其意图是让一个对象在其内部状态改变的时候，其行为也随之改变。现在我们来分析其基本结构和实现方法。
1. 模式的结构
状态模式包含以下主要角色。
环境类（Context）角色：也称为上下文，它定义了客户端需要的接口，内部维护一个当前状态，并负责具体状态的切换。
抽象状态（State）角色：定义一个接口，用以封装环境对象中的特定状态所对应的行为，可以有一个或多个行为。
具体状态（Concrete State）角色：实现抽象状态所对应的行为，并且在需要的情况下进行状态切换。
*/

type TVState interface {
	turnOn(tv *BigTV)
	turnOff(tv *BigTV)
	nextChannel(tv *BigTV)
}

// powerOnState 开机状态
type powerOnState struct {
}

func (s *powerOnState) turnOn(tv *BigTV) {
	fmt.Println("已经打开了。")
}
func (s *powerOnState) turnOff(tv *BigTV) {
	fmt.Println("TV off.")
	tv.setState(&powerOffState{})
}
func (s *powerOnState) nextChannel(tv *BigTV) {
	fmt.Println("switch next channel.")
}

// powerOffState 关机状态
type powerOffState struct {
}

func (s *powerOffState) turnOn(tv *BigTV) {
	fmt.Println("TV 打开了。")
	tv.setState(&powerOnState{})
}
func (s *powerOffState) turnOff(tv *BigTV) {
	fmt.Println("TV 是关闭的，不用再关闭了.")

}
func (s *powerOffState) nextChannel(tv *BigTV) {
	fmt.Println("can not switch next channel.")
}

// BigTV 环境类 它定义了客户端需要的接口，内部维护一个当前状态，并负责具体状态的切换
type BigTV struct {
	currentState TVState
}

func (tv *BigTV) setState(state TVState) {
	tv.currentState = state
}

func (tv *BigTV) open() {
	tv.currentState.turnOn(tv)
}
func (tv *BigTV) close() {
	tv.currentState.turnOff(tv)
}
func (tv *BigTV) switchChannel() {
	tv.currentState.nextChannel(tv)
}
