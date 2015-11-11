package yum

import (
	"testing"
)

func TestSupports(t *testing.T) {
	if Supports() {
		t.Log("yum supports")
	} else {
		t.Skip("yum not supports")
	}
}

func TestSearch(t *testing.T) {
	if !Supports() {
		t.Skip("yum not supports")
	}
	provider := NewProvider()
	lines, err := provider.Search("tree")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("found: %s", lines)
	}
}

func TestContains(t *testing.T) {
	if !Supports() {
		t.Skip("yum not supports")
	}
	provider := NewProvider()
	contains, err := provider.Contains("tree")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("contains: %s", contains)
	}
}

func TestInstall(t *testing.T) {
	if !Supports() {
		t.Skip("yum not supports")
	}
	provider := NewProvider()
	err := provider.Install("tree")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("installed")
	}
}
