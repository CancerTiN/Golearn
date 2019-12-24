package main

import (
	"fmt"
	"os"
)

func main() {
	path1 := "D:/Workspace/Golearn/src/Ch14/demo00/abc.txt"
	path2 := "D:/Workspace/Golearn/src/Ch14/demo00/xyz.txt"
	fmt.Println(path1)
	fmt.Println(path2)
	b, err := PathExists(path2)
	fmt.Println(b)
	fmt.Println(err)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}
