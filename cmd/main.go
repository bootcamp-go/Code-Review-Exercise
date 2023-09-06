package main

import (
	"app/cmd/handlers"
	"app/internal/vehicle/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// env
	// ...

	// dependencies
	dbVh := make(map[int]*storage.VehicleAttributes)
	stVh := storage.NewStorageVehicleInMemory(dbVh)
	ctVh := handlers.NewControllerVehicle(stVh)

	// server
	rt := gin.New()
	// -> middlewares
	rt.Use(gin.Recovery())
	rt.Use(gin.Logger())
	// -> handlers
	api  := rt.Group("/api/v1")
	grVh := api.Group("/vehicles")
	grVh.GET("/", ctVh.GetAll())

	// run
	if err := rt.Run(":8080"); err != nil {
		panic(err)
	}
}