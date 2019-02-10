package main

import (
	"GoStudy/Crawler/engine"
	"GoStudy/Crawler/zhenai/parser"
)

const URL  = "http://www.zhenai.com/zhenghun"
func main() {
	engine.Run(engine.Request{Url:URL,
		Parsefunc: parser.ParseCityList,
	})
}
