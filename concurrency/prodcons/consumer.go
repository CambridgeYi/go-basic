/*
@Time : 2019/6/4 17:19 
@Author : dong.liu
@File : consumer.go
@Software: GoLand
@Description:
*/
package prodcons

import "fmt"

//消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}