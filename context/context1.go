/*
@Time : 2019/6/4 17:23 
@Author : dong.liu
@File : context1
@Software: GoLand
@Description:
*/
package context


import (
	"context"
	"fmt"
	"time"
)

func StopAllGoroutineByOne(){
	ctx,cancel:=context.WithCancel(context.Background())

	go watch(ctx,"监控1")
	go watch(ctx,"监控2")
	go watch(ctx,"监控3")

	time.Sleep(10*time.Second)

	fmt.Println("可以了，通知监控停止")

	cancel()


	time.Sleep(5*time.Second)
}

func watch(ctx context.Context,name string){
	for{
		select {
		case <- ctx.Done():
			fmt.Println(name,"监控退出，停止了...")
			return
		default:
			fmt.Println(name,"goroutine 监控中...")
			time.Sleep(2*time.Second)
		}
	}
}
