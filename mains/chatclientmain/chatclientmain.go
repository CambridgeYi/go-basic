/*
@Time : 2019/6/4 17:12 
@Author : dong.liu
@File : chatclientmain
@Software: GoLand
@Description:
*/
package main

import (
	"fmt"
	"net"
	"os"
	"go-basic/chat/chatclient"
)

func main(){

	tcpAddr,_ :=net.ResolveTCPAddr("tcp",":9080")

	//拨号
	conn,_:=net.DialTCP("tcp",nil,tcpAddr)

	var p =make(map[string] interface{})

	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	p["userID"] =os.Args[2]

	go chatclient.SendMessage(conn,p)

	chatclient.ReceiveMessage(conn)
}