package behavioral_patterns

import "fmt"

type Expression interface {
	Evaluate() float64
	Interpret() string
}

type Number struct {
	Value float64
}

func (num *Number) Evaluate() float64 {
	return num.Value
}

func (num *Number) Interpret() string {
	return fmt.Sprintf("%f", num.Value)
}

type Plus struct {
	Expr1, Expr2 Expression
}

func (plus *Plus) Evaluate() float64 {
	return plus.Expr1.Evaluate() + plus.Expr2.Evaluate()
}

func (plus *Plus) Interpret() string {
	return fmt.Sprintf("(%s + %s)", plus.Expr1.Interpret(), plus.Expr2.Interpret())
}

type Minus struct {
	Expr1, Expr2 Expression
}

func (minus *Minus) Evaluate() float64 {
	return minus.Expr1.Evaluate() - minus.Expr2.Evaluate()
}

func (minus *Minus) Interpret() string {
	return fmt.Sprintf("(%s - %s)", minus.Expr1.Interpret(), minus.Expr2.Interpret())
}
