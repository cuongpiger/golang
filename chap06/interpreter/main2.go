package main

import (
	"strconv"
	"strings"
)

const (
	SUM1 = "sum"
	SUB1 = "sub"
)

type Interpreter interface {
	Read() int
}

type value int

func (v *value) Read() int {
	return int(*v)
}

type operationSum struct {
	Left  Interpreter
	Right Interpreter
}

func (a *operationSum) Read() int {
	return a.Left.Read() + a.Right.Read()
}

type operationSubtract struct {
	Left  Interpreter
	Right Interpreter
}

func (s *operationSubtract) Read() int {
	return s.Left.Read() - s.Right.Read()
}

func operatorFactory(o string, left, right Interpreter) Interpreter {
	switch o {
	case SUM1:
		return &operationSum{
			Left:  left,
			Right: right,
		}
	case SUB1:
		return &operationSubtract{
			Left:  left,
			Right: right,
		}
	}

	return nil
}

type polishNotationStack1 []Interpreter

func (p *polishNotationStack1) Push(s Interpreter) {
	*p = append(*p, s)
}

func (p *polishNotationStack1) Pop() Interpreter {
	length := len(*p)

	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}

	return nil
}

func main() {
	stack := polishNotationStack1{}
	operators := strings.Split("3 4 sum 2 sub", " ")

	for _, operatorString := range operators {
		if operatorString == SUM1 || operatorString == SUB1 {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := operatorFactory(operatorString, left, right)
			res := value(mathFunc.Read())
			stack.Push(&res)
		} else {
			val, err := strconv.Atoi(operatorString)
			if err != nil {
				panic(err)
			}

			temp := value(val)
			stack.Push(&temp)
		}
	}

	println(int(stack.Pop().Read()))
}
