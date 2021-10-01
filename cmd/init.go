package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/os/gfile"
	"github.com/ktools/library"
)

func CmdInit() {
	cJson := library.Config{
		CurrentEnv: "dev",
		Env:        nil,
	}
	fc, _ := gfile.Create(library.GetConfigPath())
	cc, _ := json.Marshal(cJson)
	fc.Write(cc)
	fmt.Println("ktools init success. please set your config in path: ~/ktools")
}
