package scheduler

import (
	"github.com/yuhaowin/go-learning/crawler/internal/engine"
)

// SimpleScheduler 所有 worker 共享同一个请求 channel，谁先读到就处理谁的，
// 不做排队缓冲，Submit 在请求暂时没人接收时也不会阻塞调用方。
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// Run 这里用指针接收者，因为要修改 s.workerChan；值接收者的话改的只是拷贝， 外部（包括 Submit）看到的仍然是 nil。
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		// send request down to worker channel
		s.workerChan <- request
	}()
}

// WorkerReady 对 SimpleScheduler 是空操作：所有 worker 从同一个 channel 里抢任务，
// 不需要单独登记"谁空闲了"。
func (s *SimpleScheduler) WorkerReady(chan engine.Request) {}
