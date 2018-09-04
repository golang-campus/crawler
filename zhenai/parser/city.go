package parser

import (
	"crawler/engin"
	"regexp"
)

//将city详情页的会员列表里面各个会员的个人主页的页面进行访问
//const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">`
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engin.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engin.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engin.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engin.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}
	return result
}
