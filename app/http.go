package main

import (
	"fmt"
	_farmHttpHandler "go-template/farm/delivery/http"
	_farmRepository "go-template/farm/repository/postgresql"
	_farmUsecase "go-template/farm/usecase"
	"go-template/infrastructure"
	"go-template/middleware"
	_pondHttpHandler "go-template/pond/delivery/http"
	_pondRepository "go-template/pond/repository/postgresql"
	_pondUsecase "go-template/pond/usecase"
	_telemetryHttpHandler "go-template/telemetry/delivery/http"
	_telemetryRepository "go-template/telemetry/repository/postgresql"
	_telemetryUsecase "go-template/telemetry/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("failed to load env")
	}

	dbDriver := infrastructure.NewDbConfig()

	db, err := dbDriver.InitDb()
	if err != nil {
		log.Fatal("failed to connect with database")
		panic(err)
	}
	r := gin.Default()
	api := r.Group("/api")
	telemetryRepository := _telemetryRepository.NewTelemetryRepository(db)
	telemetryUsecase := _telemetryUsecase.NewTelemetryRepository(telemetryRepository)
	telemetryMiddleware := middleware.NewTelemetryMiddleware(telemetryUsecase)

	api.Use(telemetryMiddleware.Telemetry())

	farmRepository := _farmRepository.NewFarmRepository(db)
	pondRepository := _pondRepository.NewPondRepository(db)

	farmUsecase := _farmUsecase.NewFarmUsecase(farmRepository)
	pondUsecase := _pondUsecase.NewPondUsecase(pondRepository)

	_farmHttpHandler.NewFarmHandler(api, farmUsecase)
	_pondHttpHandler.NewPondHandler(api, pondUsecase)
	_telemetryHttpHandler.NewTelemetryHandler(api, telemetryUsecase)

	r.Run()
}
