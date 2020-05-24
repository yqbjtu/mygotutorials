

package main


import (
	"github.com/golang/glog"
	"os"
	"flag"
	"fmt"
)

const numOfLine = 10

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARNING|FATAL] -log_dir=[string]\n", )
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line 
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}

func main() {
	
	for i := 0; i < numOfLine; i++ {
		glog.V(2).Infof("v2 LINE: %d", i)
		message := fmt.Sprintf("TEST LINE: %d", i)
		glog.Error(message)
	}
	glog.Info("info level")
	glog.Flush()
}

/*

每次产生了3个文件
glogdemo.exe.yamgqian-PC.yamgqian-PC_yangqian.log.ERROR.20200524-113450.11884 --只有ERROR级别的日志
>go run glogdemo.go -log_dir=. -v=1 是v2的info也没有， 当-v=2以及-v=3， 4等就有
glogdemo.exe.yamgqian-PC.yamgqian-PC_yangqian.log.INFO.20200524-113450.11884   ERROR和INFO以及V2的INFO级别的日志
glogdemo.exe.yamgqian-PC.yamgqian-PC_yangqian.log.WARNING.20200524-113450.11884  只有WARNING级别的日志

*/