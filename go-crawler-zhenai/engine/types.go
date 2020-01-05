package engine

type ParserFunc func(contents []byte, url string) ParseResult

/*
Request ...
*/
type Request struct {
	URL        string
	ParserFunc ParserFunc
}

/*
ParseResult ...
*/
type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}

/*
NilParser ...
*/
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
