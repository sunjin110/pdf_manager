package repository_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/pdf_manager/core/infrastructure/repository"
)

func Test_securityRepo_ProtectPDF(t *testing.T) {
	Convey("Test_securityRepo_ProtectPDF", t, func() {

		input, err := os.Open(filepath.Join("testdata", "dummy.pdf"))
		So(err, ShouldBeNil)
		defer input.Close()

		output, err := os.Create(filepath.Join("testdata", "protected_dummy.pdf"))
		So(err, ShouldBeNil)
		defer output.Close()

		securityRepo := repository.NewSecurity()
		err = securityRepo.ProtectPDF(input, output, "pass", "pass")
		So(err, ShouldBeNil)
	})
}
