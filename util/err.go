/*
@Time : 2019/6/4 17:35 
@Author : dong.liu
@File : err
@Software: GoLand
@Description:
*/
package util

import (
	"fmt"
	"os"
)

func CheckError(err error){
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal errors:%s",err.Error())
		os.Exit(1)
	}
}