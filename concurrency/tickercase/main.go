/*
@Time : 2019/7/6 0006 15:16
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
	intChan :=make(chan int ,1)

	ticker:= time.NewTicker(time.Second)

	go func() {
		for _=range ticker.C{
			select {
			case intChan <-1:
				case intChan <- 2:
					case intChan <-3:
			}
		}

		fmt.Println("End. [sender]")
	}()

	var sum int
	for e:=range intChan{
		fmt.Printf("Received:%v\n",e)
		sum +=e

		if sum >10 {
			fmt.Printf("Got :%v\n",sum)
			break;
		}

	}
	ticker.Stop()
	fmt.Println("End.[Receicer]")

}
