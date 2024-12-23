package sqlite_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/pdf_manager/core/infrastructure/sqlite"
)

// go test -v -count=1 -timeout 30s -run ^TestMigrate$ github.com/sunjin110/pdf_manager/core/infrastructure/sqlite
func TestMigrate(t *testing.T) {
	Convey("TestMigrate", t, func() {

		err := sqlite.Migrate("testdata/sample.db")
		So(err, ShouldBeNil)
	})
}
