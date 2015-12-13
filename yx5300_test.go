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
	time.Sleep(200 * time.Millisecond)

	conn.WriteCommand(CMD_QUERY_FLDR_COUNT, 0, 0)
	time.Sleep(200 * time.Millisecond)

	conn.WriteCommand(CMD_QUERY_FLDR_TRACKS, 1, 1)
	time.Sleep(200 * time.Millisecond)
	conn.WriteCommand(CMD_QUERY_FLDR_TRACKS, 2, 2)
	time.Sleep(200 * time.Millisecond)
	conn.WriteCommand(CMD_QUERY_FLDR_TRACKS, 3, 3)
	time.Sleep(200 * time.Millisecond)

	conn.WriteCommand(CMD_FOLDER_CYCLE, 2, 0)

	for {
		time.Sleep(5 * time.Second)
		//conn.WriteCommand(CMD_QUERY_STATUS, 0, 0)
	}
}
