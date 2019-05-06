package parse

import (
	"fmt"
	"regexp"
	"troytan/practice/并发爬虫/comment"
)

const (
	cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
	personRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`
	personInfoRe = `<div class="des f-cl" data-v-3c42fade>([^<]+)</div>`
)
func CityParseList(contents []byte)comment.ParseResult{
	re := regexp.MustCompile(cityListRe)
	submatch := re.FindAllSubmatch(contents, -1)
	fmt.Println("submatch",submatch)
	result := comment.ParseResult{}
	for _,value := range submatch{
		result.Iterm = append(result.Iterm, string(value[2]))
		result.Request = append(result.Request, comment.Request{
			Url :	string(value[1]),
			ParseFunc:	PersonParseList,
		} )
	}

	return result
}
func PersonParseList(contents []byte)comment.ParseResult{
	re := regexp.MustCompile(personRe)
	submatch := re.FindAllSubmatch(contents, -1)
	result := comment.ParseResult{}
	for _,value := range submatch{
		result.Iterm = append(result.Iterm, string(value[2]))
		result.Request = append(result.Request, comment.Request{
			Url :	string(value[1]),
			ParseFunc:	PersonInfoParseList,
		} )
	}

	return result
}
func PersonInfoParseList(contents []byte)comment.ParseResult {
	re := regexp.MustCompile(personInfoRe)
	submatch := re.FindAllSubmatch(contents, -1)
	result := comment.ParseResult{}
	for _,value := range submatch{
		result.Iterm = append(result.Iterm, string(value[1]))
		result.Request = append(result.Request, comment.Request{} )
	}
	return result
}