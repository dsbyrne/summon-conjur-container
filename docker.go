package main

import (
  "os"
  "io/ioutil"
  "strings"
)

func GetHostName() string {
  idPath := os.Getenv("CONTAINER_ID_PATH")
  if idPath != "" {
    buffer, err := ioutil.ReadFile(idPath)
    if err != nil {
      panic(err)
    }

    return strings.TrimSpace(string(buffer))
  }

  hostname := os.Getenv("HOSTNAME")
  if hostname == "" {
    panic("This provisioner be run from within a Docker container")
  }

  return hostname
}