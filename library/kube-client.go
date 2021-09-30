package library

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"sync"
)

var (
	clientSet *kubernetes.Clientset
	client    *sync.Pool
)

func init() {
	client = &sync.Pool{
		New: func() interface{} {
			// 应用配置初始化
			KConfig := ReadFromJson(GetConfigPath()).Env.(map[string]interface{})
			currentEnv := ReadFromJson(GetConfigPath()).CurrentEnv
			k8sconfig := KConfig[currentEnv].(map[string]interface{})
			config, err := clientcmd.BuildConfigFromFlags("", k8sconfig["path"].(string))
			if err != nil {
				panic(fmt.Sprintf("k8s config path is error:%s", err))
			}

			// 根据指定的 config 创建一个新的 clientset
			clientSet, err = kubernetes.NewForConfig(config)
			if err != nil {
				panic(fmt.Sprintf("k8s NewForConfig is error:%s", err))
			}
			return clientSet
		},
	}

}

func GetClient() *kubernetes.Clientset {
	cs := client.Get().(*kubernetes.Clientset)
	defer client.Put(clientSet)
	return cs
}
