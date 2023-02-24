package Iterator

/*
迭代器模式有以下角色
抽象迭代器(Iterator)角色： 一般定义为接口，用来定义访问和遍历元素的接口。
具体迭代器(ConcreteIterator)角色： 实现对聚合对象的遍历，并跟踪遍历时的当前位置。
抽象聚合(Aggregate)角色： 定义创建相应迭代器对象的接口。
具体聚合(ConcreteAggregate)角色： 实现创建相应的迭代器对象。

https://blog.csdn.net/weixin_42146366/article/details/107564429
*/

type Aggregate interface {
	GetIterator() Iterator
}
type ConcreteAggregate struct {
	array []interface{}
}

func (c *ConcreteAggregate) GetIterator() Iterator {
	return &ConcreteIterator{
		array: c.array,
		index: 0,
	}
}

type Iterator interface {
	HasNext() bool
	Next()
	Value() interface{}
	//Add(v interface{})
	//Remove(v interface{})
}
type ConcreteIterator struct {
	array []interface{}
	index int
}

func (c *ConcreteIterator) Value() interface{} {
	//TODO implement me
	if c.index < len(c.array) {
		return c.array[c.index]
	}
	return nil
}

func (c *ConcreteIterator) HasNext() bool {
	return c.index < len(c.array)
}
func (c *ConcreteIterator) Next() {
	c.index++
}
