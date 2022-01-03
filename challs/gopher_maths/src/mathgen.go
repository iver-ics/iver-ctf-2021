package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Operator byte

const (
	OpAdd Operator = iota
	OpSub
	OpMul
	OpDiv
	OpMod
)

func (op Operator) EvalInt64(lhs, rhs int64) int64 {
	switch op {
	case OpAdd:
		return lhs + rhs
	case OpSub:
		return lhs - rhs
	case OpMul:
		return lhs * rhs
	case OpDiv:
		return lhs / rhs
	case OpMod:
		return lhs % rhs
	default:
		panic(fmt.Sprintf("invalid operator: %T(%d)", op, op))
	}
}

func (op Operator) String() string {
	switch op {
	case OpAdd:
		return "+"
	case OpSub:
		return "-"
	case OpMul:
		return "*"
	case OpDiv:
		return "/"
	case OpMod:
		return "%"
	default:
		panic(fmt.Sprintf("invalid operator: %T(%d)", op, op))
	}
}

func RandomOperator() Operator {
	return Operator(byte(rand.Intn(5)))
}

type MathProblem struct {
	LeftOperand  int64
	RightOperand int64
	Operator     Operator
	Result       int64
}

func (mp MathProblem) String() string {
	return fmt.Sprintf("%d %s %d = %d", mp.LeftOperand, mp.Operator, mp.RightOperand, mp.Result)
}

const halfInt64Size = math.MaxInt64 / 2

func GenMathProblem() MathProblem {
	prob := MathProblem{
		Operator: RandomOperator(),
	}
	switch prob.Operator {
	case OpAdd, OpSub:
		prob.LeftOperand = RandInt64(-halfInt64Size, halfInt64Size)
		prob.RightOperand = RandInt64(-halfInt64Size, halfInt64Size)
		prob.Result = prob.Operator.EvalInt64(prob.LeftOperand, prob.RightOperand)
	case OpMul:
		a := RandInt64(-922337203685, 922337203685)
		b := RandInt64(-10000000, 10000000)
		if rand.Intn(2) == 0 {
			a, b = b, a
		}
		prob.LeftOperand = a
		prob.RightOperand = b
		prob.Result = prob.Operator.EvalInt64(a, b)
	case OpDiv:
		a := RandInt64(-922337203685, 922337203685)
		b := RandInt64(-10000000, 10000000)
		if b == 0 {
			b = 102
		}
		prob.Result = a / b
		prob.LeftOperand = b * prob.Result // fixes rounding errors
		prob.RightOperand = b
	case OpMod:
		prob.LeftOperand = RandInt64(halfInt64Size, math.MaxInt64)
		prob.RightOperand = RandInt64(1251, halfInt64Size)
		prob.Result = prob.LeftOperand % prob.RightOperand
	default:
		panic(fmt.Sprintf("invalid operator: %T(%d)", prob.Operator, prob.Operator))
	}
	return prob
}

func RandInt64(lower, upper int64) int64 {
	return lower + rand.Int63n(upper-lower+1)
}
