package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var chessMap [11][11]int

	chessMap[1][2] = 1
	chessMap[2][3] = 2

	for _, row := range chessMap {
		for _, v := range row {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}

	filename := "D:/Workspace/Golearn/src/Ch20/sparsearray/chessMap.gob"
	err := pickleDump(chessMap, filename)
	if err != nil {
		log.Fatal("pickleDump error:", err)
	}

	var localChessMap [11][11]int

	err = pickleLoad(filename, &localChessMap)
	if err != nil {
		log.Fatal("pickleLoad error:", err)
	}
	fmt.Println("localChessMap:", localChessMap)
}

func pickleDump(e interface{}, filename string) (err error) {
	var w bytes.Buffer
	encoder := gob.NewEncoder(&w)
	err = encoder.Encode(e)
	if err != nil {
		fmt.Println("encoder.Encode error:", err)
		return
	}
	err = ioutil.WriteFile(filename, w.Bytes(), 0644)
	if err != nil {
		fmt.Println("ioutil.WriteFile error:", err)
		return
	}
	return
}

func pickleLoad(filename string, e interface{}) (err error) {
	file, err := os.Open(filename)
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(e)
	if err != nil {
		fmt.Println("decoder.Decode error:", err)
		return
	}
	return
}
