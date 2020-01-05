package main

import (
	"./engine"
	"./scheduler"
	"./zhenai/parser"
	"./persist"
)

func main() {
	itemChan, err = persist.ItemServer("dating_profile")  // pass table name
	if err != nil{
		panic(err)
	}

	e:=engine.ConcurrentEngine{
		// Scheduler : &scheduler.SimpleScheduler{},
		Scheduler : &scheduler.QueuedScheduler{},
		WorkerCount :100,
		ItemChan: itemChan,
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
