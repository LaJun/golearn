package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket = 10

var mutex sync.Mutex
var wg sync.WaitGroup

func saleTickets(name string){
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println(name,"售出：",ticket) // 1
			ticket--
			mutex.Unlock()
		} else {
			mutex.Unlock()
			fmt.Println(name,"售罄，没有票了。。")
			break
		}
	}
}

func main()  {
	wg.Add(4)
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")
	wg.Wait()
	fmt.Println("程序结束")
}



