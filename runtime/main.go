/*
@Time : 2019/7/6 0006 11:52
@Author : dong.liu
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main(){
	fmt.Printf("系统中P的大小:%d",runtime.GOMAXPROCS)
	debug.FreeOSMemory()
}

