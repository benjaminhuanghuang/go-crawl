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
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Error("request should have %d  request; but had %d", resultSize, len(result.request))
	}
}
