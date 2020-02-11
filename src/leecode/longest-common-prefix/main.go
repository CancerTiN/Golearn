package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	var (
		lcp    string
		minLen int
	)
	for i, str := range strs {
		if i == 0 {
			minLen = len(str)
			continue
		}
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	for i := 0; i < minLen; i++ {
		var b byte
		for j, str := range strs {
			if j == 0 {
				b = str[i]
				continue
			}
			if str[i] != b {
				lcp = str[:i]
				return lcp
			}
		}
		lcp += string(b)
	}
	return lcp
}

func main() {
	var strs []string
	strs = []string{"flower", "flow", "flight"}
	fmt.Println(strs, longestCommonPrefix(strs))
	strs = []string{"dog", "racecar", "car"}
	fmt.Println(strs, longestCommonPrefix(strs))
	strs = []string{"c", "c"}
	fmt.Println(strs, longestCommonPrefix(strs))
}
