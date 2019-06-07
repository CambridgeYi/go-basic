/*
@Time : 2019/6/7 12:18 
@Author : dong.liu
@File : main
@Software: GoLand
@Description:
*/
package main

import (
	"go-basic/websocket"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func main(){
	// http.Handle("/",websocket.Handler(mywebsocket.Test))

	http.Handle("/",websocket.Handler(mywebsocket.Test2))

	err:=http.ListenAndServe(":8080",nil)

	if err !=nil{
		log.Fatal(err)
	}
}
