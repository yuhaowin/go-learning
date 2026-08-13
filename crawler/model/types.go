package model

// ParserFunc 定义一个函数类型，
//
//	函数类型声明（type X func(...)）：参数名可选，纯粹给阅读者看，编译器不关心，甚至同名不同名的两个类型只要参数类型和返回值类型一致就是同一类型。
type ParserFunc func(contents []byte) ParseResult

type Request struct {
	Url        string
	ParserFunc ParserFunc
}

type ParseResult struct {
	Requests []Request
	Items    []any
}

func NilFuncParser(content []byte) ParseResult {
	return ParseResult{}
}
