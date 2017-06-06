package main

import (
	"github.com/fatih/color"
	"github.com/perriea/Golang-API/app"
)

var (
	yellow *color.Color
)

func init() {
	yellow = color.New(color.FgYellow, color.Bold)
}

func main() {
	yellow.Println("Golang API start ...")
	app.Router()
}
