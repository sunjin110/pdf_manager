init:
	go install fyne.io/fyne/v2/cmd/fyne@latest

package_win:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc \
		fyne package -os windows -icon Icon.png -name "PDFManager"

package_mac:
	fyne package -name PDFManager -os darwin -icon icon.png
