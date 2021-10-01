package library

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"strconv"
	"strings"
)

const (
	kBlack = iota + 30
	kRed
	kGreen
	kYellow
	kBlue
	kPurple
	kCyan
	kWhite
)

func Black(str string) string {
	return KColor(kBlack, str)
}

func Red(str string) string {
	return KColor(kRed, str)
}
func Yellow(str string) string {
	return KColor(kYellow, str)
}
func Green(str string) string {
	return KColor(kGreen, str)
}
func Cyan(str string) string {
	return KColor(kCyan, str)
}
func Blue(str string) string {
	return KColor(kBlue, str)
}
func Purple(str string) string {
	return KColor(kPurple, str)
}
func White(str string) string {
	return KColor(kWhite, str)
}

func KColor(color int, str string) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, str)
}

func FilterResource(resourceName, inputName string) bool {
	if strings.Contains(resourceName, inputName) {
		return true
	}
	return false
}

// generate graph
func CmdGraph(content []string) {
	fmt.Println("#########################################################")
	for i := 0; i < len(content); i++ {
		si := strconv.Itoa(i)
		fmt.Println("    " + Red(si) + "   :   " + Green(content[i]))
		fmt.Println("--------------------------------------------------")
	}
	fmt.Println("#########################################################")
}

func CmdMap(content map[string]string) {
	fmt.Println("############################################################################")
	for k, v := range content {
		fmt.Println("    " + Red(k) + "   :   " + Green(v))
		fmt.Println("-----------------------------------------------------------------------")
	}
	fmt.Println("############################################################################")
}

func CmdEnvs(podName string, envs []v1.EnvVar) {
	fmt.Println("############################################################################")
	fmt.Println("############################################################################")
	fmt.Println(Blue("(容器名称)container-name") + "   :   " + Yellow(podName))
	fmt.Println("-----------------------------------------------------------------------")
	for _, v := range envs {
		fmt.Println("    " + Red(v.Name) + "   :   " + Green(v.Value))
	}
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("############################################################################")
	fmt.Println("############################################################################")
}
