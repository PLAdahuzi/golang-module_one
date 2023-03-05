package header

import (
	"github.com/golang/glog"
	"net/http"
)

func CoverHeader(w http.ResponseWriter, req *http.Request) {
	defer glog.Flush()
	h := req.Header
	for s := range h {
		glog.Infof(s + "=====" + req.Header.Get(s))

		w.Header().Set(s, req.Header.Get(s))
	}
}
