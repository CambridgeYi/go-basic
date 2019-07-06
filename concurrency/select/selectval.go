/*
@Time : 2019/7/6 0006 14:08
@Author : dong.liu
@File : selectval
@Software: GoLand
*/
package main

import "fmt"

// 注意 通道intChan1,intChan2都未被初始化，向他们发送元素值的操作将被永久阻塞
var intChan1 chan int
var intChan2 chan int
var channels =[]chan int {intChan1,intChan2}

var numbers = []int{1,2,3,4,5}
func main() {
	select {
	case getChan(0) <-getNumber(0):
		fmt.Println("1th case is selected")
	case getChan(1) <- getNumber(1):
			fmt.Println("2th case is selected")
	default:
		fmt.Println("Default")
	}
}

func getNumber(i int) int{
	fmt.Printf("numbers[%d]\n",i)
	return numbers[i]
}

func getChan(i int) chan int{
	fmt.Printf("channels[%d]\n",i)
	return channels[i]
}