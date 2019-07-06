/*
@Time : 2019/7/6 0006 16:27
@Author : dong.liu
@File : main.go
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwm sync.RWMutex

	for i:=0;i<3 ;i++  {
		go func(i int) {
			fmt.Printf("Try to lock for reading ...[%d]\n",i)
			rwm.RLock()
			fmt.Printf("Locked for reading. [%d]\n",i)
			time.Sleep(time.Second*2)

			fmt.Printf("Try to unlock for reading ...[%d]\n",i)
			rwm.RUnlock()
			fmt.Printf("Unlocked for reading. [%d]\n",i)
		}(i)
	}

	time.Sleep(time.Millisecond*100)
	fmt.Printf("Try to lock for writing...")
	rwm.Lock()
	fmt.Println("Locked for writing")
}

