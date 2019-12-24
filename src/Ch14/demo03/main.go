package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "D:/Workspace/Golearn/src/Ch14/demo00/python.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Read file error:")
		fmt.Println(err)
	}
	fmt.Printf("The type of content is %T.\n", content)
	fmt.Printf("content = %v\n", content)
	for i, v := range content {
		fmt.Printf("index: %v; uint8: %v; char: %c\n", i, v, v)
	}
	contentString := string(content)
	fmt.Printf("content string = %v", contentString)
}
