package main

import (
	"github.com/codegangsta/cli"
	"github.com/mitallast/please/brew"
	"github.com/mitallast/please/provider"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	app := cli.NewApp()
	app.Name = "please"
	app.Usage = "Polite packet manager"
	app.Commands = []cli.Command{
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "[PACKAGE...]",
			Action:  search,
		},
		{
			Name:    "contains",
			Aliases: []string{"c"},
			Usage:   "[PACKAGE...]",
			Action:  contains,
		},
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "[PACKAGE...]",
			Action:  install,
		},
	}
	app.Run(os.Args)
}

func providers() []provider.Provider {
	return []provider.Provider{
		brew.NewBrewProvider(),
	}
}

func search(c *cli.Context) {
	for _, provider := range providers() {
		founds, err := provider.Search(c.Args()...)
		if err != nil {
			log.Fatal(err)
		} else {
			for _, found := range founds {
				log.Printf("found: %s", found)
			}
		}
	}
}

func contains(c *cli.Context) {
	for _, provider := range providers() {
		founds, err := provider.Contains(c.Args()...)
		if err != nil {
			log.Fatal(err)
		} else {
			for _, found := range founds {
				log.Printf("contains: %s", found)
			}
		}
	}
}

func install(c *cli.Context) {
	for _, provider := range providers() {
		founds, err := provider.Install(c.Args()...)
		if err != nil {
			log.Fatal(err)
		} else {
			for _, found := range founds {
				log.Printf("install: %s", found)
			}
		}
	}
}
