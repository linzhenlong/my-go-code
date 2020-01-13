package taskrunner

import (
	"time"
)





// Worker 结构体.
type Worker struct {
	ticker *time.Ticker
	runner *Runner
	
}

// NewWorker XXX.
func NewWorker(interval time.Duration, r *Runner)*Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner:r,
	}
}
func (w *Worker)startWorker() {
	for {
		select {
		case <-w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

func Start() {
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(3,r)
	go w.startWorker()

	// 该5-7
}