package uicomponent

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/sunjin110/pdf_manager/core"
)

func UploadCompanyAndPasswordCSVView(w fyne.Window, pdfManagerCore core.Core) fyne.CanvasObject {

	textArea := widget.NewMultiLineEntry()
	textArea.SetPlaceHolder("アップロード内容を反映する")

	uploadButton := widget.NewButton("Upload File", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if reader == nil {
				return
			}
			defer reader.Close()

			if err := pdfManagerCore.RegistPasswordByCSV(reader); err != nil {
				dialog.ShowError(err, w)
				textArea.SetText("failed")
				return
			}

			textArea.SetText("success!")
		}, w)
	})

	content := container.NewVBox(
		widget.NewLabel("会社ごとのパスワード登録"),
		textArea,
		uploadButton,
	)
	return content
}
