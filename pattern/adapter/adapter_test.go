package adapter

import "testing"

func TestAdapter(t *testing.T)  {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	expect :=  "adaptee method"
	if res != expect{
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}

}