package facade

import "fmt"

type API interface {
	Test() string
}

type apiTmpl struct {
	a AModuleAPI
	b BModuleAPI
}
type AModuleAPI interface {
	TestA() string
}
type aModuleImpl struct{}

//NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

func (*aModuleImpl) TestA() string{
	return "A module running"
}

func NewAPI() API {
	return &apiTmpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}
func (a *apiTmpl) Test() string{
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

type BModuleAPI interface {
	TestB() string
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleTmpl{}
}

type bModuleTmpl struct {

}

func (* bModuleTmpl) TestB() string{
	return "B module running"
}
