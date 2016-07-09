package main

type ResponseHandler func(*ProvisionerServer, *Message)

type ResponseDispatcher struct {
  Handlers map[MessageType]ResponseHandler
}

func (d *ResponseDispatcher) RegisterHandler(msgType MessageType, handler ResponseHandler) {
  if d.Handlers == nil {
    d.Handlers = make(map[MessageType]ResponseHandler)
  }

  d.Handlers[msgType] = handler
}

func (d *ResponseDispatcher) HandlePacket(ps *ProvisionerServer, data []byte) {
  message, err := ParseMessage(data)
  if err != nil {
    return
  }

  handler, ok := d.Handlers[message.Type]
  if !ok {
    return
  }

  handler(ps, message)
}