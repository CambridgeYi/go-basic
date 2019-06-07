/*
@Time : 2019/6/7 12:08 
@Author : dong.liu
@File : websocket
@Software: GoLand
@Description:
*/
package mywebsocket

import (
	"fmt"
	"golang.org/x/net/websocket"
)

type Msg struct{
	From string
	To string
	Data string
}

// 处理简单字符串
func Test(conn *websocket.Conn){
	var err error
	for{
		data:=""
		err =websocket.Message.Receive(conn,&data)
		if err!=nil{
			break
		}

		fmt.Println("client send:"+data)

		msg:="hello"+data

		// 发送消息
		err = websocket.Message.Send(conn,msg)

		if err !=nil{
			break
		}
	}
}

// 处理json数据
func Test2(conn *websocket.Conn){
	var err error
	for{
		var data Msg

		// 接收消息
		err =websocket.JSON.Receive(conn,&data)
		if err!=nil{
			break
		}
		fmt.Println(data.From,data.To,data.Data)

		msg:=Msg{
			From:data.From,
			To:data.To,
			Data:"hello "+data.Data,
		}

		// 发送消息
		err =websocket.JSON.Send(conn,msg)
		if err !=nil{
			break
		}
	}
}