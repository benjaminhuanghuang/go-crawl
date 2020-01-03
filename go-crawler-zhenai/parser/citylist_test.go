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

	if len(result.Requests) != resultSize {
		t.Errorf("request should have %d  request; but had %d", resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("request should have %d  request; but had %d", resultSize, len(result.Items))
	}
}
