package testdemo

import "strings"

func mySplit(s, sep string) (strSlice []string) {
	strSlice = strings.Split(s, sep)
	return
}
