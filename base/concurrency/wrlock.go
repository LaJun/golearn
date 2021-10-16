package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex

var wg *sync.WaitGroup

func main()  {
	rwMutex = new(sync.RWMutex)
	wg = &sync.WaitGroup{}
	wg.Add(2)
	go readData(1)
	go readData(2)
	go writeData(3)
	wg.Wait()
	fmt.Println("主程序结束")
}

func writeData(i int){
	defer wg.Done()
	fmt.Println(i,"开始写：write start。。")
	rwMutex.Lock()
	fmt.Println(i,"正在写：writing。。。。")
	time.Sleep(3*time.Second)
	rwMutex.Unlock()
	fmt.Println(i,"写结束：write over。。")
}

func readData(i int){
	defer wg.Done()
	fmt.Println(i, "开始读：read start。。")
	rwMutex.RLock() //读操作上锁
	fmt.Println(i,"正在读取数据：reading。。。")
	time.Sleep(3*time.Second)
	rwMutex.RUnlock() //读操作解锁
	fmt.Println(i,"读结束：read over。。。")
}
