package queue 

import (
	"proj3-redesigned/function"
	"sync/atomic"
)

type Node struct {
	task *function.Task
	next *Node
	prev *Node
}

type Queue struct {
	bottom atomic.Pointer[Node]
	top atomic.Pointer[Node]
}

func NewQueue() *Queue {
	q := &Queue{}
	q.bottom.Store(nil)
	q.top.Store(nil)
	return q
}

func (q *Queue) PushTop(task *function.Task) {
	node := &Node{task: task, next: nil, prev: nil}
	if q.bottom.Load() == nil {
		q.bottom.Store(node)
		q.top.Store(node)
	} else {
		node.prev = q.top.Load()
		q.top.Load().next = node
		q.top.Store(node)
	}
}

func (q *Queue) PopBottom() *function.Task {
	if q.bottom.Load() == nil {
		return nil
	}
	task := q.bottom.Load().task
	q.bottom.Store(q.bottom.Load().next)
	return task
}

func (q *Queue) IsEmpty() bool {
	return q.bottom.Load() == nil
}

func (q *Queue) PopTop() *function.Task {
	if q.IsEmpty() {
		return nil
	}
	oldTop := q.top.Load()
	if (q.top).CompareAndSwap(oldTop, oldTop.prev) {
		return oldTop.task
	}
	return nil
}
