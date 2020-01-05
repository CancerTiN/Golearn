package main

import (
	"chatroom/common/message"
	"chatroom/server/process"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (p *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		fmt.Println("到了processor这里，信息为：", mes)
		up := &process.UserProcess{Conn: p.Conn}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		fmt.Println("到了processor这里，信息为：", mes)
		up := &process.UserProcess{Conn: p.Conn}
		err = up.ServerProcessRegister(mes)
	default:
		fmt.Println("Message type does not exist and cannot be processed ...")
	}
	return
}

func (p *Processor) mainProcess() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: p.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("The client has exited and the corresponding goroutine has also exited ...")
				return err
			} else {
				fmt.Println("readPkg error:", err)
				return err
			}
		}
		fmt.Printf("The message received was %v\n", mes)
		err = p.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes error:", err)
			return err
		}
	}
}
