/*
@Time : 2019/6/7 15:47 
@Author : dong.liu
@File : main
@Software: GoLand
@Description:
*/
package main

import (
	"fmt"
	"go-basic/delay"
	"time"
)

func main(){
	// 创建延迟消息
	dm :=delay.NewDelayMessage()
	// 添加任务
	dm.AddTask(time.Now().Add(time.Second*10),"test1",func(args ...interface{}){
		fmt.Println(args...)
	},[]interface{}{1,2,3})
	dm.AddTask(time.Now().Add(time.Second*10),"test2",func(args ...interface{}){
		fmt.Println(args...)
	},[]interface{}{4,5,6})
	dm.AddTask(time.Now().Add(time.Second*20),"test3",func(args ...interface{}){
		fmt.Println(args...)
	},[]interface{}{"hello","world","test"})
	dm.AddTask(time.Now().Add(time.Second*30),"test4",func(args ...interface{}){
		sum :=0
		for arg :=range args{
			sum +=arg
		}
		fmt.Println("sum: ",sum)
	},[]interface{}{1,2,3})

	// 40s后关闭
	time.AfterFunc(time.Second*40,func(){
		dm.Close()
	})
	dm.Start()
}
