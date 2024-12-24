package uicomponent

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func UploadCompanyAndPasswordCSVView(w fyne.Window) fyne.CanvasObject {
	content := container.NewStack(
		widget.NewLabel("アップロードするお"),
	)
	return content
}
