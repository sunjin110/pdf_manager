# fyne

## macでwindowsのappをbuildする方法


```sh
# windows用のクロスコンパイラをinstall
brew install mingw-w64

# Windows用のクロスコンパイルに必要な環境をセット
export CGO_ENABLED=1
export GOOS=windows
export GOARCH=amd64
export CC=x86_64-w64-mingw32-gcc

# その上で fyne package
fyne package -os windows -icon Icon.png -name "PDFManager"
```
