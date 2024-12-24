package uidata

import (
	"fyne.io/fyne/v2"
	"github.com/sunjin110/pdf_manager/desktop/ui/uimodel"
)

var Navigations = uimodel.Navigations{
	{
		ID:    "pdf_lock",
		Title: "PDFパスワード",
		View: func(w fyne.Window) fyne.CanvasObject {
			// TODO
			return nil
		},
	},
	{
		ID:    "regist_pdf_password_with_company",
		Title: "会社ごとのパスワード登録",
		View: func(w fyne.Window) fyne.CanvasObject {
			return nil
		},
	},
}
