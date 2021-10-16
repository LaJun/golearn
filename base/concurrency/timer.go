package main

import (
	"fmt"
	"sync"
	"time"
)
var wg = sync.WaitGroup{}
func main()  {

	//timer4()

	wg.Add(1)
	time.AfterFunc(3 * time.Second, func() {
		fmt.Println("3s后结束")
		wg.Done()
	})
	fmt.Println("main start ....")
	wg.Wait()
	fmt.Println("main end ....")

	//timer3()
}

func timer4()  {
	ticker := time.NewTicker(1 * time.Second)
	wg.Add(1)
	go func() {
		i := 0
		for {
			<-ticker.C
			time.Sleep(time.Duration(2 * i)* time.Second)
			i++
			fmt.Println("重试第 ", i, "次")
			if i == 5 {
				ticker.Stop()
				break
			}
		}
		wg.Done()
	}()
	fmt.Println("main running")
	wg.Wait()
	fmt.Println("main end....")
}

func timer2(){
	timer := time.NewTimer(8 * time.Second)
	go func() {
		fmt.Println(" timer start....")
		<-timer.C
	}()

	time.Sleep(2 * time.Second)
	stop := timer.Stop()
	if stop{
		fmt.Println("timer stop")
	}

	//time.Sleep(10 * time.Second)
	fmt.Println("main over ....")
}

func timer3(){
	ch1 := time.After(3 * time.Second)
	fmt.Printf("%T\n", ch1)
	fmt.Println(time.Now())
	time2 := <-ch1
	fmt.Println(time2)

}

func timer1(){
	timer := time.NewTimer(3 * time.Second)
	fmt.Printf("%T \n", timer)
	fmt.Println(time.Now())
	ch2 := timer.C
	fmt.Println(<-ch2)
}