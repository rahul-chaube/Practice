package stack

type Stack []int

func (s *Stack) Push(data int) {
	*s = append(*s, data)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return top
}
