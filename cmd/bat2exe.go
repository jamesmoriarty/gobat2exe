package main

import (
	"github.com/jamesmoriarty/gobat2exe"
	"os"
)

func main() {
	batPath := os.Args[1]

	err := gobat2exe.Bat2Exe(batPath)

	if err != nil {
		panic(err)
	} else {
		os.Exit(0)
	}
}
