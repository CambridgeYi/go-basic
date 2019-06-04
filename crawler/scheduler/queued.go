/*
@Time : 2019/6/4 17:28 
@Author : dong.liu
@File : queued
@Software: GoLand
@Description:
*/
package crawler_scheduler

import "go-basic/crawler/model"

type QueuedScheduler struct {
	requestChann chan crawler_model.Request
	// 每一个worker都有一个自己chan request
	// workerChan中存放的是Worker们的chan
	workerChan  chan chan crawler_model.Request
}

func (q *QueuedScheduler) WorkerChann() chan crawler_model.Request{
	return make(chan crawler_model.Request)
}

func (s *QueuedScheduler) Submit(request crawler_model.Request){
	s.requestChann <-request
}

func (s *QueuedScheduler) WorkerReady(w chan crawler_model.Request){
	s.workerChan <-w
}

func (s *QueuedScheduler) Run(){
	// 初始化 request channel
	s.requestChann =make(chan crawler_model.Request)
	// 初始化 workerChan
	s.workerChan =make(chan chan crawler_model.Request)

	// 创建一个goroutine
	// 1.进行request以及worker的chan存储
	// 2.分发request到worker的chan中
	go func(){
		var requestQ []crawler_model.Request
		var workerQ []chan crawler_model.Request

		for{
			var activeRequest crawler_model.Request
			var activeWorker chan crawler_model.Request

			if len(requestQ)>0 &&len(workerQ)>0{
				activeRequest =requestQ[0]
				activeWorker= workerQ[0]
			}

			select{
			case r:= <-s.requestChann:
				// 如果开始requestQ=nil,append之后就是包含一个r元素的切片
				requestQ =append(requestQ,r)
			case w:=<-s.workerChan:
				workerQ=append(workerQ,w)
				// 进行request的分发
			case activeWorker <- activeRequest:
				requestQ =requestQ[1:]
				workerQ=workerQ[1:]
			}
		}
	}()
}