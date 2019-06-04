/*
@Time : 2019/6/4 17:09 
@Author : dong.liu
@File : client
@Software: GoLand
@Description:
*/
package chatclient

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// 发送消息到服务器
func SendMessage(conn net.Conn,p map[string]interface{}){
	for{
		var input string

		// 接收输入消息，放到input变量中
		fmt.Scanln(&input)

		// 用户端退出
		if input=="/q"||input=="/quit"{
			fmt.Println("ByeBye....")
			conn.Close()
			os.Exit(0)
		}
		// 给消息赋值
		p["mytime"] =time.Now()
		p["message"] =input

		// 只处理有内容的消息
		if len(strings.TrimSpace(input))>0{
			msg,err:=json.Marshal(p)
			if err!=nil{
				fmt.Println("序列化错误",err)
			}else{
				// 注意这里需要将interface转为string
				_,err:=conn.Write([]byte(string(msg)))
				//没有发送成功
				if err!= nil{
					conn.Close()
					break
				}
			}

		}
	}
}
/*
	接受服务器消息
*/
func ReceiveMessage(conn net.Conn){
	// 接受服务器的广播消息
	buf :=make([]byte,1024)
	for{
		length, err := conn.Read(buf)
		if err!=nil{
			log.Printf("recv server msg failed:%v\n",err)
			conn.Close()
			os.Exit(0)

			break
		}
		fmt.Println(string(buf[:length]))
	}
}

