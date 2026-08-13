package engine

import (
	"log"

	"github.com/yuhaowin/go-learning/crawler/fetcher"
	"github.com/yuhaowin/go-learning/crawler/model"
)

// Worker 被并发执行
func Worker(r model.Request) (model.ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return model.ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
