/*
@Time : 2019/7/6 0006 12:53
@Author : dong.liu
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)

var mapChan =make(chan map[string]int,1)

func main(){
	syncChan :=make(chan struct{},2)
	// 用于演示发送操作
	go func(){
		countMap := make(map[string]int)
		for i:=0;i<5;i++{
			mapChan <-countMap
			fmt.Println(i)
			for k,v :=range countMap{
				fmt.Println(i)
				fmt.Printf("k:%s,v:%d \n",k,v)//只打印了4次 ？？？ k:count,v:1 重复了两次

			}
			time.Sleep(time.Millisecond)

			fmt.Printf("The count map:%v.[sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()

	// 用于演示接收操作
	go func(){
		for {
			if elem,ok := <-mapChan;ok{
				// 因为是引用类型，在这里给map的key赋值为count
				elem["count"]++
			}else{
					break
			}
		}
		fmt.Println("Stopped.[receiver]")

		syncChan <- struct{}{}
	}()

	<- syncChan
	<- syncChan
}

