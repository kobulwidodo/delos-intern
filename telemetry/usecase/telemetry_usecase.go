package usecase

import "go-template/domain"

type TelemetryUsecase struct {
	telemetryRepository domain.TelemetryRepository
}

func NewTelemetryRepository(telemetryRepository domain.TelemetryRepository) domain.TelemetryUsecase {
	return &TelemetryUsecase{telemetryRepository}
}

func (u *TelemetryUsecase) Create(input domain.CreateTelemetryDto) error {
	telemetry := domain.Telemetry{
		Ip:       input.Ip,
		Method:   input.Method,
		Endpoint: input.Endpoint,
		Status:   input.Status,
		Latency:  input.Latency,
	}
	if err := u.telemetryRepository.Create(telemetry); err != nil {
		return err
	}
	return nil
}

func (u *TelemetryUsecase) GetAll() ([]domain.Telemetry, error) {
	telemetries, err := u.telemetryRepository.GetAll()
	if err != nil {
		return telemetries, err
	}
	if len(telemetries) == 0 {
		return telemetries, domain.ErrNotFound
	}
	return telemetries, nil
}

func (u *TelemetryUsecase) GetCount() ([]domain.CountResponse, error) {
	responses, err := u.telemetryRepository.GetCount()
	if err != nil {
		return responses, err
	}
	if len(responses) == 0 {
		return responses, domain.ErrNotFound
	}
	return responses, nil
}
