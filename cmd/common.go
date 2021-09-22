package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/os/gfile"
	"github.com/ktools/library"
)

// switch env
// init env
func Switch(env string, kconfig *library.Config) {
	//f, err := ioutil.ReadFile("./test.json")
	//if err != nil {
	//	fmt.Println("Openerr:", err)
	//	return
	//}
	//c := &library.Config{}
	//err = json.Unmarshal(f, c)
	//if err != nil {
	//	fmt.Println("jsonerr:", err)
	//	return
	//}
	c := kconfig
	if c.CurrentEnv == env {
		fmt.Println(fmt.Sprintf("switch to %s success.", library.Red(env)))
		return
	}
	// 替换env
	envHas := false
	for k, _ := range c.Env.(map[string]interface{}) {
		if k == env {
			c.CurrentEnv = k
			envHas = true
			break
		}
	}
	if !envHas {
		fmt.Println("env input is not in config.")
		return
	}
	fc, _ := gfile.Create(library.GetConfigPath())
	config := c
	cc, _ := json.Marshal(config)
	fc.Write(cc)
	fmt.Println(fmt.Sprintf("switch to %s success.", library.Red(env)))
}

// global content print
func GlobalFmt(config *library.Config) {
	currentEnv := config.CurrentEnv
	fmt.Println("currentEnv(当前环境):", library.Green(currentEnv))
}

// environment command
func Env(config *library.Config) {
	env := config.Env
	envNameList := []string{}
	for k, v := range env.(map[string]interface{}) {
		vm := v.(map[string]interface{})
		envData := k + ":" + vm["name"].(string)
		envNameList = append(envNameList, envData)
	}
	fmt.Println("envList(环境列表): ", envNameList)
}

func Help() {
	helpInfo := "ktools is a funny local k8s tool. \r\n"
	helpInfo += "\r\n"
	// todo code line show ktools
	helpInfo += ""
	helpInfo += "Usage: \r\n"
	helpInfo += "\r\n"
	helpInfo += "    ktools <command> [arguments]"
	helpInfo += "\r\n"
	helpInfo += "The commands are: \r\n"
	helpInfo += "\r\n"
	helpInfo += "    env                     show current environment, config environment list. "
	helpInfo += "\r\n"
	helpInfo += "    switch <environment>    switch current environment to target environment.  "
	helpInfo += "\r\n"
	fmt.Println(helpInfo)
}
