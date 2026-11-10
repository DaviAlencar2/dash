package server

import (
	"dash/message"
	"dash/server/buffer"
	"net"
)

type Client struct {
	Id      uint64
	UdpAddr net.UDPAddr
	TcpAddr net.TCPAddr
	Buffer  *buffer.Buffer[message.DataMessage]
}
