/*
@Time : 2019/6/7 16:21 
@Author : dong.liu
@File : main
@Software: GoLand
@Description:
*/
package main

import (
	"fmt"
	"go-basic/template"
	"html/template"
	"os"
)

func main(){
	// 创建一个模板
	t := template.New("test")
	// 注释
	t,_=t.Parse(`{{/*我是注释*/}}`)
	t.Execute(os.Stdout,nil)

	// 输出单个字符串
	t2 :=template.New("test")
	// {{.}}输出当前对象的值
	t2,_ =t2.Parse(`{{.}}`)
	p:="test"
	// 输出字符串
	t2.Execute(os.Stdout,nil)
	fmt.Println();

	// 输出结构中字段的值
	t3 := template.New("test")
	// {{.字段名}}输出对象中字段的值
	// 注意字段是可导出的，首字符大写
	t3,_ =t3.Parse(`{{.Id}}{{.UserName}}{{.Age}}`)
	t3.Execute(os.Stdout,mytemplate.Person{"001","test",11,nil})
	fmt.Println()

	// 调用结构的方法
	t4:=template.New("test")
	// {{.方法 参数1 参数2}}
	// 参数依次传入方法，输出返回值
	t4,_ =t4.Parse(`{{.Say "hello"}}`)
	t4.Execute(os.Stdout,mytemplate.Person{"002","test2",22,nil})
	fmt.Println()

	// 模板中定义变量
	t5:=template.New("test")
	// {{$变量名}}输出模板中定义的变量
	t5,_ = t5.Parse(`{{$a:="模板中定义的变量"}}{{&a}}`)
	t5.Execute(os.Stdout,nil)
	fmt.Println()

	// 模板函数
	t6:=template.New("test")
	// 注册模板函数
	t6.Funcs(template.FuncMap{"test1":mytemplate.Ttest1})
	t6.Funcs(template.FuncMap{"test2":mytemplate.Ttest2})

	// {{函数名}}输出函数返回值
	// {{函数名 参数1 参数2}}
	// {{.字符名|函数名}} 以字段的值作为函数的参数
	t6,_=t6.Parse(`
		{{mytemplate.Ttest1}}
		{{mytemplate.Ttest2 "参数"}}
 		{{.UserName|mytemplate.Ttest2}}
	`)
	t6.Execute(os.Stdout,mytemplate.Person{"003","test3",33,nil})
	fmt.Println()

	//条件判断
	t7 := template.New("test")
	t7.Funcs(template.FuncMap{"test3": mytemplate.Ttest3})
	// {{if 表达式}}{{else if}}{{else}}{{end}}
	// if后面可以是一个条件表达式，可以是字符串或布尔值变量
	// 注意if后面不能直接使用==来判断
	t7, _ = t7.Parse(`
                      {{if 1}}
                        为真
                      {{else}}
                        为假
                      {{end}}
 
                      {{$a := 4}}
                      {{if $a|test3}}
                        $a=3
                      {{else}}
                        $a!=3
                      {{end}}
                      `);
	t7.Execute(os.Stdout, nil)
	fmt.Println()

	//遍历
	t8 := template.New("test")
	// {{range 键,值 := 变量}}{{end}} 遍历对象
	// {{with 变量}}{{end}} 指定当前操作的对象
	t8, _ = t8.Parse(`
                      {{range $k, $v := .Contact}}
                        {{$k}} {{$v}}
                      {{end}}
 
                      {{with .Contact}}
                        {{range $k, $v := .}}
                            {{$k}} {{$v}}
                        {{end}}
                      {{end}}
                      `)
	con := make(map[string]string)
	con["qq"] = "123456"
	con["tel"] = "13888888888"
	t8.Execute(os.Stdout, mytemplate.Person{Contact: con});
	fmt.Println()

	//嵌套模板
	t9 := template.New("test")
	t9.Funcs(template.FuncMap{"test1": mytemplate.Ttest1})
	// {{define "模板名"}}模板内容{{end}} 定义模板
	// {{template "模板名"}} 引入模板
	// {{template "模板名" 函数}} 将函数中的值赋给模板中的{{.}}
	t9, _ = t9.Parse(`
                      {{define "tp1"}} 我是模板1 {{end}}
                      {{define "tp2"}} 我是模板2 {{.}} {{end}}
                      {{define "tp3"}} {{template "tp1"}} {{template "tp2"}} {{end}}
                      {{template "tp1"}}
                      {{template "tp2" test1}}
                      {{template "tp3" test1}}
                     `)
	t9.Execute(os.Stdout, nil)
	fmt.Println()

	//内置的模板函数
	t10 := template.New("test")
	t10.Funcs(template.FuncMap{"sum": mytemplate.Sum})
	t10, _ = t10.Parse(`
                        /*如果3为真，返回4，否则返回3*/
                        {{and 3 4}}
 
                        /*call后第一个参数的返回值必须是一个函数*/
                        {{call sum 1 3 5 7}}
 
                        /*转义文本中的html标签*/
                        {{"<br>"|html}}
 
                        /*返回Contact索引为qq的值*/
                        {{index .Contact "qq"}}
 
                        /*返回用js的escape处理后的文本*/
                        {{"?a=123&b=你好"|js}}
 
                        /*返回参数的长度值*/
                        {{"hello"|len}}
 
                        /*返回单一参数的布尔否定值*/
                        {{not 0}}
 
                        /*如果3为真，返回3，否则返回4*/
                        {{or 3 4}}
 
                        /*fmt.Sprint的别名*/
                        {{"你好"|print "世界"}}
 
                        /*fmt.Sprintf的别名*
                        {{"你好"|printf "%d %s" 123}}
 
                        /*fmt.Sprintln的别名*/
                        {{"你好"|println "世界"}}
 
                        /*url中get参数转义*/
                        {{"?q=关键字&p=1"|urlquery}}
 
                        /*等于*/
                        {{if eq 1 1}}
                        1=1
                        {{end}}
 
                        /*不等于*/
                        {{if ne 1 1}}
                        1!=1
                        {{end}}
 
                        /*小于*/
                        {{if lt 3 1}}
                        3<1
                        {{end}}
 
                        /*小于等于*/
                        {{if le 3 3}}
                        3<=3
                        {{end}}
 
                        /*大于*/
                        {{if gt 3 1}}
                        3>1
                        {{end}}
 
                        /*大于等于*/
                        {{if ge 3 3}}
                        3>=3
                        {{end}}
                       `);
	con2 := make(map[string]string)
	con2["qq"] = "123456"
	con2["tel"] = "13888888888"
	t10.Execute(os.Stdout, mytemplate.Person{Contact: con2})
}

