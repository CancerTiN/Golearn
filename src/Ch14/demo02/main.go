package main

import (
	"bufio"
	"fmt"
	"io"
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
	fmt.Println("reader =", reader)
	fmt.Printf("The type of reader is %T.\n", reader)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Read string error:")
			fmt.Println(err)
		}
		fmt.Printf("str = %v", str)
	}
	fmt.Println("File reading finished ...")
}
