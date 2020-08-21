package main

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	kubeconfig := "/Users/zhaoweiguo/.kube/local.config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	stopCh := make(chan struct{})
	defer close(stopCh)

	sharedInformers := informers.NewSharedInformerFactory(clientSet, 0)
	//sharedInformers := informers.NewSharedInformerFactory(clientSet, time.Minute)   // 指定1分钟全量同步一次
	informer := sharedInformers.Core().V1().Pods().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mObj := obj.(metav1.Object)
			log.Printf("New Pod add : %s:%s", mObj.GetNamespace(), mObj.GetName())
			oobj := obj.(*v1.Pod)
			log.Println(oobj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oObj := oldObj.(metav1.Object)
			nObj := newObj.(metav1.Object)
			log.Println(oObj, nObj)
		},
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(metav1.Object)
			log.Printf("Pod Deleted: %s:%s", mObj.GetNamespace(), mObj.GetName())
		},
	})
	informer.Run(stopCh)
}
