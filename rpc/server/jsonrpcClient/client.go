/*
@Time : 2019/6/4 16:18 
@Author : dong.liu
@File : client
@Software: GoLand
@Description:
*/
package jsonrpcClient

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Params struct{
	Width,Height int
}

func Send(){
	// 连接远程rpc服务
	// 这里使用jsonrpc.Dial
	rpc,err:=jsonrpc.Dial("tcp","127.0.0.1:8180")
	if err !=nil{
		log.Fatal(err)
	}
	ret:=0
	// 调用远程方法
	// 注意第三个参数是指针类型
	err2 :=rpc.Call("Rect.Area",Params{10,20},&ret)
	if err2 !=nil{
		log.Fatal(err2)
	}

	fmt.Println(ret)

	err3:=rpc.Call("Rect.Perimeter",Params{5,12},&ret)
	if err3 !=nil{
		log.Fatal(err3)
	}

	fmt.Println(ret)
}
