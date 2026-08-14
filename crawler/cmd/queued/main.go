package main

import (
	"github.com/yuhaowin/go-learning/crawler/cmd/queued/engine"
	"github.com/yuhaowin/go-learning/crawler/cmd/queued/scheduler"
	"github.com/yuhaowin/go-learning/crawler/internal/model"
	"github.com/yuhaowin/go-learning/crawler/internal/parser"
	"github.com/yuhaowin/go-learning/crawler/internal/saver"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    saver.ItemSaver(),
	}

	e.Run(model.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
