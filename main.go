package main

import (
	"./config"
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {

	resp, err := http.Get(config.BaseUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code is ", resp.StatusCode)
		return
	}

	//将网页文本的GBK转UTF8
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", all)

	printCityList(all)
}

//检测网页编码
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

//城市列表解析器
func printCityList(contents []byte) {
	//<a href="http://www.zhenai.com/zhenghun/nanan1" class="">南岸</a>
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//for _, subMatch := range m {
		//	fmt.Printf("%s ", subMatch)
		//}
		fmt.Printf("City: %s, URL: %s \n", m[2], m[1])
		fmt.Println()
		//fmt.Printf("%s\n", m)
	}
}
