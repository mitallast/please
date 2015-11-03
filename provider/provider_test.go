package provider

import (
	"github.com/stretchr/testify/mock"
	"testing"
	"log"
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

func (p *MockProvider) Install(arg ...string) ([]string, error) {
	args := p.Called(arg)
	return args.Get(0).([]string), args.Error(1)
}

func TestSearch(t *testing.T) {
	var provider *MockProvider
	provider = NewMockProvider()
	provider.Mock.On("Search", []string{"vim"}).Return([]string{}, nil)
	lines, err := provider.Search("vim")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("lines: %s", lines)
	}
}

func TestInstall(t *testing.T) {
	var provider *MockProvider
	provider = NewMockProvider()
	provider.Mock.On("Install", []string{"vim"}).Return([]string{}, nil)
	lines, err := provider.Install("vim")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("lines: %s", lines)
	}
}