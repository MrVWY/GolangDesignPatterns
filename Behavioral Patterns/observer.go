package main

import (
	"fmt"
	"time"
)

type Event struct {
	data string
}

type Observer interface {
	Notice(*Event)
}

type Observed interface {
	Regist(Observer)
	Delete(Observer)
	Notify(*Event)
}

type A_Observer struct {
	Id int
}

func (A *A_Observer) Notice(e *Event) {
	fmt.Printf("Observer[%d] receieved '%s' msg ",A.Id,e.data)
}

type B_Observed struct {
	Observers map[Observer]struct{}
}

func (B *B_Observed) Regist(observer Observer) {
	B.Observers[observer] = struct{}{}
}

func (B *B_Observed) Delete(observer Observer) {
	delete(B.Observers,observer)
}

func (B *B_Observed) Notify(e *Event) {
	for O, _ := range B.Observers {
		O.Notice(e)
	}
}

func main() {
	Observed := &B_Observed{Observers: make(map[Observer]struct{})}

	Observer_one := &A_Observer{Id:1}
	Observer_two := &A_Observer{Id:2}

	Observed.Regist(Observer_one)
	Observed.Regist(Observer_two)

	for i:=0 ; i<5 ; i++ {
		e := &Event{fmt.Sprintf("%d\n",i)}
		Observed.Notify(e)

		time.Sleep(50)
	}
}