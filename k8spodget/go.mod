module k8spodget

go 1.13

require (
	k8s.io/api v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v0.19.2
	k8s.io/code-generator v0.19.2
	k8s.io/klog v1.0.0
	k8s.io/kubelet v0.19.2
)

replace k8s.io/client-go => k8s.io/client-go v0.19.2
