package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
	"github.com/riccardomc/teleq/models"
	"github.com/riccardomc/teleq/stack"
)

//StackServer serves a stack through an httprouter
type StackServer struct {
	Stack  *stack.Stack
	Router *httprouter.Router
}

func size(server *StackServer) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		response := models.Response{"size", server.Stack.Size()}
		json.NewEncoder(w).Encode(response)
	}
}

func peek(server *StackServer) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		response := models.Response{"peek", server.Stack.Peek()}
		json.NewEncoder(w).Encode(response)
	}
}

func push(server *StackServer) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, rp httprouter.Params) {
		value, err := url.PathUnescape(rp.ByName("value"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		server.Stack.Push(value)
		response := models.Response{"push", value}
		json.NewEncoder(w).Encode(response)
	}
}

func pop(server *StackServer) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		response := models.Response{"pop", server.Stack.Pop()}
		json.NewEncoder(w).Encode(response)
	}
}

//NewStackServer gives you a new server, yo
func NewStackServer() *StackServer {
	server := StackServer{stack.New(), httprouter.New()}
	server.Router.GET("/peek", peek(&server))
	server.Router.POST("/push/:value", push(&server))
	server.Router.GET("/pop", pop(&server))
	server.Router.GET("/size", size(&server))
	return &server
}