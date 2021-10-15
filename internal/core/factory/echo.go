package factory

import (
	domain "github.com/gabrielgouv/go-openapi-driven-template/internal/domain/echo"
	repository "github.com/gabrielgouv/go-openapi-driven-template/internal/infrastructure/repository/echo"
	service "github.com/gabrielgouv/go-openapi-driven-template/internal/service/echo"
)

func MakeEchoService() service.EchoService {
	return service.NewEchoService(MakeEchoRepository())
}

func MakeEchoRepository() domain.EchoRepository {
	return repository.NewEchoMongoRepository()
}
