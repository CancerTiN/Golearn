package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("D:/Workspace/Golearn/src/Ch14/demo00/test.txt")
	if err != nil {
		fmt.Println("Open file error:")
		fmt.Println(err)
	}
	fmt.Println("file =", file)
	err = file.Close()
	if err != nil {
		fmt.Println("Close file error:")
		fmt.Println(err)
	}
}
