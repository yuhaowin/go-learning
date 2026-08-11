package parser

import (
	"regexp"
	"strconv"

	"github.com/yuhaowin/go-learning/crawler/engine"
	"github.com/yuhaowin/go-learning/crawler/model"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^>]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^>]+)</span></td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^>]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^>]+)</td>`)
var OccupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^>]+)</td>`)
var HouseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^>]+)</span></td>`)
var CarRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^>]+)</span></td>`)

// 提取id
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, e := strconv.Atoi(extractString(contents, ageRe))
	if e == nil {
		profile.Age = age
	}
	height, e := strconv.Atoi(extractString(contents, heightRe))
	if e == nil {
		profile.Height = height
	}
	weight, e := strconv.Atoi(extractString(contents, weightRe))
	if e == nil {
		profile.Weight = weight
	}
	profile.Income = extractString(contents, incomeRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Car = extractString(contents, CarRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, OccupationRe)
	profile.House = extractString(contents, HouseRe)
	profile.Marriage = extractString(contents, marriageRe)

	result := engine.ParseResult{
		Items: []any{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseProfile(contents, url, p.userName)

}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
