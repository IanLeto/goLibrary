package pprof

import (
	"goLibrary/utils"
	"net/http"
	_ "net/http/pprof"
)

func PprofWeb() {
	utils.NoErr(http.ListenAndServe(":9999", nil))
}
