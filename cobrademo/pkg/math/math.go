package math

import "k8s.io/klog"

//just demo, not process overflow
func Add(a, b int) int {
	klog.Infof("a:%d, b:%d", a, b)
	return a + b
}
