/*
@Time : 2019/6/4 17:34 
@Author : dong.liu
@File : ratelimiter
@Software: GoLand
@Description:https://www.jianshu.com/p/9032c6f41f1e
*/
package limit

import (
	"fmt"
	"go-basic/util"
	"net"
	"sync/atomic"
	"time"
)

var (
	limiting int32 =10 //令牌桶
)

func LimitRate(){
	// 获取一个tcpAddr
	tcpAddr,err:=net.ResolveTCPAddr("tcp4","0.0.0.0:10001")
	util.CheckError(err)

	// 监听一个端口
	listener,err:=net.ListenTCP("tcp",tcpAddr)
	Util.CheckError(err)

	defer listener.Close()
	for{
		// 在此处阻塞，每次来一个请求才往下运行handle函数
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println(err)
			continue
		}
		// 起一个单独的协程处理，有多少个请求，就起多少个协程，协程之间共享同一个全量变量limiting,
		// 对其进行原子操作
		go handle(&conn)
	}
}

func handle(conn *net.Conn){
	defer (*conn).Close()
	// dcr 1 by atomic 获取一个令牌，总数减1
	// 这是一个原子性的操作，并发操作下数据不会写错
	n:=atomic.AddInt32(&limiting,-1)

	if n<0{
		// 令牌不够用了，限流，抛弃此次请求
		(*conn).Write([]byte("HTTP/1.1 404 NOT FOUND\r\n\r\nError, too many request, please try again."))
	}else{
		// 还有剩余令牌可用 假设我们的应用处理业务用了1s的时间
		time.Sleep(10*time.Second)
		// 业务处理结束后，回复200成功
		(*conn).Write([]byte("HTTP/1.1 200 OK\r\n\r\nI can change the world!"))
	}
	// add 1 by atomic 业务处理完毕，返回令牌
	atomic.AddInt32(&limiting,1)
}
