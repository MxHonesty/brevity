package ui

import (
	"brevity/service"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Represents the UI. Stores the Services that it operates on. This is an
// abstraction that builds over the two types of services.
// 	1. Local Service
//		Service that works locally.
// 	2. Remote Service
//		Service that communicates with the backend.
type UI struct {
	scheduableSrv service.AbsScheduableService
	depSrv service.AbsDependencyService
}

// Create a new instance of UI. Takes as argument a service.AbstractServiceFactory.
// Uses the factory to create new service instances that it will serve.
func NewUI(srvFactory service.AbstractServiceFactory) *UI {
	return &UI{scheduableSrv: srvFactory.ScheduableService(),
		depSrv: srvFactory.DependencyService()}
}

// This method launches the UI.
func (ui *UI) run() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Brevity!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("HI!", func() {
			hello.SetText("Welcome xD")
		}),
	))
	w.ShowAndRun()
}
