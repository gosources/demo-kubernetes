package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"strings"
)

func UserIndexFunc(obj interface{}) ([]string, error) {
	pod := obj.(*v1.Pod)
	usersString := pod.Annotations["users"]
	return strings.Split(usersString, ","), nil
}

func main() {
	keyFunc := cache.MetaNamespaceKeyFunc
	indexers := cache.Indexers{"byUser": UserIndexFunc}
	index := cache.NewIndexer(keyFunc, indexers)

	pod1 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "one", Annotations: map[string]string{"users": "aaaa,bbbb"}}}
	pod2 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "two", Annotations: map[string]string{"users": "bbbb, cccc"}}}
	pod3 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "tre", Annotations: map[string]string{"users": "aaaa, dddd"}}}

	index.Add(pod1)
	index.Add(pod2)
	index.Add(pod3)
	pods, err := index.ByIndex("byUser", "bbbb")
	if err != nil {
		panic(err)
	}
	for _, pod := range pods {
		fmt.Println(pod.(*v1.Pod).Name)
	}
}
