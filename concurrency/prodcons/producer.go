/*
@Time : 2019/6/4 17:19 
@Author : dong.liu
@File : producer.go
@Software: GoLand
@Description:
*/
package prodcons

//生产者：生成factor整数倍的序列
func Producer(factor int, out chan<- int) {
	for i := 0; i < 5; i++ {
		out <- i * factor
	}
}