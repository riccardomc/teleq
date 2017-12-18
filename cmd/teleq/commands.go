package main

import (
	"fmt"

	"github.com/riccardomc/teleq/client"
	"github.com/riccardomc/teleq/server"
	"github.com/urfave/cli"
)

var (
	ServerAction = server
	sizeAction   = size
	pushAction   = push
	peekAction   = peek
	popAction    = pop
)

func server(c *cli.Context) error {
	config := &stackserver.ServerConfig{c.Int("port")}
	s := stackserver.New(config)
	s.Serve()
	return nil
}

func size(c *cli.Context) error {
	client := client.TeleqClient{}
	response, err := client.Size(c.Parent().String("api"))
	if err != nil {
		fmt.Fprintln(c.App.ErrWriter, err)
		return err
	}
	fmt.Fprintln(c.App.Writer, response)
	return nil
}

func push(c *cli.Context) error {
	client := client.TeleqClient{}
	response, err := client.Push(c.Parent().String("api"), c.String("data"))
	if err != nil {
		fmt.Fprintln(c.App.ErrWriter, err)
		return err
	}
	fmt.Fprintln(c.App.Writer, response)
	return nil
}

func peek(c *cli.Context) error {
	client := client.TeleqClient{}
	response, err := client.Peek(c.Parent().String("api"))
	if err != nil {
		fmt.Fprintln(c.App.ErrWriter, err)
		return err
	}
	fmt.Fprintln(c.App.Writer, response)
	return nil
}

func pop(c *cli.Context) error {
	client := client.TeleqClient{}
	response, err := client.Pop(c.Parent().String("api"))
	if err != nil {
		fmt.Fprintln(c.App.ErrWriter, err)
		return err
	}
	fmt.Fprintln(c.App.Writer, response)
	return nil
}

//New returns a new cli app
func New() *cli.App {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "api, a",
			Value: "http://localhost:9009",
		},
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:   "server",
			Action: ServerAction,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port, p",
					Value: 9009,
				},
			},
		},

		cli.Command{
			Name:   "size",
			Action: sizeAction,
		},

		cli.Command{
			Name:   "peek",
			Action: peekAction,
		},

		cli.Command{
			Name:   "pop",
			Action: popAction,
		},

		cli.Command{
			Name:   "push",
			Action: pushAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "data, d",
				},
			},
		},
	}

	return app
}
