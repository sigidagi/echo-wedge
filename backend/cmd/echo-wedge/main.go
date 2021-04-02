package main

import (
	"echo-wedge/backend/cmd/echo-wedge/cmd"
)

var version string

func main() {
	cmd.Execute(version)
}
