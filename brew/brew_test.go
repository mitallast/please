package brew

import (
	"log"
	"testing"
)

func TestSearch(t *testing.T) {
	provider := NewBrewProvider()
	lines, err := provider.Search("vim")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("lines: %s", lines)
	}
}
