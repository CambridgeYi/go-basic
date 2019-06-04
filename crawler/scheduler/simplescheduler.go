/*
@Time : 2019/6/4 17:32 
@Author : dong.liu
@File : simplescheduler
@Software: GoLand
@Description:
*/
package crawler_scheduler

import "go-basic/crawler/model"

type SimpleScheduler struct{
	workerChan chan crawler_model.Request
}

// 为什么使用指针接收者，需要改变 SimpleScheduler 内部的 workerChan
func (s *SimpleScheduler) ConfigureMasterWorkerChan(in chan crawler_model.Request){
	s.workerChan = in
}

func (s *SimpleScheduler) Submit(request crawler_model.Request){
	// 每个request一个Goroutine
	go func(){
		s.workerChan <-request
	}()
}