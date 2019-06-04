/*
@Time : 2019/6/4 17:54 
@Author : dong.liu
@File : main
@Software: GoLand
@Description:
*/
package main

import (
	"fmt"
	"go-basic/concurrency/channel"
	"go-basic/concurrency/mutex"
	"go-basic/concurrency/prodcons"
	"go-basic/concurrency/pubsub"
	"go-basic/mytime"
	"strings"
	"time"
)

func main(){

	/*Go 高级编程*/
	// 使用sync.Mutex实现并发
	mutex.PrintMutex()

	// 使用 no buffer channel实现并发
	channel.PrintByNoBufferChannel()

	// 使用buffer channel实现并发
	channel.PrintByBufferChannel()

	// 多个buffer channel 实现并发
	channel.PrintMultiByBufferChannel()

	// 使用sync.WaitGroup实现并发
	channel.PrintGroupByBufferChannel()

	//producer和consumer 实现并发
	ch := make(chan int, 1)
	go prodcons.Producer(3, ch)
	go prodcons.Producer(5, ch)
	go prodcons.Consumer(ch)

	//publish和subscribe发布/订阅 实现并发
	// 创建订阅者
	p:=pubsub.NewPublisher(100*time.Microsecond,10)

	defer p.Close()

	// 创建全部主题和含有"golang"的主题
	all := p.Subscribe()
	golang:=p.SubscribeTopic(func(v interface{}) bool {
		if s,ok:=v.(string);ok{
			return strings.Contains(s,"golang")
		}
		return false
	})

	p.Publish("pub/sub hello,world!")
	p.Publish("pub/sub hello, golang!")

	go func(){
		for msg:=range all{
			fmt.Println("all:",msg)
		}
	}()

	go func(){
		for msg :=range golang{
			fmt.Println("golang:",msg)
		}
	}()
	/*
	//将HelloService类型的对象注册为一个RPC服务
	rpc.RegisterName("HelloService",new(Rpc.HelloService))

	listener,err:=net.Listen("tcp",":1234")

	if err!=nil{
		log.Fatal("ListenTCP error:",err)
	}

	conn,err:=listener.Accept()

	if err!=nil{
		log.Fatal("Accept error:",err)
	}

	rpc.ServeConn(conn)
	*/

	// 运行一段时间后退出
	//mytime.Sleep(5 * mytime.Second)*/

	// Question 1
	//Question.Defer_call()
	// Question 2
	//Question.Pase_Student()

	// context 管理多个goroutine运行状态
	//Context.StopAllGotoutines()
	//Context.StopAllGoroutineByOne()
	//Context.StopAllGoroutineByKey()

	// 获取tag标签
	// tag.GetAllTagValues()

	// 开启pprof,监听请求
	// pprof.PprofMonitor()

	// rate limit
	// limit.LimitRate()

	// window rate limit
	//limit.WindowLimit()

	// 定时器
	mytime.Run()
}
