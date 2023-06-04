package main

import (
	"break-mongo/cmd/app"
	"log"
)

func main() {

	log.Println("Server Started!!")

	app := app.NewApp()
	app.Wait()
}
