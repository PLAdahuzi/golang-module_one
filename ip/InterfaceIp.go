package ip

import (
	"fmt"
	"github.com/golang/glog"
	"net"
	"net/http"
)

func Ip(w http.ResponseWriter, r *http.Request) {
	defer glog.Flush()
	remoteIp := r.RemoteAddr

	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		remoteIp = ip
	} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
		remoteIp = ip
	} else {
		remoteIp, _, _ = net.SplitHostPort(remoteIp)
	}

	//本地ip
	if remoteIp == "::1" {
		remoteIp = "127.0.0.1"
	}
	fmt.Println(remoteIp)
	glog.Infof("获取到的ip地址 %s", remoteIp)

}
