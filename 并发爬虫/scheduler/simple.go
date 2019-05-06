package scheduler

import "troytan/practice/并发爬虫/comment"

//SimpleScheduler 实现了 Scheduler 接口
type SimpleScheduler struct {
	WorkerChannel chan comment.Request
}
func (this *SimpleScheduler)Run(){
	this.WorkerChannel = make(chan comment.Request)
}
func (this *SimpleScheduler)Submit(request comment.Request){
	go func() {
		this.WorkerChannel <- request
	}()
}
func (this *SimpleScheduler) WorkerChan() chan comment.Request {
	return this.WorkerChannel
}
func (this *SimpleScheduler) WorkerReady(chan comment.Request) {

}