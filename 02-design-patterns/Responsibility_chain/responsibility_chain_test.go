package Responsibility_chain

import "testing"

func TestResponsibilityChain(t *testing.T) {
	teacher := &Teacher{name: "李老师"}
	leader := &Leader{name: "王主任"}
	dean := &Dean{name: "张院长"}

	teacher.next = leader
	leader.next = dean

	teacher.HandleMsg(12)

}