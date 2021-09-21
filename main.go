package main

import (
	"flag"
	"github.com/ktools/cmd"
	"github.com/ktools/library"
)

const (
	env  = "env"
	sw   = "switch"
	help = "help"
)

var (
	KConfig *library.Config
)

func init() {
	KConfig = library.ReadFromJson(library.GetConfigPath())
	cmd.GlobalFmt(KConfig)
}

func main() {
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
		}
	}
}
