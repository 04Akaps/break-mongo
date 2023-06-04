package app

import (
	"break-mongo/network"
	"os"
	"os/signal"
)

type App struct {
	Router *network.Network
	stop   chan struct{}
}

func NewApp() *App {
	app := &App{
		stop: make(chan struct{}),
	}
	app.Router = network.NewNetwork()
	app.Router.Run(":8080")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		app.Stop()
	}()

	return app
}

func (a *App) Wait() {

	<-a.stop
}
func (a *App) Stop() {
	//a.AuctionMongoDB.Close()

	a.stop <- struct{}{}
}
