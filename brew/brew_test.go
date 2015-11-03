package brew

import (
	"testing"
)

func TestSearch(t *testing.T) {
	provider := NewBrewProvider()
	lines, err := provider.Search("python")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("lines: %s", lines)
	}
}

func TestInstall(t *testing.T) {
	provider := NewBrewProvider()
	lines, err := provider.Install("python")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("lines: %s", lines)
	}
}