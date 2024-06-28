package initialize

import (
	"path/filepath"

	lichee "github.com/issueye/lichee-js"
	"github.com/issueye/pitaya_admin/internal/global"
)

func InitLichee() {
	if global.Lichee == nil {
		path := filepath.Join("runtime", "js")
		global.Lichee = lichee.NewCore(lichee.OptionLog(path, global.Logger.Named("lichee")))
	}
}
