package yx5300

import (
	"github.com/tarm/serial"
	"io"
)

type Command struct {
}

type Connection interface {
	io.ReadWriteCloser
	SendCommand()
}

func MakeSerialConnection(devname string) Connection {

}
