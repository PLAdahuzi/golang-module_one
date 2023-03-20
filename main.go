package main

import (
	"github.com/golang/glog"
	"golang-module_one/config/register"
	"golang-module_one/header"
	"golang-module_one/health"
	"golang-module_one/ip"
	"golang-module_one/log"
	"golang-module_one/version"
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
