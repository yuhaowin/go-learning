package main

import (
	"github.com/yuhaowin/go-learning/crawler/implement/queued/engine"
	"github.com/yuhaowin/go-learning/crawler/implement/queued/scheduler"
	"github.com/yuhaowin/go-learning/crawler/model"
	"github.com/yuhaowin/go-learning/crawler/parser"
	"github.com/yuhaowin/go-learning/crawler/saver"
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
