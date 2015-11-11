package main

import (
	"github.com/codegangsta/cli"
	"github.com/mitallast/please/apt"
	"github.com/mitallast/please/yum"
	"github.com/mitallast/please/brew"
	"github.com/mitallast/please/npm"
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
	providers := []provider.Provider{}
	if brew.Supports() {
		providers = append(providers, brew.NewProvider())
	}
	if apt.Supports() {
		providers = append(providers, apt.NewProvider())
	}
	if yum.Supports() {
		providers = append(providers, yum.NewProvider())
	}
	if npm.Supports() {
		providers = append(providers, npm.NewProvider())
	}
	return providers
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
	packages := append([]string{}, c.Args()...)
	for _, provider := range providers() {
		contains, err := provider.Contains(packages...)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("contains: %s", contains)
		if len(contains) > 0 {
			packages = excludePackages(packages, contains)
			if err := provider.Install(contains...); err != nil {
				log.Fatal(err)
			}
		}
		if len(packages) == 0 {
			break
		}
	}
}

func excludePackages(packages []string, exclude []string) []string {
	list := []string{}
	for _, pkg := range packages {
		contains := containsPackage(exclude, pkg)
		if !contains {
			list = append(list, pkg)
		}
	}
	return list
}

func containsPackage(packages []string, pkg string) bool {
	for _, p := range packages {
		if p == pkg {
			return true
		}
	}
	return false
}
