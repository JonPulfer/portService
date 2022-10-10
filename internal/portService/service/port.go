package service

import (
	"github.com/JonPulfer/portsService/internal/portService"
	"github.com/JonPulfer/portsService/internal/portService/repository"
)

type Port interface {
	Store(port portService.Port) error
	Find(portID string) (*portService.Port, error)
}

type port struct {
	portRepos repository.PortRepository
}

func NewPortService(repos repository.PortRepository) Port {
	return &port{portRepos: repos}
}

func (p *port) Store(port portService.Port) error {

	return nil
}

func (p *port) Find(portID string) (*portService.Port, error) {

	return nil, nil
}
