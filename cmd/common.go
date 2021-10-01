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
	cmdLog  = "log(日志)"
	cmdDesc = "describe(详情)"
	cmdEnv  = "env(环境变量)"
)

// global content print
func GlobalFmt(config *library.Config) {
	currentEnv := config.CurrentEnv
	fmt.Println("currentEnv(当前环境):  ", library.Yellow(currentEnv))
}

// client-go  + regular
func Deployment(ns, deploy string) []string {
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
	return resDeploy
}

func Pod(ns, pod string) []string {
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
	return resPod
}
