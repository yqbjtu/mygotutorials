package main

import (
	"flag"
	"k8s.io/klog"
	"k8spodget/pkg/common"
	"k8spodget/pkg/resource"
)

func main() {
	klog.InitFlags(nil)
	flag.Set("log_file", "C:\\F\\myfilek8spod.log")
	flag.Parse()
	klog.Info("start to get k8s client")
	clientSet, err := common.GetClient()
	if err != nil {
		klog.Warningf("failed to get k8s clientSet, err:%v", err)
	}
	myResource := resource.MyResource{Clientset: clientSet}
	myResource.GetNode()
	myResource.GetDeployment("default")
	//会一直运行
	myResource.GetPod()
	klog.Flush()
}
