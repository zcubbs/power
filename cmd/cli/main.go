package main

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

var (
	Color string
)

func main() {

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select a color").
				Options(
					huh.NewOption("red", "Red"),
					huh.NewOption("green", "Green"),
				).
				Value(&Color),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Selected color", "color", Color)
}
