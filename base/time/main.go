package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	s1 := time.Now()
	fmt.Println(s1.Year(), s1.Month().String(), s1.Day(), s1.Hour(), s1.Minute())
	//t4 := time.Date(1970,1,1,1,0,0,0,time.UTC)
	fmt.Println(time.Now().Add(10000 * time.Second).Unix())
	res := time.Unix(time.Now().Add(1000 * time.Second).Unix(), 0).Format("12-01/02 15/04/1")
	fmt.Println(res, 1231)
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10) + 1
	time.Sleep(time.Duration(randNum) * time.Second)
	fmt.Println(randNum)
}
