package main

import (
  "net"
  "time"
  "fmt"
)

const (
  broadcastPort = 51000
)

type ProvisionerServer struct { 
  dispatch ResponseDispatcher
  port uint16
  socket *net.UDPConn
  running bool
}

func (ps *ProvisionerServer) initialize () {
  ps.dispatch.RegisterHandler(MsgType_ConjurConnection, HandleConjurConnection)
}

func (ps *ProvisionerServer) getInterfaceIP () (chosenAddr net.IP) {
  interfaces, err := net.Interfaces()
  if err != nil {
    panic(err)
  }

  for _, iface := range interfaces {
    isNotRunning := (iface.Flags & net.FlagUp == 0)
    isLoopback := (iface.Flags & net.FlagLoopback != 0)
    if isNotRunning || isLoopback {
      continue
    }

    addrs, err := iface.Addrs()
    if err != nil {
      panic(err)
    }

    for _, addr := range addrs {
      ip, _, err := net.ParseCIDR(addr.String())
      if err != nil {
        continue
      }

      if ip.To4() == nil {
        continue
      }

      return ip
    }
  }

  panic("Could not find an interface to listen on")
}

func (ps *ProvisionerServer) Listen() {
  if ps.running {
    return
  }

  ps.initialize()

  addr := net.UDPAddr{ IP: ps.getInterfaceIP() }

  socket, err := net.ListenUDP("udp4", &addr)
  if err != nil {
    panic("Failed to listen")
  }

  socketAddr, _ := net.ResolveUDPAddr(socket.LocalAddr().Network(), socket.LocalAddr().String())
  ps.port = uint16(socketAddr.Port)
  ps.socket = socket

  fmt.Printf("Listening on %d\n", ps.port)

  ps.waitForClose()
}

func (ps *ProvisionerServer) requestConjur() {
  msg := CreateMessage(MsgType_RequestCallback)
  msg.AddInt16(ps.port)
  msg.AddBytes([]byte(GetHostName()))

  ps.broadcast(msg)
}

func (ps *ProvisionerServer) Send(msg *Message) {
  if ps.socket == nil {
    panic("Not listening")
  }

  ps.socket.Write(msg.ToBytes())
}

func (ps *ProvisionerServer) broadcast(msg *Message) {
  socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
    IP:   net.IPv4(255, 255, 255, 255),
    Port: broadcastPort,
  })

  if err != nil {
    panic("Cannot broadcast")
  }

  _, err = socket.Write(msg.ToBytes())
  if err != nil {
    panic("Failed to write")
  }

  socket.Close()
}

func (ps *ProvisionerServer) waitForClose()  {
  ps.running = true

  ps.requestConjur()
  go ps.readSocket()

  for {
    if ps.socket == nil {
      break
    }

    time.Sleep(32 * time.Millisecond)
  }
}

func (ps *ProvisionerServer) Close() {
  ps.running = false

  if ps.socket == nil {
    return
  }

  ps.socket.Close()
  ps.socket = nil
}

func (ps *ProvisionerServer) readSocket() {
  if ps.socket == nil {
    panic("Not listening")
  }

  defer ps.Close()

  buffer := make([]byte, 10240)
  for ps.running {
    _, _, err := ps.socket.ReadFromUDP(buffer)
    if err != nil {
      panic("Failed to read")
    }

    ps.dispatch.HandlePacket(ps, buffer)
  }
}