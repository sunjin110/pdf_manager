package uidata

import (
	"github.com/sunjin110/pdf_manager/desktop/ui/uicomponent"
	"github.com/sunjin110/pdf_manager/desktop/ui/uimodel"
)

var Navigations = uimodel.Navigations{
	{
		ID:    "pdf_lock",
		Title: "PDFパスワード",
		View:  uicomponent.NewProtectPDFView,
	},
	{
		ID:    "regist_pdf_password_with_company",
		Title: "会社ごとのパスワード登録",
		View:  uicomponent.NewUploadPasswordCSVView,
	},
}
