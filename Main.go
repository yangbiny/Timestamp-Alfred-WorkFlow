package main

import (
	aw "github.com/deanishe/awgo"
)

func main() {
	println("xxx")
	workflow := aw.New()
	item := workflow.NewItem("时间")
	item.Arg("123456")
	workflow.SendFeedback()
}
