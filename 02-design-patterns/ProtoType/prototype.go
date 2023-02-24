package ProtoType

/*
1. 模式的结构
原型模式包含以下主要角色。
抽象原型类：规定了具体原型对象必须实现的接口。
具体原型类：实现抽象原型类的 clone() 方法，它是可被复制的对象。
访问类：使用具体原型类中的 clone() 方法来复制新的对象。
*/

type CloneAble interface {
	Clone() CloneAble
}

type Cat struct {
	name   string
	weight float32
}

func (c *Cat) Clone() Cat {
	return *c // 浅拷贝过去了
}
