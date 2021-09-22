package library

import (
	"fmt"
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
