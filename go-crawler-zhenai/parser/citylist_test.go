package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityLity(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_page.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)
	const resultSize = 540
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng"}
	if len(result.Requests) != resultSize {
		t.Errorf("request should have %d  request; but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].URL != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].URL)
		}
	}
	expectedCitys := []string{"阿坝", "阿克苏", "阿拉善盟"}
	if len(result.Items) != resultSize {
		t.Errorf("request should have %d  request; but had %d", resultSize, len(result.Items))
	}

	for i, city := range expectedCitys {
		if result.Items[i].(string) != city {
			t.Errorf("expected url #%d: %s; but was %s", i, city, result.Items[i])
		}
	}
}
