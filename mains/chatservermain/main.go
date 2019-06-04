/*
@Time : 2019/6/4 17:17 
@Author : dong.liu
@File : main
@Software: GoLand
@Description:
*/
package main

import "go-basic/chat/chatserver"

func main(){
	//启动服务器
	chatserver.StartServer("9080")
}
