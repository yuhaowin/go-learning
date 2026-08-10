package engine

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

// ParserFunc 定义一个函数类型，
//
//	函数类型声明（type X func(...)）：参数名可选，纯粹给阅读者看，编译器不关心，甚至同名不同名的两个类型只要参数类型和返回值类型一致就是同一类型。
type ParserFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url        string
	ParserFunc ParserFunc
}

type ParseResult struct {
	Requests []Request
	Items    []any
}

type Item struct {
	Id      string //存储时去重。
	Url     string
	Type    string //存储的配置
	Payload interface{}
}

func NilFuncParser(content []byte, name string) ParseResult {
	return ParseResult{}
}

type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

// FuncParser 函数类型的Parser
type FuncParser struct {
	parser ParserFunc //对应解析函数
	name   string     //函数名
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}

//func CreateFuncParser(
//	p ParserFunc, name string) *FuncParser {
//	return &FuncParser{
//		parser: p,
//		name:   name,
//	}
//}

//func NilParser([]byte) ParseResult  {
//	return ParseResult{}
//}

//type SerializedParser struct {
//	Name string
//	Args interface{}
//}

//{"ParseCitylist", nil}
//{"ProfileParser", userName}
