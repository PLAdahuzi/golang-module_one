package log

import (
	"flag"
	"github.com/golang/glog"
)

func Log() {
	defer glog.Flush()
	//flag.Lookup("logtostderr").Value.Set("true")
	flag.Lookup("log_dir").Value.Set("../../../Log_data")
}
