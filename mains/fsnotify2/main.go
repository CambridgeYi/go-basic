/*
@Time : 2019/6/7 14:11 
@Author : dong.liu
@File : main
@Software: GoLand
@Description:
*/
package main

import (
	"github.com/fsnotify/fsnotify"
	"go-basic/fsnotify"
)

func main(){
	watch,_:=fsnotify.NewWatcher()
	w:=myfsnotify.Watch{
	Watch:watch,
	}

	w.WatchDir("F:\\code\\other\\go-basic\\mains\\fsnotify2\\tmp")
	select{}
}
