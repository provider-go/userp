package autocreate

import (
	"github.com/gohouse/converter"
	"testing"
)

func Test_Auto(t *testing.T) {
	converter.NewTable2Struct().
		SavePath("../model.go").
		Dsn("root:ZTQ4ZTBjMTViNGMzODgzODUz@tcp(120.53.243.73:13306)/services?charset=utf8").
		Run()
}
