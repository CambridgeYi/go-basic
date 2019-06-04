/*
@Time : 2019/6/4 17:26 
@Author : dong.liu
@File : type
@Software: GoLand
@Description:
*/
package crawler_model

// 解析结果
type ParseResult struct{
	// 解析出来的多个 Request 任务
	Requests []Request
	//  解析出来的实体（例如，城市名），是任意类别
	Items []interface{}
}

// 请求任务封装体
type Request struct{
	// 需要爬取的Url
	Url string
	// Url对应的解析函数
	ParseFunc func([]byte) ParseResult
}