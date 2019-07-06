/*
@Time : 2019/7/6 0006 16:39
@Author : dong.liu
@File : main.go
@Software: GoLand
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var count int
	var once sync.Once
	max := rand.Intn(100)

	for i:=0;i<max ;i++  {
		once.Do(func() {
			count++
		})
	}
	fmt.Printf("Count:%d.\n",count)
}
