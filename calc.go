package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var max int = 100000

type Stack struct {
	data []int
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) isFull() bool {
	return len(s.data) == max
}

func (s *Stack) push(n int) {
	if s.isFull() {
		fmt.Println("error: stack is full")
		return
	}
	s.data = append(s.data, n)
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		fmt.Println("error: stack is empty")
		return 0
	}
	last := len(s.data) - 1
	popData := s.data[last]
	s.data = s.data[:last]
	return popData
}

func main() {

	var s *Stack = &Stack{}

	var a string

	for {
		if sc.Split(bufio.ScanWords) !=  {

			switch a {
			case '0' < a && a < '9':
				s.push(strconv.Atoi(a))
			case '+':
				s.push(s.pop() + s.pop())
				// 		case '*':
				// 			push(pop() * pop())
				// 		case '-':
				// 			op2 = pop()
				// 			push(pop() - op2)
				// 		case '/':
				// 			op2 = pop()
				// 			if (op2 != 0.0) {
				// 				push(pop() / op2)
				// 			} else {
				// 				printf("error: zero divisor\n")
				// 			}
				// 			case '\n':
				// 				fmt.Println("\t%.8g\n", s.pop())
			default:
				fmt.Println("error: unknown command %s\n", a)
			}
		}
	}

}
