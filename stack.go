package main

type Stack struct {
	ary [][][]float64
}

func NewRCS() *Stack {
	stack := new(Stack)
	m := NewMatrix()
	MakeIdentity(m)
	stack.ary = [][][]float64{m}
	return stack
}

func (s *Stack) Push() {
	head := Copy2d(*s.Peek())
  s.ary = append(s.ary, head)
}

func (s *Stack) Pop() {
	s.ary = s.ary[:len(s.ary) - 1]
}

func (s *Stack) Peek() *[][]float64 {
	return &s.ary[len(s.ary) - 1]
}

func (s *Stack) Add(m [][]float64) {
	s.ary = append(s.ary, m)
}

func Copy2d(m [][]float64) [][]float64 {
	duplicate := make([][]float64, len(m))
	for i := range m {
    duplicate[i] = make([]float64, len(m[i]))
    copy(duplicate[i], m[i])
	}
	return duplicate
}
