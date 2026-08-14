package engine

import (
	"log"

	"github.com/yuhaowin/go-learning/crawler/internal/fetcher"
)

// Worker 被并发执行
func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
