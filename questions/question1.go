/*
@Time : 2019/6/4 17:49 
@Author : dong.liu
@File : question1
@Software: GoLand
@Description:
*/
package questions

import "fmt"

//解答：
//defer 是后进先出。
//panic 需要等defer 结束后才会向上传递。 出现panic恐慌时候，会先按照defer的后入先出的顺序执行，最后才会执行panic
func Defer_call(){
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println("捕获到panic异常，我要recover过来了")
		}
	}()
	defer func(){
		fmt.Println("打印前")
	}()
	defer func(){
		fmt.Println("打印中")
	}()
	defer func(){
		fmt.Println("打印后")
	}()
	panic("触发异常")
}