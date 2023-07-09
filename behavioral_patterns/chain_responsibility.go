package behavioral_patterns

import (
	"fmt"
)

type OrderHandler interface {
	SetNext(OrderHandler)
	Execute(*Task)
}

type Task struct {
	isOrderProcessed    bool
	isPackageProcessed  bool
	isBillingProcessed  bool
	isShipmentProcessed bool
}

type OrderProcessor struct {
	nextAction OrderHandler
}

func (proc *OrderProcessor) SetNext(orderHandler OrderHandler) {
	proc.nextAction = orderHandler
}

func (proc *OrderProcessor) Execute(task *Task) {
	if !task.isOrderProcessed {
		fmt.Println("Processing the order!")
		task.isOrderProcessed = true
		proc.nextAction.Execute(task)
	} else {
		fmt.Println("Order is already processed!")
		proc.nextAction.Execute(task)
	}
}

type PackageProcessor struct {
	nextAction OrderHandler
}

func (proc *PackageProcessor) SetNext(orderHandler OrderHandler) {
	proc.nextAction = orderHandler
}

func (proc *PackageProcessor) Execute(task *Task) {
	if !task.isPackageProcessed {
		fmt.Println("Processing the package!")
		task.isPackageProcessed = true
	} else {
		fmt.Println("Package is already processed!")
	}
	if proc.nextAction != nil {
		proc.nextAction.Execute(task)
	}
}

type BillingProcessor struct {
	nextAction OrderHandler
}

func (proc *BillingProcessor) SetNext(orderHandler OrderHandler) {
	proc.nextAction = orderHandler
}

func (proc *BillingProcessor) Execute(task *Task) {
	if !task.isBillingProcessed {
		fmt.Println("Processing the billing!")
		task.isBillingProcessed = true
	} else {
		fmt.Println("Billing is already processed!")
	}
	if proc.nextAction != nil {
		proc.nextAction.Execute(task)
	}
}

type ShipmentProcessor struct {
	nextAction OrderHandler
}

func (proc *ShipmentProcessor) SetNext(orderHandler OrderHandler) {
	proc.nextAction = orderHandler
}

func (proc *ShipmentProcessor) Execute(task *Task) {
	if !task.isShipmentProcessed {
		fmt.Println("Processing the shipment!")
		task.isShipmentProcessed = true
	} else {
		fmt.Println("Shipment is already processed!")
	}
	if proc.nextAction != nil {
		proc.nextAction.Execute(task)
	}
}
