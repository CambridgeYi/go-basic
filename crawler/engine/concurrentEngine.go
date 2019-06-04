/*
@Time : 2019/6/4 17:24 
@Author : dong.liu
@File : concurrentEngine
@Software: GoLand
@Description:https://www.jianshu.com/p/cd7b55d401b9
*/
package crawler_engine

import (
	"crawler/fetcher"
	"crawler/model"
	"crawler/scheduler"
	"fmt"
	"log"
)

// 并发引擎
type ConcurrentEngine struct{
	// 调度器
	Scheduler scheduler.Scheduler
	// 开启的worker 数量
	WorkCount int
}

func (e ConcurrentEngine) Run(seeds ...model.Request){
	in:=make(chan model.Request,10)
	out:=make(chan model.ParseResult,10)

	// 初始化调度器的chan
	e.Scheduler.ConfigureMasterWorkerChan(in)

	// 创建WorkerCount个worker
	for i:=0;i<e.WorkCount;i++{
		createWorker(in,out)
	}

	// 将seeds中的Request添加到调度器 chan
	for _,r:=range seeds{
		e.Scheduler.Submit(r)
	}

	for{
		// 获取解析结果
		result:= <- out

		for _,item :=range result.Items{
			fmt.Printf("getItems,items:%v \n",item)
		}

		for _,r:=range result.Requests{
			e.Scheduler.Submit(r)
		}
	}
}


// 开启工作协程 ，处理请求队列，返回响应到响应队列
func createWorker(in chan model.Request,out chan model.ParseResult){
	go func(){
		for{
			//从请求队列中取出请求
			r:= <- in

			// 处理请求
			result,err:=handle(r)

			if err!=nil{
				continue
			}
			// 响应结果写入到解析队列中
			out <- result
		}

	}()
}

func handle(r model.Request)(model.ParseResult,error){
	log.Printf("fetching url:%s",r.Url)

	body,err:=fetcher.Fetch(r.Url)

	if err!=nil{
		log.Printf("fetch error,url:%s,err:%v",r.Url,err)
		return model.ParseResult{},nil
	}

	return r.ParseFunc(body),nil

}