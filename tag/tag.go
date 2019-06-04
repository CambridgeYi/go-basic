/*
@Time : 2019/6/4 17:52 
@Author : dong.liu
@File : tag
@Software: GoLand
@Description:
*/
package tag

import (
	"fmt"
	"reflect"
)

type T struct{
	f string `one:"1" two:"2" three:""`
}

func GetAllTagValues (){
	n:=reflect.TypeOf(T{})
	// 获取字段：f
	t,_ :=n.FieldByName("f")

	// 获取整个结构体标签内容
	fmt.Println(t.Tag)
	// 获取第1个键值对内容的值，以及是否设置了值
	fmt.Println(t.Tag.Lookup("one"))
	// 获取第2个键值对内容的值，以及是否设置了值
	fmt.Println(t.Tag.Lookup("two"))
	// 获取第3个键值对内容的值，以及是否设置了值
	fmt.Println(t.Tag.Lookup("three"))
	// 获取不存在的键的值，以及是否设置了值
	fmt.Println(t.Tag.Lookup("xxxx"))

}
