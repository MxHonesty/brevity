package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// This function launches the UI.
func CreateUI() {
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