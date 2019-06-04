/*
@Time : 2019/6/4 17:22 
@Author : dong.liu
@File : context.go
@Software: GoLand
@Description:
*/
package context

import (
	"context"
	"fmt"
	"time"
)

func StopAllGotoutines(){
	ctx,cancel:=context.WithCancel(context.Background())

	go func(ctx context.Context){
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了")
				return
			default:
				fmt.Println("goroutine监控中")
				time.Sleep(2*time.Second)
			}

		}
	}(ctx)

	time.Sleep(10*time.Second)

	fmt.Println("可以了，通知监控停止")

	cancel()

	time.Sleep(5 * time.Second)
}
