/*
@Time : 2019/6/4 17:39 
@Author : dong.liu
@File : user
@Software: GoLand
@Description:
*/
package model

// 定义结构体（xorm支持双向映射）
type User struct {
	User_id int64 `xorm:"pk autoincr"`  //指定主键并自增
	Name string `xorm:"unique"`//唯一的
	Balance float64
	Time int64 `xorm:"updated"`//修改后自动更新时间
	Creat_time int64 `xorm:"created"`//创建时间
}