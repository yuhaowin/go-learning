package engine

import (
	"github.com/yuhaowin/go-learning/crawler/internal/model"
)

// Scheduler 由 internal/scheduler 下的具体实现（SimpleScheduler、QueuedScheduler）满足，
// ConcurrentEngine 通过替换 Scheduler 实现来切换调度策略。

// Scheduler 是接口类型，不是具体的 struct，所以"值类型/指针类型"的拷贝语义对它不适用：
// 接口值本身就是 (动态类型, 动态值) 的二元组，能装入任何实现了该方法集的具体类型，无论那个具体类型是值类型还是指针类型。
// 例如 SimpleScheduler 的方法用指针接收者（见 scheduler/simple.go）因此只有 *SimpleScheduler 满足这个接口。
// main.go 里赋给 Scheduler 字段时传的就是 &scheduler.SimpleScheduler{} 指针。
type Scheduler interface {
	Notifier
	Submit(model.Request)
	WorkerChan() chan model.Request
	Run()
}

type Notifier interface {
	WorkerReady(chan model.Request)
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan any
}

func (e *ConcurrentEngine) Run(seeds ...model.Request) {
	e.Scheduler.Run()
	out := make(chan model.ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
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

func createWorker(in chan model.Request, out chan model.ParseResult, notifier Notifier) {
	go func() {
		for {
			notifier.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
