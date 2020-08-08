package main

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	kubeconfig := "/Users/zhaoweiguo/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	gvr := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
	unstructObj, err := dynamicClient.Resource(gvr).Namespace("default").List(metav1.ListOptions{Limit: 500})
	if err != nil {
		panic(err)
	}
	podlist := &apiv1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructObj.UnstructuredContent(), podlist)
	if err != nil {
		panic(err)
	}
	for _, v := range podlist.Items {
		log.Printf("namespaces:%v  name:%v status:%v \n", v.Namespace, v.Name, v.Status.Phase)
	}
}

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}
