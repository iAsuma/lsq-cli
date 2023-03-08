package utils

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
	"regexp"
)

var (
	// gofmtPath is the binary path of command `gofmt`.
	gofmtPath = gproc.SearchBinaryPath("gofmt")
)

// GoFmt formats the source file using command `gofmt -w -s PATH`.
func GoFmt(path string) {
	if gofmtPath != "" {
		gproc.ShellExec(gctx.New(), fmt.Sprintf(`%s -w -s %s`, gofmtPath, path))
	}
}

// IsOrderNo 判断是否订单号
func IsOrderNo(str string) bool {
	regRuler := `^[2-9jJ]\d{15,19}$`

	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(str)
}
