/*
@Time : 2019/7/6 0006 13:30
@Author : dong.liu
@File : main
@Software: GoLand
*/
package main

import "fmt"

func main() {
	dataChan:=make(chan int,5)
	syncChan1:=make(chan struct{},1)
	syncChan2:=make(chan struct{},2)
	// 用于演示接收操作
	go func() {
		<- syncChan1 //阻塞协程，直到有数据
		for {
			if elem,ok:= <-dataChan;ok{
				fmt.Printf("Received:%d [receiver] \n",elem)

			}else{
				break
			}
		}
		fmt.Println("Done.[receiver]")
		syncChan2 <- struct{}{}
	}()
	// 用于演示发送操作
	go func() {
		for i:=0;i<5 ;i++  {
			dataChan <-i
			fmt.Printf("Sent:%d [sender]\n",i)
		}
		// 注意1：对同一个通道仅允许关闭一次,对通道的重复关闭会引发运行时恐慌
		// 注意2：在调用close函数时，需要把代表欲关闭的通道变量作为参数传入，如果该变量为nil，就会引发运行时恐慌
		close(dataChan)
		syncChan1 <- struct{}{}
		fmt.Println("Done.[sender]")
		syncChan2 <- struct{}{}
	}()

	<- syncChan2
	<- syncChan2


}

