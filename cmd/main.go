package main

import (
	"app/cmd/handlers"
	"app/internal/vehicle/repository"
	"app/internal/vehicle/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// env
	godotenv.Load(".env")

	// dependencies
	rpVh := repository.NewRepositoryVehicleJSONFile(os.Getenv("FILE_PATH_VEHICLES_JSON"))
	svVh := service.NewServiceVehicleDefault(rpVh)
	ctVh := handlers.NewControllerVehicle(svVh)

	// server
	rt := gin.New()
	// -> middlewares
	rt.Use(gin.Recovery())
	rt.Use(gin.Logger())
	// -> handlers
	api  := rt.Group("/api/v1")
	grVh := api.Group("/vehicles")
	grVh.GET("", ctVh.GetAll())

	// run
	if err := rt.Run(os.Getenv("SERVER_ADDR")); err != nil {
		panic(err)
	}
}