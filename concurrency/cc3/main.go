/*
@Time : 2019/7/6 0006 12:17
@Author : dong.liu
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main() {
	// syncChan1,syncChan2的元素类型都是struct{}，struct{}代表的是不包含任何字段的结构体类型，也可称为空结构体类型
	// 在Go语言中，空结构体类型的变量是不占用内存空间的，并且所有该类型的变量都拥有相同的内存地址，建议用于传递“信号”的通道都以struct{}作为元素类型，
	// 注意1：当向一个值为nil的通道类型的变量发送元素值时，当前goroutine会永久的阻塞
	// 注意2：如果试图向一个已关闭的通道发送元素值，那么会立即引发一个运行时恐慌，即使发送操作正在因通道已满而被阻塞,为了避免这样的流程中断，可以在select代码块中执行发送操作
	syncChan1 := make(chan struct{}, 1)
	// 这个通道纯粹是为了不让主goroutine过早结束运行
	syncChan2 := make(chan struct{}, 2)
	// 用于演示接收操作
	go func() {
		<-syncChan1
		fmt.Println("Received a sync signal and wait a second...[receiver]")

		time.Sleep(time.Second)

		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Received:", elem, "[receiver]")
			} else {
				break
			}
		}

		fmt.Println("Stopped.[receiver]")
		syncChan2 <- struct{}{} // 写入channel两次
	}()
	// 用于演示发送操作
	go func() {
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Send:", elem, "[sender]")
			if elem == "c" {
				// 向syncChan1发送值用于唤醒syncChan1所在的协程
				syncChan1 <- struct{}{}
				fmt.Println("send a sync signal. [sender]")

			}
		}
		fmt.Println("wait 2 seconds...[sender]")
		time.Sleep(time.Second * 2)
		close(strChan)
		syncChan2 <- struct{}{} // 写入channel一次
	}()
	// 主线程试图从syncChan2接收值两次，在这两次接收都成功完成之前，主goroutine会阻塞于此，一旦那两个goroutine都向syncChan2发送了值，主goroutine就会恢复运行，但随后又会结束运行
	<-syncChan2
	<-syncChan2
}
