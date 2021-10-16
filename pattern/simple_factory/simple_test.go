package simple_factory

import "testing"

func TestFactory(t *testing.T) {
	factory := FactoryCreate(1)
	factory.Say("student a")
}