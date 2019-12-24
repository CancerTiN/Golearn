package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file1Path := "D:/Workspace/Golearn/src/Ch14/demo00/abc.txt"
	file2Path := "D:/Workspace/Golearn/src/Ch14/demo00/kkk.txt"

	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println("Read file error:")
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Println("Write file error:")
		fmt.Println(err)
		return
	}
}
