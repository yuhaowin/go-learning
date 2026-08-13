package engine

import (
	"github.com/yuhaowin/go-learning/crawler/model"
)

// Scheduler 是接口类型，不是具体的 struct，所以"值类型/指针类型"的拷贝语义
// 对它不适用：接口值本身就是 (动态类型, 动态值) 的二元组，能装入任何实现了
// 该方法集的具体类型，无论那个具体类型是值类型还是指针类型。
// 例如 SimpleScheduler 的方法用指针接收者（见 scheduler/simple.go），
// 因此只有 *SimpleScheduler 满足这个接口，main.go 里赋给 Scheduler 字段时
// 传的就是 &scheduler.SimpleScheduler{} 指针。
type Scheduler interface {
	Submit(model.Request)
	ConfigureMasterWorkerChan(chan model.Request)
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan any
}

func (e *ConcurrentEngine) Run(seeds ...model.Request) {
	in := make(chan model.Request)
	out := make(chan model.ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out

		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan model.Request, out chan model.ParseResult) {
	go func() {
		for {
			request := <-in
			result, e := Worker(request)
			if e != nil {
				continue
			}
			out <- result
		}
	}()
}
