package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCityList(t *testing.T) {

	//contents, err := fetcher.Fetch(config.BaseUrl)
	//if err!= nil{
	//	panic(err)
	//}
	//
	//fmt.Printf("%s\n",contents)
	//
	//ParseCityList(contents)

	//Verify result

	//用本地存好的测试数据来测，更稳定，不受外界其他条件影响到测试结果
	contents, err := ioutil.ReadFile("cityList_test_data.html")
	if err!= nil{
		panic(err)
	}

	//fmt.Printf("%s\n",contents)

	result := ParseCityList(contents)

	const resultSize =470
	var expectedUrls = []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	//var expectedCities = []string{
	//
	//}
	if len(result.Requests) != resultSize{
		t.Errorf("result should have %d requests:",resultSize)
	}

	for i,url := range expectedUrls{
		if result.Requests[i].Url != url{
			t.Errorf("expected url #%d: %s; but "+ "was %s ",i,url,result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize{
		t.Errorf("result should have %d requests:",resultSize)
	}


}
