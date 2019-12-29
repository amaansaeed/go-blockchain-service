package main

import (
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	a := app{}
	a.Initialize()

	a.Run(":5000")
}
