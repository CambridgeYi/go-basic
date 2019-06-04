/*
@Time : 2019/6/4 17:48 
@Author : dong.liu
@File : pprof
@Software: GoLand
@Description:参考:https://www.jianshu.com/p/1f7308dda05f
*/
package pprof

import (
	"fmt"
	"net/http"
	"os"
	"time"
	_ "net/http/pprof"
)

func PprofMonitor(){
	// 开启pprof,监听请求
	go func(){
		ip:="0.0.0.0:10000"
		if err:=http.ListenAndServe(ip,nil);err!=nil{
			fmt.Printf("start pprof failed on %s\n",ip)
			os.Exit(1)
		}
	}()

	tick:=time.Tick(time.Second/100)
	fmt.Println(tick)
	var buf []byte
	for range tick{
		buf =append(buf,make([]byte,1024*1024)...)
	}
}