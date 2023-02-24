package Observer

import "testing"

func TestObserver(t *testing.T) {
	subject := Subject{
		topic: "乌克兰",
		users: nil,
	}
	subject.AddObserver(&PersonObserver{
		name: "张三",
		age:  30,
	})
	subject.AddObserver(&PersonObserver{
		name: "小二",
		age:  12,
	})
	subject.Publish("撤军")
}
