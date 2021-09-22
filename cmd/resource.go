package cmd

import (
	"fmt"
	"github.com/ktools/library"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	resourceTimeout = int64(5)
	deploy          = "deployment"
	job             = "job"
	pod             = "pod"
	resourceList    = []string{
		deploy, job, pod,
	}
	cmdLog  = "log"
	cmdDesc = "describe"
	cmdTop  = "top"
	cmdList = []string{
		cmdLog, cmdDesc, cmdTop,
	}
)

func init() {
	//Deployment("d2d-ph")
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

func CmdList() []string {
	library.CmdGraph(cmdList)
	return cmdList
}

func Cmd(cmd, resourceName, resourceType string) {

}

// client-go  + regular
func Deployment(ns, deploy string) {
	client := library.GetClient()
	list := client.AppsV1beta2().Deployments(ns)
	listOp := v1.ListOptions{
		TimeoutSeconds: &resourceTimeout,
	}
	ds, _ := list.List(listOp)
	//resDeploy := make([]v1beta2.Deployment, 0)
	resDeploy := make([]string, 0)
	for _, v := range ds.Items {
		if library.FilterResource(v.Name, deploy) {
			resDeploy = append(resDeploy, v.Name)
		}
	}
	library.CmdGraph(resDeploy)
}

func Pod(ns, pod string) {
	client := library.GetClient()
	listPod := client.CoreV1().Pods(ns)
	listOp := v1.ListOptions{
		TimeoutSeconds: &resourceTimeout,
	}
	pods, _ := listPod.List(listOp)
	resPod := make([]string, 0)
	for _, v := range pods.Items {
		if library.FilterResource(v.Name, pod) {
			resPod = append(resPod, v.Name)
		}
	}
	library.CmdGraph(resPod)
}

func Pip() {
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
	resourceList := Resource()
	fmt.Printf("please select u want resource(请选择你想查看的资源): ")
	var selectResource int
	fmt.Scan(&selectResource)
	switch resourceList[selectResource] {
	case deploy:
		fmt.Printf("please input deploy (请输入deployment): ")
		var deployName string
		fmt.Scan(&deployName)
		Deployment(nsList[selectNs], deployName)
	case job:
	case pod:
		fmt.Printf("please input pod (请输入pod): ")
		var podName string
		fmt.Scan(&podName)
		Pod(nsList[selectNs], podName)
	}
	fmt.Printf("please input cmd (请输入cmd): ")
	var selectCmd int
	fmt.Scan(&selectCmd)
	Cmd()
}
