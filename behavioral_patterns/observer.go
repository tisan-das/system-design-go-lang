package behavioral_patterns

import "fmt"

type Subject interface {
	Attach(IObserver)
	Detach(IObserver)
	Notify()
}

type ConcreteSubject struct {
	observerList []IObserver
}

func (subject *ConcreteSubject) Attach(observer IObserver) {
	subject.observerList = append(subject.observerList, observer)
}

func (subject *ConcreteSubject) Detach(observer IObserver) {
	for i, obs := range subject.observerList {
		if observer == obs {
			subject.observerList = append(subject.observerList[:i], subject.observerList[i+1:]...)
		}
	}
}

func (subject *ConcreteSubject) Notify() {
	for _, obs := range subject.observerList {
		obs.Update()
	}
}

type IObserver interface {
	Update()
}

type ConcreteObserver struct {
	Id int
}

func (obs *ConcreteObserver) Update() {
	fmt.Printf("Received update: %v\n", obs)
}
