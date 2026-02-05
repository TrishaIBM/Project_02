package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Health Checker",
		Usage: "A tiny tool that checks whether the website is running or not",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "The domain to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "The port to check",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if port == "" {
				port = "80"
			}
			// Call the capitalized Check from check.go
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
