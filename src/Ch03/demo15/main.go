package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Base("a/b/c.go"))
	fmt.Println(filepath.Base("c.d.go"))
	fmt.Println(filepath.Base("abc"))
}
