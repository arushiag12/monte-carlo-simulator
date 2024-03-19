package scheduler

import (
	"proj3-redesigned/function"
	"proj3-redesigned/queue"
	"proj3-redesigned/worker"
	"math"
	"proj3-redesigned/barrier"
	"math/rand"
)

func RunParallel(config Config) {
	funcsInfo := function.ExtractFuncInfo(config.DataSize)

	for i := range *funcsInfo {
		(*funcsInfo)[i].Min = ProcessObjPar(&(*funcsInfo)[i], config.ThreadCount)
	}

	ResultsJSON(funcsInfo, config.DataSize)
}

func ProcessObjPar(funcObj *function.FuncInfo, numThreads int) float32 {
	b := barrier.NewBarrier(numThreads+1)
	results := make([]float32, numThreads)
	queues := make([]*queue.Queue, numThreads)

	for i := 0; i < numThreads; i++ {
		queues[i] = queue.NewQueue()
	}

	for x := funcObj.Dom.Min_x; x < funcObj.Dom.Max_x; x += 0.1 {
		for y := funcObj.Dom.Min_y; y < funcObj.Dom.Max_y; y += 0.1 {
			task := function.NewTask(funcObj.Func, x, x+0.1, y, y+0.1)
			r := rand.Intn(numThreads)
			queues[r].PushTop(task)
		}
	}

	for i := 0; i < numThreads; i++ {
		go worker.ParallelWorker(queues[i], &results[i], b)
	}
	b.Wait()

	res := float32(math.Inf(1))
	for _, r := range results {
		if r < res {
			res = r
		}
	}

	return res
}
