package main

import (
  "os"
)

func GetHostName() string {
  hostname := os.Getenv("HOSTNAME")
  if hostname == "" {
    panic("This provisioner be run from within a Docker container")
  }

  return hostname
}