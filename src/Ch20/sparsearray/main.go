package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type ValNode struct {
	Row int
	Col int
	Val int
}

type SparseArchive struct {
	Shape     [2]int
	NodeSlice []ValNode
}

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

	var nodeSlice []ValNode
	var xLim, yLim int

	for i, row := range chessMap {
		for j, v := range row {
			if v != 0 {
				// valNode := ValNode{Row: i, Col: j, Val: v}
				nodeSlice = append(nodeSlice, ValNode{Row: i, Col: j, Val: v})
			}
			if j+1 > yLim {
				yLim = j + 1
			}
		}
		if i+1 > xLim {
			xLim = i + 1
		}
	}

	sparseArchive := SparseArchive{Shape: [2]int{xLim, yLim}, NodeSlice: nodeSlice}
	fmt.Println(sparseArchive)
	filename = "D:/Workspace/Golearn/src/Ch20/sparsearray/sparseArchive.gob"
	err = pickleDump(sparseArchive, filename)
	if err != nil {
		log.Fatal("pickleDump error:", err)
	}

	var localSparseArchive SparseArchive
	err = pickleLoad(filename, &localSparseArchive)
	if err != nil {
		log.Fatal("pickleLoad error:", err)
	}
	fmt.Println("localSparseArchive:", localSparseArchive)

	returnedChessMap := make([][]int, localSparseArchive.Shape[0])
	for i, row := range returnedChessMap {
		for j := 0; j <= localSparseArchive.Shape[1]; j++ {
			v := 0
			row = append(row, v)
			returnedChessMap[i] = row
		}
	}
	for _, valNode := range localSparseArchive.NodeSlice {
		returnedChessMap[valNode.Row][valNode.Col] = valNode.Val
	}

	fmt.Println("returnedChessMap:", returnedChessMap)
	for _, row := range returnedChessMap {
		for _, v := range row {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
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
