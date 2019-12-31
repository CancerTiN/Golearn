package main

import (
	"encoding/binary"
	"fmt"
	"strings"
)

func main() {
	var datStr = strings.Repeat("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 10086)
	var pkgLen = uint32(len(datStr))
	var bytes [4]byte
	fmt.Println("len(datStr) =", len(datStr))
	fmt.Println("bytes =", bytes)
	fmt.Println("bytes =", bytes[:])
	fmt.Printf("The pointer of &bytes is %p\n", &bytes)
	fmt.Printf("The pointer of bytes[:] is %p\n", bytes[:])
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	fmt.Println("bytes =", bytes)
}
