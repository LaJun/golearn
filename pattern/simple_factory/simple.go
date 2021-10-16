package simple_factory

import "fmt"

type Factory interface {
	Say(name string)
}
type StudentFactory struct {
	
}
type TeacherFactory struct {
	
}
func (*StudentFactory) Say(ctx string)  {
	fmt.Println(ctx + "i am student")
}
func (*TeacherFactory) Say(ctx string)  {
	fmt.Println( ctx + "i am teacher")
}

func FactoryCreate(t int) Factory {
	if t == 1 {
		return &StudentFactory{}
	} else if t == 2 {
		return &TeacherFactory{}
	}
	return nil
}