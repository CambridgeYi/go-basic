/*
@Time : 2019/6/4 17:31 
@Author : dong.liu
@File : scheduler
@Software: GoLand
@Description:调度器接口
*/
package crawler_scheduler

import "go-basic/crawler/model"

// 调度器接口
type Scheduler interface {
	// 提交Request 到调度器的request任务通道中
	Submit(request crawler_model.Request)
	// 初始化当前的调度器实例的 request 任务通道
	ConfigureMasterWorkerChan(chan crawler_model.Request)
}