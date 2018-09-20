package stack

// Stack defines the interface of a Stack data structure
type Stack interface {
	Size() int
	Push(interface{})
	Peek() interface{}
	Pop() interface{}
	Init(...interface{}) error
}
