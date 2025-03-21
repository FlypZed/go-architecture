package helloworld

import (
	"context"
	"errors"

	"github.com/Hackail/rest-api-go/pkg/generic"
)

type Repository struct {
}

func NewHelloWorldRepository() *Repository {
	return &Repository{}
}

func (r Repository) GetHelloWorld(ctx context.Context) (string, error) {
	value := ctx.Value("value")
	if generic.IsNil(value) {
		return "", errors.New("value is nil")
	}
	return value.(string), nil
}
