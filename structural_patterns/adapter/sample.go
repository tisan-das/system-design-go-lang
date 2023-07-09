package adapter

import (
	"fmt"
)

// -- Adaptee ---
type Adaptee interface {
	SpecificRequest()
}

func NewAdaptee() Adaptee {
	return &adaptee{}
}

type adaptee struct{}

func (adptr *adaptee) SpecificRequest() {}

// -- Target ---
type Target interface {
	Request()
}

func NewTarget() Target {
	return &target{}
}

type target struct{}

func (tgt *target) Request() {}

// -- Adapter ---

func NewAdapter() Target {
	return &adapter{adaptee: NewAdaptee()}
}

type adapter struct {
	adaptee Adaptee
}

func (adpt *adapter) Request() {
	fmt.Println("Using the requesst specific method for adaptee!")
	adpt.adaptee.SpecificRequest()
}
