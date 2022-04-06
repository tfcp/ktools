package informer

import (
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"time"
)

// informer demo
func main() {
	clientset, _ := getClient()
	stopper := make(chan struct{})
	defer close(stopper)

	// 初始化 informer
	factory := informers.NewSharedInformerFactory(clientset, 5*time.Minute)
	nodeInformer := factory.Core().V1().Nodes()
	informer := nodeInformer.Informer()
	defer runtime.HandleCrash()

	// 启动 informer，list & watch
	go factory.Start(stopper)

	// 从 apiserver 同步资源，即 list
	if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	// 使用自定义 handler
	informer.AddEventHandler(NodeInformerHandler{})

	// 创建 lister
	nodeLister := nodeInformer.Lister()
	// 从 lister 中获取所有 items
	nodeList, err := nodeLister.List(labels.Everything())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("nodelist:", nodeList)
	<-stopper
}

// 初始化配置
func getClient() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/zhaosuji/.kube/config")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return clientset, nil
}

type NodeInformerHandler struct {

}

func (node NodeInformerHandler) OnAdd(obj interface{}){

}

func (node NodeInformerHandler) OnDelete(obj interface{}){

}

func (node NodeInformerHandler) OnUpdate(obj1,obj2 interface{}){

}
