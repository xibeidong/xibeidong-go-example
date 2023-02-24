package Visitor

import "fmt"

/*
访问者（Visitor）模式实现的关键是如何将作用于元素的操作分离出来封装成独立的类，其基本结构与实现方法如下。
1. 模式的结构
访问者模式包含以下主要角色。
抽象访问者（Visitor）角色：定义一个访问具体元素的接口，为每个具体元素类对应一个访问操作 visit() ，该操作中的参数类型标识了被访问的具体元素。
具体访问者（ConcreteVisitor）角色：实现抽象访问者角色中声明的各个访问操作，确定访问者访问一个元素时该做什么。
抽象元素（Person）角色：声明一个包含接受操作 accept() 的接口，被接受的访问者对象作为 accept() 方法的参数。
具体元素（ConcreteElement）角色：实现抽象元素角色提供的 accept() 操作，其方法体通常都是 visitor.visit(this) ，另外具体元素中可能还包含本身业务逻辑的相关操作。
对象结构（Object Structure）角色：是一个包含元素角色的容器，提供让访问者对象遍历容器中的所有元素的方法，通常由 List、Set、Map 等聚合类实现。
*/

//实例：一篇博客文章，博主关注的是阅读量和评论，读者关注的是标题和内容

type IVisitor interface {
	Visit(p IBlog)
}

// FanVisitor 读者
type FanVisitor struct {
}

func (w *FanVisitor) Visit(p IBlog) {
	fmt.Println("fan notice title is ", p.(*Article).title)
}

// BloggerVisitor 博主
type BloggerVisitor struct {
}

func (w *BloggerVisitor) Visit(p IBlog) {
	fmt.Println("blogger notice numberOfReaders  ", p.(*Article).numberOfReaders)
}

type IBlog interface {
	accept(v IVisitor)
}

type Article struct {
	title           string //标题
	numberOfReaders int    //阅读量
}

func (art *Article) accept(v IVisitor) {
	fmt.Println("Article accept visit.")
	v.Visit(art)
}
