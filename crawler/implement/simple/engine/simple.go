package engine

import (
	"log"

	"github.com/yuhaowin/go-learning/crawler/fetcher"
	"github.com/yuhaowin/go-learning/crawler/model"
)

type SimpleEngine struct {
	ItemChan chan any
}

func (e SimpleEngine) Run(seeds ...model.Request) {
	var requests []model.Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s\n", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		result := r.ParserFunc(body)

		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		requests = append(requests, result.Requests...)
	}
}
