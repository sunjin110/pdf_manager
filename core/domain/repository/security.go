package repository

import "io"

type Security interface {
	// ProtectPDF pdfをロックする
	// Err: ErrPDFIsAlreadLocked
	ProtectPDF(input io.ReadSeeker, output io.Writer, userPassword string, ownerPassword string) error
}
