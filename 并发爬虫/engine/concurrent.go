package engine

import (
	"fmt"
	"troytan/practice/并发爬虫/comment"
)

//定义一个调度接口（simple和queue分别进行实现）
type Scheduler interface {
	Run()
	Submit(comment.Request)
	WorkerChan()chan comment.Request
	WorkerReady(chan comment.Request)
}
//调度驱动，指定实现scheduler的是simple还是queue；初始化并发的数量
type ConcurrentEngine struct {
	Scheduler 	Scheduler
	WorkerCount	int
}
func (this *ConcurrentEngine) Run(seeds ...comment.Request){
	//1.创建Worker
	out := make(chan comment.ParseResult)
	this.Scheduler.Run()
	for i := 0; i < this.WorkerCount; i++ {
		in := this.Scheduler.WorkerChan()
		createWorker(in, out, this.Scheduler)
	}
	for _, request := range seeds {
		this.Scheduler.Submit(request)
	}
	for {
		result := <- out
		for _, item := range result.Iterm{
			fmt.Println("item:",item)
		}
		for _, request := range result.Request{
			this.Scheduler.Submit(request)
		}
	}

}
func createWorker(in chan comment.Request, out chan comment.ParseResult, scheduler Scheduler)  {
	go func() {
		for {
			scheduler.WorkerReady(in)
			request := <-in
			result, err := comment.Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}