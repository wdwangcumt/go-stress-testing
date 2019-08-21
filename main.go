/**
* Created by GoLand.
* User: link1st
* Date: 2019-08-15
* Time: 13:44
 */

package main

import (
	"flag"
	"fmt"
	"go-stress-testing/model"
	"go-stress-testing/server"
	"strings"
)

func main() {
	var (
		concurrency uint64 // 并发数
		totalNumber uint64 // 请求总数(单个并发)
		debugStr    string // 是否是debug
		requestUrl  string // 压测的url 目前支持，http/https ws/wss
		path        string // curl文件路径 http接口压测，自定义参数设置
	)

	flag.Uint64Var(&concurrency, "c", 1, "并发数")
	flag.Uint64Var(&totalNumber, "n", 1, "请求总数")
	flag.StringVar(&debugStr, "d", "false", "调试模式")
	flag.StringVar(&requestUrl, "u", "", "请求地址")
	flag.StringVar(&path, "p", "", "curl文件路径")

	// 解析参数
	flag.Parse()
	if concurrency == 0 || totalNumber == 0 || (requestUrl == "" && path == "") {
		fmt.Printf("示例: go run main.go -c 1 -n 1 -u https://www.baidu.com/ \n")
		fmt.Printf("-c %d -n %d -d %v -u %s \n", concurrency, totalNumber, debugStr, requestUrl)

		flag.Usage()

		return
	}

	debug := strings.ToLower(debugStr) == "true"
	request, err := model.NewRequest(requestUrl, "", 0, debug, path)
	if err != nil {
		fmt.Printf("参数不合法 %v \n", err)

		return
	}

	fmt.Printf("\n 开始启动  并发数:%d 请求数:%d 请求参数: \n", concurrency, totalNumber)
	request.Print()

	// 开始处理
	server.Dispose(concurrency, totalNumber, request)

	return
}
