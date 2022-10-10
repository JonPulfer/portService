package repository

import "github.com/JonPulfer/portsService/internal/portService"

//go:generate mockgen -destination=mocks/repository.go -package=mocks . PortRepository
type PortRepository interface {
	Store(port portService.Port) error
	Fetch(portID string) (portService.Port, error)
	Update(port portService.Port) error
}
