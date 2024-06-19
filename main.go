package main

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("ONIR Editor")

	editor := widget.NewMultiLineEntry()
	content := container.NewStack(editor)

	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New", func() { editor.SetText("") }),
			fyne.NewMenuItem("Open", func() {
				dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
					if err == nil && reader != nil {
						b, err := io.ReadAll(reader)
						if err == nil {
							editor.SetText(string(b))
						}
					}
				}, myWindow)
			}),
			fyne.NewMenuItem("Save", func() {
				dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
					if err == nil && writer != nil {
						_, err := writer.Write([]byte(editor.Text))
						if err != nil {
							dialog.ShowError(err, myWindow)
						}
					}
				}, myWindow)
			}),
		),
	)

	myWindow.SetMainMenu(mainMenu)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(600, 400))

	myWindow.ShowAndRun()
}
