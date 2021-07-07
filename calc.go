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
	a = "m"

	for {
		//if fmt.Scan(a) != 'm' {

		switch a {
		case "1":
			tmp, _ := strconv.Atoi(a)
			s.push(tmp)
		case "2":
			tmp, _ := strconv.Atoi(a)
			s.push(tmp)

		case "3":
			tmp, _ := strconv.Atoi(a)
			s.push(tmp)

		case "4":
			tmp, _ := strconv.Atoi(a)
			s.push(tmp)

		case "5":
			tmp, _ := strconv.Atoi(a)
			s.push(tmp)

		case "6":
			tmp, _ := strconv.Atoi(a)
			s.push(tmp)

		case "7":
			tmp, _ := strconv.Atoi(a)
			s.push(tmp)

		case "8":
			tmp, _ := strconv.Atoi(a)
			s.push(tmp)

		case "9":
			tmp, _ := strconv.Atoi(a)
			s.push(tmp)

		case "0":
			tmp, _ := strconv.Atoi(a)
			s.push(tmp)

		case "m":
			fmt.Println("Kimura")
			return
		case "+":
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
		//}
	}

}
