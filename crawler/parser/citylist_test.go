package parser

import (
	"testing"

	"github.com/yuhaowin/go-learning/crawler/fetcher"
)

func TestParseCityList(t *testing.T) {
	contents, e := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	if e != nil {
		panic(e)
	}

	result := ParseCityList(contents)
	const resultSize = 494

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
}
