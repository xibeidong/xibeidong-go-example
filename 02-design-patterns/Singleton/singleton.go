package Singleton

import "sync"

/*
单例（Singleton）模式的定义：
指一个类只有一个实例，且该类能自行创建这个实例的一种模式。
例如，Windows 中只能打开一个任务管理器，
这样可以避免因打开多个任务管理器窗口而造成内存资源的浪费，
或出现各个窗口显示内容的不一致等错误。

单例模式的主要角色如下。
单例类：包含一个实例且能自行创建这个实例的类。
访问类：使用单例的类。
*/
var (
	p    *Pet
	once sync.Once
)

type Pet struct {
	name string
	age  int
	mux  sync.Mutex
}

func init() {
	once.Do(func() {
		p = &Pet{
			name: "pet",
			age:  2,
			mux:  sync.Mutex{},
		}
	})
}

func GetInstance() *Pet {
	return p
}

func (p *Pet) SetAge(a int) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.age = a
}
