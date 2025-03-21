package server

import (
	"log"

	helloWorldCtrl "github.com/Hackail/rest-api-go/internal/controller/helloworld"
	helloWorldRepo "github.com/Hackail/rest-api-go/internal/repository/helloworld"
	helloWorldService "github.com/Hackail/rest-api-go/internal/service/helloworld"
)

func resolveHelloWorldController() *helloWorldCtrl.Controller {
	helloWorldCtrlClient, err := helloWorldCtrl.NewHelloWorldController(resolveHelloWorldService())
	if err != nil {
		panicHandler(err)
	}
	return helloWorldCtrlClient
}

func resolveHelloWorldService() *helloWorldService.Service {
	helloWorldServiceClient, err := helloWorldService.NewHelloWorldService(resolveHelloWorldRepository())
	if err != nil {
		panicHandler(err)
	}
	return helloWorldServiceClient
}

func resolveHelloWorldRepository() *helloWorldRepo.Repository {
	return helloWorldRepo.NewHelloWorldRepository()
}

func panicHandler(err error) {
	log.Panicf("error handled while creatin instance: %s", err.Error())
}
