package scheduler

import (
	"github.com/yuhaowin/go-learning/crawler/internal/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request // 每个 worker 有自己的 channel of Request，所以 worker 通过 channel of Worker 共享
}

func (s *QueuedScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) WorkerReady(worker chan engine.Request) {
	s.workerChan <- worker
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
