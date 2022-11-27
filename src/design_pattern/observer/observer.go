package observer

import (
	"fmt"
	"strconv"
)

func NewSubject() *Subject {
	return &Subject{
		observers: make([]*IObserver, 0),
		state:     0,
	}
}

type Subject struct {
	observers []*IObserver
	state     int
}

func (s *Subject) getState() int {
	return s.state
}

func (s *Subject) setState(state int) {
	s.state = state
	s.notifyAllObservers()
}

func (s *Subject) attach(observer IObserver) {
	s.observers = append(s.observers, &observer)
}
func (s *Subject) notifyAllObservers() {
	for _, o := range s.observers {
		(*o).update()
	}
}

type IObserver interface {
	update()
}

type BaseObserver struct {
	subject *Subject
}

func (o *BaseObserver) update() {

}

type BinaryObserver struct {
	BaseObserver
}

func NewBinaryObserver(subject *Subject) *BinaryObserver {
	observer := BinaryObserver{
		BaseObserver{
			subject: subject,
		},
	}
	observer.subject.attach(&observer)
	return &observer
}

func (b *BinaryObserver) update() {
	fmt.Println("Binary String: " + strconv.Itoa(b.subject.state))
}

func NewOctalObserver(subject *Subject) *OctalObserver {
	observer := OctalObserver{
		BaseObserver{
			subject: subject,
		},
	}
	observer.subject.attach(&observer)
	return &observer
}

type OctalObserver struct {
	BaseObserver
}

func (o *OctalObserver) update() {
	fmt.Println("Octal String: " + strconv.Itoa(o.subject.getState()))
}

func NewHexaObserver(subject *Subject) *HexaObserver {
	observer := HexaObserver{
		BaseObserver{
			subject: subject,
		},
	}
	observer.subject.attach(&observer)
	return &observer
}

type HexaObserver struct {
	BaseObserver
}

func (h *HexaObserver) update() {
	fmt.Println("Hex String: " + strconv.Itoa(h.subject.getState()))
}
