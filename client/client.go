package client

import (
	"net"
)

/*
- Duvidas
	- Atualizar metadados vai ser tipo observer ?
	-
*/


// so pra nao aparecer erro
type broker struct{
	EnderecoTCP string
	EnderecoUDP string
}

type Client struct {
	Nome string
	
	// cache
	Topicos map[uint16]uint16
	Brokers map[uint16]broker
	Inscricoes map[uint16]bool
	Bootstrap broker

	ConexaoUDP net.UDPConn
	ConexaoTCP net.Conn
}

func NewClient(nome string, bootstrap broker) *Client {
	return &Client{
		Nome: nome,
		Bootstrap: bootstrap,
		Topicos: map[uint16]uint16{},
		Brokers: map[uint16]broker{},
		Inscricoes: map[uint16]bool{},
	}
}

func (c *Client) Conectar() error{
	conexao, err := net.Dial("tcp", c.Bootstrap.EnderecoTCP)
	if err != nil {
		return err
	}

	c.ConexaoTCP = conexao

	return nil
}

func (c Client) Inscrever() {
	
}

func (c Client) Publciar() {
	
}
