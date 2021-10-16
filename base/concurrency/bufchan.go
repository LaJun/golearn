package main

import (
	"fmt"
	"strconv"
	"time"
)

type Student struct {
	Name string
	Age int
}

func main() {
	bufchanl1()
}

func bufchanl1(){
	ch := make(chan int)
	go func() {
		i := <- ch
		fmt.Println(i)
	}()
	ch <- 100
	time.Sleep(1 * time.Second)
	close(ch)

	ch1 := make(chan int, 2)
	fmt.Println(len(ch1), cap(ch1))
	ch1 <- 100
	fmt.Println(len(ch1), cap(ch1))
	ch1 <- 100                   //往里面写数据的时候，写多了就deadlock
	ch3 := make(chan string, 4)
	go func() {
		for i := 0; i < 10; i++ {
			ch3 <- "数据是" + strconv.Itoa(i)
			fmt.Println("子协程写入第", i, "个数据")
		}
		close(ch3)
	}()

	for  {
		time.Sleep(1 * time.Second)
		v, ok := <-ch3
		if !ok {
			fmt.Println("读完了....",ok)
			break
		}
		fmt.Println("\t读取的数据是：", v)
	}

	fmt.Println("main over.....")
}