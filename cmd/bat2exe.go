package main

import (
	"github.com/jamesmoriarty/gobat2exe"
	"os"
)

func main() {
	batPath := os.Args[1]

	gobat2exe.Bat2Exe(batPath)
}
