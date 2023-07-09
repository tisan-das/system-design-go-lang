package bridge

import (
	"fmt"
)

type Abstraction interface {
	Operation()
}

type Implementor interface {
	OperationImp()
}

func NewRefinedAbstraction(imp Implementor) Abstraction {
	return &refinedAbstraction{imp: imp}
}

type refinedAbstraction struct {
	imp Implementor
}

func (abs *refinedAbstraction) Operation() {
	abs.imp.OperationImp()
}

func NewConcreteImplementorA() Implementor {
	return &ConcreteImplementorA{}
}

type ConcreteImplementorA struct{}

func (impl *ConcreteImplementorA) OperationImp() {
	fmt.Println("Inside operationImp of concrete implementor A")
}
