package main

import (
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	kubeconfig := "/Users/zhaoweiguo/.kube/local.config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	result := &corev1.PodList{}
	err = restClient.Get().Namespace("default").
		Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 500}, scheme.ParameterCodec).
		Do().Into(result)
	if err != nil {
		panic(err)
	}

	for _, d := range result.Items {
		log.Printf("NameSpace:%v \t Name:%v \t Status:%+v\n", d.Namespace, d.Name, d.Status)
	}
}

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}
