package parser

import(
	"testing"
	"fmt"
	"regexp"
	"io/ioutil"
	_"fmt"
	_"github.com/linzhenlong/my-go-code/2020/study/crawler/fetcher"
	"github.com/linzhenlong/my-go-code/2020/study/crawler/fetcher"

)

func TestCityList(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("./citylist_test_data.html")
	if err != nil {
		t.Fatalf("fetch error:%s",err.Error())
	}
	//t.Logf("%s\n", contents)
	result := CityList(contents)
	
	const resultSize = 470

	expectedUrls := []string {
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedCities := []string {
		"阿坝",
		"阿克苏", 
		"阿拉善盟",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("error result should have requests %d, but got %d", resultSize, len(result.Requests))
		//t.Fatalf("Fatalf result should have %d, but got %d",resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].URL != url {
			t.Errorf("expected url #%d,%s but was %s",i, url, result.Requests[i].URL)
		}
	}
	if len(result.Items) != resultSize {
		t.Errorf("error result should have items %d, but got %d", resultSize, len(result.Requests))
		//t.Fatalf("Fatalf result should have %d, but got %d",resultSize, len(result.Requests))
	}
	for i, city := range expectedCities {
		if result.Items[i] != city {
			t.Errorf("expected city #%d,%s but was %s",i, city, result.Items[i])
		}
	}
	t.Logf("%d", len(result.Requests))
}


func TestProfile(t *testing.T) {
	const ageRegexp = `<div [^>]* class="des f-cl">[.+] | ([\d]+)岁 | [.+] | [.+] | ([\d]+)cm | ([\d]+-[\d]+)元</div>`
	
	contents, err := fetcher.Fetch("http://album.zhenai.com/u/1727435860")
	if err != nil {
		t.Errorf("error:%s", err.Error())
	}
	match := regexp.MustCompile(ageRegexp).FindAllStringSubmatch(string(contents), -1)
	fmt.Println(match)
}