/*
@Time : 2019/7/6 0006 13:15
@Author : dong.liu
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)

// Counter代表计数器的类型
type Counter struct{
	count int
}

var mapChan = make(chan map[string]Counter,1)

func main(){
	syncChan := make(chan struct{},2)

	// 用于演示接收操作
	go func() {
		for {
			if elem,ok :=<- mapChan;ok{
				counter:=elem["count"]
				counter.count++
			}else{
					break;
			}
			fmt.Println("stopped.[receiver]")
			syncChan <- struct{}{}
		}
	}()
	// 用于演示发送操作
	go func() {
		countMap := map[string]Counter{
			"count":Counter{},
		}

		for i:=0;i<5;i++{
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map:%v.[sender]\n",countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<- syncChan
	<- syncChan
}

