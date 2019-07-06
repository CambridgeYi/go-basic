/*
@Time : 2019/7/6 0006 14:55
@Author : dong.liu
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int ,1)
	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	}()
	select {
		case e:= <-intChan:
			fmt.Printf("Received:%v\n",e)
			case <-time.NewTimer(time.Millisecond*500).C:
				fmt.Printf("方式一：Timeout!")
				case <- time.After(time.Millisecond*300):
					fmt.Printf("方式二：Timeout!")
	}
}

