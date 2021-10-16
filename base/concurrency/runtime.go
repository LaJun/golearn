package main

import (
	"fmt"
	"runtime"
	"time"
)

func init(){
	fmt.Println("GOROOT-->", runtime.GOROOT()) //获取goroot目录
	fmt.Println("os/platform-->", runtime.GOOS) //获取操作系统
	fmt.Println("逻辑CPU核数：", runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
}

func main() {
	go func() {
		for i :=  0; i < 5; i++ {
			fmt.Println("goroutine...")
		}
	}()
	for i := 0; i < 4; i++ {
		runtime.Gosched()
		fmt.Println("main...")
	}

	go func() {
		fmt.Println("goroutine start")
		fun()
		fmt.Println("goroutine end")
	}()
	time.Sleep(3 * time.Second)
}

func fun()  {
	defer fmt.Println("defer....")
	runtime.Goexit()
	fmt.Println("fun函数")
}
