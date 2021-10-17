package taskrunner

import "time"

type Worker struct {
	Ticker time.Ticker // 定时器
	Runner *Runner
}

// new
func NewWorker(interval time.Duration, runner *Runner) *Worker {
	return &Worker{
		Ticker: *time.NewTicker(interval * time.Second),
		Runner: runner,
	}
}

// start
func (worker *Worker) startWorker() {
	for {
		select {
		case _ = <-worker.Ticker.C:
			worker.Runner.StartAll()
		}
	}
}

func main() {
	r := NewRunner(true, VideoDeleteDispatch, VideoDeleteExcute, 3)
	w := NewWorker(DELETE_INTERVAL, r)
	go w.startWorker()
}
