/*
@Time : 2019/6/4 17:36 
@Author : dong.liu
@File : windowlimiter
@Software: GoLand
@Description:
*/
package limit

import (
	"fmt"
	"github.com/UncleBig/goCache"
	"go-basic/util"
	"net"
	"time"
)

// https://www.jianshu.com/p/9032c6f41f1e

var (
	limit int =10
	c *goCache.Cache
)

func WindowLimit(){
	c =goCache.New(10*time.Minute,30*time.Second)
	// 获取一个tcpAddr
	tcpAddr,err:=net.ResolveTCPAddr("tcp4","0.0.0.0:10002")
	util.CheckError(err)
	// 监听一个端口
	listener,err:=net.ListenTCP("tcp",tcpAddr)
	util.CheckError(err)

	defer listener.Close()

	for{
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println(err)
			continue
		}
		go handleWindow(&conn)
	}
}

func handleWindow(conn *net.Conn){
	defer (*conn).Close()
	t:=time.Now().Unix()
	key:=fmt.Sprintf("%d",t)
	if n,found:=c.Get(key);found{
		num:=n.(int)

		fmt.Printf("key:%d num:%d\n",t,num)
		if num >=limit{
			(*conn).Write([]byte("HTTP/1.1 404 NOT FOUND\r\n\r\nError, too many request, please try again."))
		}else{
			(*conn).Write([]byte("HTTP/1.1 200 OK\r\n\r\nI can change the world!"))
			c.Increment(key,1)
		}
	}else{
		(*conn).Write([]byte("HTTP/1.1 200 OK\r\n\r\nI can change the world!"))
		c.Set(key,1,2*time.Second)
	}
}