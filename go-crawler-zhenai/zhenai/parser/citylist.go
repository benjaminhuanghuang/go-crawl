package parser

import (
	"regexp"

	"../../engine"
)

const cityUrlRe = `href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

/*
ParseCityList ...
*/
func ParseCityList(contents []byte) engine.ParseResult {
	//<a data-v-5e16505f="" href="http://www.zhenai.com/zhenghun/haerbin">哈尔滨</a>
	re := regexp.MustCompile(cityUrlRe)
	matches := re.FindAllSubmatch(contents, -1) // return [][][]bytes
	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2])) // put city name to items
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
