package uicomponent

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sunjin110/pdf_manager/core"
)

func ProtectPDFView(w fyne.Window, pdfManagerCore core.Core) fyne.CanvasObject {
	content := container.NewStack(
		widget.NewLabel("PDFを一覧で選択したら、それをロックするで"),
	)
	return content
}
