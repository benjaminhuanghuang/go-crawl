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
	Items    []interface{}
}

/*
NilParser ...
*/
func NilParser([]byte) ParseResult {
	return ParseResult{}
}


