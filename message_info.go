package main

var MessageHeader = []byte {'C', 'N', 'J', 'R'}

const (
  MsgType_RequestCallback = iota
  MsgType_ConjurConnection
)