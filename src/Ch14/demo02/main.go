package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("D:/Workspace/Golearn/src/Ch14/demo00/python.txt")
	if err != nil {
		fmt.Println("Open file error:")
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Println("file =", file)
	reader := bufio.NewReader(file)
	for {
		reader.ReadString('\n')
	}
}
