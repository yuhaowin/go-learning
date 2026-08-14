package parser

import (
	"regexp"

	"github.com/yuhaowin/go-learning/crawler/internal/model"
)

var (
	cityListRe = regexp.MustCompile(`<a [^>]*href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
)

func ParseCityList(contents []byte) model.ParseResult {

	result := model.ParseResult{}

	matches := cityListRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, model.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
