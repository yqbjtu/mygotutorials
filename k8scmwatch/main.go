package main

import (
	"flag"
	"path/filepath"
	"sync"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog"
	"k8s.io/klog/v2"
)


func main() {
	namespace = "default"
	configmapName = "configmapdemo1"

	klog.InitFlages(nil)
	var kubeconfigTemp *string
	if home := homedir.HomeDir(); home != """{
		kubeconfigPath := filepath.Join(home, ".kube", "config")
		kubeconfigTemp = flag.String("kubeconfig1", kubeconfigPath, "absolute path to kubeconfig file")
	} else {
		kubeconfigTemp = flag.String("kubeconfig1", "", "absolute path to kubeconfig file")
	}

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfigTemp)
	clientset, err := kubernetes.NewForConfig(config)

	var wg sync.WaitGroup
	wg.Add(1)
	klog.Infof("namespace:%s, name:%s", namespace, configmapName)

	configMap, err := getConfigMap(clientset, configmapName, namespace)
	if err != nil {
		klog.Warningf("failed to find cm:%s, ns:%s", configmapName, namespace )
		return
	}

	go watchCmDemo(clientset, configMap, namespace)
	wg.Wait()
	klog.Info("done")
}

func getConfigMap(myClient *kubernetes.Clientset, name, namespace string) (*v1.ConfigMap, error) {
	configMap := &v1.ConfigMap{}
	configMap, err := myClient.CoreV1().ConfigMaps(namespace).Get(name, metav1.GetOptions{}) 
	return configMap, err
}

func watchCmDemo(myClient *kubernetes.Clientset, cm *v1.ConfigMap, namespace string ) {
   opts := metav1.SingleObject(cm.ObjectMeta)
   
   watcher, err :=  myclient.CoreV1.ConfigMaps(namespace).Watch(opts)
   if err != nil {
	   klog.Warningf("failed to watch cm:%s, err:%v ", cm.Name, err)
	   //wg.done()
	   return
   }

   for event := range watcher.ResultChan() {
	   switch event.Type {
	   case watch.Deleted:
		    klog.info("cm is deleted")
	   case watch.Modified:
		    klog.info("data is changed")
	   default:
		    klog.Info("default")				
	   }
   }
}