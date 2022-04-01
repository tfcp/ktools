package main

import (
	"flag"
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
	// pod cmd
	pod = "pod"
	// ktools init
	initCmd = "init"
)

var (
	KConfig *library.Config
)

func init() {

}

func main() {
	//flag.StringVar(&ns, "ns", "default", "input your deploy name")
	//flag.StringVar(&resourceName, "deploy", "default", "input your deploy name")

	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("please sure your input correct...")
	//		fmt.Println("\r\n")
	//		cmd.Help()
	//	}
	//}()

	flag.Parse()

	if len(flag.Args()) > 0 {
		args := flag.Args()
		if args[0] == initCmd {
			cmd.CmdInit()
			return
		}
		KConfig = library.ReadFromJson(library.GetConfigPath())
		cmd.GlobalFmt(KConfig)
		switch args[0] {
		case initCmd:

		case env:
			cmd.Env(KConfig)
		case sw:
			envName := flag.Args()[1]
			cmd.Switch(envName, KConfig)
		case help:
			cmd.Help()
		case pip:
			cmd.Pip()
		case pod:
			cmd.PodCmd()
		}
	}

}
