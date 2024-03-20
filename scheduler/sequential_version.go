package scheduler

import (
	"proj/function"
	"proj/queue"
	"proj/worker"
)

// Processes the given functions in sequential mode
func RunSequential(config Config) {
	funcsInfo := function.ExtractFuncInfo(config.DataSize)

	for i := range *funcsInfo {
		(*funcsInfo)[i].Min = ProcessObjSeq(&(*funcsInfo)[i])
	}

	ResultsJSON(funcsInfo, config.DataSize)
}

// Processes the given function object in sequential mode
func ProcessObjSeq(funcObj *function.FuncInfo) float32 {
	q := queue.NewQueue()

	// Create tasks and push them to the queue
	for x := funcObj.Dom.Min_x; x < funcObj.Dom.Max_x; x += 0.1 {
		for y := funcObj.Dom.Min_y; y < funcObj.Dom.Max_y; y += 0.1 {
			task := function.NewTask(funcObj.Func, x, x+0.1, y, y+0.1)
			q.PushTop(task)
		}
	}

	// Process the queue
	res := worker.ProcessQueue(q)
	return res
}
