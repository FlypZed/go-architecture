package helloworld

import (
	"context"
	"errors"
	"net/http"

	"github.com/Hackail/rest-api-go/pkg/generic"
)

type (
	Controller struct {
		helloWorldService helloWorldService
	}

	helloWorldService interface {
		GetHelloWorld(context.Context) (string, error)
	}
)

func NewHelloWorldController(helloWorldService helloWorldService) (*Controller, error) {
	if generic.IsNil(helloWorldService) {
		return nil, errors.New("helloWorldService must not be nil")
	}
	return &Controller{helloWorldService: helloWorldService}, nil
}

func (c *Controller) GetHelloWorld(w http.ResponseWriter, r *http.Request) error {
	resp, err := c.helloWorldService.GetHelloWorld(context.WithValue(r.Context(), "value", "hello-world"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte(resp)); err != nil {
		http.Error(w, "error unmarshalling data", http.StatusInternalServerError)
		return err
	}
	return nil
}
