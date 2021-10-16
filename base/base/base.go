package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
)

func main()  {
	num := 2131233333333333333
	var  num2 int32 = 123
	fmt.Println(reflect.TypeOf(num),reflect.TypeOf(num2).Name(), reflect.TypeOf(num2))
	fmt.Printf("%T \t", num)
	fmt.Printf("%T\t", num2)

	float12 := 1231.31230
	var float43 float32 = 2312.7213
	fmt.Println(reflect.TypeOf(float12), reflect.TypeOf(float43))

	float43int := int(float43)  //浮点转整形，去掉小数点
	fmt.Println(float43int)
	fmt.Println(float64(num), float64(num2)) //整行转浮点,小的就是直接返回，大的变成科学计数法

	//s := "hello world"
	//fmt.Println(string(num), int(s))
	//fmt.Println(float64(s)) 字符串和数字类型不能乱转,  整型能转成乱码

	//数字和字符串互相转换需要库
	s := "0101"
	res, _ := strconv.Atoi(s)  //必须是完整的数字字符串, 0101转成101，0.111转成0.111
	fmt.Println(res, reflect.TypeOf(res))

	s1 := "0101"
	res1, _ := strconv.ParseFloat(s1, 1)  //必须是完整的数字字符串, 0101转成101，0.111转成0.111
	fmt.Println(res1, reflect.TypeOf(res1))

	s2 := strconv.Itoa(num) //只能是转型
	fmt.Println(s2)

	v := strconv.FormatFloat(float12, 'f', 4, 32)  //f代表是十进制不科学计数法， 5代表小数的位数，大于本来的位数后补0，小于的话直接截断

	fmt.Println("\t", float12, reflect.TypeOf(float12), v, reflect.TypeOf(v))

	//小数上下取整，四舍五入
	f2 := 1.2
	//向下取整，向上取整， 四舍五入
	fmt.Println(math.Floor(f2), math.Ceil(f2), math.Round(1.6), math.Round(1.3))

	//随机数

	min := 10
	max := 20
	fmt.Println(rand.Int63n(int64(max)- int64(min)) + int64(min))
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
