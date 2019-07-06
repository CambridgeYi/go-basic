/*
@Time : 2019/7/6 0006 14:51
@Author : dong.liu
@File : main.go
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2*time.Second)

	fmt.Printf("Present time: %v.\n",time.Now())

	expirationTime := <-timer.C
	fmt.Printf("Expiration time :%v.\n",expirationTime)
	fmt.Printf("Stop timer:%v.\n",timer.Stop())
}

