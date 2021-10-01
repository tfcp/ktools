package cmd

import (
	"fmt"
	"github.com/ktools/library"
)

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
	var resourceNameList []string
	switch resourceList[selectResource] {
	case deploy:
		fmt.Printf("please input deploy (请输入deployment): ")
		var deployName string
		fmt.Scan(&deployName)
		resourceNameList = Deployment(nsList[selectNs], deployName)
		//resourceNameList = Deployment(nsList[selectNs], deployName)
	case job:
	case pod:
		fmt.Printf("please input pod (请输入pod): ")
		var podName string
		fmt.Scan(&podName)
		Pod(nsList[selectNs], podName)
	}
	fmt.Printf("please select which resource u want (请选择你的目标资源): ")
	var selectCmd int
	fmt.Scan(&selectCmd)
	cmdList := CmdList(resourceList[selectResource])
	fmt.Printf("please input cmd (请输入命令行id): ")
	var cmdId int
	fmt.Scan(&cmdId)
	Cmd(cmdList[cmdId], resourceNameList[selectCmd], resourceList[selectResource])
}
