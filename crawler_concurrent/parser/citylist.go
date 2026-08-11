package parser

import (
	"regexp"

	"github.com/yuhaowin/go-learning/crawler_concurrent/engine"
)

var (
	cityListRe = regexp.MustCompile(`<a [^>]*href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
)

func ParseCityList(contents []byte) engine.ParseResult {

	result := engine.ParseResult{}

	matches := cityListRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
