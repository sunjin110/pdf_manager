package uicomponent

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/sunjin110/pdf_manager/core"
	"github.com/sunjin110/pdf_manager/core/domain/model"
)

type protectPDFView struct {
	w                           fyne.Window
	pdfManagerCore              core.Core
	pdfs                        pdfs
	table                       *widget.Table
	generateProtectedPDFsButton *widget.Button
}

func NewProtectPDFView(w fyne.Window, pdfManagerCore core.Core) fyne.CanvasObject {
	view := &protectPDFView{
		w:              w,
		pdfManagerCore: pdfManagerCore,
	}

	view.table = view.makeTableUI()
	view.generateProtectedPDFsButton = view.makeGenerateProtectedPDFsButtonUI()

	top := container.NewVBox(
		widget.NewLabel("PDFを一覧で選択したら、それをロックする"),
		view.makeUploadPDFsButtonUI(),
	)

	return container.NewBorder(
		top,
		view.generateProtectedPDFsButton,
		nil,
		nil,
		view.table,
	)
}

func (view *protectPDFView) makeUploadPDFsButtonUI() *widget.Button {
	return widget.NewButton("Upload PDFs(PDFが入っているフォルダを選択してください)", func() {

		// dialog.NewFolderOpen でフォルダを選択させる
		folderDialog := dialog.NewFolderOpen(func(folder fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, view.w)
				return
			}
			// ユーザーがフォルダ選択をキャンセルした場合
			if folder == nil {
				return
			}

			// フォルダ内のすべてのエントリを取得
			children, err := folder.List()
			if err != nil {
				dialog.ShowError(err, view.w)
				return
			}

			// フォルダ内のファイルを走査し、.pdf のみを処理
			var pdfFiles []fyne.URI
			for _, child := range children {
				// child は fyne.URI
				// 拡張子が .pdf のものだけピックアップ
				if strings.EqualFold(filepath.Ext(child.Name()), ".pdf") {
					pdfFiles = append(pdfFiles, child)
				}
			}

			if len(pdfFiles) == 0 {
				dialog.ShowInformation("No PDF", "選択したフォルダ内にPDFがありません", view.w)
				return
			}

			// PDFファイルをまとめて処理
			// for _, pdfURI := range pdfFiles {
			// 	// URI から Reader を取得
			// 	reader, err := storage.Reader(pdfURI)
			// 	if err != nil {
			// 		log.Printf("failed to open file: %v\n", err)
			// 		continue
			// 	}
			// 	defer reader.Close()

			// 	// ここで PDF の読み込みや処理を行う
			// 	// 例: ファイル名の表示など
			// 	fmt.Println("Found PDF:", pdfURI.Name())
			// }

			companyNames := []string{}
			for _, pdfFile := range pdfFiles {
				companyName, err := extractCompanyName(pdfFile.Name())
				if err != nil {
					dialog.ShowError(err, view.w)
					continue
				}
				companyNames = append(companyNames, companyName)
			}
			passwords, err := view.pdfManagerCore.GetPasswordsByTargetNames(context.Background(), companyNames)
			if err != nil {
				dialog.ShowError(err, view.w)
				return
			}

			view.pdfs = newPDFs(pdfFiles, passwords)
			view.generateProtectedPDFsButton.Show()
			view.table.Refresh()

			dialog.ShowInformation("Complete", "フォルダ内のPDFを処理しました", view.w)
		}, view.w)

		// ダイアログを表示
		folderDialog.Show()
	})
}

func (view *protectPDFView) makeGenerateProtectedPDFsButtonUI() *widget.Button {
	button := widget.NewButton("鍵付きpdfを作成", func() {
		protectedPDFs := []protectedPDF{}
		for _, pdf := range view.pdfs {
			reader, err := storage.Reader(pdf.pdfFile)
			if err != nil {
				dialog.ShowError(err, view.w)
				return
			}

			readSeeker, err := readSeekerFromReadCloser(reader)
			if err != nil {
				dialog.ShowError(err, view.w)
				return
			}

			if pdf.password.Password == "" {
				dialog.ShowInformation("パスワードがない会社", fmt.Sprintf("%sに紐づくパスワードが設定されていないためスキップします", pdf.fileName()), view.w)
				continue
			}

			var buf bytes.Buffer
			if err := view.pdfManagerCore.ProtectPDF(readSeeker, &buf, pdf.password.Password, pdf.password.Password); err != nil {
				dialog.ShowError(err, view.w)
				return
			}

			protectedPDFs = append(protectedPDFs, protectedPDF{
				fileName: pdf.fileName(),
				data:     buf,
			})

			// zipにする
			zipBuf, err := createZipArchive(protectedPDFs)
			if err != nil {
				dialog.ShowError(err, view.w)
				return
			}
			view.showProtectedPDFsZipSaveDialog(zipBuf)
		}

	})
	button.Hide()
	return button
}

