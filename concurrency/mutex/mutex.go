/*
@Time : 2019/6/4 17:22 
@Author : dong.liu
@File : mutex
@Software: GoLand
@Description:
*/
package mutex

import (
	"fmt"
	"sync"
)

func  PrintMutex(){
	var mu sync.Mutex

	mu.Lock()
	go func(){
		fmt.Println("sync.Mutext 你好，世界")
		mu.Unlock()
	}()

	mu.Lock()
}