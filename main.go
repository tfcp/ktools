package main

import (
	"flag"
	"fmt"
	"github.com/ktools/cmd"
	"github.com/ktools/library"
)

const (
	// show ktools
	env = "env"
	// ktools switch target env
	sw = "switch"
	// ktools help
	help = "help"
	// pipline to use ktools
	pip = "pip"
)

var (
	KConfig *library.Config
)

func init() {
	KConfig = library.ReadFromJson(library.GetConfigPath())
	cmd.GlobalFmt(KConfig)

}

func main() {
	var (
		ns string
	)
	//flag.StringVar(&ns, "ns", "default", "input your deploy name")
	//flag.StringVar(&resourceName, "deploy", "default", "input your deploy name")

	flag.Parse()
	if len(flag.Args()) > 0 {
		switch flag.Args()[0] {
		case env:
			cmd.Env(KConfig)
		case sw:
			envName := flag.Args()[1]
			cmd.Switch(envName, KConfig)
		case help:
			cmd.Help()
		case start:
			fmt.Println("\r\n")
			fmt.Printf("please input namespace(请输入目标命名空间): ")
			fmt.Scan(&ns)
			nsList := []string{}
			if ns != "" {
				nsList = cmd.Namespace(ns)
			}
			var selectNs int
			fmt.Printf("please select namespace(请选择命名空间): ")
			fmt.Scan(&selectNs)
			fmt.Println(nsList[selectNs])
			fmt.Printf("please select u want resource(请选择你想查看的资源): ")

		}
	}

}
