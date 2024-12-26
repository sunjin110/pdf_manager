package uicomponent

import (
	"context"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/sunjin110/pdf_manager/core"
	"github.com/sunjin110/pdf_manager/core/domain/model"
)

type uploadPasswordCSVView struct {
	w              fyne.Window
	pdfManagerCore core.Core
	passwords      []model.Password
	table          *widget.Table
	uploadButton   *widget.Button
	headerLabel    *widget.Label
}

func NewUploadPasswordCSVView(w fyne.Window, pdfManagerCore core.Core) fyne.CanvasObject {

	// データを最初に取得しておく
	passwords, err := pdfManagerCore.GetAllPasswords(context.Background())
	if err != nil {
		dialog.ShowError(err, w)
		return container.NewVBox(widget.NewLabel("読み込みに失敗しました"))
	}

	sort.Slice(passwords, func(i, j int) bool {
		return passwords[i].TargetName < passwords[j].TargetName
	})

	view := &uploadPasswordCSVView{
		w:              w,
		pdfManagerCore: pdfManagerCore,
		headerLabel:    widget.NewLabel("会社ごとのパスワード登録"),
		passwords:      passwords,
	}

	view.table = view.makePasswordTableUI()
	view.uploadButton = view.makeUploadButtonUI()
	return container.NewBorder(
		widget.NewLabel("会社ごとのパスワード登録"),
		view.uploadButton,
		nil,
		nil,
		view.table,
	)
}

func (view *uploadPasswordCSVView) makePasswordTableUI() *widget.Table {
	table := widget.NewTableWithHeaders(
		// 行数、列数
		func() (int, int) {
			return len(view.passwords), 2
		},
		// 各セルを生成するためのコンテンツ
		func() fyne.CanvasObject {
			// テーブルのセルは同じ型の CanvasObject である必要がある
			// ここでは単純に Label を使う
			return widget.NewLabel("")
		},
		// ID (行・列) とセルObjectを渡されるので、表示内容をセットする
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			if id.Row < len(view.passwords) {
				switch id.Col {
				case 0:
					label.SetText(view.passwords[id.Row].TargetName)
				case 1:
					label.SetText(view.passwords[id.Row].Password)
				}
			}
		},
	)

	table.SetColumnWidth(0, 200)
	table.SetColumnWidth(1, 200)

	return table
}

func (view *uploadPasswordCSVView) makeUploadButtonUI() *widget.Button {
	return widget.NewButton("Upload File", func() {
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, view.w)
				return
			}
			if reader == nil {
				// ユーザがキャンセルした
				return
			}
			defer reader.Close()

			// CSVファイルの内容をPDF Manager Coreに登録
			if err := view.pdfManagerCore.RegistPasswordByCSV(reader); err != nil {
				dialog.ShowError(err, view.w)
				return
			}
			dialog.ShowInformation("info", "アップロードに成功", view.w)

			// 登録完了後、再度データを取得し直す
			view.reloadData()
		}, view.w)

		// CSVだけ開けるようにフィルター
		fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".csv"}))
		fileDialog.Show()
	})
}

func (view *uploadPasswordCSVView) reloadData() {
	// パスワード一覧を再取得
	newData, err := view.pdfManagerCore.GetAllPasswords(context.Background())
	if err != nil {
		dialog.ShowError(err, view.w)
		return
	}

	sort.Slice(newData, func(i, j int) bool {
		return newData[i].TargetName < newData[j].TargetName
	})

	// スライスを差し替え
	view.passwords = newData

	// テーブル再描画
	// テーブルの OnCreateCell や OnUpdateCell が再度呼ばれ、最新のデータが反映される
	view.table.Refresh()
}
