/*
@Time : 2019/6/7 14:46 
@Author : dong.liu
@File : main
@Software: GoLand
@Description:
*/
package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"go-basic/fsnotify"
	"log"
	"os"
	"path/filepath"
)

func main(){
	// 创建一个监控对象
	watch,err:=fsnotify.NewWatcher()
	if err !=nil{
		log.Fatal(err)
	}

	defer watch.Close()

	// 添加要监控的文件
	err =watch.Add(myfsnotify.ConfFilePath)
	if err!=nil{
		log.Fatal(err)
	}

	// 我们启动一个goroutine来处理监控对象的事件
	go func(){
		for{
			select{
			case ev :=<- watch.Events:
				{
					// 我们只需关心文件的修改
					if ev.Op&fsnotify.Write == fsnotify.Write{
						fmt.Println(ev.Name,"文件写入")
						// 查找进程
						pid,err:=myfsnotify.GetPid("server.exe")
						// 获取运行文件的绝对路径
						exePath,_:= filepath.Abs("./server.exe")

						if err!=nil{
							// 启动进程
							go myfsnotify.StartProcess(exePath,[]string{})
						} else {
							// 找到进程，并退出
							process,err :=os.FindProcess(pid)
							if err !=nil{
								// 让进程退出
								process.Kill()
								fmt.Println(exePath,"进程退出")

							}
							// 启动进程
							go myfsnotify.StartProcess(exePath,[]string{})
						}
					}
				}
				case err := <- watch.Errors:
					{
						fmt.Println("error: ",err)
						return
					}
			}
		}
	}()

	select{}
}



