/*
@Time : 2019/6/4 15:25 
@Author : dong.liu
@File : httpClient
@Software: GoLand
@Description:
*/
package httpClient

import (
	"fmt"
	"log"
	rpc2 "net/rpc"
)

type Params struct{
	Width,Height int
}

func Send(){
	// 连接远程rpc服务
	rpc,err:=rpc2.DialHTTP("tcp","127.0.0.1:8080")
	if err!=nil{
		log.Fatal(err)
	}
	ret:=0
	// 调用远程方法
	// 注意第三个参数是指针类型
	err2:=rpc.Call("Rect.Area",Params{50,100},&ret)

	if err2!=nil{
		log.Fatal(err2)
	}

	fmt.Println(ret)

	err3 :=rpc.Call("Rect.Perimeter",Params{50,100},&ret)

	if err3 !=nil{
		log.Fatal(err3)
	}
	fmt.Println(ret)
}


