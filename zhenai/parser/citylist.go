package parser

import (
	"crawler/engin"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engin.ParseResult {
	//<a href="http://www.zhenai.com/zhenghun/nanan1" class="">南岸</a>
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engin.ParseResult{}
	for _, m := range matches {
		//for _, subMatch := range m {
		//	fmt.Printf("%s ", subMatch)
		//}
		result.Items = append(result.Items,string(m[2]))
		result.Requests = append(result.Requests,engin.Request{
			Url:string(m[1]),
			ParserFunc:engin.NilParser,
		})
		//fmt.Printf("City: %s, URL: %s \n", m[2], m[1])
		//fmt.Println()
		//fmt.Printf("%s\n", m)
	}

	return result

}
