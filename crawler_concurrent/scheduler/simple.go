package scheduler

import "github.com/yuhaowin/go-learning/crawler_concurrent/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		// send request down to worker channel
		s.workerChan <- request
	}()
}

// ConfigureMasterWorkerChan 这里用指针接收者，因为要修改 s.workerChan；值接收者的话改的只是拷贝，
// 外部（包括 Submit）看到的仍然是 nil。
func (s *SimpleScheduler) ConfigureMasterWorkerChan(in chan engine.Request) {
	s.workerChan = in
}
