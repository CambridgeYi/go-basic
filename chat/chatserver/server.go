/*
@Time : 2019/6/4 17:17 
@Author : dong.liu
@File : server
@Software: GoLand
@Description:
*/
package chatserver

import (
	"fmt"
	"github.com/json-iterator/go"
	"net"
)

type person struct {
	news string
	ip   string
}

func StartServer(port string) {

	serverTcp, err := net.ResolveTCPAddr("tcp", ":"+port)

	if err != nil {
		fmt.Printf("resolve tcp addr failed:%v\n", err)
		return
	}
	// 保存连接到服务器的连接池
	conns := make(map[string]net.Conn)
	// 消息通道
	messages := make(chan person, 10)
	// 监听
	listener, err := net.ListenTCP("tcp", serverTcp)
	if err != nil {
		fmt.Printf("listen tcp port failed:%v\n", err)
	}

	//开启协程，将信道中的消息广播出去
	go broadCastMessages(conns, messages)

	for {
		// 监听客户端连接
		clientTcp, err := listener.AcceptTCP()

		if err != nil {
			fmt.Println("连接失败")
			continue
		}
		// 获取客户端地址
		clientAddr := clientTcp.RemoteAddr().String()
		// 添加所有连接
		conns[clientAddr] = clientTcp

		// 处理客户端发来的消息到信道中
		go handleMessage(clientTcp, conns, messages)
	}

}

/*
	conns:连接池
	messages:消息信道
*/
func handleMessage(conn net.Conn, conns map[string]net.Conn, messages chan person) {
	// 暂存消息
	buf := make([]byte, 1000)
	for {
		// 从连接中读取数据，写到buffer中
		len, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			delete(conns, conn.RemoteAddr().String())
			break
		}
		// 将消息写入到结构体
		p := person{
			news: string(buf[:]),
			ip:   conn.RemoteAddr().String(),
		}

		// 发送消息到信道
		messages <- p
		fmt.Println(string(buf[:len]))
	}
}

/*
	conns:连接池
	messages:消息信道
*/
func broadCastMessages(conns map[string]net.Conn, messages chan person) {
	for {
		//读取信道里面的消息
		msg := <-messages

		//发送除了自己的其他所有人
		for k, v := range conns {
			if k != msg.ip {
				// 将发送过来的消息news反序列化为map
				//var m map[string]string
				if len(msg.news) > 0 {
					// json.Unmarshal([]byte(msg.News), &m) m老是解析为nil
					time := jsoniter.Get([]byte(msg.news), "mytime")
					userID := jsoniter.Get([]byte(msg.news), "userID")
					message := jsoniter.Get([]byte(msg.news), "message")
					//拼接消息,广播出去,取值要与客户端一致，不然解析不出来
					broadCastMsg := time.ToString() + "\n" + userID.ToString() +
						":" + message.ToString()
					// 发送
					_, err := v.Write([]byte(broadCastMsg))
					if err != nil {
						delete(conns, k)
						v.Close()
						continue
					}
				}
			}
		}
	}

}