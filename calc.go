package main

import "fmt"
import "strconv"

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

	s := &Stack{
		data: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	var a string
	for {
		if fmt.Scan(&a) != nil {

			switch a {
			case NUMBER:
				push(atoi(data))
			case '+':
				push(pop() + pop())
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
			case '\n':
				printf("\t%.8g\n", pop())
			default:
				printf("error: unknown command %s\n", s)
			}
		}
	}

}
