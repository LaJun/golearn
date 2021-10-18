package factory_method

import "testing"

func compute(factory OperatorFactory, a,b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	//var factory OperatorFactory
	if compute(PlusOperatorFactory{}, 1, 2) != 3{
		t.Fatal("error in plus")
	}

	if compute(MinusOperatorFactory{}, 4, 2) != 2 {
		t.Fatal("error in minus")
	}
}