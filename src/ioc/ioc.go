package ioc

import (
	"simple-api/src/app/core/baserepo"
	"simple-api/src/app/repositories/userrepo"
	"simple-api/src/app/schemas/usersche"
)

func InitUserRepo() *userrepo.Repository {
	return &userrepo.Repository{
		baserepo.RepositoryBase{
			usersche.Schema,
		},
		usersche.Schema,
	}
}
