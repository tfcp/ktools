package cmd

import (
	"fmt"
	"github.com/ktools/library"
	"k8s.io/api/apps/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
)

var ()

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
	for i := 0; i < len(resNs); i++ {
		si := strconv.Itoa(i)
		fmt.Println(si + ":" + resNs[i])
	}
	return resNs
}

// client-go  + regular
func Deployment(ns, deploy string) {
	to := int64(5)
	client := library.GetClient()
	list := client.AppsV1beta2().Deployments(ns)
	listOp := v1.ListOptions{
		TimeoutSeconds: &to,
	}
	ds, _ := list.List(listOp)
	resDeploy := make([]v1beta2.Deployment, 0)
	for _, v := range ds.Items {
		if library.FilterResource(v.Name, deploy) {
			resDeploy = append(resDeploy, v)
		}
	}
	for _, v := range resDeploy {
		fmt.Println(v.Name)
	}

}
