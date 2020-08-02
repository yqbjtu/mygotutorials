package main

import (
	"errors"
	"flag"
	"k8s.io/klog"
	"sync"
	"time"
)

/*
  模拟cyclicBarrier 三个routine都等待第一步完成之后（wg1 wait）完成后开始执行第二部分工作
*/
func method1(wg, wg2 *sync.WaitGroup) (string, error) {
	logMsg := "method1"
	klog.Info(logMsg)

	//return logMsg, nil
	go func(wgInner, wg2 *sync.WaitGroup, str string) {
		klog.Infof("inner %s starts. wg:%v", str, wgInner)
		wgInner.Wait()
		klog.Infof("inner %s done", str)
		wg2.Done()
	}(wg, wg2, logMsg)

	time.Sleep(4 * time.Second)
	wg.Done()
	klog.Infof("%s starts to wait. wg:%v", logMsg, wg)
	klog.Infof("%s done", logMsg)
	return logMsg, errors.New(logMsg)
}

func method2(wg, wg2 *sync.WaitGroup) (string, error) {
	logMsg := "method2"
	klog.Info(logMsg)

	go func(wgInner, wg2 *sync.WaitGroup, str string) {
		klog.Infof("inner %s starts. wg:%v", str, wgInner)
		wgInner.Wait()
		klog.Infof("inner %s done", str)
		wg2.Done()
	}(wg, wg2, logMsg)

	time.Sleep(1 * time.Second)
	wg.Done()
	klog.Infof("%s starts to wait. wg:%v", logMsg, wg)

	//return logMsg, nil
	klog.Infof("%s done", logMsg)
	return logMsg, errors.New(logMsg)
}

func method3(wg, wg2 *sync.WaitGroup) (string, error) {
	logMsg := "method3"
	klog.Info(logMsg)

	go func(wgInner, wg2 *sync.WaitGroup, str string) {
		klog.Infof("inner %s starts. wg:%v", str, wgInner)
		wgInner.Wait()
		klog.Infof("inner %s done", str)
		wg2.Done()
	}(wg, wg2, logMsg)
	time.Sleep(2 * time.Second)
	wg.Done()
	klog.Infof("%s starts to wait. wg:%v", logMsg, wg)

	klog.Infof("%s done", logMsg)
	return logMsg, nil
	//return logMsg, errors.New(logMsg)
}

func main() {
	klog.InitFlags(nil)
	flag.Set("log_file", "C:\\F\\myfilewaitgroup.log")
	flag.Parse()
	klog.Info("nice to meet you")

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	wg1.Add(3)
	wg2.Add(3)

	go method1(&wg1, &wg2)
	go method2(&wg1, &wg2)
	go method3(&wg1, &wg2)

	wg1.Wait()
	klog.Info("main1 done")

	wg2.Wait()
	klog.Info("main done")
}
