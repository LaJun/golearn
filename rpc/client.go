package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// 传的参数
type Param struct {
	Width, Height int
}

// 主函数
func main() {
	// 1.连接远程rpc服务
	conn, err := jsonrpc.Dial("tcp", ":8800")
	if err != nil {
		log.Fatal(err)
	}
	// 2.调用方法
	// 面积
	ret := 0
	err2 := conn.Call("Rect.Area", Param{50, 100}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("面积：", ret)
}