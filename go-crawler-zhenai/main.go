package main

import (
	"./engine"

	"./parser"
)

func main() {
	engine.Run(engine.Request{
		URL:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
