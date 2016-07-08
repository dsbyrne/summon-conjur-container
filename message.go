package main

import (
  "encoding/binary"
  "bytes"
  "time"
)

type MessageType uint16;

type Message struct {
  Timestamp int64 
  Type MessageType
  Data *bytes.Buffer
}

func CreateMessage(msgType MessageType) *Message {
  return &Message {
    Timestamp: time.Now().UnixNano(),
    Type: msgType,
    Data: new(bytes.Buffer),
  }
}

func ParseMessage(data []byte) (message *Message, err error) {
  buffer := new(bytes.Buffer)

  header := [4]byte {}
  err = binary.Read(buffer, binary.BigEndian, &header)
  if err != nil {
    return nil, err
  }

  if bytes.Compare(header[:], MessageHeader) != 0 {
    return nil, &messageHeaderError{header}
  }

  binary.Read(buffer, binary.BigEndian, message.Timestamp)
  binary.Read(buffer, binary.BigEndian, message.Type)
  binary.Read(buffer, binary.BigEndian, message.Data)

  return message, nil
}

func (m Message) ToBytes() []byte {
  buffer := bytes.NewBuffer(MessageHeader)
  
  binary.Write(buffer, binary.BigEndian, m.Timestamp)
  binary.Write(buffer, binary.BigEndian, m.Type)
  buffer.Write(m.Data.Bytes())

  return buffer.Bytes()
}

func (m Message) AddInt64(val uint64) {
  binary.Write(m.Data, binary.BigEndian, val)
}

func (m Message) AddInt32(val uint32) {
  binary.Write(m.Data, binary.BigEndian, val)
}

func (m Message) AddInt16(val uint16) {
  binary.Write(m.Data, binary.BigEndian, val)
}

func (m Message) AddInt8(val uint8) {
  binary.Write(m.Data, binary.BigEndian, val)
}

func (m Message) AddBytes(val []byte) {
  binary.Write(m.Data, binary.BigEndian, val)
}