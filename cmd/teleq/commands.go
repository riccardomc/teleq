package main

import (
	"fmt"

	"github.com/riccardomc/teleq/client"
	"github.com/riccardomc/teleq/server"
	"github.com/urfave/cli"
)

func getServerAction(server stackserver.ServerInterface) func(*cli.Context) error {
	return func(c *cli.Context) error {
		port := c.Int("port")
		fmt.Fprintln(c.App.Writer, "Serving on", port)
		server.Serve(port)
		return nil
	}
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
	app.Usage = "a remote stack"
	app.Name = "teleq"
	app.ErrWriter = app.Writer

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "api, a",
			Value: "http://localhost:9009",
		},
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:   "server",
			Action: getServerAction(Server),
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port, p",
					Value: 9009,
				},
			},
		},

		cli.Command{
			Name:   "size",
			Action: size,
		},

		cli.Command{
			Name:   "peek",
			Action: peek,
		},

		cli.Command{
			Name:   "pop",
			Action: pop,
		},

		cli.Command{
			Name:   "push",
			Action: push,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "data, d",
				},
			},
		},
	}

	return app
}
