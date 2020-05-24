package main

import (
	"flag"
	"k8s.io/klog"
)

func main() {
	klog.InitFlags(nil)
	flag.Set("log_file", "C:\\F\\myfile.log")
	flag.Parse()
	klog.Info("info nice to meet you")
	klog.Error("Error nice to meet you")
	klog.Flush()
}