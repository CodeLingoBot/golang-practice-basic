package stack

type Stack struct {
	i    int
	data [10] int
}

func (s *Stack) Push(k int) {
	s.data[s.i] = k
	s.i++
}

func (s *Stack) PopByIndex(k int) {
	//i = s.data[s.i]
	s.data[k] = 0
	s.i--
	//return
}

func (s *Stack) Pop() {
	s.i--
	s.data[s.i] = 0
}
