/*
@Time : 2019/6/4 17:53 
@Author : dong.liu
@File : timesignal
@Software: GoLand
@Description:
*/
package mytime

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Task struct{
	// 任务的内容,具体可以是一个复杂的结果体对象
	msg string
	// 任务的优先级,在对同一个bucket的数据,可以按照优先级来处理
	pri int
	// bucket的标识
	idx int
	// 任务标识,标识任务是否执行成功,是否需要删除
	status bool
}

var taskList = map[int][]Task{}

var ticker =time.NewTicker(1*time.Second)
// 轮片指针
var cc =0

// 简单的执行任务
func (t *Task) runTask(){
	fmt.Println("run message",t.msg)
	t.status=true
}

/*
	* 假设 i是任务的id号，表示有一个150个任务要进入队列审核
*/
func initTask(){
	for i:=0;i<150;i++{
		sendTask(i)
	}
}

func sendTask(idx int){
	msg:=fmt.Sprintf("task message %d",idx)
	pri:=idx/60
	idx=idx%60

	task:=Task{
		msg,
		pri,
		idx,
		false,
	}
	taskList[idx] =append(taskList[idx],task)
}

func Run(){
	c:=make(chan os.Signal)
	status:=true
	signal.Notify(c,
		syscall.SIGKILL,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		os.Interrupt,
		os.Kill,
	)

	initTask()

	go func(){
		for{
			select{
			case <- ticker.C:
				for _,t:=range taskList[cc]{
					if t.status==false{
						t.runTask()
					}
				}
				cc +=1
				// 循环轮询
				cc = cc%60
			case <-c: // 监听 信息
				ticker.Stop()
				fmt.Println("kill task")
				status =false
				break
			}
		}
	}()
	// 常驻
	for{
		time.Sleep(1*time.Second)
		if status == false {
			break
		}
	}
}