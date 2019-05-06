package comment

import (
	"fmt"
	"troytan/practice/并发爬虫/fetcher"
)

func Worker(request Request)(ParseResult, error){
	bytes, err := fetcher.Fetch(request.Url)
	if err != nil {
		fmt.Println("createWorker Fetch err, err:", err, "url:", request.Url)
		return ParseResult{}, err
	}
	//fmt.Println("worker",string(bytes))
	parseFunc := request.ParseFunc(bytes)
	return parseFunc, nil
}
