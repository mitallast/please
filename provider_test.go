package please

import "github.com/stretchr/testify/mock"

type mockProvider struct {
    mock.Mock
}

func (p *mockProvider) search(arg ...string) ([]byte, error) {
    args := p.Called(arg)
    return args.Get(0).([]byte), args.Error(1)
}

func (p *mockProvider) install(arg ...string) ([]byte, error) {
    args := p.Called(arg)
    return args.Get(0).([]byte), args.Error(1)
}

