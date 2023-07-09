package behavioral_patterns

type Observer struct {
	PrivateStr string
}

func (obs *Observer) CreateMemento() *ObserverMemento {
	return &ObserverMemento{privateStr: obs.PrivateStr}
}

func (obs *Observer) SetMemento(memento *ObserverMemento) {
	obs.PrivateStr = memento.privateStr
}

// ObserverMemento methods should be accessible from Observer only
type ObserverMemento struct {
	privateStr string
}

func (memento *ObserverMemento) getState() string {
	return memento.privateStr
}

func (memento *ObserverMemento) setState(str string) {
	memento.privateStr = str
}

// Place in different package to make sure observerMemento methods are inaccessible to Caretaker
type Caretaker struct {
	mementoList []*ObserverMemento
}

func (caretaker *Caretaker) AppendMemento(memento *ObserverMemento) {
	caretaker.mementoList = append(caretaker.mementoList, memento)
}

func (caretaker *Caretaker) GetMementoIndex(index int) *ObserverMemento {
	return caretaker.mementoList[index]
}
