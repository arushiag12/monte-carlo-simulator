package worker

import (
	"math"
	"math/rand"
	"proj/barrier"
	"proj/function"
	"proj/queue"
	"sync/atomic"
)

// Processes all tasks in the given queue and returns the minimum value
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

// Processes a queue in parallel and waits at a barrier
func ParallelWorker(q *queue.Queue, result *float32, b *barrier.Barrier) {
	// Process the queue
	*result = ProcessQueue(q)

	// Wait at the barrier
	b.Wait()
}

// ParallelWSWorker processes a queue in parallel with work stealing
func ParallelWSWorker(q []*queue.Queue, results *[]float32, i int, counter *int32, numThreads int, b *barrier.Barrier) {
	// Process own queue
	min := ProcessQueue(q[i])

	// Increment counter
	atomic.AddInt32(counter, 1)

	var new_val float32
	var r int

    // Attempt to steal work until all threads are done (by checking the counter)
	for atomic.LoadInt32(counter) != int32(numThreads) {

		// Randomly select a queue to steal from
		r = rand.Intn(numThreads)

		// Skip if trying to steal from self
		if r == i {
			continue
		}

		// Skip if chosen queue is empty
		if q[i].IsEmpty() {
			continue
		}

		// Steal work
		task := q[i].PopTop()

		// Skip if no task was stolen
		if task == nil {
			continue
		}

		// Process the stolen task
		new_val = function.ProcessTask(task)
		if new_val < min {
			min = new_val
		}
	}
	(*results)[i] = min

	// Wait at the barrier
	b.Wait()
}
