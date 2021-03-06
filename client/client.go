package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/riccardomc/teleq/models"
)

//Client interface
type Client interface {
	Size(string) (int, error)
	Peek(string) (interface{}, error)
	Pop(string) (interface{}, error)
	Push(string, interface{}) (interface{}, error)
}

//TeleqClient client
type TeleqClient struct {
}

//Size perform size operation on remote TeleqServer
func (c TeleqClient) Size(serverURL string) (int, error) {
	operationPath := "/size"
	response := &models.Response{}

	u, err := url.Parse(serverURL)
	if err != nil {
		return -1, err
	}
	u.Path = path.Join(u.Path, operationPath)

	r, err := http.Get(u.String())
	if err != nil {
		return -1, err
	}
	if r.StatusCode != 200 {
		return -1, fmt.Errorf("%s %s", r.Status, operationPath)
	}
	json.NewDecoder(r.Body).Decode(response)

	return int(response.Data.(float64)), nil
}

//Peek perform peek operation on remote TeleqServer
func (c TeleqClient) Peek(serverURL string) (interface{}, error) {
	operationPath := "/peek"
	response := &models.Response{}

	u, err := url.Parse(serverURL)
	if err != nil {
		return true, err
	}
	u.Path = path.Join(u.Path, operationPath)

	r, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("%s %s", r.Status, operationPath)
	}
	json.NewDecoder(r.Body).Decode(response)

	return response.Data, nil
}

//Pop perform pop operation on remote TeleqServer
func (c TeleqClient) Pop(serverURL string) (interface{}, error) {
	operationPath := "/pop"
	response := &models.Response{}

	u, err := url.Parse(serverURL)
	if err != nil {
		return true, err
	}
	u.Path = path.Join(u.Path, operationPath)

	r, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("%s %s", r.Status, operationPath)
	}
	json.NewDecoder(r.Body).Decode(response)

	return response.Data, nil
}

//Push perform push operation on remote TeleqServer
func (c TeleqClient) Push(serverURL string, data interface{}) (interface{}, error) {
	operationPath := "/push"
	response := &models.Response{}

	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, operationPath)

	requestBody, _ := json.Marshal(models.Request{data})
	r, err := http.Post(u.String(), "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return true, err
	}
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("%s %s %s", r.Status, operationPath, requestBody)
	}
	json.NewDecoder(r.Body).Decode(response)

	return response.Data, nil
}
