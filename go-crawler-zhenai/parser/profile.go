package parser

import (
	"regexp"
	"strconv"

	"../engine"
	"../model"
)

//<div data-v-8b1eac0c="" class="m-btn purple">33岁</div>
var ageRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)岁</div>`)

//<div data-v-8b1eac0c="" class="m-btn purple">离异</div>
var marriageRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)岁</div>`)

/*
ParseProfile ...
https://album.zhenai.com/u/1280064210
*/
func ParseProfile(contents []byte, userName string) engine.ParseResult {

	profile := model.Profile{}
	profile.Name = userName

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}

	profile.Marriage = extractString(contents, marriageRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
