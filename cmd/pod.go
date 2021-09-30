package cmd

import (
	"fmt"
	"github.com/ktools/library"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
	"strings"
)

func PodCmd() {
	var (
		ns string
	)
	fmt.Println("\r\n")
	fmt.Printf("please input namespace(请输入目标命名空间): ")
	fmt.Scan(&ns)
	nsList := []string{}
	if ns != "" {
		nsList = Namespace(ns)
	}
	if len(nsList) == 0 {
		fmt.Println("namespace list is empty.")
		return
	}
	var selectNs int
	fmt.Printf("please select namespace(请选择命名空间): ")
	fmt.Scan(&selectNs)
	fmt.Println("u choose namespace(您当前选择的命名空间): ", library.Green(nsList[selectNs]))
	//var resourceNameList []string
	fmt.Printf("please input pod (请输入pod): ")
	var podName string
	fmt.Scan(&podName)
	podList := Pod(nsList[selectNs], podName)
	fmt.Printf("please select which resource u want (请选择你的目标资源): ")
	var selectPod int
	fmt.Scan(&selectPod)
	cmdList := CmdList(deploy)
	fmt.Printf("please input cmd (请输入命令行id): ")
	var cmdId int
	fmt.Scan(&cmdId)
	switch cmdList[cmdId] {
	case cmdDesc:
		podDescribe(podList[selectPod], nsList[selectNs])
	}

}

// pod describe
func podDescribe(podName, nsName string) {
	kubeClient := library.GetClient()
	podDesc, _ := kubeClient.CoreV1().Pods(nsName).Get(podName, v1.GetOptions{})
	descMap := map[string]string{}
	descMap["namespace-name"] = nsName
	descMap["pod-name"] = podDesc.Name
	descMap["node-name"] = podDesc.Spec.NodeName
	descMap["pod-ip"] = podDesc.Status.PodIP
	containers := podDesc.Spec.Containers
	for k, v := range containers {
		descMap["container-name-"+strconv.Itoa(k)] = v.Name
		descMap["container-image-"+strconv.Itoa(k)] = v.Image
		descMap["container-imagePullPolicy-"+strconv.Itoa(k)] = string(v.ImagePullPolicy)
		descMap["container-cmd-"+strconv.Itoa(k)] = strings.Join(v.Command, "|")
		//descMap["container-env-"+strconv.Itoa(k)] = strings.Join(v.Env, "|")
	}

	library.CmdMap(descMap)
}
