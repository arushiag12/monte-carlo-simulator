package scheduler

import (
	"proj/barrier"
	"proj/function"
	"proj/queue"
	"proj/worker"
	"math"
	"math/rand"
)

// Processes the given functions in work-stealing mode
func RunWorkStealing(config Config) {
	funcsInfo := function.ExtractFuncInfo(config.DataSize)

	for i := range *funcsInfo {
		(*funcsInfo)[i].Min = ProcessObjPar(&(*funcsInfo)[i], config.ThreadCount)
	}

	ResultsJSON(funcsInfo, config.DataSize)
}

// Processes the given function object in work-stealing mode
func ProcessObjWS(funcObj *function.FuncInfo, numThreads int) float32 {
	b := barrier.NewBarrier(numThreads+1)
	counter := int32(0)
	results := make([]float32, numThreads)
	queues := make([]*queue.Queue, numThreads)

	// Create task queues for each thread
	for i := 0; i < numThreads; i++ {
		queues[i] = queue.NewQueue()
	}

	// Create tasks and push them to the queues
	for x := funcObj.Dom.Min_x; x < funcObj.Dom.Max_x; x += 0.1 {
		for y := funcObj.Dom.Min_y; y < funcObj.Dom.Max_y; y += 0.1 {
			task := function.NewTask(funcObj.Func, x, x+0.1, y, y+0.1)
			r := rand.Intn(numThreads)
			queues[r].PushTop(task)
		}
	}

	// Create and start worker threads
	for i := 0; i < numThreads; i++ {
		go worker.ParallelWSWorker(queues, &results, i, &counter, numThreads, b)
	}

	// Wait for all workers to finish
	b.Wait()

	// Find the minimum value
	res := float32(math.Inf(1))
	for _, r := range results {
		if r < res {
			res = r
		}
	}

	return res
}

