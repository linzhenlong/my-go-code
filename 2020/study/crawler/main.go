package main

import (
	"github.com/linzhenlong/my-go-code/2020/study/crawler/engine"
	"github.com/linzhenlong/my-go-code/2020/study/crawler/zhenai/parser"
	
)

func main() {
	engine.Run(engine.Request{
		URL: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.CityList,
	})
}