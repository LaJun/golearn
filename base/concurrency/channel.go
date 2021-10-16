package main

import (
	"fmt"
	"time"
)

func test1(ch chan int) {
	fmt.Printf("%T, %p\n",ch, ch)
}
func main()  {
	var a chan int
	if a == nil {
		/*fmt.Println("chnnel 是 nil得不能使用")
		a = make(chan int)
		fmt.Printf("数据类型是： %T, %p", a, a)
		test1(a) //引用传值*/
	}

	ch1 := make(chan bool)
	fmt.Println(ch1)
	//channel1(ch1)
	//channel2()
	channel4()
}

func channel1(ch1 chan bool){
	go func() {
		for i := 0; i< 10; i++ {
			fmt.Println("子协程 中i", i)
		}
		ch1 <- true
		fmt.Printf("结束")
	}()

	data := <-ch1
	fmt.Println("data-->", data)
	fmt.Println("main over ....")
}

func channel2()  {
	ch1 := make(chan  int)
	go func() {
		fmt.Println("子协程执行")
		time.Sleep(3 * time.Second)
		data := <-ch1
		fmt.Println("data: ", data)
	}()

	time.Sleep(10 * time.Second)
	ch1 <- 100
	fmt.Println("main over ....")
}

func channel3(){
	ch1 := make(chan int)
	go func() {
		for i:=0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1) //关闭通道
	}()
	for  {
		time.Sleep(1 * time.Second)
		v, ok := <-ch1    //通道里没数据了， 则在读取int通道的值为0，ok为false， 然后一直堵塞
		if !ok{
			fmt.Println("已经读取了所有数据", v, ok)
			break
		}
		fmt.Println("取出数据：", v, ok)
	}
	fmt.Println("main ..over...")
}

func channel4(){
	ch1 := make(chan int)
	go func() {
		defer close(ch1)
		for i:=0 ; i < 10; i++{
			time.Sleep(1 * time.Second)
			ch1 <-i
		}
	}()

	//range可以读数据知道通道关闭
	for v := range ch1 {
		fmt.Println("读取数据：", v)
	}
	fmt.Println("main ... over....")
}