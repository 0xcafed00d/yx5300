package main

import (
	//"fmt"
	"github.com/simulatedsimian/yx5300"
)

func main() {
	conn, err := yx5300.MakeSerialConnection("/dev/ttyUSB0", false)
	if err != nil {
		panic(err)
	}

	conn.WriteCommand(CMD_SEL_DEV, 0, DEV_TF)

}
