package main

import (
	"github.com/pavel-ovchinnikov/go-calculator/internal/application"
)

func main() {
	app := application.NewApplication()
	app.Run()
}
