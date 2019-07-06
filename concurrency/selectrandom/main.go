/*
@Time : 2019/7/6 0006 14:18
@Author : dong.liu
@File : main
@Software: GoLand
*/
package main

import "fmt"

func main() {
	chanCap :=5
	intChan :=make(chan int ,chanCap)

	for i:=0;i<chanCap ;i++  {
		select {
		case intChan <- 1:
		case intChan<-2:
		case intChan <-3:
		}

	}
	for i:=0;i< chanCap ;i++  {
		fmt.Printf("%d\n", <-intChan)
	}
}

