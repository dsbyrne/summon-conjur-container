package main

import (
  "fmt"
)

func HandleConjurConnection(ps *ProvisionerServer, msg *Message) {
  fmt.Println("Hello from Conjur!")
}