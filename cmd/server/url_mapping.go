package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Hackail/rest-api-go/internal/controller/helloworld"
	"github.com/Hackail/rest-api-go/pkg/infrastructure"
)

type mapping struct {
	helloWorldController helloworld.Controller
}

func newMapping() *mapping {
	return &mapping{
		helloWorldController: *resolveHelloWorldController(),
	}
}

func (m mapping) mapUrlsToControllers(router *chi.Mux) {
	router.Method(http.MethodGet, "/hello-world", infrastructure.Handle(m.helloWorldController.GetHelloWorld))
}
