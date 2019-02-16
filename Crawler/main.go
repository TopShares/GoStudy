package main

import (
	"GoStudy/Crawler/engine"
	"GoStudy/Crawler/scheduler"
	"GoStudy/Crawler/zhenai/parser"
)

const URL  = "http://www.zhenai.com/zhenghun"
func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	e :=engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount:100,
		//ItemChan: persist.ItemSaver(),
	}
	//e.Run(engine.Request{
	//	Url:URL,
	//	Parsefunc: parser.ParseCityList,
	//})
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun/shanghai",
		Parsefunc: parser.ParseCity,
	})
}
