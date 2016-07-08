package main

import (
  "fmt"
)

type messageHeaderError struct {
  header [4]byte
}

func (e *messageHeaderError) Error() string {
  return fmt.Sprintf("Unrecognized message header: %s", e.header)
}