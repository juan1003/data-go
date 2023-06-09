package main

import "fmt"

type Stack struct {
    stack []int
}

func (s *Stack) push(item int) {
    s.stack = append(s.stack, item)
}

func (s *Stack) pop() int {
    if s.isEmpty() {
        return -1
    }
    item := s.stack[len(s.stack)-1]
    s.stack = s.stack[:len(s.stack)-1]
    return item
}

func (s *Stack) isEmpty() bool {
    return len(s.stack) == 0
}

func (s *Stack) size() int {
    return len(s.stack)
}

func (s *Stack) peek() int {
    if s.isEmpty() {
        return -1
    }
    return s.stack[len(s.stack)-1]
}

func main() {
    stack := &Stack{}
    stack.push(10)
    stack.push(5)
    stack.push(7)

    fmt.Println(stack.peek())
    
    fmt.Println(stack.pop())
    fmt.Println(stack.pop())

    fmt.Println(stack.size())
    fmt.Println(stack.isEmpty())
}
