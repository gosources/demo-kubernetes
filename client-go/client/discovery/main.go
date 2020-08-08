package main

import (
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	kubeconfig := "/Users/zhaoweiguo/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		log.Panic(err)
	}

	groups, apiResourceLists, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		log.Panic(err)
	}
	for _, group := range groups {
		log.Println(group)
	}
	for _, apiResourceList := range apiResourceLists {
		kind := apiResourceList.Kind
		apiVersion := apiResourceList.APIVersion
		apiResources := apiResourceList.APIResources
		groupVersion := apiResourceList.GroupVersion
		log.Println(kind)
		log.Println(apiVersion)
		log.Println(apiResources)
		log.Println(groupVersion)
		log.Println("=====================")
	}
}

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}
