package npm

import (
	"testing"
)

func TestSearch(t *testing.T) {
	provider := NewNpmProvider()
	lines, err := provider.Search("jquery")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("found: %s", lines)
	}
}

func TestContains(t *testing.T) {
	provider := NewNpmProvider()
	contains, err := provider.Contains("jquery")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("contains: %s", contains)
	}
}

func testInstall(t *testing.T) {
	provider := NewNpmProvider()
	lines, err := provider.Install("jquery")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("install: %s", lines)
	}
}
