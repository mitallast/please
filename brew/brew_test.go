package brew

import (
	"log"
	"testing"
)

func TestSearch(t *testing.T) {
	provider := newBrewProvider()
	lines, err := provider.search("vim")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("lines: %s", lines)
	}
}
