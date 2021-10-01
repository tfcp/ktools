package cmd

import (
	"fmt"
	"github.com/ktools/library"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	//kubeClient := library.GetClient()
	//a, _ := kubeClient.CoreV1().Pods("app").Get("med-xim-5b57f67475-rkxhb", v1.GetOptions{})
	////fmt.Println(a.Status.ContainerStatuses[0].Requests.Cpu().String())
	//fmt.Println(a.Spec.Containers[0].Resources.Requests.Cpu().Value())
	//m := a.Spec.Containers[0].Resources.Requests.Memory().Value()
	//m1 := int(m / 1024 / 1024)
	//fmt.Println(strconv.Itoa(m1) + "Mi")
}

func Namespace(ns string) []string {
	client := library.GetClient()
	listNs, _ := client.CoreV1().Namespaces().List(v1.ListOptions{})
	resNs := make([]string, 0)
	for _, nsValue := range listNs.Items {
		if library.FilterResource(nsValue.Name, ns) {
			resNs = append(resNs, nsValue.Name)
		}
	}
	library.CmdGraph(resNs)
	return resNs
}

func Resource() []string {
	library.CmdGraph(resourceList)
	return resourceList
}

func CmdList(resourceType string) []string {
	var cmdList = []string{
		cmdLog, cmdDesc, cmdEnv,
	}
	switch resourceType {
	case deploy, job:
		cmdList = []string{cmdDesc}
	}
	library.CmdGraph(cmdList)
	return cmdList
}

func Cmd(cmdString, resourceName, resourceType string) {
	fmt.Println("cmdId:", cmdString)
	fmt.Println("resourceName:", resourceName)
	switch resourceType {
	case deploy:

	}
}
