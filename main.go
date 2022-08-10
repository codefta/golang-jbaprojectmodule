package main

import (
	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
)

func main() {
	programmer := emoji.Sprint(":man_technologist:")
	emoji.Printf(":woman_technologist:")
	color.Green(" Learning about modules! " + programmer)
}
