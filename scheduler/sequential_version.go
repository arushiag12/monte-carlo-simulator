package scheduler

import (
	"proj3-redesigned/function"
	"proj3-redesigned/queue"
	"proj3-redesigned/worker"
)

func RunSequential(config Config) {
	funcsInfo := function.ExtractFuncInfo(config.DataSize)

	for i := range *funcsInfo {
		(*funcsInfo)[i].Min = ProcessObjSeq(&(*funcsInfo)[i])
	}

	ResultsJSON(funcsInfo, config.DataSize)
}

func ProcessObjSeq(funcObj *function.FuncInfo) float32 {
	q := queue.NewQueue()

	for x := funcObj.Dom.Min_x; x < funcObj.Dom.Max_x; x += 0.1 {
		for y := funcObj.Dom.Min_y; y < funcObj.Dom.Max_y; y += 0.1 {
			task := function.NewTask(funcObj.Func, x, x+0.1, y, y+0.1)
			q.PushTop(task)
		}
	}

	res := worker.ProcessQueue(q)
	return res
}
