package observer

import "testing"

func TestObserver(t *testing.T) {
	subject := NewSubject()
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")

	subject.Attach(reader1)
	subject.Attach(reader2)

	subject.UpdateContext("observer mode")
}
