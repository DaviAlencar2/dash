package client

import (
	"net"
)

/*
** Duvidas ** 
	- Atualizar metadados vai ser tipo observer ?
	-
*/


// so pra nao aparecer erro
type Broker int

type Client struct {
	Nome string
	
	// cache
	Topicos map[uint16]uint16
	Brokers map[uint16]Broker
	Inscricoes map[uint16]bool
	Bootstrap Broker

	ConexaoUDP *net.UDPConn
	ConexaoTCP *net.Conn
}

func (c Client) Conectar() {

}

func (c Client) Inscrever() {
	
}

func (c Client) Publciar() {
	
}
