package engine

type Request struct {
	Url string
	Parsefunc func([]byte) ParseResult
}

type ParseResult struct{
	Requests []Request
	Items []interface{}
}

func NilParser([]byte) ParseResult{
	return  ParseResult{}
}