package cmd

import (
	"fmt"
	"github.com/ktools/library"
)

// show environment list
func Env(config *library.Config) {
	env := config.Env
	envNameList := []string{}
	for k, v := range env.(map[string]interface{}) {
		vm := v.(map[string]interface{})
		envData := k + ":" + vm["name"].(string)
		envNameList = append(envNameList, envData)
	}
	fmt.Println("*** config environment list(环境配置列表) ***")
	library.CmdGraph(envNameList)
}
