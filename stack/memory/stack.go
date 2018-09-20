package memory

//Stack data structure
type Stack struct {
	frames []Frame
}

//Frame is a stack frame
type Frame struct {
	Data interface{}
}

//New returns a new stack
func New() *Stack {
	return &Stack{make([]Frame, 0)}
}

//Init initializes the stack
func (s *Stack) Init(parameters ...interface{}) error {
	s.frames = make([]Frame, 0)
	return nil
}

//Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.frames)
}

//Push an element in the stack
func (s *Stack) Push(data interface{}) {
	s.frames = append(s.frames, Frame{data})
}

//Peek returns the element at the top of the stack
func (s *Stack) Peek() interface{} {
	if s.Size() == 0 {
		return nil
	}

	return s.frames[s.Size()-1].Data
}

//Pop returns and removes the element at the top of the stack
func (s *Stack) Pop() interface{} {
	data := s.Peek()
	if data != nil {
		s.frames = s.frames[:s.Size()-1]
	}
	return data
}
