package message

import (
	"encoding/binary"
	"errors"
)

// Ainda falta implementar o tamanho max para payload
type ControlMessage struct{
	Operation Operation
	Topic uint16
	Payload []byte // Debater sobre
}


func NewControlMessage(operation Operation, topic uint16, payload []byte) *ControlMessage{
	return &ControlMessage{
		Operation: operation,
		Topic: topic,
		Payload: payload,
	}
}


// [Tamanho: 4 bytes][Operação: 1 byte][Tópico: 2 bytes][Payload: possivelmente 1017]
func (c ControlMessage) Encode() []byte {
	bodySize := 3 + len(c.Payload)
	bytes := make([]byte, 4 + bodySize)

	binary.BigEndian.PutUint32(bytes[0:4], uint32(bodySize))

	bytes[4] = byte(c.Operation)
	binary.BigEndian.PutUint16(bytes[5:7], c.Topic)

	copy(bytes[7:], c.Payload)

	return bytes
}

// 
func DecodeControlMessage(bytes []byte) (*ControlMessage, error) {

	//[Tamanho: 4 bytes][Operação: 1 byte][Tópico: 2 bytes] = 7 bytes
	if len(bytes) < 7 {
		return nil, errors.New("mensagem de controle incompleta")
	}

	bodySize := binary.BigEndian.Uint32(bytes[0:4])

	if len(bytes) - 4 != int(bodySize) {
		return nil, errors.New("valor declarado e real de tamanho não são iguais")
	}

	operation, err := NewOperation(bytes[4])
	if err != nil {
		return nil, err
	}
	
	topic := binary.BigEndian.Uint16(bytes[5:7])

	return NewControlMessage(operation, topic, bytes[7:]), nil
}