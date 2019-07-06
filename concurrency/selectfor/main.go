/*
@Time : 2019/7/6 0006 14:26
@Author : dong.liu
@File : main
@Software: GoLand
*/
package main

import "fmt"

func main() {
	initChan :=make(chan int,10)

	for i:=0;i<10;i++{
		initChan <- i
	}
	close(initChan)

	syncChan := make(chan struct{},1)

	go func() {

		Loop:
			for {
				select {
				case e,ok := <-initChan:
					if !ok	{
						fmt.Println("End.")
						// 这是一个带标签的break ,Loop为标签的名字，意为中断紧贴于该标签之下的那条语句的执行
						break Loop
					}
					fmt.Printf("Received:%v\n",e)
				}
			}
		syncChan <- struct{}{}
	}()

	<- syncChan
}

