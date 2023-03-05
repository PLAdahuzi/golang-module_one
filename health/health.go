package health

import (
	"github.com/golang/glog"
	"io"
	"net/http"
)

func Healths(w http.ResponseWriter, req *http.Request) {
	defer glog.Flush()
	writeString, err := io.WriteString(w, "ok")
	if err != nil {
		return
	}
	glog.Infof("返回的内容：%s", writeString)
}
