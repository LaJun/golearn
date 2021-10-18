package adapter

type Target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest()string
}

type adapteeImpl struct {

}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

func (*adapteeImpl)SpecificRequest() string{
	return "adaptee method"
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

type adapter struct {
	Adaptee
}

func (a *adapter) Request() string{
	return a.SpecificRequest()
}