package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	name := "D:/Workspace/Golearn/src/Ch14/demo00/abc.txt"
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Open file error:")
		fmt.Println(err)
		return
	}
	defer file.Close()

	str := "ABC! ENGLISH!\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
