/*
@Time : 2019/7/6 0006 11:45
@Author : dong.liu
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)

func main(){
	names :=[]string{"Eric","Robert","Jim","Mark"}

	for _,name :=range names {
		go func(who string){
			fmt.Printf("Hello,%s\n",who)
		}(name)
		// time.Sleep(time.Microsecond)
	}
	time.Sleep(time.Microsecond)
}

