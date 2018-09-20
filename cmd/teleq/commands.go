package main

import (
	"fmt"

	"github.com/riccardomc/teleq/client"
	"github.com/riccardomc/teleq/server"
	"github.com/riccardomc/teleq/stack"
	"github.com/riccardomc/teleq/stack/database"
	"github.com/riccardomc/teleq/stack/memory"
	"github.com/urfave/cli"
)

func getServerAction(server stackserver.ServerInterface) func(*cli.Context) error {
	return func(c *cli.Context) error {
		var stack stack.Stack
		port := c.Int("port")
		databaseURI := c.String("database")
		if databaseURI == "" {
			fmt.Println("Using in memory stack")
			stack = memory.New()
		} else {
			fmt.Println("Using database stack")
			stack = database.New()
			err := stack.Init(databaseURI)
			if err != nil {
				fmt.Println(c.App.Writer, err)
				return err
			}
		}
		fmt.Fprintln(c.App.Writer, "Serving on", port)
		server.SetPort(port).SetStack(stack).Serve()
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
		{
			Name:   "server",
			Action: getServerAction(Server),
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port, p",
					Value: 9009,
				},
				cli.StringFlag{
					Name:   "database, d",
					Value:  "",
					EnvVar: "TELEQ_DATABASE_HOST",
				},
			},
		},

		{
			Name:   "size",
			Action: size,
		},

		{
			Name:   "peek",
			Action: peek,
		},

		{
			Name:   "pop",
			Action: pop,
		},

		{
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
