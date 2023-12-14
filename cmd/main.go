package main

import (
	"app/internal/application"
	"fmt"
)

func main() {
	// env
	// ...

	// app
	// - config
	cfg := &application.ConfigDefaultInMemory{
		FileLoader: "./docs/db/vehicles_100.json",
		Addr:       ":8080",
	}
	// - app
	app := application.NewDefaultInMemory(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}