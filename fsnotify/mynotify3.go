/*
@Time : 2019/6/7 14:30 
@Author : dong.liu
@File : my
@Software: GoLand
@Description:
*/
package myfsnotify

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

type Conf struct {
	Port int `json:"port"`
}

const (
	ConfFilePath ="./conf"
)
// 获取进程ID
func GetPid(processName string)(int,error){
	// 通过wmic process get name,processid |findstr server.exe获取进程ID
	buf :=bytes.Buffer{}

	cmd:=exec.Command("wmic","process","get","name,processid")
	cmd.Stdout =&buf
	cmd.Run()

	cmd2 :=exec.Command("findstr",processName)
	cmd2.Stdin =&buf
	data,_ :=cmd2.CombinedOutput()
	if len(data) ==0{
		return -1,errors.New("not find")
	}

	info:=string(data)
	// 这里通过正则把进程id提取出来
	reg:=regexp.MustCompile(`[0-9]]+`)
	pid:=reg.FindString(info)

	return strconv.Atoi(pid)
}

func StartProcess(exePath string,args []string) error {
	attr:=&os.ProcAttr{
		// files指定新进程继承的活动文件对象
		// 前三个分别为，标准输入，标准输出，标准错误输出
		Files:[]*os.File{os.Stdin,os.Stdout,os.Stderr},
		// 新进程的环境变量
		Env:os.Environ(),
	}

	p,err:=os.StartProcess(exePath,args,attr)
	if err!=nil{
		return err
	}
	fmt.Println(exePath,"进程启动")
	p.Wait()
	return nil
}
