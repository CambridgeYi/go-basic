/*
@Time : 2019/6/7 13:50 
@Author : dong.liu
@File : mynotify2
@Software: GoLand
@Description:
*/
package myfsnotify

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
)

type Watch struct{
	Watch *fsnotify.Watcher
}

// 监控目录
func (w *Watch) WatchDir(dir string){
	// 通过walk来遍历目录下的所有子目录
	filepath.Walk(dir,func(path string,info os.FileInfo,err error) error{
		// 这里判断是否是目录，只需监控目录即可
		// 目录下的文件也在监控范围内，不需要我们一个一个加
		if info.IsDir(){
			path,err:=filepath.Abs(path)
			if err!=nil{
				return err
			}
			err =w.Watch.Add(path)
			if err!=nil{
				return err
			}
			fmt.Println("监控：",path)
		}
		return nil
	})

	go func(){
		for{
			select {
			case ev:= <- w.Watch.Events:
				{
					if ev.Op&fsnotify.Create ==fsnotify.Create{
						fmt.Println("创建问价：",ev.Name)
						// 这里获取新创建文件的信息，如果是目录，则加入监控中
						fi,err:=os.Stat(ev.Name)
						if err ==nil&&fi.IsDir(){
							w.Watch.Add(ev.Name)
							fmt.Println("添加监控：",ev.Name)
						}
					}
					if ev.Op&fsnotify.Write==fsnotify.Write{
						fmt.Println("写入文件:",ev.Name)
					}
					if ev.Op&fsnotify.Remove==fsnotify.Remove{
						fmt.Println("删除文件:",ev.Name)
						// 如果删除文件是目录，则移除监控
						fi,err:=os.Stat(ev.Name)
						if err !=nil && fi.IsDir(){
							w.Watch.Remove(ev.Name)
							fmt.Println("删除监控:",ev.Name)
						}
					}
					if ev.Op&fsnotify.Rename ==fsnotify.Rename{
						fmt.Println("重命名文件:",ev.Name)
						// 如果重命名文件是目录，则移除监控
						// 注意这里无法使用os.Stat来判断是否是目录了
						// 因为重命名后，go已经无法找到原文件来获取信息了
						w.Watch.Remove(ev.Name)
					}
					if ev.Op&fsnotify.Chmod==fsnotify.Chmod{
						fmt.Println("修改权限：",ev.Name)
					}
				}
				case err:= <- w.Watch.Errors:
				{
					fmt.Println("error: ",err)
					return
				}
			}
		}
	}()
}