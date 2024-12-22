package usecase

import (
	"io"

	"github.com/sunjin110/pdf_manager/core/domain/repository"
)

type Security interface {
	ProtectPDF(input io.ReadSeeker, output io.Writer, userPassword string, ownerPassword string) error
}

type security struct {
	securityRepo repository.Security
}

func NewSecurity(securityRepo repository.Security) Security {
	return &security{
		securityRepo: securityRepo,
	}
}

func (s *security) ProtectPDF(input io.ReadSeeker, output io.Writer, userPassword string, ownerPassword string) error {
	return s.securityRepo.ProtectPDF(input, output, userPassword, ownerPassword)
}
