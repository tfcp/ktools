package library

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	clientSet *kubernetes.Clientset
)

func init() {
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
}

func GetClient() *kubernetes.Clientset {
	return clientSet
}
