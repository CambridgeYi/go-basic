/*
@Time : 2019/6/4 17:38 
@Author : dong.liu
@File : log
@Software: GoLand
@Description:
*/
package log

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

func ConfigFileLogger(logPath string,logFileName string,maxAge time.Duration,rotationTime time.Duration){
	baseLogPath:= path.Join(logPath,logFileName)

	writer,err:=rotatelogs.New(
		baseLogPath+".%Y%m%d%M",
		rotatelogs.WithLinkName(baseLogPath), //生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),//文件最大保存时间
		//rotatelogs.WithRotationCount(365),//最多存365个文件
		rotatelogs.WithRotationTime(rotationTime), //日志切割时间间隔
	)

	if err!=nil{
		fmt.Println("config local file system logger error:%s",err)
	}

	lfHook :=lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel:writer,//为不同级别设置不同的输出目的
		logrus.InfoLevel:writer,
		logrus.WarnLevel:writer,
		logrus.ErrorLevel:writer,
		logrus.FatalLevel:writer,
		logrus.PanicLevel:writer,
	}, &logrus.TextFormatter{})

	logrus.AddHook(lfHook)
}