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
		case "install":
			install()
			break
		}
	}

}

func providers() []provider.Provider {
	return []provider.Provider{
		brew.NewBrewProvider(),
	}
}

func search() {
	for _, provider := range providers() {
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

func install() {
	for _, provider := range providers() {
		founds, err := provider.Install(os.Args[2:]...)
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
	log.Println("  please install [PACKAGE...]")
}
