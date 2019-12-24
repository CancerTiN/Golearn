package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	name := "D:/Workspace/Golearn/src/Ch14/demo00/abc.txt"
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Open file error:")
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	i := 0
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Read string error:")
			fmt.Println(err)
		}
		fmt.Printf("line %v = %v", i, str)
		i++
	}

	str := "Hello! Beijing!\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
