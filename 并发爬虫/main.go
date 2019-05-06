package main

import (
	"troytan/practice/并发爬虫/comment"
	"troytan/practice/并发爬虫/engine"
	"troytan/practice/并发爬虫/parse/zhenai"
	"troytan/practice/并发爬虫/scheduler"
)

func main() {
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 10,
	}
	concurrentEngine.Run(comment.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc:parse.CityParseList,
	})
}