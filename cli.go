package main

import (
	"fmt"
)

type User interface {
	RecordInput(interest string)
}

type CLI struct {
	user User
}

func (cli *CLI) RecordInput() {
	log.Print("What are you interested in today?")
}
