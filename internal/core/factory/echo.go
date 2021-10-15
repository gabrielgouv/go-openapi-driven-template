package factory

import (
	domain "git.jetbrains.space/keplerproject/kppoc/go-openapi-driven-template/internal/domain/echo"
	repository "git.jetbrains.space/keplerproject/kppoc/go-openapi-driven-template/internal/infrastructure/repository/echo"
	service "git.jetbrains.space/keplerproject/kppoc/go-openapi-driven-template/internal/service/echo"
)

func MakeEchoService() service.EchoService {
	return service.NewEchoService(MakeEchoRepository())
}

func MakeEchoRepository() domain.EchoRepository {
	return repository.NewEchoMongoRepository()
}
