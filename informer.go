package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"time"

	"github.com/martin-helmich/kubernetes-crd-example/api/types/v1alpha1"
	client_v1alpha1 "github.com/martin-helmich/kubernetes-crd-example/clientset/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

func WatchResources(clientSet client_v1alpha1.ExampleV1Alpha1Interface) cache.Controller {
	_, projectController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return clientSet.Projects(v1.NamespaceAll).List(lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return clientSet.Projects(v1.NamespaceAll).Watch(lo)
			},
		},
		&v1alpha1.Project{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    AddMiddleWare,
			DeleteFunc: DeleteMiddleWare,
			UpdateFunc: UpdateMiddleWare,
		},
	)

	return projectController
}

func AddMiddleWare(obj interface{}) {
	fmt.Println("add func")
}

func DeleteMiddleWare(obj interface{}){
	fmt.Println("delete func")
}

func UpdateMiddleWare(oldObj, newObj interface{}) {
	fmt.Println("update func")
}