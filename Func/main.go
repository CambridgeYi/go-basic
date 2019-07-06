/*
@Time : 2019/7/6 0006 12:00
@Author : dong.liu
@File : main
@Software: GoLand
*/
package main

import "fmt"

func swap(a *int, b *int){
	var temp int
	temp = *a
	*a = *b
	*b = temp

}
func function( a,b int , sum func(int,int) int){
	fmt.Print(sum(a,b))
}

func sum(a,b int) int {
	return a+b
}

func main(){
	x :=5
	y :=10
	swap(&x,&y)
	fmt.Print(x,y)
	fmt.Print("\n")
	var a, b int =5,6
	f:=sum

	function(a,b,f)
}

