/*
@Time : 2019/6/4 17:27 
@Author : dong.liu
@File : fetcher
@Software: GoLand
@Description:
*/
package crawler_fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte,error){
	// 1.爬取url
	resp,err :=http.Get(url)
	if err!=nil{
		return nil,err
	}

	defer resp.Body.Close()

	if resp.StatusCode !=http.StatusOK{
		return nil,fmt.Errorf("wrong statusCode,%d",resp.StatusCode)
	}
	// 2. 读取响应体并返回
	return ioutil.ReadAll(resp.Body)
}