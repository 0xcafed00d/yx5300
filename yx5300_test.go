package yx5300

import (
	"testing"
	"time"
)

func Test1(t *testing.T) {
	conn, err := MakeSerialConnection("/dev/ttyUSB0", false)
	if err != nil {
		t.Error(err)
		return
	}

	conn.WriteCommand(CMD_SEL_DEV, 0, DEV_TF)
	time.Sleep(500 * time.Millisecond)

	conn.WriteCommand(CMD_QUERY_FLDR_COUNT, 0, 0)
	time.Sleep(500 * time.Millisecond)

	conn.WriteCommand(CMD_QUERY_TOT_TRACKS, 0, 0)
	time.Sleep(500 * time.Millisecond)

	for n := 1; n < 10; n++ {
		conn.WriteCommand(CMD_QUERY_FLDR_TRACKS, 0, byte(n))
		time.Sleep(500 * time.Millisecond)
	}

	conn.WriteCommand(CMD_PLAY_FOLDER_FILE, 5, 2)

	for {
		time.Sleep(5 * time.Second)
		//conn.WriteCommand(CMD_QUERY_STATUS, 0, 0)
	}
}
