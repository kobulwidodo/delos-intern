package domain

import "gorm.io/gorm"

type Telemetry struct {
	gorm.Model
	Ip       string
	Method   string
	Endpoint string
	Status   int
	Latency  int64
}

type TelemetryRepository interface {
	Create(telemetry Telemetry) error
	GetAll() ([]Telemetry, error)
	GetCount() ([]CountResponse, error)
}

type TelemetryUsecase interface {
	Create(input CreateTelemetryDto) error
	GetAll() ([]Telemetry, error)
	GetCount() ([]CountResponse, error)
}

type CreateTelemetryDto struct {
	Ip       string
	Method   string
	Endpoint string
	Status   int
	Latency  int64
}

type CountResponse struct {
	Endpoint   string `json:"endpoint"`
	Count      int    `json:"count"`
	UniqueUser int    `json:"unique_user"`
}
