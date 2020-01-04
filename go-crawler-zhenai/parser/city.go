package parser
import (
	"regexp"

	"../engine"
)
// <a href="http://album.zhenai.com/u/1378263033" target="_blank">细水长流</a>
const userRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`


/*
ParseCity ...
*/
func ParseCity(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(userRe)
	matches := re.FindAllSubmatch(contents, -1) // return [][][]bytes
	result := engine.ParseResult{}

	for _, m := range matches {
		name :=  string(m[2])
		result.Items = append(result.Items, "User "+ name) // put user name to items
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(m[1]),
			// 通过函数式编程来添加参数
			ParserFunc: func (contents []byte) engine.ParseResult {
				return ParseProfile(contents, name)   // 
			},
		})
	}
	return result
}
