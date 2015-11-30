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

	conn.WriteCommand(CMD_SEL_DEV, DEV_TF, 0)
	time.Sleep(200 * time.Millisecond)

	conn.WriteCommand(CMD_PLAY_FOLDER_FILE, 1, 1)

	time.Sleep(200 * time.Second)
}
