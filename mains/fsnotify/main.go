/*
@Time : 2019/6/7 13:09 
@Author : dong.liu
@File : main
@Software: GoLand
@Description:
*/
package main

import (
	"fmt"
	"go-basic/fsnotify"
	"go-basic/path"
)

func main(){
	path,_ :=path.ExecPath()
	fmt.Println("执行路径：",path)
	myfsnotify.Notify()
}


