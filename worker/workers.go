package worker

import (
	"proj3-redesigned/function"
	"math"
	"proj3-redesigned/queue"
	"proj3-redesigned/barrier"
	"sync/atomic"
	"math/rand"
)

func ProcessQueue(q *queue.Queue) float32 {
	min := float32(math.Inf(1))
	for {
		task := q.PopBottom()
		if task == nil {
			break
		}
		result := function.ProcessTask(task)
		if result < min {
			min = result
		}
	}
	return min
}

func ParallelWorker(q *queue.Queue, result *float32, b *barrier.Barrier) {
	*result = ProcessQueue(q)
	b.Wait()
}

func ParallelWSWorker(q []*queue.Queue, results *[]float32, i int, counter *int32, numThreads int, b *barrier.Barrier) {
	min := ProcessQueue(q[i])
	atomic.AddInt32(counter, 1)

	var new_val float32
	var r int

	for atomic.LoadInt32(counter) != int32(numThreads) {
		r = rand.Intn(numThreads)
		if r == i {
			continue
		}

		if q[i].IsEmpty() {
			continue
		}

		task := q[i].PopTop()
		if task == nil {
			continue
		}
	
		new_val = function.ProcessTask(task)
		if new_val < min {
			min = new_val
		}
	}

	(*results)[i] = min
	b.Wait()
}
