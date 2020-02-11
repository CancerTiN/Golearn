package main

import (
	"fmt"
	"strings"
)

func SingleValue(s string) (v int) {
	if len(s) == 0 {
		return
	}
	m := map[int32]int{'I': 1, 'V': 5, 'X': 10}
	for i, b := range s {
		if i == 0 && b == 'I' && s[len(s)-1] != 'I' {
			m['I'] = -1
		}
		v += m[b]
	}
	return
}

func SplitDigit(s string) (si, te, hu, th string) {
	m := map[int32]int{'M': -1, 'D': -1, 'C': -1, 'L': -1, 'X': -1, 'V': -1, 'I': -1}
	for i, b := range s {
		if m[b] == -1 {
			m[b] = i
		}
	}
	return
}

func myRomanToInt(s string) int {
	sum := 0
	s = strings.Replace(s, "IV", "O", -1)
	s = strings.Replace(s, "IX", "P", -1)
	s = strings.Replace(s, "XL", "Q", -1)
	s = strings.Replace(s, "XC", "R", -1)
	s = strings.Replace(s, "CD", "S", -1)
	s = strings.Replace(s, "CM", "T", -1)
	m := map[int32]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
		'O': 4,
		'P': 9,
		'Q': 40,
		'R': 90,
		'S': 400,
		'T': 900,
	}
	for _, b := range s {
		sum += m[b]
	}
	return sum
}

func romanToInt(s string) int {
	sum := 0
	m := map[string]int{
		"I":  1,
		"IV": 3,
		"V":  5,
		"IX": 8,
		"X":  10,
		"XL": 30,
		"L":  50,
		"XC": 80,
		"C":  100,
		"CD": 300,
		"D":  500,
		"CM": 800,
		"M":  1000,
	}
	for i, _ := range s {
		var (
			v  int
			ok bool
		)
		if i > 0 {
			v, ok = m[s[i-1:i+1]]
			if !ok {
				v = m[s[i:i+1]]
			}
		} else {
			v = m[s[i:i+1]]
		}
		sum += v
	}
	return sum
}

func main() {
	slice := []string{"III", "IV", "IX", "LVIII", "MCMXCIV"}
	for _, s := range slice {
		fmt.Println(s, romanToInt(s))
	}
}
