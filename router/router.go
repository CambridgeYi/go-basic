/*
@Time : 2019/6/4 17:40 
@Author : dong.liu
@File : router
@Software: GoLand
@Description:
*/
package router

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"strconv"
	"fmt"
	"go-basic/model"
)

//定义xorm引擎
var x *xorm.Engine

func init(){
	var err error
	//创建xorm引擎
	x,err = xorm.NewEngine("mysql","root:123456@tcp(10.10.22.35:3306)/mygo?charset=utf8")

	if err!=nil {
		fmt.Println("数据库连接失败,异常信息：%s",err)
	}

	if err:=x.Sync(new(model.User));err!=nil{
		fmt.Println("数据库表同步失败,异常 信息：%s",err)
	}
	fmt.Println("数据库连接成功")
}

// @增加用户
// @Summary add user
// @Accept  json
// @Produce json
// @Param name query string false "string valid"
// @Param money query int false "string valid"
// @Success 200 {string} string	"ok"
// @Router /insert [get]
func Insert(c *gin.Context){
	logrus.Info("开始添加")

	name:=c.Query("name")
	if name==""{
		c.JSON(200,gin.H{"msg":"name不得为空"})
		return
	}
	money := c.Query("money")
	if money==""{
		c.JSON(200,gin.H{"msg":"money不得为空"})
	}
	balance,_:=strconv.ParseFloat(money,64)
	var user = new(model.User)
	user.Name =name
	user.Balance=balance
	affected,err := x.Insert(user)
	defer x.Close()
	if affected==0||err!=nil{
		c.JSON(200,gin.H{"msg":"添加失败","err":err,"rel":affected})
	}else{
		c.JSON(200,gin.H{"msg":"添加成功"})
	}

}