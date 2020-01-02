package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("Read the data sent by the server ...")
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read error:", err)
		// err = errors.New("readPkg head error")
		return
	}

	var pkgLen uint32 = binary.BigEndian.Uint32(buf[:4])
	n, err := conn.Read(buf[:pkgLen])
	if err != nil || n != int(pkgLen) {
		fmt.Println("conn.Read error:", err)
		err = errors.New("readPkg body error")
		return
	}

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}

	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	var pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write error:", err)
		return
	}

	n, err = conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("conn.Write error:", err)
		return
	}

	return
}
