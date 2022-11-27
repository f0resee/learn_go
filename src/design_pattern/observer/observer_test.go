package observer

import "testing"

func TestObserver(t *testing.T) {
	subject := NewSubject()

	NewHexaObserver(subject)
	NewOctalObserver(subject)
	NewBinaryObserver(subject)

	t.Log("set state to 15")
	subject.setState(15)
	t.Log("set state to 10")
	subject.setState(10)
}
