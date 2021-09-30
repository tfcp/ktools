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
	//ns string
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
		case pip:
			cmd.Pip()
		case pod:
			cmd.PodCmd()
		}
	}

}
