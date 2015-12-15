package main

import (
	"github.com/simulatedsimian/yx5300"
	"time"
)

func main() {
	conn, err := yx5300.MakeSerialConnection("/dev/ttyUSB0", false)
	if err != nil {
		panic(err)
	}

	conn.WriteCommand(yx5300.CMD_SEL_DEV, 0, yx5300.DEV_TF)

	conn.WriteCommand(yx5300.CMD_SEL_DEV, 0, yx5300.DEV_TF)
	time.Sleep(500 * time.Millisecond)

	conn.WriteCommand(yx5300.CMD_QUERY_FLDR_COUNT, 0, 0)
	time.Sleep(500 * time.Millisecond)

	conn.WriteCommand(yx5300.CMD_QUERY_TOT_TRACKS, 0, 0)
	time.Sleep(500 * time.Millisecond)

	for n := 1; n < 10; n++ {
		conn.WriteCommand(yx5300.CMD_QUERY_FLDR_TRACKS, 0, byte(n))
		time.Sleep(500 * time.Millisecond)
	}

	conn.WriteCommand(yx5300.CMD_PLAY_FOLDER_FILE, 4, 1)
	conn.WriteCommand(yx5300.CMD_PLAY_FOLDER_FILE, 4, 2)
	conn.WriteCommand(yx5300.CMD_PLAY_FOLDER_FILE, 4, 3)

	for {
		time.Sleep(5 * time.Second)
		//conn.WriteCommand(CMD_QUERY_STATUS, 0, 0)
	}
}
