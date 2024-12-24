package uicomponent

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func UploadCompanyAndPasswordCSVView(w fyne.Window) fyne.CanvasObject {
	content := container.NewVBox(
		widget.NewLabel("アップロードするお"),
		widget.NewLabel("アップロードするお"),
		widget.NewLabel("アップロードするお======"),
	)
	return content
}
