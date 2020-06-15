package p2p

import (
	"encoding/json"
	"fmt"
)

const (
	PROTOCOL_NAME = 'simple_bitcoin_protocol'
	MY_VERSION = '0.1.0'
	MSG_ADD = 0
	MSG_REMOVE = 1
	MSG_CORE_LIST = 2
	MSG_REQUEST_CORE_LIST = 3
	MSG_PING = 4
	MSG_ADD_AS_EDGE = 5
	MSG_REMOVE_EDGE = 6
	ERR_PROTOCOL_UNMATCH = 0
	ERR_VERSION_UNMATCH = 1
	OK_WITH_PAYLOAD = 2
	OK_WITHOUT_PAYLOAD = 3
)

type Message struct {
	Protocol string `json:"protocl"`
	Version  string `json:"version"`
	MsgType  int    `json:"msg_type"`
	MyPort   string `json:"my_port"`
	Payload  string `json:"payload"`
}

func build(msgType int, payload=None) Message {
	return Message{
		Protocol: PROTOCOL_NAME,
		Version:  MY_VERSION,
		MsgType:  msgType,
	}
}

func parse() {

}
