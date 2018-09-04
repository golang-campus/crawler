package main

import (
	"./config"
	"crawler/engin"
	"crawler/zhenai/parser"
)

func main() {

	engin.Run(engin.Request{
		Url: config.BaseUrl,
		ParserFunc:parser.ParseCityList,
	})
	//
	//all, e := fetcher.Fetch(config.BaseUrl)
	//if e != nil {
	//	panic(e)
	//}
	//
	//printCityList(all)
}

//城市列表解析器
//func printCityList(contents []byte) {
//
//}
