package npm

import (
	"testing"
)

func TestSupports(t *testing.T) {
	if Supports() {
		t.Log("npm supports")
	} else {
		t.Skip("npm not supports")
	}
}

func TestSearch(t *testing.T) {
	if !Supports() {
		t.Skip("npm not supports")
	}
	provider := NewNpmProvider()
	lines, err := provider.Search("jquery-sortable")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("found: %s", lines)
	}
}

func TestContains(t *testing.T) {
	if !Supports() {
		t.Skip("npm not supports")
	}
	provider := NewNpmProvider()
	contains, err := provider.Contains("jquery-sortable")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("contains: %s", contains)
	}
}

func testInstall(t *testing.T) {
	if !Supports() {
		t.Skip("npm not supports")
	}
	provider := NewNpmProvider()
	lines, err := provider.Install("jquery-sortable")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("install: %s", lines)
	}
}
