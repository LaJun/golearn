package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
	Sex string
}

type Man struct {
	Person
	JB float64
}
func (p Person)Say(msg string){
	fmt.Println("hello,", msg)
}


func (m Man) Say(msg string, nam int) {
	m.Person.Say("牛逼")
}

func (p Person)PrintInfo() {
	fmt.Printf("姓名：%s, 年龄：%，性别： %s\n", p.Name, p.Age, p.Sex)
}


func main()  {
	man := Man{
		Person{"Lei", 12, "nan"},
		12,
	}
	man.Person.Say("sadas")

	var num float64 = 1.234

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)
	fmt.Println(pointer, value, convertPointer, convertValue)

	/*p1 := Person{"王二狗", 30, "男"}
	DofiledAndMethod(p1)*/
	//ChangeVar()
	//ChangeAtr()
	callMethod()
	//callFunc()
}

func DofiledAndMethod(input interface{}){
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is", getType.Name(), ",get Kind is:", getType.Kind() )

	getValue := reflect.ValueOf(input)
	fmt.Println("get all field is:", getValue)

	for i:= 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("字段名称:%s, 字段类型:%s, 字段数值:%v \n", field.Name, field.Type, value)
	}

	for i:= 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名称:%s, 方法类型:%v \n", method.Name, method.Type)
	}
}

func ChangeVar()  {
	var num float64 = 1.2342
	fmt.Println("num的数值为", num)

	pointer := reflect.ValueOf(&num) //这里必须是指针参数，否则pointer不是指针，下面Elem()报错
	newValue := pointer.Elem()
	fmt.Println("类型是", newValue.Type(), "是否可以改变", newValue.CanSet())

	/*value := reflect.ValueOf(&num)
	value.SetFloat(6.28) //panic: reflect: reflect.Value.SetFloat using unaddressable value
	fmt.Println(value.CanSet()) //false*/

	newValue.SetFloat(77.12312)
	fmt.Println("新的数值：", num)


}

func ChangeAtr() {
	s1 := Person{"雷车", 12, "男"}
	fmt.Printf("%T\n",s1)  //main.Person类型
	p1 := &s1
	fmt.Printf("%T\n",p1)  //*main.Person类型
	fmt.Println(s1.Name, (*p1).Name, p1.Name)

	v1 := reflect.ValueOf(&s1)

	if v1.Kind() == reflect.Ptr{
		fmt.Println(v1.Elem().CanSet())
		v1 = v1.Elem()
	}

	f1 := v1.FieldByName("Name")
	f1.SetString("及澳门")
	f3 := v1.FieldByName("Sex")
	f3.SetString("幼儿园")
	fmt.Println(s1)
}


func callMethod() {
	p := Person{"ruby", 30, "nan"}
	getValue := reflect.ValueOf(p)

	methodValue1 := getValue.MethodByName("PrintInfo")

	fmt.Printf("Kind : %s, Type : %s\n",methodValue1.Kind(),methodValue1.Type())
	methodValue1.Call(nil) //没有参数可以传nil
	args := make([]reflect.Value, 0) //没有参数，可以传空切片
	methodValue1.Call(args)

	methodValue2 := getValue.MethodByName("Say")
	methodValue2.Call([]reflect.Value{reflect.ValueOf("我很牛逼，你信吗") }) //传递reflect.Value切片
}

func callFunc() {
	value := reflect.ValueOf(foo1)
	fmt.Printf("Kind : %s , Type : %s\n",value.Kind(),value.Type()) //Kind : func , Type : func()

	value2 := reflect.ValueOf(foo2)
	fmt.Printf("Kind : %s , Type : %s\n",value2.Kind(),value2.Type()) //Kind : func , Type : func(int, string)

	value.Call(nil)
	value2.Call([]reflect.Value{reflect.ValueOf(11), reflect.ValueOf("213123")}) //参数个数和类型必须一致
}


func foo1()  {
	fmt.Println("我是函数fun1()，无参的。。")
}

func foo2(i int, s string)  {
	fmt.Println("我是函数fun2()，有参数。。",i,s)
}