package parse

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"../model"
)

var countRegex = regexp.MustCompile(`<dd class="clearfix">([\s\S]+?)</dd>`)
var imagRegex = regexp.MustCompile(`<div class="item-pic" style="background-image: url(////(.*))">`)
var titleRegex = regexp.MustCompile(`<p class="item-name">(.*)</p>`)
var introduceRegex = regexp.MustCompile(`<p class="item-name">(.*)</p>`)
var degreeRegex = regexp.MustCompile(`<span class="item-level">(.*)</span>`)
var learnCountRegex = regexp.MustCompile(`	<span class="item-number"><i class="imwap-peaple"></i>(.*)</span>`)
var priceRegex = regexp.MustCompile(`<span class="pay-price red">¥(.*)</span>`)

func ParseCourse(data []byte) {
	// fmt.Printf("%s", data)
	imooc := model.Imooc{}
	list := make([]interface{}, 0)
	result := findCount(data, countRegex)

	for _, value := range result {
		imooc.Img = extractString(value[1], imagRegex)
		imooc.Title = extractString(value[1], titleRegex)
		imooc.Introduce = extractString(value[1], introduceRegex)
		imooc.Degree = extractString(value[1], degreeRegex)
		imooc.LearnCount = extractString(value[1], learnCountRegex)
		imooc.Price = extractString(value[1], priceRegex)

		list = append(list, imooc)
	}
	b, err := json.MarshalIndent(list, "", " ")
	fmt.Println(b)
	f, err := os.OpenFile("learn_go.json", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return
	}

	defer f.Close()
	if _, err := f.Write(b); err != nil {
		panic(err)
	}
}

func extractString(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)

	if len(match) >= 2 {
		return string(match[1])
	}

	return ""
}

func findCount(content []byte, re *regexp.Regexp) [][][]byte {
	match := re.FindAllSubmatch(content, -1) // -1 means return all match
	return match
}
