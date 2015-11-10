package apt

import (
	"testing"
)

func TestSupports(t *testing.T) {
	if Supports() {
		t.Log("apt supports")
	} else {
		t.Skip("apt not supports")
	}
}

func TestSearch(t *testing.T) {
	if !Supports() {
		t.Skip("apt not supports")
	}
	provider := NewProvider()
	lines, err := provider.Search("python")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("found: %s", lines)
	}
}

func TestContains(t *testing.T) {
	if !Supports() {
		t.Skip("apt not supports")
	}
	provider := NewProvider()
	contains, err := provider.Contains("python")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("contains: %s", contains)
	}
}

func TestInstall(t *testing.T) {
	if !Supports() {
		t.Skip("apt not supports")
	}
	provider := NewProvider()
	err := provider.Install("python")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("installed")
	}
}
