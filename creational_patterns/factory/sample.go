package factory

import "fmt"

type IProduct interface {
	GetProductName() string
}

type Product1 struct {
	ProductName1 string
}

func (product *Product1) GetProductName() string {
	return "product1"
}

type Product2 struct {
	ProductName2 string
}

func (product *Product2) GetProductName() string {
	return "product2"
}

type ICreator interface {
	FactoryMethod() IProduct
	Operation()
	GetProduct() IProduct
}

type ConcreteCreator1 struct {
	Product1 Product1
}

func NewConcreteCreator1() ICreator {
	return &ConcreteCreator1{Product1: Product1{}}
}

func (creator *ConcreteCreator1) FactoryMethod() IProduct {
	fmt.Println("going to create the object")
	return &creator.Product1
}

func (creator *ConcreteCreator1) Operation() {
	fmt.Println("operation: ", creator.Product1.GetProductName())
}

func (creator *ConcreteCreator1) GetProduct() IProduct {
	return &creator.Product1
}

type ConcreteCreator2 struct {
	Product2 Product2
}

func NewConcreteCreator2() ICreator {
	return &ConcreteCreator2{Product2: Product2{}}
}

func (creator *ConcreteCreator2) FactoryMethod() IProduct {
	fmt.Println("going to create the object")
	return &creator.Product2
}

func (creator *ConcreteCreator2) Operation() {
	fmt.Println("operation: ", creator.Product2.GetProductName())
}

func (creator *ConcreteCreator2) GetProduct() IProduct {
	return &creator.Product2
}
