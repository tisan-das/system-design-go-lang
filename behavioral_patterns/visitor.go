package behavioral_patterns

import "fmt"

type Shape interface {
	GetType() string
	Accept(Visitor)
}

type Square struct {
	Side int
}

func (square *Square) GetType() string {
	return "square"
}

func (square *Square) Accept(visitor Visitor) {
	visitor.VisitSquare(square)
}

type Circle struct {
	Radius int
}

func (circle *Circle) GetType() string {
	return "circle"
}

func (cirlce *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(cirlce)
}

type Visitor interface {
	VisitSquare(*Square)
	VisitCircle(*Circle)
}

type AreaCalculator struct{}

func (area *AreaCalculator) VisitSquare(square *Square) {
	fmt.Printf("Going to calculate area for square: %+v\n", square)
}

func (area *AreaCalculator) VisitCircle(circle *Circle) {
	fmt.Printf("Going to calculate area for circle: %+v\n", circle)
}
