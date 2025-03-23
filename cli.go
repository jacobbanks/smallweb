package main

import (
	"fmt"
)


type User interface {
	RecordInput(interest string)
}

func CLI() {
	log.Print("What are you interested in today?")
}
