package main

import (
	"fmt"
	"reflect"
)

func main()  {
	//var a = [5]int{}
	//fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a).Kind())
	var x float64 = 3
	t := reflect.TypeOf(x)
	fmt.Println("type:", t)
	v := reflect.ValueOf(x)
	fmt.Println("value", v)

	fmt.Println("kind is float64", v.Kind())
	fmt.Println("type: ", v.Type())
	fmt.Println("value: ", v.Float())

	fmt.Println(v.Type() == t, reflect.New(t)  == v)
	fmt.Println(v.Float() == 3.0)
}
