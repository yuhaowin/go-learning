package main

import (
	"github.com/yuhaowin/go-learning/crawler/cmd/simple/engine"
	"github.com/yuhaowin/go-learning/crawler/internal/model"
	"github.com/yuhaowin/go-learning/crawler/internal/parser"
	"github.com/yuhaowin/go-learning/crawler/internal/saver"
)

// 单任务版爬虫
// 获取并打印所有城市第一页用户的详细信息
func main() {
	engine.SimpleEngine{ItemChan: saver.ItemSaver()}.Run(model.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
