package main

import (
	"./engine"
	"./scheduler"
	"./zhenai/parser"
	"./persist"
)

func main() {
	e:=engine.ConcurrentEngine{
		// Scheduler : &scheduler.SimpleScheduler{},
		Scheduler : &scheduler.QueuedScheduler{},
		WorkerCount :100,
		ItemChan: persist.ItemServer(),
	}
	
	// Start from home page
	// e.Run(engine.Request{
	// 	URL:        "https://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })

	// Start from city page
	e.Run(engine.Request{
		URL:        "https://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
