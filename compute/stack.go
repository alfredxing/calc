package compute

type StringStack struct {
    Slice []string
    Pos int
}

type FloatStack struct {
    Slice []float64
    Pos int
}

func NewStringStack() *StringStack {
    return &StringStack{
        Slice: []string{},
        Pos: -1,
    }
}

func NewFloatStack() *FloatStack {
    return &FloatStack{
        Slice: []float64{},
        Pos: -1,
    }
}

func (s *StringStack) Push(a string) {
    s.Pos++
    if s.Pos < len(s.Slice) {
        s.Slice[s.Pos] = a
    } else {
        s.Slice = append(s.Slice, a)
    }
}

func (s *StringStack) Pop() string {
    ret := s.Top()
    s.Pos--
    return ret
}

func (s *StringStack) Top() string {
    return s.Slice[s.Pos]
}

func (s *FloatStack) Push(a float64) {
    s.Pos++
    if s.Pos < len(s.Slice) {
        s.Slice[s.Pos] = a
    } else {
        s.Slice = append(s.Slice, a)
    }
}

func (s *FloatStack) Pop() float64 {
    ret := s.Top()
    s.Pos--
    return ret
}

func (s *FloatStack) Top() float64 {
    return s.Slice[s.Pos]
}
