package userrepo

import (
	"simple-api/src/app/core/baserepo"
	"simple-api/src/app/schemas/usersche"
)

//Repository is
type Repository struct {
	baserepo.RepositoryBase
}

// Init is a
func (repo *Repository) Init() {
	repo.RepositoryBase.Init(usersche.Schema)
}
