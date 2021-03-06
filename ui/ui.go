package ui

import (
	"brevity/service"
	"brevity/service_abstract"
	"fyne.io/fyne/v2"
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
	scheduableSrv service_abstract.AbsScheduableService
	depSrv        service_abstract.AbsDependencyService

	MainHBox *fyne.Container  // The Main Layout that has to be displayed.
	SideMenuVBox *fyne.Container
	TableVBox *fyne.Container
	CalendarBox *widget.Table
}

// Method for setting up the layouts.
func (ui *UI) setUpLayout() {
	ui.MainHBox = container.NewHBox()  // The main Layout for the ui.
	ui.SideMenuVBox = container.NewVBox()
	ui.TableVBox = container.NewVBox()

	ui.TableVBox.Add(ui.CalendarBox)

	ui.MainHBox.Add(ui.SideMenuVBox)
	ui.MainHBox.Add(ui.TableVBox)
}

// Method for initializing the content inside the layouts.
func (ui *UI) initContent() {
	hello := widget.NewLabel("Hello Brevity!")
	hello2 := widget.NewLabel("Hello Brevity!")
	hello3 := widget.NewLabel("Hello Brevity!")
	ui.SideMenuVBox.Add(hello)
	ui.SideMenuVBox.Add(hello2)
	ui.SideMenuVBox.Add(hello3)
}

// Creates a table full of test values for Labels.
func (ui *UI) initTable() {
	ui.CalendarBox = widget.NewTable(
		func() (int, int) {return 7, 8},
		func() fyne.CanvasObject {
			return widget.NewLabel("TEST")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText("TEST1")
		})
}

// Create a new instance of UI. Takes as argument a service.AbstractServiceFactory.
// Uses the factory to create new service instances that it will serve.
func NewUI(srvFactory service.AbstractServiceFactory) *UI {
	ui := &UI{scheduableSrv: srvFactory.ScheduableService(),
		depSrv: srvFactory.DependencyService()}

	return ui
}

// This method launches the UI.
func (ui *UI) Run() {
	a := app.New()
	w := a.NewWindow("Hello")

	ui.initTable()
	ui.setUpLayout()
	ui.initContent()

	w.SetContent(ui.MainHBox)
	w.ShowAndRun()
}
