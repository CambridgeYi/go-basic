/*
@Time : 2019/7/6 0006 13:42
@Author : dong.liu
@File : main.go
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)

var strChan =make(chan string,3)

func main() {
	syncChan1:=make(chan struct{},1)
	syncChan2:=make(chan struct{},2)

	go receive(strChan,syncChan1,syncChan2)  // 用于演示接收操作
	go send(strChan,syncChan1,syncChan2)  // 用于演示发送操作

	// 阻塞主线程
	<- syncChan2
	<- syncChan2

}

func receive(strChan <-chan string,
	syncChan1 <-chan struct{},
	syncChan2 chan<- struct{}){
	<- syncChan1
	fmt.Printf("Received a sync signal and wait a second...[receiver]")
	time.Sleep(time.Second)

	// 方式一:
	for {
		if elem,ok:= <- strChan;ok{
			fmt.Println("Received:",elem,"receiver")
		}else {
			break
		}
	}
	// 方式二:chan通道数据被读取后就没有了，channel是复制引用，至少复制一次，最多复制两次
	for e :=range strChan{
		fmt.Println("Received1111:",e,"receiver")
	}

	fmt.Println("Stopped.[receiver]")
	syncChan2 <- struct{}{}
}

func send(strChan chan<- string,
	syncChan1 chan<- struct{},
	syncChan2 chan<- struct{}){
	for _,elem :=range []string{"a","b","c","d"}{
		strChan <-elem
		fmt.Println("Sent:",elem,"[sender]")

		if elem =="c"{
			syncChan1 <- struct{}{}
			fmt.Println("Sent a sync signal.[sender]")
		}
	}
	fmt.Println("wait 2 seconds ... [sender]")
	time.Sleep(time.Second*2)
	close(strChan)

	syncChan2 <- struct{}{}
}