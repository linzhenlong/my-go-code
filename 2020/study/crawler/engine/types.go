package engine

// Request 请求结构体
type Request struct {
	URL        string
	ParserFunc func([]byte) ParserResult
}

// ParserResult 解析结果结构体.
type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

// NilParser nil解析器.
func NilParser(contents []byte) ParserResult {
	return ParserResult{}
}
