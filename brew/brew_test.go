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
		t.Logf("found: %s", lines)
	}
}

func TestContains(t *testing.T) {
	provider := NewBrewProvider()
	contains, err := provider.Contains("python")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("contains: %s", contains)
	}
}

func TestInstall(t *testing.T) {
	provider := NewBrewProvider()
	lines, err := provider.Install("python")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("install: %s", lines)
	}
}
