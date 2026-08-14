package scheduler

import (
	"github.com/yuhaowin/go-learning/crawler/internal/model"
)

// SimpleScheduler 所有 worker 共享同一个请求 channel，谁先读到就处理谁的，
// 不做排队缓冲，Submit 在请求暂时没人接收时也不会阻塞调用方。
type SimpleScheduler struct {
	workerChan chan model.Request
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan model.Request)
}

func (s *SimpleScheduler) WorkerChan() chan model.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Submit(request model.Request) {
	go func() {
		s.workerChan <- request
	}()
}

// WorkerReady 对 SimpleScheduler 是空操作：所有 worker 从同一个 channel 里抢任务，
// 不需要单独登记"谁空闲了"。
func (s *SimpleScheduler) WorkerReady(chan model.Request) {}
