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
