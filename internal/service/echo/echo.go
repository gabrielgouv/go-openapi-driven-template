package service

import (
	"github.com/gabrielgouv/go-openapi-driven-template/internal/domain/echo"
	"math/rand"
)

type EchoService struct {
	Repository domain.EchoRepository
}

func (p *EchoService) GetEcho(message string) domain.Echo {
	randNumber := rand.Float32()
	date := p.Repository.FindEchoDate()
	response := domain.Echo{Date: date, Message: message, RandomNumber: randNumber}
	return response
}

func NewEchoService(repository domain.EchoRepository) EchoService {
	return EchoService{repository}
}
