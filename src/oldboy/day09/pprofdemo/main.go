package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// A piece of problematic code.
func logicCode() {
	var nilCh chan int
	for {
		select {
		case v := <-nilCh:
			fmt.Printf("Received from chan, value: %v.\n", v)
		default:
			// Reduce CPU usage.
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("Create cpu pprof failed, error: %v.\n", err)
			return
		}
		_ = pprof.StartCPUProfile(file)
		defer func() {
			pprof.StopCPUProfile()
			_ = file.Close()
		}()
	}

	for i := 0; i < 8; i++ {
		go logicCode()
	}

	time.Sleep(20 * time.Second)

	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("Create mem pprof failed, error: %v.\n", err)
			return
		}
		_ = pprof.WriteHeapProfile(file)
		_ = file.Close()
	}
}
