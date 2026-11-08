package message

import (
	"dash/consts"
	"encoding/binary"
)

// adicionar protobuf depois
type DataMessage struct {
	Operation Operation
	Topic     uint16
	Payload   string
}

func NewMessage(operation Operation, topic uint16, payload string) *DataMessage {
	return &DataMessage{
		Operation: operation,
		Topic:     topic,
		Payload:   payload,
	}
}

func Decode(bytes []byte) (*DataMessage, error) {
	operation, err := NewOperation(bytes[0])
	if err != nil {
		return nil, err
	}

	topic := binary.BigEndian.Uint16(bytes[1:3])

	return NewMessage(operation, topic, string(bytes[3:])), nil
}

func (d DataMessage) Encode() []byte {
	bytes := make([]byte, consts.MAX_DATA_MESSAGE_SIZE)

	bytes[0] = byte(d.Operation)
	binary.BigEndian.PutUint16(bytes[1:3], d.Topic)

	// capa no limite do slice bytes, já trunca
	copy(bytes[:3], []byte(d.Payload))

	return bytes
}
