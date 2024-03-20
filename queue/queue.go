package queue 

import (
	"proj/function"
	"sync/atomic"
)

// Node struct for queue
type Node struct {
	task *function.Task
	next *Node
	prev *Node
}

// Queue struct with atomic pointers for thread safety
type Queue struct {
	bottom atomic.Pointer[Node]
	top atomic.Pointer[Node]
}

// Creates a new queue
func NewQueue() *Queue {
	q := &Queue{}
	q.bottom.Store(nil)
	q.top.Store(nil)
	return q
}

// Adds a task to the top of the queue
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

// Removes a task from the bottom of the queue
func (q *Queue) PopBottom() *function.Task {
	if q.bottom.Load() == nil {
		return nil
	}
	task := q.bottom.Load().task
	q.bottom.Store(q.bottom.Load().next)
	return task
}

// Returns boolean indicating if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.bottom.Load() == nil
}

// Tries to remove a task from the top of the queue in a thread-safe manner and returns it if successful, otherwise returns nil
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
