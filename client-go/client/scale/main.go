package main

import (
	discocache "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/dynamic"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/scale"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

/*
@todo 没找到更多scale client相关的资料
*/
func main() {
	kubeconfig := "/Users/zhaoweiguo/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	cachedDiscovery := discocache.NewMemCacheClient(client.Discovery())

	restMapper := discovery.NewDeferredDiscoveryRESTMapper(cachedDiscovery, apimeta.InterfacesForUnstructured)

	scaleClient, err := scale.NewForCo在nfig(config, restMapper, dynamic.LegacyAPIPathResolverFunc, scaleKindResolver)

}

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}
