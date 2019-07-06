/*
@Time : 2019/7/6 0006 16:36
@Author : dong.liu
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var countVal atomic.Value
	countVal.Store([]int{1,3,5,7})

	anotherStore(countVal)

	fmt.Printf("the count value :%+v \n",countVal.Load())

}

func anotherStore(countVal atomic.Value){
	countVal.Store([]int{2,4,6,8})
}
