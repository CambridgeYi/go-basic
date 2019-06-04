/*
@Time : 2019/6/4 17:30 
@Author : dong.liu
@File : queuescheduler
@Software: GoLand
@Description:
*/
package crawler_scheduler

import "go-basic/crawler/model"

// 队列调度器接口
type QueueScheduler interface {

	ReadyNotifier
	Submit(request crawler_model.Request)
	WorkerChann() chan crawler_model.Request
	Run()
}
