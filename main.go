package main

import (
	"log"

	"github.com/Slava02/helperBot/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
