package main

import (
	"errors"
	"flag"
	"k8s.io/klog"
	"time"
)

func method1() (string, error) {
	logMsg := "method1"
	klog.Info(logMsg)

	time.Sleep(4 * time.Second)
	//return logMsg, nil
	return logMsg, errors.New(logMsg)
}

func method2() (string, error) {
	logMsg := "method2"
	klog.Info(logMsg)

	time.Sleep(1 * time.Second)
	//return logMsg, nil
	return logMsg, errors.New(logMsg)
}

func method3() (string, error) {
	logMsg := "method3"
	klog.Info(logMsg)

	time.Sleep(2 * time.Second)
	return logMsg, nil
	//return logMsg, errors.New(logMsg)
}

func main() {
	klog.InitFlags(nil)
	flag.Set("log_file", "C:\\F\\myfilechan.log")
	flag.Parse()
	klog.Info("info nice to meet you")

	successChan := make(chan string, 3)
	failureChan := make(chan string, 3)
	defer close(successChan)
	defer close(failureChan)
	failureCount := 0
	//var myFuncs [3]func()
	//myFuncs[0] := method1()
	//myFuncs[1] := method2()
	//myFuncs[2] := method3()
	//var myFuncs [3]int
	//myFuncs[0] := 1

	myFuncs := [3]int{1, 2, 3}
	for index, myFunc := range myFuncs {
		klog.Infof("index:%v, value:%v", index, myFunc)
		go func(innerMyFunc int) {
			if innerMyFunc == 1 {
				strResult, err := method1()
				if err == nil {
					successChan <- strResult
				} else {
					failureChan <- strResult
				}
			} else if innerMyFunc == 2 {
				strResult, err := method2()
				if err == nil {
					successChan <- strResult
				} else {
					failureChan <- strResult
				}
			} else if innerMyFunc == 3 {
				strResult, err := method3()
				if err == nil {
					successChan <- strResult
				} else {
					failureChan <- strResult
				}
			}
		}(myFunc)
	}

	for {
		select {
		case successV := <-successChan:
			klog.Infof("successValue:%v", successV)
			return
		case failureV := <-failureChan:
			failureCount++
			klog.Infof("failureValue:%v, failureCount:%d", failureV, failureCount)
			if failureCount >= 3 {
				return
			}
		}
	}

	klog.Info("main done")
}
