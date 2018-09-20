package stackserver

import (
	"github.com/julienschmidt/httprouter"
	"github.com/riccardomc/teleq/stack"
)

//ServerInterface represents a server
type ServerInterface interface {
	Size() httprouter.Handle
	Peek() httprouter.Handle
	Push() httprouter.Handle
	Pop() httprouter.Handle
	SetPort(int) ServerInterface
	SetRouter(*httprouter.Router) ServerInterface
	SetStack(stack.Stack) ServerInterface
	Serve()
}
