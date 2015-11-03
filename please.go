package main

import (
	"github.com/mitallast/please/brew"
	"github.com/mitallast/please/provider"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	switch len(os.Args) {
	case 1:
		help()
		return
	default:
		switch os.Args[1] {
		case "search":
			search()
			break
		}
	}

}

func search() {
	providers := []provider.Provider{
		brew.NewBrewProvider(),
	}
	for _, provider := range providers {
		founds, err := provider.Search(os.Args[2:]...)
		if err != nil {
			log.Fatal(err)
		} else {
			for _, found := range founds {
				log.Printf("found: %s", found)
			}
		}
	}
}

func help() {
	log.Println("Example usage:")
	log.Println("  please search [PACKAGE...]")
}
