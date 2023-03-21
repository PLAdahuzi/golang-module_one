package server

import (
	"github.com/golang/glog"
	"net/http"
)

func Service() {
	defer glog.Flush()
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		glog.Info("################################ 服务启动失败 ################################")
		glog.Info(err)
	}
}
