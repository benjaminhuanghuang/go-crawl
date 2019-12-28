package main

import (
	"./fetch"
	"./parse"
)

const url = "https://coding.imooc.com"

func main() {
	fetch.Fetch(url, parse.ParseCourse)
}
