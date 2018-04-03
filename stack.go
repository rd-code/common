package common

import "github.com/pkg/errors"

//实现栈的简单封装
type Stack struct {
    stack *stack
    size  int
}

type stack struct {
    data interface{}
    next *stack
}

func NewStack() *Stack {
    return &Stack{}
}

var nilStackErr = errors.New("the stack is nil")

func (s *Stack) Push(data interface{}) error {
    if s == nil {
        return nilStackErr
    }
    tmp := &stack{
        data: data,
        next: s.stack,
    }
    s.stack = tmp
    s.size ++
    return nil
}

var emptyStackErr = errors.New("the stack is empty")

func (s *Stack) Pop() (interface{}, error) {
    if s == nil {
        return nil, nilStackErr
    }
    if s.stack == nil {
        return nil, emptyStackErr
    }
    result := s.stack.data
    s.stack = s.stack.next
    s.size--
    return result, nil
}

func (s *Stack) IsEmpty() bool {
    return s.stack == nil
}

func (s *Stack) Len() int {
    return s.size
}
