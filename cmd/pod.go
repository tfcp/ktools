package cmd

import (
	"bytes"
	"fmt"
	"github.com/ktools/library"
	"io"
	v12 "k8s.io/api/core/v1"
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
	cmdList := CmdList(pod)
	fmt.Printf("please input cmd (请输入命令行id): ")
	var cmdId int
	fmt.Scan(&cmdId)
	switch cmdList[cmdId] {
	case cmdDesc:
		podDescribe(podList[selectPod], nsList[selectNs])
	case cmdLog:
		podLog(podList[selectPod], nsList[selectNs])
	case cmdEnv:
		podEnv(podList[selectPod], nsList[selectNs])
	}

}

// pod describe
func podDescribe(podName, nsName string) {
	kubeClient := library.GetClient()
	podDesc, _ := kubeClient.CoreV1().Pods(nsName).Get(podName, v1.GetOptions{})
	descMap := map[string]string{}
	descMap["(命名空间)namespace-name"] = nsName
	descMap["(pod名称)pod-name"] = podDesc.Name
	descMap["(节点名称)node-name"] = podDesc.Spec.NodeName
	descMap["(pod的IP)pod-ip"] = podDesc.Status.PodIP
	descMap["(pod的服务质量)pod-QOS"] = string(podDesc.Status.QOSClass)
	containers := podDesc.Spec.Containers
	for k, v := range containers {
		descMap["(容器名称)container-name-"+strconv.Itoa(k)] = v.Name
		descMap["(容器资源cpu限制值)container-limit-cpu-"+strconv.Itoa(k)] = v.Resources.Limits.Cpu().String()
		descMap["(容器资源内存限制值)container-limit-mem-"+strconv.Itoa(k)] = v.Resources.Limits.Memory().String()
		descMap["(容器资源cpu期望值)container-request-cpu-"+strconv.Itoa(k)] = v.Resources.Requests.Cpu().String()
		descMap["(容器资源内存期望值)container-request-mem-"+strconv.Itoa(k)] = v.Resources.Requests.Memory().String()
		descMap["(容器镜像)container-image-"+strconv.Itoa(k)] = v.Image
		descMap["(容器镜像策略)container-imagePullPolicy-"+strconv.Itoa(k)] = string(v.ImagePullPolicy)
		descMap["(容器命令行)container-cmd-"+strconv.Itoa(k)] = strings.Join(v.Command, "|")
	}
	library.CmdMap(descMap)
}

func podLog(podName, nsName string) {
	kubeClient := library.GetClient()
	LogLine := int64(50)
	rc := kubeClient.CoreV1().Pods(nsName).GetLogs(podName, &v12.PodLogOptions{
		//Container
		TailLines: &LogLine,
	})
	podLogs, _ := rc.Stream()
	defer podLogs.Close()
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, podLogs)
	if err != nil {
		//err = errors.Wrap(err, "kubernetes.podLogsToBuf")
		return
	}
	LogStr := buf.String()
	fmt.Println(LogStr)
}

func podTop(podName, nsName string) {

}

func podEnv(podName, nsName string) {
	kubeClient := library.GetClient()
	podDesc, _ := kubeClient.CoreV1().Pods(nsName).Get(podName, v1.GetOptions{})
	containers := podDesc.Spec.Containers
	for _, v := range containers {
		library.CmdEnvs(v.Name, v.Env)
	}
}
