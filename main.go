package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	fmt.Println("goreleaser test")
	if info, ok := debug.ReadBuildInfo(); ok {
		fmt.Println("Version", info.Main.Version)
	}
}
