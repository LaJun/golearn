package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main(){
	wg.Add(2)
	go func() {
		for i:=1;i<=10;i++{
			fmt.Println("fun1.。。i:",i)
		}
		wg.Done() //给wg等待中的执行的goroutine数量减1.同Add(-1)  如果没有fatal error: all goroutines are asleep - deadlock!
	}()

	go func() {
		for j:=1;j<=10;j++{
			fmt.Println("\tfun2..j,",j)
			time.Sleep(1* time.Second)
		}
		defer   wg.Done()
	}()

	fmt.Println("main进入阻塞状态。。。等待wg中的子goroutine结束。。")
	wg.Wait()
	fmt.Println("main，解除阻塞。。")

}
