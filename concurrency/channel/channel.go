/*
@Time : 2019/6/4 17:21 
@Author : dong.liu
@File : channel.go
@Software: GoLand
@Description:
*/
package channel

import (
	"fmt"
	"sync"
)

func PrintByNoBufferChannel(){
	done:=make(chan int)

	go func(){
		fmt.Println("no buffer channel 你好，世界")
		<- done
	}()

	done <- 1
}

func PrintByBufferChannel(){
	done :=make(chan int,1)

	go func(){
		fmt.Println("buffer channel 你好，世界")
		done <- 1
	}()

	<- done
}

func PrintMultiByBufferChannel(){
	done :=make(chan int,10) //带十个缓存

	//开N个后台打印协程
	for i:=0;i<cap(done);i++{
		go func(i int){
			fmtInfo :=fmt.Sprintf("multi buffer channels : %d 你好，世界 \n",i)

			fmt.Printf(fmtInfo)
			done <- 1
		}(i)
	}

	//等待N个后台协程完成
	for i:=0;i<cap(done);i++{
		<- done
	}
}

func PrintGroupByBufferChannel(){
	var wg  sync.WaitGroup

	//开启N个后台打印协程
	for i:=0;i<10;i++{
		wg.Add(1)

		go func(i int){
			fmtInfo:=fmt.Sprintf("sync.WaitGroup:%d 你好，世界 \n",i)
			fmt.Println(fmtInfo)
			wg.Done()
		}(i)

		//等待N个后台协程完成
		wg.Wait()
	}
}