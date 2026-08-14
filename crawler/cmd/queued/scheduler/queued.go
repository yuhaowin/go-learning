package scheduler

import (
	"github.com/yuhaowin/go-learning/crawler/internal/model"
)

type QueuedScheduler struct {
	requestChan chan model.Request
	workerChan  chan chan model.Request // 每个 worker 有自己的 channel of Request，所以 worker 通过 channel of Worker 共享
}

func (s *QueuedScheduler) Submit(request model.Request) {
	s.requestChan <- request
}

func (s *QueuedScheduler) WorkerChan() chan model.Request {
	return make(chan model.Request)
}

func (s *QueuedScheduler) WorkerReady(worker chan model.Request) {
	s.workerChan <- worker
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan model.Request)
	s.workerChan = make(chan chan model.Request)
	go func() {
		var requestQ []model.Request
		var workerQ []chan model.Request

		for {
			var activeRequest model.Request
			var activeWorker chan model.Request
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
