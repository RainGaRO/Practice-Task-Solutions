package stack

// func UseStack() {
// 	stack := NewStack()
// 	stack.Push(1)
// 	stack.Push(2)

// 	fmt.Println(stack.Peek())
// 	fmt.Println(stack.Pop())
// 	fmt.Println(stack.Peek())

// 	stack.Push("String")
// 	fmt.Println(stack.Size())
// 	fmt.Println(stack.Peek())
// 	fmt.Println(stack.Pop())
// 	fmt.Println(stack.Peek())
// }

type Stack struct {
	Data []interface{}
}

func NewStack() *Stack {
	return &Stack{nil}
}

func (s *Stack) Size() int {
	return len(s.Data)
}

func (s *Stack) Push(v interface{}) {
	s.Data = append(s.Data, v)
}

func (s *Stack) Pop() (interface{}, bool) {
	pop := s.Data[len(s.Data)-1]
	if len(s.Data) > 0 {
		s.Data = s.Data[:len(s.Data)-1] //или append?
		return s.Data, true
	}
	return pop, false
}

func (s *Stack) Peek() (interface{}, bool) {
	if len(s.Data) > 0 {
		return s.Data[len(s.Data)-1], true
	}
	return s.Data[len(s.Data)-1], false
}

func (s *Stack) IsEmpty() bool {
	if len(s.Data) != 0 {
		return false
	}
	return true
}
