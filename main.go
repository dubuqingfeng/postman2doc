package main

import (
	"os"

	"gopkg.in/urfave/cli.v2"
	"github.com/fatih/color"
	"fmt"
	"flag"
	"github.com/dubuqingfeng/postman2doc/postman"
)

func main() {
	app := &cli.App{
		Name:    "postman2doc",
		Version: "1.0",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Dubu qingfeng",
				Email: "1135326346@qq.com",
			},
		},
		Copyright: "(c) 2018 Earth",
		HelpName:  "contrive",
		Usage:     "Postman exports documents.",
		UsageText: "Postman exports documents cli.",
		Action: func(c *cli.Context) error {
			fmt.Printf("Hello %q\n", c.Args().First())
			fmt.Println(color.GreenString("Begin."))
			set := flag.NewFlagSet("contrive", 0)
			nc := cli.NewContext(c.App, set, c)
			param := postman.Param{
				Template:    nc.String("template"),
				IsDirectory: nc.Bool("directory"),
				Collection:  nc.String("collection"),
				Output:      c.Args().First(),
			}
			postman.Handle(param)
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{Value: false, Aliases: []string{"r"}, Name: "directory"},
			&cli.StringFlag{Value: "template/default.md", Name: "template", Aliases: []string{"t"}},
			&cli.StringFlag{Value: "", Name: "collection", Aliases: []string{"c"}},
		},
	}

	app.Run(os.Args)
}
