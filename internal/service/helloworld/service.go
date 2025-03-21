package helloworld

import (
	"context"
	"errors"

	"github.com/Hackail/rest-api-go/pkg/generic"
)

type (
	Service struct {
		helloWorldRepository helloWorldRepository
	}

	helloWorldRepository interface {
		GetHelloWorld(context.Context) (string, error)
	}
)

func NewHelloWorldService(helloWorldRepository helloWorldRepository) (*Service, error) {
	if generic.IsNil(helloWorldRepository) {
		return nil, errors.New("helloWorldService must not be nil")
	}
	return &Service{helloWorldRepository: helloWorldRepository}, nil
}

func (s Service) GetHelloWorld(ctx context.Context) (string, error) {
	return s.helloWorldRepository.GetHelloWorld(ctx)
}
