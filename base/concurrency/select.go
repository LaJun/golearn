package main

import (
	"fmt"
	"time"
)

func main() {
	//select1()
	select2()
}

func select1(){
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()

	select {
		case num1 := <-ch1:
			fmt.Println("ch1 中读取数据", num1)
			case num2, ok := <-ch2:
				if ok{
					fmt.Println("ch2中取数据。。", num2)
				}else{
					fmt.Println("ch2通道已经关闭。。")
				}
	}
}

func select2(){
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch2 <- 200
	}()
	select {
	case <-ch1:
		fmt.Println("case 1")
	case <-ch2:
		fmt.Println("case 2..")
	case <-time.After(3 * time.Second):
		fmt.Println("case3执行。。timeout。。")
	}
}