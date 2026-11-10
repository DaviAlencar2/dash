package message

import "fmt"

type Operation uint8

const (
	SUB Operation = iota
	UNSUB
	ADD
	DEL
	UPD_META
	GET_META
	PING
	SEND
	FORWARDING
)

func NewOperation(value uint8) (Operation, error) {
	switch value {
	case 0:
		return SUB, nil
	case 1:
		return UNSUB, nil
	case 2:
		return ADD, nil
	case 3:
		return DEL, nil
	case 4:
		return UPD_META, nil
	case 5:
		return GET_META, nil
	case 6:
		return PING, nil
	case 7:
		return SEND, nil
	case 8:
		return FORWARDING, nil
	default:
		return 0, fmt.Errorf("invalid operation: %d", value)
	}
}
