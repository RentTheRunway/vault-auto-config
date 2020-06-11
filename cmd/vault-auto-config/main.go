package main

import (
	"fmt"
	"github.com/RentTheRunway/vault-auto-config/cmd/vault-auto-config/cmd"
	"github.com/juju/loggo"
)

type stdOutWriter struct {
}

// Override with unformatted output to stdout
func (s *stdOutWriter) Write(entry loggo.Entry) {
	fmt.Println(entry.Message)
}

func main() {
	_, _ = loggo.ReplaceDefaultWriter(&stdOutWriter{})
	cmd.Execute()
}
