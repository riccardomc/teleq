package stackserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/riccardomc/teleq/models"
	"github.com/riccardomc/teleq/stack"
)

//StackServer serves a stack through an httprouter
type StackServer struct {
	Stack  *stack.Stack
	Router *httprouter.Router
	Config *ServerConfig
}

//Size returns the size operation handle of the server
func (server *StackServer) Size() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		response := models.Response{"size", server.Stack.Size()}
		json.NewEncoder(w).Encode(response)
	}
}

//Peek returns the peek operation handle of the server
func (server *StackServer) Peek() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		response := models.Response{"peek", server.Stack.Peek()}
		json.NewEncoder(w).Encode(response)
	}
}

//Push returns the push operation handle of the server
func (server *StackServer) Push() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, rp httprouter.Params) {
		request := models.Request{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		defer r.Body.Close()
		server.Stack.Push(request.Data)
		response := models.Response{"push", request.Data}
		json.NewEncoder(w).Encode(response)
	}
}

//Pop returns the pop operation handle of the server
func (server *StackServer) Pop() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		response := models.Response{"pop", server.Stack.Pop()}
		json.NewEncoder(w).Encode(response)
	}
}

//Serve starts listening and serving
func (server *StackServer) Serve() {
	http.ListenAndServe(":"+strconv.Itoa(server.Config.Port), server.Router)
}

//New gives you a new server, yo
func New(config *ServerConfig) *StackServer {
	server := StackServer{stack.New(), httprouter.New(), config}
	server.Router.GET("/peek", server.Peek())
	server.Router.POST("/push", server.Push())
	server.Router.GET("/pop", server.Pop())
	server.Router.GET("/size", server.Size())
	return &server
}
