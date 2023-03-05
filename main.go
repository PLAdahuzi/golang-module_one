package main

import (
	"github.com/golang/glog"
	"module_one/config/register"
	"module_one/header"
	"module_one/health"
	"module_one/ip"
	"module_one/log"
	"module_one/version"
	"net/http"
)

func main() {
	defer glog.Flush()
	http.HandleFunc("/healthz", health.Healths)
	http.HandleFunc("/header", header.CoverHeader)
	http.HandleFunc("/version", version.GetVersion)
	http.HandleFunc("/ip", ip.Ip)
	log.Log()
	glog.Info("################################ 服务启动 ################################")
	server.Service()
}
