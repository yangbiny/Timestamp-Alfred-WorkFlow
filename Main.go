package main

import (
	aw "github.com/deanishe/awgo"
	"os"
)

func main() {
	args := os.Args
	println(args)
	workflow := aw.New()
	println(workflow.Args())
}
