package repository

import (
	"git.jetbrains.space/keplerproject/kppoc/go-openapi-driven-template/internal/core/util/dateutil"
	"git.jetbrains.space/keplerproject/kppoc/go-openapi-driven-template/internal/domain/echo"
)

type EchoMongoRepository struct{}

var _ domain.EchoRepository = (*EchoMongoRepository)(nil)

func (p EchoMongoRepository) FindEchoDate() string {
	return dateutil.NowAsString()
}

func NewEchoMongoRepository() EchoMongoRepository {
	return EchoMongoRepository{}
}
