package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	filename := "D:/Workspace/Golearn/src/oldboy/day11/taildemo/my.log"
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	t, err := tail.TailFile(filename, config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("tail.TailFile(%#v, %#v) return: (%#v, %#v).\n", filename, config, t, err)

	var (
		line *tail.Line
		ok   bool
	)

	for {
		line, ok = <-t.Lines
		if !ok {
			fmt.Printf("%#v, %#v = <-%#v\n", line, ok, t.Lines)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line.Text:", line.Text)
	}
}
