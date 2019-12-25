package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	dstFilePath := `D:\Workspace\Golearn\src\Ch14\demo00\flower.jpg`
	srcFilePath := `D:\Workspace\Golearn\src\Ch14\figure\flower.jpg`
	written, err := CopyFile(dstFilePath, srcFilePath)
	if err != nil {
		fmt.Println("Copy file error:")
		fmt.Println(err)
	} else {
		fmt.Printf("The number of bytes copied is %v.\n", written)
	}
}

func CopyFile(dstFilePath string, srcFilePath string) (written int64, err error) {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Println("Open error:")
		fmt.Println(err)
		return 0, err
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file error")
		fmt.Println(err)
		return 0, err
	}
	defer dstFile.Close()

	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(dstFile)
	return io.Copy(writer, reader)
}
