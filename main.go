package main

import (
	"brevity/repository"
	"brevity/service"
	"brevity/ui"
)

func main() {
	// Next line creates an UI for a local Service that is run over in-memory repositories.
	currentUI := ui.NewUI(service.NewLocalServiceFactory(repository.NewInMemoryRepositoryFactory()))
	currentUI.Run()
}