func (view *protectPDFView) showProtectedPDFsZipSaveDialog(zipBuf *bytes.Buffer) {
	saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, view.w)
			return
		}
		if writer == nil {
			// ユーザーがcancel
			return
		}
		defer writer.Close()

		// writer zip
		if _, err := writer.Write(zipBuf.Bytes()); err != nil {
			dialog.ShowError(err, view.w)
			return
		}
	}, view.w)

	saveDialog.SetFileName("protected.zip")
	saveDialog.SetFilter(storage.NewExtensionFileFilter([]string{".zip"}))
	saveDialog.Show()
}

func readSeekerFromReadCloser(rc io.ReadCloser) (io.ReadSeeker, error) {
	defer rc.Close()

	// データを全て読み込み
	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

// makeTableUI ロックをかける前の確認tableを作成する
func (view *protectPDFView) makeTableUI() *widget.Table {
	table := widget.NewTableWithHeaders(
		func() (rows int, cols int) {
			return len(view.pdfs), 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			if id.Row < len(view.pdfs) {

				p := view.pdfs[id.Row]

				switch id.Col {
				case 0: // ファイル名
					// label.SetText(view.pdfs[id.Row].fileName())
					label.SetText(p.fileName())
				case 1: // ヒットした会社名
					label.SetText(p.password.TargetName)
				case 2: // パスワード
					label.SetText(p.password.Password)
				}
			}
		},
	)

	table.SetColumnWidth(0, 300)
	table.SetColumnWidth(1, 150)
	table.SetColumnWidth(2, 120)

	return table
}

type pdfs []pdf

func newPDFs(pdfFiles []fyne.URI, passwords model.Passwords) pdfs {

	companyNamePasswordMap := map[string]model.Password{}
	for _, password := range passwords {
		companyNamePasswordMap[password.TargetName] = password
	}

	pdfs := make(pdfs, 0, len(pdfFiles))
	for _, pdfFile := range pdfFiles {

		companyName, err := extractCompanyName(pdfFile.Name())
		if err != nil {
			fmt.Println("会社名が抽出できませんでした", err)
			continue
		}

		pdfs = append(pdfs, pdf{
			pdfFile:  pdfFile,
			password: companyNamePasswordMap[companyName],
		})
	}
	return pdfs
}

type pdf struct {
	pdfFile  fyne.URI
	password model.Password
}

func extractCompanyName(fileName string) (string, error) {
	// 1. 拡張子 .pdf を除外する
	fname := strings.TrimSuffix(fileName, ".pdf")

	// 2. 最初の "_" の手前を除外する
	//    分割は2つだけにしておけば、前半(インデックス0)は用済みなので捨てられる
	parts := strings.SplitN(fname, "_", 2)
	if len(parts) < 2 {
		// "_" が無い場合はそのまま返すなど、適宜エラーハンドリング
		// return fname
		return "", fmt.Errorf("failed extarct company name. fileName: %s", fileName)
	}
	afterUnderscore := parts[1]

	// 3. 「御中」から先を除外する
	//    「御中」が無いケースがあり得る場合は、分割後のチェックが必要
	companyParts := strings.SplitN(afterUnderscore, "御中", 2)
	// 先頭要素が会社名部分
	companyName := companyParts[0]

	// 余分な空白などを除去
	return strings.TrimSpace(companyName), nil
}

func (p *pdf) fileName() string {
	return p.pdfFile.Name()
}

type protectedPDF struct {
	fileName string
	data     bytes.Buffer
}

func createZipArchive(pdfs []protectedPDF) (*bytes.Buffer, error) {
	zipBuf := &bytes.Buffer{}

	zipWriter := zip.NewWriter(zipBuf)

	for _, pdf := range pdfs {
		zipFile, err := zipWriter.Create(pdf.fileName)
		if err != nil {
			return nil, err
		}

		if _, err := io.Copy(zipFile, &pdf.data); err != nil {
			return nil, err
		}
	}

	if err := zipWriter.Close(); err != nil {
		return nil, err
	}
	return zipBuf, nil
}
