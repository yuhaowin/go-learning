package main

import (
	"github.com/yuhaowin/go-learning/crawler/internal/engine"
	"github.com/yuhaowin/go-learning/crawler/internal/parser"
	"github.com/yuhaowin/go-learning/crawler/internal/saver"
	"github.com/yuhaowin/go-learning/crawler/internal/scheduler"
)

func main() {
	// SimpleScheduler 的方法用指针接收者（Run 要给 workerChan 赋值），
	// 所以只有 *SimpleScheduler 满足 engine.Scheduler 接口，这里必须传指针。
	// 详见 crawler/docs/interface-values.md。
	e := engine.ConcurrentEngine{
		// 必须传指针：SimpleScheduler 的方法用指针接收者（因为 ConfigureMasterWorkerChan 要修改 workerChan 字段），值类型不满足 Scheduler 接口。
		// SimpleScheduler 的两个方法 Submit 和 ConfigureMasterWorkerChan 都是定义在 *Sim pleScheduler（指针接收者）上的（scheduler/simple.go:9, :15），而不是值接收者。
		// 关键在于 ConfigureMasterWorkerChan 会给 s.workerChan 赋值（scheduler/simple.go:16），这是在修改 SimpleScheduler 结构体内部的字段。如果用值接收者，方法内修改的只是调用者传入值的一份拷贝，外部（包括后续 Submit 里用到的 s.workerChan）根本看不到这次赋值, workerChan 永远是 nil。
		// 另外一层原因是 Go 的方法集规则：只要有一个方法使用了指针接收者（这里两个都是），那么只有 *SimpleScheduler 才能满足 Scheduler 这个接口，SimpleScheduler（值类型）是不满足的。所以 main.go:11 必须写 &scheduler.SimpleScheduler{} 传指针，否则会编译报错「SimpleScheduler does not implement Scheduler」。
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
		ItemChan:    saver.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
