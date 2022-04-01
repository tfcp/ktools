package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)


func main() {

	config, err := clientcmd.BuildConfigFromFlags("", "./config/config")
	if err != nil {
		log.Fatal(err)
		return
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	//sharedInformerFac := informers.NewSharedInformerFactory(clientset,time.Minute*10)

	//nodeInformer := sharedInformerFac.Core().V1().Nodes()
	//informer := nodeInformer.Informer()
	// 初始化 informer
	factory := informers.NewSharedInformerFactory(clientset, 0)
	nodeInformer := factory.Core().V1().Nodes()
	informer := nodeInformer.Informer()
	defer runtime.HandleCrash()
	stopper := make(chan struct{})
	defer close(stopper)
	// 启动 informer，list & watch
	go factory.Start(stopper)

	// 从 apiserver 同步资源，必不可少
	if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	// 使用自定义 handler
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		UpdateFunc: func(interface{}, interface{}) { fmt.Println("update not implemented") }, // 此处省略 workqueue 的使用
		DeleteFunc: func(interface{}) { fmt.Println("delete not implemented") },
	})

	// 创建 lister
	nodeLister := nodeInformer.Lister()
	// 从 lister 中获取所有 items
	_, err = nodeLister.List(labels.Everything())
	if err != nil {
		fmt.Println(err)
	}
	<-stopper

	select {}
}

func onAdd(obj interface{}) {
	node := obj.(*v1.Node)
	//fmt.Println("add a Node:", node.Name)
	//fmt.Println("add a Cpu:", node.Status.Allocatable.Memory())
	//fmt.Println("add a Mem:", node.Status.Allocatable.Cpu())
	fmt.Println("add a label:", node)
}

// 参考获取Allocated resources https://www.jianshu.com/p/a3b8f1019d10
