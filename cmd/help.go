package cmd

import "fmt"

// show ktools help list
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
	helpInfo += "    pod			    to operate our pod resources.  "
	helpInfo += "\r\n"
	helpInfo += "    switch <environment>    switch current environment to target environment.  "
	helpInfo += "\r\n"
	fmt.Println(helpInfo)
}
