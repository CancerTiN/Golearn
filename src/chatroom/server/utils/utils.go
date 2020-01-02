package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (t *Transfer) ReadPkg() (mes message.Message, err error) {
	// buf := make([]byte, 8096)
	fmt.Println("Read the data sent by the client ...")
	_, err = t.Conn.Read(t.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read error:", err)
		// err = errors.New("readPkg head error")
		return
	}

	var pkgLen uint32 = binary.BigEndian.Uint32(t.Buf[:4])
	n, err := t.Conn.Read(t.Buf[:pkgLen])
	if err != nil || n != int(pkgLen) {
		fmt.Println("conn.Read error:", err)
		err = errors.New("readPkg body error")
		return
	}

	err = json.Unmarshal(t.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}

	return
}

func (t *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(t.Buf[:4], pkgLen)
	n, err := t.Conn.Write(t.Buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write error:", err)
		return
	}

	n, err = t.Conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("conn.Write error:", err)
		return
	}

	return
}
