/*
@Time : 2019/6/7 16:17 
@Author : dong.liu
@File : template
@Software: GoLand
@Description:
*/
package mytemplate

type Person struct{
	Id string
	UserName string
	Age int
	Contact map[string]string
}

func (p Person) Say(msg string) string{
	return msg
}

func Ttest1() string{
	return "test1"
}

func Ttest2(msg string) string{
	return msg+"test2"
}

func Ttest3(a int) bool{
	if a==3{
		return true
	}
	return false
}

func Sum() func(nums ...int)(int,error){
	return func(nums ...int)(int,error){
		sum :=0
		for _,v :=range nums{
			sum +=v
		}
		return sum,nil
	}
}