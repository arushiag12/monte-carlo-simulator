package function

import(
	"math/rand"
	"math"
	"time"
)

type Task struct {
	Func func(float32, float32) float32
	Min_x float32
	Max_x float32
	Min_y float32
	Max_y float32
}

func NewTask(f func(float32, float32) float32, min_x float32, max_x float32, min_y float32, max_y float32) *Task {
	return &Task{Func: f, Min_x: min_x, Max_x: max_x, Min_y: min_y, Max_y: max_y}
}

func ProcessTask(task *Task) float32 {
    min := float32(math.Inf(1))
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
    for i := 0; i < 100; i++ {
        x := r.Float32() * (task.Max_x - task.Min_x) + task.Min_x
        y := r.Float32() * (task.Max_y - task.Min_y) + task.Min_y
        result := task.Func(x, y)
        if result < min {
            min = result
        }
    }
    return min
}
