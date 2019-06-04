/*
@Time : 2019/6/4 16:09 
@Author : dong.liu
@File : jsonrpcServer
@Software: GoLand
@Description:
*/
package jsonrpcServer

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Params struct{
	Width,Height int
}

type Rect struct{}

func (r *Rect) Area(p Params,ret *int) error{
	*ret = p.Width * p.Height
	return nil
}

func (r *Rect) Perimeter(p Params,ret *int) error {
	*ret = (p.Width+p.Height)*2
	return nil
}

func chkError(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func Send(){
	rect:=new(Rect)
	// 注册rpc服务
	rpc.Register(rect)
	// 获取tcpaddr
	tcpaddr,err := net.ResolveTCPAddr("tcp4","127.0.0.1:8180")
	chkError(err)
	// 监听端口
	tcplisten,err2:=net.ListenTCP("tcp",tcpaddr)
	chkError(err2)
	for {
		conn,err3 :=  tcplisten.Accept()
		if err3 != nil{
			continue
		}

		// 使用goroutine单独使用rpc连接请求
		// 这里使用jsonrpc进行处理
		go jsonrpc.ServeConn(conn)
	}


}