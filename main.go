package main

import (
	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	App := app.InitializeApp()
	App.Run()
}
