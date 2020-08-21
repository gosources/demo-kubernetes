package main

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"strings"
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
			status := oobj.Status
			//spec := oobj.Spec
			//objMeta := oobj.ObjectMeta

			log.Println(status.Message)
			//log.Println(status.HostIP)
			//log.Println(status.PodIP)
			//log.Println(status.Reason)
			//
			//conditions := status.Conditions
			//for _, condition := range conditions {
			//	log.Println(condition.Reason)
			//	log.Println(condition.Message)
			//	log.Println(condition.Status)
			//	log.Println(condition.Type)
			//}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oObj := oldObj.(metav1.Object)
			nObj := newObj.(metav1.Object)
			if strings.HasPrefix(oObj.GetName(), "tuya-data-puller-cronjob") {
				return
			}

			log.Printf("%s:%s Pod Updated to %s:%s", oObj.GetNamespace(), oObj.GetName(), nObj.GetNamespace(), nObj.GetName())
			oObj2 := oldObj.(*v1.Pod)
			nObj2 := newObj.(*v1.Pod)

			// Spec
			oSpec := oObj2.Spec
			nSpec := nObj2.Spec
			log.Println(oSpec.ActiveDeadlineSeconds, nSpec.ActiveDeadlineSeconds)
			log.Println(oSpec.Size(), nSpec.Size())
			log.Println(oSpec.String(), nSpec.String())

			// Status
			//log.Println(oObj2.Status.PodIP, nObj2.Status.PodIP)
			//log.Println(oObj2.Status.Reason, nObj2.Status.Reason)
			//log.Println(oObj2.Status.Message, nObj2.Status.Message)
			//log.Println(oObj2.Status.StartTime, nObj2.Status.StartTime)
			//
			//for _, condition := range oObj2.Status.Conditions {
			//	log.Println(condition.Status, condition.Type, condition.Message, condition.Reason, condition.LastProbeTime, condition.LastTransitionTime)
			//
			//}
			//log.Println("===========")
			//for _, condition := range nObj2.Status.Conditions {
			//	log.Println(condition.Status, condition.Type, condition.Message, condition.Reason, condition.LastProbeTime, condition.LastTransitionTime)
			//
			//}

		},
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(metav1.Object)
			log.Printf("Pod Deleted: %s:%s", mObj.GetNamespace(), mObj.GetName())
		},
	})
	informer.Run(stopCh)
}
