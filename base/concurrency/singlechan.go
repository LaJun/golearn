package main

import (
	"fmt"
	"strconv"
	"time"
)

func main()  {
	num := 1
	go func() {
		num += 1
		fmt.Println("ζεδΊ" + strconv.Itoa(  num))
	}()
	fmt.Println(num)

	chan1 := make(chan int)
	go func(channel <-chan int ) {
	}(chan1)

	time.Sleep(3 * time.Second)
}