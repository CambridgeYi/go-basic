/*
@Time : 2019/6/4 17:31 
@Author : dong.liu
@File : readynotifier
@Software: GoLand
@Description:
*/
package crawler_scheduler

import "go-basic/crawler/model"

type ReadyNotifier interface {
	WorkerReady(chan crawler_model.Request)
}