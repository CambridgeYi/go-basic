/*
@Time : 2019/7/6 0006 16:10
@Author : dong.liu
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex

	fmt.Println("Lock the lock.(main)")

	mutex.Lock()
	fmt.Println("The lock is locked.(main)")

	for i:=1;i<=3 ;i++  {
		go func(i int) {
			fmt.Printf("Lock the lock . (g%d)\n",i)
			mutex.Lock()
			fmt.Printf("The lock is Locked . (g%d)\n",i)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlock the lock . (main)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked. (main)")

	time.Sleep(time.Second)
}

