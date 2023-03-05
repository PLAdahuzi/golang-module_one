package version

import (
	"github.com/golang/glog"
	"net/http"
	"os"
)

func GetVersion(w http.ResponseWriter, req *http.Request) {
	defer glog.Flush()
	getenv := os.Getenv("GOVERSION")
	w.Header().Set("GOVERSION", getenv)
	glog.Infof("获取到的环境变量是： %s", getenv)
}
