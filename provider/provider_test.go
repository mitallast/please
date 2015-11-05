package provider

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockProvider struct {
	mock.Mock
}

func NewMockProvider() *MockProvider {
	return &MockProvider{}
}

func (p *MockProvider) Search(arg ...string) ([]string, error) {
	args := p.Called(arg)
	return args.Get(0).([]string), args.Error(1)
}

func (p *MockProvider) Contains(arg ...string) ([]string, error) {
	args := p.Called(arg)
	return args.Get(0).([]string), args.Error(1)
}

func (p *MockProvider) Install(arg ...string) ([]string, error) {
	args := p.Called(arg)
	return args.Get(0).([]string), args.Error(1)
}

func TestSearch(t *testing.T) {
	provider := NewMockProvider()
	provider.Mock.On("Search", []string{"vim"}).Return([]string{}, nil)
	lines, err := provider.Search("vim")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("found: %s", lines)
	}
}

func TestContains(t *testing.T) {
	provider := NewMockProvider()
	provider.Mock.On("Contains", []string{"vim"}).Return([]string{}, nil)
	contains, err := provider.Contains("vim")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("contains: %v", contains)
	}
}

func TestInstall(t *testing.T) {
	provider := NewMockProvider()
	provider.Mock.On("Install", []string{"vim"}).Return([]string{}, nil)
	lines, err := provider.Install("vim")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("install: %s", lines)
	}
}
