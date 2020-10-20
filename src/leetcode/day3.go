package main

import (
	"fmt"
	"strings"
)

func deleteBackspace(S string) string {
	temp := ""
	for _, char := range S {
		if string(char) == "#" {
			if len(temp) > 0 {
				temp = temp[:len(temp)-1]
			}
		} else {
			temp = temp + string(char)
		}
	}
	return temp
}

func backspaceCompare(S string, T string) bool {
	s := deleteBackspace(S)
	t := deleteBackspace(T)
	return strings.EqualFold(s, t)
}

func main() {
	result := backspaceCompare("s#s", "s##s")
	fmt.Printf("%v", result)
}
