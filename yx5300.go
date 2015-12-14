package yx5300

import (
	"fmt"
	"github.com/tarm/serial"
	"io"
)

const (
	CMD_NEXT_SONG         = 0x01
	CMD_PREV_SONG         = 0x02
	CMD_PLAY_W_INDEX      = 0x03
	CMD_VOLUME_UP         = 0x04
	CMD_VOLUME_DOWN       = 0x05
	CMD_SET_VOLUME        = 0x06
	CMD_SINGLE_CYCLE_PLAY = 0x08
	CMD_SEL_DEV           = 0x09
	CMD_SLEEP_MODE        = 0x0A
	CMD_WAKE_UP           = 0x0B
	CMD_RESET             = 0x0C
	CMD_PLAY              = 0x0D
	CMD_PAUSE             = 0x0E
	CMD_PLAY_FOLDER_FILE  = 0x0F
	CMD_STOP_PLAY         = 0x16
	CMD_FOLDER_CYCLE      = 0x17
	CMD_SHUFFLE_PLAY      = 0x18
	CMD_SET_SINGLE_CYCLE  = 0x19
	CMD_SET_DAC           = 0x1A
	CMD_PLAY_W_VOL        = 0x22
	CMD_QUERY_STATUS      = 0x42
	CMD_QUERY_FLDR_TRACKS = 0x4e
	CMD_QUERY_TOT_TRACKS  = 0x48
	CMD_QUERY_FLDR_COUNT  = 0x4f

	DAC_ON           = 0x00
	DAC_OFF          = 0x01
	SINGLE_CYCLE_ON  = 0x00
	SINGLE_CYCLE_OFF = 0x01
	DEV_TF           = 0x02

	RESP_MEDIA_REMOVED     = 0x3b
	RESP_MEDIA_INSERTED    = 0x3a
	RESP_TF_TRACK_FINISHED = 0x3d
	RESP_ERROR             = 0x40
	RESP_ACK               = 0x41
	RESP_FLDR_TRACK_COUNT  = 0x4e
)

type Connection struct {
	comms        io.ReadWriteCloser
	ResponseChan chan Response
}

type Response struct {
	Code  int
	Param int
}

func (c *Connection) WriteCommand(cmd, arg1, arg2 byte) {
	var buffer [8]byte

	buffer[0] = 0x7e //starting byte
	buffer[1] = 0xff //version
	buffer[2] = 0x06 //the number of bytes of the command without starting byte and ending byte
	buffer[3] = cmd
	buffer[4] = 0x00 //0x00 = no feedback, 0x01 = feedback
	buffer[5] = arg1 //datah
	buffer[6] = arg2 //datal
	buffer[7] = 0xef //ending byte

	c.comms.Write(buffer[:])
}

func parseResponses(connection *Connection) {
	var buffer [1]byte
	var resp []byte

	for {
		_, err := connection.comms.Read(buffer[:])
		if err != nil {
			panic(err)
		}

		if buffer[0] == 0x7e {
			resp = nil
		}
		resp = append(resp, buffer[0])
		if buffer[0] == 0xef {
			for _, v := range resp {
				fmt.Printf("%02x ", v)
			}
			fmt.Println()
		}
	}
}

func MakeSerialConnection(devname string, debug bool) (*Connection, error) {
	port, err := serial.OpenPort(&serial.Config{Name: devname, Baud: 9600})

	connection := &Connection{
		comms:        port,
		ResponseChan: make(chan Response, 10),
	}

	if err == nil {
		go parseResponses(connection)
	}
	return connection, err
}
