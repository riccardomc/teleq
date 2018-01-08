package stackserver

import "github.com/julienschmidt/httprouter"

//ServerInterface represents a server
type ServerInterface interface {
	Size() httprouter.Handle
	Peek() httprouter.Handle
	Push() httprouter.Handle
	Pop() httprouter.Handle
	Serve(int)
}
