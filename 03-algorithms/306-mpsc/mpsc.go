package mpsc

import (
	"sync/atomic"
	"unsafe"
)

/*
 *	oneshot: 代表一个发送者，一个接收者的队列
 * 	mpsc: 代表多个发送者，一个接收者的队列
 *	spmc/broadcast: 代表一个发送者，多个接收者的队列
 * 	mpmc/channel: 代表多个发送者，多个接收者的队列,go的chan就是
 *
 *	根据场景的不同，选择不同的队列，可以得到更好的性能
 */

type node struct {
	val  interface{}
	next *node
}

// 数据结构如下，单向链表，方向从tail指向head：
// tail -> tail.next -> tail.next.next -> head（最新的数据）
// push添加到head，pop从tail后面移走，tail节点不含有效数据，只有一个next指针
// lock-free 的 queue 没有mutx竞争，写效率更高，读只能单线程

type Queue struct {
	head *node
	tail *node
}

func New() *Queue {
	q := new(Queue)
	q.tail = &node{}
	q.head = q.tail
	return q
}

// Push 添加到队列的后面（head），并发安全。
func (q *Queue) Push(v interface{}) {
	n := &node{}
	n.val = v

	// n成为新的head
	prev := (*node)(atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.head)), unsafe.Pointer(n)))

	// 原来的head指向新的head，串联起来。
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&prev.next)), unsafe.Pointer(n))

}

// Pop must be called from a single, consumer goroutine
func (q *Queue) Pop() interface{} {
	if q.tail.next == nil {
		return nil
	}
	next := (*node)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail.next))))
	v := next.val
	// tail指向下一个节点
	q.tail.next = next.next
	// pop出去的节点设为nil，GC回收
	next = nil
	return v
}

func (q *Queue) Empty() bool {
	tail := q.tail
	next := (*node)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&tail.next))))
	return next == nil
}

// 参考gonet内部实现 https://github.com/bobohume/gonet/blob/master/base/mpsc/deque.go
// 优化思考：sync.Pool管理node，； 链表读写性能差，考虑其它数据结构
