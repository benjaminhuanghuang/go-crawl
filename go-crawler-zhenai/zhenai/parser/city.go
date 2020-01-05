package parser
import (
	"regexp"

	"../../engine"
)

// Page Url https://www.zhenai.com/zhenghun/shanghai

// User link <a href="http://album.zhenai.com/u/1378263033" target="_blank">细水长流</a>
var userRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
// Next Page <a href="http://www.zhenai.com/zhenghun/shanghai/2">下一页</a>
var nextUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
/*
ParseCity ...
*/
func ParseCity(contents []byte) engine.ParseResult{
	matches := userRe.FindAllSubmatch(contents, -1) // return [][][]bytes
	result := engine.ParseResult{}

	for _, m := range matches {
		name :=  string(m[2])
		url := string(m[1])
		// result.Items = append(result.Items, "User "+ name) // put user name to items
		result.Requests = append(result.Requests, engine.Request{
			URL:      url  ,
			// 通过函数式编程来添加参数
			ParserFunc: func (contents []byte) engine.ParseResult {
				return ParseProfile(contents,url, name)   // 
			},
		})
	}
	// request to next page
	matches = nextUrlRe.FindAllSubmatch(contents, -1)
	for _,m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			URL: string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}