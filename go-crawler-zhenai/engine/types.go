package engine

/*
Request ...
*/
type Request struct {
	URL       string
	ParserFunc func([]byte) ParseResult
}

/*
ParseResult ...
*/
type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct{
	Url string
	Id string
	Type string
	Payload interface{}
}
/*
NilParser ...
*/
func NilParser([]byte) ParseResult {
	return ParseResult{}
}


