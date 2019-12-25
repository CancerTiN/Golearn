package main

func addUpper(n int) int {
	var r int
	for i := 0; i <= n; i++ {
		r += i
	}
	return r
}

func getSub(n1, n2 int) int {
	return n1 - n2
}
