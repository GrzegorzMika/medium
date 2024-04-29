package main

import (
	"golang.org/x/sync/singleflight"
)

func main() {
	app := &App{
		cache:   NewInMemoryStore(),
		group:   singleflight.Group{},
		manager: NewDummyModelManager(),
	}

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
