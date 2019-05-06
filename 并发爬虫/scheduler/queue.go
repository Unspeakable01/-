package scheduler

import "troytan/practice/并发爬虫/comment"

//QueueScheduler 接口实现了 scheduler interface
type QueueScheduler struct {
	RequestChannel chan comment.Request
	WorkerChannel	chan chan comment.Request
}
func (this *QueueScheduler) Run(){
	this.RequestChannel = make(chan comment.Request)
	this.WorkerChannel 	= make(chan chan comment.Request)
	go func() {
		var requestQueue	[]comment.Request
		var workerQueue		[]chan comment.Request
		for {
			var activeRequest comment.Request
			var activeWorker chan comment.Request
			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}
			select {
			case request := <-this.RequestChannel:
				requestQueue = append(requestQueue, request)
			case worker := <-this.WorkerChannel:
				workerQueue = append(workerQueue, worker)
			case activeWorker<- activeRequest :
				requestQueue = requestQueue[1:]
				workerQueue	 = workerQueue[1:]
			}
		}
	}()

}
func (this *QueueScheduler) Submit(request comment.Request){
	this.RequestChannel<- request
}
func (this *QueueScheduler) WorkerChan()chan comment.Request{
	return make(chan comment.Request)
}
func (this *QueueScheduler) WorkerReady(workerChan chan comment.Request){
	this.WorkerChannel<- workerChan
}