package scheduler

import "github.com/yuhaowin/go-learning/crawler_queued/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request // 每个 worker 有自己的 channel of Request，所以 worker 通过 channel of Worker 共享
}

func (s QueuedScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

func (s QueuedScheduler) WorkerReady(worker chan engine.Request) {
	s.workerChan <- worker
}

func (s QueuedScheduler) ConfigureMasterWorkerChan(requests chan engine.Request) {
	//TODO implement me
	panic("implement me")
}
