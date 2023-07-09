package builder

import (
	"fmt"
)

type Product struct{}

type IBuilder interface {
	BuildPart1()
	BuildPart2()
	BuildPart3()
	GetResult() *Product
}

type ConcreteBuilder1 struct{}

func NewConcreteBuilder1() IBuilder {
	return &ConcreteBuilder1{}
}

type ConcreteProduct1 struct{}

func (builder *ConcreteBuilder1) BuildPart1() {
	fmt.Println("building part1")
}

func (builder *ConcreteBuilder1) BuildPart2() {
	fmt.Println("building part2")
}

func (builder *ConcreteBuilder1) BuildPart3() {
	fmt.Println("building part3")
}

func (builder *ConcreteBuilder1) GetResult() *Product {
	return &Product{}
}
