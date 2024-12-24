package uicomponent

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func UploadCompanyAndPasswordCSVView(w fyne.Window) fyne.CanvasObject {

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

			// ファイルの中身を読み取る
			// TODO CSVの内容を読み取る様にする
			content, err := io.ReadAll(reader)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			textArea.SetText(string(content))

		}, w)
	})

	content := container.NewVBox(
		widget.NewLabel("会社ごとのパスワード登録"),
		textArea,
		uploadButton,
	)
	return content
}
