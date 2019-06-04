/*
@Time : 2019/6/4 17:50 
@Author : dong.liu
@File : question2
@Software: GoLand
@Description:
*/
package questions

import "fmt"

type student struct{
	Name string
	Age int
}
//考点：foreach
//解答：
//这样的写法初学者经常会遇到的，很危险！ 与Java的foreach一样，都是使用副本的方式。所以m[stu.Name]=&stu实际上一致指向同一个指针，
// 最终该指针的值为遍历的最后一个struct的值拷贝
func Pase_Student(){
	m :=make(map[string]*student)

	var stus = []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	} //错误写法
	for _,sku:=range stus{
		m[sku.Name]=&sku
	}
	fmt.Println(m)

	//正确写法
	for i:=0;i<len(stus);i++{
		m[stus[i].Name] =&stus[i]
	}
	fmt.Println(m)
}