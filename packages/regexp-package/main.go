package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	searchSentence := "Ali used up 23.50 ringgit, he has 10.30 ringgit left"
	pattern := "[0-9]+.[0-9]+"

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		num := strconv.FormatFloat(v*2, 'f', 2, 32)
		return num
	}

	if ok, _ := regexp.Match(pattern, []byte(searchSentence)); ok {
		fmt.Println("Found your String")
	}
	re, _ := regexp.Compile(pattern)

	str := re.Find([]byte(searchSentence))
	fmt.Println(string(str))

	str2 := re.ReplaceAllStringFunc(searchSentence, f)
	fmt.Println(str2)
}
