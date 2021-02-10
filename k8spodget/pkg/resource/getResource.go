package resource

import (
	"context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	"time"

	"k8s.io/client-go/kubernetes"
)

type MyResource struct {
	Clientset *kubernetes.Clientset
}

func (s *MyResource) GetNode() {
	opts := v1.ListOptions{
		Limit: 100,
	}

	nodes, err := s.Clientset.CoreV1().Nodes().List(context.TODO(), opts)
	if err != nil {
		klog.Error("failed to get node list ,err:%v", err)
		return
	}

	for _, node := range nodes.Items {
		klog.Infof("name:%s, Status:%v", node.Name, node.Status.NodeInfo.OSImage)
	}
}

func (s *MyResource) GetPod() {
	opts := v1.ListOptions{
		Limit: 100,
	}

	podwatch, err := s.Clientset.CoreV1().Pods("default").Watch(context.TODO(), opts)
	if err != nil {
		klog.Error("failed to watch pod list, err:%v", err)
		return
	}

	for {
		select {
		case e, ok := <-podwatch.ResultChan():
			if !ok {
				// 说明该通道已经被close掉了
				klog.Warning("podWatch chan has been close!")
				time.Sleep(time.Second * 5)
			}
			if e.Object != nil {
				klog.Infof("chan is ok. type:%v", e.Type)
				klog.Info(e.Object.DeepCopyObject())
			}
		}
	}
}

func (s *MyResource) GetDeployment(ns string) {
	deploy, err := s.Clientset.AppsV1().Deployments(ns).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		klog.Error("failed to get deploy list ,err:%v", err)
		return
	}

	for _, deploy := range deploy.Items {
		klog.Infof("deployName:%s,  replicas:%d, status.UnavailableReplicas:%d,", deploy.Name, *deploy.Spec.Replicas, deploy.Status.UnavailableReplicas)
	}
}
