package barrier

import (
	"sync"
	"sync/atomic"
)

// Barrier struct implementation with a condition variable
type Barrier struct {
	capacity int32
	count int32
	cond *sync.Cond
}

// Returns a new barrier with the given capacity
func NewBarrier(capacity int) *Barrier {
	return &Barrier{capacity: int32(capacity), count: 0, cond: sync.NewCond(&sync.Mutex{})}
}

// Blocks until all goroutines have reached the barrier
func (b *Barrier) Wait() {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	atomic.AddInt32(&b.count, 1)

	if b.count == b.capacity {
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}
}