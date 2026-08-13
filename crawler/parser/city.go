package parser

import (
	"regexp"

	"github.com/yuhaowin/go-learning/crawler/model"
)

var (
	homepage  = "http://www.zhenai.com/zhenghun"
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) model.ParseResult {

	result := model.ParseResult{}

	matches := profileRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		name := string(m[2])
		result.Requests = append(result.Requests, model.Request{
			Url: homepage, // test
			ParserFunc: func(contents []byte) model.ParseResult {
				return ParseProfile(contents, name)
			},
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, model.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
