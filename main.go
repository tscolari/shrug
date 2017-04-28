package main

import (
	"github.com/tscolari/shrug/garden"
	"github.com/tscolari/shrug/ui"
)

func main() {
	client := garden.NewClient()
	if err := ui.Start(client); err != nil {
		panic(err)
	}
}
