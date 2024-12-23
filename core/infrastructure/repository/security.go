package repository

import (
	"fmt"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"

	"github.com/sunjin110/pdf_manager/core/domain/repository"
)

type securityRepo struct {
}

func NewSecurity() repository.Security {
	return &securityRepo{}
}

func (s *securityRepo) ProtectPDF(input io.ReadSeeker, output io.Writer, userPassword string, ownerPassword string) error {
	// 暗号化の設定を作成
	// 第1引数: OwnerPassword
	// 第2引数: UserPassword
	// 第3引数: 鍵長(128, 256など)
	conf := model.NewAESConfiguration(
		ownerPassword, // PDFの管理者用パスワード
		userPassword,  // PDFを開くためのユーザパスワード
		256,           // AESのビット長 (128 または 256)
	)

	// パスワードを設定して暗号化
	if err := api.Encrypt(input, output, conf); err != nil {
		return fmt.Errorf("failed encrypt. err: %w", err)
	}

	return nil
}
