/*
@Time : 2019/6/4 17:42 
@Author : dong.liu
@File : main
@Software: GoLand
@Description:
*/
package ginmain

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-basic/docs"
	"go-basic/log"
	logbase "log"
	"net/rpc"
	"time"
	"go-basic/router"
)

func init(){
	log.ConfigFileLogger("logs","log",time.Hour*24*365,time.Hour*24)

}
// @title 测试
// @version 0.0.1
// @description  测试
func main() {

	g := gin.Default()

	// use ginSwagger middleware to
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.GET("/insert",router.Insert)

	//rpc客户端 服务端在gobase项目中
	client,err:=rpc.Dial("tcp","localhost:1234")
	if err!=nil{
		logbase.Fatal("dialing:",err)
	}

	var reply string
	err= client.Call("HelloService.Hello","hello",&reply)

	if err!=nil{
		logbase.Fatal(err)
	}

	logbase.Println(reply)

	g.Run()
}
