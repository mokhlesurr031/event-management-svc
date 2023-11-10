package main

import (
	"log"

	"github.com/mokhlesurr031/event-management-svc/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Println(err)
	}
}
